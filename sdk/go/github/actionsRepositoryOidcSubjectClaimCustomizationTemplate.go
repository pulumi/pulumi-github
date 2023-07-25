// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage an OpenID Connect subject claim customization template for a GitHub
// repository.
//
// More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
// available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).
//
// The following table lists the behaviour of `useDefault`:
//
// | `useDefault` | `includeClaimKeys` | Template used                                             |
// |---------------|----------------------|-----------------------------------------------------------|
// | `true`        | Unset                | GitHub's default                                          |
// | `false`       | Set                  | `includeClaimKeys`                                      |
// | `false`       | Unset                | Organization's default if set, otherwise GitHub's default |
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
//			example, err := github.NewRepository(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewActionsRepositoryOidcSubjectClaimCustomizationTemplate(ctx, "exampleTemplate", &github.ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs{
//				Repository: example.Name,
//				UseDefault: pulumi.Bool(false),
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
// This resource can be imported using the repository's name.
//
// ```sh
//
//	$ pulumi import github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate test example_repository
//
// ```
type ActionsRepositoryOidcSubjectClaimCustomizationTemplate struct {
	pulumi.CustomResourceState

	// A list of OpenID Connect claims.
	IncludeClaimKeys pulumi.StringArrayOutput `pulumi:"includeClaimKeys"`
	// The name of the repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Whether to use the default template or not. If `true`, `includeClaimKeys` must not
	// be set.
	UseDefault pulumi.BoolOutput `pulumi:"useDefault"`
}

// NewActionsRepositoryOidcSubjectClaimCustomizationTemplate registers a new resource with the given unique name, arguments, and options.
func NewActionsRepositoryOidcSubjectClaimCustomizationTemplate(ctx *pulumi.Context,
	name string, args *ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs, opts ...pulumi.ResourceOption) (*ActionsRepositoryOidcSubjectClaimCustomizationTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.UseDefault == nil {
		return nil, errors.New("invalid value for required argument 'UseDefault'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ActionsRepositoryOidcSubjectClaimCustomizationTemplate
	err := ctx.RegisterResource("github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsRepositoryOidcSubjectClaimCustomizationTemplate gets an existing ActionsRepositoryOidcSubjectClaimCustomizationTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsRepositoryOidcSubjectClaimCustomizationTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsRepositoryOidcSubjectClaimCustomizationTemplateState, opts ...pulumi.ResourceOption) (*ActionsRepositoryOidcSubjectClaimCustomizationTemplate, error) {
	var resource ActionsRepositoryOidcSubjectClaimCustomizationTemplate
	err := ctx.ReadResource("github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsRepositoryOidcSubjectClaimCustomizationTemplate resources.
type actionsRepositoryOidcSubjectClaimCustomizationTemplateState struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys []string `pulumi:"includeClaimKeys"`
	// The name of the repository.
	Repository *string `pulumi:"repository"`
	// Whether to use the default template or not. If `true`, `includeClaimKeys` must not
	// be set.
	UseDefault *bool `pulumi:"useDefault"`
}

type ActionsRepositoryOidcSubjectClaimCustomizationTemplateState struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys pulumi.StringArrayInput
	// The name of the repository.
	Repository pulumi.StringPtrInput
	// Whether to use the default template or not. If `true`, `includeClaimKeys` must not
	// be set.
	UseDefault pulumi.BoolPtrInput
}

func (ActionsRepositoryOidcSubjectClaimCustomizationTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsRepositoryOidcSubjectClaimCustomizationTemplateState)(nil)).Elem()
}

type actionsRepositoryOidcSubjectClaimCustomizationTemplateArgs struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys []string `pulumi:"includeClaimKeys"`
	// The name of the repository.
	Repository string `pulumi:"repository"`
	// Whether to use the default template or not. If `true`, `includeClaimKeys` must not
	// be set.
	UseDefault bool `pulumi:"useDefault"`
}

// The set of arguments for constructing a ActionsRepositoryOidcSubjectClaimCustomizationTemplate resource.
type ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs struct {
	// A list of OpenID Connect claims.
	IncludeClaimKeys pulumi.StringArrayInput
	// The name of the repository.
	Repository pulumi.StringInput
	// Whether to use the default template or not. If `true`, `includeClaimKeys` must not
	// be set.
	UseDefault pulumi.BoolInput
}

func (ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsRepositoryOidcSubjectClaimCustomizationTemplateArgs)(nil)).Elem()
}

