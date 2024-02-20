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

// This resource allows you to create and manage settings for a GitHub Organization.
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
//			_, err := github.NewOrganizationSettings(ctx, "test", &github.OrganizationSettingsArgs{
//				AdvancedSecurityEnabledForNewRepositories: pulumi.Bool(false),
//				BillingEmail:                pulumi.String("test@example.com"),
//				Blog:                        pulumi.String("https://example.com"),
//				Company:                     pulumi.String("Test Company"),
//				DefaultRepositoryPermission: pulumi.String("read"),
//				DependabotAlertsEnabledForNewRepositories:             pulumi.Bool(false),
//				DependabotSecurityUpdatesEnabledForNewRepositories:    pulumi.Bool(false),
//				DependencyGraphEnabledForNewRepositories:              pulumi.Bool(false),
//				Description:                                           pulumi.String("Test Description"),
//				Email:                                                 pulumi.String("test@example.com"),
//				HasOrganizationProjects:                               pulumi.Bool(true),
//				HasRepositoryProjects:                                 pulumi.Bool(true),
//				Location:                                              pulumi.String("Test Location"),
//				MembersCanCreateInternalRepositories:                  pulumi.Bool(true),
//				MembersCanCreatePages:                                 pulumi.Bool(true),
//				MembersCanCreatePrivatePages:                          pulumi.Bool(true),
//				MembersCanCreatePrivateRepositories:                   pulumi.Bool(true),
//				MembersCanCreatePublicPages:                           pulumi.Bool(true),
//				MembersCanCreatePublicRepositories:                    pulumi.Bool(true),
//				MembersCanCreateRepositories:                          pulumi.Bool(true),
//				MembersCanForkPrivateRepositories:                     pulumi.Bool(true),
//				SecretScanningEnabledForNewRepositories:               pulumi.Bool(false),
//				SecretScanningPushProtectionEnabledForNewRepositories: pulumi.Bool(false),
//				TwitterUsername:                                       pulumi.String("Test"),
//				WebCommitSignoffRequired:                              pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Organization settings can be imported using the `id` of the organization. The `id` of the organization can be found using the [get an organization](https://docs.github.com/en/rest/orgs/orgs#get-an-organization) API.
//
// ```sh
//
//	$ pulumi import github:index/organizationSettings:OrganizationSettings test 123456789
//
// ```
type OrganizationSettings struct {
	pulumi.CustomResourceState

	// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
	AdvancedSecurityEnabledForNewRepositories pulumi.BoolPtrOutput `pulumi:"advancedSecurityEnabledForNewRepositories"`
	// The billing email address for the organization.
	BillingEmail pulumi.StringOutput `pulumi:"billingEmail"`
	// The blog URL for the organization.
	Blog pulumi.StringPtrOutput `pulumi:"blog"`
	// The company name for the organization.
	Company pulumi.StringPtrOutput `pulumi:"company"`
	// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
	DefaultRepositoryPermission pulumi.StringPtrOutput `pulumi:"defaultRepositoryPermission"`
	// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
	DependabotAlertsEnabledForNewRepositories pulumi.BoolPtrOutput `pulumi:"dependabotAlertsEnabledForNewRepositories"`
	// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
	DependabotSecurityUpdatesEnabledForNewRepositories pulumi.BoolPtrOutput `pulumi:"dependabotSecurityUpdatesEnabledForNewRepositories"`
	// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
	DependencyGraphEnabledForNewRepositories pulumi.BoolPtrOutput `pulumi:"dependencyGraphEnabledForNewRepositories"`
	// The description for the organization.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The email address for the organization.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// Whether or not organization projects are enabled for the organization.
	HasOrganizationProjects pulumi.BoolPtrOutput `pulumi:"hasOrganizationProjects"`
	// Whether or not repository projects are enabled for the organization.
	HasRepositoryProjects pulumi.BoolPtrOutput `pulumi:"hasRepositoryProjects"`
	// The location for the organization.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
	MembersCanCreateInternalRepositories pulumi.BoolPtrOutput `pulumi:"membersCanCreateInternalRepositories"`
	// Whether or not organization members can create new pages. Defaults to `true`.
	MembersCanCreatePages pulumi.BoolPtrOutput `pulumi:"membersCanCreatePages"`
	// Whether or not organization members can create new private pages. Defaults to `true`.
	MembersCanCreatePrivatePages pulumi.BoolPtrOutput `pulumi:"membersCanCreatePrivatePages"`
	// Whether or not organization members can create new private repositories. Defaults to `true`.
	MembersCanCreatePrivateRepositories pulumi.BoolPtrOutput `pulumi:"membersCanCreatePrivateRepositories"`
	// Whether or not organization members can create new public pages. Defaults to `true`.
	MembersCanCreatePublicPages pulumi.BoolPtrOutput `pulumi:"membersCanCreatePublicPages"`
	// Whether or not organization members can create new public repositories. Defaults to `true`.
	MembersCanCreatePublicRepositories pulumi.BoolPtrOutput `pulumi:"membersCanCreatePublicRepositories"`
	// Whether or not organization members can create new repositories. Defaults to `true`.
	MembersCanCreateRepositories pulumi.BoolPtrOutput `pulumi:"membersCanCreateRepositories"`
	// Whether or not organization members can fork private repositories. Defaults to `false`.
	MembersCanForkPrivateRepositories pulumi.BoolPtrOutput `pulumi:"membersCanForkPrivateRepositories"`
	// The name for the organization.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
	SecretScanningEnabledForNewRepositories pulumi.BoolPtrOutput `pulumi:"secretScanningEnabledForNewRepositories"`
	// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
	SecretScanningPushProtectionEnabledForNewRepositories pulumi.BoolPtrOutput `pulumi:"secretScanningPushProtectionEnabledForNewRepositories"`
	// The Twitter username for the organization.
	TwitterUsername pulumi.StringPtrOutput `pulumi:"twitterUsername"`
	// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
	WebCommitSignoffRequired pulumi.BoolPtrOutput `pulumi:"webCommitSignoffRequired"`
}

