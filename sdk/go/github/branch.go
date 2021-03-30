// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource allows you to create and manage branches within your repository.
//
// Additional constraints can be applied to ensure your branch is created from
// another branch or commit.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v3/go/github"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.NewBranch(ctx, "development", &github.BranchArgs{
// 			Branch:     pulumi.String("development"),
// 			Repository: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// GitHub Branch can be imported using an ID made up of `repository:branch`, e.g.
//
// ```sh
//  $ pulumi import github:index/branch:Branch terraform terraform:master
// ```
//
//  Optionally, a source branch may be specified using an ID of `repository:branch:source_branch`. This is useful for importing branches that do not branch directly off master.
//
// ```sh
//  $ pulumi import github:index/branch:Branch terraform terraform:feature-branch:dev
// ```
type Branch struct {
	pulumi.CustomResourceState

	// The repository branch to create.
	Branch pulumi.StringOutput `pulumi:"branch"`
	// An etag representing the Branch object.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A string representing a branch reference, in the form of `refs/heads/<branch>`.
	Ref pulumi.StringOutput `pulumi:"ref"`
	// The GitHub repository name.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// A string storing the reference's `HEAD` commit's SHA1.
	Sha pulumi.StringOutput `pulumi:"sha"`
	// The branch name to start from. Defaults to `main`.
	SourceBranch pulumi.StringPtrOutput `pulumi:"sourceBranch"`
	// The commit hash to start from. Defaults to the tip of `sourceBranch`. If provided, `sourceBranch` is ignored.
	SourceSha pulumi.StringOutput `pulumi:"sourceSha"`
}

// NewBranch registers a new resource with the given unique name, arguments, and options.
func NewBranch(ctx *pulumi.Context,
	name string, args *BranchArgs, opts ...pulumi.ResourceOption) (*Branch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Branch == nil {
		return nil, errors.New("invalid value for required argument 'Branch'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource Branch
	err := ctx.RegisterResource("github:index/branch:Branch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranch gets an existing Branch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchState, opts ...pulumi.ResourceOption) (*Branch, error) {
	var resource Branch
	err := ctx.ReadResource("github:index/branch:Branch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Branch resources.
type branchState struct {
	// The repository branch to create.
	Branch *string `pulumi:"branch"`
	// An etag representing the Branch object.
	Etag *string `pulumi:"etag"`
	// A string representing a branch reference, in the form of `refs/heads/<branch>`.
	Ref *string `pulumi:"ref"`
	// The GitHub repository name.
	Repository *string `pulumi:"repository"`
	// A string storing the reference's `HEAD` commit's SHA1.
	Sha *string `pulumi:"sha"`
	// The branch name to start from. Defaults to `main`.
	SourceBranch *string `pulumi:"sourceBranch"`
	// The commit hash to start from. Defaults to the tip of `sourceBranch`. If provided, `sourceBranch` is ignored.
	SourceSha *string `pulumi:"sourceSha"`
}

type BranchState struct {
	// The repository branch to create.
	Branch pulumi.StringPtrInput
	// An etag representing the Branch object.
	Etag pulumi.StringPtrInput
	// A string representing a branch reference, in the form of `refs/heads/<branch>`.
	Ref pulumi.StringPtrInput
	// The GitHub repository name.
	Repository pulumi.StringPtrInput
	// A string storing the reference's `HEAD` commit's SHA1.
	Sha pulumi.StringPtrInput
	// The branch name to start from. Defaults to `main`.
	SourceBranch pulumi.StringPtrInput
	// The commit hash to start from. Defaults to the tip of `sourceBranch`. If provided, `sourceBranch` is ignored.
	SourceSha pulumi.StringPtrInput
}

func (BranchState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchState)(nil)).Elem()
}

type branchArgs struct {
	// The repository branch to create.
	Branch string `pulumi:"branch"`
	// The GitHub repository name.
	Repository string `pulumi:"repository"`
	// The branch name to start from. Defaults to `main`.
	SourceBranch *string `pulumi:"sourceBranch"`
	// The commit hash to start from. Defaults to the tip of `sourceBranch`. If provided, `sourceBranch` is ignored.
	SourceSha *string `pulumi:"sourceSha"`
}

// The set of arguments for constructing a Branch resource.
type BranchArgs struct {
	// The repository branch to create.
	Branch pulumi.StringInput
	// The GitHub repository name.
	Repository pulumi.StringInput
	// The branch name to start from. Defaults to `main`.
	SourceBranch pulumi.StringPtrInput
	// The commit hash to start from. Defaults to the tip of `sourceBranch`. If provided, `sourceBranch` is ignored.
	SourceSha pulumi.StringPtrInput
}

func (BranchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchArgs)(nil)).Elem()
}

type BranchInput interface {
	pulumi.Input

	ToBranchOutput() BranchOutput
	ToBranchOutputWithContext(ctx context.Context) BranchOutput
}

func (*Branch) ElementType() reflect.Type {
	return reflect.TypeOf((*Branch)(nil))
}

func (i *Branch) ToBranchOutput() BranchOutput {
	return i.ToBranchOutputWithContext(context.Background())
}

func (i *Branch) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchOutput)
}

