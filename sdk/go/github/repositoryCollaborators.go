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

// Provides a GitHub repository collaborators resource.
//
// > Note: RepositoryCollaborators cannot be used in conjunction with RepositoryCollaborator and
// TeamRepository or they will fight over what your policy should be.
//
// This resource allows you to manage all collaborators for repositories in your
// organization or personal account. For organization repositories, collaborators can
// have explicit (and differing levels of) read, write, or administrator access to
// specific repositories, without giving the user full organization membership.
// For personal repositories, collaborators can only be granted write
// (implicitly includes read) permission.
//
// When applied, an invitation will be sent to the user to become a collaborators
// on a repository. When destroyed, either the invitation will be cancelled or the
// collaborators will be removed from the repository.
//
// This resource is authoritative. For adding a collaborator to a repo in a non-authoritative manner, use
// RepositoryCollaborator instead.
//
// Further documentation on GitHub collaborators:
//
// - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
// - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
// - [Converting an organization member to an outside collaborators](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)
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
//			someTeam, err := github.NewTeam(ctx, "someTeam", &github.TeamArgs{
//				Description: pulumi.String("Some cool team"),
//			})
//			if err != nil {
//				return err
//			}
//			someRepo, err := github.NewRepository(ctx, "someRepo", nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryCollaborators(ctx, "someRepoCollaborators", &github.RepositoryCollaboratorsArgs{
//				Repository: someRepo.Name,
//				Users: github.RepositoryCollaboratorsUserArray{
//					&github.RepositoryCollaboratorsUserArgs{
//						Permission: pulumi.String("admin"),
//						Username:   pulumi.String("SomeUser"),
//					},
//				},
//				Teams: github.RepositoryCollaboratorsTeamArray{
//					&github.RepositoryCollaboratorsTeamArgs{
//						Permission: pulumi.String("pull"),
//						TeamId:     someTeam.Slug,
//					},
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
// GitHub Repository Collaborators can be imported using the name `name`, e.g.
//
// ```sh
//
//	$ pulumi import github:index/repositoryCollaborators:RepositoryCollaborators collaborators terraform
//
// ```
type RepositoryCollaborators struct {
	pulumi.CustomResourceState

	// Map of usernames to invitation ID for any users added as part of creation of this resource to
	// be used in `UserInvitationAccepter`.
	InvitationIds pulumi.StringMapOutput `pulumi:"invitationIds"`
	// The GitHub repository
	Repository pulumi.StringOutput `pulumi:"repository"`
	// List of teams
	Teams RepositoryCollaboratorsTeamArrayOutput `pulumi:"teams"`
	// List of users
	Users RepositoryCollaboratorsUserArrayOutput `pulumi:"users"`
}

// NewRepositoryCollaborators registers a new resource with the given unique name, arguments, and options.
func NewRepositoryCollaborators(ctx *pulumi.Context,
	name string, args *RepositoryCollaboratorsArgs, opts ...pulumi.ResourceOption) (*RepositoryCollaborators, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryCollaborators
	err := ctx.RegisterResource("github:index/repositoryCollaborators:RepositoryCollaborators", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryCollaborators gets an existing RepositoryCollaborators resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryCollaborators(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryCollaboratorsState, opts ...pulumi.ResourceOption) (*RepositoryCollaborators, error) {
	var resource RepositoryCollaborators
	err := ctx.ReadResource("github:index/repositoryCollaborators:RepositoryCollaborators", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryCollaborators resources.
type repositoryCollaboratorsState struct {
	// Map of usernames to invitation ID for any users added as part of creation of this resource to
	// be used in `UserInvitationAccepter`.
	InvitationIds map[string]string `pulumi:"invitationIds"`
	// The GitHub repository
	Repository *string `pulumi:"repository"`
	// List of teams
	Teams []RepositoryCollaboratorsTeam `pulumi:"teams"`
	// List of users
	Users []RepositoryCollaboratorsUser `pulumi:"users"`
}

type RepositoryCollaboratorsState struct {
	// Map of usernames to invitation ID for any users added as part of creation of this resource to
	// be used in `UserInvitationAccepter`.
	InvitationIds pulumi.StringMapInput
	// The GitHub repository
	Repository pulumi.StringPtrInput
	// List of teams
	Teams RepositoryCollaboratorsTeamArrayInput
	// List of users
	Users RepositoryCollaboratorsUserArrayInput
}

func (RepositoryCollaboratorsState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryCollaboratorsState)(nil)).Elem()
}

type repositoryCollaboratorsArgs struct {
	// The GitHub repository
	Repository string `pulumi:"repository"`
	// List of teams
	Teams []RepositoryCollaboratorsTeam `pulumi:"teams"`
	// List of users
	Users []RepositoryCollaboratorsUser `pulumi:"users"`
}

// The set of arguments for constructing a RepositoryCollaborators resource.
type RepositoryCollaboratorsArgs struct {
	// The GitHub repository
	Repository pulumi.StringInput
	// List of teams
	Teams RepositoryCollaboratorsTeamArrayInput
	// List of users
	Users RepositoryCollaboratorsUserArrayInput
}

func (RepositoryCollaboratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryCollaboratorsArgs)(nil)).Elem()
}

