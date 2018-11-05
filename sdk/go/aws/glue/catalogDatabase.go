// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Glue Catalog Database Resource. You can refer to the [Glue Developer Guide](http://docs.aws.amazon.com/glue/latest/dg/populate-data-catalog.html) for a full explanation of the Glue Data Catalog functionality
type CatalogDatabase struct {
	s *pulumi.ResourceState
}

// NewCatalogDatabase registers a new resource with the given unique name, arguments, and options.
func NewCatalogDatabase(ctx *pulumi.Context,
	name string, args *CatalogDatabaseArgs, opts ...pulumi.ResourceOpt) (*CatalogDatabase, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["catalogId"] = nil
		inputs["description"] = nil
		inputs["locationUri"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
	} else {
		inputs["catalogId"] = args.CatalogId
		inputs["description"] = args.Description
		inputs["locationUri"] = args.LocationUri
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
	}
	s, err := ctx.RegisterResource("aws:glue/catalogDatabase:CatalogDatabase", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CatalogDatabase{s: s}, nil
}

// GetCatalogDatabase gets an existing CatalogDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCatalogDatabase(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CatalogDatabaseState, opts ...pulumi.ResourceOpt) (*CatalogDatabase, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["catalogId"] = state.CatalogId
		inputs["description"] = state.Description
		inputs["locationUri"] = state.LocationUri
		inputs["name"] = state.Name
		inputs["parameters"] = state.Parameters
	}
	s, err := ctx.ReadResource("aws:glue/catalogDatabase:CatalogDatabase", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CatalogDatabase{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CatalogDatabase) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CatalogDatabase) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
func (r *CatalogDatabase) CatalogId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["catalogId"])
}

// Description of the database.
func (r *CatalogDatabase) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The location of the database (for example, an HDFS path).
func (r *CatalogDatabase) LocationUri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["locationUri"])
}

// The name of the database.
func (r *CatalogDatabase) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A list of key-value pairs that define parameters and properties of the database.
func (r *CatalogDatabase) Parameters() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["parameters"])
}

// Input properties used for looking up and filtering CatalogDatabase resources.
type CatalogDatabaseState struct {
	// ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
	CatalogId interface{}
	// Description of the database.
	Description interface{}
	// The location of the database (for example, an HDFS path).
	LocationUri interface{}
	// The name of the database.
	Name interface{}
	// A list of key-value pairs that define parameters and properties of the database.
	Parameters interface{}
}

// The set of arguments for constructing a CatalogDatabase resource.
type CatalogDatabaseArgs struct {
	// ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
	CatalogId interface{}
	// Description of the database.
	Description interface{}
	// The location of the database (for example, an HDFS path).
	LocationUri interface{}
	// The name of the database.
	Name interface{}
	// A list of key-value pairs that define parameters and properties of the database.
	Parameters interface{}
}
