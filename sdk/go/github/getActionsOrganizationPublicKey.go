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

// Use this data source to retrieve information about a GitHub Actions Organization public key. This data source is required to be used with other GitHub secrets interactions.
// Note that the provider `token` must have admin rights to an organization to retrieve it's action public key.
func GetActionsOrganizationPublicKey(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetActionsOrganizationPublicKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetActionsOrganizationPublicKeyResult
	err := ctx.Invoke("github:index/getActionsOrganizationPublicKey:getActionsOrganizationPublicKey", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getActionsOrganizationPublicKey.
type GetActionsOrganizationPublicKeyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Actual key retrieved.
	Key string `pulumi:"key"`
	// ID of the key that has been retrieved.
	KeyId string `pulumi:"keyId"`
}

func GetActionsOrganizationPublicKeyOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetActionsOrganizationPublicKeyResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetActionsOrganizationPublicKeyResult, error) {
		r, err := GetActionsOrganizationPublicKey(ctx, opts...)
		var s GetActionsOrganizationPublicKeyResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetActionsOrganizationPublicKeyResultOutput)
}

// A collection of values returned by getActionsOrganizationPublicKey.
type GetActionsOrganizationPublicKeyResultOutput struct{ *pulumi.OutputState }

func (GetActionsOrganizationPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActionsOrganizationPublicKeyResult)(nil)).Elem()
}

func (o GetActionsOrganizationPublicKeyResultOutput) ToGetActionsOrganizationPublicKeyResultOutput() GetActionsOrganizationPublicKeyResultOutput {
	return o
}

func (o GetActionsOrganizationPublicKeyResultOutput) ToGetActionsOrganizationPublicKeyResultOutputWithContext(ctx context.Context) GetActionsOrganizationPublicKeyResultOutput {
	return o
}

func (o GetActionsOrganizationPublicKeyResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetActionsOrganizationPublicKeyResult] {
	return pulumix.Output[GetActionsOrganizationPublicKeyResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetActionsOrganizationPublicKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsOrganizationPublicKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Actual key retrieved.
func (o GetActionsOrganizationPublicKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsOrganizationPublicKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

// ID of the key that has been retrieved.
func (o GetActionsOrganizationPublicKeyResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetActionsOrganizationPublicKeyResult) string { return v.KeyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActionsOrganizationPublicKeyResultOutput{})
}
