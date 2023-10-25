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
//
//	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := github.GetOrganizationTeams(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// To retrieve only the team's at the root of the organization:
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
//			_, err := github.GetOrganizationTeams(ctx, &github.GetOrganizationTeamsArgs{
//				RootTeamsOnly: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrganizationTeams(ctx *pulumi.Context, args *GetOrganizationTeamsArgs, opts ...pulumi.InvokeOption) (*GetOrganizationTeamsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationTeamsResult
	err := ctx.Invoke("github:index/getOrganizationTeams:getOrganizationTeams", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationTeams.
type GetOrganizationTeamsArgs struct {
	// (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
	ResultsPerPage *int `pulumi:"resultsPerPage"`
	// (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
	RootTeamsOnly *bool `pulumi:"rootTeamsOnly"`
	// (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
	SummaryOnly *bool `pulumi:"summaryOnly"`
}

// A collection of values returned by getOrganizationTeams.
type GetOrganizationTeamsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
	ResultsPerPage *int `pulumi:"resultsPerPage"`
	// (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
	RootTeamsOnly *bool `pulumi:"rootTeamsOnly"`
	// (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
	SummaryOnly *bool `pulumi:"summaryOnly"`
	// (Required) An Array of GitHub Teams.  Each `team` block consists of the fields documented below.
	Teams []GetOrganizationTeamsTeam `pulumi:"teams"`
}

func GetOrganizationTeamsOutput(ctx *pulumi.Context, args GetOrganizationTeamsOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationTeamsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrganizationTeamsResult, error) {
			args := v.(GetOrganizationTeamsArgs)
			r, err := GetOrganizationTeams(ctx, &args, opts...)
			var s GetOrganizationTeamsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrganizationTeamsResultOutput)
}

// A collection of arguments for invoking getOrganizationTeams.
type GetOrganizationTeamsOutputArgs struct {
	// (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
	ResultsPerPage pulumi.IntPtrInput `pulumi:"resultsPerPage"`
	// (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
	RootTeamsOnly pulumi.BoolPtrInput `pulumi:"rootTeamsOnly"`
	// (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
	SummaryOnly pulumi.BoolPtrInput `pulumi:"summaryOnly"`
}

func (GetOrganizationTeamsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationTeamsArgs)(nil)).Elem()
}

// A collection of values returned by getOrganizationTeams.
type GetOrganizationTeamsResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationTeamsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationTeamsResult)(nil)).Elem()
}

func (o GetOrganizationTeamsResultOutput) ToGetOrganizationTeamsResultOutput() GetOrganizationTeamsResultOutput {
	return o
}

func (o GetOrganizationTeamsResultOutput) ToGetOrganizationTeamsResultOutputWithContext(ctx context.Context) GetOrganizationTeamsResultOutput {
	return o
}

func (o GetOrganizationTeamsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetOrganizationTeamsResult] {
	return pulumix.Output[GetOrganizationTeamsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationTeamsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationTeamsResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
func (o GetOrganizationTeamsResultOutput) ResultsPerPage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetOrganizationTeamsResult) *int { return v.ResultsPerPage }).(pulumi.IntPtrOutput)
}

// (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
func (o GetOrganizationTeamsResultOutput) RootTeamsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetOrganizationTeamsResult) *bool { return v.RootTeamsOnly }).(pulumi.BoolPtrOutput)
}

// (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
func (o GetOrganizationTeamsResultOutput) SummaryOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetOrganizationTeamsResult) *bool { return v.SummaryOnly }).(pulumi.BoolPtrOutput)
}

// (Required) An Array of GitHub Teams.  Each `team` block consists of the fields documented below.
func (o GetOrganizationTeamsResultOutput) Teams() GetOrganizationTeamsTeamArrayOutput {
	return o.ApplyT(func(v GetOrganizationTeamsResult) []GetOrganizationTeamsTeam { return v.Teams }).(GetOrganizationTeamsTeamArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationTeamsResultOutput{})
}
