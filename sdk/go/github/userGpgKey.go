// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a GitHub user's GPG key resource.
//
// This resource allows you to add/remove GPG keys from your user account.
//
// ## Import
//
// GPG keys are not importable due to the fact that [API](https://developer.github.com/v3/users/gpg_keys/#gpg-keys) does not return previously uploaded GPG key.
type UserGpgKey struct {
	pulumi.CustomResourceState

	// Your public GPG key, generated in ASCII-armored format.
	// See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
	ArmoredPublicKey pulumi.StringOutput `pulumi:"armoredPublicKey"`
	Etag             pulumi.StringOutput `pulumi:"etag"`
	// The key ID of the GPG key, e.g. `3262EFF25BA0D270`
	KeyId pulumi.StringOutput `pulumi:"keyId"`
}

// NewUserGpgKey registers a new resource with the given unique name, arguments, and options.
func NewUserGpgKey(ctx *pulumi.Context,
	name string, args *UserGpgKeyArgs, opts ...pulumi.ResourceOption) (*UserGpgKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArmoredPublicKey == nil {
		return nil, errors.New("invalid value for required argument 'ArmoredPublicKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserGpgKey
	err := ctx.RegisterResource("github:index/userGpgKey:UserGpgKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGpgKey gets an existing UserGpgKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGpgKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGpgKeyState, opts ...pulumi.ResourceOption) (*UserGpgKey, error) {
	var resource UserGpgKey
	err := ctx.ReadResource("github:index/userGpgKey:UserGpgKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGpgKey resources.
type userGpgKeyState struct {
	// Your public GPG key, generated in ASCII-armored format.
	// See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
	ArmoredPublicKey *string `pulumi:"armoredPublicKey"`
	Etag             *string `pulumi:"etag"`
	// The key ID of the GPG key, e.g. `3262EFF25BA0D270`
	KeyId *string `pulumi:"keyId"`
}

type UserGpgKeyState struct {
	// Your public GPG key, generated in ASCII-armored format.
	// See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
	ArmoredPublicKey pulumi.StringPtrInput
	Etag             pulumi.StringPtrInput
	// The key ID of the GPG key, e.g. `3262EFF25BA0D270`
	KeyId pulumi.StringPtrInput
}

func (UserGpgKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGpgKeyState)(nil)).Elem()
}

type userGpgKeyArgs struct {
	// Your public GPG key, generated in ASCII-armored format.
	// See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
	ArmoredPublicKey string `pulumi:"armoredPublicKey"`
}

// The set of arguments for constructing a UserGpgKey resource.
type UserGpgKeyArgs struct {
	// Your public GPG key, generated in ASCII-armored format.
	// See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
	ArmoredPublicKey pulumi.StringInput
}

func (UserGpgKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGpgKeyArgs)(nil)).Elem()
}

type UserGpgKeyInput interface {
	pulumi.Input

	ToUserGpgKeyOutput() UserGpgKeyOutput
	ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput
}

func (*UserGpgKey) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGpgKey)(nil)).Elem()
}

func (i *UserGpgKey) ToUserGpgKeyOutput() UserGpgKeyOutput {
	return i.ToUserGpgKeyOutputWithContext(context.Background())
}

func (i *UserGpgKey) ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyOutput)
}

func (i *UserGpgKey) ToOutput(ctx context.Context) pulumix.Output[*UserGpgKey] {
	return pulumix.Output[*UserGpgKey]{
		OutputState: i.ToUserGpgKeyOutputWithContext(ctx).OutputState,
	}
}

// UserGpgKeyArrayInput is an input type that accepts UserGpgKeyArray and UserGpgKeyArrayOutput values.
// You can construct a concrete instance of `UserGpgKeyArrayInput` via:
//
//	UserGpgKeyArray{ UserGpgKeyArgs{...} }
type UserGpgKeyArrayInput interface {
	pulumi.Input

	ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput
	ToUserGpgKeyArrayOutputWithContext(context.Context) UserGpgKeyArrayOutput
}

type UserGpgKeyArray []UserGpgKeyInput

func (UserGpgKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGpgKey)(nil)).Elem()
}

func (i UserGpgKeyArray) ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput {
	return i.ToUserGpgKeyArrayOutputWithContext(context.Background())
}

func (i UserGpgKeyArray) ToUserGpgKeyArrayOutputWithContext(ctx context.Context) UserGpgKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyArrayOutput)
}