// NewOrganizationSettings registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSettings(ctx *pulumi.Context,
	name string, args *OrganizationSettingsArgs, opts ...pulumi.ResourceOption) (*OrganizationSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingEmail == nil {
		return nil, errors.New("invalid value for required argument 'BillingEmail'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationSettings
	err := ctx.RegisterResource("github:index/organizationSettings:OrganizationSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSettings gets an existing OrganizationSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSettingsState, opts ...pulumi.ResourceOption) (*OrganizationSettings, error) {
	var resource OrganizationSettings
	err := ctx.ReadResource("github:index/organizationSettings:OrganizationSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSettings resources.
type organizationSettingsState struct {
	// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
	AdvancedSecurityEnabledForNewRepositories *bool `pulumi:"advancedSecurityEnabledForNewRepositories"`
	// The billing email address for the organization.
	BillingEmail *string `pulumi:"billingEmail"`
	// The blog URL for the organization.
	Blog *string `pulumi:"blog"`
	// The company name for the organization.
	Company *string `pulumi:"company"`
	// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
	DefaultRepositoryPermission *string `pulumi:"defaultRepositoryPermission"`
	// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
	DependabotAlertsEnabledForNewRepositories *bool `pulumi:"dependabotAlertsEnabledForNewRepositories"`
	// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
	DependabotSecurityUpdatesEnabledForNewRepositories *bool `pulumi:"dependabotSecurityUpdatesEnabledForNewRepositories"`
	// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
	DependencyGraphEnabledForNewRepositories *bool `pulumi:"dependencyGraphEnabledForNewRepositories"`
	// The description for the organization.
	Description *string `pulumi:"description"`
	// The email address for the organization.
	Email *string `pulumi:"email"`
	// Whether or not organization projects are enabled for the organization.
	HasOrganizationProjects *bool `pulumi:"hasOrganizationProjects"`
	// Whether or not repository projects are enabled for the organization.
	HasRepositoryProjects *bool `pulumi:"hasRepositoryProjects"`
	// The location for the organization.
	Location *string `pulumi:"location"`
	// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
	MembersCanCreateInternalRepositories *bool `pulumi:"membersCanCreateInternalRepositories"`
	// Whether or not organization members can create new pages. Defaults to `true`.
	MembersCanCreatePages *bool `pulumi:"membersCanCreatePages"`
	// Whether or not organization members can create new private pages. Defaults to `true`.
	MembersCanCreatePrivatePages *bool `pulumi:"membersCanCreatePrivatePages"`
	// Whether or not organization members can create new private repositories. Defaults to `true`.
	MembersCanCreatePrivateRepositories *bool `pulumi:"membersCanCreatePrivateRepositories"`
	// Whether or not organization members can create new public pages. Defaults to `true`.
	MembersCanCreatePublicPages *bool `pulumi:"membersCanCreatePublicPages"`
	// Whether or not organization members can create new public repositories. Defaults to `true`.
	MembersCanCreatePublicRepositories *bool `pulumi:"membersCanCreatePublicRepositories"`
	// Whether or not organization members can create new repositories. Defaults to `true`.
	MembersCanCreateRepositories *bool `pulumi:"membersCanCreateRepositories"`
	// Whether or not organization members can fork private repositories. Defaults to `false`.
	MembersCanForkPrivateRepositories *bool `pulumi:"membersCanForkPrivateRepositories"`
	// The name for the organization.
	Name *string `pulumi:"name"`
	// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
	SecretScanningEnabledForNewRepositories *bool `pulumi:"secretScanningEnabledForNewRepositories"`
	// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
	SecretScanningPushProtectionEnabledForNewRepositories *bool `pulumi:"secretScanningPushProtectionEnabledForNewRepositories"`
	// The Twitter username for the organization.
	TwitterUsername *string `pulumi:"twitterUsername"`
	// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
	WebCommitSignoffRequired *bool `pulumi:"webCommitSignoffRequired"`
}

type OrganizationSettingsState struct {
	// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
	AdvancedSecurityEnabledForNewRepositories pulumi.BoolPtrInput
	// The billing email address for the organization.
	BillingEmail pulumi.StringPtrInput
	// The blog URL for the organization.
	Blog pulumi.StringPtrInput
	// The company name for the organization.
	Company pulumi.StringPtrInput
	// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
	DefaultRepositoryPermission pulumi.StringPtrInput
	// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
	DependabotAlertsEnabledForNewRepositories pulumi.BoolPtrInput
	// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
	DependabotSecurityUpdatesEnabledForNewRepositories pulumi.BoolPtrInput
	// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
	DependencyGraphEnabledForNewRepositories pulumi.BoolPtrInput
	// The description for the organization.
	Description pulumi.StringPtrInput
	// The email address for the organization.
	Email pulumi.StringPtrInput
	// Whether or not organization projects are enabled for the organization.
	HasOrganizationProjects pulumi.BoolPtrInput
	// Whether or not repository projects are enabled for the organization.
	HasRepositoryProjects pulumi.BoolPtrInput
	// The location for the organization.
	Location pulumi.StringPtrInput
	// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
	MembersCanCreateInternalRepositories pulumi.BoolPtrInput
	// Whether or not organization members can create new pages. Defaults to `true`.
	MembersCanCreatePages pulumi.BoolPtrInput
	// Whether or not organization members can create new private pages. Defaults to `true`.
	MembersCanCreatePrivatePages pulumi.BoolPtrInput
	// Whether or not organization members can create new private repositories. Defaults to `true`.
	MembersCanCreatePrivateRepositories pulumi.BoolPtrInput
	// Whether or not organization members can create new public pages. Defaults to `true`.
	MembersCanCreatePublicPages pulumi.BoolPtrInput
	// Whether or not organization members can create new public repositories. Defaults to `true`.
	MembersCanCreatePublicRepositories pulumi.BoolPtrInput
	// Whether or not organization members can create new repositories. Defaults to `true`.
	MembersCanCreateRepositories pulumi.BoolPtrInput
	// Whether or not organization members can fork private repositories. Defaults to `false`.
	MembersCanForkPrivateRepositories pulumi.BoolPtrInput
	// The name for the organization.
	Name pulumi.StringPtrInput
	// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
	SecretScanningEnabledForNewRepositories pulumi.BoolPtrInput
	// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
	SecretScanningPushProtectionEnabledForNewRepositories pulumi.BoolPtrInput
	// The Twitter username for the organization.
	TwitterUsername pulumi.StringPtrInput
	// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
	WebCommitSignoffRequired pulumi.BoolPtrInput
}

func (OrganizationSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSettingsState)(nil)).Elem()
}

type organizationSettingsArgs struct {
	// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
	AdvancedSecurityEnabledForNewRepositories *bool `pulumi:"advancedSecurityEnabledForNewRepositories"`
	// The billing email address for the organization.
	BillingEmail string `pulumi:"billingEmail"`
	// The blog URL for the organization.
	Blog *string `pulumi:"blog"`
	// The company name for the organization.
	Company *string `pulumi:"company"`
	// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
	DefaultRepositoryPermission *string `pulumi:"defaultRepositoryPermission"`
	// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
	DependabotAlertsEnabledForNewRepositories *bool `pulumi:"dependabotAlertsEnabledForNewRepositories"`
	// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
	DependabotSecurityUpdatesEnabledForNewRepositories *bool `pulumi:"dependabotSecurityUpdatesEnabledForNewRepositories"`
	// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
	DependencyGraphEnabledForNewRepositories *bool `pulumi:"dependencyGraphEnabledForNewRepositories"`
	// The description for the organization.
	Description *string `pulumi:"description"`
	// The email address for the organization.
	Email *string `pulumi:"email"`
	// Whether or not organization projects are enabled for the organization.
	HasOrganizationProjects *bool `pulumi:"hasOrganizationProjects"`
	// Whether or not repository projects are enabled for the organization.
	HasRepositoryProjects *bool `pulumi:"hasRepositoryProjects"`
	// The location for the organization.
	Location *string `pulumi:"location"`
	// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
	MembersCanCreateInternalRepositories *bool `pulumi:"membersCanCreateInternalRepositories"`
	// Whether or not organization members can create new pages. Defaults to `true`.
	MembersCanCreatePages *bool `pulumi:"membersCanCreatePages"`
	// Whether or not organization members can create new private pages. Defaults to `true`.
	MembersCanCreatePrivatePages *bool `pulumi:"membersCanCreatePrivatePages"`
	// Whether or not organization members can create new private repositories. Defaults to `true`.
	MembersCanCreatePrivateRepositories *bool `pulumi:"membersCanCreatePrivateRepositories"`
	// Whether or not organization members can create new public pages. Defaults to `true`.
	MembersCanCreatePublicPages *bool `pulumi:"membersCanCreatePublicPages"`
	// Whether or not organization members can create new public repositories. Defaults to `true`.
	MembersCanCreatePublicRepositories *bool `pulumi:"membersCanCreatePublicRepositories"`
	// Whether or not organization members can create new repositories. Defaults to `true`.
	MembersCanCreateRepositories *bool `pulumi:"membersCanCreateRepositories"`
	// Whether or not organization members can fork private repositories. Defaults to `false`.
	MembersCanForkPrivateRepositories *bool `pulumi:"membersCanForkPrivateRepositories"`
	// The name for the organization.
	Name *string `pulumi:"name"`
	// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
	SecretScanningEnabledForNewRepositories *bool `pulumi:"secretScanningEnabledForNewRepositories"`
	// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
	SecretScanningPushProtectionEnabledForNewRepositories *bool `pulumi:"secretScanningPushProtectionEnabledForNewRepositories"`
	// The Twitter username for the organization.
	TwitterUsername *string `pulumi:"twitterUsername"`
	// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
	WebCommitSignoffRequired *bool `pulumi:"webCommitSignoffRequired"`
}

// The set of arguments for constructing a OrganizationSettings resource.
type OrganizationSettingsArgs struct {
	// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
	AdvancedSecurityEnabledForNewRepositories pulumi.BoolPtrInput
	// The billing email address for the organization.
	BillingEmail pulumi.StringInput
	// The blog URL for the organization.
	Blog pulumi.StringPtrInput
	// The company name for the organization.
	Company pulumi.StringPtrInput
	// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
	DefaultRepositoryPermission pulumi.StringPtrInput
	// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
	DependabotAlertsEnabledForNewRepositories pulumi.BoolPtrInput
	// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
	DependabotSecurityUpdatesEnabledForNewRepositories pulumi.BoolPtrInput
	// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
	DependencyGraphEnabledForNewRepositories pulumi.BoolPtrInput
	// The description for the organization.
	Description pulumi.StringPtrInput
	// The email address for the organization.
	Email pulumi.StringPtrInput
	// Whether or not organization projects are enabled for the organization.
	HasOrganizationProjects pulumi.BoolPtrInput
	// Whether or not repository projects are enabled for the organization.
	HasRepositoryProjects pulumi.BoolPtrInput
	// The location for the organization.
	Location pulumi.StringPtrInput
	// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
	MembersCanCreateInternalRepositories pulumi.BoolPtrInput
	// Whether or not organization members can create new pages. Defaults to `true`.
	MembersCanCreatePages pulumi.BoolPtrInput
	// Whether or not organization members can create new private pages. Defaults to `true`.
	MembersCanCreatePrivatePages pulumi.BoolPtrInput
	// Whether or not organization members can create new private repositories. Defaults to `true`.
	MembersCanCreatePrivateRepositories pulumi.BoolPtrInput
	// Whether or not organization members can create new public pages. Defaults to `true`.
	MembersCanCreatePublicPages pulumi.BoolPtrInput
	// Whether or not organization members can create new public repositories. Defaults to `true`.
	MembersCanCreatePublicRepositories pulumi.BoolPtrInput
	// Whether or not organization members can create new repositories. Defaults to `true`.
	MembersCanCreateRepositories pulumi.BoolPtrInput
	// Whether or not organization members can fork private repositories. Defaults to `false`.
	MembersCanForkPrivateRepositories pulumi.BoolPtrInput
	// The name for the organization.
	Name pulumi.StringPtrInput
	// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
	SecretScanningEnabledForNewRepositories pulumi.BoolPtrInput
	// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
	SecretScanningPushProtectionEnabledForNewRepositories pulumi.BoolPtrInput
	// The Twitter username for the organization.
	TwitterUsername pulumi.StringPtrInput
	// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
	WebCommitSignoffRequired pulumi.BoolPtrInput
}

func (OrganizationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSettingsArgs)(nil)).Elem()
}

type OrganizationSettingsInput interface {
	pulumi.Input

	ToOrganizationSettingsOutput() OrganizationSettingsOutput
	ToOrganizationSettingsOutputWithContext(ctx context.Context) OrganizationSettingsOutput
}

func (*OrganizationSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSettings)(nil)).Elem()
}

