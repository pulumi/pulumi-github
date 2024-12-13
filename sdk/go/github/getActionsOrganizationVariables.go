// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of variables of the organization.
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
//			_, err := github.GetActionsOrganizationVariables(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetActionsOrganizationVariables(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetActionsOrganizationVariablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetActionsOrganizationVariablesResult
	err := ctx.Invoke("github:index/getActionsOrganizationVariables:getActionsOrganizationVariables", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getActionsOrganizationVariables.
type GetActionsOrganizationVariablesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// list of variables for the repository
	Variables []GetActionsOrganizationVariablesVariable `pulumi:"variables"`
}

func GetActionsOrganizationVariablesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetActionsOrganizationVariablesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetActionsOrganizationVariablesResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("github:index/getActionsOrganizationVariables:getActionsOrganizationVariables", nil, GetActionsOrganizationVariablesResultOutput{}, options).(GetActionsOrganizationVariablesResultOutput), nil
	}).(GetActionsOrganizationVariablesResultOutput)
}

// A collection of values returned by getActionsOrganizationVariables.
type GetActionsOrganizationVariablesResultOutput struct{ *pulumi.OutputState }

func (GetActionsOrganizationVariablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsOrganizationVariablesResult)(nil)).Elem()
}

func (o GetActionsOrganizationVariablesResultOutput) ToGetActionsOrganizationVariablesResultOutput() GetActionsOrganizationVariablesResultOutput {
	return o
}

func (o GetActionsOrganizationVariablesResultOutput) ToGetActionsOrganizationVariablesResultOutputWithContext(ctx context.Context) GetActionsOrganizationVariablesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetActionsOrganizationVariablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsOrganizationVariablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// list of variables for the repository
func (o GetActionsOrganizationVariablesResultOutput) Variables() GetActionsOrganizationVariablesVariableArrayOutput {
	return o.ApplyT(func(v GetActionsOrganizationVariablesResult) []GetActionsOrganizationVariablesVariable {
		return v.Variables
	}).(GetActionsOrganizationVariablesVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActionsOrganizationVariablesResultOutput{})
}