func (i *Branch) ToBranchPtrOutput() BranchPtrOutput {
	return i.ToBranchPtrOutputWithContext(context.Background())
}

func (i *Branch) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPtrOutput)
}

type BranchPtrInput interface {
	pulumi.Input

	ToBranchPtrOutput() BranchPtrOutput
	ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput
}

type branchPtrType BranchArgs

func (*branchPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil))
}

func (i *branchPtrType) ToBranchPtrOutput() BranchPtrOutput {
	return i.ToBranchPtrOutputWithContext(context.Background())
}

func (i *branchPtrType) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPtrOutput)
}

// BranchArrayInput is an input type that accepts BranchArray and BranchArrayOutput values.
// You can construct a concrete instance of `BranchArrayInput` via:
//
//          BranchArray{ BranchArgs{...} }
type BranchArrayInput interface {
	pulumi.Input

	ToBranchArrayOutput() BranchArrayOutput
	ToBranchArrayOutputWithContext(context.Context) BranchArrayOutput
}

type BranchArray []BranchInput

func (BranchArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Branch)(nil))
}

func (i BranchArray) ToBranchArrayOutput() BranchArrayOutput {
	return i.ToBranchArrayOutputWithContext(context.Background())
}

func (i BranchArray) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchArrayOutput)
}

// BranchMapInput is an input type that accepts BranchMap and BranchMapOutput values.
// You can construct a concrete instance of `BranchMapInput` via:
//
//          BranchMap{ "key": BranchArgs{...} }
type BranchMapInput interface {
	pulumi.Input

	ToBranchMapOutput() BranchMapOutput
	ToBranchMapOutputWithContext(context.Context) BranchMapOutput
}

type BranchMap map[string]BranchInput

func (BranchMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Branch)(nil))
}

func (i BranchMap) ToBranchMapOutput() BranchMapOutput {
	return i.ToBranchMapOutputWithContext(context.Background())
}

func (i BranchMap) ToBranchMapOutputWithContext(ctx context.Context) BranchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchMapOutput)
}

type BranchOutput struct {
	*pulumi.OutputState
}

func (BranchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Branch)(nil))
}

func (o BranchOutput) ToBranchOutput() BranchOutput {
	return o
}

func (o BranchOutput) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return o
}

func (o BranchOutput) ToBranchPtrOutput() BranchPtrOutput {
	return o.ToBranchPtrOutputWithContext(context.Background())
}

func (o BranchOutput) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return o.ApplyT(func(v Branch) *Branch {
		return &v
	}).(BranchPtrOutput)
}

type BranchPtrOutput struct {
	*pulumi.OutputState
}

func (BranchPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil))
}

func (o BranchPtrOutput) ToBranchPtrOutput() BranchPtrOutput {
	return o
}

func (o BranchPtrOutput) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return o
}

type BranchArrayOutput struct{ *pulumi.OutputState }

func (BranchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Branch)(nil))
}

func (o BranchArrayOutput) ToBranchArrayOutput() BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) Index(i pulumi.IntInput) BranchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Branch {
		return vs[0].([]Branch)[vs[1].(int)]
	}).(BranchOutput)
}

type BranchMapOutput struct{ *pulumi.OutputState }

func (BranchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Branch)(nil))
}

func (o BranchMapOutput) ToBranchMapOutput() BranchMapOutput {
	return o
}

func (o BranchMapOutput) ToBranchMapOutputWithContext(ctx context.Context) BranchMapOutput {
	return o
}

func (o BranchMapOutput) MapIndex(k pulumi.StringInput) BranchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Branch {
		return vs[0].(map[string]Branch)[vs[1].(string)]
	}).(BranchOutput)
}

func init() {
	pulumi.RegisterOutputType(BranchOutput{})
	pulumi.RegisterOutputType(BranchPtrOutput{})
	pulumi.RegisterOutputType(BranchArrayOutput{})
	pulumi.RegisterOutputType(BranchMapOutput{})
}
