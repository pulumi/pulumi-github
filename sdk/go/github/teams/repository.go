// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package teams

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource manages relationships between teams and repositories
// in your GitHub organization.
// 
// Creating this resource grants a particular team permissions on a
// particular repository.
// 
// The repository and the team must both belong to the same organization
// on GitHub. This resource does not actually *create* any repositories;
// to do that, see `github_repository`.
type Repository struct {
	s *pulumi.ResourceState
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOpt) (*Repository, error) {
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	if args == nil || args.TeamId == nil {
		return nil, errors.New("missing required argument 'TeamId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["permission"] = nil
		inputs["repository"] = nil
		inputs["teamId"] = nil
	} else {
		inputs["permission"] = args.Permission
		inputs["repository"] = args.Repository
		inputs["teamId"] = args.TeamId
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("github:teams/repository:Repository", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Repository{s: s}, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RepositoryState, opts ...pulumi.ResourceOpt) (*Repository, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["permission"] = state.Permission
		inputs["repository"] = state.Repository
		inputs["teamId"] = state.TeamId
	}
	s, err := ctx.ReadResource("github:teams/repository:Repository", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Repository{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Repository) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Repository) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Repository) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The permissions of team members regarding the repository.
// Must be one of `pull`, `push`, or `admin`. Defaults to `pull`.
func (r *Repository) Permission() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["permission"])
}

// The repository to add to the team.
func (r *Repository) Repository() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["repository"])
}

// The GitHub team id
func (r *Repository) TeamId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["teamId"])
}

// Input properties used for looking up and filtering Repository resources.
type RepositoryState struct {
	Etag interface{}
	// The permissions of team members regarding the repository.
	// Must be one of `pull`, `push`, or `admin`. Defaults to `pull`.
	Permission interface{}
	// The repository to add to the team.
	Repository interface{}
	// The GitHub team id
	TeamId interface{}
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The permissions of team members regarding the repository.
	// Must be one of `pull`, `push`, or `admin`. Defaults to `pull`.
	Permission interface{}
	// The repository to add to the team.
	Repository interface{}
	// The GitHub team id
	TeamId interface{}
}