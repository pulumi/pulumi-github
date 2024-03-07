// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of codespaces secrets of the organization.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := github.GetCodespacesOrganizationSecrets(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetCodespacesOrganizationSecrets(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCodespacesOrganizationSecretsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCodespacesOrganizationSecretsResult
	err := ctx.Invoke("github:index/getCodespacesOrganizationSecrets:getCodespacesOrganizationSecrets", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCodespacesOrganizationSecrets.
type GetCodespacesOrganizationSecretsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// list of secrets for the repository
	Secrets []GetCodespacesOrganizationSecretsSecret `pulumi:"secrets"`
}

func GetCodespacesOrganizationSecretsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetCodespacesOrganizationSecretsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetCodespacesOrganizationSecretsResult, error) {
		r, err := GetCodespacesOrganizationSecrets(ctx, opts...)
		var s GetCodespacesOrganizationSecretsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetCodespacesOrganizationSecretsResultOutput)
}

// A collection of values returned by getCodespacesOrganizationSecrets.
type GetCodespacesOrganizationSecretsResultOutput struct{ *pulumi.OutputState }

func (GetCodespacesOrganizationSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCodespacesOrganizationSecretsResult)(nil)).Elem()
}

func (o GetCodespacesOrganizationSecretsResultOutput) ToGetCodespacesOrganizationSecretsResultOutput() GetCodespacesOrganizationSecretsResultOutput {
	return o
}

func (o GetCodespacesOrganizationSecretsResultOutput) ToGetCodespacesOrganizationSecretsResultOutputWithContext(ctx context.Context) GetCodespacesOrganizationSecretsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCodespacesOrganizationSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCodespacesOrganizationSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// list of secrets for the repository
func (o GetCodespacesOrganizationSecretsResultOutput) Secrets() GetCodespacesOrganizationSecretsSecretArrayOutput {
	return o.ApplyT(func(v GetCodespacesOrganizationSecretsResult) []GetCodespacesOrganizationSecretsSecret {
		return v.Secrets
	}).(GetCodespacesOrganizationSecretsSecretArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCodespacesOrganizationSecretsResultOutput{})
}
