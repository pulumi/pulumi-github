// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package users

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a GitHub user's SSH key resource.
// 
// This resource allows you to add/remove SSH keys from your user account.
type SshKey struct {
	s *pulumi.ResourceState
}

// NewSshKey registers a new resource with the given unique name, arguments, and options.
func NewSshKey(ctx *pulumi.Context,
	name string, args *SshKeyArgs, opts ...pulumi.ResourceOpt) (*SshKey, error) {
	if args == nil || args.Key == nil {
		return nil, errors.New("missing required argument 'Key'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["key"] = nil
		inputs["title"] = nil
	} else {
		inputs["key"] = args.Key
		inputs["title"] = args.Title
	}
	inputs["etag"] = nil
	inputs["url"] = nil
	s, err := ctx.RegisterResource("github:users/sshKey:SshKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SshKey{s: s}, nil
}

// GetSshKey gets an existing SshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SshKeyState, opts ...pulumi.ResourceOpt) (*SshKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["key"] = state.Key
		inputs["title"] = state.Title
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("github:users/sshKey:SshKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SshKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SshKey) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SshKey) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *SshKey) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The public SSH key to add to your GitHub account.
func (r *SshKey) Key() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["key"])
}

// A descriptive name for the new key. e.g. `Personal MacBook Air`
func (r *SshKey) Title() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["title"])
}

// The URL of the SSH key
func (r *SshKey) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering SshKey resources.
type SshKeyState struct {
	Etag interface{}
	// The public SSH key to add to your GitHub account.
	Key interface{}
	// A descriptive name for the new key. e.g. `Personal MacBook Air`
	Title interface{}
	// The URL of the SSH key
	Url interface{}
}

// The set of arguments for constructing a SshKey resource.
type SshKeyArgs struct {
	// The public SSH key to add to your GitHub account.
	Key interface{}
	// A descriptive name for the new key. e.g. `Personal MacBook Air`
	Title interface{}
}
