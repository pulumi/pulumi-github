// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage PullRequests for repositories within your GitHub organization or personal account.
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
//			_, err := github.NewRepositoryPullRequest(ctx, "example", &github.RepositoryPullRequestArgs{
//				BaseRef:        pulumi.String("main"),
//				BaseRepository: pulumi.String("example-repository"),
//				Body:           pulumi.String("This will change everything"),
//				HeadRef:        pulumi.String("feature-branch"),
//				Title:          pulumi.String("My newest feature"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type RepositoryPullRequest struct {
	pulumi.CustomResourceState

	// Name of the branch serving as the base of the Pull Request.
	BaseRef pulumi.StringOutput `pulumi:"baseRef"`
	// Name of the base repository to retrieve the Pull Requests from.
	BaseRepository pulumi.StringOutput `pulumi:"baseRepository"`
	// Head commit SHA of the Pull Request base.
	BaseSha pulumi.StringOutput `pulumi:"baseSha"`
	// Body of the Pull Request.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// Indicates Whether this Pull Request is a draft.
	Draft pulumi.BoolOutput `pulumi:"draft"`
	// Name of the branch serving as the head of the Pull Request.
	HeadRef pulumi.StringOutput `pulumi:"headRef"`
	// Head commit SHA of the Pull Request head.
	HeadSha pulumi.StringOutput `pulumi:"headSha"`
	// List of label names set on the Pull Request.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	MaintainerCanModify pulumi.BoolPtrOutput `pulumi:"maintainerCanModify"`
	// The number of the Pull Request within the repository.
	Number pulumi.IntOutput `pulumi:"number"`
	// Unix timestamp indicating the Pull Request creation time.
	OpenedAt pulumi.IntOutput `pulumi:"openedAt"`
	// GitHub login of the user who opened the Pull Request.
	OpenedBy pulumi.StringOutput `pulumi:"openedBy"`
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner pulumi.StringPtrOutput `pulumi:"owner"`
	// the current Pull Request state - can be "open", "closed" or "merged".
	State pulumi.StringOutput `pulumi:"state"`
	// The title of the Pull Request.
	Title pulumi.StringOutput `pulumi:"title"`
	// The timestamp of the last Pull Request update.
	UpdatedAt pulumi.IntOutput `pulumi:"updatedAt"`
}

