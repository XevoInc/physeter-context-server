package ranker

import (
	"fmt"
	"time"

	pbContext "xevo/physeter-context-server/proto"
	"xevo/physeter-context-server/server/types"

	"github.com/najeira/jpholiday"
)

var jst = time.FixedZone("JST", 3600*9)

type FeatureVector struct {
}

var _ types.Ranker = (*FeatureVector)(nil)

func (fv *FeatureVector) Name() string {
	return "FeatureVector"
}

func (fv *FeatureVector) Rank(req *pbContext.GetRecommendsRequest, pois []*pbContext.PointOfInterest) (*pbContext.GetRecommendsResponse, error) {
	err := validateRequest(req)
	if err != nil {
		return nil, err
	}

	category := getRecommendCategory(req)
	if category == pbContext.UNKNOWN {
		return &pbContext.GetRecommendsResponse{
			Recommends: nil,
		}, nil
	}
	return &pbContext.GetRecommendsResponse{
		Recommends: filterPois(pois, category),
	}, nil
}
func getRecommendCategory(req *pbContext.GetRecommendsRequest) pbContext.PointOfInterest_Category {
	switch {
	case isGasIsLow(req):
		return pbContext.GAS_STATION
	case isInTheMorningAndExistPassenger(req):
		return pbContext.SCHOOL
	case isInTheMorning(req):
		return pbContext.CAFE
	case isAtLunchTime(req):
		return pbContext.RESTAURANT
	case isInTheEvening(req):
		return pbContext.GROCERY
	}
	return pbContext.UNKNOWN
}

func validateRequest(req *pbContext.GetRecommendsRequest) error {
	if req == nil {
		return fmt.Errorf("Request object must not be nil")
	}
	if req.Time == nil {
		return fmt.Errorf("The time is required")
	}
	if req.CarState == nil {
		return fmt.Errorf("The car state is required")
	}
	return nil
}

func filterPois(pois []*pbContext.PointOfInterest, category pbContext.PointOfInterest_Category) (ret []*pbContext.PointOfInterest) {
	for _, poi := range pois {
		if contains(poi, category) {
			ret = append(ret, poi)
		}
	}
	return
}
func contains(poi *pbContext.PointOfInterest, category pbContext.PointOfInterest_Category) bool {
	for _, c := range poi.Categories {
		if c == category {
			return true
		}
	}
	return false
}

func getTime(req *pbContext.GetRecommendsRequest) time.Time {
	return time.Unix(req.Time.GetSeconds(), int64(req.Time.GetNanos())).In(jst)
}
func isWorkday(t time.Time) bool {
	weekday := t.Weekday()
	if weekday == time.Sunday || weekday == time.Saturday {
		return false
	}
	name := jpholiday.Name(t)
	return len(name) == 0
}

func isGasIsLow(req *pbContext.GetRecommendsRequest) bool {
	return req.CarState.FuelLevelPercentage <= 20.0
}
func isInTheMorning(req *pbContext.GetRecommendsRequest) bool {
	t := getTime(req)
	h := t.Hour()
	return 6 <= h && h <= 9
}
func isInTheMorningAndExistPassenger(req *pbContext.GetRecommendsRequest) bool {
	return isInTheMorning(req) && req.CarState.NumberOfPassengers > 1
}
func isAtLunchTime(req *pbContext.GetRecommendsRequest) bool {
	t := getTime(req)
	h := t.Hour()
	return 12 <= h && h <= 13
}
func isInTheEvening(req *pbContext.GetRecommendsRequest) bool {
	t := getTime(req)
	h := t.Hour()
	return 18 <= h && h <= 21
}
