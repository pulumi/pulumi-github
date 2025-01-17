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

// This resource allows you to create and manage a specific custom property for a GitHub repository.
//
// ## Example Usage
//
// > Note that this assumes there already is a custom property defined on the org level called `my-cool-property` of type `string`
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
//			example, err := github.NewRepository(ctx, "example", &github.RepositoryArgs{
//				Name:        pulumi.String("example"),
//				Description: pulumi.String("My awesome codebase"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryCustomProperty(ctx, "string", &github.RepositoryCustomPropertyArgs{
//				Repository:   example.Name,
//				PropertyName: pulumi.String("my-cool-property"),
//				PropertyType: pulumi.String("string"),
//				PropertyValues: pulumi.StringArray{
//					pulumi.String("test"),
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
// GitHub Repository Custom Property can be imported using an ID made up of a comibnation of the names of the organization, repository, custom property separated by a `:` character, e.g.
//
// ```sh
// $ pulumi import github:index/repositoryCustomProperty:RepositoryCustomProperty example <organization-name>:<repo-name>:<custom-property-name>
// ```
type RepositoryCustomProperty struct {
	pulumi.CustomResourceState

	// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
	PropertyName pulumi.StringOutput `pulumi:"propertyName"`
	// Type of the custom property. Can be one of `singleSelect`, `multiSelect`, `string`, or `trueFalse`
	PropertyType pulumi.StringOutput `pulumi:"propertyType"`
	// Value of the custom property in the form of an array. Properties of type `singleSelect`, `string`, and `trueFalse` are represented as a string array of length 1
	PropertyValues pulumi.StringArrayOutput `pulumi:"propertyValues"`
	// The repository of the environment.
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewRepositoryCustomProperty registers a new resource with the given unique name, arguments, and options.
func NewRepositoryCustomProperty(ctx *pulumi.Context,
	name string, args *RepositoryCustomPropertyArgs, opts ...pulumi.ResourceOption) (*RepositoryCustomProperty, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PropertyName == nil {
		return nil, errors.New("invalid value for required argument 'PropertyName'")
	}
	if args.PropertyType == nil {
		return nil, errors.New("invalid value for required argument 'PropertyType'")
	}
	if args.PropertyValues == nil {
		return nil, errors.New("invalid value for required argument 'PropertyValues'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryCustomProperty
	err := ctx.RegisterResource("github:index/repositoryCustomProperty:RepositoryCustomProperty", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryCustomProperty gets an existing RepositoryCustomProperty resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryCustomProperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryCustomPropertyState, opts ...pulumi.ResourceOption) (*RepositoryCustomProperty, error) {
	var resource RepositoryCustomProperty
	err := ctx.ReadResource("github:index/repositoryCustomProperty:RepositoryCustomProperty", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryCustomProperty resources.
type repositoryCustomPropertyState struct {
	// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
	PropertyName *string `pulumi:"propertyName"`
	// Type of the custom property. Can be one of `singleSelect`, `multiSelect`, `string`, or `trueFalse`
	PropertyType *string `pulumi:"propertyType"`
	// Value of the custom property in the form of an array. Properties of type `singleSelect`, `string`, and `trueFalse` are represented as a string array of length 1
	PropertyValues []string `pulumi:"propertyValues"`
	// The repository of the environment.
	Repository *string `pulumi:"repository"`
}

type RepositoryCustomPropertyState struct {
	// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
	PropertyName pulumi.StringPtrInput
	// Type of the custom property. Can be one of `singleSelect`, `multiSelect`, `string`, or `trueFalse`
	PropertyType pulumi.StringPtrInput
	// Value of the custom property in the form of an array. Properties of type `singleSelect`, `string`, and `trueFalse` are represented as a string array of length 1
	PropertyValues pulumi.StringArrayInput
	// The repository of the environment.
	Repository pulumi.StringPtrInput
}

func (RepositoryCustomPropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryCustomPropertyState)(nil)).Elem()
}

type repositoryCustomPropertyArgs struct {
	// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
	PropertyName string `pulumi:"propertyName"`
	// Type of the custom property. Can be one of `singleSelect`, `multiSelect`, `string`, or `trueFalse`
	PropertyType string `pulumi:"propertyType"`
	// Value of the custom property in the form of an array. Properties of type `singleSelect`, `string`, and `trueFalse` are represented as a string array of length 1
	PropertyValues []string `pulumi:"propertyValues"`
	// The repository of the environment.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryCustomProperty resource.
type RepositoryCustomPropertyArgs struct {
	// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
	PropertyName pulumi.StringInput
	// Type of the custom property. Can be one of `singleSelect`, `multiSelect`, `string`, or `trueFalse`
	PropertyType pulumi.StringInput
	// Value of the custom property in the form of an array. Properties of type `singleSelect`, `string`, and `trueFalse` are represented as a string array of length 1
	PropertyValues pulumi.StringArrayInput
	// The repository of the environment.
	Repository pulumi.StringInput
}

func (RepositoryCustomPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryCustomPropertyArgs)(nil)).Elem()
}

type RepositoryCustomPropertyInput interface {
	pulumi.Input

	ToRepositoryCustomPropertyOutput() RepositoryCustomPropertyOutput
	ToRepositoryCustomPropertyOutputWithContext(ctx context.Context) RepositoryCustomPropertyOutput
}

func (*RepositoryCustomProperty) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryCustomProperty)(nil)).Elem()
}

