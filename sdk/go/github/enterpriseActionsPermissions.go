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

// This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise.
// You must have admin access to an enterprise to use this resource.
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
//			example_enterprise, err := github.GetEnterprise(ctx, &github.GetEnterpriseArgs{
//				Slug: "my-enterprise",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			example_org, err := github.GetOrganization(ctx, &github.GetOrganizationArgs{
//				Name: "my-org",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewEnterpriseActionsPermissions(ctx, "test", &github.EnterpriseActionsPermissionsArgs{
//				EnterpriseId:         pulumi.String(example_enterprise.Slug),
//				AllowedActions:       pulumi.String("selected"),
//				EnabledOrganizations: pulumi.String("selected"),
//				AllowedActionsConfig: &github.EnterpriseActionsPermissionsAllowedActionsConfigArgs{
//					GithubOwnedAllowed: pulumi.Bool(true),
//					PatternsAlloweds: pulumi.StringArray{
//						pulumi.String("actions/cache@*"),
//						pulumi.String("actions/checkout@*"),
//					},
//					VerifiedAllowed: pulumi.Bool(true),
//				},
//				EnabledOrganizationsConfig: &github.EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs{
//					OrganizationIds: pulumi.IntArray{
//						pulumi.String(example_org.Id),
//					},
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
//
// ## Import
//
// This resource can be imported using the name of the GitHub enterprise:
//
// ```sh
// $ pulumi import github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions test github_enterprise_name
// ```
type EnterpriseActionsPermissions struct {
	pulumi.CustomResourceState

	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions pulumi.StringPtrOutput `pulumi:"allowedActions"`
	// Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig EnterpriseActionsPermissionsAllowedActionsConfigPtrOutput `pulumi:"allowedActionsConfig"`
	// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledOrganizations pulumi.StringOutput `pulumi:"enabledOrganizations"`
	// Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
	EnabledOrganizationsConfig EnterpriseActionsPermissionsEnabledOrganizationsConfigPtrOutput `pulumi:"enabledOrganizationsConfig"`
	// The ID of the enterprise.
	EnterpriseId pulumi.StringOutput `pulumi:"enterpriseId"`
}

// NewEnterpriseActionsPermissions registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseActionsPermissions(ctx *pulumi.Context,
	name string, args *EnterpriseActionsPermissionsArgs, opts ...pulumi.ResourceOption) (*EnterpriseActionsPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnabledOrganizations == nil {
		return nil, errors.New("invalid value for required argument 'EnabledOrganizations'")
	}
	if args.EnterpriseId == nil {
		return nil, errors.New("invalid value for required argument 'EnterpriseId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnterpriseActionsPermissions
	err := ctx.RegisterResource("github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseActionsPermissions gets an existing EnterpriseActionsPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseActionsPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseActionsPermissionsState, opts ...pulumi.ResourceOption) (*EnterpriseActionsPermissions, error) {
	var resource EnterpriseActionsPermissions
	err := ctx.ReadResource("github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseActionsPermissions resources.
type enterpriseActionsPermissionsState struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions *string `pulumi:"allowedActions"`
	// Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig *EnterpriseActionsPermissionsAllowedActionsConfig `pulumi:"allowedActionsConfig"`
	// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledOrganizations *string `pulumi:"enabledOrganizations"`
	// Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
	EnabledOrganizationsConfig *EnterpriseActionsPermissionsEnabledOrganizationsConfig `pulumi:"enabledOrganizationsConfig"`
	// The ID of the enterprise.
	EnterpriseId *string `pulumi:"enterpriseId"`
}

type EnterpriseActionsPermissionsState struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions pulumi.StringPtrInput
	// Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig EnterpriseActionsPermissionsAllowedActionsConfigPtrInput
	// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledOrganizations pulumi.StringPtrInput
	// Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
	EnabledOrganizationsConfig EnterpriseActionsPermissionsEnabledOrganizationsConfigPtrInput
	// The ID of the enterprise.
	EnterpriseId pulumi.StringPtrInput
}

func (EnterpriseActionsPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseActionsPermissionsState)(nil)).Elem()
}

type enterpriseActionsPermissionsArgs struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions *string `pulumi:"allowedActions"`
	// Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig *EnterpriseActionsPermissionsAllowedActionsConfig `pulumi:"allowedActionsConfig"`
	// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledOrganizations string `pulumi:"enabledOrganizations"`
	// Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
	EnabledOrganizationsConfig *EnterpriseActionsPermissionsEnabledOrganizationsConfig `pulumi:"enabledOrganizationsConfig"`
	// The ID of the enterprise.
	EnterpriseId string `pulumi:"enterpriseId"`
}

