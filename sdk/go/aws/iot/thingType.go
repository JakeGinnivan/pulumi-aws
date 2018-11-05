// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and manages an AWS IoT Thing Type.
type ThingType struct {
	s *pulumi.ResourceState
}

// NewThingType registers a new resource with the given unique name, arguments, and options.
func NewThingType(ctx *pulumi.Context,
	name string, args *ThingTypeArgs, opts ...pulumi.ResourceOpt) (*ThingType, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["deprecated"] = nil
		inputs["name"] = nil
		inputs["properties"] = nil
	} else {
		inputs["deprecated"] = args.Deprecated
		inputs["name"] = args.Name
		inputs["properties"] = args.Properties
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:iot/thingType:ThingType", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ThingType{s: s}, nil
}

// GetThingType gets an existing ThingType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThingType(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ThingTypeState, opts ...pulumi.ResourceOpt) (*ThingType, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["deprecated"] = state.Deprecated
		inputs["name"] = state.Name
		inputs["properties"] = state.Properties
	}
	s, err := ctx.ReadResource("aws:iot/thingType:ThingType", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ThingType{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ThingType) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ThingType) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the created AWS IoT Thing Type.
func (r *ThingType) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Whether the thing type is deprecated. If true, no new things could be associated with this type.
func (r *ThingType) Deprecated() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["deprecated"])
}

// The name of the thing type.
func (r *ThingType) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *ThingType) Properties() *pulumi.Output {
	return r.s.State["properties"]
}

// Input properties used for looking up and filtering ThingType resources.
type ThingTypeState struct {
	// The ARN of the created AWS IoT Thing Type.
	Arn interface{}
	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated interface{}
	// The name of the thing type.
	Name interface{}
	Properties interface{}
}

// The set of arguments for constructing a ThingType resource.
type ThingTypeArgs struct {
	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated interface{}
	// The name of the thing type.
	Name interface{}
	Properties interface{}
}