func (i *OrganizationSettings) ToOrganizationSettingsOutput() OrganizationSettingsOutput {
	return i.ToOrganizationSettingsOutputWithContext(context.Background())
}

func (i *OrganizationSettings) ToOrganizationSettingsOutputWithContext(ctx context.Context) OrganizationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSettingsOutput)
}

// OrganizationSettingsArrayInput is an input type that accepts OrganizationSettingsArray and OrganizationSettingsArrayOutput values.
// You can construct a concrete instance of `OrganizationSettingsArrayInput` via:
//
//	OrganizationSettingsArray{ OrganizationSettingsArgs{...} }
type OrganizationSettingsArrayInput interface {
	pulumi.Input

	ToOrganizationSettingsArrayOutput() OrganizationSettingsArrayOutput
	ToOrganizationSettingsArrayOutputWithContext(context.Context) OrganizationSettingsArrayOutput
}

type OrganizationSettingsArray []OrganizationSettingsInput

func (OrganizationSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationSettings)(nil)).Elem()
}

func (i OrganizationSettingsArray) ToOrganizationSettingsArrayOutput() OrganizationSettingsArrayOutput {
	return i.ToOrganizationSettingsArrayOutputWithContext(context.Background())
}

func (i OrganizationSettingsArray) ToOrganizationSettingsArrayOutputWithContext(ctx context.Context) OrganizationSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSettingsArrayOutput)
}