// The set of arguments for constructing a EnterpriseActionsPermissions resource.
type EnterpriseActionsPermissionsArgs struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions pulumi.StringPtrInput
	// Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig EnterpriseActionsPermissionsAllowedActionsConfigPtrInput
	// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledOrganizations pulumi.StringInput
	// Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
	EnabledOrganizationsConfig EnterpriseActionsPermissionsEnabledOrganizationsConfigPtrInput
	// The ID of the enterprise.
	EnterpriseId pulumi.StringInput
}

func (EnterpriseActionsPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseActionsPermissionsArgs)(nil)).Elem()
}

type EnterpriseActionsPermissionsInput interface {
	pulumi.Input

	ToEnterpriseActionsPermissionsOutput() EnterpriseActionsPermissionsOutput
	ToEnterpriseActionsPermissionsOutputWithContext(ctx context.Context) EnterpriseActionsPermissionsOutput
}

func (*EnterpriseActionsPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseActionsPermissions)(nil)).Elem()
}

func (i *EnterpriseActionsPermissions) ToEnterpriseActionsPermissionsOutput() EnterpriseActionsPermissionsOutput {
	return i.ToEnterpriseActionsPermissionsOutputWithContext(context.Background())
}

func (i *EnterpriseActionsPermissions) ToEnterpriseActionsPermissionsOutputWithContext(ctx context.Context) EnterpriseActionsPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseActionsPermissionsOutput)
}

// EnterpriseActionsPermissionsArrayInput is an input type that accepts EnterpriseActionsPermissionsArray and EnterpriseActionsPermissionsArrayOutput values.
// You can construct a concrete instance of `EnterpriseActionsPermissionsArrayInput` via:
//
//	EnterpriseActionsPermissionsArray{ EnterpriseActionsPermissionsArgs{...} }
type EnterpriseActionsPermissionsArrayInput interface {
	pulumi.Input

	ToEnterpriseActionsPermissionsArrayOutput() EnterpriseActionsPermissionsArrayOutput
	ToEnterpriseActionsPermissionsArrayOutputWithContext(context.Context) EnterpriseActionsPermissionsArrayOutput
}

type EnterpriseActionsPermissionsArray []EnterpriseActionsPermissionsInput

func (EnterpriseActionsPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseActionsPermissions)(nil)).Elem()
}

func (i EnterpriseActionsPermissionsArray) ToEnterpriseActionsPermissionsArrayOutput() EnterpriseActionsPermissionsArrayOutput {
	return i.ToEnterpriseActionsPermissionsArrayOutputWithContext(context.Background())
}

func (i EnterpriseActionsPermissionsArray) ToEnterpriseActionsPermissionsArrayOutputWithContext(ctx context.Context) EnterpriseActionsPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseActionsPermissionsArrayOutput)
}

// EnterpriseActionsPermissionsMapInput is an input type that accepts EnterpriseActionsPermissionsMap and EnterpriseActionsPermissionsMapOutput values.
// You can construct a concrete instance of `EnterpriseActionsPermissionsMapInput` via:
//
//	EnterpriseActionsPermissionsMap{ "key": EnterpriseActionsPermissionsArgs{...} }
type EnterpriseActionsPermissionsMapInput interface {
	pulumi.Input

	ToEnterpriseActionsPermissionsMapOutput() EnterpriseActionsPermissionsMapOutput
	ToEnterpriseActionsPermissionsMapOutputWithContext(context.Context) EnterpriseActionsPermissionsMapOutput
}

type EnterpriseActionsPermissionsMap map[string]EnterpriseActionsPermissionsInput

func (EnterpriseActionsPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseActionsPermissions)(nil)).Elem()
}

func (i EnterpriseActionsPermissionsMap) ToEnterpriseActionsPermissionsMapOutput() EnterpriseActionsPermissionsMapOutput {
	return i.ToEnterpriseActionsPermissionsMapOutputWithContext(context.Background())
}

func (i EnterpriseActionsPermissionsMap) ToEnterpriseActionsPermissionsMapOutputWithContext(ctx context.Context) EnterpriseActionsPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseActionsPermissionsMapOutput)
}