func (i *RepositoryCustomProperty) ToRepositoryCustomPropertyOutput() RepositoryCustomPropertyOutput {
	return i.ToRepositoryCustomPropertyOutputWithContext(context.Background())
}

func (i *RepositoryCustomProperty) ToRepositoryCustomPropertyOutputWithContext(ctx context.Context) RepositoryCustomPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCustomPropertyOutput)
}

// RepositoryCustomPropertyArrayInput is an input type that accepts RepositoryCustomPropertyArray and RepositoryCustomPropertyArrayOutput values.
// You can construct a concrete instance of `RepositoryCustomPropertyArrayInput` via:
//
//	RepositoryCustomPropertyArray{ RepositoryCustomPropertyArgs{...} }
type RepositoryCustomPropertyArrayInput interface {
	pulumi.Input

	ToRepositoryCustomPropertyArrayOutput() RepositoryCustomPropertyArrayOutput
	ToRepositoryCustomPropertyArrayOutputWithContext(context.Context) RepositoryCustomPropertyArrayOutput
}

type RepositoryCustomPropertyArray []RepositoryCustomPropertyInput

func (RepositoryCustomPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryCustomProperty)(nil)).Elem()
}

func (i RepositoryCustomPropertyArray) ToRepositoryCustomPropertyArrayOutput() RepositoryCustomPropertyArrayOutput {
	return i.ToRepositoryCustomPropertyArrayOutputWithContext(context.Background())
}

func (i RepositoryCustomPropertyArray) ToRepositoryCustomPropertyArrayOutputWithContext(ctx context.Context) RepositoryCustomPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCustomPropertyArrayOutput)
}

// RepositoryCustomPropertyMapInput is an input type that accepts RepositoryCustomPropertyMap and RepositoryCustomPropertyMapOutput values.
// You can construct a concrete instance of `RepositoryCustomPropertyMapInput` via:
//
//	RepositoryCustomPropertyMap{ "key": RepositoryCustomPropertyArgs{...} }
type RepositoryCustomPropertyMapInput interface {
	pulumi.Input

	ToRepositoryCustomPropertyMapOutput() RepositoryCustomPropertyMapOutput
	ToRepositoryCustomPropertyMapOutputWithContext(context.Context) RepositoryCustomPropertyMapOutput
}

type RepositoryCustomPropertyMap map[string]RepositoryCustomPropertyInput

func (RepositoryCustomPropertyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryCustomProperty)(nil)).Elem()
}

