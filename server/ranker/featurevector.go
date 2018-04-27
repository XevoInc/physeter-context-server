package ranker

import (
	pbContext "xevo/physeter-context-server/proto"
	"xevo/physeter-context-server/server/types"
)

type FeatureVector struct {
}

var _ types.Ranker = (*FeatureVector)(nil)

func (fv *FeatureVector) Name() string {
	return "FeatureVector"
}

func (fv *FeatureVector) Rank(pois []*pbContext.PointOfInterest) (*pbContext.GetRecommendsResponse, error) {
	return &pbContext.GetRecommendsResponse{
		Recommends: pois,
	}, nil
}
