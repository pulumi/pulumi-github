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

// This resource allows you to create and manage a release in a specific
// GitHub repository.
//
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
//			repo, err := github.NewRepository(ctx, "repo", &github.RepositoryArgs{
//				Description: pulumi.String("GitHub repo managed by Terraform"),
//				Private:     pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRelease(ctx, "example", &github.ReleaseArgs{
//				Repository: repo.Name,
//				TagName:    pulumi.String("v1.0.0"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### On Non-Default Branch
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
//			exampleRepository, err := github.NewRepository(ctx, "exampleRepository", &github.RepositoryArgs{
//				AutoInit: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleBranch, err := github.NewBranch(ctx, "exampleBranch", &github.BranchArgs{
//				Repository:   exampleRepository.Name,
//				Branch:       pulumi.String("branch_name"),
//				SourceBranch: exampleRepository.DefaultBranch,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRelease(ctx, "exampleRelease", &github.ReleaseArgs{
//				Repository:      exampleRepository.Name,
//				TagName:         pulumi.String("v1.0.0"),
//				TargetCommitish: exampleBranch.Branch,
//				Draft:           pulumi.Bool(false),
//				Prerelease:      pulumi.Bool(false),
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
// This resource can be imported using the `name` of the repository, combined with the `id` of the release, and a `:` character for separating components, e.g.
//
// ```sh
//
//	$ pulumi import github:index/release:Release example repo:12345678
//
// ```
type Release struct {
	pulumi.CustomResourceState

	// Text describing the contents of the tag.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
	DiscussionCategoryName pulumi.StringPtrOutput `pulumi:"discussionCategoryName"`
	// Set to `false` to create a published release.
	Draft pulumi.BoolPtrOutput `pulumi:"draft"`
	Etag  pulumi.StringOutput  `pulumi:"etag"`
	// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
	GenerateReleaseNotes pulumi.BoolPtrOutput `pulumi:"generateReleaseNotes"`
	// The name of the release.
	Name pulumi.StringOutput `pulumi:"name"`
	// Set to `false` to identify the release as a full release.
	Prerelease pulumi.BoolPtrOutput `pulumi:"prerelease"`
	// The name of the repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The name of the tag.
	TagName pulumi.StringOutput `pulumi:"tagName"`
	// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
	TargetCommitish pulumi.StringPtrOutput `pulumi:"targetCommitish"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.TagName == nil {
		return nil, errors.New("invalid value for required argument 'TagName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Release
	err := ctx.RegisterResource("github:index/release:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("github:index/release:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
	// Text describing the contents of the tag.
	Body *string `pulumi:"body"`
	// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
	DiscussionCategoryName *string `pulumi:"discussionCategoryName"`
	// Set to `false` to create a published release.
	Draft *bool   `pulumi:"draft"`
	Etag  *string `pulumi:"etag"`
	// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
	GenerateReleaseNotes *bool `pulumi:"generateReleaseNotes"`
	// The name of the release.
	Name *string `pulumi:"name"`
	// Set to `false` to identify the release as a full release.
	Prerelease *bool `pulumi:"prerelease"`
	// The name of the repository.
	Repository *string `pulumi:"repository"`
	// The name of the tag.
	TagName *string `pulumi:"tagName"`
	// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
	TargetCommitish *string `pulumi:"targetCommitish"`
}

type ReleaseState struct {
	// Text describing the contents of the tag.
	Body pulumi.StringPtrInput
	// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
	DiscussionCategoryName pulumi.StringPtrInput
	// Set to `false` to create a published release.
	Draft pulumi.BoolPtrInput
	Etag  pulumi.StringPtrInput
	// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
	GenerateReleaseNotes pulumi.BoolPtrInput
	// The name of the release.
	Name pulumi.StringPtrInput
	// Set to `false` to identify the release as a full release.
	Prerelease pulumi.BoolPtrInput
	// The name of the repository.
	Repository pulumi.StringPtrInput
	// The name of the tag.
	TagName pulumi.StringPtrInput
	// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
	TargetCommitish pulumi.StringPtrInput
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	// Text describing the contents of the tag.
	Body *string `pulumi:"body"`
	// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
	DiscussionCategoryName *string `pulumi:"discussionCategoryName"`
	// Set to `false` to create a published release.
	Draft *bool `pulumi:"draft"`
	// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
	GenerateReleaseNotes *bool `pulumi:"generateReleaseNotes"`
	// The name of the release.
	Name *string `pulumi:"name"`
	// Set to `false` to identify the release as a full release.
	Prerelease *bool `pulumi:"prerelease"`
	// The name of the repository.
	Repository string `pulumi:"repository"`
	// The name of the tag.
	TagName string `pulumi:"tagName"`
	// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
	TargetCommitish *string `pulumi:"targetCommitish"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	// Text describing the contents of the tag.
	Body pulumi.StringPtrInput
	// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
	DiscussionCategoryName pulumi.StringPtrInput
	// Set to `false` to create a published release.
	Draft pulumi.BoolPtrInput
	// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
	GenerateReleaseNotes pulumi.BoolPtrInput
	// The name of the release.
	Name pulumi.StringPtrInput
	// Set to `false` to identify the release as a full release.
	Prerelease pulumi.BoolPtrInput
	// The name of the repository.
	Repository pulumi.StringInput
	// The name of the tag.
	TagName pulumi.StringInput
	// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
	TargetCommitish pulumi.StringPtrInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

func (i *Release) ToOutput(ctx context.Context) pulumix.Output[*Release] {
	return pulumix.Output[*Release]{
		OutputState: i.ToReleaseOutputWithContext(ctx).OutputState,
	}
}

// ReleaseArrayInput is an input type that accepts ReleaseArray and ReleaseArrayOutput values.
// You can construct a concrete instance of `ReleaseArrayInput` via:
//
//	ReleaseArray{ ReleaseArgs{...} }
type ReleaseArrayInput interface {
	pulumi.Input

	ToReleaseArrayOutput() ReleaseArrayOutput
	ToReleaseArrayOutputWithContext(context.Context) ReleaseArrayOutput
}

type ReleaseArray []ReleaseInput

func (ReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (i ReleaseArray) ToReleaseArrayOutput() ReleaseArrayOutput {
	return i.ToReleaseArrayOutputWithContext(context.Background())
}

func (i ReleaseArray) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseArrayOutput)
}

