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

// This resource allows you to create and manage a repository tag protection for repositories within your GitHub organization or personal account.
//
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
//			_, err := github.NewRepositoryTagProtection(ctx, "example", &github.RepositoryTagProtectionArgs{
//				Pattern:    pulumi.String("v*"),
//				Repository: pulumi.String("example-repository"),
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
// Repository tag protections can be imported using the `name` of the repository, combined with the `id` of the tag protection, separated by a `/` character.
// The `id` of the tag protection can be found using the [GitHub API](https://docs.github.com/en/rest/repos/tags#list-tag-protection-states-for-a-repository).
//
// Importing uses the name of the repository, as well as the ID of the tag protection, e.g.
//
// ```sh
// $ pulumi import github:index/repositoryTagProtection:RepositoryTagProtection terraform my-repo/31077
// ```
type RepositoryTagProtection struct {
	pulumi.CustomResourceState

	// The pattern of the tag to protect.
	Pattern pulumi.StringOutput `pulumi:"pattern"`
	// Name of the repository to add the tag protection to.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The ID of the tag protection.
	TagProtectionId pulumi.IntOutput `pulumi:"tagProtectionId"`
}

// NewRepositoryTagProtection registers a new resource with the given unique name, arguments, and options.
func NewRepositoryTagProtection(ctx *pulumi.Context,
	name string, args *RepositoryTagProtectionArgs, opts ...pulumi.ResourceOption) (*RepositoryTagProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Pattern == nil {
		return nil, errors.New("invalid value for required argument 'Pattern'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryTagProtection
	err := ctx.RegisterResource("github:index/repositoryTagProtection:RepositoryTagProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryTagProtection gets an existing RepositoryTagProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryTagProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryTagProtectionState, opts ...pulumi.ResourceOption) (*RepositoryTagProtection, error) {
	var resource RepositoryTagProtection
	err := ctx.ReadResource("github:index/repositoryTagProtection:RepositoryTagProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryTagProtection resources.
type repositoryTagProtectionState struct {
	// The pattern of the tag to protect.
	Pattern *string `pulumi:"pattern"`
	// Name of the repository to add the tag protection to.
	Repository *string `pulumi:"repository"`
	// The ID of the tag protection.
	TagProtectionId *int `pulumi:"tagProtectionId"`
}

type RepositoryTagProtectionState struct {
	// The pattern of the tag to protect.
	Pattern pulumi.StringPtrInput
	// Name of the repository to add the tag protection to.
	Repository pulumi.StringPtrInput
	// The ID of the tag protection.
	TagProtectionId pulumi.IntPtrInput
}

func (RepositoryTagProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryTagProtectionState)(nil)).Elem()
}

type repositoryTagProtectionArgs struct {
	// The pattern of the tag to protect.
	Pattern string `pulumi:"pattern"`
	// Name of the repository to add the tag protection to.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryTagProtection resource.
type RepositoryTagProtectionArgs struct {
	// The pattern of the tag to protect.
	Pattern pulumi.StringInput
	// Name of the repository to add the tag protection to.
	Repository pulumi.StringInput
}

func (RepositoryTagProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryTagProtectionArgs)(nil)).Elem()
}

type RepositoryTagProtectionInput interface {
	pulumi.Input

	ToRepositoryTagProtectionOutput() RepositoryTagProtectionOutput
	ToRepositoryTagProtectionOutputWithContext(ctx context.Context) RepositoryTagProtectionOutput
}

func (*RepositoryTagProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryTagProtection)(nil)).Elem()
}

func (i *RepositoryTagProtection) ToRepositoryTagProtectionOutput() RepositoryTagProtectionOutput {
	return i.ToRepositoryTagProtectionOutputWithContext(context.Background())
}

func (i *RepositoryTagProtection) ToRepositoryTagProtectionOutputWithContext(ctx context.Context) RepositoryTagProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryTagProtectionOutput)
}

// RepositoryTagProtectionArrayInput is an input type that accepts RepositoryTagProtectionArray and RepositoryTagProtectionArrayOutput values.
// You can construct a concrete instance of `RepositoryTagProtectionArrayInput` via:
//
//	RepositoryTagProtectionArray{ RepositoryTagProtectionArgs{...} }
type RepositoryTagProtectionArrayInput interface {
	pulumi.Input

	ToRepositoryTagProtectionArrayOutput() RepositoryTagProtectionArrayOutput
	ToRepositoryTagProtectionArrayOutputWithContext(context.Context) RepositoryTagProtectionArrayOutput
}

type RepositoryTagProtectionArray []RepositoryTagProtectionInput

func (RepositoryTagProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryTagProtection)(nil)).Elem()
}

