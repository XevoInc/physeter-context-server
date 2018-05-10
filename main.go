package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"mime"
	"net"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pbContext "xevo/physeter-context-server/proto"
	"xevo/physeter-context-server/server"
	// Static files
	_ "xevo/physeter-context-server/statik"
)

var log grpclog.LoggerV2

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(log)
}

func startGrpcServer(addr string, cert *tls.Certificate) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("Failed to listen: %s", err)
	}
	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(cert)),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
	)
	pbContext.RegisterContextServiceServer(s, server.New())

	log.Info("Serving gRPC on https://", addr)
	log.Fatalln(s.Serve(lis))

	return nil
}

func startGatewayServer(addr string, gatewayAddr string, cp *x509.CertPool, cert *tls.Certificate) error {
	gwmux, err := createGrpcGateway(addr, cp)
	if err != nil {
		return err
	}

	statik, err := createStaticFS()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", gwmux)
	mux.Handle("/", http.StripPrefix("/", statik))

	log.Info("Serving gRPC-Gateway on https://", gatewayAddr)
	log.Info("Serving OpenAPI Documentation on https://", gatewayAddr, "/openapi-ui/")
	gwServer := http.Server{
		Addr: gatewayAddr,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*cert},
		},
		Handler: allowCORS(mux),
	}
	log.Fatalln(gwServer.ListenAndServeTLS("", ""))

	return nil
}

func createGrpcGateway(addr string, cp *x509.CertPool) (*runtime.ServeMux, error) {
	// See https://github.com/grpc/grpc/blob/master/doc/naming.md
	// for gRPC naming standard information.
	dialAddr := fmt.Sprintf("passthrough://localhost/%s", addr)
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(cp, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial server: %s", err)
	}

	jsonpb := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		// This is necessary to get error details properly
		// marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	err = pbContext.RegisterContextServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return nil, fmt.Errorf("Failed to register gateway: %s", err)
	}
	return gwmux, nil
}

func createStaticFS() (http.Handler, error) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("Failed to serve static files: %s", err)
	}
	return http.FileServer(statikFS), nil
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Infof("preflight request for %s", r.URL.Path)
	return
}

func loadCertificate(host string) (*tls.Certificate, *x509.CertPool, error) {
	cert, err := tls.LoadX509KeyPair(path.Join("./certificates", host+".crt"), path.Join("./certificates", host+".key"))
	if err != nil {
		return nil, nil, err
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, nil, err
	}
	cp := x509.NewCertPool()
	cp.AddCert(cert.Leaf)
	return &cert, cp, nil
}

func main() {
	var (
		gRPCHost    string
		gRPCPort    int
		gatewayHost string
		gatewayPort int
	)

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}

	app := cli.NewApp()
	app.Name = "context-server"
	app.Version = "0.1.0"
	app.Usage = "start GRPC server or GRPC gateway"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "grpc-host, h",
			Value:       "localhost",
			Usage:       "The gRPC server `host`",
			EnvVar:      "GRPC_HOST",
			Destination: &gRPCHost,
		},
		cli.IntFlag{
			Name:        "grpc-port, p",
			Value:       10000,
			Usage:       "The gRPC server `port`",
			EnvVar:      "GRPC_PORT",
			Destination: &gRPCPort,
		},
		cli.StringFlag{
			Name:        "gateway-host, H",
			Value:       "localhost",
			Usage:       "The gRPC-Gateway server `host`",
			EnvVar:      "GATEWAY_HOST",
			Destination: &gatewayHost,
		},
		cli.IntFlag{
			Name:        "gateway-port, P",
			Value:       11000,
			Usage:       "The gRPC-Gateway server `port`",
			EnvVar:      "GATEWAY_PORT",
			Destination: &gatewayPort,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "grpc-server",
			Aliases: []string{"grpc"},
			Usage:   "start gRPC server",
			Action: func(c *cli.Context) error {
				cert, _, err := loadCertificate(gRPCHost)
				if err != nil {
					return err
				}
				addr := fmt.Sprintf("%s:%d", gRPCHost, gRPCPort)
				err = startGrpcServer(addr, cert)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "gateway-server",
			Aliases: []string{"gateway"},
			Usage:   "start gRPC-Gateway server",
			Action: func(c *cli.Context) error {
				cert, cp, err := loadCertificate(gRPCHost)
				if err != nil {
					return err
				}
				addr := fmt.Sprintf("%s:%d", gRPCHost, gRPCPort)
				gatewayAddr := fmt.Sprintf("%s:%d", gatewayHost, gatewayPort)
				err = startGatewayServer(addr, gatewayAddr, cp, cert)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
