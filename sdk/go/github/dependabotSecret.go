// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DependabotSecret struct {
	pulumi.CustomResourceState

	// Date of 'dependabot_secret' creation.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrOutput `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted.
	PlaintextValue pulumi.StringPtrOutput `pulumi:"plaintextValue"`
	// Name of the repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Name of the secret.
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// Date of 'dependabot_secret' update.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewDependabotSecret registers a new resource with the given unique name, arguments, and options.
func NewDependabotSecret(ctx *pulumi.Context,
	name string, args *DependabotSecretArgs, opts ...pulumi.ResourceOption) (*DependabotSecret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.SecretName == nil {
		return nil, errors.New("invalid value for required argument 'SecretName'")
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
	var resource DependabotSecret
	err := ctx.RegisterResource("github:index/dependabotSecret:DependabotSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDependabotSecret gets an existing DependabotSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDependabotSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DependabotSecretState, opts ...pulumi.ResourceOption) (*DependabotSecret, error) {
	var resource DependabotSecret
	err := ctx.ReadResource("github:index/dependabotSecret:DependabotSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DependabotSecret resources.
type dependabotSecretState struct {
	// Date of 'dependabot_secret' creation.
	CreatedAt *string `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue *string `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted.
	PlaintextValue *string `pulumi:"plaintextValue"`
	// Name of the repository.
	Repository *string `pulumi:"repository"`
	// Name of the secret.
	SecretName *string `pulumi:"secretName"`
	// Date of 'dependabot_secret' update.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type DependabotSecretState struct {
	// Date of 'dependabot_secret' creation.
	CreatedAt pulumi.StringPtrInput
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrInput
	// Plaintext value of the secret to be encrypted.
	PlaintextValue pulumi.StringPtrInput
	// Name of the repository.
	Repository pulumi.StringPtrInput
	// Name of the secret.
	SecretName pulumi.StringPtrInput
	// Date of 'dependabot_secret' update.
	UpdatedAt pulumi.StringPtrInput
}

func (DependabotSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*dependabotSecretState)(nil)).Elem()
}

type dependabotSecretArgs struct {
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue *string `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted.
	PlaintextValue *string `pulumi:"plaintextValue"`
	// Name of the repository.
	Repository string `pulumi:"repository"`
	// Name of the secret.
	SecretName string `pulumi:"secretName"`
}

// The set of arguments for constructing a DependabotSecret resource.
type DependabotSecretArgs struct {
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrInput
	// Plaintext value of the secret to be encrypted.
	PlaintextValue pulumi.StringPtrInput
	// Name of the repository.
	Repository pulumi.StringInput
	// Name of the secret.
	SecretName pulumi.StringInput
}

func (DependabotSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dependabotSecretArgs)(nil)).Elem()
}

type DependabotSecretInput interface {
	pulumi.Input

	ToDependabotSecretOutput() DependabotSecretOutput
	ToDependabotSecretOutputWithContext(ctx context.Context) DependabotSecretOutput
}

func (*DependabotSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**DependabotSecret)(nil)).Elem()
}

func (i *DependabotSecret) ToDependabotSecretOutput() DependabotSecretOutput {
	return i.ToDependabotSecretOutputWithContext(context.Background())
}

func (i *DependabotSecret) ToDependabotSecretOutputWithContext(ctx context.Context) DependabotSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependabotSecretOutput)
}

// DependabotSecretArrayInput is an input type that accepts DependabotSecretArray and DependabotSecretArrayOutput values.
// You can construct a concrete instance of `DependabotSecretArrayInput` via:
//
//	DependabotSecretArray{ DependabotSecretArgs{...} }
type DependabotSecretArrayInput interface {
	pulumi.Input

	ToDependabotSecretArrayOutput() DependabotSecretArrayOutput
	ToDependabotSecretArrayOutputWithContext(context.Context) DependabotSecretArrayOutput
}

type DependabotSecretArray []DependabotSecretInput

func (DependabotSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DependabotSecret)(nil)).Elem()
}

func (i DependabotSecretArray) ToDependabotSecretArrayOutput() DependabotSecretArrayOutput {
	return i.ToDependabotSecretArrayOutputWithContext(context.Background())
}

func (i DependabotSecretArray) ToDependabotSecretArrayOutputWithContext(ctx context.Context) DependabotSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependabotSecretArrayOutput)
}

// DependabotSecretMapInput is an input type that accepts DependabotSecretMap and DependabotSecretMapOutput values.
// You can construct a concrete instance of `DependabotSecretMapInput` via:
//
//	DependabotSecretMap{ "key": DependabotSecretArgs{...} }
type DependabotSecretMapInput interface {
	pulumi.Input

	ToDependabotSecretMapOutput() DependabotSecretMapOutput
	ToDependabotSecretMapOutputWithContext(context.Context) DependabotSecretMapOutput
}

type DependabotSecretMap map[string]DependabotSecretInput

func (DependabotSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DependabotSecret)(nil)).Elem()
}

func (i DependabotSecretMap) ToDependabotSecretMapOutput() DependabotSecretMapOutput {
	return i.ToDependabotSecretMapOutputWithContext(context.Background())
}

func (i DependabotSecretMap) ToDependabotSecretMapOutputWithContext(ctx context.Context) DependabotSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependabotSecretMapOutput)
}

type DependabotSecretOutput struct{ *pulumi.OutputState }

func (DependabotSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DependabotSecret)(nil)).Elem()
}

func (o DependabotSecretOutput) ToDependabotSecretOutput() DependabotSecretOutput {
	return o
}

func (o DependabotSecretOutput) ToDependabotSecretOutputWithContext(ctx context.Context) DependabotSecretOutput {
	return o
}

// Date of 'dependabot_secret' creation.
func (o DependabotSecretOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotSecret) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Encrypted value of the secret using the GitHub public key in Base64 format.
func (o DependabotSecretOutput) EncryptedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DependabotSecret) pulumi.StringPtrOutput { return v.EncryptedValue }).(pulumi.StringPtrOutput)
}

// Plaintext value of the secret to be encrypted.
func (o DependabotSecretOutput) PlaintextValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DependabotSecret) pulumi.StringPtrOutput { return v.PlaintextValue }).(pulumi.StringPtrOutput)
}

// Name of the repository.
func (o DependabotSecretOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotSecret) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// Name of the secret.
func (o DependabotSecretOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotSecret) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// Date of 'dependabot_secret' update.
func (o DependabotSecretOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotSecret) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type DependabotSecretArrayOutput struct{ *pulumi.OutputState }

func (DependabotSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DependabotSecret)(nil)).Elem()
}

func (o DependabotSecretArrayOutput) ToDependabotSecretArrayOutput() DependabotSecretArrayOutput {
	return o
}

func (o DependabotSecretArrayOutput) ToDependabotSecretArrayOutputWithContext(ctx context.Context) DependabotSecretArrayOutput {
	return o
}

func (o DependabotSecretArrayOutput) Index(i pulumi.IntInput) DependabotSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DependabotSecret {
		return vs[0].([]*DependabotSecret)[vs[1].(int)]
	}).(DependabotSecretOutput)
}

type DependabotSecretMapOutput struct{ *pulumi.OutputState }

func (DependabotSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DependabotSecret)(nil)).Elem()
}

func (o DependabotSecretMapOutput) ToDependabotSecretMapOutput() DependabotSecretMapOutput {
	return o
}

func (o DependabotSecretMapOutput) ToDependabotSecretMapOutputWithContext(ctx context.Context) DependabotSecretMapOutput {
	return o
}

func (o DependabotSecretMapOutput) MapIndex(k pulumi.StringInput) DependabotSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DependabotSecret {
		return vs[0].(map[string]*DependabotSecret)[vs[1].(string)]
	}).(DependabotSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DependabotSecretInput)(nil)).Elem(), &DependabotSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependabotSecretArrayInput)(nil)).Elem(), DependabotSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependabotSecretMapInput)(nil)).Elem(), DependabotSecretMap{})
	pulumi.RegisterOutputType(DependabotSecretOutput{})
	pulumi.RegisterOutputType(DependabotSecretArrayOutput{})
	pulumi.RegisterOutputType(DependabotSecretMapOutput{})
}
