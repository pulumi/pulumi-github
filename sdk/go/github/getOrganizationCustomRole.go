// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a custom role in a GitHub Organization.
//
// > Note: Custom roles are currently only available in GitHub Enterprise Cloud.
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
//			_, err := github.LookupOrganizationCustomRole(ctx, &github.LookupOrganizationCustomRoleArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOrganizationCustomRole(ctx *pulumi.Context, args *LookupOrganizationCustomRoleArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationCustomRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationCustomRoleResult
	err := ctx.Invoke("github:index/getOrganizationCustomRole:getOrganizationCustomRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationCustomRole.
type LookupOrganizationCustomRoleArgs struct {
	// The name of the custom role.
	Name string `pulumi:"name"`
}

// A collection of values returned by getOrganizationCustomRole.
type LookupOrganizationCustomRoleResult struct {
	// The system role from which the role inherits permissions.
	BaseRole string `pulumi:"baseRole"`
	// The description for the custom role.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// A list of additional permissions included in this role.
	Permissions []string `pulumi:"permissions"`
}

func LookupOrganizationCustomRoleOutput(ctx *pulumi.Context, args LookupOrganizationCustomRoleOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationCustomRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationCustomRoleResultOutput, error) {
			args := v.(LookupOrganizationCustomRoleArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupOrganizationCustomRoleResult
			secret, err := ctx.InvokePackageRaw("github:index/getOrganizationCustomRole:getOrganizationCustomRole", args, &rv, "", opts...)
			if err != nil {
				return LookupOrganizationCustomRoleResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupOrganizationCustomRoleResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupOrganizationCustomRoleResultOutput), nil
			}
			return output, nil
		}).(LookupOrganizationCustomRoleResultOutput)
}

// A collection of arguments for invoking getOrganizationCustomRole.
type LookupOrganizationCustomRoleOutputArgs struct {
	// The name of the custom role.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupOrganizationCustomRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationCustomRoleArgs)(nil)).Elem()
}

// A collection of values returned by getOrganizationCustomRole.
type LookupOrganizationCustomRoleResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationCustomRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationCustomRoleResult)(nil)).Elem()
}

func (o LookupOrganizationCustomRoleResultOutput) ToLookupOrganizationCustomRoleResultOutput() LookupOrganizationCustomRoleResultOutput {
	return o
}

func (o LookupOrganizationCustomRoleResultOutput) ToLookupOrganizationCustomRoleResultOutputWithContext(ctx context.Context) LookupOrganizationCustomRoleResultOutput {
	return o
}

// The system role from which the role inherits permissions.
func (o LookupOrganizationCustomRoleResultOutput) BaseRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomRoleResult) string { return v.BaseRole }).(pulumi.StringOutput)
}

// The description for the custom role.
func (o LookupOrganizationCustomRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrganizationCustomRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrganizationCustomRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationCustomRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// A list of additional permissions included in this role.
func (o LookupOrganizationCustomRoleResultOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrganizationCustomRoleResult) []string { return v.Permissions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationCustomRoleResultOutput{})
}