// OrganizationSettingsMapInput is an input type that accepts OrganizationSettingsMap and OrganizationSettingsMapOutput values.
// You can construct a concrete instance of `OrganizationSettingsMapInput` via:
//
//	OrganizationSettingsMap{ "key": OrganizationSettingsArgs{...} }
type OrganizationSettingsMapInput interface {
	pulumi.Input

	ToOrganizationSettingsMapOutput() OrganizationSettingsMapOutput
	ToOrganizationSettingsMapOutputWithContext(context.Context) OrganizationSettingsMapOutput
}

type OrganizationSettingsMap map[string]OrganizationSettingsInput

func (OrganizationSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationSettings)(nil)).Elem()
}

func (i OrganizationSettingsMap) ToOrganizationSettingsMapOutput() OrganizationSettingsMapOutput {
	return i.ToOrganizationSettingsMapOutputWithContext(context.Background())
}

func (i OrganizationSettingsMap) ToOrganizationSettingsMapOutputWithContext(ctx context.Context) OrganizationSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSettingsMapOutput)
}

type OrganizationSettingsOutput struct{ *pulumi.OutputState }

func (OrganizationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSettings)(nil)).Elem()
}

func (o OrganizationSettingsOutput) ToOrganizationSettingsOutput() OrganizationSettingsOutput {
	return o
}

