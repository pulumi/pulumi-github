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

// This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise organizations.
// You must have admin access to an organization to use this resource.
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
//			example, err := github.NewRepository(ctx, "example", &github.RepositoryArgs{
//				Name: pulumi.String("my-repository"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewActionsOrganizationPermissions(ctx, "test", &github.ActionsOrganizationPermissionsArgs{
//				AllowedActions:      pulumi.String("selected"),
//				EnabledRepositories: pulumi.String("selected"),
//				AllowedActionsConfig: &github.ActionsOrganizationPermissionsAllowedActionsConfigArgs{
//					GithubOwnedAllowed: pulumi.Bool(true),
//					PatternsAlloweds: pulumi.StringArray{
//						pulumi.String("actions/cache@*"),
//						pulumi.String("actions/checkout@*"),
//					},
//					VerifiedAllowed: pulumi.Bool(true),
//				},
//				EnabledRepositoriesConfig: &github.ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs{
//					RepositoryIds: pulumi.IntArray{
//						example.RepoId,
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// This resource can be imported using the name of the GitHub organization:
//
// ```sh
// $ pulumi import github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions test github_organization_name
// ```
type ActionsOrganizationPermissions struct {
	pulumi.CustomResourceState

	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions pulumi.StringPtrOutput `pulumi:"allowedActions"`
	// Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig ActionsOrganizationPermissionsAllowedActionsConfigPtrOutput `pulumi:"allowedActionsConfig"`
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledRepositories pulumi.StringOutput `pulumi:"enabledRepositories"`
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
	EnabledRepositoriesConfig ActionsOrganizationPermissionsEnabledRepositoriesConfigPtrOutput `pulumi:"enabledRepositoriesConfig"`
}

// NewActionsOrganizationPermissions registers a new resource with the given unique name, arguments, and options.
func NewActionsOrganizationPermissions(ctx *pulumi.Context,
	name string, args *ActionsOrganizationPermissionsArgs, opts ...pulumi.ResourceOption) (*ActionsOrganizationPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnabledRepositories == nil {
		return nil, errors.New("invalid value for required argument 'EnabledRepositories'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ActionsOrganizationPermissions
	err := ctx.RegisterResource("github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsOrganizationPermissions gets an existing ActionsOrganizationPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsOrganizationPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsOrganizationPermissionsState, opts ...pulumi.ResourceOption) (*ActionsOrganizationPermissions, error) {
	var resource ActionsOrganizationPermissions
	err := ctx.ReadResource("github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsOrganizationPermissions resources.
type actionsOrganizationPermissionsState struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions *string `pulumi:"allowedActions"`
	// Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig *ActionsOrganizationPermissionsAllowedActionsConfig `pulumi:"allowedActionsConfig"`
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledRepositories *string `pulumi:"enabledRepositories"`
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
	EnabledRepositoriesConfig *ActionsOrganizationPermissionsEnabledRepositoriesConfig `pulumi:"enabledRepositoriesConfig"`
}

type ActionsOrganizationPermissionsState struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions pulumi.StringPtrInput
	// Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig ActionsOrganizationPermissionsAllowedActionsConfigPtrInput
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledRepositories pulumi.StringPtrInput
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
	EnabledRepositoriesConfig ActionsOrganizationPermissionsEnabledRepositoriesConfigPtrInput
}

func (ActionsOrganizationPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationPermissionsState)(nil)).Elem()
}

type actionsOrganizationPermissionsArgs struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions *string `pulumi:"allowedActions"`
	// Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig *ActionsOrganizationPermissionsAllowedActionsConfig `pulumi:"allowedActionsConfig"`
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledRepositories string `pulumi:"enabledRepositories"`
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
	EnabledRepositoriesConfig *ActionsOrganizationPermissionsEnabledRepositoriesConfig `pulumi:"enabledRepositoriesConfig"`
}