type EnterpriseActionsPermissionsOutput struct{ *pulumi.OutputState }

func (EnterpriseActionsPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseActionsPermissions)(nil)).Elem()
}

func (o EnterpriseActionsPermissionsOutput) ToEnterpriseActionsPermissionsOutput() EnterpriseActionsPermissionsOutput {
	return o
}

func (o EnterpriseActionsPermissionsOutput) ToEnterpriseActionsPermissionsOutputWithContext(ctx context.Context) EnterpriseActionsPermissionsOutput {
	return o
}

// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
func (o EnterpriseActionsPermissionsOutput) AllowedActions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseActionsPermissions) pulumi.StringPtrOutput { return v.AllowedActions }).(pulumi.StringPtrOutput)
}

// Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
func (o EnterpriseActionsPermissionsOutput) AllowedActionsConfig() EnterpriseActionsPermissionsAllowedActionsConfigPtrOutput {
	return o.ApplyT(func(v *EnterpriseActionsPermissions) EnterpriseActionsPermissionsAllowedActionsConfigPtrOutput {
		return v.AllowedActionsConfig
	}).(EnterpriseActionsPermissionsAllowedActionsConfigPtrOutput)
}

// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
func (o EnterpriseActionsPermissionsOutput) EnabledOrganizations() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseActionsPermissions) pulumi.StringOutput { return v.EnabledOrganizations }).(pulumi.StringOutput)
}

// Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
func (o EnterpriseActionsPermissionsOutput) EnabledOrganizationsConfig() EnterpriseActionsPermissionsEnabledOrganizationsConfigPtrOutput {
	return o.ApplyT(func(v *EnterpriseActionsPermissions) EnterpriseActionsPermissionsEnabledOrganizationsConfigPtrOutput {
		return v.EnabledOrganizationsConfig
	}).(EnterpriseActionsPermissionsEnabledOrganizationsConfigPtrOutput)
}

// The ID of the enterprise.
func (o EnterpriseActionsPermissionsOutput) EnterpriseId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseActionsPermissions) pulumi.StringOutput { return v.EnterpriseId }).(pulumi.StringOutput)
}

type EnterpriseActionsPermissionsArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseActionsPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseActionsPermissions)(nil)).Elem()
}

func (o EnterpriseActionsPermissionsArrayOutput) ToEnterpriseActionsPermissionsArrayOutput() EnterpriseActionsPermissionsArrayOutput {
	return o
}

func (o EnterpriseActionsPermissionsArrayOutput) ToEnterpriseActionsPermissionsArrayOutputWithContext(ctx context.Context) EnterpriseActionsPermissionsArrayOutput {
	return o
}

func (o EnterpriseActionsPermissionsArrayOutput) Index(i pulumi.IntInput) EnterpriseActionsPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseActionsPermissions {
		return vs[0].([]*EnterpriseActionsPermissions)[vs[1].(int)]
	}).(EnterpriseActionsPermissionsOutput)
}

type EnterpriseActionsPermissionsMapOutput struct{ *pulumi.OutputState }

func (EnterpriseActionsPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseActionsPermissions)(nil)).Elem()
}

func (o EnterpriseActionsPermissionsMapOutput) ToEnterpriseActionsPermissionsMapOutput() EnterpriseActionsPermissionsMapOutput {
	return o
}

func (o EnterpriseActionsPermissionsMapOutput) ToEnterpriseActionsPermissionsMapOutputWithContext(ctx context.Context) EnterpriseActionsPermissionsMapOutput {
	return o
}

func (o EnterpriseActionsPermissionsMapOutput) MapIndex(k pulumi.StringInput) EnterpriseActionsPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseActionsPermissions {
		return vs[0].(map[string]*EnterpriseActionsPermissions)[vs[1].(string)]
	}).(EnterpriseActionsPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseActionsPermissionsInput)(nil)).Elem(), &EnterpriseActionsPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseActionsPermissionsArrayInput)(nil)).Elem(), EnterpriseActionsPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseActionsPermissionsMapInput)(nil)).Elem(), EnterpriseActionsPermissionsMap{})
	pulumi.RegisterOutputType(EnterpriseActionsPermissionsOutput{})
	pulumi.RegisterOutputType(EnterpriseActionsPermissionsArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseActionsPermissionsMapOutput{})
}
