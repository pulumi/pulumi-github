// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage GitHub Actions variables within your GitHub repositories.
// You must have write access to a repository to use this resource.
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
//			_, err := github.NewActionsVariable(ctx, "exampleVariable", &github.ActionsVariableArgs{
//				Repository:   pulumi.String("example_repository"),
//				Value:        pulumi.String("example_variable_value"),
//				VariableName: pulumi.String("example_variable_name"),
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
// GitHub Actions variables can be imported using an ID made up of `repository:variable_name`, e.g.
//
// ```sh
//
//	$ pulumi import github:index/actionsVariable:ActionsVariable myvariable myrepo:myvariable
//
// ```
type ActionsVariable struct {
	pulumi.CustomResourceState

	// Date of actionsVariable creation.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the repository
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Date of actionsVariable update.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Value of the variable
	Value pulumi.StringOutput `pulumi:"value"`
	// Name of the variable
	VariableName pulumi.StringOutput `pulumi:"variableName"`
}

// NewActionsVariable registers a new resource with the given unique name, arguments, and options.
func NewActionsVariable(ctx *pulumi.Context,
	name string, args *ActionsVariableArgs, opts ...pulumi.ResourceOption) (*ActionsVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.VariableName == nil {
		return nil, errors.New("invalid value for required argument 'VariableName'")
	}
	var resource ActionsVariable
	err := ctx.RegisterResource("github:index/actionsVariable:ActionsVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsVariable gets an existing ActionsVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsVariableState, opts ...pulumi.ResourceOption) (*ActionsVariable, error) {
	var resource ActionsVariable
	err := ctx.ReadResource("github:index/actionsVariable:ActionsVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsVariable resources.
type actionsVariableState struct {
	// Date of actionsVariable creation.
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the repository
	Repository *string `pulumi:"repository"`
	// Date of actionsVariable update.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Value of the variable
	Value *string `pulumi:"value"`
	// Name of the variable
	VariableName *string `pulumi:"variableName"`
}

type ActionsVariableState struct {
	// Date of actionsVariable creation.
	CreatedAt pulumi.StringPtrInput
	// Name of the repository
	Repository pulumi.StringPtrInput
	// Date of actionsVariable update.
	UpdatedAt pulumi.StringPtrInput
	// Value of the variable
	Value pulumi.StringPtrInput
	// Name of the variable
	VariableName pulumi.StringPtrInput
}

func (ActionsVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsVariableState)(nil)).Elem()
}

type actionsVariableArgs struct {
	// Name of the repository
	Repository string `pulumi:"repository"`
	// Value of the variable
	Value string `pulumi:"value"`
	// Name of the variable
	VariableName string `pulumi:"variableName"`
}

// The set of arguments for constructing a ActionsVariable resource.
type ActionsVariableArgs struct {
	// Name of the repository
	Repository pulumi.StringInput
	// Value of the variable
	Value pulumi.StringInput
	// Name of the variable
	VariableName pulumi.StringInput
}

func (ActionsVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsVariableArgs)(nil)).Elem()
}

type ActionsVariableInput interface {
	pulumi.Input

	ToActionsVariableOutput() ActionsVariableOutput
	ToActionsVariableOutputWithContext(ctx context.Context) ActionsVariableOutput
}

func (*ActionsVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsVariable)(nil)).Elem()
}

func (i *ActionsVariable) ToActionsVariableOutput() ActionsVariableOutput {
	return i.ToActionsVariableOutputWithContext(context.Background())
}

func (i *ActionsVariable) ToActionsVariableOutputWithContext(ctx context.Context) ActionsVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsVariableOutput)
}

// ActionsVariableArrayInput is an input type that accepts ActionsVariableArray and ActionsVariableArrayOutput values.
// You can construct a concrete instance of `ActionsVariableArrayInput` via:
//
//	ActionsVariableArray{ ActionsVariableArgs{...} }
type ActionsVariableArrayInput interface {
	pulumi.Input

	ToActionsVariableArrayOutput() ActionsVariableArrayOutput
	ToActionsVariableArrayOutputWithContext(context.Context) ActionsVariableArrayOutput
}

