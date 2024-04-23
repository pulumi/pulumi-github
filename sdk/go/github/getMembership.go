// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to find out if a user is a member of your organization, as well
// as what role they have within it.
// If the user's membership in the organization is pending their acceptance of an invite,
// the role they would have once they accept will be returned.
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
//			_, err := github.LookupMembership(ctx, &github.LookupMembershipArgs{
//				Username: "SomeUser",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupMembership(ctx *pulumi.Context, args *LookupMembershipArgs, opts ...pulumi.InvokeOption) (*LookupMembershipResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMembershipResult
	err := ctx.Invoke("github:index/getMembership:getMembership", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMembership.
type LookupMembershipArgs struct {
	// The organization to check for the above username.
	Organization *string `pulumi:"organization"`
	// The username to lookup in the organization.
	Username string `pulumi:"username"`
}

// A collection of values returned by getMembership.
type LookupMembershipResult struct {
	// An etag representing the membership object.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	Organization *string `pulumi:"organization"`
	// `admin` or `member` -- the role the user has within the organization.
	Role string `pulumi:"role"`
	// `active` or `pending` -- the state of membership within the organization.  `active` if the member has accepted the invite, or `pending` if the invite is still pending.
	State string `pulumi:"state"`
	// The username.
	Username string `pulumi:"username"`
}

func LookupMembershipOutput(ctx *pulumi.Context, args LookupMembershipOutputArgs, opts ...pulumi.InvokeOption) LookupMembershipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMembershipResult, error) {
			args := v.(LookupMembershipArgs)
			r, err := LookupMembership(ctx, &args, opts...)
			var s LookupMembershipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMembershipResultOutput)
}

// A collection of arguments for invoking getMembership.
type LookupMembershipOutputArgs struct {
	// The organization to check for the above username.
	Organization pulumi.StringPtrInput `pulumi:"organization"`
	// The username to lookup in the organization.
	Username pulumi.StringInput `pulumi:"username"`
}

func (LookupMembershipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMembershipArgs)(nil)).Elem()
}

// A collection of values returned by getMembership.
type LookupMembershipResultOutput struct{ *pulumi.OutputState }

func (LookupMembershipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMembershipResult)(nil)).Elem()
}

func (o LookupMembershipResultOutput) ToLookupMembershipResultOutput() LookupMembershipResultOutput {
	return o
}

func (o LookupMembershipResultOutput) ToLookupMembershipResultOutputWithContext(ctx context.Context) LookupMembershipResultOutput {
	return o
}

// An etag representing the membership object.
func (o LookupMembershipResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMembershipResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMembershipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMembershipResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMembershipResultOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMembershipResult) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// `admin` or `member` -- the role the user has within the organization.
func (o LookupMembershipResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMembershipResult) string { return v.Role }).(pulumi.StringOutput)
}

// `active` or `pending` -- the state of membership within the organization.  `active` if the member has accepted the invite, or `pending` if the invite is still pending.
func (o LookupMembershipResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMembershipResult) string { return v.State }).(pulumi.StringOutput)
}

// The username.
func (o LookupMembershipResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMembershipResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMembershipResultOutput{})
}
