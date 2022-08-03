// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource allows you to create and manage webhooks for GitHub organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.NewOrganizationWebhook(ctx, "foo", &github.OrganizationWebhookArgs{
// 			Active: pulumi.Bool(false),
// 			Configuration: &OrganizationWebhookConfigurationArgs{
// 				ContentType: pulumi.String("form"),
// 				InsecureSsl: pulumi.Bool(false),
// 				Url:         pulumi.String("https://google.de/"),
// 			},
// 			Events: pulumi.StringArray{
// 				pulumi.String("issues"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Organization webhooks can be imported using the `id` of the webhook. The `id` of the webhook can be found in the URL of the webhook. For example, `"https://github.com/organizations/foo-org/settings/hooks/123456789"`.
//
// ```sh
//  $ pulumi import github:index/organizationWebhook:OrganizationWebhook terraform 123456789
// ```
//
//  If secret is populated in the webhook's configuration, the value will be imported as "********".
type OrganizationWebhook struct {
	pulumi.CustomResourceState

	// Indicate of the webhook should receive events. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
	Configuration OrganizationWebhookConfigurationPtrOutput `pulumi:"configuration"`
	Etag          pulumi.StringOutput                       `pulumi:"etag"`
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
	Events pulumi.StringArrayOutput `pulumi:"events"`
	// URL of the webhook
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewOrganizationWebhook registers a new resource with the given unique name, arguments, and options.
func NewOrganizationWebhook(ctx *pulumi.Context,
	name string, args *OrganizationWebhookArgs, opts ...pulumi.ResourceOption) (*OrganizationWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Events == nil {
		return nil, errors.New("invalid value for required argument 'Events'")
	}
	var resource OrganizationWebhook
	err := ctx.RegisterResource("github:index/organizationWebhook:OrganizationWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationWebhook gets an existing OrganizationWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationWebhookState, opts ...pulumi.ResourceOption) (*OrganizationWebhook, error) {
	var resource OrganizationWebhook
	err := ctx.ReadResource("github:index/organizationWebhook:OrganizationWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationWebhook resources.
type organizationWebhookState struct {
	// Indicate of the webhook should receive events. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
	Configuration *OrganizationWebhookConfiguration `pulumi:"configuration"`
	Etag          *string                           `pulumi:"etag"`
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
	Events []string `pulumi:"events"`
	// URL of the webhook
	Url *string `pulumi:"url"`
}

type OrganizationWebhookState struct {
	// Indicate of the webhook should receive events. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
	Configuration OrganizationWebhookConfigurationPtrInput
	Etag          pulumi.StringPtrInput
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
	Events pulumi.StringArrayInput
	// URL of the webhook
	Url pulumi.StringPtrInput
}

func (OrganizationWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationWebhookState)(nil)).Elem()
}

type organizationWebhookArgs struct {
	// Indicate of the webhook should receive events. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
	Configuration *OrganizationWebhookConfiguration `pulumi:"configuration"`
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
	Events []string `pulumi:"events"`
}

// The set of arguments for constructing a OrganizationWebhook resource.
type OrganizationWebhookArgs struct {
	// Indicate of the webhook should receive events. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
	Configuration OrganizationWebhookConfigurationPtrInput
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
	Events pulumi.StringArrayInput
}

func (OrganizationWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationWebhookArgs)(nil)).Elem()
}

type OrganizationWebhookInput interface {
	pulumi.Input

	ToOrganizationWebhookOutput() OrganizationWebhookOutput
	ToOrganizationWebhookOutputWithContext(ctx context.Context) OrganizationWebhookOutput
}

func (*OrganizationWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationWebhook)(nil)).Elem()
}

func (i *OrganizationWebhook) ToOrganizationWebhookOutput() OrganizationWebhookOutput {
	return i.ToOrganizationWebhookOutputWithContext(context.Background())
}

func (i *OrganizationWebhook) ToOrganizationWebhookOutputWithContext(ctx context.Context) OrganizationWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationWebhookOutput)
}

// OrganizationWebhookArrayInput is an input type that accepts OrganizationWebhookArray and OrganizationWebhookArrayOutput values.
// You can construct a concrete instance of `OrganizationWebhookArrayInput` via:
//
//          OrganizationWebhookArray{ OrganizationWebhookArgs{...} }
type OrganizationWebhookArrayInput interface {
	pulumi.Input

	ToOrganizationWebhookArrayOutput() OrganizationWebhookArrayOutput
	ToOrganizationWebhookArrayOutputWithContext(context.Context) OrganizationWebhookArrayOutput
}

type OrganizationWebhookArray []OrganizationWebhookInput

