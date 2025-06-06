// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **Note:** The data source will return a maximum of `1000` repositories
//
//	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).
//
// Use this data source to retrieve a list of GitHub repositories using a search query.
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
//			_, err := github.GetRepositories(ctx, &github.GetRepositoriesArgs{
//				Query:         "org:hashicorp language:Go",
//				IncludeRepoId: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRepositories(ctx *pulumi.Context, args *GetRepositoriesArgs, opts ...pulumi.InvokeOption) (*GetRepositoriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRepositoriesResult
	err := ctx.Invoke("github:index/getRepositories:getRepositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositories.
type GetRepositoriesArgs struct {
	// Returns a list of found repository IDs
	IncludeRepoId *bool `pulumi:"includeRepoId"`
	// Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
	Query string `pulumi:"query"`
	// Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
	ResultsPerPage *int `pulumi:"resultsPerPage"`
	// Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
	Sort *string `pulumi:"sort"`
}

// A collection of values returned by getRepositories.
type GetRepositoriesResult struct {
	FullNames []string `pulumi:"fullNames"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	IncludeRepoId *bool    `pulumi:"includeRepoId"`
	Names         []string `pulumi:"names"`
	Query         string   `pulumi:"query"`
	// (Optional) A list of found repository IDs (e.g. `449898861`)
	RepoIds        []int   `pulumi:"repoIds"`
	ResultsPerPage *int    `pulumi:"resultsPerPage"`
	Sort           *string `pulumi:"sort"`
}

func GetRepositoriesOutput(ctx *pulumi.Context, args GetRepositoriesOutputArgs, opts ...pulumi.InvokeOption) GetRepositoriesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRepositoriesResultOutput, error) {
			args := v.(GetRepositoriesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getRepositories:getRepositories", args, GetRepositoriesResultOutput{}, options).(GetRepositoriesResultOutput), nil
		}).(GetRepositoriesResultOutput)
}

// A collection of arguments for invoking getRepositories.
type GetRepositoriesOutputArgs struct {
	// Returns a list of found repository IDs
	IncludeRepoId pulumi.BoolPtrInput `pulumi:"includeRepoId"`
	// Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
	Query pulumi.StringInput `pulumi:"query"`
	// Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
	ResultsPerPage pulumi.IntPtrInput `pulumi:"resultsPerPage"`
	// Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
	Sort pulumi.StringPtrInput `pulumi:"sort"`
}

func (GetRepositoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesArgs)(nil)).Elem()
}

// A collection of values returned by getRepositories.
type GetRepositoriesResultOutput struct{ *pulumi.OutputState }

func (GetRepositoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesResult)(nil)).Elem()
}

func (o GetRepositoriesResultOutput) ToGetRepositoriesResultOutput() GetRepositoriesResultOutput {
	return o
}

func (o GetRepositoriesResultOutput) ToGetRepositoriesResultOutputWithContext(ctx context.Context) GetRepositoriesResultOutput {
	return o
}

func (o GetRepositoriesResultOutput) FullNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRepositoriesResult) []string { return v.FullNames }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRepositoriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRepositoriesResultOutput) IncludeRepoId() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *bool { return v.IncludeRepoId }).(pulumi.BoolPtrOutput)
}

func (o GetRepositoriesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRepositoriesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRepositoriesResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesResult) string { return v.Query }).(pulumi.StringOutput)
}

// (Optional) A list of found repository IDs (e.g. `449898861`)
func (o GetRepositoriesResultOutput) RepoIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetRepositoriesResult) []int { return v.RepoIds }).(pulumi.IntArrayOutput)
}

func (o GetRepositoriesResultOutput) ResultsPerPage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *int { return v.ResultsPerPage }).(pulumi.IntPtrOutput)
}

func (o GetRepositoriesResultOutput) Sort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *string { return v.Sort }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRepositoriesResultOutput{})
}
