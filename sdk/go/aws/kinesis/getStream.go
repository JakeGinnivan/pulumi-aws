// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about a Kinesis Stream for use in other
// resources.
//
// For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
func LookupStream(ctx *pulumi.Context, args *LookupStreamArgs, opts ...pulumi.InvokeOption) (*LookupStreamResult, error) {
	var rv LookupStreamResult
	err := ctx.Invoke("aws:kinesis/getStream:getStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStream.
type LookupStreamArgs struct {
	// The name of the Kinesis Stream.
	Name string `pulumi:"name"`
	// A map of tags to assigned to the stream.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getStream.
type LookupStreamResult struct {
	// The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).
	Arn string `pulumi:"arn"`
	// The list of shard ids in the CLOSED state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
	ClosedShards []string `pulumi:"closedShards"`
	// The approximate UNIX timestamp that the stream was created.
	CreationTimestamp int `pulumi:"creationTimestamp"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Kinesis Stream.
	Name string `pulumi:"name"`
	// The list of shard ids in the OPEN state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
	OpenShards []string `pulumi:"openShards"`
	// Length of time (in hours) data records are accessible after they are added to the stream.
	RetentionPeriod int `pulumi:"retentionPeriod"`
	// A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more.
	ShardLevelMetrics []string `pulumi:"shardLevelMetrics"`
	// The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.
	Status string `pulumi:"status"`
	// A map of tags to assigned to the stream.
	Tags map[string]interface{} `pulumi:"tags"`
}