type ActionsVariableArray []ActionsVariableInput

func (ActionsVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsVariable)(nil)).Elem()
}

func (i ActionsVariableArray) ToActionsVariableArrayOutput() ActionsVariableArrayOutput {
	return i.ToActionsVariableArrayOutputWithContext(context.Background())
}

func (i ActionsVariableArray) ToActionsVariableArrayOutputWithContext(ctx context.Context) ActionsVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsVariableArrayOutput)
}

// ActionsVariableMapInput is an input type that accepts ActionsVariableMap and ActionsVariableMapOutput values.
// You can construct a concrete instance of `ActionsVariableMapInput` via:
//
//	ActionsVariableMap{ "key": ActionsVariableArgs{...} }
type ActionsVariableMapInput interface {
	pulumi.Input

	ToActionsVariableMapOutput() ActionsVariableMapOutput
	ToActionsVariableMapOutputWithContext(context.Context) ActionsVariableMapOutput
}

type ActionsVariableMap map[string]ActionsVariableInput

func (ActionsVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsVariable)(nil)).Elem()
}

func (i ActionsVariableMap) ToActionsVariableMapOutput() ActionsVariableMapOutput {
	return i.ToActionsVariableMapOutputWithContext(context.Background())
}

func (i ActionsVariableMap) ToActionsVariableMapOutputWithContext(ctx context.Context) ActionsVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsVariableMapOutput)
}

type ActionsVariableOutput struct{ *pulumi.OutputState }

func (ActionsVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsVariable)(nil)).Elem()
}

func (o ActionsVariableOutput) ToActionsVariableOutput() ActionsVariableOutput {
	return o
}

func (o ActionsVariableOutput) ToActionsVariableOutputWithContext(ctx context.Context) ActionsVariableOutput {
	return o
}

// Date of actionsVariable creation.
func (o ActionsVariableOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsVariable) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the repository
func (o ActionsVariableOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsVariable) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// Date of actionsVariable update.
func (o ActionsVariableOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsVariable) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Value of the variable
func (o ActionsVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsVariable) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// Name of the variable
func (o ActionsVariableOutput) VariableName() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsVariable) pulumi.StringOutput { return v.VariableName }).(pulumi.StringOutput)
}

type ActionsVariableArrayOutput struct{ *pulumi.OutputState }

func (ActionsVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsVariable)(nil)).Elem()
}

func (o ActionsVariableArrayOutput) ToActionsVariableArrayOutput() ActionsVariableArrayOutput {
	return o
}

func (o ActionsVariableArrayOutput) ToActionsVariableArrayOutputWithContext(ctx context.Context) ActionsVariableArrayOutput {
	return o
}

func (o ActionsVariableArrayOutput) Index(i pulumi.IntInput) ActionsVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsVariable {
		return vs[0].([]*ActionsVariable)[vs[1].(int)]
	}).(ActionsVariableOutput)
}

type ActionsVariableMapOutput struct{ *pulumi.OutputState }

func (ActionsVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsVariable)(nil)).Elem()
}

func (o ActionsVariableMapOutput) ToActionsVariableMapOutput() ActionsVariableMapOutput {
	return o
}

func (o ActionsVariableMapOutput) ToActionsVariableMapOutputWithContext(ctx context.Context) ActionsVariableMapOutput {
	return o
}

func (o ActionsVariableMapOutput) MapIndex(k pulumi.StringInput) ActionsVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsVariable {
		return vs[0].(map[string]*ActionsVariable)[vs[1].(string)]
	}).(ActionsVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsVariableInput)(nil)).Elem(), &ActionsVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsVariableArrayInput)(nil)).Elem(), ActionsVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsVariableMapInput)(nil)).Elem(), ActionsVariableMap{})
	pulumi.RegisterOutputType(ActionsVariableOutput{})
	pulumi.RegisterOutputType(ActionsVariableArrayOutput{})
	pulumi.RegisterOutputType(ActionsVariableMapOutput{})
}
