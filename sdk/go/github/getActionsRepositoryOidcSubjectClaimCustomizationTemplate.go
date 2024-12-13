// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the OpenID Connect subject claim customization template for a repository
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
//			_, err := github.LookupActionsRepositoryOidcSubjectClaimCustomizationTemplate(ctx, &github.LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs{
//				Name: "example_repository",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupActionsRepositoryOidcSubjectClaimCustomizationTemplate(ctx *pulumi.Context, args *LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs, opts ...pulumi.InvokeOption) (*LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult
	err := ctx.Invoke("github:index/getActionsRepositoryOidcSubjectClaimCustomizationTemplate:getActionsRepositoryOidcSubjectClaimCustomizationTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActionsRepositoryOidcSubjectClaimCustomizationTemplate.
type LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs struct {
	// Name of the repository to get the OpenID Connect subject claim customization template for.
	Name string `pulumi:"name"`
}

// A collection of values returned by getActionsRepositoryOidcSubjectClaimCustomizationTemplate.
type LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of OpenID Connect claim keys.
	IncludeClaimKeys []string `pulumi:"includeClaimKeys"`
	Name             string   `pulumi:"name"`
	// Whether the repository uses the default template.
	UseDefault bool `pulumi:"useDefault"`
}

func LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput(ctx *pulumi.Context, args LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput, error) {
			args := v.(LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("github:index/getActionsRepositoryOidcSubjectClaimCustomizationTemplate:getActionsRepositoryOidcSubjectClaimCustomizationTemplate", args, LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput{}, options).(LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput), nil
		}).(LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput)
}

// A collection of arguments for invoking getActionsRepositoryOidcSubjectClaimCustomizationTemplate.
type LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputArgs struct {
	// Name of the repository to get the OpenID Connect subject claim customization template for.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getActionsRepositoryOidcSubjectClaimCustomizationTemplate.
type LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult)(nil)).Elem()
}

func (o LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput) ToLookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput() LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput {
	return o
}

func (o LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput) ToLookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutputWithContext(ctx context.Context) LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of OpenID Connect claim keys.
func (o LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput) IncludeClaimKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult) []string {
		return v.IncludeClaimKeys
	}).(pulumi.StringArrayOutput)
}

func (o LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether the repository uses the default template.
func (o LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput) UseDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResult) bool { return v.UseDefault }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionsRepositoryOidcSubjectClaimCustomizationTemplateResultOutput{})
}