type RepositoryCollaboratorsInput interface {
	pulumi.Input

	ToRepositoryCollaboratorsOutput() RepositoryCollaboratorsOutput
	ToRepositoryCollaboratorsOutputWithContext(ctx context.Context) RepositoryCollaboratorsOutput
}

func (*RepositoryCollaborators) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryCollaborators)(nil)).Elem()
}

func (i *RepositoryCollaborators) ToRepositoryCollaboratorsOutput() RepositoryCollaboratorsOutput {
	return i.ToRepositoryCollaboratorsOutputWithContext(context.Background())
}

func (i *RepositoryCollaborators) ToRepositoryCollaboratorsOutputWithContext(ctx context.Context) RepositoryCollaboratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCollaboratorsOutput)
}

// RepositoryCollaboratorsArrayInput is an input type that accepts RepositoryCollaboratorsArray and RepositoryCollaboratorsArrayOutput values.
// You can construct a concrete instance of `RepositoryCollaboratorsArrayInput` via:
//
//	RepositoryCollaboratorsArray{ RepositoryCollaboratorsArgs{...} }
type RepositoryCollaboratorsArrayInput interface {
	pulumi.Input

	ToRepositoryCollaboratorsArrayOutput() RepositoryCollaboratorsArrayOutput
	ToRepositoryCollaboratorsArrayOutputWithContext(context.Context) RepositoryCollaboratorsArrayOutput
}

type RepositoryCollaboratorsArray []RepositoryCollaboratorsInput

func (RepositoryCollaboratorsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryCollaborators)(nil)).Elem()
}

func (i RepositoryCollaboratorsArray) ToRepositoryCollaboratorsArrayOutput() RepositoryCollaboratorsArrayOutput {
	return i.ToRepositoryCollaboratorsArrayOutputWithContext(context.Background())
}

func (i RepositoryCollaboratorsArray) ToRepositoryCollaboratorsArrayOutputWithContext(ctx context.Context) RepositoryCollaboratorsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCollaboratorsArrayOutput)
}

// RepositoryCollaboratorsMapInput is an input type that accepts RepositoryCollaboratorsMap and RepositoryCollaboratorsMapOutput values.
// You can construct a concrete instance of `RepositoryCollaboratorsMapInput` via:
//
//	RepositoryCollaboratorsMap{ "key": RepositoryCollaboratorsArgs{...} }
type RepositoryCollaboratorsMapInput interface {
	pulumi.Input

	ToRepositoryCollaboratorsMapOutput() RepositoryCollaboratorsMapOutput
	ToRepositoryCollaboratorsMapOutputWithContext(context.Context) RepositoryCollaboratorsMapOutput
}

type RepositoryCollaboratorsMap map[string]RepositoryCollaboratorsInput

func (RepositoryCollaboratorsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryCollaborators)(nil)).Elem()
}

