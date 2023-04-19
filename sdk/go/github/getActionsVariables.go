// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of variables for a GitHub repository.
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
//			_, err := github.GetActionsVariables(ctx, &github.GetActionsVariablesArgs{
//				Name: pulumi.StringRef("example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetActionsVariables(ctx *pulumi.Context, args *GetActionsVariablesArgs, opts ...pulumi.InvokeOption) (*GetActionsVariablesResult, error) {
	var rv GetActionsVariablesResult
	err := ctx.Invoke("github:index/getActionsVariables:getActionsVariables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActionsVariables.
type GetActionsVariablesArgs struct {
	// Full name of the repository (in `org/name` format).
	FullName *string `pulumi:"fullName"`
	// The name of the repository.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getActionsVariables.
type GetActionsVariablesResult struct {
	FullName string `pulumi:"fullName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the variable
	Name string `pulumi:"name"`
	// list of variables for the repository
	Variables []GetActionsVariablesVariable `pulumi:"variables"`
}

func GetActionsVariablesOutput(ctx *pulumi.Context, args GetActionsVariablesOutputArgs, opts ...pulumi.InvokeOption) GetActionsVariablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActionsVariablesResult, error) {
			args := v.(GetActionsVariablesArgs)
			r, err := GetActionsVariables(ctx, &args, opts...)
			var s GetActionsVariablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetActionsVariablesResultOutput)
}

// A collection of arguments for invoking getActionsVariables.
type GetActionsVariablesOutputArgs struct {
	// Full name of the repository (in `org/name` format).
	FullName pulumi.StringPtrInput `pulumi:"fullName"`
	// The name of the repository.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetActionsVariablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsVariablesArgs)(nil)).Elem()
}

// A collection of values returned by getActionsVariables.
type GetActionsVariablesResultOutput struct{ *pulumi.OutputState }

func (GetActionsVariablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsVariablesResult)(nil)).Elem()
}

func (o GetActionsVariablesResultOutput) ToGetActionsVariablesResultOutput() GetActionsVariablesResultOutput {
	return o
}

func (o GetActionsVariablesResultOutput) ToGetActionsVariablesResultOutputWithContext(ctx context.Context) GetActionsVariablesResultOutput {
	return o
}

func (o GetActionsVariablesResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsVariablesResult) string { return v.FullName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetActionsVariablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsVariablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the variable
func (o GetActionsVariablesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsVariablesResult) string { return v.Name }).(pulumi.StringOutput)
}

// list of variables for the repository
func (o GetActionsVariablesResultOutput) Variables() GetActionsVariablesVariableArrayOutput {
	return o.ApplyT(func(v GetActionsVariablesResult) []GetActionsVariablesVariable { return v.Variables }).(GetActionsVariablesVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActionsVariablesResultOutput{})
}