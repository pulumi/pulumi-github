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

// This resource allows you to create and manage GitHub Actions variables within your GitHub organization.
// You must have write access to a repository to use this resource.
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
//			_, err := github.NewActionsOrganizationVariable(ctx, "example_variable", &github.ActionsOrganizationVariableArgs{
//				VariableName: pulumi.String("example_variable_name"),
//				Visibility:   pulumi.String("private"),
//				Value:        pulumi.String("example_variable_value"),
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
//			repo, err := github.LookupRepository(ctx, &github.LookupRepositoryArgs{
//				FullName: pulumi.StringRef("my-org/repo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewActionsOrganizationVariable(ctx, "example_variable", &github.ActionsOrganizationVariableArgs{
//				VariableName: pulumi.String("example_variable_name"),
//				Visibility:   pulumi.String("selected"),
//				Value:        pulumi.String("example_variable_value"),
//				SelectedRepositoryIds: pulumi.IntArray{
//					pulumi.Int(repo.RepoId),
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
// This resource can be imported using an ID made up of the variable name:
//
// ```sh
// $ pulumi import github:index/actionsOrganizationVariable:ActionsOrganizationVariable test_variable test_variable_name
// ```
type ActionsOrganizationVariable struct {
	pulumi.CustomResourceState

	// Date of actionsVariable creation.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// An array of repository ids that can access the organization variable.
	SelectedRepositoryIds pulumi.IntArrayOutput `pulumi:"selectedRepositoryIds"`
	// Date of actionsVariable update.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Value of the variable
	Value pulumi.StringOutput `pulumi:"value"`
	// Name of the variable
	VariableName pulumi.StringOutput `pulumi:"variableName"`
	// Configures the access that repositories have to the organization variable.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewActionsOrganizationVariable registers a new resource with the given unique name, arguments, and options.
func NewActionsOrganizationVariable(ctx *pulumi.Context,
	name string, args *ActionsOrganizationVariableArgs, opts ...pulumi.ResourceOption) (*ActionsOrganizationVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.VariableName == nil {
		return nil, errors.New("invalid value for required argument 'VariableName'")
	}
	if args.Visibility == nil {
		return nil, errors.New("invalid value for required argument 'Visibility'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ActionsOrganizationVariable
	err := ctx.RegisterResource("github:index/actionsOrganizationVariable:ActionsOrganizationVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionsOrganizationVariable gets an existing ActionsOrganizationVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionsOrganizationVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionsOrganizationVariableState, opts ...pulumi.ResourceOption) (*ActionsOrganizationVariable, error) {
	var resource ActionsOrganizationVariable
	err := ctx.ReadResource("github:index/actionsOrganizationVariable:ActionsOrganizationVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionsOrganizationVariable resources.
type actionsOrganizationVariableState struct {
	// Date of actionsVariable creation.
	CreatedAt *string `pulumi:"createdAt"`
	// An array of repository ids that can access the organization variable.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
	// Date of actionsVariable update.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Value of the variable
	Value *string `pulumi:"value"`
	// Name of the variable
	VariableName *string `pulumi:"variableName"`
	// Configures the access that repositories have to the organization variable.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility *string `pulumi:"visibility"`
}

type ActionsOrganizationVariableState struct {
	// Date of actionsVariable creation.
	CreatedAt pulumi.StringPtrInput
	// An array of repository ids that can access the organization variable.
	SelectedRepositoryIds pulumi.IntArrayInput
	// Date of actionsVariable update.
	UpdatedAt pulumi.StringPtrInput
	// Value of the variable
	Value pulumi.StringPtrInput
	// Name of the variable
	VariableName pulumi.StringPtrInput
	// Configures the access that repositories have to the organization variable.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringPtrInput
}

func (ActionsOrganizationVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationVariableState)(nil)).Elem()
}

type actionsOrganizationVariableArgs struct {
	// An array of repository ids that can access the organization variable.
	SelectedRepositoryIds []int `pulumi:"selectedRepositoryIds"`
	// Value of the variable
	Value string `pulumi:"value"`
	// Name of the variable
	VariableName string `pulumi:"variableName"`
	// Configures the access that repositories have to the organization variable.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility string `pulumi:"visibility"`
}

// The set of arguments for constructing a ActionsOrganizationVariable resource.
type ActionsOrganizationVariableArgs struct {
	// An array of repository ids that can access the organization variable.
	SelectedRepositoryIds pulumi.IntArrayInput
	// Value of the variable
	Value pulumi.StringInput
	// Name of the variable
	VariableName pulumi.StringInput
	// Configures the access that repositories have to the organization variable.
	// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
	Visibility pulumi.StringInput
}

func (ActionsOrganizationVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionsOrganizationVariableArgs)(nil)).Elem()
}

type ActionsOrganizationVariableInput interface {
	pulumi.Input

	ToActionsOrganizationVariableOutput() ActionsOrganizationVariableOutput
	ToActionsOrganizationVariableOutputWithContext(ctx context.Context) ActionsOrganizationVariableOutput
}

func (*ActionsOrganizationVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationVariable)(nil)).Elem()
}

func (i *ActionsOrganizationVariable) ToActionsOrganizationVariableOutput() ActionsOrganizationVariableOutput {
	return i.ToActionsOrganizationVariableOutputWithContext(context.Background())
}

func (i *ActionsOrganizationVariable) ToActionsOrganizationVariableOutputWithContext(ctx context.Context) ActionsOrganizationVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationVariableOutput)
}