// The set of arguments for constructing a ActionsOrganizationPermissions resource.
type ActionsOrganizationPermissionsArgs struct {
	// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
	AllowedActions pulumi.StringPtrInput
	// Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
	AllowedActionsConfig ActionsOrganizationPermissionsAllowedActionsConfigPtrInput
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
	EnabledRepositories pulumi.StringInput
	// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
	EnabledRepositoriesConfig ActionsOrganizationPermissionsEnabledRepositoriesConfigPtrInput
}

func (ActionsOrganizationPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationPermissionsArgs)(nil)).Elem()
}

type ActionsOrganizationPermissionsInput interface {
	pulumi.Input

	ToActionsOrganizationPermissionsOutput() ActionsOrganizationPermissionsOutput
	ToActionsOrganizationPermissionsOutputWithContext(ctx context.Context) ActionsOrganizationPermissionsOutput
}

func (*ActionsOrganizationPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationPermissions)(nil)).Elem()
}

func (i *ActionsOrganizationPermissions) ToActionsOrganizationPermissionsOutput() ActionsOrganizationPermissionsOutput {
	return i.ToActionsOrganizationPermissionsOutputWithContext(context.Background())
}

func (i *ActionsOrganizationPermissions) ToActionsOrganizationPermissionsOutputWithContext(ctx context.Context) ActionsOrganizationPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationPermissionsOutput)
}

// ActionsOrganizationPermissionsArrayInput is an input type that accepts ActionsOrganizationPermissionsArray and ActionsOrganizationPermissionsArrayOutput values.
// You can construct a concrete instance of `ActionsOrganizationPermissionsArrayInput` via:
//
//	ActionsOrganizationPermissionsArray{ ActionsOrganizationPermissionsArgs{...} }
type ActionsOrganizationPermissionsArrayInput interface {
	pulumi.Input

	ToActionsOrganizationPermissionsArrayOutput() ActionsOrganizationPermissionsArrayOutput
	ToActionsOrganizationPermissionsArrayOutputWithContext(context.Context) ActionsOrganizationPermissionsArrayOutput
}

type ActionsOrganizationPermissionsArray []ActionsOrganizationPermissionsInput

func (ActionsOrganizationPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationPermissions)(nil)).Elem()
}

func (i ActionsOrganizationPermissionsArray) ToActionsOrganizationPermissionsArrayOutput() ActionsOrganizationPermissionsArrayOutput {
	return i.ToActionsOrganizationPermissionsArrayOutputWithContext(context.Background())
}

func (i ActionsOrganizationPermissionsArray) ToActionsOrganizationPermissionsArrayOutputWithContext(ctx context.Context) ActionsOrganizationPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationPermissionsArrayOutput)
}

// ActionsOrganizationPermissionsMapInput is an input type that accepts ActionsOrganizationPermissionsMap and ActionsOrganizationPermissionsMapOutput values.
// You can construct a concrete instance of `ActionsOrganizationPermissionsMapInput` via:
//
//	ActionsOrganizationPermissionsMap{ "key": ActionsOrganizationPermissionsArgs{...} }
type ActionsOrganizationPermissionsMapInput interface {
	pulumi.Input

	ToActionsOrganizationPermissionsMapOutput() ActionsOrganizationPermissionsMapOutput
	ToActionsOrganizationPermissionsMapOutputWithContext(context.Context) ActionsOrganizationPermissionsMapOutput
}

type ActionsOrganizationPermissionsMap map[string]ActionsOrganizationPermissionsInput

func (ActionsOrganizationPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationPermissions)(nil)).Elem()
}

func (i ActionsOrganizationPermissionsMap) ToActionsOrganizationPermissionsMapOutput() ActionsOrganizationPermissionsMapOutput {
	return i.ToActionsOrganizationPermissionsMapOutputWithContext(context.Background())
}

func (i ActionsOrganizationPermissionsMap) ToActionsOrganizationPermissionsMapOutputWithContext(ctx context.Context) ActionsOrganizationPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationPermissionsMapOutput)
}

type ActionsOrganizationPermissionsOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationPermissions)(nil)).Elem()
}

func (o ActionsOrganizationPermissionsOutput) ToActionsOrganizationPermissionsOutput() ActionsOrganizationPermissionsOutput {
	return o
}

func (o ActionsOrganizationPermissionsOutput) ToActionsOrganizationPermissionsOutputWithContext(ctx context.Context) ActionsOrganizationPermissionsOutput {
	return o
}

// The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
func (o ActionsOrganizationPermissionsOutput) AllowedActions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionsOrganizationPermissions) pulumi.StringPtrOutput { return v.AllowedActions }).(pulumi.StringPtrOutput)
}

// Sets the actions that are allowed in an organization. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
func (o ActionsOrganizationPermissionsOutput) AllowedActionsConfig() ActionsOrganizationPermissionsAllowedActionsConfigPtrOutput {
	return o.ApplyT(func(v *ActionsOrganizationPermissions) ActionsOrganizationPermissionsAllowedActionsConfigPtrOutput {
		return v.AllowedActionsConfig
	}).(ActionsOrganizationPermissionsAllowedActionsConfigPtrOutput)
}

// The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
func (o ActionsOrganizationPermissionsOutput) EnabledRepositories() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationPermissions) pulumi.StringOutput { return v.EnabledRepositories }).(pulumi.StringOutput)
}

// Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabledRepositories` = `selected`. See Enabled Repositories Config below for details.
func (o ActionsOrganizationPermissionsOutput) EnabledRepositoriesConfig() ActionsOrganizationPermissionsEnabledRepositoriesConfigPtrOutput {
	return o.ApplyT(func(v *ActionsOrganizationPermissions) ActionsOrganizationPermissionsEnabledRepositoriesConfigPtrOutput {
		return v.EnabledRepositoriesConfig
	}).(ActionsOrganizationPermissionsEnabledRepositoriesConfigPtrOutput)
}

type ActionsOrganizationPermissionsArrayOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationPermissions)(nil)).Elem()
}

func (o ActionsOrganizationPermissionsArrayOutput) ToActionsOrganizationPermissionsArrayOutput() ActionsOrganizationPermissionsArrayOutput {
	return o
}

func (o ActionsOrganizationPermissionsArrayOutput) ToActionsOrganizationPermissionsArrayOutputWithContext(ctx context.Context) ActionsOrganizationPermissionsArrayOutput {
	return o
}

func (o ActionsOrganizationPermissionsArrayOutput) Index(i pulumi.IntInput) ActionsOrganizationPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsOrganizationPermissions {
		return vs[0].([]*ActionsOrganizationPermissions)[vs[1].(int)]
	}).(ActionsOrganizationPermissionsOutput)
}

type ActionsOrganizationPermissionsMapOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationPermissions)(nil)).Elem()
}

func (o ActionsOrganizationPermissionsMapOutput) ToActionsOrganizationPermissionsMapOutput() ActionsOrganizationPermissionsMapOutput {
	return o
}

func (o ActionsOrganizationPermissionsMapOutput) ToActionsOrganizationPermissionsMapOutputWithContext(ctx context.Context) ActionsOrganizationPermissionsMapOutput {
	return o
}

func (o ActionsOrganizationPermissionsMapOutput) MapIndex(k pulumi.StringInput) ActionsOrganizationPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsOrganizationPermissions {
		return vs[0].(map[string]*ActionsOrganizationPermissions)[vs[1].(string)]
	}).(ActionsOrganizationPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationPermissionsInput)(nil)).Elem(), &ActionsOrganizationPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationPermissionsArrayInput)(nil)).Elem(), ActionsOrganizationPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationPermissionsMapInput)(nil)).Elem(), ActionsOrganizationPermissionsMap{})
	pulumi.RegisterOutputType(ActionsOrganizationPermissionsOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationPermissionsArrayOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationPermissionsMapOutput{})
}
