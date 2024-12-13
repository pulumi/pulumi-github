// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a list of repository branch protection rules.
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
//			_, err := github.GetBranchProtectionRules(ctx, &github.GetBranchProtectionRulesArgs{
//				Repository: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetBranchProtectionRules(ctx *pulumi.Context, args *GetBranchProtectionRulesArgs, opts ...pulumi.InvokeOption) (*GetBranchProtectionRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBranchProtectionRulesResult
	err := ctx.Invoke("github:index/getBranchProtectionRules:getBranchProtectionRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBranchProtectionRules.
type GetBranchProtectionRulesArgs struct {
	// The GitHub repository name.
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getBranchProtectionRules.
type GetBranchProtectionRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Repository string `pulumi:"repository"`
	// Collection of Branch Protection Rules. Each of the results conforms to the following scheme:
	Rules []GetBranchProtectionRulesRule `pulumi:"rules"`
}

func GetBranchProtectionRulesOutput(ctx *pulumi.Context, args GetBranchProtectionRulesOutputArgs, opts ...pulumi.InvokeOption) GetBranchProtectionRulesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBranchProtectionRulesResultOutput, error) {
			args := v.(GetBranchProtectionRulesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getBranchProtectionRules:getBranchProtectionRules", args, GetBranchProtectionRulesResultOutput{}, options).(GetBranchProtectionRulesResultOutput), nil
		}).(GetBranchProtectionRulesResultOutput)
}

// A collection of arguments for invoking getBranchProtectionRules.
type GetBranchProtectionRulesOutputArgs struct {
	// The GitHub repository name.
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (GetBranchProtectionRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBranchProtectionRulesArgs)(nil)).Elem()
}

// A collection of values returned by getBranchProtectionRules.
type GetBranchProtectionRulesResultOutput struct{ *pulumi.OutputState }

func (GetBranchProtectionRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBranchProtectionRulesResult)(nil)).Elem()
}

func (o GetBranchProtectionRulesResultOutput) ToGetBranchProtectionRulesResultOutput() GetBranchProtectionRulesResultOutput {
	return o
}

func (o GetBranchProtectionRulesResultOutput) ToGetBranchProtectionRulesResultOutputWithContext(ctx context.Context) GetBranchProtectionRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetBranchProtectionRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBranchProtectionRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBranchProtectionRulesResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GetBranchProtectionRulesResult) string { return v.Repository }).(pulumi.StringOutput)
}

// Collection of Branch Protection Rules. Each of the results conforms to the following scheme:
func (o GetBranchProtectionRulesResultOutput) Rules() GetBranchProtectionRulesRuleArrayOutput {
	return o.ApplyT(func(v GetBranchProtectionRulesResult) []GetBranchProtectionRulesRule { return v.Rules }).(GetBranchProtectionRulesRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBranchProtectionRulesResultOutput{})
}
