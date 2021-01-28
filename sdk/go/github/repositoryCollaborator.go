// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a GitHub repository collaborator resource.
//
// This resource allows you to add/remove collaborators from repositories in your
// organization or personal account. For organization repositories, collaborators can
// have explicit (and differing levels of) read, write, or administrator access to
// specific repositories, without giving the user full organization membership.
// For personal repositories, collaborators can only be granted write
// (implictly includes read) permission.
//
// When applied, an invitation will be sent to the user to become a collaborator
// on a repository. When destroyed, either the invitation will be cancelled or the
// collaborator will be removed from the repository.
//
// Further documentation on GitHub collaborators:
//
// - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
// - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
// - [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)
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
// 		_, err := github.NewRepositoryCollaborator(ctx, "aRepoCollaborator", &github.RepositoryCollaboratorArgs{
// 			Permission: pulumi.String("admin"),
// 			Repository: pulumi.String("our-cool-repo"),
// 			Username:   pulumi.String("SomeUser"),
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
// GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.
//
// ```sh
//  $ pulumi import github:index/repositoryCollaborator:RepositoryCollaborator collaborator terraform:someuser
// ```
type RepositoryCollaborator struct {
	pulumi.CustomResourceState

	// ID of the invitation to be used in `UserInvitationAccepter`
	InvitationId pulumi.StringOutput `pulumi:"invitationId"`
	// The permission of the outside collaborator for the repository.
	// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
	// Must be `push` for personal repositories. Defaults to `push`.
	Permission pulumi.StringPtrOutput `pulumi:"permission"`
	// The GitHub repository
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The user to add to the repository as a collaborator.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewRepositoryCollaborator registers a new resource with the given unique name, arguments, and options.
func NewRepositoryCollaborator(ctx *pulumi.Context,
	name string, args *RepositoryCollaboratorArgs, opts ...pulumi.ResourceOption) (*RepositoryCollaborator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource RepositoryCollaborator
	err := ctx.RegisterResource("github:index/repositoryCollaborator:RepositoryCollaborator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryCollaborator gets an existing RepositoryCollaborator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryCollaborator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryCollaboratorState, opts ...pulumi.ResourceOption) (*RepositoryCollaborator, error) {
	var resource RepositoryCollaborator
	err := ctx.ReadResource("github:index/repositoryCollaborator:RepositoryCollaborator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryCollaborator resources.
type repositoryCollaboratorState struct {
	// ID of the invitation to be used in `UserInvitationAccepter`
	InvitationId *string `pulumi:"invitationId"`
	// The permission of the outside collaborator for the repository.
	// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
	// Must be `push` for personal repositories. Defaults to `push`.
	Permission *string `pulumi:"permission"`
	// The GitHub repository
	Repository *string `pulumi:"repository"`
	// The user to add to the repository as a collaborator.
	Username *string `pulumi:"username"`
}

type RepositoryCollaboratorState struct {
	// ID of the invitation to be used in `UserInvitationAccepter`
	InvitationId pulumi.StringPtrInput
	// The permission of the outside collaborator for the repository.
	// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
	// Must be `push` for personal repositories. Defaults to `push`.
	Permission pulumi.StringPtrInput
	// The GitHub repository
	Repository pulumi.StringPtrInput
	// The user to add to the repository as a collaborator.
	Username pulumi.StringPtrInput
}

func (RepositoryCollaboratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryCollaboratorState)(nil)).Elem()
}

type repositoryCollaboratorArgs struct {
	// The permission of the outside collaborator for the repository.
	// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
	// Must be `push` for personal repositories. Defaults to `push`.
	Permission *string `pulumi:"permission"`
	// The GitHub repository
	Repository string `pulumi:"repository"`
	// The user to add to the repository as a collaborator.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a RepositoryCollaborator resource.
type RepositoryCollaboratorArgs struct {
	// The permission of the outside collaborator for the repository.
	// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
	// Must be `push` for personal repositories. Defaults to `push`.
	Permission pulumi.StringPtrInput
	// The GitHub repository
	Repository pulumi.StringInput
	// The user to add to the repository as a collaborator.
	Username pulumi.StringInput
}

func (RepositoryCollaboratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryCollaboratorArgs)(nil)).Elem()
}

type RepositoryCollaboratorInput interface {
	pulumi.Input

	ToRepositoryCollaboratorOutput() RepositoryCollaboratorOutput
	ToRepositoryCollaboratorOutputWithContext(ctx context.Context) RepositoryCollaboratorOutput
}

func (*RepositoryCollaborator) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryCollaborator)(nil))
}

func (i *RepositoryCollaborator) ToRepositoryCollaboratorOutput() RepositoryCollaboratorOutput {
	return i.ToRepositoryCollaboratorOutputWithContext(context.Background())
}

func (i *RepositoryCollaborator) ToRepositoryCollaboratorOutputWithContext(ctx context.Context) RepositoryCollaboratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCollaboratorOutput)
}

type RepositoryCollaboratorOutput struct {
	*pulumi.OutputState
}

func (RepositoryCollaboratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryCollaborator)(nil))
}

func (o RepositoryCollaboratorOutput) ToRepositoryCollaboratorOutput() RepositoryCollaboratorOutput {
	return o
}

func (o RepositoryCollaboratorOutput) ToRepositoryCollaboratorOutputWithContext(ctx context.Context) RepositoryCollaboratorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RepositoryCollaboratorOutput{})
}