type ActionsRepositoryOidcSubjectClaimCustomizationTemplateInput interface {
	pulumi.Input

	ToActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput
	ToActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputWithContext(ctx context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput
}

func (*ActionsRepositoryOidcSubjectClaimCustomizationTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsRepositoryOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (i *ActionsRepositoryOidcSubjectClaimCustomizationTemplate) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput {
	return i.ToActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputWithContext(context.Background())
}

func (i *ActionsRepositoryOidcSubjectClaimCustomizationTemplate) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputWithContext(ctx context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput)
}

// ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayInput is an input type that accepts ActionsRepositoryOidcSubjectClaimCustomizationTemplateArray and ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput values.
// You can construct a concrete instance of `ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayInput` via:
//
//	ActionsRepositoryOidcSubjectClaimCustomizationTemplateArray{ ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs{...} }
type ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayInput interface {
	pulumi.Input

	ToActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput
	ToActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput
}

type ActionsRepositoryOidcSubjectClaimCustomizationTemplateArray []ActionsRepositoryOidcSubjectClaimCustomizationTemplateInput

func (ActionsRepositoryOidcSubjectClaimCustomizationTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsRepositoryOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (i ActionsRepositoryOidcSubjectClaimCustomizationTemplateArray) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput {
	return i.ToActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(context.Background())
}

func (i ActionsRepositoryOidcSubjectClaimCustomizationTemplateArray) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(ctx context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput)
}

// ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapInput is an input type that accepts ActionsRepositoryOidcSubjectClaimCustomizationTemplateMap and ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput values.
// You can construct a concrete instance of `ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapInput` via:
//
//	ActionsRepositoryOidcSubjectClaimCustomizationTemplateMap{ "key": ActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs{...} }
type ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapInput interface {
	pulumi.Input

	ToActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput
	ToActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutputWithContext(context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput
}

type ActionsRepositoryOidcSubjectClaimCustomizationTemplateMap map[string]ActionsRepositoryOidcSubjectClaimCustomizationTemplateInput

func (ActionsRepositoryOidcSubjectClaimCustomizationTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsRepositoryOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (i ActionsRepositoryOidcSubjectClaimCustomizationTemplateMap) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput {
	return i.ToActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutputWithContext(context.Background())
}

func (i ActionsRepositoryOidcSubjectClaimCustomizationTemplateMap) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutputWithContext(ctx context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput)
}

type ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput struct{ *pulumi.OutputState }

func (ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsRepositoryOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput {
	return o
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputWithContext(ctx context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput {
	return o
}

// A list of OpenID Connect claims.
func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput) IncludeClaimKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionsRepositoryOidcSubjectClaimCustomizationTemplate) pulumi.StringArrayOutput {
		return v.IncludeClaimKeys
	}).(pulumi.StringArrayOutput)
}

// The name of the repository.
func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsRepositoryOidcSubjectClaimCustomizationTemplate) pulumi.StringOutput {
		return v.Repository
	}).(pulumi.StringOutput)
}

// Whether to use the default template or not. If `true`, `includeClaimKeys` must not
// be set.
func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput) UseDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *ActionsRepositoryOidcSubjectClaimCustomizationTemplate) pulumi.BoolOutput { return v.UseDefault }).(pulumi.BoolOutput)
}

type ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput struct{ *pulumi.OutputState }

func (ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsRepositoryOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput {
	return o
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutputWithContext(ctx context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput {
	return o
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput) Index(i pulumi.IntInput) ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsRepositoryOidcSubjectClaimCustomizationTemplate {
		return vs[0].([]*ActionsRepositoryOidcSubjectClaimCustomizationTemplate)[vs[1].(int)]
	}).(ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput)
}

type ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput struct{ *pulumi.OutputState }

func (ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsRepositoryOidcSubjectClaimCustomizationTemplate)(nil)).Elem()
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput() ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput {
	return o
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput) ToActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutputWithContext(ctx context.Context) ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput {
	return o
}

func (o ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput) MapIndex(k pulumi.StringInput) ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsRepositoryOidcSubjectClaimCustomizationTemplate {
		return vs[0].(map[string]*ActionsRepositoryOidcSubjectClaimCustomizationTemplate)[vs[1].(string)]
	}).(ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsRepositoryOidcSubjectClaimCustomizationTemplateInput)(nil)).Elem(), &ActionsRepositoryOidcSubjectClaimCustomizationTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayInput)(nil)).Elem(), ActionsRepositoryOidcSubjectClaimCustomizationTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapInput)(nil)).Elem(), ActionsRepositoryOidcSubjectClaimCustomizationTemplateMap{})
	pulumi.RegisterOutputType(ActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput{})
	pulumi.RegisterOutputType(ActionsRepositoryOidcSubjectClaimCustomizationTemplateArrayOutput{})
	pulumi.RegisterOutputType(ActionsRepositoryOidcSubjectClaimCustomizationTemplateMapOutput{})
}