func (OrganizationWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationWebhook)(nil)).Elem()
}

func (i OrganizationWebhookArray) ToOrganizationWebhookArrayOutput() OrganizationWebhookArrayOutput {
	return i.ToOrganizationWebhookArrayOutputWithContext(context.Background())
}

func (i OrganizationWebhookArray) ToOrganizationWebhookArrayOutputWithContext(ctx context.Context) OrganizationWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationWebhookArrayOutput)
}

// OrganizationWebhookMapInput is an input type that accepts OrganizationWebhookMap and OrganizationWebhookMapOutput values.
// You can construct a concrete instance of `OrganizationWebhookMapInput` via:
//
//          OrganizationWebhookMap{ "key": OrganizationWebhookArgs{...} }
type OrganizationWebhookMapInput interface {
	pulumi.Input

	ToOrganizationWebhookMapOutput() OrganizationWebhookMapOutput
	ToOrganizationWebhookMapOutputWithContext(context.Context) OrganizationWebhookMapOutput
}

type OrganizationWebhookMap map[string]OrganizationWebhookInput

func (OrganizationWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationWebhook)(nil)).Elem()
}

func (i OrganizationWebhookMap) ToOrganizationWebhookMapOutput() OrganizationWebhookMapOutput {
	return i.ToOrganizationWebhookMapOutputWithContext(context.Background())
}

func (i OrganizationWebhookMap) ToOrganizationWebhookMapOutputWithContext(ctx context.Context) OrganizationWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationWebhookMapOutput)
}

type OrganizationWebhookOutput struct{ *pulumi.OutputState }

func (OrganizationWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationWebhook)(nil)).Elem()
}

func (o OrganizationWebhookOutput) ToOrganizationWebhookOutput() OrganizationWebhookOutput {
	return o
}

func (o OrganizationWebhookOutput) ToOrganizationWebhookOutputWithContext(ctx context.Context) OrganizationWebhookOutput {
	return o
}

// Indicate of the webhook should receive events. Defaults to `true`.
func (o OrganizationWebhookOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrganizationWebhook) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// key/value pair of configuration for this webhook. Available keys are `url`, `contentType`, `secret` and `insecureSsl`.
func (o OrganizationWebhookOutput) Configuration() OrganizationWebhookConfigurationPtrOutput {
	return o.ApplyT(func(v *OrganizationWebhook) OrganizationWebhookConfigurationPtrOutput { return v.Configuration }).(OrganizationWebhookConfigurationPtrOutput)
}

func (o OrganizationWebhookOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationWebhook) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
func (o OrganizationWebhookOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationWebhook) pulumi.StringArrayOutput { return v.Events }).(pulumi.StringArrayOutput)
}

// URL of the webhook
func (o OrganizationWebhookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationWebhook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type OrganizationWebhookArrayOutput struct{ *pulumi.OutputState }

func (OrganizationWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationWebhook)(nil)).Elem()
}

func (o OrganizationWebhookArrayOutput) ToOrganizationWebhookArrayOutput() OrganizationWebhookArrayOutput {
	return o
}

func (o OrganizationWebhookArrayOutput) ToOrganizationWebhookArrayOutputWithContext(ctx context.Context) OrganizationWebhookArrayOutput {
	return o
}

func (o OrganizationWebhookArrayOutput) Index(i pulumi.IntInput) OrganizationWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationWebhook {
		return vs[0].([]*OrganizationWebhook)[vs[1].(int)]
	}).(OrganizationWebhookOutput)
}

type OrganizationWebhookMapOutput struct{ *pulumi.OutputState }

func (OrganizationWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationWebhook)(nil)).Elem()
}

func (o OrganizationWebhookMapOutput) ToOrganizationWebhookMapOutput() OrganizationWebhookMapOutput {
	return o
}

func (o OrganizationWebhookMapOutput) ToOrganizationWebhookMapOutputWithContext(ctx context.Context) OrganizationWebhookMapOutput {
	return o
}

func (o OrganizationWebhookMapOutput) MapIndex(k pulumi.StringInput) OrganizationWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationWebhook {
		return vs[0].(map[string]*OrganizationWebhook)[vs[1].(string)]
	}).(OrganizationWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationWebhookInput)(nil)).Elem(), &OrganizationWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationWebhookArrayInput)(nil)).Elem(), OrganizationWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationWebhookMapInput)(nil)).Elem(), OrganizationWebhookMap{})
	pulumi.RegisterOutputType(OrganizationWebhookOutput{})
	pulumi.RegisterOutputType(OrganizationWebhookArrayOutput{})
	pulumi.RegisterOutputType(OrganizationWebhookMapOutput{})
}
