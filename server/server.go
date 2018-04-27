package server

import (
	"context"
	"xevo/physeter-context-server/server/poicollector"
	"xevo/physeter-context-server/server/ranker"
	"xevo/physeter-context-server/server/types"

	pbContext "xevo/physeter-context-server/proto"
)

type Server struct {
	c types.PoiCollector
	r types.Ranker
}

var _ pbContext.ContextServiceServer = (*Server)(nil)

func New() *Server {
	return &Server{
		c: &poicollector.ZdcCollector{},
		r: &ranker.FeatureVector{},
	}
}

func (s *Server) GetRecommends(c context.Context, req *pbContext.GetRecommendsRequest) (*pbContext.GetRecommendsResponse, error) {
	pois, err := s.c.Collect(req)
	if err != nil {
		return nil, err
	}
	return s.r.Rank(pois)
}