// NewRepositoryPullRequest registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPullRequest(ctx *pulumi.Context,
	name string, args *RepositoryPullRequestArgs, opts ...pulumi.ResourceOption) (*RepositoryPullRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseRef == nil {
		return nil, errors.New("invalid value for required argument 'BaseRef'")
	}
	if args.BaseRepository == nil {
		return nil, errors.New("invalid value for required argument 'BaseRepository'")
	}
	if args.HeadRef == nil {
		return nil, errors.New("invalid value for required argument 'HeadRef'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource RepositoryPullRequest
	err := ctx.RegisterResource("github:index/repositoryPullRequest:RepositoryPullRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPullRequest gets an existing RepositoryPullRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPullRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPullRequestState, opts ...pulumi.ResourceOption) (*RepositoryPullRequest, error) {
	var resource RepositoryPullRequest
	err := ctx.ReadResource("github:index/repositoryPullRequest:RepositoryPullRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPullRequest resources.
type repositoryPullRequestState struct {
	// Name of the branch serving as the base of the Pull Request.
	BaseRef *string `pulumi:"baseRef"`
	// Name of the base repository to retrieve the Pull Requests from.
	BaseRepository *string `pulumi:"baseRepository"`
	// Head commit SHA of the Pull Request base.
	BaseSha *string `pulumi:"baseSha"`
	// Body of the Pull Request.
	Body *string `pulumi:"body"`
	// Indicates Whether this Pull Request is a draft.
	Draft *bool `pulumi:"draft"`
	// Name of the branch serving as the head of the Pull Request.
	HeadRef *string `pulumi:"headRef"`
	// Head commit SHA of the Pull Request head.
	HeadSha *string `pulumi:"headSha"`
	// List of label names set on the Pull Request.
	Labels []string `pulumi:"labels"`
	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	MaintainerCanModify *bool `pulumi:"maintainerCanModify"`
	// The number of the Pull Request within the repository.
	Number *int `pulumi:"number"`
	// Unix timestamp indicating the Pull Request creation time.
	OpenedAt *int `pulumi:"openedAt"`
	// GitHub login of the user who opened the Pull Request.
	OpenedBy *string `pulumi:"openedBy"`
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner *string `pulumi:"owner"`
	// the current Pull Request state - can be "open", "closed" or "merged".
	State *string `pulumi:"state"`
	// The title of the Pull Request.
	Title *string `pulumi:"title"`
	// The timestamp of the last Pull Request update.
	UpdatedAt *int `pulumi:"updatedAt"`
}

type RepositoryPullRequestState struct {
	// Name of the branch serving as the base of the Pull Request.
	BaseRef pulumi.StringPtrInput
	// Name of the base repository to retrieve the Pull Requests from.
	BaseRepository pulumi.StringPtrInput
	// Head commit SHA of the Pull Request base.
	BaseSha pulumi.StringPtrInput
	// Body of the Pull Request.
	Body pulumi.StringPtrInput
	// Indicates Whether this Pull Request is a draft.
	Draft pulumi.BoolPtrInput
	// Name of the branch serving as the head of the Pull Request.
	HeadRef pulumi.StringPtrInput
	// Head commit SHA of the Pull Request head.
	HeadSha pulumi.StringPtrInput
	// List of label names set on the Pull Request.
	Labels pulumi.StringArrayInput
	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	MaintainerCanModify pulumi.BoolPtrInput
	// The number of the Pull Request within the repository.
	Number pulumi.IntPtrInput
	// Unix timestamp indicating the Pull Request creation time.
	OpenedAt pulumi.IntPtrInput
	// GitHub login of the user who opened the Pull Request.
	OpenedBy pulumi.StringPtrInput
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner pulumi.StringPtrInput
	// the current Pull Request state - can be "open", "closed" or "merged".
	State pulumi.StringPtrInput
	// The title of the Pull Request.
	Title pulumi.StringPtrInput
	// The timestamp of the last Pull Request update.
	UpdatedAt pulumi.IntPtrInput
}

func (RepositoryPullRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPullRequestState)(nil)).Elem()
}

type repositoryPullRequestArgs struct {
	// Name of the branch serving as the base of the Pull Request.
	BaseRef string `pulumi:"baseRef"`
	// Name of the base repository to retrieve the Pull Requests from.
	BaseRepository string `pulumi:"baseRepository"`
	// Body of the Pull Request.
	Body *string `pulumi:"body"`
	// Name of the branch serving as the head of the Pull Request.
	HeadRef string `pulumi:"headRef"`
	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	MaintainerCanModify *bool `pulumi:"maintainerCanModify"`
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner *string `pulumi:"owner"`
	// The title of the Pull Request.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a RepositoryPullRequest resource.
type RepositoryPullRequestArgs struct {
	// Name of the branch serving as the base of the Pull Request.
	BaseRef pulumi.StringInput
	// Name of the base repository to retrieve the Pull Requests from.
	BaseRepository pulumi.StringInput
	// Body of the Pull Request.
	Body pulumi.StringPtrInput
	// Name of the branch serving as the head of the Pull Request.
	HeadRef pulumi.StringInput
	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	MaintainerCanModify pulumi.BoolPtrInput
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner pulumi.StringPtrInput
	// The title of the Pull Request.
	Title pulumi.StringInput
}

func (RepositoryPullRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPullRequestArgs)(nil)).Elem()
}

type RepositoryPullRequestInput interface {
	pulumi.Input

	ToRepositoryPullRequestOutput() RepositoryPullRequestOutput
	ToRepositoryPullRequestOutputWithContext(ctx context.Context) RepositoryPullRequestOutput
}

func (*RepositoryPullRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPullRequest)(nil)).Elem()
}

func (i *RepositoryPullRequest) ToRepositoryPullRequestOutput() RepositoryPullRequestOutput {
	return i.ToRepositoryPullRequestOutputWithContext(context.Background())
}

func (i *RepositoryPullRequest) ToRepositoryPullRequestOutputWithContext(ctx context.Context) RepositoryPullRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPullRequestOutput)
}

// RepositoryPullRequestArrayInput is an input type that accepts RepositoryPullRequestArray and RepositoryPullRequestArrayOutput values.
// You can construct a concrete instance of `RepositoryPullRequestArrayInput` via:
//
//	RepositoryPullRequestArray{ RepositoryPullRequestArgs{...} }
type RepositoryPullRequestArrayInput interface {
	pulumi.Input

	ToRepositoryPullRequestArrayOutput() RepositoryPullRequestArrayOutput
	ToRepositoryPullRequestArrayOutputWithContext(context.Context) RepositoryPullRequestArrayOutput
}

type RepositoryPullRequestArray []RepositoryPullRequestInput

func (RepositoryPullRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPullRequest)(nil)).Elem()
}

func (i RepositoryPullRequestArray) ToRepositoryPullRequestArrayOutput() RepositoryPullRequestArrayOutput {
	return i.ToRepositoryPullRequestArrayOutputWithContext(context.Background())
}

func (i RepositoryPullRequestArray) ToRepositoryPullRequestArrayOutputWithContext(ctx context.Context) RepositoryPullRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPullRequestArrayOutput)
}

// RepositoryPullRequestMapInput is an input type that accepts RepositoryPullRequestMap and RepositoryPullRequestMapOutput values.
// You can construct a concrete instance of `RepositoryPullRequestMapInput` via:
//
//	RepositoryPullRequestMap{ "key": RepositoryPullRequestArgs{...} }
type RepositoryPullRequestMapInput interface {
	pulumi.Input

	ToRepositoryPullRequestMapOutput() RepositoryPullRequestMapOutput
	ToRepositoryPullRequestMapOutputWithContext(context.Context) RepositoryPullRequestMapOutput
}

type RepositoryPullRequestMap map[string]RepositoryPullRequestInput

func (RepositoryPullRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPullRequest)(nil)).Elem()
}

