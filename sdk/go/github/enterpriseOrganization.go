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

// This resource allows you to create and manage a GitHub enterprise organization.
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
//			_, err := github.NewEnterpriseOrganization(ctx, "org", &github.EnterpriseOrganizationArgs{
//				EnterpriseId: pulumi.Any(data.Github_enterprise.Enterprise.Id),
//				DisplayName:  pulumi.String("Some Awesome Org"),
//				Description:  pulumi.String("Organization created with terraform"),
//				BillingEmail: pulumi.String("jon@winteriscoming.com"),
//				AdminLogins: pulumi.StringArray{
//					pulumi.String("jon-snow"),
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
// GitHub Enterprise Organization can be imported using the `slug` of the enterprise, combined with the `orgname` of the organization, separated by a `/` character.
//
// ```sh
// $ pulumi import github:index/enterpriseOrganization:EnterpriseOrganization org enterp/some-awesome-org
// ```
type EnterpriseOrganization struct {
	pulumi.CustomResourceState

	// List of organization owner usernames.
	AdminLogins pulumi.StringArrayOutput `pulumi:"adminLogins"`
	// The billing email address.
	BillingEmail pulumi.StringOutput `pulumi:"billingEmail"`
	// The ID of the organization.
	DatabaseId pulumi.IntOutput `pulumi:"databaseId"`
	// The description of the organization.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the organization.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The ID of the enterprise.
	EnterpriseId pulumi.StringOutput `pulumi:"enterpriseId"`
	// The name of the organization.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEnterpriseOrganization registers a new resource with the given unique name, arguments, and options.
func NewEnterpriseOrganization(ctx *pulumi.Context,
	name string, args *EnterpriseOrganizationArgs, opts ...pulumi.ResourceOption) (*EnterpriseOrganization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminLogins == nil {
		return nil, errors.New("invalid value for required argument 'AdminLogins'")
	}
	if args.BillingEmail == nil {
		return nil, errors.New("invalid value for required argument 'BillingEmail'")
	}
	if args.EnterpriseId == nil {
		return nil, errors.New("invalid value for required argument 'EnterpriseId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnterpriseOrganization
	err := ctx.RegisterResource("github:index/enterpriseOrganization:EnterpriseOrganization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterpriseOrganization gets an existing EnterpriseOrganization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterpriseOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseOrganizationState, opts ...pulumi.ResourceOption) (*EnterpriseOrganization, error) {
	var resource EnterpriseOrganization
	err := ctx.ReadResource("github:index/enterpriseOrganization:EnterpriseOrganization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterpriseOrganization resources.
type enterpriseOrganizationState struct {
	// List of organization owner usernames.
	AdminLogins []string `pulumi:"adminLogins"`
	// The billing email address.
	BillingEmail *string `pulumi:"billingEmail"`
	// The ID of the organization.
	DatabaseId *int `pulumi:"databaseId"`
	// The description of the organization.
	Description *string `pulumi:"description"`
	// The display name of the organization.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the enterprise.
	EnterpriseId *string `pulumi:"enterpriseId"`
	// The name of the organization.
	Name *string `pulumi:"name"`
}

type EnterpriseOrganizationState struct {
	// List of organization owner usernames.
	AdminLogins pulumi.StringArrayInput
	// The billing email address.
	BillingEmail pulumi.StringPtrInput
	// The ID of the organization.
	DatabaseId pulumi.IntPtrInput
	// The description of the organization.
	Description pulumi.StringPtrInput
	// The display name of the organization.
	DisplayName pulumi.StringPtrInput
	// The ID of the enterprise.
	EnterpriseId pulumi.StringPtrInput
	// The name of the organization.
	Name pulumi.StringPtrInput
}

func (EnterpriseOrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseOrganizationState)(nil)).Elem()
}

type enterpriseOrganizationArgs struct {
	// List of organization owner usernames.
	AdminLogins []string `pulumi:"adminLogins"`
	// The billing email address.
	BillingEmail string `pulumi:"billingEmail"`
	// The description of the organization.
	Description *string `pulumi:"description"`
	// The display name of the organization.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the enterprise.
	EnterpriseId string `pulumi:"enterpriseId"`
	// The name of the organization.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a EnterpriseOrganization resource.
type EnterpriseOrganizationArgs struct {
	// List of organization owner usernames.
	AdminLogins pulumi.StringArrayInput
	// The billing email address.
	BillingEmail pulumi.StringInput
	// The description of the organization.
	Description pulumi.StringPtrInput
	// The display name of the organization.
	DisplayName pulumi.StringPtrInput
	// The ID of the enterprise.
	EnterpriseId pulumi.StringInput
	// The name of the organization.
	Name pulumi.StringPtrInput
}

func (EnterpriseOrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseOrganizationArgs)(nil)).Elem()
}

type EnterpriseOrganizationInput interface {
	pulumi.Input

	ToEnterpriseOrganizationOutput() EnterpriseOrganizationOutput
	ToEnterpriseOrganizationOutputWithContext(ctx context.Context) EnterpriseOrganizationOutput
}

func (*EnterpriseOrganization) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseOrganization)(nil)).Elem()
}

func (i *EnterpriseOrganization) ToEnterpriseOrganizationOutput() EnterpriseOrganizationOutput {
	return i.ToEnterpriseOrganizationOutputWithContext(context.Background())
}

func (i *EnterpriseOrganization) ToEnterpriseOrganizationOutputWithContext(ctx context.Context) EnterpriseOrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseOrganizationOutput)
}

