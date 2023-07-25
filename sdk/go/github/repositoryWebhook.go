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

// This resource allows you to create and manage webhooks for repositories within your
// GitHub organization or personal account.
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
//			repo, err := github.NewRepository(ctx, "repo", &github.RepositoryArgs{
//				Description: pulumi.String("Terraform acceptance tests"),
//				HomepageUrl: pulumi.String("http://example.com/"),
//				Visibility:  pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryWebhook(ctx, "foo", &github.RepositoryWebhookArgs{
//				Repository: repo.Name,
//				Configuration: &github.RepositoryWebhookConfigurationArgs{
//					Url:         pulumi.String("https://google.de/"),
//					ContentType: pulumi.String("form"),
//					InsecureSsl: pulumi.Bool(false),
//				},
//				Active: pulumi.Bool(false),
//				Events: pulumi.StringArray{
//					pulumi.String("issues"),
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
// Repository webhooks can be imported using the `name` of the repository, combined with the `id` of the webhook, separated by a `/` character. The `id` of the webhook can be found in the URL of the webhook. For example`"https://github.com/foo-org/foo-repo/settings/hooks/14711452"`. Importing uses the name of the repository, as well as the ID of the webhook, e.g.
//
// ```sh
//
//	$ pulumi import github:index/repositoryWebhook:RepositoryWebhook terraform terraform/11235813
//
// ```
//
//	If secret is populated in the webhook's configuration, the value will be imported as "********".
type RepositoryWebhook struct {
	pulumi.CustomResourceState

	// Indicate if the webhook should receive events. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Configuration block for the webhook. Detailed below.
	Configuration RepositoryWebhookConfigurationPtrOutput `pulumi:"configuration"`
	Etag          pulumi.StringOutput                     `pulumi:"etag"`
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
	Events pulumi.StringArrayOutput `pulumi:"events"`
	// The repository of the webhook.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The URL of the webhook.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewRepositoryWebhook registers a new resource with the given unique name, arguments, and options.
func NewRepositoryWebhook(ctx *pulumi.Context,
	name string, args *RepositoryWebhookArgs, opts ...pulumi.ResourceOption) (*RepositoryWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Events == nil {
		return nil, errors.New("invalid value for required argument 'Events'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryWebhook
	err := ctx.RegisterResource("github:index/repositoryWebhook:RepositoryWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryWebhook gets an existing RepositoryWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryWebhookState, opts ...pulumi.ResourceOption) (*RepositoryWebhook, error) {
	var resource RepositoryWebhook
	err := ctx.ReadResource("github:index/repositoryWebhook:RepositoryWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryWebhook resources.
type repositoryWebhookState struct {
	// Indicate if the webhook should receive events. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Configuration block for the webhook. Detailed below.
	Configuration *RepositoryWebhookConfiguration `pulumi:"configuration"`
	Etag          *string                         `pulumi:"etag"`
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
	Events []string `pulumi:"events"`
	// The repository of the webhook.
	Repository *string `pulumi:"repository"`
	// The URL of the webhook.
	Url *string `pulumi:"url"`
}

type RepositoryWebhookState struct {
	// Indicate if the webhook should receive events. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Configuration block for the webhook. Detailed below.
	Configuration RepositoryWebhookConfigurationPtrInput
	Etag          pulumi.StringPtrInput
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
	Events pulumi.StringArrayInput
	// The repository of the webhook.
	Repository pulumi.StringPtrInput
	// The URL of the webhook.
	Url pulumi.StringPtrInput
}

func (RepositoryWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryWebhookState)(nil)).Elem()
}

type repositoryWebhookArgs struct {
	// Indicate if the webhook should receive events. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Configuration block for the webhook. Detailed below.
	Configuration *RepositoryWebhookConfiguration `pulumi:"configuration"`
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
	Events []string `pulumi:"events"`
	// The repository of the webhook.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryWebhook resource.
type RepositoryWebhookArgs struct {
	// Indicate if the webhook should receive events. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Configuration block for the webhook. Detailed below.
	Configuration RepositoryWebhookConfigurationPtrInput
	// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
	Events pulumi.StringArrayInput
	// The repository of the webhook.
	Repository pulumi.StringInput
}

func (RepositoryWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryWebhookArgs)(nil)).Elem()
}

type RepositoryWebhookInput interface {
	pulumi.Input

	ToRepositoryWebhookOutput() RepositoryWebhookOutput
	ToRepositoryWebhookOutputWithContext(ctx context.Context) RepositoryWebhookOutput
}

func (*RepositoryWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryWebhook)(nil)).Elem()
}

func (i *RepositoryWebhook) ToRepositoryWebhookOutput() RepositoryWebhookOutput {
	return i.ToRepositoryWebhookOutputWithContext(context.Background())
}

func (i *RepositoryWebhook) ToRepositoryWebhookOutputWithContext(ctx context.Context) RepositoryWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWebhookOutput)
}

