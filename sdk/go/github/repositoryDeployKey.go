// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a GitHub repository deploy key resource.
//
// A deploy key is an SSH key that is stored on your server and grants
// access to a single GitHub repository. This key is attached directly to the repository instead of to a personal user
// account.
//
// This resource allows you to add/remove repository deploy keys.
//
// Further documentation on GitHub repository deploy keys:
// - [About deploy keys](https://developer.github.com/guides/managing-deploy-keys/#deploy-keys)
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
// 		_, err := github.NewRepositoryDeployKey(ctx, "exampleRepositoryDeployKey", &github.RepositoryDeployKeyArgs{
// 			Key:        pulumi.String("ssh-rsa AAA..."),
// 			ReadOnly:   pulumi.Bool(false),
// 			Repository: pulumi.String("test-repo"),
// 			Title:      pulumi.String("Repository test key"),
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
// Repository deploy keys can be imported using a colon-separated pair of repository name and GitHub's key id. The latter can be obtained by GitHub's SDKs and API.
//
// ```sh
//  $ pulumi import github:index/repositoryDeployKey:RepositoryDeployKey foo test-repo:23824728
// ```
type RepositoryDeployKey struct {
	pulumi.CustomResourceState

	Etag pulumi.StringOutput `pulumi:"etag"`
	// A SSH key.
	Key pulumi.StringOutput `pulumi:"key"`
	// A boolean qualifying the key to be either read only or read/write.
	ReadOnly pulumi.BoolPtrOutput `pulumi:"readOnly"`
	// Name of the GitHub repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// A title.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewRepositoryDeployKey registers a new resource with the given unique name, arguments, and options.
func NewRepositoryDeployKey(ctx *pulumi.Context,
	name string, args *RepositoryDeployKeyArgs, opts ...pulumi.ResourceOption) (*RepositoryDeployKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource RepositoryDeployKey
	err := ctx.RegisterResource("github:index/repositoryDeployKey:RepositoryDeployKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryDeployKey gets an existing RepositoryDeployKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryDeployKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryDeployKeyState, opts ...pulumi.ResourceOption) (*RepositoryDeployKey, error) {
	var resource RepositoryDeployKey
	err := ctx.ReadResource("github:index/repositoryDeployKey:RepositoryDeployKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryDeployKey resources.
type repositoryDeployKeyState struct {
	Etag *string `pulumi:"etag"`
	// A SSH key.
	Key *string `pulumi:"key"`
	// A boolean qualifying the key to be either read only or read/write.
	ReadOnly *bool `pulumi:"readOnly"`
	// Name of the GitHub repository.
	Repository *string `pulumi:"repository"`
	// A title.
	Title *string `pulumi:"title"`
}

type RepositoryDeployKeyState struct {
	Etag pulumi.StringPtrInput
	// A SSH key.
	Key pulumi.StringPtrInput
	// A boolean qualifying the key to be either read only or read/write.
	ReadOnly pulumi.BoolPtrInput
	// Name of the GitHub repository.
	Repository pulumi.StringPtrInput
	// A title.
	Title pulumi.StringPtrInput
}

func (RepositoryDeployKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryDeployKeyState)(nil)).Elem()
}

type repositoryDeployKeyArgs struct {
	// A SSH key.
	Key string `pulumi:"key"`
	// A boolean qualifying the key to be either read only or read/write.
	ReadOnly *bool `pulumi:"readOnly"`
	// Name of the GitHub repository.
	Repository string `pulumi:"repository"`
	// A title.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a RepositoryDeployKey resource.
type RepositoryDeployKeyArgs struct {
	// A SSH key.
	Key pulumi.StringInput
	// A boolean qualifying the key to be either read only or read/write.
	ReadOnly pulumi.BoolPtrInput
	// Name of the GitHub repository.
	Repository pulumi.StringInput
	// A title.
	Title pulumi.StringInput
}

func (RepositoryDeployKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryDeployKeyArgs)(nil)).Elem()
}

type RepositoryDeployKeyInput interface {
	pulumi.Input

	ToRepositoryDeployKeyOutput() RepositoryDeployKeyOutput
	ToRepositoryDeployKeyOutputWithContext(ctx context.Context) RepositoryDeployKeyOutput
}

func (*RepositoryDeployKey) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryDeployKey)(nil)).Elem()
}

func (i *RepositoryDeployKey) ToRepositoryDeployKeyOutput() RepositoryDeployKeyOutput {
	return i.ToRepositoryDeployKeyOutputWithContext(context.Background())
}

func (i *RepositoryDeployKey) ToRepositoryDeployKeyOutputWithContext(ctx context.Context) RepositoryDeployKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryDeployKeyOutput)
}

