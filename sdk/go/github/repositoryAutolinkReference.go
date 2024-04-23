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

// This resource allows you to create and manage an autolink reference for a single repository.
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
//			repo, err := github.NewRepository(ctx, "repo", &github.RepositoryArgs{
//				Name:        pulumi.String("my-repo"),
//				Description: pulumi.String("GitHub repo managed by Terraform"),
//				Private:     pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryAutolinkReference(ctx, "autolink", &github.RepositoryAutolinkReferenceArgs{
//				Repository:        repo.Name,
//				KeyPrefix:         pulumi.String("TICKET-"),
//				TargetUrlTemplate: pulumi.String("https://example.com/TICKET?query=<num>"),
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
// ### Import by key prefix
//
// ```sh
// $ pulumi import github:index/repositoryAutolinkReference:RepositoryAutolinkReference auto oof/OOF-
// ```
type RepositoryAutolinkReference struct {
	pulumi.CustomResourceState

	// An etag representing the autolink reference object.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	IsAlphanumeric pulumi.BoolPtrOutput `pulumi:"isAlphanumeric"`
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	KeyPrefix pulumi.StringOutput `pulumi:"keyPrefix"`
	// The repository of the autolink reference.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	TargetUrlTemplate pulumi.StringOutput `pulumi:"targetUrlTemplate"`
}

// NewRepositoryAutolinkReference registers a new resource with the given unique name, arguments, and options.
func NewRepositoryAutolinkReference(ctx *pulumi.Context,
	name string, args *RepositoryAutolinkReferenceArgs, opts ...pulumi.ResourceOption) (*RepositoryAutolinkReference, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyPrefix == nil {
		return nil, errors.New("invalid value for required argument 'KeyPrefix'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.TargetUrlTemplate == nil {
		return nil, errors.New("invalid value for required argument 'TargetUrlTemplate'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryAutolinkReference
	err := ctx.RegisterResource("github:index/repositoryAutolinkReference:RepositoryAutolinkReference", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryAutolinkReference gets an existing RepositoryAutolinkReference resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryAutolinkReference(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryAutolinkReferenceState, opts ...pulumi.ResourceOption) (*RepositoryAutolinkReference, error) {
	var resource RepositoryAutolinkReference
	err := ctx.ReadResource("github:index/repositoryAutolinkReference:RepositoryAutolinkReference", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryAutolinkReference resources.
type repositoryAutolinkReferenceState struct {
	// An etag representing the autolink reference object.
	Etag *string `pulumi:"etag"`
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	IsAlphanumeric *bool `pulumi:"isAlphanumeric"`
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	KeyPrefix *string `pulumi:"keyPrefix"`
	// The repository of the autolink reference.
	Repository *string `pulumi:"repository"`
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	TargetUrlTemplate *string `pulumi:"targetUrlTemplate"`
}

type RepositoryAutolinkReferenceState struct {
	// An etag representing the autolink reference object.
	Etag pulumi.StringPtrInput
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	IsAlphanumeric pulumi.BoolPtrInput
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	KeyPrefix pulumi.StringPtrInput
	// The repository of the autolink reference.
	Repository pulumi.StringPtrInput
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	TargetUrlTemplate pulumi.StringPtrInput
}

func (RepositoryAutolinkReferenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryAutolinkReferenceState)(nil)).Elem()
}

type repositoryAutolinkReferenceArgs struct {
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	IsAlphanumeric *bool `pulumi:"isAlphanumeric"`
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	KeyPrefix string `pulumi:"keyPrefix"`
	// The repository of the autolink reference.
	Repository string `pulumi:"repository"`
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	TargetUrlTemplate string `pulumi:"targetUrlTemplate"`
}

// The set of arguments for constructing a RepositoryAutolinkReference resource.
type RepositoryAutolinkReferenceArgs struct {
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	IsAlphanumeric pulumi.BoolPtrInput
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	KeyPrefix pulumi.StringInput
	// The repository of the autolink reference.
	Repository pulumi.StringInput
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	TargetUrlTemplate pulumi.StringInput
}

func (RepositoryAutolinkReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryAutolinkReferenceArgs)(nil)).Elem()
}

type RepositoryAutolinkReferenceInput interface {
	pulumi.Input

	ToRepositoryAutolinkReferenceOutput() RepositoryAutolinkReferenceOutput
	ToRepositoryAutolinkReferenceOutputWithContext(ctx context.Context) RepositoryAutolinkReferenceOutput
}

func (*RepositoryAutolinkReference) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryAutolinkReference)(nil)).Elem()
}

func (i *RepositoryAutolinkReference) ToRepositoryAutolinkReferenceOutput() RepositoryAutolinkReferenceOutput {
	return i.ToRepositoryAutolinkReferenceOutputWithContext(context.Background())
}

func (i *RepositoryAutolinkReference) ToRepositoryAutolinkReferenceOutputWithContext(ctx context.Context) RepositoryAutolinkReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryAutolinkReferenceOutput)
}

// RepositoryAutolinkReferenceArrayInput is an input type that accepts RepositoryAutolinkReferenceArray and RepositoryAutolinkReferenceArrayOutput values.
// You can construct a concrete instance of `RepositoryAutolinkReferenceArrayInput` via:
//
//	RepositoryAutolinkReferenceArray{ RepositoryAutolinkReferenceArgs{...} }
type RepositoryAutolinkReferenceArrayInput interface {
	pulumi.Input

	ToRepositoryAutolinkReferenceArrayOutput() RepositoryAutolinkReferenceArrayOutput
	ToRepositoryAutolinkReferenceArrayOutputWithContext(context.Context) RepositoryAutolinkReferenceArrayOutput
}

