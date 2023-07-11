// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage custom roles in a GitHub Organization for use in repositories.
//
// > Note: Custom roles are currently only available in GitHub Enterprise Cloud.
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
//			_, err := github.NewOrganizationCustomRole(ctx, "example", &github.OrganizationCustomRoleArgs{
//				BaseRole:    pulumi.String("read"),
//				Description: pulumi.String("Example custom role that uses the read role as its base"),
//				Permissions: pulumi.StringArray{
//					pulumi.String("add_assignee"),
//					pulumi.String("add_label"),
//					pulumi.String("bypass_branch_protection"),
//					pulumi.String("close_issue"),
//					pulumi.String("close_pull_request"),
//					pulumi.String("mark_as_duplicate"),
//					pulumi.String("create_tag"),
//					pulumi.String("delete_issue"),
//					pulumi.String("delete_tag"),
//					pulumi.String("manage_deploy_keys"),
//					pulumi.String("push_protected_branch"),
//					pulumi.String("read_code_scanning"),
//					pulumi.String("reopen_issue"),
//					pulumi.String("reopen_pull_request"),
//					pulumi.String("request_pr_review"),
//					pulumi.String("resolve_dependabot_alerts"),
//					pulumi.String("resolve_secret_scanning_alerts"),
//					pulumi.String("view_secret_scanning_alerts"),
//					pulumi.String("write_code_scanning"),
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
// Custom roles can be imported using the `id` of the role. The `id` of the custom role can be found using the [list custom roles in an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles#list-custom-repository-roles-in-an-organization) API.
//
// ```sh
//
//	$ pulumi import github:index/organizationCustomRole:OrganizationCustomRole example 1234
//
// ```
type OrganizationCustomRole struct {
	pulumi.CustomResourceState

	// The system role from which the role inherits permissions. Can be one of: `read`, `triage`, `write`, or `maintain`.
	BaseRole pulumi.StringOutput `pulumi:"baseRole"`
	// The description for the custom role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the custom role.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of additional permissions included in this role. Must have a minimum of 1 additional permission. The list of available permissions can be found using the [list repository fine-grained permissions for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization) API.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
}

