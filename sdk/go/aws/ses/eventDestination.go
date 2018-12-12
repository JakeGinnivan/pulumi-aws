// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SES event destination
type EventDestination struct {
	s *pulumi.ResourceState
}

// NewEventDestination registers a new resource with the given unique name, arguments, and options.
func NewEventDestination(ctx *pulumi.Context,
	name string, args *EventDestinationArgs, opts ...pulumi.ResourceOpt) (*EventDestination, error) {
	if args == nil || args.ConfigurationSetName == nil {
		return nil, errors.New("missing required argument 'ConfigurationSetName'")
	}
	if args == nil || args.MatchingTypes == nil {
		return nil, errors.New("missing required argument 'MatchingTypes'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cloudwatchDestinations"] = nil
		inputs["configurationSetName"] = nil
		inputs["enabled"] = nil
		inputs["kinesisDestination"] = nil
		inputs["matchingTypes"] = nil
		inputs["name"] = nil
		inputs["snsDestination"] = nil
	} else {
		inputs["cloudwatchDestinations"] = args.CloudwatchDestinations
		inputs["configurationSetName"] = args.ConfigurationSetName
		inputs["enabled"] = args.Enabled
		inputs["kinesisDestination"] = args.KinesisDestination
		inputs["matchingTypes"] = args.MatchingTypes
		inputs["name"] = args.Name
		inputs["snsDestination"] = args.SnsDestination
	}
	s, err := ctx.RegisterResource("aws:ses/eventDestination:EventDestination", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventDestination{s: s}, nil
}

// GetEventDestination gets an existing EventDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventDestination(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EventDestinationState, opts ...pulumi.ResourceOpt) (*EventDestination, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cloudwatchDestinations"] = state.CloudwatchDestinations
		inputs["configurationSetName"] = state.ConfigurationSetName
		inputs["enabled"] = state.Enabled
		inputs["kinesisDestination"] = state.KinesisDestination
		inputs["matchingTypes"] = state.MatchingTypes
		inputs["name"] = state.Name
		inputs["snsDestination"] = state.SnsDestination
	}
	s, err := ctx.ReadResource("aws:ses/eventDestination:EventDestination", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventDestination{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EventDestination) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EventDestination) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// CloudWatch destination for the events
func (r *EventDestination) CloudwatchDestinations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["cloudwatchDestinations"])
}

// The name of the configuration set
func (r *EventDestination) ConfigurationSetName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["configurationSetName"])
}

// If true, the event destination will be enabled
func (r *EventDestination) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// Send the events to a kinesis firehose destination
func (r *EventDestination) KinesisDestination() *pulumi.Output {
	return r.s.State["kinesisDestination"]
}

// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
func (r *EventDestination) MatchingTypes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["matchingTypes"])
}

// The name of the event destination
func (r *EventDestination) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Send the events to an SNS Topic destination
func (r *EventDestination) SnsDestination() *pulumi.Output {
	return r.s.State["snsDestination"]
}

// Input properties used for looking up and filtering EventDestination resources.
type EventDestinationState struct {
	// CloudWatch destination for the events
	CloudwatchDestinations interface{}
	// The name of the configuration set
	ConfigurationSetName interface{}
	// If true, the event destination will be enabled
	Enabled interface{}
	// Send the events to a kinesis firehose destination
	KinesisDestination interface{}
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes interface{}
	// The name of the event destination
	Name interface{}
	// Send the events to an SNS Topic destination
	SnsDestination interface{}
}

// The set of arguments for constructing a EventDestination resource.
type EventDestinationArgs struct {
	// CloudWatch destination for the events
	CloudwatchDestinations interface{}
	// The name of the configuration set
	ConfigurationSetName interface{}
	// If true, the event destination will be enabled
	Enabled interface{}
	// Send the events to a kinesis firehose destination
	KinesisDestination interface{}
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes interface{}
	// The name of the event destination
	Name interface{}
	// Send the events to an SNS Topic destination
	SnsDestination interface{}
}