type RepositoryAutolinkReferenceArray []RepositoryAutolinkReferenceInput

func (RepositoryAutolinkReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryAutolinkReference)(nil)).Elem()
}

func (i RepositoryAutolinkReferenceArray) ToRepositoryAutolinkReferenceArrayOutput() RepositoryAutolinkReferenceArrayOutput {
	return i.ToRepositoryAutolinkReferenceArrayOutputWithContext(context.Background())
}

func (i RepositoryAutolinkReferenceArray) ToRepositoryAutolinkReferenceArrayOutputWithContext(ctx context.Context) RepositoryAutolinkReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryAutolinkReferenceArrayOutput)
}

// RepositoryAutolinkReferenceMapInput is an input type that accepts RepositoryAutolinkReferenceMap and RepositoryAutolinkReferenceMapOutput values.
// You can construct a concrete instance of `RepositoryAutolinkReferenceMapInput` via:
//
//	RepositoryAutolinkReferenceMap{ "key": RepositoryAutolinkReferenceArgs{...} }
type RepositoryAutolinkReferenceMapInput interface {
	pulumi.Input

	ToRepositoryAutolinkReferenceMapOutput() RepositoryAutolinkReferenceMapOutput
	ToRepositoryAutolinkReferenceMapOutputWithContext(context.Context) RepositoryAutolinkReferenceMapOutput
}

type RepositoryAutolinkReferenceMap map[string]RepositoryAutolinkReferenceInput

func (RepositoryAutolinkReferenceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryAutolinkReference)(nil)).Elem()
}

func (i RepositoryAutolinkReferenceMap) ToRepositoryAutolinkReferenceMapOutput() RepositoryAutolinkReferenceMapOutput {
	return i.ToRepositoryAutolinkReferenceMapOutputWithContext(context.Background())
}

func (i RepositoryAutolinkReferenceMap) ToRepositoryAutolinkReferenceMapOutputWithContext(ctx context.Context) RepositoryAutolinkReferenceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryAutolinkReferenceMapOutput)
}

type RepositoryAutolinkReferenceOutput struct{ *pulumi.OutputState }

func (RepositoryAutolinkReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryAutolinkReference)(nil)).Elem()
}

func (o RepositoryAutolinkReferenceOutput) ToRepositoryAutolinkReferenceOutput() RepositoryAutolinkReferenceOutput {
	return o
}

func (o RepositoryAutolinkReferenceOutput) ToRepositoryAutolinkReferenceOutputWithContext(ctx context.Context) RepositoryAutolinkReferenceOutput {
	return o
}

// An etag representing the autolink reference object.
func (o RepositoryAutolinkReferenceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryAutolinkReference) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
func (o RepositoryAutolinkReferenceOutput) IsAlphanumeric() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryAutolinkReference) pulumi.BoolPtrOutput { return v.IsAlphanumeric }).(pulumi.BoolPtrOutput)
}

// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
func (o RepositoryAutolinkReferenceOutput) KeyPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryAutolinkReference) pulumi.StringOutput { return v.KeyPrefix }).(pulumi.StringOutput)
}

// The repository of the autolink reference.
func (o RepositoryAutolinkReferenceOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryAutolinkReference) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
func (o RepositoryAutolinkReferenceOutput) TargetUrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryAutolinkReference) pulumi.StringOutput { return v.TargetUrlTemplate }).(pulumi.StringOutput)
}

type RepositoryAutolinkReferenceArrayOutput struct{ *pulumi.OutputState }

func (RepositoryAutolinkReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryAutolinkReference)(nil)).Elem()
}

func (o RepositoryAutolinkReferenceArrayOutput) ToRepositoryAutolinkReferenceArrayOutput() RepositoryAutolinkReferenceArrayOutput {
	return o
}

func (o RepositoryAutolinkReferenceArrayOutput) ToRepositoryAutolinkReferenceArrayOutputWithContext(ctx context.Context) RepositoryAutolinkReferenceArrayOutput {
	return o
}

func (o RepositoryAutolinkReferenceArrayOutput) Index(i pulumi.IntInput) RepositoryAutolinkReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryAutolinkReference {
		return vs[0].([]*RepositoryAutolinkReference)[vs[1].(int)]
	}).(RepositoryAutolinkReferenceOutput)
}

type RepositoryAutolinkReferenceMapOutput struct{ *pulumi.OutputState }

func (RepositoryAutolinkReferenceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryAutolinkReference)(nil)).Elem()
}

func (o RepositoryAutolinkReferenceMapOutput) ToRepositoryAutolinkReferenceMapOutput() RepositoryAutolinkReferenceMapOutput {
	return o
}

func (o RepositoryAutolinkReferenceMapOutput) ToRepositoryAutolinkReferenceMapOutputWithContext(ctx context.Context) RepositoryAutolinkReferenceMapOutput {
	return o
}

func (o RepositoryAutolinkReferenceMapOutput) MapIndex(k pulumi.StringInput) RepositoryAutolinkReferenceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryAutolinkReference {
		return vs[0].(map[string]*RepositoryAutolinkReference)[vs[1].(string)]
	}).(RepositoryAutolinkReferenceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryAutolinkReferenceInput)(nil)).Elem(), &RepositoryAutolinkReference{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryAutolinkReferenceArrayInput)(nil)).Elem(), RepositoryAutolinkReferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryAutolinkReferenceMapInput)(nil)).Elem(), RepositoryAutolinkReferenceMap{})
	pulumi.RegisterOutputType(RepositoryAutolinkReferenceOutput{})
	pulumi.RegisterOutputType(RepositoryAutolinkReferenceArrayOutput{})
	pulumi.RegisterOutputType(RepositoryAutolinkReferenceMapOutput{})
}
