// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glacier

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality
//
// > **NOTE:** When removing a Glacier Vault, the Vault must be empty.
type Vault struct {
	pulumi.CustomResourceState

	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy pulumi.StringPtrOutput `pulumi:"accessPolicy"`
	// The ARN of the vault.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The URI of the vault that was created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name pulumi.StringOutput `pulumi:"name"`
	// The notifications for the Vault. Fields documented below.
	Notifications VaultNotificationArrayOutput `pulumi:"notifications"`
	// A map of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		args = &VaultArgs{}
	}
	var resource Vault
	err := ctx.RegisterResource("aws:glacier/vault:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("aws:glacier/vault:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy *string `pulumi:"accessPolicy"`
	// The ARN of the vault.
	Arn *string `pulumi:"arn"`
	// The URI of the vault that was created.
	Location *string `pulumi:"location"`
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name *string `pulumi:"name"`
	// The notifications for the Vault. Fields documented below.
	Notifications []VaultNotification `pulumi:"notifications"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type VaultState struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy pulumi.StringPtrInput
	// The ARN of the vault.
	Arn pulumi.StringPtrInput
	// The URI of the vault that was created.
	Location pulumi.StringPtrInput
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name pulumi.StringPtrInput
	// The notifications for the Vault. Fields documented below.
	Notifications VaultNotificationArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy *string `pulumi:"accessPolicy"`
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name *string `pulumi:"name"`
	// The notifications for the Vault. Fields documented below.
	Notifications []VaultNotification `pulumi:"notifications"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy pulumi.StringPtrInput
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name pulumi.StringPtrInput
	// The notifications for the Vault. Fields documented below.
	Notifications VaultNotificationArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}
