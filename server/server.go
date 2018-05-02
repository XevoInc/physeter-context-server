package server

import (
	"context"
	"xevo/physeter-context-server/server/poicollector"
	"xevo/physeter-context-server/server/ranker"
	"xevo/physeter-context-server/server/types"

	pbContext "xevo/physeter-context-server/proto"
)

const (
	defaultLimit = 10
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
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	limit := req.Limit
	if limit <= 0 || 100 < limit {
		limit = defaultLimit
	}

	pois, err := s.c.Collect(req)
	if err != nil {
		return nil, err
	}
	sorted, err := s.r.Rank(req, pois)
	if err != nil {
		return nil, err
	}

	result := sorted.Recommends
	sorted.Recommends = result[:min(int(limit), len(result))]
	return sorted, nil
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
