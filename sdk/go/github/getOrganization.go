// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve basic information about a GitHub Organization.
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
//			_, err := github.GetOrganization(ctx, &github.GetOrganizationArgs{
//				Name: "github",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrganization(ctx *pulumi.Context, args *GetOrganizationArgs, opts ...pulumi.InvokeOption) (*GetOrganizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationResult
	err := ctx.Invoke("github:index/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationArgs struct {
	// Whether or not to include archived repos in the `repositories` list
	IgnoreArchivedRepos *bool `pulumi:"ignoreArchivedRepos"`
	// The organization's public profile name
	Name string `pulumi:"name"`
}

// A collection of values returned by getOrganization.
type GetOrganizationResult struct {
	// Whether advanced security is enabled for new repositories.
	AdvancedSecurityEnabledForNewRepositories bool `pulumi:"advancedSecurityEnabledForNewRepositories"`
	// Default permission level members have for organization repositories.
	DefaultRepositoryPermission string `pulumi:"defaultRepositoryPermission"`
	// Whether Dependabot alerts is automatically enabled for new repositories.
	DependabotAlertsEnabledForNewRepositories bool `pulumi:"dependabotAlertsEnabledForNewRepositories"`
	// Whether Dependabot security updates is automatically enabled for new repositories.
	DependabotSecurityUpdatesEnabledForNewRepositories bool `pulumi:"dependabotSecurityUpdatesEnabledForNewRepositories"`
	// Whether dependency graph is automatically enabled for new repositories.
	DependencyGraphEnabledForNewRepositories bool `pulumi:"dependencyGraphEnabledForNewRepositories"`
	// The organization account description
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string `pulumi:"id"`
	IgnoreArchivedRepos *bool  `pulumi:"ignoreArchivedRepos"`
	// The members login
	Login string `pulumi:"login"`
	// **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
	//
	// Deprecated: Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.
	Members []string `pulumi:"members"`
	// The type of repository allowed to be created by members of the organization. Can be one of `ALL`, `PUBLIC`, `PRIVATE`, `NONE`.
	MembersAllowedRepositoryCreationType string `pulumi:"membersAllowedRepositoryCreationType"`
	// Whether organization members can create internal repositories.
	MembersCanCreateInternalRepositories bool `pulumi:"membersCanCreateInternalRepositories"`
	// Whether organization members can create pages sites.
	MembersCanCreatePages bool `pulumi:"membersCanCreatePages"`
	// Whether organization members can create private pages sites.
	MembersCanCreatePrivatePages bool `pulumi:"membersCanCreatePrivatePages"`
	// Whether organization members can create private repositories.
	MembersCanCreatePrivateRepositories bool `pulumi:"membersCanCreatePrivateRepositories"`
	// Whether organization members can create public pages sites.
	MembersCanCreatePublicPages bool `pulumi:"membersCanCreatePublicPages"`
	// Whether organization members can create public repositories.
	MembersCanCreatePublicRepositories bool `pulumi:"membersCanCreatePublicRepositories"`
	// Whether non-admin organization members can create repositories.
	MembersCanCreateRepositories bool `pulumi:"membersCanCreateRepositories"`
	// Whether organization members can create private repository forks.
	MembersCanForkPrivateRepositories bool `pulumi:"membersCanForkPrivateRepositories"`
	// The organization's public profile name
	Name string `pulumi:"name"`
	// GraphQL global node ID for use with the v4 API
	NodeId string `pulumi:"nodeId"`
	// The organization's name as used in URLs and the API
	Orgname string `pulumi:"orgname"`
	// The organization account plan name
	Plan string `pulumi:"plan"`
	// (`list`) A list of the full names of the repositories in the organization formatted as `owner/name` strings
	Repositories []string `pulumi:"repositories"`
	// Whether secret scanning is automatically enabled for new repositories.
	SecretScanningEnabledForNewRepositories bool `pulumi:"secretScanningEnabledForNewRepositories"`
	// Whether secret scanning push protection is automatically enabled for new repositories.
	SecretScanningPushProtectionEnabledForNewRepositories bool `pulumi:"secretScanningPushProtectionEnabledForNewRepositories"`
	// Whether two-factor authentication is required for all members of the organization.
	TwoFactorRequirementEnabled bool `pulumi:"twoFactorRequirementEnabled"`
	// (`list`) A list with the members of the organization with following fields:
	Users []map[string]string `pulumi:"users"`
	// Whether organization members must sign all commits.
	WebCommitSignoffRequired bool `pulumi:"webCommitSignoffRequired"`
}

func GetOrganizationOutput(ctx *pulumi.Context, args GetOrganizationOutputArgs, opts ...pulumi.InvokeOption) GetOrganizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrganizationResult, error) {
			args := v.(GetOrganizationArgs)
			r, err := GetOrganization(ctx, &args, opts...)
			var s GetOrganizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrganizationResultOutput)
}

// A collection of arguments for invoking getOrganization.
type GetOrganizationOutputArgs struct {
	// Whether or not to include archived repos in the `repositories` list
	IgnoreArchivedRepos pulumi.BoolPtrInput `pulumi:"ignoreArchivedRepos"`
	// The organization's public profile name
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationArgs)(nil)).Elem()
}

// A collection of values returned by getOrganization.
type GetOrganizationResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationResult)(nil)).Elem()
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutput() GetOrganizationResultOutput {
	return o
}

