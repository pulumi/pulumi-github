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

// Use this data source to retrieve information about a repository branch.
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
//			_, err := github.LookupBranch(ctx, &github.LookupBranchArgs{
//				Branch:     "development",
//				Repository: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBranch(ctx *pulumi.Context, args *LookupBranchArgs, opts ...pulumi.InvokeOption) (*LookupBranchResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBranchResult
	err := ctx.Invoke("github:index/getBranch:getBranch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBranch.
type LookupBranchArgs struct {
	// The repository branch to create.
	Branch string `pulumi:"branch"`
	// The GitHub repository name.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getBranch.
type LookupBranchResult struct {
	Branch string `pulumi:"branch"`
	// An etag representing the Branch object.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A string representing a branch reference, in the form of `refs/heads/<branch>`.
	Ref        string `pulumi:"ref"`
	Repository string `pulumi:"repository"`
	// A string storing the reference's `HEAD` commit's SHA1.
	Sha string `pulumi:"sha"`
}

func LookupBranchOutput(ctx *pulumi.Context, args LookupBranchOutputArgs, opts ...pulumi.InvokeOption) LookupBranchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBranchResult, error) {
			args := v.(LookupBranchArgs)
			r, err := LookupBranch(ctx, &args, opts...)
			var s LookupBranchResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBranchResultOutput)
}

// A collection of arguments for invoking getBranch.
type LookupBranchOutputArgs struct {
	// The repository branch to create.
	Branch pulumi.StringInput `pulumi:"branch"`
	// The GitHub repository name.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (LookupBranchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBranchArgs)(nil)).Elem()
}

// A collection of values returned by getBranch.
type LookupBranchResultOutput struct{ *pulumi.OutputState }

func (LookupBranchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBranchResult)(nil)).Elem()
}

func (o LookupBranchResultOutput) ToLookupBranchResultOutput() LookupBranchResultOutput {
	return o
}

func (o LookupBranchResultOutput) ToLookupBranchResultOutputWithContext(ctx context.Context) LookupBranchResultOutput {
	return o
}

func (o LookupBranchResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBranchResult] {
	return pulumix.Output[LookupBranchResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBranchResultOutput) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Branch }).(pulumi.StringOutput)
}

// An etag representing the Branch object.
func (o LookupBranchResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBranchResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Id }).(pulumi.StringOutput)
}

// A string representing a branch reference, in the form of `refs/heads/<branch>`.
func (o LookupBranchResultOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Ref }).(pulumi.StringOutput)
}

func (o LookupBranchResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Repository }).(pulumi.StringOutput)
}

// A string storing the reference's `HEAD` commit's SHA1.
func (o LookupBranchResultOutput) Sha() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBranchResult) string { return v.Sha }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBranchResultOutput{})
}
