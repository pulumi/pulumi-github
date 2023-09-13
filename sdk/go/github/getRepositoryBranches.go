// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve information about branches in a repository.
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
//			_, err := github.GetRepositoryBranches(ctx, &github.GetRepositoryBranchesArgs{
//				Repository: "example-repository",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRepositoryBranches(ctx *pulumi.Context, args *GetRepositoryBranchesArgs, opts ...pulumi.InvokeOption) (*GetRepositoryBranchesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRepositoryBranchesResult
	err := ctx.Invoke("github:index/getRepositoryBranches:getRepositoryBranches", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryBranches.
type GetRepositoryBranchesArgs struct {
	// . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
	OnlyNonProtectedBranches *bool `pulumi:"onlyNonProtectedBranches"`
	// . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
	OnlyProtectedBranches *bool `pulumi:"onlyProtectedBranches"`
	// Name of the repository to retrieve the branches from.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRepositoryBranches.
type GetRepositoryBranchesResult struct {
	// The list of this repository's branches. Each element of `branches` has the following attributes:
	Branches []GetRepositoryBranchesBranch `pulumi:"branches"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string `pulumi:"id"`
	OnlyNonProtectedBranches *bool  `pulumi:"onlyNonProtectedBranches"`
	OnlyProtectedBranches    *bool  `pulumi:"onlyProtectedBranches"`
	Repository               string `pulumi:"repository"`
}

func GetRepositoryBranchesOutput(ctx *pulumi.Context, args GetRepositoryBranchesOutputArgs, opts ...pulumi.InvokeOption) GetRepositoryBranchesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRepositoryBranchesResult, error) {
			args := v.(GetRepositoryBranchesArgs)
			r, err := GetRepositoryBranches(ctx, &args, opts...)
			var s GetRepositoryBranchesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRepositoryBranchesResultOutput)
}

// A collection of arguments for invoking getRepositoryBranches.
type GetRepositoryBranchesOutputArgs struct {
	// . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
	OnlyNonProtectedBranches pulumi.BoolPtrInput `pulumi:"onlyNonProtectedBranches"`
	// . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
	OnlyProtectedBranches pulumi.BoolPtrInput `pulumi:"onlyProtectedBranches"`
	// Name of the repository to retrieve the branches from.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetRepositoryBranchesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryBranchesArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryBranches.
type GetRepositoryBranchesResultOutput struct{ *pulumi.OutputState }

func (GetRepositoryBranchesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryBranchesResult)(nil)).Elem()
}

func (o GetRepositoryBranchesResultOutput) ToGetRepositoryBranchesResultOutput() GetRepositoryBranchesResultOutput {
	return o
}

func (o GetRepositoryBranchesResultOutput) ToGetRepositoryBranchesResultOutputWithContext(ctx context.Context) GetRepositoryBranchesResultOutput {
	return o
}

func (o GetRepositoryBranchesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetRepositoryBranchesResult] {
	return pulumix.Output[GetRepositoryBranchesResult]{
		OutputState: o.OutputState,
	}
}

// The list of this repository's branches. Each element of `branches` has the following attributes:
func (o GetRepositoryBranchesResultOutput) Branches() GetRepositoryBranchesBranchArrayOutput {
	return o.ApplyT(func(v GetRepositoryBranchesResult) []GetRepositoryBranchesBranch { return v.Branches }).(GetRepositoryBranchesBranchArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRepositoryBranchesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryBranchesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRepositoryBranchesResultOutput) OnlyNonProtectedBranches() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRepositoryBranchesResult) *bool { return v.OnlyNonProtectedBranches }).(pulumi.BoolPtrOutput)
}

func (o GetRepositoryBranchesResultOutput) OnlyProtectedBranches() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRepositoryBranchesResult) *bool { return v.OnlyProtectedBranches }).(pulumi.BoolPtrOutput)
}

func (o GetRepositoryBranchesResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryBranchesResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRepositoryBranchesResultOutput{})
}
