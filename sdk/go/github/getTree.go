// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a single tree.
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
//			thisRepository, err := github.LookupRepository(ctx, &github.LookupRepositoryArgs{
//				Name: pulumi.StringRef("example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			thisBranch, err := github.LookupBranch(ctx, &github.LookupBranchArgs{
//				Branch:     thisRepository.DefaultBranch,
//				Repository: thisRepository.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			thisTree, err := github.GetTree(ctx, &github.GetTreeArgs{
//				Recursive:  pulumi.BoolRef(false),
//				Repository: thisRepository.Name,
//				TreeSha:    thisBranch.Sha,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("entries", thisTree.Entries)
//			return nil
//		})
//	}
//
// ```
func GetTree(ctx *pulumi.Context, args *GetTreeArgs, opts ...pulumi.InvokeOption) (*GetTreeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTreeResult
	err := ctx.Invoke("github:index/getTree:getTree", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTree.
type GetTreeArgs struct {
	// Setting this parameter to `true` returns the objects or subtrees referenced by the tree specified in `treeSha`.
	Recursive *bool `pulumi:"recursive"`
	// The name of the repository.
	Repository string `pulumi:"repository"`
	// The SHA1 value for the tree.
	TreeSha string `pulumi:"treeSha"`
}

// A collection of values returned by getTree.
type GetTreeResult struct {
	// Objects (of `path`, `mode`, `type`, `size`, and `sha`) specifying a tree structure.
	Entries []GetTreeEntry `pulumi:"entries"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Recursive  *bool  `pulumi:"recursive"`
	Repository string `pulumi:"repository"`
	TreeSha    string `pulumi:"treeSha"`
}

func GetTreeOutput(ctx *pulumi.Context, args GetTreeOutputArgs, opts ...pulumi.InvokeOption) GetTreeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTreeResult, error) {
			args := v.(GetTreeArgs)
			r, err := GetTree(ctx, &args, opts...)
			var s GetTreeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTreeResultOutput)
}

// A collection of arguments for invoking getTree.
type GetTreeOutputArgs struct {
	// Setting this parameter to `true` returns the objects or subtrees referenced by the tree specified in `treeSha`.
	Recursive pulumi.BoolPtrInput `pulumi:"recursive"`
	// The name of the repository.
	Repository pulumi.StringInput `pulumi:"repository"`
	// The SHA1 value for the tree.
	TreeSha pulumi.StringInput `pulumi:"treeSha"`
}

func (GetTreeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTreeArgs)(nil)).Elem()
}

// A collection of values returned by getTree.
type GetTreeResultOutput struct{ *pulumi.OutputState }

func (GetTreeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTreeResult)(nil)).Elem()
}

func (o GetTreeResultOutput) ToGetTreeResultOutput() GetTreeResultOutput {
	return o
}

func (o GetTreeResultOutput) ToGetTreeResultOutputWithContext(ctx context.Context) GetTreeResultOutput {
	return o
}

// Objects (of `path`, `mode`, `type`, `size`, and `sha`) specifying a tree structure.
func (o GetTreeResultOutput) Entries() GetTreeEntryArrayOutput {
	return o.ApplyT(func(v GetTreeResult) []GetTreeEntry { return v.Entries }).(GetTreeEntryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTreeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTreeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTreeResultOutput) Recursive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetTreeResult) *bool { return v.Recursive }).(pulumi.BoolPtrOutput)
}

func (o GetTreeResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetTreeResult) string { return v.Repository }).(pulumi.StringOutput)
}

func (o GetTreeResultOutput) TreeSha() pulumi.StringOutput {
	return o.ApplyT(func(v GetTreeResult) string { return v.TreeSha }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTreeResultOutput{})
}
