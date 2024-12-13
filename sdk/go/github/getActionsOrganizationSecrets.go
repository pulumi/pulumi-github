// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of secrets of the organization.
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
//			_, err := github.GetActionsOrganizationSecrets(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetActionsOrganizationSecrets(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetActionsOrganizationSecretsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetActionsOrganizationSecretsResult
	err := ctx.Invoke("github:index/getActionsOrganizationSecrets:getActionsOrganizationSecrets", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getActionsOrganizationSecrets.
type GetActionsOrganizationSecretsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// list of secrets for the repository
	Secrets []GetActionsOrganizationSecretsSecret `pulumi:"secrets"`
}

func GetActionsOrganizationSecretsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetActionsOrganizationSecretsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetActionsOrganizationSecretsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("github:index/getActionsOrganizationSecrets:getActionsOrganizationSecrets", nil, GetActionsOrganizationSecretsResultOutput{}, options).(GetActionsOrganizationSecretsResultOutput), nil
	}).(GetActionsOrganizationSecretsResultOutput)
}

// A collection of values returned by getActionsOrganizationSecrets.
type GetActionsOrganizationSecretsResultOutput struct{ *pulumi.OutputState }

func (GetActionsOrganizationSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsOrganizationSecretsResult)(nil)).Elem()
}

func (o GetActionsOrganizationSecretsResultOutput) ToGetActionsOrganizationSecretsResultOutput() GetActionsOrganizationSecretsResultOutput {
	return o
}

func (o GetActionsOrganizationSecretsResultOutput) ToGetActionsOrganizationSecretsResultOutputWithContext(ctx context.Context) GetActionsOrganizationSecretsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetActionsOrganizationSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsOrganizationSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// list of secrets for the repository
func (o GetActionsOrganizationSecretsResultOutput) Secrets() GetActionsOrganizationSecretsSecretArrayOutput {
	return o.ApplyT(func(v GetActionsOrganizationSecretsResult) []GetActionsOrganizationSecretsSecret { return v.Secrets }).(GetActionsOrganizationSecretsSecretArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActionsOrganizationSecretsResultOutput{})
}