func (i RepositoryPullRequestMap) ToRepositoryPullRequestMapOutput() RepositoryPullRequestMapOutput {
	return i.ToRepositoryPullRequestMapOutputWithContext(context.Background())
}

func (i RepositoryPullRequestMap) ToRepositoryPullRequestMapOutputWithContext(ctx context.Context) RepositoryPullRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPullRequestMapOutput)
}

type RepositoryPullRequestOutput struct{ *pulumi.OutputState }

func (RepositoryPullRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPullRequest)(nil)).Elem()
}

func (o RepositoryPullRequestOutput) ToRepositoryPullRequestOutput() RepositoryPullRequestOutput {
	return o
}

func (o RepositoryPullRequestOutput) ToRepositoryPullRequestOutputWithContext(ctx context.Context) RepositoryPullRequestOutput {
	return o
}

// Name of the branch serving as the base of the Pull Request.
func (o RepositoryPullRequestOutput) BaseRef() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.BaseRef }).(pulumi.StringOutput)
}

// Name of the base repository to retrieve the Pull Requests from.
func (o RepositoryPullRequestOutput) BaseRepository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.BaseRepository }).(pulumi.StringOutput)
}

// Head commit SHA of the Pull Request base.
func (o RepositoryPullRequestOutput) BaseSha() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.BaseSha }).(pulumi.StringOutput)
}

// Body of the Pull Request.
func (o RepositoryPullRequestOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringPtrOutput { return v.Body }).(pulumi.StringPtrOutput)
}

// Indicates Whether this Pull Request is a draft.
func (o RepositoryPullRequestOutput) Draft() pulumi.BoolOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.BoolOutput { return v.Draft }).(pulumi.BoolOutput)
}

// Name of the branch serving as the head of the Pull Request.
func (o RepositoryPullRequestOutput) HeadRef() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.HeadRef }).(pulumi.StringOutput)
}

// Head commit SHA of the Pull Request head.
func (o RepositoryPullRequestOutput) HeadSha() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.HeadSha }).(pulumi.StringOutput)
}

// List of label names set on the Pull Request.
func (o RepositoryPullRequestOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
func (o RepositoryPullRequestOutput) MaintainerCanModify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.BoolPtrOutput { return v.MaintainerCanModify }).(pulumi.BoolPtrOutput)
}

// The number of the Pull Request within the repository.
func (o RepositoryPullRequestOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.IntOutput { return v.Number }).(pulumi.IntOutput)
}

// Unix timestamp indicating the Pull Request creation time.
func (o RepositoryPullRequestOutput) OpenedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.IntOutput { return v.OpenedAt }).(pulumi.IntOutput)
}

// GitHub login of the user who opened the Pull Request.
func (o RepositoryPullRequestOutput) OpenedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.OpenedBy }).(pulumi.StringOutput)
}

// Owner of the repository. If not provided, the provider's default owner is used.
func (o RepositoryPullRequestOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringPtrOutput { return v.Owner }).(pulumi.StringPtrOutput)
}

// the current Pull Request state - can be "open", "closed" or "merged".
func (o RepositoryPullRequestOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The title of the Pull Request.
func (o RepositoryPullRequestOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The timestamp of the last Pull Request update.
func (o RepositoryPullRequestOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryPullRequest) pulumi.IntOutput { return v.UpdatedAt }).(pulumi.IntOutput)
}

type RepositoryPullRequestArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPullRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPullRequest)(nil)).Elem()
}

func (o RepositoryPullRequestArrayOutput) ToRepositoryPullRequestArrayOutput() RepositoryPullRequestArrayOutput {
	return o
}

func (o RepositoryPullRequestArrayOutput) ToRepositoryPullRequestArrayOutputWithContext(ctx context.Context) RepositoryPullRequestArrayOutput {
	return o
}

func (o RepositoryPullRequestArrayOutput) Index(i pulumi.IntInput) RepositoryPullRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryPullRequest {
		return vs[0].([]*RepositoryPullRequest)[vs[1].(int)]
	}).(RepositoryPullRequestOutput)
}

type RepositoryPullRequestMapOutput struct{ *pulumi.OutputState }

func (RepositoryPullRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPullRequest)(nil)).Elem()
}

func (o RepositoryPullRequestMapOutput) ToRepositoryPullRequestMapOutput() RepositoryPullRequestMapOutput {
	return o
}

func (o RepositoryPullRequestMapOutput) ToRepositoryPullRequestMapOutputWithContext(ctx context.Context) RepositoryPullRequestMapOutput {
	return o
}

func (o RepositoryPullRequestMapOutput) MapIndex(k pulumi.StringInput) RepositoryPullRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryPullRequest {
		return vs[0].(map[string]*RepositoryPullRequest)[vs[1].(string)]
	}).(RepositoryPullRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPullRequestInput)(nil)).Elem(), &RepositoryPullRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPullRequestArrayInput)(nil)).Elem(), RepositoryPullRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPullRequestMapInput)(nil)).Elem(), RepositoryPullRequestMap{})
	pulumi.RegisterOutputType(RepositoryPullRequestOutput{})
	pulumi.RegisterOutputType(RepositoryPullRequestArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPullRequestMapOutput{})
}