// RepositoryWebhookArrayInput is an input type that accepts RepositoryWebhookArray and RepositoryWebhookArrayOutput values.
// You can construct a concrete instance of `RepositoryWebhookArrayInput` via:
//
//	RepositoryWebhookArray{ RepositoryWebhookArgs{...} }
type RepositoryWebhookArrayInput interface {
	pulumi.Input

	ToRepositoryWebhookArrayOutput() RepositoryWebhookArrayOutput
	ToRepositoryWebhookArrayOutputWithContext(context.Context) RepositoryWebhookArrayOutput
}

type RepositoryWebhookArray []RepositoryWebhookInput

func (RepositoryWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryWebhook)(nil)).Elem()
}

func (i RepositoryWebhookArray) ToRepositoryWebhookArrayOutput() RepositoryWebhookArrayOutput {
	return i.ToRepositoryWebhookArrayOutputWithContext(context.Background())
}

func (i RepositoryWebhookArray) ToRepositoryWebhookArrayOutputWithContext(ctx context.Context) RepositoryWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWebhookArrayOutput)
}

// RepositoryWebhookMapInput is an input type that accepts RepositoryWebhookMap and RepositoryWebhookMapOutput values.
// You can construct a concrete instance of `RepositoryWebhookMapInput` via:
//
//	RepositoryWebhookMap{ "key": RepositoryWebhookArgs{...} }
type RepositoryWebhookMapInput interface {
	pulumi.Input

	ToRepositoryWebhookMapOutput() RepositoryWebhookMapOutput
	ToRepositoryWebhookMapOutputWithContext(context.Context) RepositoryWebhookMapOutput
}

type RepositoryWebhookMap map[string]RepositoryWebhookInput

func (RepositoryWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryWebhook)(nil)).Elem()
}

func (i RepositoryWebhookMap) ToRepositoryWebhookMapOutput() RepositoryWebhookMapOutput {
	return i.ToRepositoryWebhookMapOutputWithContext(context.Background())
}

func (i RepositoryWebhookMap) ToRepositoryWebhookMapOutputWithContext(ctx context.Context) RepositoryWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWebhookMapOutput)
}

type RepositoryWebhookOutput struct{ *pulumi.OutputState }

func (RepositoryWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryWebhook)(nil)).Elem()
}

func (o RepositoryWebhookOutput) ToRepositoryWebhookOutput() RepositoryWebhookOutput {
	return o
}

func (o RepositoryWebhookOutput) ToRepositoryWebhookOutputWithContext(ctx context.Context) RepositoryWebhookOutput {
	return o
}

// Indicate if the webhook should receive events. Defaults to `true`.
func (o RepositoryWebhookOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// Configuration block for the webhook. Detailed below.
func (o RepositoryWebhookOutput) Configuration() RepositoryWebhookConfigurationPtrOutput {
	return o.ApplyT(func(v *RepositoryWebhook) RepositoryWebhookConfigurationPtrOutput { return v.Configuration }).(RepositoryWebhookConfigurationPtrOutput)
}

func (o RepositoryWebhookOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
func (o RepositoryWebhookOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringArrayOutput { return v.Events }).(pulumi.StringArrayOutput)
}

// The repository of the webhook.
func (o RepositoryWebhookOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The URL of the webhook.
func (o RepositoryWebhookOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWebhook) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type RepositoryWebhookArrayOutput struct{ *pulumi.OutputState }

func (RepositoryWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryWebhook)(nil)).Elem()
}

func (o RepositoryWebhookArrayOutput) ToRepositoryWebhookArrayOutput() RepositoryWebhookArrayOutput {
	return o
}

func (o RepositoryWebhookArrayOutput) ToRepositoryWebhookArrayOutputWithContext(ctx context.Context) RepositoryWebhookArrayOutput {
	return o
}

func (o RepositoryWebhookArrayOutput) Index(i pulumi.IntInput) RepositoryWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryWebhook {
		return vs[0].([]*RepositoryWebhook)[vs[1].(int)]
	}).(RepositoryWebhookOutput)
}

type RepositoryWebhookMapOutput struct{ *pulumi.OutputState }

func (RepositoryWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryWebhook)(nil)).Elem()
}

func (o RepositoryWebhookMapOutput) ToRepositoryWebhookMapOutput() RepositoryWebhookMapOutput {
	return o
}

func (o RepositoryWebhookMapOutput) ToRepositoryWebhookMapOutputWithContext(ctx context.Context) RepositoryWebhookMapOutput {
	return o
}

func (o RepositoryWebhookMapOutput) MapIndex(k pulumi.StringInput) RepositoryWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryWebhook {
		return vs[0].(map[string]*RepositoryWebhook)[vs[1].(string)]
	}).(RepositoryWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWebhookInput)(nil)).Elem(), &RepositoryWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWebhookArrayInput)(nil)).Elem(), RepositoryWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWebhookMapInput)(nil)).Elem(), RepositoryWebhookMap{})
	pulumi.RegisterOutputType(RepositoryWebhookOutput{})
	pulumi.RegisterOutputType(RepositoryWebhookArrayOutput{})
	pulumi.RegisterOutputType(RepositoryWebhookMapOutput{})
}
