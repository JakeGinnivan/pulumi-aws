// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM group.
type Group struct {
	s *pulumi.ResourceState
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOpt) (*Group, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["path"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["path"] = args.Path
	}
	inputs["arn"] = nil
	inputs["uniqueId"] = nil
	s, err := ctx.RegisterResource("aws:iam/group:Group", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Group{s: s}, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GroupState, opts ...pulumi.ResourceOpt) (*Group, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["name"] = state.Name
		inputs["path"] = state.Path
		inputs["uniqueId"] = state.UniqueId
	}
	s, err := ctx.ReadResource("aws:iam/group:Group", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Group{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Group) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Group) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN assigned by AWS for this group.
func (r *Group) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
func (r *Group) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Path in which to create the group.
func (r *Group) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// The [unique ID][1] assigned by AWS.
func (r *Group) UniqueId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["uniqueId"])
}

// Input properties used for looking up and filtering Group resources.
type GroupState struct {
	// The ARN assigned by AWS for this group.
	Arn interface{}
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name interface{}
	// Path in which to create the group.
	Path interface{}
	// The [unique ID][1] assigned by AWS.
	UniqueId interface{}
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name interface{}
	// Path in which to create the group.
	Path interface{}
}