func (i UserGpgKeyArray) ToOutput(ctx context.Context) pulumix.Output[[]*UserGpgKey] {
	return pulumix.Output[[]*UserGpgKey]{
		OutputState: i.ToUserGpgKeyArrayOutputWithContext(ctx).OutputState,
	}
}

// UserGpgKeyMapInput is an input type that accepts UserGpgKeyMap and UserGpgKeyMapOutput values.
// You can construct a concrete instance of `UserGpgKeyMapInput` via:
//
//	UserGpgKeyMap{ "key": UserGpgKeyArgs{...} }
type UserGpgKeyMapInput interface {
	pulumi.Input

	ToUserGpgKeyMapOutput() UserGpgKeyMapOutput
	ToUserGpgKeyMapOutputWithContext(context.Context) UserGpgKeyMapOutput
}

type UserGpgKeyMap map[string]UserGpgKeyInput

func (UserGpgKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGpgKey)(nil)).Elem()
}

func (i UserGpgKeyMap) ToUserGpgKeyMapOutput() UserGpgKeyMapOutput {
	return i.ToUserGpgKeyMapOutputWithContext(context.Background())
}

func (i UserGpgKeyMap) ToUserGpgKeyMapOutputWithContext(ctx context.Context) UserGpgKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGpgKeyMapOutput)
}

func (i UserGpgKeyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserGpgKey] {
	return pulumix.Output[map[string]*UserGpgKey]{
		OutputState: i.ToUserGpgKeyMapOutputWithContext(ctx).OutputState,
	}
}

type UserGpgKeyOutput struct{ *pulumi.OutputState }

func (UserGpgKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyOutput) ToUserGpgKeyOutput() UserGpgKeyOutput {
	return o
}

func (o UserGpgKeyOutput) ToUserGpgKeyOutputWithContext(ctx context.Context) UserGpgKeyOutput {
	return o
}

func (o UserGpgKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*UserGpgKey] {
	return pulumix.Output[*UserGpgKey]{
		OutputState: o.OutputState,
	}
}

// Your public GPG key, generated in ASCII-armored format.
// See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
func (o UserGpgKeyOutput) ArmoredPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.ArmoredPublicKey }).(pulumi.StringOutput)
}

func (o UserGpgKeyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The key ID of the GPG key, e.g. `3262EFF25BA0D270`
func (o UserGpgKeyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGpgKey) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

type UserGpgKeyArrayOutput struct{ *pulumi.OutputState }

func (UserGpgKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyArrayOutput) ToUserGpgKeyArrayOutput() UserGpgKeyArrayOutput {
	return o
}

func (o UserGpgKeyArrayOutput) ToUserGpgKeyArrayOutputWithContext(ctx context.Context) UserGpgKeyArrayOutput {
	return o
}

func (o UserGpgKeyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*UserGpgKey] {
	return pulumix.Output[[]*UserGpgKey]{
		OutputState: o.OutputState,
	}
}

func (o UserGpgKeyArrayOutput) Index(i pulumi.IntInput) UserGpgKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserGpgKey {
		return vs[0].([]*UserGpgKey)[vs[1].(int)]
	}).(UserGpgKeyOutput)
}

type UserGpgKeyMapOutput struct{ *pulumi.OutputState }

func (UserGpgKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGpgKey)(nil)).Elem()
}

func (o UserGpgKeyMapOutput) ToUserGpgKeyMapOutput() UserGpgKeyMapOutput {
	return o
}

func (o UserGpgKeyMapOutput) ToUserGpgKeyMapOutputWithContext(ctx context.Context) UserGpgKeyMapOutput {
	return o
}

func (o UserGpgKeyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*UserGpgKey] {
	return pulumix.Output[map[string]*UserGpgKey]{
		OutputState: o.OutputState,
	}
}

func (o UserGpgKeyMapOutput) MapIndex(k pulumi.StringInput) UserGpgKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserGpgKey {
		return vs[0].(map[string]*UserGpgKey)[vs[1].(string)]
	}).(UserGpgKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyInput)(nil)).Elem(), &UserGpgKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyArrayInput)(nil)).Elem(), UserGpgKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGpgKeyMapInput)(nil)).Elem(), UserGpgKeyMap{})
	pulumi.RegisterOutputType(UserGpgKeyOutput{})
	pulumi.RegisterOutputType(UserGpgKeyArrayOutput{})
	pulumi.RegisterOutputType(UserGpgKeyMapOutput{})
}