func (i RepositoryCustomPropertyMap) ToRepositoryCustomPropertyMapOutput() RepositoryCustomPropertyMapOutput {
	return i.ToRepositoryCustomPropertyMapOutputWithContext(context.Background())
}

func (i RepositoryCustomPropertyMap) ToRepositoryCustomPropertyMapOutputWithContext(ctx context.Context) RepositoryCustomPropertyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCustomPropertyMapOutput)
}

type RepositoryCustomPropertyOutput struct{ *pulumi.OutputState }

func (RepositoryCustomPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryCustomProperty)(nil)).Elem()
}

func (o RepositoryCustomPropertyOutput) ToRepositoryCustomPropertyOutput() RepositoryCustomPropertyOutput {
	return o
}

func (o RepositoryCustomPropertyOutput) ToRepositoryCustomPropertyOutputWithContext(ctx context.Context) RepositoryCustomPropertyOutput {
	return o
}

// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
func (o RepositoryCustomPropertyOutput) PropertyName() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryCustomProperty) pulumi.StringOutput { return v.PropertyName }).(pulumi.StringOutput)
}

// Type of the custom property. Can be one of `singleSelect`, `multiSelect`, `string`, or `trueFalse`
func (o RepositoryCustomPropertyOutput) PropertyType() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryCustomProperty) pulumi.StringOutput { return v.PropertyType }).(pulumi.StringOutput)
}

// Value of the custom property in the form of an array. Properties of type `singleSelect`, `string`, and `trueFalse` are represented as a string array of length 1
func (o RepositoryCustomPropertyOutput) PropertyValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryCustomProperty) pulumi.StringArrayOutput { return v.PropertyValues }).(pulumi.StringArrayOutput)
}

// The repository of the environment.
func (o RepositoryCustomPropertyOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryCustomProperty) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

type RepositoryCustomPropertyArrayOutput struct{ *pulumi.OutputState }

func (RepositoryCustomPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryCustomProperty)(nil)).Elem()
}

func (o RepositoryCustomPropertyArrayOutput) ToRepositoryCustomPropertyArrayOutput() RepositoryCustomPropertyArrayOutput {
	return o
}

func (o RepositoryCustomPropertyArrayOutput) ToRepositoryCustomPropertyArrayOutputWithContext(ctx context.Context) RepositoryCustomPropertyArrayOutput {
	return o
}

func (o RepositoryCustomPropertyArrayOutput) Index(i pulumi.IntInput) RepositoryCustomPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryCustomProperty {
		return vs[0].([]*RepositoryCustomProperty)[vs[1].(int)]
	}).(RepositoryCustomPropertyOutput)
}

type RepositoryCustomPropertyMapOutput struct{ *pulumi.OutputState }

func (RepositoryCustomPropertyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryCustomProperty)(nil)).Elem()
}

func (o RepositoryCustomPropertyMapOutput) ToRepositoryCustomPropertyMapOutput() RepositoryCustomPropertyMapOutput {
	return o
}

func (o RepositoryCustomPropertyMapOutput) ToRepositoryCustomPropertyMapOutputWithContext(ctx context.Context) RepositoryCustomPropertyMapOutput {
	return o
}

func (o RepositoryCustomPropertyMapOutput) MapIndex(k pulumi.StringInput) RepositoryCustomPropertyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryCustomProperty {
		return vs[0].(map[string]*RepositoryCustomProperty)[vs[1].(string)]
	}).(RepositoryCustomPropertyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryCustomPropertyInput)(nil)).Elem(), &RepositoryCustomProperty{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryCustomPropertyArrayInput)(nil)).Elem(), RepositoryCustomPropertyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryCustomPropertyMapInput)(nil)).Elem(), RepositoryCustomPropertyMap{})
	pulumi.RegisterOutputType(RepositoryCustomPropertyOutput{})
	pulumi.RegisterOutputType(RepositoryCustomPropertyArrayOutput{})
	pulumi.RegisterOutputType(RepositoryCustomPropertyMapOutput{})
}