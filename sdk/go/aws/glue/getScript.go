// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).
func LookupScript(ctx *pulumi.Context, args *GetScriptArgs) (*GetScriptResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["dagEdges"] = args.DagEdges
		inputs["dagNodes"] = args.DagNodes
		inputs["language"] = args.Language
	}
	outputs, err := ctx.Invoke("aws:glue/getScript:getScript", inputs)
	if err != nil {
		return nil, err
	}
	return &GetScriptResult{
		PythonScript: outputs["pythonScript"],
		ScalaCode: outputs["scalaCode"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getScript.
type GetScriptArgs struct {
	// A list of the edges in the DAG. Defined below.
	DagEdges interface{}
	// A list of the nodes in the DAG. Defined below.
	DagNodes interface{}
	// The programming language of the resulting code from the DAG. Defaults to `PYTHON`. Valid values are `PYTHON` and `SCALA`.
	Language interface{}
}

// A collection of values returned by getScript.
type GetScriptResult struct {
	// The Python script generated from the DAG when the `language` argument is set to `PYTHON`.
	PythonScript interface{}
	// The Scala code generated from the DAG when the `language` argument is set to `SCALA`.
	ScalaCode interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}