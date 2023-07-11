// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of codespaces secrets of the organization.
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
//			_, err := github.GetCodespacesOrganizationSecrets(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCodespacesOrganizationSecrets(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCodespacesOrganizationSecretsResult, error) {
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