func (i ReleaseArray) ToOutput(ctx context.Context) pulumix.Output[[]*Release] {
	return pulumix.Output[[]*Release]{
		OutputState: i.ToReleaseArrayOutputWithContext(ctx).OutputState,
	}
}

// ReleaseMapInput is an input type that accepts ReleaseMap and ReleaseMapOutput values.
// You can construct a concrete instance of `ReleaseMapInput` via:
//
//	ReleaseMap{ "key": ReleaseArgs{...} }
type ReleaseMapInput interface {
	pulumi.Input

	ToReleaseMapOutput() ReleaseMapOutput
	ToReleaseMapOutputWithContext(context.Context) ReleaseMapOutput
}

type ReleaseMap map[string]ReleaseInput

func (ReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (i ReleaseMap) ToReleaseMapOutput() ReleaseMapOutput {
	return i.ToReleaseMapOutputWithContext(context.Background())
}

func (i ReleaseMap) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseMapOutput)
}

func (i ReleaseMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Release] {
	return pulumix.Output[map[string]*Release]{
		OutputState: i.ToReleaseMapOutputWithContext(ctx).OutputState,
	}
}

type ReleaseOutput struct{ *pulumi.OutputState }

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToOutput(ctx context.Context) pulumix.Output[*Release] {
	return pulumix.Output[*Release]{
		OutputState: o.OutputState,
	}
}

// Text describing the contents of the tag.
func (o ReleaseOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.Body }).(pulumi.StringPtrOutput)
}

// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
func (o ReleaseOutput) DiscussionCategoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.DiscussionCategoryName }).(pulumi.StringPtrOutput)
}

// Set to `false` to create a published release.
func (o ReleaseOutput) Draft() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.Draft }).(pulumi.BoolPtrOutput)
}

func (o ReleaseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
func (o ReleaseOutput) GenerateReleaseNotes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.GenerateReleaseNotes }).(pulumi.BoolPtrOutput)
}

// The name of the release.
func (o ReleaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set to `false` to identify the release as a full release.
func (o ReleaseOutput) Prerelease() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.Prerelease }).(pulumi.BoolPtrOutput)
}

// The name of the repository.
func (o ReleaseOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The name of the tag.
func (o ReleaseOutput) TagName() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.TagName }).(pulumi.StringOutput)
}

// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
func (o ReleaseOutput) TargetCommitish() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.TargetCommitish }).(pulumi.StringPtrOutput)
}

type ReleaseArrayOutput struct{ *pulumi.OutputState }

func (ReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (o ReleaseArrayOutput) ToReleaseArrayOutput() ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Release] {
	return pulumix.Output[[]*Release]{
		OutputState: o.OutputState,
	}
}

func (o ReleaseArrayOutput) Index(i pulumi.IntInput) ReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Release {
		return vs[0].([]*Release)[vs[1].(int)]
	}).(ReleaseOutput)
}

type ReleaseMapOutput struct{ *pulumi.OutputState }

func (ReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (o ReleaseMapOutput) ToReleaseMapOutput() ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Release] {
	return pulumix.Output[map[string]*Release]{
		OutputState: o.OutputState,
	}
}

func (o ReleaseMapOutput) MapIndex(k pulumi.StringInput) ReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Release {
		return vs[0].(map[string]*Release)[vs[1].(string)]
	}).(ReleaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseInput)(nil)).Elem(), &Release{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseArrayInput)(nil)).Elem(), ReleaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseMapInput)(nil)).Elem(), ReleaseMap{})
	pulumi.RegisterOutputType(ReleaseOutput{})
	pulumi.RegisterOutputType(ReleaseArrayOutput{})
	pulumi.RegisterOutputType(ReleaseMapOutput{})
}
