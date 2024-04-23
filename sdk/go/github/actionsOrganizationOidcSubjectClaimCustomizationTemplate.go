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

// This resource allows you to create and manage an OpenID Connect subject claim customization template within a GitHub
// organization.
//
// More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
// available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).
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
//			_, err := github.NewActionsOrganizationOidcSubjectClaimCustomizationTemplate(ctx, "example_template", &github.ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs{
//				IncludeClaimKeys: pulumi.StringArray{
//					pulumi.String("actor"),
//					pulumi.String("context"),
//					pulumi.String("repository_owner"),
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
// This resource can be imported using the organization's name.
//
// ```sh
// $ pulumi import github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate test example_organization
// ```
type ActionsOrganizationOidcSubjectClaimCustomizationTemplate struct {
	pulumi.CustomResourceState

	// A list of OpenID Connect claims.
	IncludeClaimKeys pulumi.StringArrayOutput `pulumi:"includeClaimKeys"`
}

// NewActionsOrganizationOidcSubjectClaimCustomizationTemplate registers a new resource with the given unique name, arguments, and options.
func NewActionsOrganizationOidcSubjectClaimCustomizationTemplate(ctx *pulumi.Context,
	name string, args *ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs, opts ...pulumi.ResourceOption) (*ActionsOrganizationOidcSubjectClaimCustomizationTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IncludeClaimKeys == nil {
		return nil, errors.New("invalid value for required argument 'IncludeClaimKeys'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ActionsOrganizationOidcSubjectClaimCustomizationTemplate
	err := ctx.RegisterResource("github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsOrganizationOidcSubjectClaimCustomizationTemplate gets an existing ActionsOrganizationOidcSubjectClaimCustomizationTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsOrganizationOidcSubjectClaimCustomizationTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsOrganizationOidcSubjectClaimCustomizationTemplateState, opts ...pulumi.ResourceOption) (*ActionsOrganizationOidcSubjectClaimCustomizationTemplate, error) {
	var resource ActionsOrganizationOidcSubjectClaimCustomizationTemplate
	err := ctx.ReadResource("github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsOrganizationOidcSubjectClaimCustomizationTemplate resources.
type actionsOrganizationOidcSubjectClaimCustomizationTemplateState struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys []string `pulumi:"includeClaimKeys"`
}

type ActionsOrganizationOidcSubjectClaimCustomizationTemplateState struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys pulumi.StringArrayInput
}

func (ActionsOrganizationOidcSubjectClaimCustomizationTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationOidcSubjectClaimCustomizationTemplateState)(nil)).Elem()
}

type actionsOrganizationOidcSubjectClaimCustomizationTemplateArgs struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys []string `pulumi:"includeClaimKeys"`
}

// The set of arguments for constructing a ActionsOrganizationOidcSubjectClaimCustomizationTemplate resource.
type ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys pulumi.StringArrayInput
}

func (ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationOidcSubjectClaimCustomizationTemplateArgs)(nil)).Elem()
}

type ActionsOrganizationOidcSubjectClaimCustomizationTemplateInput interface {
	pulumi.Input

	ToActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput
	ToActionsOrganizationOidcSubjectClaimCustomizationTemplateOutputWithContext(ctx context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput
}

func (*ActionsOrganizationOidcSubjectClaimCustomizationTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (i *ActionsOrganizationOidcSubjectClaimCustomizationTemplate) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput {
	return i.ToActionsOrganizationOidcSubjectClaimCustomizationTemplateOutputWithContext(context.Background())
}

func (i *ActionsOrganizationOidcSubjectClaimCustomizationTemplate) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateOutputWithContext(ctx context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput)
}

// ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayInput is an input type that accepts ActionsOrganizationOidcSubjectClaimCustomizationTemplateArray and ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput values.
// You can construct a concrete instance of `ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayInput` via:
//
//	ActionsOrganizationOidcSubjectClaimCustomizationTemplateArray{ ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs{...} }
type ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayInput interface {
	pulumi.Input

	ToActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput
	ToActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput
}

type ActionsOrganizationOidcSubjectClaimCustomizationTemplateArray []ActionsOrganizationOidcSubjectClaimCustomizationTemplateInput

func (ActionsOrganizationOidcSubjectClaimCustomizationTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (i ActionsOrganizationOidcSubjectClaimCustomizationTemplateArray) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput {
	return i.ToActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(context.Background())
}

func (i ActionsOrganizationOidcSubjectClaimCustomizationTemplateArray) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(ctx context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput)
}

// ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapInput is an input type that accepts ActionsOrganizationOidcSubjectClaimCustomizationTemplateMap and ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput values.
// You can construct a concrete instance of `ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapInput` via:
//
//	ActionsOrganizationOidcSubjectClaimCustomizationTemplateMap{ "key": ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs{...} }
type ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapInput interface {
	pulumi.Input

	ToActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput
	ToActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutputWithContext(context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput
}

type ActionsOrganizationOidcSubjectClaimCustomizationTemplateMap map[string]ActionsOrganizationOidcSubjectClaimCustomizationTemplateInput

func (ActionsOrganizationOidcSubjectClaimCustomizationTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (i ActionsOrganizationOidcSubjectClaimCustomizationTemplateMap) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput {
	return i.ToActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutputWithContext(context.Background())
}

func (i ActionsOrganizationOidcSubjectClaimCustomizationTemplateMap) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutputWithContext(ctx context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput)
}

type ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput {
	return o
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateOutputWithContext(ctx context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput {
	return o
}

// A list of OpenID Connect claims.
func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput) IncludeClaimKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionsOrganizationOidcSubjectClaimCustomizationTemplate) pulumi.StringArrayOutput {
		return v.IncludeClaimKeys
	}).(pulumi.StringArrayOutput)
}

type ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput {
	return o
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(ctx context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput {
	return o
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput) Index(i pulumi.IntInput) ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsOrganizationOidcSubjectClaimCustomizationTemplate {
		return vs[0].([]*ActionsOrganizationOidcSubjectClaimCustomizationTemplate)[vs[1].(int)]
	}).(ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput)
}

type ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput() ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput {
	return o
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput) ToActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutputWithContext(ctx context.Context) ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput {
	return o
}

func (o ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput) MapIndex(k pulumi.StringInput) ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsOrganizationOidcSubjectClaimCustomizationTemplate {
		return vs[0].(map[string]*ActionsOrganizationOidcSubjectClaimCustomizationTemplate)[vs[1].(string)]
	}).(ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationOidcSubjectClaimCustomizationTemplateInput)(nil)).Elem(), &ActionsOrganizationOidcSubjectClaimCustomizationTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayInput)(nil)).Elem(), ActionsOrganizationOidcSubjectClaimCustomizationTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapInput)(nil)).Elem(), ActionsOrganizationOidcSubjectClaimCustomizationTemplateMap{})
	pulumi.RegisterOutputType(ActionsOrganizationOidcSubjectClaimCustomizationTemplateOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationOidcSubjectClaimCustomizationTemplateArrayOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationOidcSubjectClaimCustomizationTemplateMapOutput{})
}
