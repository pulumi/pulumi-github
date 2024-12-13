// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a specific GitHub Pull Request in a repository.
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
//			_, err := github.LookupRepositoryPullRequest(ctx, &github.LookupRepositoryPullRequestArgs{
//				BaseRepository: "example_repository",
//				Number:         1,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRepositoryPullRequest(ctx *pulumi.Context, args *LookupRepositoryPullRequestArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryPullRequestResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryPullRequestResult
	err := ctx.Invoke("github:index/getRepositoryPullRequest:getRepositoryPullRequest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryPullRequest.
type LookupRepositoryPullRequestArgs struct {
	// Name of the base repository to retrieve the Pull Request from.
	BaseRepository string `pulumi:"baseRepository"`
	// The number of the Pull Request within the repository.
	Number int `pulumi:"number"`
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner *string `pulumi:"owner"`
}

// A collection of values returned by getRepositoryPullRequest.
type LookupRepositoryPullRequestResult struct {
	// Name of the ref (branch) of the Pull Request base.
	BaseRef        string `pulumi:"baseRef"`
	BaseRepository string `pulumi:"baseRepository"`
	// Head commit SHA of the Pull Request base.
	BaseSha string `pulumi:"baseSha"`
	// Body of the Pull Request.
	Body string `pulumi:"body"`
	// Indicates Whether this Pull Request is a draft.
	Draft bool `pulumi:"draft"`
	// Owner of the Pull Request head repository.
	HeadOwner string `pulumi:"headOwner"`
	HeadRef   string `pulumi:"headRef"`
	// Name of the Pull Request head repository.
	HeadRepository string `pulumi:"headRepository"`
	// Head commit SHA of the Pull Request head.
	HeadSha string `pulumi:"headSha"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of label names set on the Pull Request.
	Labels []string `pulumi:"labels"`
	// Indicates whether the base repository maintainers can modify the Pull Request.
	MaintainerCanModify bool `pulumi:"maintainerCanModify"`
	Number              int  `pulumi:"number"`
	// Unix timestamp indicating the Pull Request creation time.
	OpenedAt int `pulumi:"openedAt"`
	// GitHub login of the user who opened the Pull Request.
	OpenedBy string  `pulumi:"openedBy"`
	Owner    *string `pulumi:"owner"`
	// the current Pull Request state - can be "open", "closed" or "merged".
	State string `pulumi:"state"`
	// The title of the Pull Request.
	Title string `pulumi:"title"`
	// The timestamp of the last Pull Request update.
	UpdatedAt int `pulumi:"updatedAt"`
}

func LookupRepositoryPullRequestOutput(ctx *pulumi.Context, args LookupRepositoryPullRequestOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryPullRequestResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRepositoryPullRequestResultOutput, error) {
			args := v.(LookupRepositoryPullRequestArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getRepositoryPullRequest:getRepositoryPullRequest", args, LookupRepositoryPullRequestResultOutput{}, options).(LookupRepositoryPullRequestResultOutput), nil
		}).(LookupRepositoryPullRequestResultOutput)
}

// A collection of arguments for invoking getRepositoryPullRequest.
type LookupRepositoryPullRequestOutputArgs struct {
	// Name of the base repository to retrieve the Pull Request from.
	BaseRepository pulumi.StringInput `pulumi:"baseRepository"`
	// The number of the Pull Request within the repository.
	Number pulumi.IntInput `pulumi:"number"`
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner pulumi.StringPtrInput `pulumi:"owner"`
}

func (LookupRepositoryPullRequestOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryPullRequestArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryPullRequest.
type LookupRepositoryPullRequestResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryPullRequestResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryPullRequestResult)(nil)).Elem()
}

func (o LookupRepositoryPullRequestResultOutput) ToLookupRepositoryPullRequestResultOutput() LookupRepositoryPullRequestResultOutput {
	return o
}

func (o LookupRepositoryPullRequestResultOutput) ToLookupRepositoryPullRequestResultOutputWithContext(ctx context.Context) LookupRepositoryPullRequestResultOutput {
	return o
}

// Name of the ref (branch) of the Pull Request base.
func (o LookupRepositoryPullRequestResultOutput) BaseRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.BaseRef }).(pulumi.StringOutput)
}

func (o LookupRepositoryPullRequestResultOutput) BaseRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.BaseRepository }).(pulumi.StringOutput)
}

// Head commit SHA of the Pull Request base.
func (o LookupRepositoryPullRequestResultOutput) BaseSha() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.BaseSha }).(pulumi.StringOutput)
}

// Body of the Pull Request.
func (o LookupRepositoryPullRequestResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.Body }).(pulumi.StringOutput)
}

// Indicates Whether this Pull Request is a draft.
func (o LookupRepositoryPullRequestResultOutput) Draft() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) bool { return v.Draft }).(pulumi.BoolOutput)
}

// Owner of the Pull Request head repository.
func (o LookupRepositoryPullRequestResultOutput) HeadOwner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.HeadOwner }).(pulumi.StringOutput)
}

func (o LookupRepositoryPullRequestResultOutput) HeadRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.HeadRef }).(pulumi.StringOutput)
}

// Name of the Pull Request head repository.
func (o LookupRepositoryPullRequestResultOutput) HeadRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.HeadRepository }).(pulumi.StringOutput)
}

// Head commit SHA of the Pull Request head.
func (o LookupRepositoryPullRequestResultOutput) HeadSha() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.HeadSha }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRepositoryPullRequestResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of label names set on the Pull Request.
func (o LookupRepositoryPullRequestResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// Indicates whether the base repository maintainers can modify the Pull Request.
func (o LookupRepositoryPullRequestResultOutput) MaintainerCanModify() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) bool { return v.MaintainerCanModify }).(pulumi.BoolOutput)
}

func (o LookupRepositoryPullRequestResultOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) int { return v.Number }).(pulumi.IntOutput)
}

// Unix timestamp indicating the Pull Request creation time.
func (o LookupRepositoryPullRequestResultOutput) OpenedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) int { return v.OpenedAt }).(pulumi.IntOutput)
}

// GitHub login of the user who opened the Pull Request.
func (o LookupRepositoryPullRequestResultOutput) OpenedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.OpenedBy }).(pulumi.StringOutput)
}

func (o LookupRepositoryPullRequestResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

// the current Pull Request state - can be "open", "closed" or "merged".
func (o LookupRepositoryPullRequestResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.State }).(pulumi.StringOutput)
}

// The title of the Pull Request.
func (o LookupRepositoryPullRequestResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) string { return v.Title }).(pulumi.StringOutput)
}

// The timestamp of the last Pull Request update.
func (o LookupRepositoryPullRequestResultOutput) UpdatedAt() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRepositoryPullRequestResult) int { return v.UpdatedAt }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryPullRequestResultOutput{})
}
