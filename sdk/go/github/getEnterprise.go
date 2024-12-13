// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve basic information about a GitHub enterprise.
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
//			_, err := github.GetEnterprise(ctx, &github.GetEnterpriseArgs{
//				Slug: "example-co",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetEnterprise(ctx *pulumi.Context, args *GetEnterpriseArgs, opts ...pulumi.InvokeOption) (*GetEnterpriseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEnterpriseResult
	err := ctx.Invoke("github:index/getEnterprise:getEnterprise", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnterprise.
type GetEnterpriseArgs struct {
	// The URL slug identifying the enterprise.
	Slug string `pulumi:"slug"`
}

// A collection of values returned by getEnterprise.
type GetEnterpriseResult struct {
	// The time the enterprise was created.
	CreatedAt string `pulumi:"createdAt"`
	// The database ID of the enterprise.
	DatabaseId int `pulumi:"databaseId"`
	// The description of the enterprise.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the enterprise.
	Name string `pulumi:"name"`
	// The URL slug identifying the enterprise.
	Slug string `pulumi:"slug"`
	// The url for the enterprise.
	Url string `pulumi:"url"`
}

func GetEnterpriseOutput(ctx *pulumi.Context, args GetEnterpriseOutputArgs, opts ...pulumi.InvokeOption) GetEnterpriseResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetEnterpriseResultOutput, error) {
			args := v.(GetEnterpriseArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getEnterprise:getEnterprise", args, GetEnterpriseResultOutput{}, options).(GetEnterpriseResultOutput), nil
		}).(GetEnterpriseResultOutput)
}

// A collection of arguments for invoking getEnterprise.
type GetEnterpriseOutputArgs struct {
	// The URL slug identifying the enterprise.
	Slug pulumi.StringInput `pulumi:"slug"`
}

func (GetEnterpriseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnterpriseArgs)(nil)).Elem()
}

// A collection of values returned by getEnterprise.
type GetEnterpriseResultOutput struct{ *pulumi.OutputState }

func (GetEnterpriseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnterpriseResult)(nil)).Elem()
}

func (o GetEnterpriseResultOutput) ToGetEnterpriseResultOutput() GetEnterpriseResultOutput {
	return o
}

func (o GetEnterpriseResultOutput) ToGetEnterpriseResultOutputWithContext(ctx context.Context) GetEnterpriseResultOutput {
	return o
}

// The time the enterprise was created.
func (o GetEnterpriseResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The database ID of the enterprise.
func (o GetEnterpriseResultOutput) DatabaseId() pulumi.IntOutput {
	return o.ApplyT(func(v GetEnterpriseResult) int { return v.DatabaseId }).(pulumi.IntOutput)
}

// The description of the enterprise.
func (o GetEnterpriseResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEnterpriseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the enterprise.
func (o GetEnterpriseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL slug identifying the enterprise.
func (o GetEnterpriseResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseResult) string { return v.Slug }).(pulumi.StringOutput)
}

// The url for the enterprise.
func (o GetEnterpriseResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnterpriseResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEnterpriseResultOutput{})
}