func (o OrganizationSettingsOutput) ToOrganizationSettingsOutputWithContext(ctx context.Context) OrganizationSettingsOutput {
	return o
}

// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
func (o OrganizationSettingsOutput) AdvancedSecurityEnabledForNewRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.AdvancedSecurityEnabledForNewRepositories }).(pulumi.BoolPtrOutput)
}

// The billing email address for the organization.
func (o OrganizationSettingsOutput) BillingEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringOutput { return v.BillingEmail }).(pulumi.StringOutput)
}

// The blog URL for the organization.
func (o OrganizationSettingsOutput) Blog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringPtrOutput { return v.Blog }).(pulumi.StringPtrOutput)
}

// The company name for the organization.
func (o OrganizationSettingsOutput) Company() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringPtrOutput { return v.Company }).(pulumi.StringPtrOutput)
}

// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
func (o OrganizationSettingsOutput) DefaultRepositoryPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringPtrOutput { return v.DefaultRepositoryPermission }).(pulumi.StringPtrOutput)
}

// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
func (o OrganizationSettingsOutput) DependabotAlertsEnabledForNewRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.DependabotAlertsEnabledForNewRepositories }).(pulumi.BoolPtrOutput)
}

// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
func (o OrganizationSettingsOutput) DependabotSecurityUpdatesEnabledForNewRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput {
		return v.DependabotSecurityUpdatesEnabledForNewRepositories
	}).(pulumi.BoolPtrOutput)
}

// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
func (o OrganizationSettingsOutput) DependencyGraphEnabledForNewRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.DependencyGraphEnabledForNewRepositories }).(pulumi.BoolPtrOutput)
}

// The description for the organization.
func (o OrganizationSettingsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The email address for the organization.
func (o OrganizationSettingsOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// Whether or not organization projects are enabled for the organization.
func (o OrganizationSettingsOutput) HasOrganizationProjects() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.HasOrganizationProjects }).(pulumi.BoolPtrOutput)
}

// Whether or not repository projects are enabled for the organization.
func (o OrganizationSettingsOutput) HasRepositoryProjects() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.HasRepositoryProjects }).(pulumi.BoolPtrOutput)
}

// The location for the organization.
func (o OrganizationSettingsOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
func (o OrganizationSettingsOutput) MembersCanCreateInternalRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanCreateInternalRepositories }).(pulumi.BoolPtrOutput)
}

// Whether or not organization members can create new pages. Defaults to `true`.
func (o OrganizationSettingsOutput) MembersCanCreatePages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanCreatePages }).(pulumi.BoolPtrOutput)
}

// Whether or not organization members can create new private pages. Defaults to `true`.
func (o OrganizationSettingsOutput) MembersCanCreatePrivatePages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanCreatePrivatePages }).(pulumi.BoolPtrOutput)
}

