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

// Use this data source to retrieve all deploy keys of a repository.
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
//			_, err := github.GetRepositoryDeployKeys(ctx, &github.GetRepositoryDeployKeysArgs{
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
func GetRepositoryDeployKeys(ctx *pulumi.Context, args *GetRepositoryDeployKeysArgs, opts ...pulumi.InvokeOption) (*GetRepositoryDeployKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRepositoryDeployKeysResult
	err := ctx.Invoke("github:index/getRepositoryDeployKeys:getRepositoryDeployKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryDeployKeys.
type GetRepositoryDeployKeysArgs struct {
	// Name of the repository to retrieve the branches from.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRepositoryDeployKeys.
type GetRepositoryDeployKeysResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of this repository's deploy keys. Each element of `keys` has the following attributes:
	Keys       []GetRepositoryDeployKeysKey `pulumi:"keys"`
	Repository string                       `pulumi:"repository"`
}

func GetRepositoryDeployKeysOutput(ctx *pulumi.Context, args GetRepositoryDeployKeysOutputArgs, opts ...pulumi.InvokeOption) GetRepositoryDeployKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRepositoryDeployKeysResult, error) {
			args := v.(GetRepositoryDeployKeysArgs)
			r, err := GetRepositoryDeployKeys(ctx, &args, opts...)
			var s GetRepositoryDeployKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRepositoryDeployKeysResultOutput)
}

// A collection of arguments for invoking getRepositoryDeployKeys.
type GetRepositoryDeployKeysOutputArgs struct {
	// Name of the repository to retrieve the branches from.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetRepositoryDeployKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryDeployKeysArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryDeployKeys.
type GetRepositoryDeployKeysResultOutput struct{ *pulumi.OutputState }

func (GetRepositoryDeployKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoryDeployKeysResult)(nil)).Elem()
}

func (o GetRepositoryDeployKeysResultOutput) ToGetRepositoryDeployKeysResultOutput() GetRepositoryDeployKeysResultOutput {
	return o
}

func (o GetRepositoryDeployKeysResultOutput) ToGetRepositoryDeployKeysResultOutputWithContext(ctx context.Context) GetRepositoryDeployKeysResultOutput {
	return o
}

func (o GetRepositoryDeployKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetRepositoryDeployKeysResult] {
	return pulumix.Output[GetRepositoryDeployKeysResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetRepositoryDeployKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryDeployKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of this repository's deploy keys. Each element of `keys` has the following attributes:
func (o GetRepositoryDeployKeysResultOutput) Keys() GetRepositoryDeployKeysKeyArrayOutput {
	return o.ApplyT(func(v GetRepositoryDeployKeysResult) []GetRepositoryDeployKeysKey { return v.Keys }).(GetRepositoryDeployKeysKeyArrayOutput)
}

func (o GetRepositoryDeployKeysResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoryDeployKeysResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRepositoryDeployKeysResultOutput{})
}
