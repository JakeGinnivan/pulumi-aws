// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 API.
 * 
 * > **Note:** Amazon API Gateway Version 2 resources are used for creating and deploying WebSocket and HTTP APIs. To create and deploy REST APIs, use Amazon API Gateway Version 1.
 * 
 * ## Example Usage
 * 
 * ### Basic WebSocket API
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.apigatewayv2.Api("example", {
 *     protocolType: "WEBSOCKET",
 *     routeSelectionExpression: "$request.body.action",
 * });
 * ```
 * 
 * ### Basic HTTP API
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.apigatewayv2.Api("example", {
 *     protocolType: "HTTP",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/apigatewayv2_api.html.markdown.
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiState, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/api:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * The URI of the API, of the form `{api-id}.execute-api.{region}.amazonaws.com`.
     */
    public /*out*/ readonly apiEndpoint!: pulumi.Output<string>;
    /**
     * An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
     * Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
     * Applicable for WebSocket APIs.
     */
    public readonly apiKeySelectionExpression!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the API.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
     */
    public readonly corsConfiguration!: pulumi.Output<outputs.apigatewayv2.ApiCorsConfiguration | undefined>;
    /**
     * Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
     */
    public readonly credentialsArn!: pulumi.Output<string | undefined>;
    /**
     * The description of the API.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ARN prefix to be used in an [`aws.lambda.Permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn` attribute
     * or in an [`aws.iam.Policy`](https://www.terraform.io/docs/providers/aws/r/iam_policy.html) to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
     * See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
     */
    public /*out*/ readonly executionArn!: pulumi.Output<string>;
    /**
     * The name of the API.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
     */
    public readonly protocolType!: pulumi.Output<string>;
    /**
     * Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
     */
    public readonly routeKey!: pulumi.Output<string | undefined>;
    /**
     * The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
     * Defaults to `$request.method $request.path`.
     */
    public readonly routeSelectionExpression!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the API.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
     * For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
     * The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
     */
    public readonly target!: pulumi.Output<string | undefined>;
    /**
     * A version identifier for the API.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiArgs | ApiState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ApiState | undefined;
            inputs["apiEndpoint"] = state ? state.apiEndpoint : undefined;
            inputs["apiKeySelectionExpression"] = state ? state.apiKeySelectionExpression : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["corsConfiguration"] = state ? state.corsConfiguration : undefined;
            inputs["credentialsArn"] = state ? state.credentialsArn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["executionArn"] = state ? state.executionArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protocolType"] = state ? state.protocolType : undefined;
            inputs["routeKey"] = state ? state.routeKey : undefined;
            inputs["routeSelectionExpression"] = state ? state.routeSelectionExpression : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["target"] = state ? state.target : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ApiArgs | undefined;
            if (!args || args.protocolType === undefined) {
                throw new Error("Missing required property 'protocolType'");
            }
            inputs["apiKeySelectionExpression"] = args ? args.apiKeySelectionExpression : undefined;
            inputs["corsConfiguration"] = args ? args.corsConfiguration : undefined;
            inputs["credentialsArn"] = args ? args.credentialsArn : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocolType"] = args ? args.protocolType : undefined;
            inputs["routeKey"] = args ? args.routeKey : undefined;
            inputs["routeSelectionExpression"] = args ? args.routeSelectionExpression : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["apiEndpoint"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["executionArn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Api.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Api resources.
 */
export interface ApiState {
    /**
     * The URI of the API, of the form `{api-id}.execute-api.{region}.amazonaws.com`.
     */
    readonly apiEndpoint?: pulumi.Input<string>;
    /**
     * An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
     * Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
     * Applicable for WebSocket APIs.
     */
    readonly apiKeySelectionExpression?: pulumi.Input<string>;
    /**
     * The ARN of the API.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
     */
    readonly corsConfiguration?: pulumi.Input<inputs.apigatewayv2.ApiCorsConfiguration>;
    /**
     * Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
     */
    readonly credentialsArn?: pulumi.Input<string>;
    /**
     * The description of the API.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ARN prefix to be used in an [`aws.lambda.Permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `sourceArn` attribute
     * or in an [`aws.iam.Policy`](https://www.terraform.io/docs/providers/aws/r/iam_policy.html) to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
     * See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
     */
    readonly executionArn?: pulumi.Input<string>;
    /**
     * The name of the API.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
     */
    readonly protocolType?: pulumi.Input<string>;
    /**
     * Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
     */
    readonly routeKey?: pulumi.Input<string>;
    /**
     * The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
     * Defaults to `$request.method $request.path`.
     */
    readonly routeSelectionExpression?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the API.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
     * For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
     * The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
     */
    readonly target?: pulumi.Input<string>;
    /**
     * A version identifier for the API.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    /**
     * An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
     * Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
     * Applicable for WebSocket APIs.
     */
    readonly apiKeySelectionExpression?: pulumi.Input<string>;
    /**
     * The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
     */
    readonly corsConfiguration?: pulumi.Input<inputs.apigatewayv2.ApiCorsConfiguration>;
    /**
     * Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
     */
    readonly credentialsArn?: pulumi.Input<string>;
    /**
     * The description of the API.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the API.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
     */
    readonly protocolType: pulumi.Input<string>;
    /**
     * Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
     */
    readonly routeKey?: pulumi.Input<string>;
    /**
     * The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
     * Defaults to `$request.method $request.path`.
     */
    readonly routeSelectionExpression?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the API.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
     * For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
     * The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
     */
    readonly target?: pulumi.Input<string>;
    /**
     * A version identifier for the API.
     */
    readonly version?: pulumi.Input<string>;
}