// NewOrganizationCustomRole registers a new resource with the given unique name, arguments, and options.
func NewOrganizationCustomRole(ctx *pulumi.Context,
	name string, args *OrganizationCustomRoleArgs, opts ...pulumi.ResourceOption) (*OrganizationCustomRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseRole == nil {
		return nil, errors.New("invalid value for required argument 'BaseRole'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	var resource OrganizationCustomRole
	err := ctx.RegisterResource("github:index/organizationCustomRole:OrganizationCustomRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationCustomRole gets an existing OrganizationCustomRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationCustomRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationCustomRoleState, opts ...pulumi.ResourceOption) (*OrganizationCustomRole, error) {
	var resource OrganizationCustomRole
	err := ctx.ReadResource("github:index/organizationCustomRole:OrganizationCustomRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationCustomRole resources.
type organizationCustomRoleState struct {
	// The system role from which the role inherits permissions. Can be one of: `read`, `triage`, `write`, or `maintain`.
	BaseRole *string `pulumi:"baseRole"`
	// The description for the custom role.
	Description *string `pulumi:"description"`
	// The name of the custom role.
	Name *string `pulumi:"name"`
	// A list of additional permissions included in this role. Must have a minimum of 1 additional permission. The list of available permissions can be found using the [list repository fine-grained permissions for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization) API.
	Permissions []string `pulumi:"permissions"`
}

type OrganizationCustomRoleState struct {
	// The system role from which the role inherits permissions. Can be one of: `read`, `triage`, `write`, or `maintain`.
	BaseRole pulumi.StringPtrInput
	// The description for the custom role.
	Description pulumi.StringPtrInput
	// The name of the custom role.
	Name pulumi.StringPtrInput
	// A list of additional permissions included in this role. Must have a minimum of 1 additional permission. The list of available permissions can be found using the [list repository fine-grained permissions for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization) API.
	Permissions pulumi.StringArrayInput
}

func (OrganizationCustomRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationCustomRoleState)(nil)).Elem()
}

type organizationCustomRoleArgs struct {
	// The system role from which the role inherits permissions. Can be one of: `read`, `triage`, `write`, or `maintain`.
	BaseRole string `pulumi:"baseRole"`
	// The description for the custom role.
	Description *string `pulumi:"description"`
	// The name of the custom role.
	Name *string `pulumi:"name"`
	// A list of additional permissions included in this role. Must have a minimum of 1 additional permission. The list of available permissions can be found using the [list repository fine-grained permissions for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization) API.
	Permissions []string `pulumi:"permissions"`
}

// The set of arguments for constructing a OrganizationCustomRole resource.
type OrganizationCustomRoleArgs struct {
	// The system role from which the role inherits permissions. Can be one of: `read`, `triage`, `write`, or `maintain`.
	BaseRole pulumi.StringInput
	// The description for the custom role.
	Description pulumi.StringPtrInput
	// The name of the custom role.
	Name pulumi.StringPtrInput
	// A list of additional permissions included in this role. Must have a minimum of 1 additional permission. The list of available permissions can be found using the [list repository fine-grained permissions for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization) API.
	Permissions pulumi.StringArrayInput
}

func (OrganizationCustomRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationCustomRoleArgs)(nil)).Elem()
}

type OrganizationCustomRoleInput interface {
	pulumi.Input

	ToOrganizationCustomRoleOutput() OrganizationCustomRoleOutput
	ToOrganizationCustomRoleOutputWithContext(ctx context.Context) OrganizationCustomRoleOutput
}

func (*OrganizationCustomRole) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationCustomRole)(nil)).Elem()
}

func (i *OrganizationCustomRole) ToOrganizationCustomRoleOutput() OrganizationCustomRoleOutput {
	return i.ToOrganizationCustomRoleOutputWithContext(context.Background())
}

func (i *OrganizationCustomRole) ToOrganizationCustomRoleOutputWithContext(ctx context.Context) OrganizationCustomRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationCustomRoleOutput)
}

// OrganizationCustomRoleArrayInput is an input type that accepts OrganizationCustomRoleArray and OrganizationCustomRoleArrayOutput values.
// You can construct a concrete instance of `OrganizationCustomRoleArrayInput` via:
//
//	OrganizationCustomRoleArray{ OrganizationCustomRoleArgs{...} }
type OrganizationCustomRoleArrayInput interface {
	pulumi.Input

	ToOrganizationCustomRoleArrayOutput() OrganizationCustomRoleArrayOutput
	ToOrganizationCustomRoleArrayOutputWithContext(context.Context) OrganizationCustomRoleArrayOutput
}

type OrganizationCustomRoleArray []OrganizationCustomRoleInput

func (OrganizationCustomRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationCustomRole)(nil)).Elem()
}

func (i OrganizationCustomRoleArray) ToOrganizationCustomRoleArrayOutput() OrganizationCustomRoleArrayOutput {
	return i.ToOrganizationCustomRoleArrayOutputWithContext(context.Background())
}

func (i OrganizationCustomRoleArray) ToOrganizationCustomRoleArrayOutputWithContext(ctx context.Context) OrganizationCustomRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationCustomRoleArrayOutput)
}

// OrganizationCustomRoleMapInput is an input type that accepts OrganizationCustomRoleMap and OrganizationCustomRoleMapOutput values.
// You can construct a concrete instance of `OrganizationCustomRoleMapInput` via:
//
//	OrganizationCustomRoleMap{ "key": OrganizationCustomRoleArgs{...} }
type OrganizationCustomRoleMapInput interface {
	pulumi.Input

	ToOrganizationCustomRoleMapOutput() OrganizationCustomRoleMapOutput
	ToOrganizationCustomRoleMapOutputWithContext(context.Context) OrganizationCustomRoleMapOutput
}

