// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource manages the team settings (in particular the request review delegation settings) within the organization
//
// Creating this resource will alter the team Code Review settings.
//
// The team must both belong to the same organization configured in the provider on GitHub.
//
// > **Note**: This resource relies on the v4 GraphQl GitHub API. If this API is not available, or the Stone Crop schema preview is not available, then this resource will not work as intended.
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
//			// Add a repository to the team
//			someTeam, err := github.NewTeam(ctx, "some_team", &github.TeamArgs{
//				Name:        pulumi.String("SomeTeam"),
//				Description: pulumi.String("Some cool team"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewTeamSettings(ctx, "code_review_settings", &github.TeamSettingsArgs{
//				TeamId: someTeam.ID(),
//				ReviewRequestDelegation: &github.TeamSettingsReviewRequestDelegationArgs{
//					Algorithm:   pulumi.String("ROUND_ROBIN"),
//					MemberCount: pulumi.Int(1),
//					Notify:      pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// GitHub Teams can be imported using the GitHub team ID, or the team slug e.g.
//
// ```sh
// $ pulumi import github:index/teamSettings:TeamSettings code_review_settings 1234567
// ```
// or,
//
// ```sh
// $ pulumi import github:index/teamSettings:TeamSettings code_review_settings SomeTeam
// ```
type TeamSettings struct {
	pulumi.CustomResourceState

	// The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
	ReviewRequestDelegation TeamSettingsReviewRequestDelegationPtrOutput `pulumi:"reviewRequestDelegation"`
	// The GitHub team id or the GitHub team slug
	TeamId pulumi.StringOutput `pulumi:"teamId"`
	// The slug of the Team within the Organization.
	TeamSlug pulumi.StringOutput `pulumi:"teamSlug"`
	// The unique ID of the Team on GitHub. Corresponds to the ID of the 'github_team_settings' resource.
	TeamUid pulumi.StringOutput `pulumi:"teamUid"`
}

// NewTeamSettings registers a new resource with the given unique name, arguments, and options.
func NewTeamSettings(ctx *pulumi.Context,
	name string, args *TeamSettingsArgs, opts ...pulumi.ResourceOption) (*TeamSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TeamSettings
	err := ctx.RegisterResource("github:index/teamSettings:TeamSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamSettings gets an existing TeamSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamSettingsState, opts ...pulumi.ResourceOption) (*TeamSettings, error) {
	var resource TeamSettings
	err := ctx.ReadResource("github:index/teamSettings:TeamSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamSettings resources.
type teamSettingsState struct {
	// The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
	ReviewRequestDelegation *TeamSettingsReviewRequestDelegation `pulumi:"reviewRequestDelegation"`
	// The GitHub team id or the GitHub team slug
	TeamId *string `pulumi:"teamId"`
	// The slug of the Team within the Organization.
	TeamSlug *string `pulumi:"teamSlug"`
	// The unique ID of the Team on GitHub. Corresponds to the ID of the 'github_team_settings' resource.
	TeamUid *string `pulumi:"teamUid"`
}

type TeamSettingsState struct {
	// The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
	ReviewRequestDelegation TeamSettingsReviewRequestDelegationPtrInput
	// The GitHub team id or the GitHub team slug
	TeamId pulumi.StringPtrInput
	// The slug of the Team within the Organization.
	TeamSlug pulumi.StringPtrInput
	// The unique ID of the Team on GitHub. Corresponds to the ID of the 'github_team_settings' resource.
	TeamUid pulumi.StringPtrInput
}

func (TeamSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamSettingsState)(nil)).Elem()
}

type teamSettingsArgs struct {
	// The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
	ReviewRequestDelegation *TeamSettingsReviewRequestDelegation `pulumi:"reviewRequestDelegation"`
	// The GitHub team id or the GitHub team slug
	TeamId string `pulumi:"teamId"`
}

// The set of arguments for constructing a TeamSettings resource.
type TeamSettingsArgs struct {
	// The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
	ReviewRequestDelegation TeamSettingsReviewRequestDelegationPtrInput
	// The GitHub team id or the GitHub team slug
	TeamId pulumi.StringInput
}

func (TeamSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamSettingsArgs)(nil)).Elem()
}

type TeamSettingsInput interface {
	pulumi.Input

	ToTeamSettingsOutput() TeamSettingsOutput
	ToTeamSettingsOutputWithContext(ctx context.Context) TeamSettingsOutput
}

func (*TeamSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamSettings)(nil)).Elem()
}

func (i *TeamSettings) ToTeamSettingsOutput() TeamSettingsOutput {
	return i.ToTeamSettingsOutputWithContext(context.Background())
}

func (i *TeamSettings) ToTeamSettingsOutputWithContext(ctx context.Context) TeamSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamSettingsOutput)
}

