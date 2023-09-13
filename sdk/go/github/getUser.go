// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve information about a GitHub user.
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
//			_, err := github.GetUser(ctx, &github.GetUserArgs{
//				Username: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			current, err := github.GetUser(ctx, &github.GetUserArgs{
//				Username: "",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("currentGithubLogin", current.Login)
//			return nil
//		})
//	}
//
// ```
func GetUser(ctx *pulumi.Context, args *GetUserArgs, opts ...pulumi.InvokeOption) (*GetUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserResult
	err := ctx.Invoke("github:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// The username. Use an empty string `""` to retrieve information about the currently authenticated user.
	Username string `pulumi:"username"`
}

// A collection of values returned by getUser.
type GetUserResult struct {
	// the user's avatar URL.
	AvatarUrl string `pulumi:"avatarUrl"`
	// the user's bio.
	Bio string `pulumi:"bio"`
	// the user's blog location.
	Blog string `pulumi:"blog"`
	// the user's company name.
	Company string `pulumi:"company"`
	// the creation date.
	CreatedAt string `pulumi:"createdAt"`
	// the user's email.
	Email string `pulumi:"email"`
	// the number of followers.
	Followers int `pulumi:"followers"`
	// the number of following users.
	Following int `pulumi:"following"`
	// list of user's GPG keys.
	GpgKeys []string `pulumi:"gpgKeys"`
	// the user's gravatar ID.
	GravatarId string `pulumi:"gravatarId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// the user's location.
	Location string `pulumi:"location"`
	// the user's login.
	Login string `pulumi:"login"`
	// the user's full name.
	Name string `pulumi:"name"`
	// the Node ID of the user.
	NodeId string `pulumi:"nodeId"`
	// the number of public gists.
	PublicGists int `pulumi:"publicGists"`
	// the number of public repositories.
	PublicRepos int `pulumi:"publicRepos"`
	// whether the user is a GitHub admin.
	SiteAdmin bool `pulumi:"siteAdmin"`
	// list of user's SSH keys.
	SshKeys []string `pulumi:"sshKeys"`
	// the suspended date if the user is suspended.
	SuspendedAt string `pulumi:"suspendedAt"`
	// the update date.
	UpdatedAt string `pulumi:"updatedAt"`
	Username  string `pulumi:"username"`
}

func GetUserOutput(ctx *pulumi.Context, args GetUserOutputArgs, opts ...pulumi.InvokeOption) GetUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserResult, error) {
			args := v.(GetUserArgs)
			r, err := GetUser(ctx, &args, opts...)
			var s GetUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserResultOutput)
}

// A collection of arguments for invoking getUser.
type GetUserOutputArgs struct {
	// The username. Use an empty string `""` to retrieve information about the currently authenticated user.
	Username pulumi.StringInput `pulumi:"username"`
}

func (GetUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type GetUserResultOutput struct{ *pulumi.OutputState }

func (GetUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserResult)(nil)).Elem()
}

func (o GetUserResultOutput) ToGetUserResultOutput() GetUserResultOutput {
	return o
}

func (o GetUserResultOutput) ToGetUserResultOutputWithContext(ctx context.Context) GetUserResultOutput {
	return o
}

func (o GetUserResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetUserResult] {
	return pulumix.Output[GetUserResult]{
		OutputState: o.OutputState,
	}
}

// the user's avatar URL.
func (o GetUserResultOutput) AvatarUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.AvatarUrl }).(pulumi.StringOutput)
}

// the user's bio.
func (o GetUserResultOutput) Bio() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Bio }).(pulumi.StringOutput)
}

// the user's blog location.
func (o GetUserResultOutput) Blog() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Blog }).(pulumi.StringOutput)
}

// the user's company name.
func (o GetUserResultOutput) Company() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Company }).(pulumi.StringOutput)
}

// the creation date.
func (o GetUserResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// the user's email.
func (o GetUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// the number of followers.
func (o GetUserResultOutput) Followers() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserResult) int { return v.Followers }).(pulumi.IntOutput)
}

// the number of following users.
func (o GetUserResultOutput) Following() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserResult) int { return v.Following }).(pulumi.IntOutput)
}

// list of user's GPG keys.
func (o GetUserResultOutput) GpgKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserResult) []string { return v.GpgKeys }).(pulumi.StringArrayOutput)
}

// the user's gravatar ID.
func (o GetUserResultOutput) GravatarId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.GravatarId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// the user's location.
func (o GetUserResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Location }).(pulumi.StringOutput)
}

// the user's login.
func (o GetUserResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Login }).(pulumi.StringOutput)
}

// the user's full name.
func (o GetUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// the Node ID of the user.
func (o GetUserResultOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.NodeId }).(pulumi.StringOutput)
}

// the number of public gists.
func (o GetUserResultOutput) PublicGists() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserResult) int { return v.PublicGists }).(pulumi.IntOutput)
}

// the number of public repositories.
func (o GetUserResultOutput) PublicRepos() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserResult) int { return v.PublicRepos }).(pulumi.IntOutput)
}

// whether the user is a GitHub admin.
func (o GetUserResultOutput) SiteAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v GetUserResult) bool { return v.SiteAdmin }).(pulumi.BoolOutput)
}

// list of user's SSH keys.
func (o GetUserResultOutput) SshKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserResult) []string { return v.SshKeys }).(pulumi.StringArrayOutput)
}

// the suspended date if the user is suspended.
func (o GetUserResultOutput) SuspendedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.SuspendedAt }).(pulumi.StringOutput)
}

// the update date.
func (o GetUserResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o GetUserResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserResultOutput{})
}