// ActionsOrganizationVariableArrayInput is an input type that accepts ActionsOrganizationVariableArray and ActionsOrganizationVariableArrayOutput values.
// You can construct a concrete instance of `ActionsOrganizationVariableArrayInput` via:
//
//	ActionsOrganizationVariableArray{ ActionsOrganizationVariableArgs{...} }
type ActionsOrganizationVariableArrayInput interface {
	pulumi.Input

	ToActionsOrganizationVariableArrayOutput() ActionsOrganizationVariableArrayOutput
	ToActionsOrganizationVariableArrayOutputWithContext(context.Context) ActionsOrganizationVariableArrayOutput
}

type ActionsOrganizationVariableArray []ActionsOrganizationVariableInput

func (ActionsOrganizationVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationVariable)(nil)).Elem()
}

func (i ActionsOrganizationVariableArray) ToActionsOrganizationVariableArrayOutput() ActionsOrganizationVariableArrayOutput {
	return i.ToActionsOrganizationVariableArrayOutputWithContext(context.Background())
}

func (i ActionsOrganizationVariableArray) ToActionsOrganizationVariableArrayOutputWithContext(ctx context.Context) ActionsOrganizationVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationVariableArrayOutput)
}

// ActionsOrganizationVariableMapInput is an input type that accepts ActionsOrganizationVariableMap and ActionsOrganizationVariableMapOutput values.
// You can construct a concrete instance of `ActionsOrganizationVariableMapInput` via:
//
//	ActionsOrganizationVariableMap{ "key": ActionsOrganizationVariableArgs{...} }
type ActionsOrganizationVariableMapInput interface {
	pulumi.Input

	ToActionsOrganizationVariableMapOutput() ActionsOrganizationVariableMapOutput
	ToActionsOrganizationVariableMapOutputWithContext(context.Context) ActionsOrganizationVariableMapOutput
}

type ActionsOrganizationVariableMap map[string]ActionsOrganizationVariableInput

func (ActionsOrganizationVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationVariable)(nil)).Elem()
}

func (i ActionsOrganizationVariableMap) ToActionsOrganizationVariableMapOutput() ActionsOrganizationVariableMapOutput {
	return i.ToActionsOrganizationVariableMapOutputWithContext(context.Background())
}

func (i ActionsOrganizationVariableMap) ToActionsOrganizationVariableMapOutputWithContext(ctx context.Context) ActionsOrganizationVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOrganizationVariableMapOutput)
}

type ActionsOrganizationVariableOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsOrganizationVariable)(nil)).Elem()
}

func (o ActionsOrganizationVariableOutput) ToActionsOrganizationVariableOutput() ActionsOrganizationVariableOutput {
	return o
}

func (o ActionsOrganizationVariableOutput) ToActionsOrganizationVariableOutputWithContext(ctx context.Context) ActionsOrganizationVariableOutput {
	return o
}

// Date of actionsVariable creation.
func (o ActionsOrganizationVariableOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationVariable) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// An array of repository ids that can access the organization variable.
func (o ActionsOrganizationVariableOutput) SelectedRepositoryIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ActionsOrganizationVariable) pulumi.IntArrayOutput { return v.SelectedRepositoryIds }).(pulumi.IntArrayOutput)
}

// Date of actionsVariable update.
func (o ActionsOrganizationVariableOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationVariable) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Value of the variable
func (o ActionsOrganizationVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationVariable) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// Name of the variable
func (o ActionsOrganizationVariableOutput) VariableName() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationVariable) pulumi.StringOutput { return v.VariableName }).(pulumi.StringOutput)
}

// Configures the access that repositories have to the organization variable.
// Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
func (o ActionsOrganizationVariableOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *ActionsOrganizationVariable) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type ActionsOrganizationVariableArrayOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionsOrganizationVariable)(nil)).Elem()
}

func (o ActionsOrganizationVariableArrayOutput) ToActionsOrganizationVariableArrayOutput() ActionsOrganizationVariableArrayOutput {
	return o
}

func (o ActionsOrganizationVariableArrayOutput) ToActionsOrganizationVariableArrayOutputWithContext(ctx context.Context) ActionsOrganizationVariableArrayOutput {
	return o
}

func (o ActionsOrganizationVariableArrayOutput) Index(i pulumi.IntInput) ActionsOrganizationVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ActionsOrganizationVariable {
		return vs[0].([]*ActionsOrganizationVariable)[vs[1].(int)]
	}).(ActionsOrganizationVariableOutput)
}

type ActionsOrganizationVariableMapOutput struct{ *pulumi.OutputState }

func (ActionsOrganizationVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionsOrganizationVariable)(nil)).Elem()
}

func (o ActionsOrganizationVariableMapOutput) ToActionsOrganizationVariableMapOutput() ActionsOrganizationVariableMapOutput {
	return o
}

func (o ActionsOrganizationVariableMapOutput) ToActionsOrganizationVariableMapOutputWithContext(ctx context.Context) ActionsOrganizationVariableMapOutput {
	return o
}

func (o ActionsOrganizationVariableMapOutput) MapIndex(k pulumi.StringInput) ActionsOrganizationVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ActionsOrganizationVariable {
		return vs[0].(map[string]*ActionsOrganizationVariable)[vs[1].(string)]
	}).(ActionsOrganizationVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationVariableInput)(nil)).Elem(), &ActionsOrganizationVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationVariableArrayInput)(nil)).Elem(), ActionsOrganizationVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsOrganizationVariableMapInput)(nil)).Elem(), ActionsOrganizationVariableMap{})
	pulumi.RegisterOutputType(ActionsOrganizationVariableOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationVariableArrayOutput{})
	pulumi.RegisterOutputType(ActionsOrganizationVariableMapOutput{})
}
