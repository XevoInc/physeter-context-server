package types

import (
	pbContext "xevo/physeter-context-server/proto"
)

// PoiCollector is the collector of external pois.
type PoiCollector interface {
	Name() string
	Collect(req *pbContext.GetRecommendsRequest) ([]*pbContext.PointOfInterest, error)
}

// Ranker is the ranker of collected pois.
type Ranker interface {
	Name() string
	Rank(req *pbContext.GetRecommendsRequest, pois []*pbContext.PointOfInterest) (*pbContext.GetRecommendsResponse, error)
}
