// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve basic information about a GitHub Organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.GetOrganization(ctx, &GetOrganizationArgs{
// 			Name: "github",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrganization(ctx *pulumi.Context, args *GetOrganizationArgs, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	var rv GetOrganizationResult
	err := ctx.Invoke("github:index/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationArgs struct {
	// The name of the organization account
	Name string `pulumi:"name"`
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	// The description the organization account
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The login of the organization account
	Login string `pulumi:"login"`
	// (`list`) A list with the members of the organization
	Members []string `pulumi:"members"`
	// The name of the organization account
	Name   string `pulumi:"name"`
	NodeId string `pulumi:"nodeId"`
	// The plan name for the organization account
	Plan string `pulumi:"plan"`
	// (`list`) A list with the repositories on the organization
	Repositories []string `pulumi:"repositories"`
}

func GetOrganizationOutput(ctx *pulumi.Context, args GetOrganizationOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrganizationResult, error) {
			args := v.(GetOrganizationArgs)
			r, err := GetOrganization(ctx, &args, opts...)
			return *r, err
		}).(GetOrganizationResultOutput)
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationOutputArgs struct {
	// The name of the organization account
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganization.
type GetOrganizationResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationResult)(nil)).Elem()
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutput() GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutputWithContext(ctx context.Context) GetOrganizationResultOutput {
	return o
}

// The description the organization account
func (o GetOrganizationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The login of the organization account
func (o GetOrganizationResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Login }).(pulumi.StringOutput)
}

// (`list`) A list with the members of the organization
func (o GetOrganizationResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrganizationResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the organization account
func (o GetOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.NodeId }).(pulumi.StringOutput)
}

// The plan name for the organization account
func (o GetOrganizationResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Plan }).(pulumi.StringOutput)
}

// (`list`) A list with the repositories on the organization
func (o GetOrganizationResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrganizationResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationResultOutput{})
}
