// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Rate Based Rule Resource
type RateBasedRule struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The name or description of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayOutput `pulumi:"predicates"`
	// Valid value is IP.
	RateKey pulumi.StringOutput `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntOutput `pulumi:"rateLimit"`
	// Key-value map of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewRateBasedRule registers a new resource with the given unique name, arguments, and options.
func NewRateBasedRule(ctx *pulumi.Context,
	name string, args *RateBasedRuleArgs, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	if args == nil || args.RateKey == nil {
		return nil, errors.New("missing required argument 'RateKey'")
	}
	if args == nil || args.RateLimit == nil {
		return nil, errors.New("missing required argument 'RateLimit'")
	}
	if args == nil {
		args = &RateBasedRuleArgs{}
	}
	var resource RateBasedRule
	err := ctx.RegisterResource("aws:waf/rateBasedRule:RateBasedRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRateBasedRule gets an existing RateBasedRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRateBasedRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RateBasedRuleState, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	var resource RateBasedRule
	err := ctx.ReadResource("aws:waf/rateBasedRule:RateBasedRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RateBasedRule resources.
type rateBasedRuleState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RateBasedRulePredicate `pulumi:"predicates"`
	// Valid value is IP.
	RateKey *string `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit *int `pulumi:"rateLimit"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

type RateBasedRuleState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringPtrInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayInput
	// Valid value is IP.
	RateKey pulumi.StringPtrInput
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntPtrInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
}

func (RateBasedRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleState)(nil)).Elem()
}

type rateBasedRuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName string `pulumi:"metricName"`
	// The name or description of the rule.
	Name *string `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates []RateBasedRulePredicate `pulumi:"predicates"`
	// Valid value is IP.
	RateKey string `pulumi:"rateKey"`
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit int `pulumi:"rateLimit"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a RateBasedRule resource.
type RateBasedRuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringInput
	// The name or description of the rule.
	Name pulumi.StringPtrInput
	// The objects to include in a rule (documented below).
	Predicates RateBasedRulePredicateArrayInput
	// Valid value is IP.
	RateKey pulumi.StringInput
	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit pulumi.IntInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
}

func (RateBasedRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleArgs)(nil)).Elem()
}
