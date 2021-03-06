// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: context.proto

/*
Package context_v1 is a generated protocol buffer package.

It is generated from these files:
	context.proto

It has these top-level messages:
	GetRecommendsRequest
	GetRecommendsResponse
	UserState
	CarState
	Coordinates
	GeoCoordinates
	PointOfInterest
*/
package context_v1

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetRecommendsRequest) Validate() error {
	if this.UserState != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.UserState); err != nil {
			return go_proto_validators.FieldError("UserState", err)
		}
	}
	if nil == this.CarState {
		return go_proto_validators.FieldError("CarState", fmt.Errorf("message must exist"))
	}
	if this.CarState != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CarState); err != nil {
			return go_proto_validators.FieldError("CarState", err)
		}
	}
	if nil == this.Time {
		return go_proto_validators.FieldError("Time", fmt.Errorf("message must exist"))
	}
	if this.Time != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Time); err != nil {
			return go_proto_validators.FieldError("Time", err)
		}
	}
	return nil
}
func (this *GetRecommendsResponse) Validate() error {
	for _, item := range this.Recommends {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Recommends", err)
			}
		}
	}
	return nil
}
func (this *UserState) Validate() error {
	return nil
}
func (this *CarState) Validate() error {
	if nil == this.CurrentLocation {
		return go_proto_validators.FieldError("CurrentLocation", fmt.Errorf("message must exist"))
	}
	if this.CurrentLocation != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CurrentLocation); err != nil {
			return go_proto_validators.FieldError("CurrentLocation", err)
		}
	}
	if this.Destination != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Destination); err != nil {
			return go_proto_validators.FieldError("Destination", err)
		}
	}
	if !(this.FuelLevelPercentage >= 0) {
		return go_proto_validators.FieldError("FuelLevelPercentage", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.FuelLevelPercentage))
	}
	if !(this.FuelLevelPercentage <= 100) {
		return go_proto_validators.FieldError("FuelLevelPercentage", fmt.Errorf(`value '%v' must be lower than or equal to '100'`, this.FuelLevelPercentage))
	}
	if !(this.RangeM >= 0) {
		return go_proto_validators.FieldError("RangeM", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.RangeM))
	}
	return nil
}
func (this *Coordinates) Validate() error {
	return nil
}
func (this *GeoCoordinates) Validate() error {
	if nil == this.Location {
		return go_proto_validators.FieldError("Location", fmt.Errorf("message must exist"))
	}
	if this.Location != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Location); err != nil {
			return go_proto_validators.FieldError("Location", err)
		}
	}
	return nil
}
func (this *PointOfInterest) Validate() error {
	if this.Coordinates != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Coordinates); err != nil {
			return go_proto_validators.FieldError("Coordinates", err)
		}
	}
	return nil
}
