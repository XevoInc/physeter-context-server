package insecure

import (
	"crypto/tls"
	"crypto/x509"
	"log"
)

const certPEM = `-----BEGIN CERTIFICATE-----
MIIDgjCCAmoCCQDinB7w9PVs3DANBgkqhkiG9w0BAQUFADCBgjELMAkGA1UEBhMC
SlAxDjAMBgNVBAgMBVRva3lvMRgwFgYDVQQHDA9TaGlidXlhIEhpZ2FzaGkxEjAQ
BgNVBAoMCVhldm8gSy5LLjESMBAGA1UEAwwJbG9jYWxob3N0MSEwHwYJKoZIhvcN
AQkBFhJoaG9yaXVjaGlAeGV2by5jb20wHhcNMTgwNDI1MDkwMDAxWhcNMjgwNDIy
MDkwMDAxWjCBgjELMAkGA1UEBhMCSlAxDjAMBgNVBAgMBVRva3lvMRgwFgYDVQQH
DA9TaGlidXlhIEhpZ2FzaGkxEjAQBgNVBAoMCVhldm8gSy5LLjESMBAGA1UEAwwJ
bG9jYWxob3N0MSEwHwYJKoZIhvcNAQkBFhJoaG9yaXVjaGlAeGV2by5jb20wggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDpLFjTZwbIwKP133ixxwYykPh8
HFXutj0OzUZfUXI89OMwkSQnJkkmFI7LlDvw/5rWw7S6useSaf9pdJ+zbWcekcgS
ySLLIiVMfJSdYJ6DbrSfD59iZimJsFKa0FokIGUT7WLnmW0zc8J3O1G+h7VfoqZ+
73QmnOrUgGDTw1U1+pNOt1MLtO1Wbfdi1Kz/fC1hx0vjZ+oufwE6IC44B5pIjBmr
rg1BTb5c2+9fcKrF6/u+bQLOpopiaodGGUcEsqENUZ5x/V12SdpEZHJzTqSHhMZL
+0kHQZLGB25bKqsXuJZw1D7I7yR+i8R4srn3ELnY6QzPpwJkseSQH55tZ4LhAgMB
AAEwDQYJKoZIhvcNAQEFBQADggEBAJWED3oal9fDZzPvYJlzyMcXMf1tB6fGJchV
hzY/AcFvjeD8RuueY3BtNsfc0cMnGJ5GpVWv7CRK+RtYp0EprRnYMRDH8z+AH8EG
7bT5TpQjgk0Y47lPXIzhu6ETtfpEb5RIlgbXphKOnanLmDBoL4A4s77IrZLnnSb9
RSAvHEylwuWXOb0KJ8ZxQZoktSR9B1lU2+MCiORRig0Nel2zxVGY6AKaCv0yB87c
cy/ZdfvbTs0bqxxKLthhCFiSGVPUExd9JuQ9/W+rp2KasN36yCkj9jSPFqyR+nsf
yvVuLQokzax70kf0dz45Wb2ucDlheSQ5Zrwf1KCdWslVkx3feQg=
-----END CERTIFICATE-----
`

const keyPEM = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA6SxY02cGyMCj9d94sccGMpD4fBxV7rY9Ds1GX1FyPPTjMJEk
JyZJJhSOy5Q78P+a1sO0urrHkmn/aXSfs21nHpHIEskiyyIlTHyUnWCeg260nw+f
YmYpibBSmtBaJCBlE+1i55ltM3PCdztRvoe1X6Kmfu90Jpzq1IBg08NVNfqTTrdT
C7TtVm33YtSs/3wtYcdL42fqLn8BOiAuOAeaSIwZq64NQU2+XNvvX3Cqxev7vm0C
zqaKYmqHRhlHBLKhDVGecf1ddknaRGRyc06kh4TGS/tJB0GSxgduWyqrF7iWcNQ+
yO8kfovEeLK59xC52OkMz6cCZLHkkB+ebWeC4QIDAQABAoIBAQDAIDpQdmOwpopy
3V+lnEgPAkS8ftyPhLlqqNmDUKjAcUeBjRYs6+754ZaHR+zb9sCulTMlaSFoEs7R
JbcxMMP2/EfOssxodAonq78IrcPuqBlLtqOtgGBZg1+NEIA1pDuk38TO8Cx6aLej
Zs1EIK3hzzNBBCzAZCtaTS17T0e24N3gpq5NWCuq3hw4eZTxWV170kl7QV6GFybu
88DEwB73NUYdBR6E4qe7udgnKlQVR8z0GrHgQurGKxVIwzyLin745EwFFwQpPoIt
lCJtZIdKMi9FWRJEqzN1OSx/D/0by7yYMgAM6XgNjFQpsFYEJGHDxUf3kFyKulp+
9H8zD3QBAoGBAP8DbVmRHA4zZkjh2QNMrjTbJo6qkfFFrRBuakUNRYUa9eoZHkRi
Q+bO5zjZjOCTKAgxn7d/ACq2H3ZGcowxlwn64Uut/VbNUlZlxgEwiYqLRYV5jDrj
13E8qsq3/dJbyu/5l3vpC6o0yvIjrcW9dLUxpa8Tibez2oKcjdwVnEfJAoGBAOoT
SedUytC16J/ouIvWPEPSexdr6JXjLh1eIT4Urlf4A+Fbvk2QhrXO4KtxoCcK9nVG
rbYia2JvvFWzwlCP3qGZcIcPFDwO+ZW2nXTRkOSq45RMl3K6hgGGa2wTJjCXLJFI
0pr7TUR50pAVvNH+n44iVlltHPgJGacCI0By4R5ZAoGAOMWMjIoPkuHCfas54tAp
YD8BoHU8tFfspWHParv/pBhScuFQqayurS2WsENOZ5ibKufRRMqTQO72piAc1DUe
COy/R1fetHpVQEK8B/vEQwGqYoKiw2GBzwXQh/zaFwC0dyQ7oyxImoXSQIHM5967
orMAza8t06XImZc8xahR+HECgYBT5P/q8qASd/BdR+rE688Du++0ME/XUOpZSkB9
4KmwC8tQPTKS2Eb/6Jrrt7jf5XT6CY//JIz6ZfPJf7kYaMgxgS4sqgHlmLRprV73
3Jii7IDOyTjKvQGEkGn1/VMFvVkS5Vfehk1mSodmbvQTC8CiD7qFKK9mUtSn4ANE
eSgoWQKBgG4rhJ4eNggyt2yb50mffcICsf7olvHEH/R387l1oo253fXGZjlHN5DZ
xhzIdWya8IhcjHnRtB2Cj/74hX5K5aVDYmlCvdua+IKkNx7LvIhwjHt6RIShXjXv
zOcFX+bBEUw2vkDjq61k2E0FJw3jhfE5DDgMsNgAtzuITUryJE1J
-----END RSA PRIVATE KEY-----
`

var (
	// Cert is a self signed certificate
	Cert tls.Certificate
	// CertPool contains the self signed certificate
	CertPool *x509.CertPool
)

func init() {
	var err error
	Cert, err = tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		log.Fatalln("Failed to parse key pair:", err)
	}
	Cert.Leaf, err = x509.ParseCertificate(Cert.Certificate[0])
	if err != nil {
		log.Fatalln("Failed to parse certificate:", err)
	}

	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)
}