// Whether or not organization members can create new private repositories. Defaults to `true`.
func (o OrganizationSettingsOutput) MembersCanCreatePrivateRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanCreatePrivateRepositories }).(pulumi.BoolPtrOutput)
}

// Whether or not organization members can create new public pages. Defaults to `true`.
func (o OrganizationSettingsOutput) MembersCanCreatePublicPages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanCreatePublicPages }).(pulumi.BoolPtrOutput)
}

// Whether or not organization members can create new public repositories. Defaults to `true`.
func (o OrganizationSettingsOutput) MembersCanCreatePublicRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanCreatePublicRepositories }).(pulumi.BoolPtrOutput)
}

// Whether or not organization members can create new repositories. Defaults to `true`.
func (o OrganizationSettingsOutput) MembersCanCreateRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanCreateRepositories }).(pulumi.BoolPtrOutput)
}

// Whether or not organization members can fork private repositories. Defaults to `false`.
func (o OrganizationSettingsOutput) MembersCanForkPrivateRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.MembersCanForkPrivateRepositories }).(pulumi.BoolPtrOutput)
}

// The name for the organization.
func (o OrganizationSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
func (o OrganizationSettingsOutput) SecretScanningEnabledForNewRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.SecretScanningEnabledForNewRepositories }).(pulumi.BoolPtrOutput)
}

// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
func (o OrganizationSettingsOutput) SecretScanningPushProtectionEnabledForNewRepositories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput {
		return v.SecretScanningPushProtectionEnabledForNewRepositories
	}).(pulumi.BoolPtrOutput)
}

// The Twitter username for the organization.
func (o OrganizationSettingsOutput) TwitterUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.StringPtrOutput { return v.TwitterUsername }).(pulumi.StringPtrOutput)
}

// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
func (o OrganizationSettingsOutput) WebCommitSignoffRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationSettings) pulumi.BoolPtrOutput { return v.WebCommitSignoffRequired }).(pulumi.BoolPtrOutput)
}

type OrganizationSettingsArrayOutput struct{ *pulumi.OutputState }

func (OrganizationSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationSettings)(nil)).Elem()
}

func (o OrganizationSettingsArrayOutput) ToOrganizationSettingsArrayOutput() OrganizationSettingsArrayOutput {
	return o
}

func (o OrganizationSettingsArrayOutput) ToOrganizationSettingsArrayOutputWithContext(ctx context.Context) OrganizationSettingsArrayOutput {
	return o
}

func (o OrganizationSettingsArrayOutput) Index(i pulumi.IntInput) OrganizationSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationSettings {
		return vs[0].([]*OrganizationSettings)[vs[1].(int)]
	}).(OrganizationSettingsOutput)
}

type OrganizationSettingsMapOutput struct{ *pulumi.OutputState }

func (OrganizationSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationSettings)(nil)).Elem()
}

func (o OrganizationSettingsMapOutput) ToOrganizationSettingsMapOutput() OrganizationSettingsMapOutput {
	return o
}

func (o OrganizationSettingsMapOutput) ToOrganizationSettingsMapOutputWithContext(ctx context.Context) OrganizationSettingsMapOutput {
	return o
}

func (o OrganizationSettingsMapOutput) MapIndex(k pulumi.StringInput) OrganizationSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationSettings {
		return vs[0].(map[string]*OrganizationSettings)[vs[1].(string)]
	}).(OrganizationSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationSettingsInput)(nil)).Elem(), &OrganizationSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationSettingsArrayInput)(nil)).Elem(), OrganizationSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationSettingsMapInput)(nil)).Elem(), OrganizationSettingsMap{})
	pulumi.RegisterOutputType(OrganizationSettingsOutput{})
	pulumi.RegisterOutputType(OrganizationSettingsArrayOutput{})
	pulumi.RegisterOutputType(OrganizationSettingsMapOutput{})
}
