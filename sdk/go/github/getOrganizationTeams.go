// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about all GitHub teams in an organization.
//
// ## Example Usage
//
// To retrieve *all* teams of the organization:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.GetOrganizationTeams(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To retrieve only the team's at the root of the organization:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		_, err := github.GetOrganizationTeams(ctx, &github.GetOrganizationTeamsArgs{
// 			RootTeamsOnly: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrganizationTeams(ctx *pulumi.Context, args *GetOrganizationTeamsArgs, opts ...pulumi.InvokeOption) (*GetOrganizationTeamsResult, error) {
	var rv GetOrganizationTeamsResult
	err := ctx.Invoke("github:index/getOrganizationTeams:getOrganizationTeams", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationTeams.
type GetOrganizationTeamsArgs struct {
	// Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
	RootTeamsOnly *bool `pulumi:"rootTeamsOnly"`
}

// A collection of values returned by getOrganizationTeams.
type GetOrganizationTeamsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
	RootTeamsOnly *bool `pulumi:"rootTeamsOnly"`
	// An Array of GitHub Teams.  Each `team` block consists of the fields documented below.
	// ***
	Teams []GetOrganizationTeamsTeam `pulumi:"teams"`
}