func (i RepositoryCollaboratorsMap) ToRepositoryCollaboratorsMapOutput() RepositoryCollaboratorsMapOutput {
	return i.ToRepositoryCollaboratorsMapOutputWithContext(context.Background())
}

func (i RepositoryCollaboratorsMap) ToRepositoryCollaboratorsMapOutputWithContext(ctx context.Context) RepositoryCollaboratorsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCollaboratorsMapOutput)
}

type RepositoryCollaboratorsOutput struct{ *pulumi.OutputState }

func (RepositoryCollaboratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryCollaborators)(nil)).Elem()
}

func (o RepositoryCollaboratorsOutput) ToRepositoryCollaboratorsOutput() RepositoryCollaboratorsOutput {
	return o
}

func (o RepositoryCollaboratorsOutput) ToRepositoryCollaboratorsOutputWithContext(ctx context.Context) RepositoryCollaboratorsOutput {
	return o
}

// Map of usernames to invitation ID for any users added as part of creation of this resource to
// be used in `UserInvitationAccepter`.
func (o RepositoryCollaboratorsOutput) InvitationIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RepositoryCollaborators) pulumi.StringMapOutput { return v.InvitationIds }).(pulumi.StringMapOutput)
}

// The GitHub repository
func (o RepositoryCollaboratorsOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryCollaborators) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// List of teams
func (o RepositoryCollaboratorsOutput) Teams() RepositoryCollaboratorsTeamArrayOutput {
	return o.ApplyT(func(v *RepositoryCollaborators) RepositoryCollaboratorsTeamArrayOutput { return v.Teams }).(RepositoryCollaboratorsTeamArrayOutput)
}

// List of users
func (o RepositoryCollaboratorsOutput) Users() RepositoryCollaboratorsUserArrayOutput {
	return o.ApplyT(func(v *RepositoryCollaborators) RepositoryCollaboratorsUserArrayOutput { return v.Users }).(RepositoryCollaboratorsUserArrayOutput)
}

type RepositoryCollaboratorsArrayOutput struct{ *pulumi.OutputState }

func (RepositoryCollaboratorsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryCollaborators)(nil)).Elem()
}

func (o RepositoryCollaboratorsArrayOutput) ToRepositoryCollaboratorsArrayOutput() RepositoryCollaboratorsArrayOutput {
	return o
}

func (o RepositoryCollaboratorsArrayOutput) ToRepositoryCollaboratorsArrayOutputWithContext(ctx context.Context) RepositoryCollaboratorsArrayOutput {
	return o
}

func (o RepositoryCollaboratorsArrayOutput) Index(i pulumi.IntInput) RepositoryCollaboratorsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryCollaborators {
		return vs[0].([]*RepositoryCollaborators)[vs[1].(int)]
	}).(RepositoryCollaboratorsOutput)
}

type RepositoryCollaboratorsMapOutput struct{ *pulumi.OutputState }

func (RepositoryCollaboratorsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryCollaborators)(nil)).Elem()
}

func (o RepositoryCollaboratorsMapOutput) ToRepositoryCollaboratorsMapOutput() RepositoryCollaboratorsMapOutput {
	return o
}

func (o RepositoryCollaboratorsMapOutput) ToRepositoryCollaboratorsMapOutputWithContext(ctx context.Context) RepositoryCollaboratorsMapOutput {
	return o
}

func (o RepositoryCollaboratorsMapOutput) MapIndex(k pulumi.StringInput) RepositoryCollaboratorsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryCollaborators {
		return vs[0].(map[string]*RepositoryCollaborators)[vs[1].(string)]
	}).(RepositoryCollaboratorsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryCollaboratorsInput)(nil)).Elem(), &RepositoryCollaborators{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryCollaboratorsArrayInput)(nil)).Elem(), RepositoryCollaboratorsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryCollaboratorsMapInput)(nil)).Elem(), RepositoryCollaboratorsMap{})
	pulumi.RegisterOutputType(RepositoryCollaboratorsOutput{})
	pulumi.RegisterOutputType(RepositoryCollaboratorsArrayOutput{})
	pulumi.RegisterOutputType(RepositoryCollaboratorsMapOutput{})
}
