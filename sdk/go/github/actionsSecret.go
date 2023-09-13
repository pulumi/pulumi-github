// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := github.GetActionsPublicKey(ctx, &github.GetActionsPublicKeyArgs{
//				Repository: "example_repository",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewActionsSecret(ctx, "exampleSecretActionsSecret", &github.ActionsSecretArgs{
//				Repository:     pulumi.String("example_repository"),
//				SecretName:     pulumi.String("example_secret_name"),
//				PlaintextValue: pulumi.Any(_var.Some_secret_string),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewActionsSecret(ctx, "exampleSecretIndex/actionsSecretActionsSecret", &github.ActionsSecretArgs{
//				Repository:     pulumi.String("example_repository"),
//				SecretName:     pulumi.String("example_secret_name"),
//				EncryptedValue: pulumi.Any(_var.Some_encrypted_secret_string),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// This resource can be imported using an ID made up of the `repository` and `secret_name`:
//
// ```sh
//
//	$ pulumi import github:index/actionsSecret:ActionsSecret example_secret <repository>/<secret_name>
//
// ```
//
//	NOTEthe implementation is limited in that it won't fetch the value of the `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
type ActionsSecret struct {
	pulumi.CustomResourceState

	// Date of actionsSecret creation.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrOutput `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrOutput `pulumi:"plaintextValue"`
	// Name of the repository
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Name of the secret
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// Date of actionsSecret update.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewActionsSecret registers a new resource with the given unique name, arguments, and options.
func NewActionsSecret(ctx *pulumi.Context,
	name string, args *ActionsSecretArgs, opts ...pulumi.ResourceOption) (*ActionsSecret, error) {
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ActionsSecret
	err := ctx.RegisterResource("github:index/actionsSecret:ActionsSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsSecret gets an existing ActionsSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsSecretState, opts ...pulumi.ResourceOption) (*ActionsSecret, error) {
	var resource ActionsSecret
	err := ctx.ReadResource("github:index/actionsSecret:ActionsSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsSecret resources.
type actionsSecretState struct {
	// Date of actionsSecret creation.
	CreatedAt *string `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue *string `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue *string `pulumi:"plaintextValue"`
	// Name of the repository
	Repository *string `pulumi:"repository"`
	// Name of the secret
	SecretName *string `pulumi:"secretName"`
	// Date of actionsSecret update.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ActionsSecretState struct {
	// Date of actionsSecret creation.
	CreatedAt pulumi.StringPtrInput
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrInput
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrInput
	// Name of the repository
	Repository pulumi.StringPtrInput
	// Name of the secret
	SecretName pulumi.StringPtrInput
	// Date of actionsSecret update.
	UpdatedAt pulumi.StringPtrInput
}

func (ActionsSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsSecretState)(nil)).Elem()
}

type actionsSecretArgs struct {
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue *string `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue *string `pulumi:"plaintextValue"`
	// Name of the repository
	Repository string `pulumi:"repository"`
	// Name of the secret
	SecretName string `pulumi:"secretName"`
}

// The set of arguments for constructing a ActionsSecret resource.
type ActionsSecretArgs struct {
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrInput
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrInput
	// Name of the repository
	Repository pulumi.StringInput
	// Name of the secret
	SecretName pulumi.StringInput
}

func (ActionsSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsSecretArgs)(nil)).Elem()
}

type ActionsSecretInput interface {
	pulumi.Input

	ToActionsSecretOutput() ActionsSecretOutput
	ToActionsSecretOutputWithContext(ctx context.Context) ActionsSecretOutput
}

func (*ActionsSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsSecret)(nil)).Elem()
}

func (i *ActionsSecret) ToActionsSecretOutput() ActionsSecretOutput {
	return i.ToActionsSecretOutputWithContext(context.Background())
}

func (i *ActionsSecret) ToActionsSecretOutputWithContext(ctx context.Context) ActionsSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsSecretOutput)
}

func (i *ActionsSecret) ToOutput(ctx context.Context) pulumix.Output[*ActionsSecret] {
	return pulumix.Output[*ActionsSecret]{
		OutputState: i.ToActionsSecretOutputWithContext(ctx).OutputState,
	}
}

// ActionsSecretArrayInput is an input type that accepts ActionsSecretArray and ActionsSecretArrayOutput values.
// You can construct a concrete instance of `ActionsSecretArrayInput` via:
//
//	ActionsSecretArray{ ActionsSecretArgs{...} }
type ActionsSecretArrayInput interface {
	pulumi.Input

	ToActionsSecretArrayOutput() ActionsSecretArrayOutput
	ToActionsSecretArrayOutputWithContext(context.Context) ActionsSecretArrayOutput
}

type ActionsSecretArray []ActionsSecretInput

func (ActionsSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsSecret)(nil)).Elem()
}

func (i ActionsSecretArray) ToActionsSecretArrayOutput() ActionsSecretArrayOutput {
	return i.ToActionsSecretArrayOutputWithContext(context.Background())
}

func (i ActionsSecretArray) ToActionsSecretArrayOutputWithContext(ctx context.Context) ActionsSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsSecretArrayOutput)
}

func (i ActionsSecretArray) ToOutput(ctx context.Context) pulumix.Output[[]*ActionsSecret] {
	return pulumix.Output[[]*ActionsSecret]{
		OutputState: i.ToActionsSecretArrayOutputWithContext(ctx).OutputState,
	}
}

// ActionsSecretMapInput is an input type that accepts ActionsSecretMap and ActionsSecretMapOutput values.
// You can construct a concrete instance of `ActionsSecretMapInput` via:
//
//	ActionsSecretMap{ "key": ActionsSecretArgs{...} }
type ActionsSecretMapInput interface {
	pulumi.Input

	ToActionsSecretMapOutput() ActionsSecretMapOutput
	ToActionsSecretMapOutputWithContext(context.Context) ActionsSecretMapOutput
}

type ActionsSecretMap map[string]ActionsSecretInput

func (ActionsSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsSecret)(nil)).Elem()
}

func (i ActionsSecretMap) ToActionsSecretMapOutput() ActionsSecretMapOutput {
	return i.ToActionsSecretMapOutputWithContext(context.Background())
}

func (i ActionsSecretMap) ToActionsSecretMapOutputWithContext(ctx context.Context) ActionsSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsSecretMapOutput)
}

func (i ActionsSecretMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ActionsSecret] {
	return pulumix.Output[map[string]*ActionsSecret]{
		OutputState: i.ToActionsSecretMapOutputWithContext(ctx).OutputState,
	}
}

type ActionsSecretOutput struct{ *pulumi.OutputState }

func (ActionsSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsSecret)(nil)).Elem()
}

func (o ActionsSecretOutput) ToActionsSecretOutput() ActionsSecretOutput {
	return o
}

func (o ActionsSecretOutput) ToActionsSecretOutputWithContext(ctx context.Context) ActionsSecretOutput {
	return o
}

func (o ActionsSecretOutput) ToOutput(ctx context.Context) pulumix.Output[*ActionsSecret] {
	return pulumix.Output[*ActionsSecret]{
		OutputState: o.OutputState,
	}
}

// Date of actionsSecret creation.
func (o ActionsSecretOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsSecret) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Encrypted value of the secret using the GitHub public key in Base64 format.
func (o ActionsSecretOutput) EncryptedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionsSecret) pulumi.StringPtrOutput { return v.EncryptedValue }).(pulumi.StringPtrOutput)
}

// Plaintext value of the secret to be encrypted
func (o ActionsSecretOutput) PlaintextValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionsSecret) pulumi.StringPtrOutput { return v.PlaintextValue }).(pulumi.StringPtrOutput)
}

// Name of the repository
func (o ActionsSecretOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsSecret) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// Name of the secret
func (o ActionsSecretOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsSecret) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// Date of actionsSecret update.
func (o ActionsSecretOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsSecret) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ActionsSecretArrayOutput struct{ *pulumi.OutputState }

func (ActionsSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsSecret)(nil)).Elem()
}

func (o ActionsSecretArrayOutput) ToActionsSecretArrayOutput() ActionsSecretArrayOutput {
	return o
}

func (o ActionsSecretArrayOutput) ToActionsSecretArrayOutputWithContext(ctx context.Context) ActionsSecretArrayOutput {
	return o
}

func (o ActionsSecretArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ActionsSecret] {
	return pulumix.Output[[]*ActionsSecret]{
		OutputState: o.OutputState,
	}
}

func (o ActionsSecretArrayOutput) Index(i pulumi.IntInput) ActionsSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsSecret {
		return vs[0].([]*ActionsSecret)[vs[1].(int)]
	}).(ActionsSecretOutput)
}

type ActionsSecretMapOutput struct{ *pulumi.OutputState }

func (ActionsSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsSecret)(nil)).Elem()
}

func (o ActionsSecretMapOutput) ToActionsSecretMapOutput() ActionsSecretMapOutput {
	return o
}

func (o ActionsSecretMapOutput) ToActionsSecretMapOutputWithContext(ctx context.Context) ActionsSecretMapOutput {
	return o
}

func (o ActionsSecretMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ActionsSecret] {
	return pulumix.Output[map[string]*ActionsSecret]{
		OutputState: o.OutputState,
	}
}

func (o ActionsSecretMapOutput) MapIndex(k pulumi.StringInput) ActionsSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsSecret {
		return vs[0].(map[string]*ActionsSecret)[vs[1].(string)]
	}).(ActionsSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsSecretInput)(nil)).Elem(), &ActionsSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsSecretArrayInput)(nil)).Elem(), ActionsSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsSecretMapInput)(nil)).Elem(), ActionsSecretMap{})
	pulumi.RegisterOutputType(ActionsSecretOutput{})
	pulumi.RegisterOutputType(ActionsSecretArrayOutput{})
	pulumi.RegisterOutputType(ActionsSecretMapOutput{})
}
