// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ## Import
//
// This resource can be imported using an ID made up of the secret name:
//
// ```sh
// $ pulumi import github:index/actionsOrganizationSecret:ActionsOrganizationSecret test_secret test_secret_name
// ```
// NOTE: the implementation is limited in that it won't fetch the value of the
// `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
type ActionsOrganizationSecret struct {
	pulumi.CustomResourceState

	// Date of actionsSecret creation.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrOutput `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrOutput `pulumi:"plaintextValue"`
	// Name of the secret
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayOutput `pulumi:"selectedRepositoryIds"`
	// Date of actionsSecret update.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewActionsOrganizationSecret registers a new resource with the given unique name, arguments, and options.
func NewActionsOrganizationSecret(ctx *pulumi.Context,
	name string, args *ActionsOrganizationSecretArgs, opts ...pulumi.ResourceOption) (*ActionsOrganizationSecret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecretName == nil {
		return nil, errors.New("invalid value for required argument 'SecretName'")
	}
	if args.Visibility == nil {
		return nil, errors.New("invalid value for required argument 'Visibility'")
	}
	if args.EncryptedValue != nil {
		args.EncryptedValue = pulumi.ToSecret(args.EncryptedValue).(pulumi.StringPtrInput)
	}
	if args.PlaintextValue != nil {
		args.PlaintextValue = pulumi.ToSecret(args.PlaintextValue).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"encryptedValue",
		"plaintextValue",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ActionsOrganizationSecret
	err := ctx.RegisterResource("github:index/actionsOrganizationSecret:ActionsOrganizationSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsOrganizationSecret gets an existing ActionsOrganizationSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsOrganizationSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsOrganizationSecretState, opts ...pulumi.ResourceOption) (*ActionsOrganizationSecret, error) {
	var resource ActionsOrganizationSecret
	err := ctx.ReadResource("github:index/actionsOrganizationSecret:ActionsOrganizationSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsOrganizationSecret resources.
type actionsOrganizationSecretState struct {
	// Date of actionsSecret creation.
	CreatedAt *string `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue *string `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue *string `pulumi:"plaintextValue"`
	// Name of the secret
	SecretName *string `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
	// Date of actionsSecret update.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility *string `pulumi:"visibility"`
}

type ActionsOrganizationSecretState struct {
	// Date of actionsSecret creation.
	CreatedAt pulumi.StringPtrInput
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrInput
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrInput
	// Name of the secret
	SecretName pulumi.StringPtrInput
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayInput
	// Date of actionsSecret update.
	UpdatedAt pulumi.StringPtrInput
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringPtrInput
}

func (ActionsOrganizationSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationSecretState)(nil)).Elem()
}

type actionsOrganizationSecretArgs struct {
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue *string `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue *string `pulumi:"plaintextValue"`
	// Name of the secret
	SecretName string `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility string `pulumi:"visibility"`
}

// The set of arguments for constructing a ActionsOrganizationSecret resource.
type ActionsOrganizationSecretArgs struct {
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrInput
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrInput
	// Name of the secret
	SecretName pulumi.StringInput
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayInput
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringInput
}

func (ActionsOrganizationSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationSecretArgs)(nil)).Elem()
}

type ActionsOrganizationSecretInput interface {
	pulumi.Input

	ToActionsOrganizationSecretOutput() ActionsOrganizationSecretOutput
	ToActionsOrganizationSecretOutputWithContext(ctx context.Context) ActionsOrganizationSecretOutput
}

func (*ActionsOrganizationSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationSecret)(nil)).Elem()
}

func (i *ActionsOrganizationSecret) ToActionsOrganizationSecretOutput() ActionsOrganizationSecretOutput {
	return i.ToActionsOrganizationSecretOutputWithContext(context.Background())
}

func (i *ActionsOrganizationSecret) ToActionsOrganizationSecretOutputWithContext(ctx context.Context) ActionsOrganizationSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationSecretOutput)
}

// ActionsOrganizationSecretArrayInput is an input type that accepts ActionsOrganizationSecretArray and ActionsOrganizationSecretArrayOutput values.
// You can construct a concrete instance of `ActionsOrganizationSecretArrayInput` via:
//
//	ActionsOrganizationSecretArray{ ActionsOrganizationSecretArgs{...} }
type ActionsOrganizationSecretArrayInput interface {
	pulumi.Input

	ToActionsOrganizationSecretArrayOutput() ActionsOrganizationSecretArrayOutput
	ToActionsOrganizationSecretArrayOutputWithContext(context.Context) ActionsOrganizationSecretArrayOutput
}

type ActionsOrganizationSecretArray []ActionsOrganizationSecretInput

func (ActionsOrganizationSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationSecret)(nil)).Elem()
}

func (i ActionsOrganizationSecretArray) ToActionsOrganizationSecretArrayOutput() ActionsOrganizationSecretArrayOutput {
	return i.ToActionsOrganizationSecretArrayOutputWithContext(context.Background())
}

func (i ActionsOrganizationSecretArray) ToActionsOrganizationSecretArrayOutputWithContext(ctx context.Context) ActionsOrganizationSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationSecretArrayOutput)
}

// ActionsOrganizationSecretMapInput is an input type that accepts ActionsOrganizationSecretMap and ActionsOrganizationSecretMapOutput values.
// You can construct a concrete instance of `ActionsOrganizationSecretMapInput` via:
//
//	ActionsOrganizationSecretMap{ "key": ActionsOrganizationSecretArgs{...} }
type ActionsOrganizationSecretMapInput interface {
	pulumi.Input

	ToActionsOrganizationSecretMapOutput() ActionsOrganizationSecretMapOutput
	ToActionsOrganizationSecretMapOutputWithContext(context.Context) ActionsOrganizationSecretMapOutput
}

type ActionsOrganizationSecretMap map[string]ActionsOrganizationSecretInput

func (ActionsOrganizationSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationSecret)(nil)).Elem()
}

func (i ActionsOrganizationSecretMap) ToActionsOrganizationSecretMapOutput() ActionsOrganizationSecretMapOutput {
	return i.ToActionsOrganizationSecretMapOutputWithContext(context.Background())
}

func (i ActionsOrganizationSecretMap) ToActionsOrganizationSecretMapOutputWithContext(ctx context.Context) ActionsOrganizationSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationSecretMapOutput)
}

type ActionsOrganizationSecretOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationSecret)(nil)).Elem()
}

func (o ActionsOrganizationSecretOutput) ToActionsOrganizationSecretOutput() ActionsOrganizationSecretOutput {
	return o
}

func (o ActionsOrganizationSecretOutput) ToActionsOrganizationSecretOutputWithContext(ctx context.Context) ActionsOrganizationSecretOutput {
	return o
}

// Date of actionsSecret creation.
func (o ActionsOrganizationSecretOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecret) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Encrypted value of the secret using the GitHub public key in Base64 format.
func (o ActionsOrganizationSecretOutput) EncryptedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecret) pulumi.StringPtrOutput { return v.EncryptedValue }).(pulumi.StringPtrOutput)
}

// Plaintext value of the secret to be encrypted
func (o ActionsOrganizationSecretOutput) PlaintextValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecret) pulumi.StringPtrOutput { return v.PlaintextValue }).(pulumi.StringPtrOutput)
}

// Name of the secret
func (o ActionsOrganizationSecretOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecret) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// An array of repository ids that can access the organization secret.
func (o ActionsOrganizationSecretOutput) SelectedRepositoryIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecret) pulumi.IntArrayOutput { return v.SelectedRepositoryIds }).(pulumi.IntArrayOutput)
}

// Date of actionsSecret update.
func (o ActionsOrganizationSecretOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecret) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Configures the access that repositories have to the organization secret.
// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
func (o ActionsOrganizationSecretOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationSecret) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type ActionsOrganizationSecretArrayOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationSecret)(nil)).Elem()
}

func (o ActionsOrganizationSecretArrayOutput) ToActionsOrganizationSecretArrayOutput() ActionsOrganizationSecretArrayOutput {
	return o
}

func (o ActionsOrganizationSecretArrayOutput) ToActionsOrganizationSecretArrayOutputWithContext(ctx context.Context) ActionsOrganizationSecretArrayOutput {
	return o
}

func (o ActionsOrganizationSecretArrayOutput) Index(i pulumi.IntInput) ActionsOrganizationSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsOrganizationSecret {
		return vs[0].([]*ActionsOrganizationSecret)[vs[1].(int)]
	}).(ActionsOrganizationSecretOutput)
}

type ActionsOrganizationSecretMapOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationSecret)(nil)).Elem()
}

func (o ActionsOrganizationSecretMapOutput) ToActionsOrganizationSecretMapOutput() ActionsOrganizationSecretMapOutput {
	return o
}

func (o ActionsOrganizationSecretMapOutput) ToActionsOrganizationSecretMapOutputWithContext(ctx context.Context) ActionsOrganizationSecretMapOutput {
	return o
}

func (o ActionsOrganizationSecretMapOutput) MapIndex(k pulumi.StringInput) ActionsOrganizationSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsOrganizationSecret {
		return vs[0].(map[string]*ActionsOrganizationSecret)[vs[1].(string)]
	}).(ActionsOrganizationSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationSecretInput)(nil)).Elem(), &ActionsOrganizationSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationSecretArrayInput)(nil)).Elem(), ActionsOrganizationSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationSecretMapInput)(nil)).Elem(), ActionsOrganizationSecretMap{})
	pulumi.RegisterOutputType(ActionsOrganizationSecretOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationSecretArrayOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationSecretMapOutput{})
}