// RepositoryDeployKeyArrayInput is an input type that accepts RepositoryDeployKeyArray and RepositoryDeployKeyArrayOutput values.
// You can construct a concrete instance of `RepositoryDeployKeyArrayInput` via:
//
//          RepositoryDeployKeyArray{ RepositoryDeployKeyArgs{...} }
type RepositoryDeployKeyArrayInput interface {
	pulumi.Input

	ToRepositoryDeployKeyArrayOutput() RepositoryDeployKeyArrayOutput
	ToRepositoryDeployKeyArrayOutputWithContext(context.Context) RepositoryDeployKeyArrayOutput
}

type RepositoryDeployKeyArray []RepositoryDeployKeyInput

func (RepositoryDeployKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryDeployKey)(nil)).Elem()
}

func (i RepositoryDeployKeyArray) ToRepositoryDeployKeyArrayOutput() RepositoryDeployKeyArrayOutput {
	return i.ToRepositoryDeployKeyArrayOutputWithContext(context.Background())
}

func (i RepositoryDeployKeyArray) ToRepositoryDeployKeyArrayOutputWithContext(ctx context.Context) RepositoryDeployKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryDeployKeyArrayOutput)
}

// RepositoryDeployKeyMapInput is an input type that accepts RepositoryDeployKeyMap and RepositoryDeployKeyMapOutput values.
// You can construct a concrete instance of `RepositoryDeployKeyMapInput` via:
//
//          RepositoryDeployKeyMap{ "key": RepositoryDeployKeyArgs{...} }
type RepositoryDeployKeyMapInput interface {
	pulumi.Input

	ToRepositoryDeployKeyMapOutput() RepositoryDeployKeyMapOutput
	ToRepositoryDeployKeyMapOutputWithContext(context.Context) RepositoryDeployKeyMapOutput
}

type RepositoryDeployKeyMap map[string]RepositoryDeployKeyInput

func (RepositoryDeployKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryDeployKey)(nil)).Elem()
}

func (i RepositoryDeployKeyMap) ToRepositoryDeployKeyMapOutput() RepositoryDeployKeyMapOutput {
	return i.ToRepositoryDeployKeyMapOutputWithContext(context.Background())
}

func (i RepositoryDeployKeyMap) ToRepositoryDeployKeyMapOutputWithContext(ctx context.Context) RepositoryDeployKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryDeployKeyMapOutput)
}

type RepositoryDeployKeyOutput struct{ *pulumi.OutputState }

func (RepositoryDeployKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryDeployKey)(nil)).Elem()
}

func (o RepositoryDeployKeyOutput) ToRepositoryDeployKeyOutput() RepositoryDeployKeyOutput {
	return o
}

func (o RepositoryDeployKeyOutput) ToRepositoryDeployKeyOutputWithContext(ctx context.Context) RepositoryDeployKeyOutput {
	return o
}

type RepositoryDeployKeyArrayOutput struct{ *pulumi.OutputState }

func (RepositoryDeployKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryDeployKey)(nil)).Elem()
}

func (o RepositoryDeployKeyArrayOutput) ToRepositoryDeployKeyArrayOutput() RepositoryDeployKeyArrayOutput {
	return o
}

func (o RepositoryDeployKeyArrayOutput) ToRepositoryDeployKeyArrayOutputWithContext(ctx context.Context) RepositoryDeployKeyArrayOutput {
	return o
}

func (o RepositoryDeployKeyArrayOutput) Index(i pulumi.IntInput) RepositoryDeployKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryDeployKey {
		return vs[0].([]*RepositoryDeployKey)[vs[1].(int)]
	}).(RepositoryDeployKeyOutput)
}

type RepositoryDeployKeyMapOutput struct{ *pulumi.OutputState }

func (RepositoryDeployKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryDeployKey)(nil)).Elem()
}

func (o RepositoryDeployKeyMapOutput) ToRepositoryDeployKeyMapOutput() RepositoryDeployKeyMapOutput {
	return o
}

func (o RepositoryDeployKeyMapOutput) ToRepositoryDeployKeyMapOutputWithContext(ctx context.Context) RepositoryDeployKeyMapOutput {
	return o
}

func (o RepositoryDeployKeyMapOutput) MapIndex(k pulumi.StringInput) RepositoryDeployKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryDeployKey {
		return vs[0].(map[string]*RepositoryDeployKey)[vs[1].(string)]
	}).(RepositoryDeployKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryDeployKeyInput)(nil)).Elem(), &RepositoryDeployKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryDeployKeyArrayInput)(nil)).Elem(), RepositoryDeployKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryDeployKeyMapInput)(nil)).Elem(), RepositoryDeployKeyMap{})
	pulumi.RegisterOutputType(RepositoryDeployKeyOutput{})
	pulumi.RegisterOutputType(RepositoryDeployKeyArrayOutput{})
	pulumi.RegisterOutputType(RepositoryDeployKeyMapOutput{})
}
