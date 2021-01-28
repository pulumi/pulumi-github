// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a GitHub membership resource.
//
// This resource allows you to add/remove users from your organization. When applied,
// an invitation will be sent to the user to become part of the organization. When
// destroyed, either the invitation will be cancelled or the user will be removed.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-github/sdk/v2/go/github/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := github.NewMembership(ctx, "membershipForSomeUser", &github.MembershipArgs{
// 			Role:     pulumi.String("member"),
// 			Username: pulumi.String("SomeUser"),
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
// GitHub Membership can be imported using an ID made up of `organization:username`, e.g.
//
// ```sh
//  $ pulumi import github:index/membership:Membership member hashicorp:someuser
// ```
type Membership struct {
	pulumi.CustomResourceState

	Etag pulumi.StringOutput `pulumi:"etag"`
	// The role of the user within the organization.
	// Must be one of `member` or `admin`. Defaults to `member`.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// The user to add to the organization.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewMembership registers a new resource with the given unique name, arguments, and options.
func NewMembership(ctx *pulumi.Context,
	name string, args *MembershipArgs, opts ...pulumi.ResourceOption) (*Membership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource Membership
	err := ctx.RegisterResource("github:index/membership:Membership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembership gets an existing Membership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipState, opts ...pulumi.ResourceOption) (*Membership, error) {
	var resource Membership
	err := ctx.ReadResource("github:index/membership:Membership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Membership resources.
type membershipState struct {
	Etag *string `pulumi:"etag"`
	// The role of the user within the organization.
	// Must be one of `member` or `admin`. Defaults to `member`.
	Role *string `pulumi:"role"`
	// The user to add to the organization.
	Username *string `pulumi:"username"`
}

type MembershipState struct {
	Etag pulumi.StringPtrInput
	// The role of the user within the organization.
	// Must be one of `member` or `admin`. Defaults to `member`.
	Role pulumi.StringPtrInput
	// The user to add to the organization.
	Username pulumi.StringPtrInput
}

func (MembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipState)(nil)).Elem()
}

type membershipArgs struct {
	// The role of the user within the organization.
	// Must be one of `member` or `admin`. Defaults to `member`.
	Role *string `pulumi:"role"`
	// The user to add to the organization.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Membership resource.
type MembershipArgs struct {
	// The role of the user within the organization.
	// Must be one of `member` or `admin`. Defaults to `member`.
	Role pulumi.StringPtrInput
	// The user to add to the organization.
	Username pulumi.StringInput
}

func (MembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipArgs)(nil)).Elem()
}

type MembershipInput interface {
	pulumi.Input

	ToMembershipOutput() MembershipOutput
	ToMembershipOutputWithContext(ctx context.Context) MembershipOutput
}

func (*Membership) ElementType() reflect.Type {
	return reflect.TypeOf((*Membership)(nil))
}

func (i *Membership) ToMembershipOutput() MembershipOutput {
	return i.ToMembershipOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipOutput)
}

type MembershipOutput struct {
	*pulumi.OutputState
}

func (MembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Membership)(nil))
}

func (o MembershipOutput) ToMembershipOutput() MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MembershipOutput{})
}
