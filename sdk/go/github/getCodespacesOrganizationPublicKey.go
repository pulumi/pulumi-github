// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a GitHub Codespaces Organization public key. This data source is required to be used with other GitHub secrets interactions.
// Note that the provider `token` must have admin rights to an organization to retrieve it's Codespaces public key.
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
//			_, err := github.GetCodespacesOrganizationPublicKey(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCodespacesOrganizationPublicKey(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCodespacesOrganizationPublicKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCodespacesOrganizationPublicKeyResult
	err := ctx.Invoke("github:index/getCodespacesOrganizationPublicKey:getCodespacesOrganizationPublicKey", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCodespacesOrganizationPublicKey.
type GetCodespacesOrganizationPublicKeyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Actual key retrieved.
	Key string `pulumi:"key"`
	// ID of the key that has been retrieved.
	KeyId string `pulumi:"keyId"`
}
