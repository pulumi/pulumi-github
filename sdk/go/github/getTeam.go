// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a GitHub team.
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
//			_, err := github.LookupTeam(ctx, &github.LookupTeamArgs{
//				Slug: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTeam(ctx *pulumi.Context, args *LookupTeamArgs, opts ...pulumi.InvokeOption) (*LookupTeamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTeamResult
	err := ctx.Invoke("github:index/getTeam:getTeam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTeam.
type LookupTeamArgs struct {
	// Type of membership to be requested to fill the list of members. Can be either "all" or "immediate". Default: "all"
	MembershipType *string `pulumi:"membershipType"`
	// Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
	ResultsPerPage *int `pulumi:"resultsPerPage"`
	// The team slug.
	Slug string `pulumi:"slug"`
	// Exclude the members and repositories of the team from the returned result. Defaults to `false`.
	SummaryOnly *bool `pulumi:"summaryOnly"`
}

// A collection of values returned by getTeam.
type LookupTeamResult struct {
	// the team's description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of team members (list of GitHub usernames). Not returned if `summaryOnly = true`
	Members        []string `pulumi:"members"`
	MembershipType *string  `pulumi:"membershipType"`
	// the team's full name.
	Name string `pulumi:"name"`
	// the Node ID of the team.
	NodeId string `pulumi:"nodeId"`
	// the team's permission level.
	Permission string `pulumi:"permission"`
	// the team's privacy type.
	Privacy string `pulumi:"privacy"`
	// List of team repositories (list of repo names). Not returned if `summaryOnly = true`
	Repositories []string `pulumi:"repositories"`
	// List of team repositories (list of `repoId` and `roleName`). Not returned if `summaryOnly = true`
	RepositoriesDetaileds []GetTeamRepositoriesDetailed `pulumi:"repositoriesDetaileds"`
	ResultsPerPage        *int                          `pulumi:"resultsPerPage"`
	Slug                  string                        `pulumi:"slug"`
	SummaryOnly           *bool                         `pulumi:"summaryOnly"`
}

func LookupTeamOutput(ctx *pulumi.Context, args LookupTeamOutputArgs, opts ...pulumi.InvokeOption) LookupTeamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTeamResultOutput, error) {
			args := v.(LookupTeamArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupTeamResult
			secret, err := ctx.InvokePackageRaw("github:index/getTeam:getTeam", args, &rv, "", opts...)
			if err != nil {
				return LookupTeamResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupTeamResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupTeamResultOutput), nil
			}
			return output, nil
		}).(LookupTeamResultOutput)
}

// A collection of arguments for invoking getTeam.
type LookupTeamOutputArgs struct {
	// Type of membership to be requested to fill the list of members. Can be either "all" or "immediate". Default: "all"
	MembershipType pulumi.StringPtrInput `pulumi:"membershipType"`
	// Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
	ResultsPerPage pulumi.IntPtrInput `pulumi:"resultsPerPage"`
	// The team slug.
	Slug pulumi.StringInput `pulumi:"slug"`
	// Exclude the members and repositories of the team from the returned result. Defaults to `false`.
	SummaryOnly pulumi.BoolPtrInput `pulumi:"summaryOnly"`
}

func (LookupTeamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamArgs)(nil)).Elem()
}

// A collection of values returned by getTeam.
type LookupTeamResultOutput struct{ *pulumi.OutputState }

func (LookupTeamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTeamResult)(nil)).Elem()
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutput() LookupTeamResultOutput {
	return o
}

func (o LookupTeamResultOutput) ToLookupTeamResultOutputWithContext(ctx context.Context) LookupTeamResultOutput {
	return o
}

// the team's description.
func (o LookupTeamResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTeamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of team members (list of GitHub usernames). Not returned if `summaryOnly = true`
func (o LookupTeamResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

func (o LookupTeamResultOutput) MembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTeamResult) *string { return v.MembershipType }).(pulumi.StringPtrOutput)
}

// the team's full name.
func (o LookupTeamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Name }).(pulumi.StringOutput)
}

// the Node ID of the team.
func (o LookupTeamResultOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.NodeId }).(pulumi.StringOutput)
}

// the team's permission level.
func (o LookupTeamResultOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Permission }).(pulumi.StringOutput)
}

// the team's privacy type.
func (o LookupTeamResultOutput) Privacy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Privacy }).(pulumi.StringOutput)
}

// List of team repositories (list of repo names). Not returned if `summaryOnly = true`
func (o LookupTeamResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

// List of team repositories (list of `repoId` and `roleName`). Not returned if `summaryOnly = true`
func (o LookupTeamResultOutput) RepositoriesDetaileds() GetTeamRepositoriesDetailedArrayOutput {
	return o.ApplyT(func(v LookupTeamResult) []GetTeamRepositoriesDetailed { return v.RepositoriesDetaileds }).(GetTeamRepositoriesDetailedArrayOutput)
}

func (o LookupTeamResultOutput) ResultsPerPage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTeamResult) *int { return v.ResultsPerPage }).(pulumi.IntPtrOutput)
}

func (o LookupTeamResultOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTeamResult) string { return v.Slug }).(pulumi.StringOutput)
}

func (o LookupTeamResultOutput) SummaryOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTeamResult) *bool { return v.SummaryOnly }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTeamResultOutput{})
}