// EnterpriseOrganizationArrayInput is an input type that accepts EnterpriseOrganizationArray and EnterpriseOrganizationArrayOutput values.
// You can construct a concrete instance of `EnterpriseOrganizationArrayInput` via:
//
//	EnterpriseOrganizationArray{ EnterpriseOrganizationArgs{...} }
type EnterpriseOrganizationArrayInput interface {
	pulumi.Input

	ToEnterpriseOrganizationArrayOutput() EnterpriseOrganizationArrayOutput
	ToEnterpriseOrganizationArrayOutputWithContext(context.Context) EnterpriseOrganizationArrayOutput
}

type EnterpriseOrganizationArray []EnterpriseOrganizationInput

func (EnterpriseOrganizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseOrganization)(nil)).Elem()
}

func (i EnterpriseOrganizationArray) ToEnterpriseOrganizationArrayOutput() EnterpriseOrganizationArrayOutput {
	return i.ToEnterpriseOrganizationArrayOutputWithContext(context.Background())
}

func (i EnterpriseOrganizationArray) ToEnterpriseOrganizationArrayOutputWithContext(ctx context.Context) EnterpriseOrganizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseOrganizationArrayOutput)
}

// EnterpriseOrganizationMapInput is an input type that accepts EnterpriseOrganizationMap and EnterpriseOrganizationMapOutput values.
// You can construct a concrete instance of `EnterpriseOrganizationMapInput` via:
//
//	EnterpriseOrganizationMap{ "key": EnterpriseOrganizationArgs{...} }
type EnterpriseOrganizationMapInput interface {
	pulumi.Input

	ToEnterpriseOrganizationMapOutput() EnterpriseOrganizationMapOutput
	ToEnterpriseOrganizationMapOutputWithContext(context.Context) EnterpriseOrganizationMapOutput
}

type EnterpriseOrganizationMap map[string]EnterpriseOrganizationInput

func (EnterpriseOrganizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseOrganization)(nil)).Elem()
}

func (i EnterpriseOrganizationMap) ToEnterpriseOrganizationMapOutput() EnterpriseOrganizationMapOutput {
	return i.ToEnterpriseOrganizationMapOutputWithContext(context.Background())
}

func (i EnterpriseOrganizationMap) ToEnterpriseOrganizationMapOutputWithContext(ctx context.Context) EnterpriseOrganizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseOrganizationMapOutput)
}

type EnterpriseOrganizationOutput struct{ *pulumi.OutputState }

func (EnterpriseOrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseOrganization)(nil)).Elem()
}

func (o EnterpriseOrganizationOutput) ToEnterpriseOrganizationOutput() EnterpriseOrganizationOutput {
	return o
}

func (o EnterpriseOrganizationOutput) ToEnterpriseOrganizationOutputWithContext(ctx context.Context) EnterpriseOrganizationOutput {
	return o
}

// List of organization owner usernames.
func (o EnterpriseOrganizationOutput) AdminLogins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EnterpriseOrganization) pulumi.StringArrayOutput { return v.AdminLogins }).(pulumi.StringArrayOutput)
}

// The billing email address.
func (o EnterpriseOrganizationOutput) BillingEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseOrganization) pulumi.StringOutput { return v.BillingEmail }).(pulumi.StringOutput)
}

// The ID of the organization.
func (o EnterpriseOrganizationOutput) DatabaseId() pulumi.IntOutput {
	return o.ApplyT(func(v *EnterpriseOrganization) pulumi.IntOutput { return v.DatabaseId }).(pulumi.IntOutput)
}

// The description of the organization.
func (o EnterpriseOrganizationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseOrganization) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the organization.
func (o EnterpriseOrganizationOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseOrganization) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the enterprise.
func (o EnterpriseOrganizationOutput) EnterpriseId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseOrganization) pulumi.StringOutput { return v.EnterpriseId }).(pulumi.StringOutput)
}

// The name of the organization.
func (o EnterpriseOrganizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseOrganization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type EnterpriseOrganizationArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseOrganizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnterpriseOrganization)(nil)).Elem()
}

func (o EnterpriseOrganizationArrayOutput) ToEnterpriseOrganizationArrayOutput() EnterpriseOrganizationArrayOutput {
	return o
}

func (o EnterpriseOrganizationArrayOutput) ToEnterpriseOrganizationArrayOutputWithContext(ctx context.Context) EnterpriseOrganizationArrayOutput {
	return o
}

func (o EnterpriseOrganizationArrayOutput) Index(i pulumi.IntInput) EnterpriseOrganizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnterpriseOrganization {
		return vs[0].([]*EnterpriseOrganization)[vs[1].(int)]
	}).(EnterpriseOrganizationOutput)
}

type EnterpriseOrganizationMapOutput struct{ *pulumi.OutputState }

func (EnterpriseOrganizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnterpriseOrganization)(nil)).Elem()
}

func (o EnterpriseOrganizationMapOutput) ToEnterpriseOrganizationMapOutput() EnterpriseOrganizationMapOutput {
	return o
}

func (o EnterpriseOrganizationMapOutput) ToEnterpriseOrganizationMapOutputWithContext(ctx context.Context) EnterpriseOrganizationMapOutput {
	return o
}

func (o EnterpriseOrganizationMapOutput) MapIndex(k pulumi.StringInput) EnterpriseOrganizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnterpriseOrganization {
		return vs[0].(map[string]*EnterpriseOrganization)[vs[1].(string)]
	}).(EnterpriseOrganizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseOrganizationInput)(nil)).Elem(), &EnterpriseOrganization{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseOrganizationArrayInput)(nil)).Elem(), EnterpriseOrganizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnterpriseOrganizationMapInput)(nil)).Elem(), EnterpriseOrganizationMap{})
	pulumi.RegisterOutputType(EnterpriseOrganizationOutput{})
	pulumi.RegisterOutputType(EnterpriseOrganizationArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseOrganizationMapOutput{})
}
