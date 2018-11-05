// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new Amazon Redshift security group. You use security groups to control access to non-VPC clusters
type SecurityGroup struct {
	s *pulumi.ResourceState
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOpt) (*SecurityGroup, error) {
	if args == nil || args.Ingress == nil {
		return nil, errors.New("missing required argument 'Ingress'")
	}
	inputs := make(map[string]interface{})
	inputs["description"] = "Managed by Pulumi"
	if args == nil {
		inputs["ingress"] = nil
		inputs["name"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["ingress"] = args.Ingress
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:redshift/securityGroup:SecurityGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityGroup{s: s}, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecurityGroupState, opts ...pulumi.ResourceOpt) (*SecurityGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["ingress"] = state.Ingress
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:redshift/securityGroup:SecurityGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecurityGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecurityGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The description of the Redshift security group. Defaults to "Managed by Terraform".
func (r *SecurityGroup) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// A list of ingress rules.
func (r *SecurityGroup) Ingress() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ingress"])
}

// The name of the Redshift security group.
func (r *SecurityGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering SecurityGroup resources.
type SecurityGroupState struct {
	// The description of the Redshift security group. Defaults to "Managed by Terraform".
	Description interface{}
	// A list of ingress rules.
	Ingress interface{}
	// The name of the Redshift security group.
	Name interface{}
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The description of the Redshift security group. Defaults to "Managed by Terraform".
	Description interface{}
	// A list of ingress rules.
	Ingress interface{}
	// The name of the Redshift security group.
	Name interface{}
}
