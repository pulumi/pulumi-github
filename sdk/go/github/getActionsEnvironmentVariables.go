// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of variables of the repository environment.
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
//			_, err := github.GetActionsEnvironmentVariables(ctx, &github.GetActionsEnvironmentVariablesArgs{
//				Name:        pulumi.StringRef("exampleRepo"),
//				Environment: "exampleEnvironment",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetActionsEnvironmentVariables(ctx *pulumi.Context, args *GetActionsEnvironmentVariablesArgs, opts ...pulumi.InvokeOption) (*GetActionsEnvironmentVariablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetActionsEnvironmentVariablesResult
	err := ctx.Invoke("github:index/getActionsEnvironmentVariables:getActionsEnvironmentVariables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActionsEnvironmentVariables.
type GetActionsEnvironmentVariablesArgs struct {
	Environment string  `pulumi:"environment"`
	FullName    *string `pulumi:"fullName"`
	// Name of the variable
	Name *string `pulumi:"name"`
}

// A collection of values returned by getActionsEnvironmentVariables.
type GetActionsEnvironmentVariablesResult struct {
	Environment string `pulumi:"environment"`
	FullName    string `pulumi:"fullName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the variable
	Name string `pulumi:"name"`
	// list of variables for the environment
	Variables []GetActionsEnvironmentVariablesVariable `pulumi:"variables"`
}

func GetActionsEnvironmentVariablesOutput(ctx *pulumi.Context, args GetActionsEnvironmentVariablesOutputArgs, opts ...pulumi.InvokeOption) GetActionsEnvironmentVariablesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetActionsEnvironmentVariablesResultOutput, error) {
			args := v.(GetActionsEnvironmentVariablesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getActionsEnvironmentVariables:getActionsEnvironmentVariables", args, GetActionsEnvironmentVariablesResultOutput{}, options).(GetActionsEnvironmentVariablesResultOutput), nil
		}).(GetActionsEnvironmentVariablesResultOutput)
}

// A collection of arguments for invoking getActionsEnvironmentVariables.
type GetActionsEnvironmentVariablesOutputArgs struct {
	Environment pulumi.StringInput    `pulumi:"environment"`
	FullName    pulumi.StringPtrInput `pulumi:"fullName"`
	// Name of the variable
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetActionsEnvironmentVariablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsEnvironmentVariablesArgs)(nil)).Elem()
}

// A collection of values returned by getActionsEnvironmentVariables.
type GetActionsEnvironmentVariablesResultOutput struct{ *pulumi.OutputState }

func (GetActionsEnvironmentVariablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsEnvironmentVariablesResult)(nil)).Elem()
}

func (o GetActionsEnvironmentVariablesResultOutput) ToGetActionsEnvironmentVariablesResultOutput() GetActionsEnvironmentVariablesResultOutput {
	return o
}

func (o GetActionsEnvironmentVariablesResultOutput) ToGetActionsEnvironmentVariablesResultOutputWithContext(ctx context.Context) GetActionsEnvironmentVariablesResultOutput {
	return o
}

func (o GetActionsEnvironmentVariablesResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsEnvironmentVariablesResult) string { return v.Environment }).(pulumi.StringOutput)
}

func (o GetActionsEnvironmentVariablesResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsEnvironmentVariablesResult) string { return v.FullName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetActionsEnvironmentVariablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsEnvironmentVariablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the variable
func (o GetActionsEnvironmentVariablesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsEnvironmentVariablesResult) string { return v.Name }).(pulumi.StringOutput)
}

// list of variables for the environment
func (o GetActionsEnvironmentVariablesResultOutput) Variables() GetActionsEnvironmentVariablesVariableArrayOutput {
	return o.ApplyT(func(v GetActionsEnvironmentVariablesResult) []GetActionsEnvironmentVariablesVariable {
		return v.Variables
	}).(GetActionsEnvironmentVariablesVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActionsEnvironmentVariablesResultOutput{})
}
