// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the OpenID Connect subject claim customization template for an organization
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
//			_, err := github.LookupActionsOrganizationOidcSubjectClaimCustomizationTemplate(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupActionsOrganizationOidcSubjectClaimCustomizationTemplate(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResult
	err := ctx.Invoke("github:index/getActionsOrganizationOidcSubjectClaimCustomizationTemplate:getActionsOrganizationOidcSubjectClaimCustomizationTemplate", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getActionsOrganizationOidcSubjectClaimCustomizationTemplate.
type LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of OpenID Connect claim keys.
	IncludeClaimKeys []string `pulumi:"includeClaimKeys"`
}

func LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("github:index/getActionsOrganizationOidcSubjectClaimCustomizationTemplate:getActionsOrganizationOidcSubjectClaimCustomizationTemplate", nil, LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput{}, options).(LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput), nil
	}).(LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput)
}

// A collection of values returned by getActionsOrganizationOidcSubjectClaimCustomizationTemplate.
type LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResult)(nil)).Elem()
}

func (o LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput) ToLookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput() LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput {
	return o
}

func (o LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput) ToLookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutputWithContext(ctx context.Context) LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of OpenID Connect claim keys.
func (o LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput) IncludeClaimKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResult) []string {
		return v.IncludeClaimKeys
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionsOrganizationOidcSubjectClaimCustomizationTemplateResultOutput{})
}