func (o GetOrganizationResultOutput) ToGetOrganizationResultOutputWithContext(ctx context.Context) GetOrganizationResultOutput {
	return o
}

// Whether advanced security is enabled for new repositories.
func (o GetOrganizationResultOutput) AdvancedSecurityEnabledForNewRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.AdvancedSecurityEnabledForNewRepositories }).(pulumi.BoolOutput)
}

// Default permission level members have for organization repositories.
func (o GetOrganizationResultOutput) DefaultRepositoryPermission() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.DefaultRepositoryPermission }).(pulumi.StringOutput)
}

// Whether Dependabot alerts is automatically enabled for new repositories.
func (o GetOrganizationResultOutput) DependabotAlertsEnabledForNewRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.DependabotAlertsEnabledForNewRepositories }).(pulumi.BoolOutput)
}

// Whether Dependabot security updates is automatically enabled for new repositories.
func (o GetOrganizationResultOutput) DependabotSecurityUpdatesEnabledForNewRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.DependabotSecurityUpdatesEnabledForNewRepositories }).(pulumi.BoolOutput)
}

// Whether dependency graph is automatically enabled for new repositories.
func (o GetOrganizationResultOutput) DependencyGraphEnabledForNewRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.DependencyGraphEnabledForNewRepositories }).(pulumi.BoolOutput)
}

// The organization account description
func (o GetOrganizationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrganizationResultOutput) IgnoreArchivedRepos() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetOrganizationResult) *bool { return v.IgnoreArchivedRepos }).(pulumi.BoolPtrOutput)
}

// The members login
func (o GetOrganizationResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Login }).(pulumi.StringOutput)
}

// **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
//
// Deprecated: Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.
func (o GetOrganizationResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrganizationResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// The type of repository allowed to be created by members of the organization. Can be one of `ALL`, `PUBLIC`, `PRIVATE`, `NONE`.
func (o GetOrganizationResultOutput) MembersAllowedRepositoryCreationType() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.MembersAllowedRepositoryCreationType }).(pulumi.StringOutput)
}

// Whether organization members can create internal repositories.
func (o GetOrganizationResultOutput) MembersCanCreateInternalRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanCreateInternalRepositories }).(pulumi.BoolOutput)
}

// Whether organization members can create pages sites.
func (o GetOrganizationResultOutput) MembersCanCreatePages() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanCreatePages }).(pulumi.BoolOutput)
}

// Whether organization members can create private pages sites.
func (o GetOrganizationResultOutput) MembersCanCreatePrivatePages() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanCreatePrivatePages }).(pulumi.BoolOutput)
}

// Whether organization members can create private repositories.
func (o GetOrganizationResultOutput) MembersCanCreatePrivateRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanCreatePrivateRepositories }).(pulumi.BoolOutput)
}

// Whether organization members can create public pages sites.
func (o GetOrganizationResultOutput) MembersCanCreatePublicPages() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanCreatePublicPages }).(pulumi.BoolOutput)
}

// Whether organization members can create public repositories.
func (o GetOrganizationResultOutput) MembersCanCreatePublicRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanCreatePublicRepositories }).(pulumi.BoolOutput)
}

// Whether non-admin organization members can create repositories.
func (o GetOrganizationResultOutput) MembersCanCreateRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanCreateRepositories }).(pulumi.BoolOutput)
}

// Whether organization members can create private repository forks.
func (o GetOrganizationResultOutput) MembersCanForkPrivateRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.MembersCanForkPrivateRepositories }).(pulumi.BoolOutput)
}

// The organization's public profile name
func (o GetOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

// GraphQL global node ID for use with the v4 API
func (o GetOrganizationResultOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.NodeId }).(pulumi.StringOutput)
}

// The organization's name as used in URLs and the API
func (o GetOrganizationResultOutput) Orgname() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Orgname }).(pulumi.StringOutput)
}

// The organization account plan name
func (o GetOrganizationResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationResult) string { return v.Plan }).(pulumi.StringOutput)
}

// (`list`) A list of the full names of the repositories in the organization formatted as `owner/name` strings
func (o GetOrganizationResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOrganizationResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

// Whether secret scanning is automatically enabled for new repositories.
func (o GetOrganizationResultOutput) SecretScanningEnabledForNewRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.SecretScanningEnabledForNewRepositories }).(pulumi.BoolOutput)
}

// Whether secret scanning push protection is automatically enabled for new repositories.
func (o GetOrganizationResultOutput) SecretScanningPushProtectionEnabledForNewRepositories() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.SecretScanningPushProtectionEnabledForNewRepositories }).(pulumi.BoolOutput)
}

// Whether two-factor authentication is required for all members of the organization.
func (o GetOrganizationResultOutput) TwoFactorRequirementEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.TwoFactorRequirementEnabled }).(pulumi.BoolOutput)
}

// (`list`) A list with the members of the organization with following fields:
func (o GetOrganizationResultOutput) Users() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v GetOrganizationResult) []map[string]string { return v.Users }).(pulumi.StringMapArrayOutput)
}

// Whether organization members must sign all commits.
func (o GetOrganizationResultOutput) WebCommitSignoffRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrganizationResult) bool { return v.WebCommitSignoffRequired }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationResultOutput{})
}
