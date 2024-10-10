// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about the IP allow list of an organization.
// The allow list for IP addresses will block access to private resources via the web, API,
// and Git from any IP addresses that are not on the allow list.
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
//			_, err := github.GetOrganizationIpAllowList(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrganizationIpAllowList(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOrganizationIpAllowListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationIpAllowListResult
	err := ctx.Invoke("github:index/getOrganizationIpAllowList:getOrganizationIpAllowList", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganizationIpAllowList.
type GetOrganizationIpAllowListResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An Array of allowed IP addresses.
	// ***
	IpAllowLists []GetOrganizationIpAllowListIpAllowList `pulumi:"ipAllowLists"`
}

func GetOrganizationIpAllowListOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetOrganizationIpAllowListResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetOrganizationIpAllowListResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetOrganizationIpAllowListResult
		secret, err := ctx.InvokePackageRaw("github:index/getOrganizationIpAllowList:getOrganizationIpAllowList", nil, &rv, "", opts...)
		if err != nil {
			return GetOrganizationIpAllowListResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetOrganizationIpAllowListResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetOrganizationIpAllowListResultOutput), nil
		}
		return output, nil
	}).(GetOrganizationIpAllowListResultOutput)
}

// A collection of values returned by getOrganizationIpAllowList.
type GetOrganizationIpAllowListResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationIpAllowListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationIpAllowListResult)(nil)).Elem()
}

func (o GetOrganizationIpAllowListResultOutput) ToGetOrganizationIpAllowListResultOutput() GetOrganizationIpAllowListResultOutput {
	return o
}

func (o GetOrganizationIpAllowListResultOutput) ToGetOrganizationIpAllowListResultOutputWithContext(ctx context.Context) GetOrganizationIpAllowListResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationIpAllowListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationIpAllowListResult) string { return v.Id }).(pulumi.StringOutput)
}

// An Array of allowed IP addresses.
// ***
func (o GetOrganizationIpAllowListResultOutput) IpAllowLists() GetOrganizationIpAllowListIpAllowListArrayOutput {
	return o.ApplyT(func(v GetOrganizationIpAllowListResult) []GetOrganizationIpAllowListIpAllowList {
		return v.IpAllowLists
	}).(GetOrganizationIpAllowListIpAllowListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationIpAllowListResultOutput{})
}
