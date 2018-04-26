package server

import (
	"context"
	"sync"

	pbContext "xevo/physeter-context-server/proto"
)

type Backend struct {
	mu *sync.RWMutex
}

var _ pbContext.ContextServiceServer = (*Backend)(nil)

func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

func (b *Backend) GetRecommends(c context.Context, req *pbContext.GetRecommendsRequest) (*pbContext.GetRecommendsResponse, error) {
	return nil, nil
}
