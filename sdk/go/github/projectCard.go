// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectCard struct {
	pulumi.CustomResourceState

	CardId      pulumi.IntOutput       `pulumi:"cardId"`
	ColumnId    pulumi.StringOutput    `pulumi:"columnId"`
	ContentId   pulumi.IntPtrOutput    `pulumi:"contentId"`
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	Etag        pulumi.StringOutput    `pulumi:"etag"`
	Note        pulumi.StringPtrOutput `pulumi:"note"`
}

// NewProjectCard registers a new resource with the given unique name, arguments, and options.
func NewProjectCard(ctx *pulumi.Context,
	name string, args *ProjectCardArgs, opts ...pulumi.ResourceOption) (*ProjectCard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnId == nil {
		return nil, errors.New("invalid value for required argument 'ColumnId'")
	}
	var resource ProjectCard
	err := ctx.RegisterResource("github:index/projectCard:ProjectCard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectCard gets an existing ProjectCard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectCard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectCardState, opts ...pulumi.ResourceOption) (*ProjectCard, error) {
	var resource ProjectCard
	err := ctx.ReadResource("github:index/projectCard:ProjectCard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectCard resources.
type projectCardState struct {
	CardId      *int    `pulumi:"cardId"`
	ColumnId    *string `pulumi:"columnId"`
	ContentId   *int    `pulumi:"contentId"`
	ContentType *string `pulumi:"contentType"`
	Etag        *string `pulumi:"etag"`
	Note        *string `pulumi:"note"`
}

type ProjectCardState struct {
	CardId      pulumi.IntPtrInput
	ColumnId    pulumi.StringPtrInput
	ContentId   pulumi.IntPtrInput
	ContentType pulumi.StringPtrInput
	Etag        pulumi.StringPtrInput
	Note        pulumi.StringPtrInput
}

func (ProjectCardState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectCardState)(nil)).Elem()
}

type projectCardArgs struct {
	ColumnId    string  `pulumi:"columnId"`
	ContentId   *int    `pulumi:"contentId"`
	ContentType *string `pulumi:"contentType"`
	Note        *string `pulumi:"note"`
}

// The set of arguments for constructing a ProjectCard resource.
type ProjectCardArgs struct {
	ColumnId    pulumi.StringInput
	ContentId   pulumi.IntPtrInput
	ContentType pulumi.StringPtrInput
	Note        pulumi.StringPtrInput
}

func (ProjectCardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectCardArgs)(nil)).Elem()
}

type ProjectCardInput interface {
	pulumi.Input

	ToProjectCardOutput() ProjectCardOutput
	ToProjectCardOutputWithContext(ctx context.Context) ProjectCardOutput
}

func (*ProjectCard) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectCard)(nil)).Elem()
}

func (i *ProjectCard) ToProjectCardOutput() ProjectCardOutput {
	return i.ToProjectCardOutputWithContext(context.Background())
}

func (i *ProjectCard) ToProjectCardOutputWithContext(ctx context.Context) ProjectCardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectCardOutput)
}

// ProjectCardArrayInput is an input type that accepts ProjectCardArray and ProjectCardArrayOutput values.
// You can construct a concrete instance of `ProjectCardArrayInput` via:
//
//	ProjectCardArray{ ProjectCardArgs{...} }
type ProjectCardArrayInput interface {
	pulumi.Input

	ToProjectCardArrayOutput() ProjectCardArrayOutput
	ToProjectCardArrayOutputWithContext(context.Context) ProjectCardArrayOutput
}

type ProjectCardArray []ProjectCardInput

func (ProjectCardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectCard)(nil)).Elem()
}

func (i ProjectCardArray) ToProjectCardArrayOutput() ProjectCardArrayOutput {
	return i.ToProjectCardArrayOutputWithContext(context.Background())
}

func (i ProjectCardArray) ToProjectCardArrayOutputWithContext(ctx context.Context) ProjectCardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectCardArrayOutput)
}

// ProjectCardMapInput is an input type that accepts ProjectCardMap and ProjectCardMapOutput values.
// You can construct a concrete instance of `ProjectCardMapInput` via:
//
//	ProjectCardMap{ "key": ProjectCardArgs{...} }
type ProjectCardMapInput interface {
	pulumi.Input

	ToProjectCardMapOutput() ProjectCardMapOutput
	ToProjectCardMapOutputWithContext(context.Context) ProjectCardMapOutput
}

type ProjectCardMap map[string]ProjectCardInput

func (ProjectCardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectCard)(nil)).Elem()
}

func (i ProjectCardMap) ToProjectCardMapOutput() ProjectCardMapOutput {
	return i.ToProjectCardMapOutputWithContext(context.Background())
}

func (i ProjectCardMap) ToProjectCardMapOutputWithContext(ctx context.Context) ProjectCardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectCardMapOutput)
}

type ProjectCardOutput struct{ *pulumi.OutputState }

func (ProjectCardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectCard)(nil)).Elem()
}

func (o ProjectCardOutput) ToProjectCardOutput() ProjectCardOutput {
	return o
}

func (o ProjectCardOutput) ToProjectCardOutputWithContext(ctx context.Context) ProjectCardOutput {
	return o
}

func (o ProjectCardOutput) CardId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectCard) pulumi.IntOutput { return v.CardId }).(pulumi.IntOutput)
}

func (o ProjectCardOutput) ColumnId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCard) pulumi.StringOutput { return v.ColumnId }).(pulumi.StringOutput)
}

func (o ProjectCardOutput) ContentId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectCard) pulumi.IntPtrOutput { return v.ContentId }).(pulumi.IntPtrOutput)
}

func (o ProjectCardOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCard) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o ProjectCardOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectCard) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ProjectCardOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectCard) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

type ProjectCardArrayOutput struct{ *pulumi.OutputState }

func (ProjectCardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectCard)(nil)).Elem()
}

func (o ProjectCardArrayOutput) ToProjectCardArrayOutput() ProjectCardArrayOutput {
	return o
}

func (o ProjectCardArrayOutput) ToProjectCardArrayOutputWithContext(ctx context.Context) ProjectCardArrayOutput {
	return o
}

func (o ProjectCardArrayOutput) Index(i pulumi.IntInput) ProjectCardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectCard {
		return vs[0].([]*ProjectCard)[vs[1].(int)]
	}).(ProjectCardOutput)
}

type ProjectCardMapOutput struct{ *pulumi.OutputState }

func (ProjectCardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectCard)(nil)).Elem()
}

func (o ProjectCardMapOutput) ToProjectCardMapOutput() ProjectCardMapOutput {
	return o
}

func (o ProjectCardMapOutput) ToProjectCardMapOutputWithContext(ctx context.Context) ProjectCardMapOutput {
	return o
}

func (o ProjectCardMapOutput) MapIndex(k pulumi.StringInput) ProjectCardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectCard {
		return vs[0].(map[string]*ProjectCard)[vs[1].(string)]
	}).(ProjectCardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectCardInput)(nil)).Elem(), &ProjectCard{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectCardArrayInput)(nil)).Elem(), ProjectCardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectCardMapInput)(nil)).Elem(), ProjectCardMap{})
	pulumi.RegisterOutputType(ProjectCardOutput{})
	pulumi.RegisterOutputType(ProjectCardArrayOutput{})
	pulumi.RegisterOutputType(ProjectCardMapOutput{})
}
