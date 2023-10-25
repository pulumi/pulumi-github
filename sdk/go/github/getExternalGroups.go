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

// Use this data source to retrieve external groups belonging to an organization.
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
//			exampleExternalGroups, err := github.GetExternalGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			localGroups := exampleExternalGroups
//			ctx.Export("groups", localGroups)
//			return nil
//		})
//	}
//
// ```
func GetExternalGroups(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetExternalGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetExternalGroupsResult
	err := ctx.Invoke("github:index/getExternalGroups:getExternalGroups", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getExternalGroups.
type GetExternalGroupsResult struct {
	// an array of external groups belonging to the organization. Each group consists of the fields documented below.
	ExternalGroups []GetExternalGroupsExternalGroup `pulumi:"externalGroups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetExternalGroupsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetExternalGroupsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetExternalGroupsResult, error) {
		r, err := GetExternalGroups(ctx, opts...)
		var s GetExternalGroupsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetExternalGroupsResultOutput)
}

// A collection of values returned by getExternalGroups.
type GetExternalGroupsResultOutput struct{ *pulumi.OutputState }

func (GetExternalGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetExternalGroupsResult)(nil)).Elem()
}

func (o GetExternalGroupsResultOutput) ToGetExternalGroupsResultOutput() GetExternalGroupsResultOutput {
	return o
}

func (o GetExternalGroupsResultOutput) ToGetExternalGroupsResultOutputWithContext(ctx context.Context) GetExternalGroupsResultOutput {
	return o
}

func (o GetExternalGroupsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetExternalGroupsResult] {
	return pulumix.Output[GetExternalGroupsResult]{
		OutputState: o.OutputState,
	}
}

// an array of external groups belonging to the organization. Each group consists of the fields documented below.
func (o GetExternalGroupsResultOutput) ExternalGroups() GetExternalGroupsExternalGroupArrayOutput {
	return o.ApplyT(func(v GetExternalGroupsResult) []GetExternalGroupsExternalGroup { return v.ExternalGroups }).(GetExternalGroupsExternalGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetExternalGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetExternalGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetExternalGroupsResultOutput{})
}
