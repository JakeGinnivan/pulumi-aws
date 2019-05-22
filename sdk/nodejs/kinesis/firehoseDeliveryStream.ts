// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Kinesis Firehose Delivery Stream resource. Amazon Kinesis Firehose is a fully managed, elastic service to easily deliver real-time data streams to destinations such as Amazon S3 and Amazon Redshift.
 * 
 * For more details, see the [Amazon Kinesis Firehose Documentation][1].
 * 
 * ## Example Usage
 * 
 * ### Extended S3 Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const firehoseRole = new aws.iam.Role("firehose_role", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "firehose.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const lambdaIam = new aws.iam.Role("lambda_iam", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "lambda.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 * });
 * const lambdaProcessor = new aws.lambda.Function("lambda_processor", {
 *     code: new pulumi.asset.FileArchive("lambda.zip"),
 *     handler: "exports.handler",
 *     role: lambdaIam.arn,
 *     runtime: "nodejs8.10",
 * });
 * const extendedS3Stream = new aws.kinesis.FirehoseDeliveryStream("extended_s3_stream", {
 *     destination: "extended_s3",
 *     extendedS3Configuration: {
 *         bucketArn: bucket.arn,
 *         processingConfiguration: {
 *             enabled: true,
 *             processors: [{
 *                 parameters: [{
 *                     parameterName: "LambdaArn",
 *                     parameterValue: pulumi.interpolate`${lambdaProcessor.arn}:$LATEST`,
 *                 }],
 *                 type: "Lambda",
 *             }],
 *         },
 *         roleArn: firehoseRole.arn,
 *     },
 * });
 * ```
 * 
 * ### S3 Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const firehoseRole = new aws.iam.Role("firehose_role", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "firehose.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 * });
 * const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
 *     destination: "s3",
 *     s3Configuration: {
 *         bucketArn: bucket.arn,
 *         roleArn: firehoseRole.arn,
 *     },
 * });
 * ```
 * 
 * ### Redshift Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const testCluster = new aws.redshift.Cluster("test_cluster", {
 *     clusterIdentifier: "tf-redshift-cluster-%d",
 *     clusterType: "single-node",
 *     databaseName: "test",
 *     masterPassword: "T3stPass",
 *     masterUsername: "testuser",
 *     nodeType: "dc1.large",
 * });
 * const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
 *     destination: "redshift",
 *     redshiftConfiguration: {
 *         clusterJdbcurl: pulumi.interpolate`jdbc:redshift://${testCluster.endpoint}/${testCluster.databaseName}`,
 *         copyOptions: "delimiter '|'", // the default delimiter
 *         dataTableColumns: "test-col",
 *         dataTableName: "test-table",
 *         password: "T3stPass",
 *         roleArn: aws_iam_role_firehose_role.arn,
 *         s3BackupConfiguration: {
 *             bucketArn: aws_s3_bucket_bucket.arn,
 *             bufferInterval: 300,
 *             bufferSize: 15,
 *             compressionFormat: "GZIP",
 *             roleArn: aws_iam_role_firehose_role.arn,
 *         },
 *         s3BackupMode: "Enabled",
 *         username: "testuser",
 *     },
 *     s3Configuration: {
 *         bucketArn: aws_s3_bucket_bucket.arn,
 *         bufferInterval: 400,
 *         bufferSize: 10,
 *         compressionFormat: "GZIP",
 *         roleArn: aws_iam_role_firehose_role.arn,
 *     },
 * });
 * ```
 * 
 * ### Elasticsearch Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const testCluster = new aws.elasticsearch.Domain("test_cluster", {});
 * const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
 *     destination: "elasticsearch",
 *     elasticsearchConfiguration: {
 *         domainArn: testCluster.arn,
 *         indexName: "test",
 *         processingConfiguration: {
 *             enabled: true,
 *             processors: [{
 *                 parameters: [{
 *                     parameterName: "LambdaArn",
 *                     parameterValue: pulumi.interpolate`${aws_lambda_function_lambda_processor.arn}:$LATEST`,
 *                 }],
 *                 type: "Lambda",
 *             }],
 *         },
 *         roleArn: aws_iam_role_firehose_role.arn,
 *         typeName: "test",
 *     },
 *     s3Configuration: {
 *         bucketArn: aws_s3_bucket_bucket.arn,
 *         bufferInterval: 400,
 *         bufferSize: 10,
 *         compressionFormat: "GZIP",
 *         roleArn: aws_iam_role_firehose_role.arn,
 *     },
 * });
 * ```
 * 
 * 
 * ### Splunk Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
 *     destination: "splunk",
 *     s3Configuration: {
 *         bucketArn: aws_s3_bucket_bucket.arn,
 *         bufferInterval: 400,
 *         bufferSize: 10,
 *         compressionFormat: "GZIP",
 *         roleArn: aws_iam_role_firehose.arn,
 *     },
 *     splunkConfiguration: {
 *         hecAcknowledgmentTimeout: 600,
 *         hecEndpoint: "https://http-inputs-mydomain.splunkcloud.com:443",
 *         hecEndpointType: "Event",
 *         hecToken: "51D4DA16-C61B-4F5F-8EC7-ED4301342A4A",
 *         s3BackupMode: "FailedEventsOnly",
 *     },
 * });
 * ```
 */
export class FirehoseDeliveryStream extends pulumi.CustomResource {
    /**
     * Get an existing FirehoseDeliveryStream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirehoseDeliveryStreamState, opts?: pulumi.CustomResourceOptions): FirehoseDeliveryStream {
        return new FirehoseDeliveryStream(name, <any>state, { ...opts, id: id });
    }

    /**
     * The Amazon Resource Name (ARN) specifying the Stream
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
     */
    public readonly destination!: pulumi.Output<string>;
    public readonly destinationId!: pulumi.Output<string>;
    public readonly elasticsearchConfiguration!: pulumi.Output<{ bufferingInterval?: number, bufferingSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }, domainArn: string, indexName: string, indexRotationPeriod?: string, processingConfiguration?: { enabled?: boolean, processors?: { parameters?: { parameterName: string, parameterValue: string }[], type: string }[] }, retryDuration?: number, roleArn: string, s3BackupMode?: string, typeName?: string } | undefined>;
    /**
     * Enhanced configuration options for the s3 destination. More details are given below.
     */
    public readonly extendedS3Configuration!: pulumi.Output<{ bucketArn: string, bufferInterval?: number, bufferSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }, compressionFormat?: string, dataFormatConversionConfiguration?: { enabled?: boolean, inputFormatConfiguration: { deserializer: { hiveJsonSerDe?: { timestampFormats?: string[] }, openXJsonSerDe?: { caseInsensitive?: boolean, columnToJsonKeyMappings?: {[key: string]: string}, convertDotsInJsonKeysToUnderscores?: boolean } } }, outputFormatConfiguration: { serializer: { orcSerDe?: { blockSizeBytes?: number, bloomFilterColumns?: string[], bloomFilterFalsePositiveProbability?: number, compression?: string, dictionaryKeyThreshold?: number, enablePadding?: boolean, formatVersion?: string, paddingTolerance?: number, rowIndexStride?: number, stripeSizeBytes?: number }, parquetSerDe?: { blockSizeBytes?: number, compression?: string, enableDictionaryCompression?: boolean, maxPaddingBytes?: number, pageSizeBytes?: number, writerVersion?: string } } }, schemaConfiguration: { catalogId: string, databaseName: string, region: string, roleArn: string, tableName: string, versionId?: string } }, errorOutputPrefix?: string, kmsKeyArn?: string, prefix?: string, processingConfiguration?: { enabled?: boolean, processors?: { parameters?: { parameterName: string, parameterValue: string }[], type: string }[] }, roleArn: string, s3BackupConfiguration?: { bucketArn: string, bufferInterval?: number, bufferSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }, compressionFormat?: string, kmsKeyArn?: string, prefix?: string, roleArn: string }, s3BackupMode?: string } | undefined>;
    /**
     * Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
     */
    public readonly kinesisSourceConfiguration!: pulumi.Output<{ kinesisStreamArn: string, roleArn: string } | undefined>;
    /**
     * A name to identify the stream. This is unique to the
     * AWS account and region the Stream is created in.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration options if redshift is the destination.
     * Using `redshift_configuration` requires the user to also specify a
     * `s3_configuration` block. More details are given below.
     */
    public readonly redshiftConfiguration!: pulumi.Output<{ cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }, clusterJdbcurl: string, copyOptions?: string, dataTableColumns?: string, dataTableName: string, password: string, processingConfiguration?: { enabled?: boolean, processors?: { parameters?: { parameterName: string, parameterValue: string }[], type: string }[] }, retryDuration?: number, roleArn: string, s3BackupConfiguration?: { bucketArn: string, bufferInterval?: number, bufferSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }, compressionFormat?: string, kmsKeyArn?: string, prefix?: string, roleArn: string }, s3BackupMode?: string, username: string } | undefined>;
    /**
     * Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
     * is redshift). More details are given below.
     */
    public readonly s3Configuration!: pulumi.Output<{ bucketArn: string, bufferInterval?: number, bufferSize?: number, cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }, compressionFormat?: string, kmsKeyArn?: string, prefix?: string, roleArn: string } | undefined>;
    public readonly splunkConfiguration!: pulumi.Output<{ cloudwatchLoggingOptions: { enabled?: boolean, logGroupName?: string, logStreamName?: string }, hecAcknowledgmentTimeout?: number, hecEndpoint: string, hecEndpointType?: string, hecToken: string, processingConfiguration?: { enabled?: boolean, processors?: { parameters?: { parameterName: string, parameterValue: string }[], type: string }[] }, retryDuration?: number, s3BackupMode?: string } | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Specifies the table version for the output data schema. Defaults to `LATEST`.
     */
    public readonly versionId!: pulumi.Output<string>;

    /**
     * Create a FirehoseDeliveryStream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirehoseDeliveryStreamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirehoseDeliveryStreamArgs | FirehoseDeliveryStreamState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirehoseDeliveryStreamState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["destinationId"] = state ? state.destinationId : undefined;
            inputs["elasticsearchConfiguration"] = state ? state.elasticsearchConfiguration : undefined;
            inputs["extendedS3Configuration"] = state ? state.extendedS3Configuration : undefined;
            inputs["kinesisSourceConfiguration"] = state ? state.kinesisSourceConfiguration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["redshiftConfiguration"] = state ? state.redshiftConfiguration : undefined;
            inputs["s3Configuration"] = state ? state.s3Configuration : undefined;
            inputs["splunkConfiguration"] = state ? state.splunkConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["versionId"] = state ? state.versionId : undefined;
        } else {
            const args = argsOrState as FirehoseDeliveryStreamArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["destinationId"] = args ? args.destinationId : undefined;
            inputs["elasticsearchConfiguration"] = args ? args.elasticsearchConfiguration : undefined;
            inputs["extendedS3Configuration"] = args ? args.extendedS3Configuration : undefined;
            inputs["kinesisSourceConfiguration"] = args ? args.kinesisSourceConfiguration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["redshiftConfiguration"] = args ? args.redshiftConfiguration : undefined;
            inputs["s3Configuration"] = args ? args.s3Configuration : undefined;
            inputs["splunkConfiguration"] = args ? args.splunkConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["versionId"] = args ? args.versionId : undefined;
        }
        super("aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirehoseDeliveryStream resources.
 */
export interface FirehoseDeliveryStreamState {
    /**
     * The Amazon Resource Name (ARN) specifying the Stream
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
     */
    readonly destination?: pulumi.Input<string>;
    readonly destinationId?: pulumi.Input<string>;
    readonly elasticsearchConfiguration?: pulumi.Input<{ bufferingInterval?: pulumi.Input<number>, bufferingSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, domainArn: pulumi.Input<string>, indexName: pulumi.Input<string>, indexRotationPeriod?: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, retryDuration?: pulumi.Input<number>, roleArn: pulumi.Input<string>, s3BackupMode?: pulumi.Input<string>, typeName?: pulumi.Input<string> }>;
    /**
     * Enhanced configuration options for the s3 destination. More details are given below.
     */
    readonly extendedS3Configuration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, dataFormatConversionConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, inputFormatConfiguration: pulumi.Input<{ deserializer: pulumi.Input<{ hiveJsonSerDe?: pulumi.Input<{ timestampFormats?: pulumi.Input<pulumi.Input<string>[]> }>, openXJsonSerDe?: pulumi.Input<{ caseInsensitive?: pulumi.Input<boolean>, columnToJsonKeyMappings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, convertDotsInJsonKeysToUnderscores?: pulumi.Input<boolean> }> }> }>, outputFormatConfiguration: pulumi.Input<{ serializer: pulumi.Input<{ orcSerDe?: pulumi.Input<{ blockSizeBytes?: pulumi.Input<number>, bloomFilterColumns?: pulumi.Input<pulumi.Input<string>[]>, bloomFilterFalsePositiveProbability?: pulumi.Input<number>, compression?: pulumi.Input<string>, dictionaryKeyThreshold?: pulumi.Input<number>, enablePadding?: pulumi.Input<boolean>, formatVersion?: pulumi.Input<string>, paddingTolerance?: pulumi.Input<number>, rowIndexStride?: pulumi.Input<number>, stripeSizeBytes?: pulumi.Input<number> }>, parquetSerDe?: pulumi.Input<{ blockSizeBytes?: pulumi.Input<number>, compression?: pulumi.Input<string>, enableDictionaryCompression?: pulumi.Input<boolean>, maxPaddingBytes?: pulumi.Input<number>, pageSizeBytes?: pulumi.Input<number>, writerVersion?: pulumi.Input<string> }> }> }>, schemaConfiguration: pulumi.Input<{ catalogId?: pulumi.Input<string>, databaseName: pulumi.Input<string>, region?: pulumi.Input<string>, roleArn: pulumi.Input<string>, tableName: pulumi.Input<string>, versionId?: pulumi.Input<string> }> }>, errorOutputPrefix?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, roleArn: pulumi.Input<string>, s3BackupConfiguration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, roleArn: pulumi.Input<string> }>, s3BackupMode?: pulumi.Input<string> }>;
    /**
     * Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
     */
    readonly kinesisSourceConfiguration?: pulumi.Input<{ kinesisStreamArn: pulumi.Input<string>, roleArn: pulumi.Input<string> }>;
    /**
     * A name to identify the stream. This is unique to the
     * AWS account and region the Stream is created in.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration options if redshift is the destination.
     * Using `redshift_configuration` requires the user to also specify a
     * `s3_configuration` block. More details are given below.
     */
    readonly redshiftConfiguration?: pulumi.Input<{ cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, clusterJdbcurl: pulumi.Input<string>, copyOptions?: pulumi.Input<string>, dataTableColumns?: pulumi.Input<string>, dataTableName: pulumi.Input<string>, password: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, retryDuration?: pulumi.Input<number>, roleArn: pulumi.Input<string>, s3BackupConfiguration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, roleArn: pulumi.Input<string> }>, s3BackupMode?: pulumi.Input<string>, username: pulumi.Input<string> }>;
    /**
     * Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
     * is redshift). More details are given below.
     */
    readonly s3Configuration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, roleArn: pulumi.Input<string> }>;
    readonly splunkConfiguration?: pulumi.Input<{ cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, hecAcknowledgmentTimeout?: pulumi.Input<number>, hecEndpoint: pulumi.Input<string>, hecEndpointType?: pulumi.Input<string>, hecToken: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, retryDuration?: pulumi.Input<number>, s3BackupMode?: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the table version for the output data schema. Defaults to `LATEST`.
     */
    readonly versionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirehoseDeliveryStream resource.
 */
export interface FirehoseDeliveryStreamArgs {
    /**
     * The Amazon Resource Name (ARN) specifying the Stream
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
     */
    readonly destination: pulumi.Input<string>;
    readonly destinationId?: pulumi.Input<string>;
    readonly elasticsearchConfiguration?: pulumi.Input<{ bufferingInterval?: pulumi.Input<number>, bufferingSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, domainArn: pulumi.Input<string>, indexName: pulumi.Input<string>, indexRotationPeriod?: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, retryDuration?: pulumi.Input<number>, roleArn: pulumi.Input<string>, s3BackupMode?: pulumi.Input<string>, typeName?: pulumi.Input<string> }>;
    /**
     * Enhanced configuration options for the s3 destination. More details are given below.
     */
    readonly extendedS3Configuration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, dataFormatConversionConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, inputFormatConfiguration: pulumi.Input<{ deserializer: pulumi.Input<{ hiveJsonSerDe?: pulumi.Input<{ timestampFormats?: pulumi.Input<pulumi.Input<string>[]> }>, openXJsonSerDe?: pulumi.Input<{ caseInsensitive?: pulumi.Input<boolean>, columnToJsonKeyMappings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, convertDotsInJsonKeysToUnderscores?: pulumi.Input<boolean> }> }> }>, outputFormatConfiguration: pulumi.Input<{ serializer: pulumi.Input<{ orcSerDe?: pulumi.Input<{ blockSizeBytes?: pulumi.Input<number>, bloomFilterColumns?: pulumi.Input<pulumi.Input<string>[]>, bloomFilterFalsePositiveProbability?: pulumi.Input<number>, compression?: pulumi.Input<string>, dictionaryKeyThreshold?: pulumi.Input<number>, enablePadding?: pulumi.Input<boolean>, formatVersion?: pulumi.Input<string>, paddingTolerance?: pulumi.Input<number>, rowIndexStride?: pulumi.Input<number>, stripeSizeBytes?: pulumi.Input<number> }>, parquetSerDe?: pulumi.Input<{ blockSizeBytes?: pulumi.Input<number>, compression?: pulumi.Input<string>, enableDictionaryCompression?: pulumi.Input<boolean>, maxPaddingBytes?: pulumi.Input<number>, pageSizeBytes?: pulumi.Input<number>, writerVersion?: pulumi.Input<string> }> }> }>, schemaConfiguration: pulumi.Input<{ catalogId?: pulumi.Input<string>, databaseName: pulumi.Input<string>, region?: pulumi.Input<string>, roleArn: pulumi.Input<string>, tableName: pulumi.Input<string>, versionId?: pulumi.Input<string> }> }>, errorOutputPrefix?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, roleArn: pulumi.Input<string>, s3BackupConfiguration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, roleArn: pulumi.Input<string> }>, s3BackupMode?: pulumi.Input<string> }>;
    /**
     * Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
     */
    readonly kinesisSourceConfiguration?: pulumi.Input<{ kinesisStreamArn: pulumi.Input<string>, roleArn: pulumi.Input<string> }>;
    /**
     * A name to identify the stream. This is unique to the
     * AWS account and region the Stream is created in.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration options if redshift is the destination.
     * Using `redshift_configuration` requires the user to also specify a
     * `s3_configuration` block. More details are given below.
     */
    readonly redshiftConfiguration?: pulumi.Input<{ cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, clusterJdbcurl: pulumi.Input<string>, copyOptions?: pulumi.Input<string>, dataTableColumns?: pulumi.Input<string>, dataTableName: pulumi.Input<string>, password: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, retryDuration?: pulumi.Input<number>, roleArn: pulumi.Input<string>, s3BackupConfiguration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, roleArn: pulumi.Input<string> }>, s3BackupMode?: pulumi.Input<string>, username: pulumi.Input<string> }>;
    /**
     * Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
     * is redshift). More details are given below.
     */
    readonly s3Configuration?: pulumi.Input<{ bucketArn: pulumi.Input<string>, bufferInterval?: pulumi.Input<number>, bufferSize?: pulumi.Input<number>, cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, compressionFormat?: pulumi.Input<string>, kmsKeyArn?: pulumi.Input<string>, prefix?: pulumi.Input<string>, roleArn: pulumi.Input<string> }>;
    readonly splunkConfiguration?: pulumi.Input<{ cloudwatchLoggingOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, logGroupName?: pulumi.Input<string>, logStreamName?: pulumi.Input<string> }>, hecAcknowledgmentTimeout?: pulumi.Input<number>, hecEndpoint: pulumi.Input<string>, hecEndpointType?: pulumi.Input<string>, hecToken: pulumi.Input<string>, processingConfiguration?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, processors?: pulumi.Input<pulumi.Input<{ parameters?: pulumi.Input<pulumi.Input<{ parameterName: pulumi.Input<string>, parameterValue: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]> }>, retryDuration?: pulumi.Input<number>, s3BackupMode?: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the table version for the output data schema. Defaults to `LATEST`.
     */
    readonly versionId?: pulumi.Input<string>;
}