func (i RepositoryTagProtectionArray) ToRepositoryTagProtectionArrayOutput() RepositoryTagProtectionArrayOutput {
	return i.ToRepositoryTagProtectionArrayOutputWithContext(context.Background())
}

func (i RepositoryTagProtectionArray) ToRepositoryTagProtectionArrayOutputWithContext(ctx context.Context) RepositoryTagProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryTagProtectionArrayOutput)
}

// RepositoryTagProtectionMapInput is an input type that accepts RepositoryTagProtectionMap and RepositoryTagProtectionMapOutput values.
// You can construct a concrete instance of `RepositoryTagProtectionMapInput` via:
//
//	RepositoryTagProtectionMap{ "key": RepositoryTagProtectionArgs{...} }
type RepositoryTagProtectionMapInput interface {
	pulumi.Input

	ToRepositoryTagProtectionMapOutput() RepositoryTagProtectionMapOutput
	ToRepositoryTagProtectionMapOutputWithContext(context.Context) RepositoryTagProtectionMapOutput
}

type RepositoryTagProtectionMap map[string]RepositoryTagProtectionInput

func (RepositoryTagProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryTagProtection)(nil)).Elem()
}

func (i RepositoryTagProtectionMap) ToRepositoryTagProtectionMapOutput() RepositoryTagProtectionMapOutput {
	return i.ToRepositoryTagProtectionMapOutputWithContext(context.Background())
}

func (i RepositoryTagProtectionMap) ToRepositoryTagProtectionMapOutputWithContext(ctx context.Context) RepositoryTagProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryTagProtectionMapOutput)
}

type RepositoryTagProtectionOutput struct{ *pulumi.OutputState }

func (RepositoryTagProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryTagProtection)(nil)).Elem()
}

func (o RepositoryTagProtectionOutput) ToRepositoryTagProtectionOutput() RepositoryTagProtectionOutput {
	return o
}

func (o RepositoryTagProtectionOutput) ToRepositoryTagProtectionOutputWithContext(ctx context.Context) RepositoryTagProtectionOutput {
	return o
}

// The pattern of the tag to protect.
func (o RepositoryTagProtectionOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryTagProtection) pulumi.StringOutput { return v.Pattern }).(pulumi.StringOutput)
}

// Name of the repository to add the tag protection to.
func (o RepositoryTagProtectionOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryTagProtection) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The ID of the tag protection.
func (o RepositoryTagProtectionOutput) TagProtectionId() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryTagProtection) pulumi.IntOutput { return v.TagProtectionId }).(pulumi.IntOutput)
}

type RepositoryTagProtectionArrayOutput struct{ *pulumi.OutputState }

func (RepositoryTagProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryTagProtection)(nil)).Elem()
}

func (o RepositoryTagProtectionArrayOutput) ToRepositoryTagProtectionArrayOutput() RepositoryTagProtectionArrayOutput {
	return o
}

func (o RepositoryTagProtectionArrayOutput) ToRepositoryTagProtectionArrayOutputWithContext(ctx context.Context) RepositoryTagProtectionArrayOutput {
	return o
}

func (o RepositoryTagProtectionArrayOutput) Index(i pulumi.IntInput) RepositoryTagProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryTagProtection {
		return vs[0].([]*RepositoryTagProtection)[vs[1].(int)]
	}).(RepositoryTagProtectionOutput)
}

type RepositoryTagProtectionMapOutput struct{ *pulumi.OutputState }

func (RepositoryTagProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryTagProtection)(nil)).Elem()
}

func (o RepositoryTagProtectionMapOutput) ToRepositoryTagProtectionMapOutput() RepositoryTagProtectionMapOutput {
	return o
}

func (o RepositoryTagProtectionMapOutput) ToRepositoryTagProtectionMapOutputWithContext(ctx context.Context) RepositoryTagProtectionMapOutput {
	return o
}

func (o RepositoryTagProtectionMapOutput) MapIndex(k pulumi.StringInput) RepositoryTagProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryTagProtection {
		return vs[0].(map[string]*RepositoryTagProtection)[vs[1].(string)]
	}).(RepositoryTagProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryTagProtectionInput)(nil)).Elem(), &RepositoryTagProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryTagProtectionArrayInput)(nil)).Elem(), RepositoryTagProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryTagProtectionMapInput)(nil)).Elem(), RepositoryTagProtectionMap{})
	pulumi.RegisterOutputType(RepositoryTagProtectionOutput{})
	pulumi.RegisterOutputType(RepositoryTagProtectionArrayOutput{})
	pulumi.RegisterOutputType(RepositoryTagProtectionMapOutput{})
}