type OrganizationCustomRoleMap map[string]OrganizationCustomRoleInput

func (OrganizationCustomRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationCustomRole)(nil)).Elem()
}

func (i OrganizationCustomRoleMap) ToOrganizationCustomRoleMapOutput() OrganizationCustomRoleMapOutput {
	return i.ToOrganizationCustomRoleMapOutputWithContext(context.Background())
}

func (i OrganizationCustomRoleMap) ToOrganizationCustomRoleMapOutputWithContext(ctx context.Context) OrganizationCustomRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationCustomRoleMapOutput)
}

type OrganizationCustomRoleOutput struct{ *pulumi.OutputState }

func (OrganizationCustomRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationCustomRole)(nil)).Elem()
}

func (o OrganizationCustomRoleOutput) ToOrganizationCustomRoleOutput() OrganizationCustomRoleOutput {
	return o
}

func (o OrganizationCustomRoleOutput) ToOrganizationCustomRoleOutputWithContext(ctx context.Context) OrganizationCustomRoleOutput {
	return o
}

// The system role from which the role inherits permissions. Can be one of: `read`, `triage`, `write`, or `maintain`.
func (o OrganizationCustomRoleOutput) BaseRole() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationCustomRole) pulumi.StringOutput { return v.BaseRole }).(pulumi.StringOutput)
}

// The description for the custom role.
func (o OrganizationCustomRoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationCustomRole) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the custom role.
func (o OrganizationCustomRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationCustomRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of additional permissions included in this role. Must have a minimum of 1 additional permission. The list of available permissions can be found using the [list repository fine-grained permissions for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization) API.
func (o OrganizationCustomRoleOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationCustomRole) pulumi.StringArrayOutput { return v.Permissions }).(pulumi.StringArrayOutput)
}

type OrganizationCustomRoleArrayOutput struct{ *pulumi.OutputState }

func (OrganizationCustomRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationCustomRole)(nil)).Elem()
}

func (o OrganizationCustomRoleArrayOutput) ToOrganizationCustomRoleArrayOutput() OrganizationCustomRoleArrayOutput {
	return o
}

func (o OrganizationCustomRoleArrayOutput) ToOrganizationCustomRoleArrayOutputWithContext(ctx context.Context) OrganizationCustomRoleArrayOutput {
	return o
}

func (o OrganizationCustomRoleArrayOutput) Index(i pulumi.IntInput) OrganizationCustomRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationCustomRole {
		return vs[0].([]*OrganizationCustomRole)[vs[1].(int)]
	}).(OrganizationCustomRoleOutput)
}

type OrganizationCustomRoleMapOutput struct{ *pulumi.OutputState }

func (OrganizationCustomRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationCustomRole)(nil)).Elem()
}

func (o OrganizationCustomRoleMapOutput) ToOrganizationCustomRoleMapOutput() OrganizationCustomRoleMapOutput {
	return o
}

func (o OrganizationCustomRoleMapOutput) ToOrganizationCustomRoleMapOutputWithContext(ctx context.Context) OrganizationCustomRoleMapOutput {
	return o
}

func (o OrganizationCustomRoleMapOutput) MapIndex(k pulumi.StringInput) OrganizationCustomRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationCustomRole {
		return vs[0].(map[string]*OrganizationCustomRole)[vs[1].(string)]
	}).(OrganizationCustomRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationCustomRoleInput)(nil)).Elem(), &OrganizationCustomRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationCustomRoleArrayInput)(nil)).Elem(), OrganizationCustomRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationCustomRoleMapInput)(nil)).Elem(), OrganizationCustomRoleMap{})
	pulumi.RegisterOutputType(OrganizationCustomRoleOutput{})
	pulumi.RegisterOutputType(OrganizationCustomRoleArrayOutput{})
	pulumi.RegisterOutputType(OrganizationCustomRoleMapOutput{})
}