// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve each organization member's SAML or SCIM user
// attributes.
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
//			_, err := github.GetOrganizationExternalIdentities(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrganizationExternalIdentities(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetOrganizationExternalIdentitiesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationExternalIdentitiesResult
	err := ctx.Invoke("github:index/getOrganizationExternalIdentities:getOrganizationExternalIdentities", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getOrganizationExternalIdentities.
type GetOrganizationExternalIdentitiesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An Array of identities returned from GitHub
	Identities []GetOrganizationExternalIdentitiesIdentity `pulumi:"identities"`
}

func GetOrganizationExternalIdentitiesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetOrganizationExternalIdentitiesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetOrganizationExternalIdentitiesResult, error) {
		r, err := GetOrganizationExternalIdentities(ctx, opts...)
		var s GetOrganizationExternalIdentitiesResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetOrganizationExternalIdentitiesResultOutput)
}

// A collection of values returned by getOrganizationExternalIdentities.
type GetOrganizationExternalIdentitiesResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationExternalIdentitiesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationExternalIdentitiesResult)(nil)).Elem()
}

func (o GetOrganizationExternalIdentitiesResultOutput) ToGetOrganizationExternalIdentitiesResultOutput() GetOrganizationExternalIdentitiesResultOutput {
	return o
}

func (o GetOrganizationExternalIdentitiesResultOutput) ToGetOrganizationExternalIdentitiesResultOutputWithContext(ctx context.Context) GetOrganizationExternalIdentitiesResultOutput {
	return o
}

func (o GetOrganizationExternalIdentitiesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetOrganizationExternalIdentitiesResult] {
	return pulumix.Output[GetOrganizationExternalIdentitiesResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationExternalIdentitiesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationExternalIdentitiesResult) string { return v.Id }).(pulumi.StringOutput)
}

// An Array of identities returned from GitHub
func (o GetOrganizationExternalIdentitiesResultOutput) Identities() GetOrganizationExternalIdentitiesIdentityArrayOutput {
	return o.ApplyT(func(v GetOrganizationExternalIdentitiesResult) []GetOrganizationExternalIdentitiesIdentity {
		return v.Identities
	}).(GetOrganizationExternalIdentitiesIdentityArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationExternalIdentitiesResultOutput{})
}
