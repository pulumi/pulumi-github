// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
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
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := github.NewDependabotOrganizationSecret(ctx, "exampleSecretDependabotOrganizationSecret", &github.DependabotOrganizationSecretArgs{
//				SecretName:     pulumi.String("example_secret_name"),
//				Visibility:     pulumi.String("private"),
//				PlaintextValue: pulumi.Any(_var.Some_secret_string),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewDependabotOrganizationSecret(ctx, "exampleSecretIndex/dependabotOrganizationSecretDependabotOrganizationSecret", &github.DependabotOrganizationSecretArgs{
//				SecretName:     pulumi.String("example_secret_name"),
//				Visibility:     pulumi.String("private"),
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
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			repo, err := github.LookupRepository(ctx, &github.LookupRepositoryArgs{
//				FullName: pulumi.StringRef("my-org/repo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewDependabotOrganizationSecret(ctx, "exampleSecretDependabotOrganizationSecret", &github.DependabotOrganizationSecretArgs{
//				SecretName:     pulumi.String("example_secret_name"),
//				Visibility:     pulumi.String("selected"),
//				PlaintextValue: pulumi.Any(_var.Some_secret_string),
//				SelectedRepositoryIds: pulumi.IntArray{
//					pulumi.Int(repo.RepoId),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewDependabotOrganizationSecret(ctx, "exampleSecretIndex/dependabotOrganizationSecretDependabotOrganizationSecret", &github.DependabotOrganizationSecretArgs{
//				SecretName:     pulumi.String("example_secret_name"),
//				Visibility:     pulumi.String("selected"),
//				EncryptedValue: pulumi.Any(_var.Some_encrypted_secret_string),
//				SelectedRepositoryIds: pulumi.IntArray{
//					pulumi.Int(repo.RepoId),
//				},
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
// This resource can be imported using an ID made up of the secret name:
//
// ```sh
// $ pulumi import github:index/dependabotOrganizationSecret:DependabotOrganizationSecret test_secret test_secret_name
// ```
//
// NOTE: the implementation is limited in that it won't fetch the value of the
//
// `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
type DependabotOrganizationSecret struct {
	pulumi.CustomResourceState

	// Date of dependabotSecret creation.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrOutput `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrOutput `pulumi:"plaintextValue"`
	// Name of the secret
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayOutput `pulumi:"selectedRepositoryIds"`
	// Date of dependabotSecret update.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewDependabotOrganizationSecret registers a new resource with the given unique name, arguments, and options.
func NewDependabotOrganizationSecret(ctx *pulumi.Context,
	name string, args *DependabotOrganizationSecretArgs, opts ...pulumi.ResourceOption) (*DependabotOrganizationSecret, error) {
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
	var resource DependabotOrganizationSecret
	err := ctx.RegisterResource("github:index/dependabotOrganizationSecret:DependabotOrganizationSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDependabotOrganizationSecret gets an existing DependabotOrganizationSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDependabotOrganizationSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DependabotOrganizationSecretState, opts ...pulumi.ResourceOption) (*DependabotOrganizationSecret, error) {
	var resource DependabotOrganizationSecret
	err := ctx.ReadResource("github:index/dependabotOrganizationSecret:DependabotOrganizationSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DependabotOrganizationSecret resources.
type dependabotOrganizationSecretState struct {
	// Date of dependabotSecret creation.
	CreatedAt *string `pulumi:"createdAt"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue *string `pulumi:"encryptedValue"`
	// Plaintext value of the secret to be encrypted
	PlaintextValue *string `pulumi:"plaintextValue"`
	// Name of the secret
	SecretName *string `pulumi:"secretName"`
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
	// Date of dependabotSecret update.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility *string `pulumi:"visibility"`
}

type DependabotOrganizationSecretState struct {
	// Date of dependabotSecret creation.
	CreatedAt pulumi.StringPtrInput
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	EncryptedValue pulumi.StringPtrInput
	// Plaintext value of the secret to be encrypted
	PlaintextValue pulumi.StringPtrInput
	// Name of the secret
	SecretName pulumi.StringPtrInput
	// An array of repository ids that can access the organization secret.
	SelectedRepositoryIds pulumi.IntArrayInput
	// Date of dependabotSecret update.
	UpdatedAt pulumi.StringPtrInput
	// Configures the access that repositories have to the organization secret.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringPtrInput
}

func (DependabotOrganizationSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*dependabotOrganizationSecretState)(nil)).Elem()
}

type dependabotOrganizationSecretArgs struct {
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

// The set of arguments for constructing a DependabotOrganizationSecret resource.
type DependabotOrganizationSecretArgs struct {
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

func (DependabotOrganizationSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dependabotOrganizationSecretArgs)(nil)).Elem()
}

type DependabotOrganizationSecretInput interface {
	pulumi.Input

	ToDependabotOrganizationSecretOutput() DependabotOrganizationSecretOutput
	ToDependabotOrganizationSecretOutputWithContext(ctx context.Context) DependabotOrganizationSecretOutput
}

func (*DependabotOrganizationSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**DependabotOrganizationSecret)(nil)).Elem()
}

func (i *DependabotOrganizationSecret) ToDependabotOrganizationSecretOutput() DependabotOrganizationSecretOutput {
	return i.ToDependabotOrganizationSecretOutputWithContext(context.Background())
}

func (i *DependabotOrganizationSecret) ToDependabotOrganizationSecretOutputWithContext(ctx context.Context) DependabotOrganizationSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependabotOrganizationSecretOutput)
}

// DependabotOrganizationSecretArrayInput is an input type that accepts DependabotOrganizationSecretArray and DependabotOrganizationSecretArrayOutput values.
// You can construct a concrete instance of `DependabotOrganizationSecretArrayInput` via:
//
//	DependabotOrganizationSecretArray{ DependabotOrganizationSecretArgs{...} }
type DependabotOrganizationSecretArrayInput interface {
	pulumi.Input

	ToDependabotOrganizationSecretArrayOutput() DependabotOrganizationSecretArrayOutput
	ToDependabotOrganizationSecretArrayOutputWithContext(context.Context) DependabotOrganizationSecretArrayOutput
}

type DependabotOrganizationSecretArray []DependabotOrganizationSecretInput

func (DependabotOrganizationSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DependabotOrganizationSecret)(nil)).Elem()
}

func (i DependabotOrganizationSecretArray) ToDependabotOrganizationSecretArrayOutput() DependabotOrganizationSecretArrayOutput {
	return i.ToDependabotOrganizationSecretArrayOutputWithContext(context.Background())
}

func (i DependabotOrganizationSecretArray) ToDependabotOrganizationSecretArrayOutputWithContext(ctx context.Context) DependabotOrganizationSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependabotOrganizationSecretArrayOutput)
}

// DependabotOrganizationSecretMapInput is an input type that accepts DependabotOrganizationSecretMap and DependabotOrganizationSecretMapOutput values.
// You can construct a concrete instance of `DependabotOrganizationSecretMapInput` via:
//
//	DependabotOrganizationSecretMap{ "key": DependabotOrganizationSecretArgs{...} }
type DependabotOrganizationSecretMapInput interface {
	pulumi.Input

	ToDependabotOrganizationSecretMapOutput() DependabotOrganizationSecretMapOutput
	ToDependabotOrganizationSecretMapOutputWithContext(context.Context) DependabotOrganizationSecretMapOutput
}

type DependabotOrganizationSecretMap map[string]DependabotOrganizationSecretInput

func (DependabotOrganizationSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DependabotOrganizationSecret)(nil)).Elem()
}

func (i DependabotOrganizationSecretMap) ToDependabotOrganizationSecretMapOutput() DependabotOrganizationSecretMapOutput {
	return i.ToDependabotOrganizationSecretMapOutputWithContext(context.Background())
}

func (i DependabotOrganizationSecretMap) ToDependabotOrganizationSecretMapOutputWithContext(ctx context.Context) DependabotOrganizationSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependabotOrganizationSecretMapOutput)
}

type DependabotOrganizationSecretOutput struct{ *pulumi.OutputState }

func (DependabotOrganizationSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DependabotOrganizationSecret)(nil)).Elem()
}

func (o DependabotOrganizationSecretOutput) ToDependabotOrganizationSecretOutput() DependabotOrganizationSecretOutput {
	return o
}

func (o DependabotOrganizationSecretOutput) ToDependabotOrganizationSecretOutputWithContext(ctx context.Context) DependabotOrganizationSecretOutput {
	return o
}

// Date of dependabotSecret creation.
func (o DependabotOrganizationSecretOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotOrganizationSecret) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Encrypted value of the secret using the GitHub public key in Base64 format.
func (o DependabotOrganizationSecretOutput) EncryptedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DependabotOrganizationSecret) pulumi.StringPtrOutput { return v.EncryptedValue }).(pulumi.StringPtrOutput)
}

// Plaintext value of the secret to be encrypted
func (o DependabotOrganizationSecretOutput) PlaintextValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DependabotOrganizationSecret) pulumi.StringPtrOutput { return v.PlaintextValue }).(pulumi.StringPtrOutput)
}

// Name of the secret
func (o DependabotOrganizationSecretOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotOrganizationSecret) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// An array of repository ids that can access the organization secret.
func (o DependabotOrganizationSecretOutput) SelectedRepositoryIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *DependabotOrganizationSecret) pulumi.IntArrayOutput { return v.SelectedRepositoryIds }).(pulumi.IntArrayOutput)
}

// Date of dependabotSecret update.
func (o DependabotOrganizationSecretOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotOrganizationSecret) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Configures the access that repositories have to the organization secret.
// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
func (o DependabotOrganizationSecretOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *DependabotOrganizationSecret) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type DependabotOrganizationSecretArrayOutput struct{ *pulumi.OutputState }

func (DependabotOrganizationSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DependabotOrganizationSecret)(nil)).Elem()
}

func (o DependabotOrganizationSecretArrayOutput) ToDependabotOrganizationSecretArrayOutput() DependabotOrganizationSecretArrayOutput {
	return o
}

func (o DependabotOrganizationSecretArrayOutput) ToDependabotOrganizationSecretArrayOutputWithContext(ctx context.Context) DependabotOrganizationSecretArrayOutput {
	return o
}

func (o DependabotOrganizationSecretArrayOutput) Index(i pulumi.IntInput) DependabotOrganizationSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DependabotOrganizationSecret {
		return vs[0].([]*DependabotOrganizationSecret)[vs[1].(int)]
	}).(DependabotOrganizationSecretOutput)
}

type DependabotOrganizationSecretMapOutput struct{ *pulumi.OutputState }

func (DependabotOrganizationSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DependabotOrganizationSecret)(nil)).Elem()
}

func (o DependabotOrganizationSecretMapOutput) ToDependabotOrganizationSecretMapOutput() DependabotOrganizationSecretMapOutput {
	return o
}

func (o DependabotOrganizationSecretMapOutput) ToDependabotOrganizationSecretMapOutputWithContext(ctx context.Context) DependabotOrganizationSecretMapOutput {
	return o
}

func (o DependabotOrganizationSecretMapOutput) MapIndex(k pulumi.StringInput) DependabotOrganizationSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DependabotOrganizationSecret {
		return vs[0].(map[string]*DependabotOrganizationSecret)[vs[1].(string)]
	}).(DependabotOrganizationSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DependabotOrganizationSecretInput)(nil)).Elem(), &DependabotOrganizationSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependabotOrganizationSecretArrayInput)(nil)).Elem(), DependabotOrganizationSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependabotOrganizationSecretMapInput)(nil)).Elem(), DependabotOrganizationSecretMap{})
	pulumi.RegisterOutputType(DependabotOrganizationSecretOutput{})
	pulumi.RegisterOutputType(DependabotOrganizationSecretArrayOutput{})
	pulumi.RegisterOutputType(DependabotOrganizationSecretMapOutput{})
}