// TeamSettingsArrayInput is an input type that accepts TeamSettingsArray and TeamSettingsArrayOutput values.
// You can construct a concrete instance of `TeamSettingsArrayInput` via:
//
//	TeamSettingsArray{ TeamSettingsArgs{...} }
type TeamSettingsArrayInput interface {
	pulumi.Input

	ToTeamSettingsArrayOutput() TeamSettingsArrayOutput
	ToTeamSettingsArrayOutputWithContext(context.Context) TeamSettingsArrayOutput
}

type TeamSettingsArray []TeamSettingsInput

func (TeamSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamSettings)(nil)).Elem()
}

func (i TeamSettingsArray) ToTeamSettingsArrayOutput() TeamSettingsArrayOutput {
	return i.ToTeamSettingsArrayOutputWithContext(context.Background())
}

func (i TeamSettingsArray) ToTeamSettingsArrayOutputWithContext(ctx context.Context) TeamSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamSettingsArrayOutput)
}

// TeamSettingsMapInput is an input type that accepts TeamSettingsMap and TeamSettingsMapOutput values.
// You can construct a concrete instance of `TeamSettingsMapInput` via:
//
//	TeamSettingsMap{ "key": TeamSettingsArgs{...} }
type TeamSettingsMapInput interface {
	pulumi.Input

	ToTeamSettingsMapOutput() TeamSettingsMapOutput
	ToTeamSettingsMapOutputWithContext(context.Context) TeamSettingsMapOutput
}

type TeamSettingsMap map[string]TeamSettingsInput

func (TeamSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamSettings)(nil)).Elem()
}

func (i TeamSettingsMap) ToTeamSettingsMapOutput() TeamSettingsMapOutput {
	return i.ToTeamSettingsMapOutputWithContext(context.Background())
}

func (i TeamSettingsMap) ToTeamSettingsMapOutputWithContext(ctx context.Context) TeamSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamSettingsMapOutput)
}

type TeamSettingsOutput struct{ *pulumi.OutputState }

func (TeamSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamSettings)(nil)).Elem()
}

func (o TeamSettingsOutput) ToTeamSettingsOutput() TeamSettingsOutput {
	return o
}

func (o TeamSettingsOutput) ToTeamSettingsOutputWithContext(ctx context.Context) TeamSettingsOutput {
	return o
}

// The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
func (o TeamSettingsOutput) ReviewRequestDelegation() TeamSettingsReviewRequestDelegationPtrOutput {
	return o.ApplyT(func(v *TeamSettings) TeamSettingsReviewRequestDelegationPtrOutput { return v.ReviewRequestDelegation }).(TeamSettingsReviewRequestDelegationPtrOutput)
}

// The GitHub team id or the GitHub team slug
func (o TeamSettingsOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamSettings) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

// The slug of the Team within the Organization.
func (o TeamSettingsOutput) TeamSlug() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamSettings) pulumi.StringOutput { return v.TeamSlug }).(pulumi.StringOutput)
}

// The unique ID of the Team on GitHub. Corresponds to the ID of the 'github_team_settings' resource.
func (o TeamSettingsOutput) TeamUid() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamSettings) pulumi.StringOutput { return v.TeamUid }).(pulumi.StringOutput)
}

type TeamSettingsArrayOutput struct{ *pulumi.OutputState }

func (TeamSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamSettings)(nil)).Elem()
}

func (o TeamSettingsArrayOutput) ToTeamSettingsArrayOutput() TeamSettingsArrayOutput {
	return o
}

func (o TeamSettingsArrayOutput) ToTeamSettingsArrayOutputWithContext(ctx context.Context) TeamSettingsArrayOutput {
	return o
}

func (o TeamSettingsArrayOutput) Index(i pulumi.IntInput) TeamSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamSettings {
		return vs[0].([]*TeamSettings)[vs[1].(int)]
	}).(TeamSettingsOutput)
}

type TeamSettingsMapOutput struct{ *pulumi.OutputState }

func (TeamSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamSettings)(nil)).Elem()
}

func (o TeamSettingsMapOutput) ToTeamSettingsMapOutput() TeamSettingsMapOutput {
	return o
}

func (o TeamSettingsMapOutput) ToTeamSettingsMapOutputWithContext(ctx context.Context) TeamSettingsMapOutput {
	return o
}

func (o TeamSettingsMapOutput) MapIndex(k pulumi.StringInput) TeamSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamSettings {
		return vs[0].(map[string]*TeamSettings)[vs[1].(string)]
	}).(TeamSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamSettingsInput)(nil)).Elem(), &TeamSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamSettingsArrayInput)(nil)).Elem(), TeamSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamSettingsMapInput)(nil)).Elem(), TeamSettingsMap{})
	pulumi.RegisterOutputType(TeamSettingsOutput{})
	pulumi.RegisterOutputType(TeamSettingsArrayOutput{})
	pulumi.RegisterOutputType(TeamSettingsMapOutput{})
}
