// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve autolink references for a repository.
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
//			_, err := github.GetRepositoryAutolinkReferences(ctx, &github.GetRepositoryAutolinkReferencesArgs{
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
func GetRepositoryAutolinkReferences(ctx *pulumi.Context, args *GetRepositoryAutolinkReferencesArgs, opts ...pulumi.InvokeOption) (*GetRepositoryAutolinkReferencesResult, error) {
	var rv GetRepositoryAutolinkReferencesResult
	err := ctx.Invoke("github:index/getRepositoryAutolinkReferences:getRepositoryAutolinkReferences", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryAutolinkReferences.
type GetRepositoryAutolinkReferencesArgs struct {
	// Name of the repository to retrieve the autolink references from.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRepositoryAutolinkReferences.
type GetRepositoryAutolinkReferencesResult struct {
	// The list of this repository's autolink references. Each element of `autolinkReferences` has the following attributes:
	AutolinkReferences []GetRepositoryAutolinkReferencesAutolinkReference `pulumi:"autolinkReferences"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Repository string `pulumi:"repository"`
}

func GetRepositoryAutolinkReferencesOutput(ctx *pulumi.Context, args GetRepositoryAutolinkReferencesOutputArgs, opts ...pulumi.InvokeOption) GetRepositoryAutolinkReferencesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRepositoryAutolinkReferencesResult, error) {
			args := v.(GetRepositoryAutolinkReferencesArgs)
			r, err := GetRepositoryAutolinkReferences(ctx, &args, opts...)
			var s GetRepositoryAutolinkReferencesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRepositoryAutolinkReferencesResultOutput)
}

// A collection of arguments for invoking getRepositoryAutolinkReferences.
type GetRepositoryAutolinkReferencesOutputArgs struct {
	// Name of the repository to retrieve the autolink references from.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetRepositoryAutolinkReferencesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryAutolinkReferencesArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryAutolinkReferences.
type GetRepositoryAutolinkReferencesResultOutput struct{ *pulumi.OutputState }

func (GetRepositoryAutolinkReferencesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryAutolinkReferencesResult)(nil)).Elem()
}

func (o GetRepositoryAutolinkReferencesResultOutput) ToGetRepositoryAutolinkReferencesResultOutput() GetRepositoryAutolinkReferencesResultOutput {
	return o
}

func (o GetRepositoryAutolinkReferencesResultOutput) ToGetRepositoryAutolinkReferencesResultOutputWithContext(ctx context.Context) GetRepositoryAutolinkReferencesResultOutput {
	return o
}

// The list of this repository's autolink references. Each element of `autolinkReferences` has the following attributes:
func (o GetRepositoryAutolinkReferencesResultOutput) AutolinkReferences() GetRepositoryAutolinkReferencesAutolinkReferenceArrayOutput {
	return o.ApplyT(func(v GetRepositoryAutolinkReferencesResult) []GetRepositoryAutolinkReferencesAutolinkReference {
		return v.AutolinkReferences
	}).(GetRepositoryAutolinkReferencesAutolinkReferenceArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRepositoryAutolinkReferencesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryAutolinkReferencesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRepositoryAutolinkReferencesResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryAutolinkReferencesResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRepositoryAutolinkReferencesResultOutput{})
}
