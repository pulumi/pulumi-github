// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package users

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage GitHub repository collaborator invitations.
type InvitationAcceptor struct {
	s *pulumi.ResourceState
}

// NewInvitationAcceptor registers a new resource with the given unique name, arguments, and options.
func NewInvitationAcceptor(ctx *pulumi.Context,
	name string, args *InvitationAcceptorArgs, opts ...pulumi.ResourceOpt) (*InvitationAcceptor, error) {
	if args == nil || args.InvitationId == nil {
		return nil, errors.New("missing required argument 'InvitationId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["invitationId"] = nil
	} else {
		inputs["invitationId"] = args.InvitationId
	}
	s, err := ctx.RegisterResource("github:users/invitationAcceptor:InvitationAcceptor", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InvitationAcceptor{s: s}, nil
}

// GetInvitationAcceptor gets an existing InvitationAcceptor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInvitationAcceptor(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InvitationAcceptorState, opts ...pulumi.ResourceOpt) (*InvitationAcceptor, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["invitationId"] = state.InvitationId
	}
	s, err := ctx.ReadResource("github:users/invitationAcceptor:InvitationAcceptor", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InvitationAcceptor{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *InvitationAcceptor) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *InvitationAcceptor) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// ID of the invitation to accept
func (r *InvitationAcceptor) InvitationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["invitationId"])
}

// Input properties used for looking up and filtering InvitationAcceptor resources.
type InvitationAcceptorState struct {
	// ID of the invitation to accept
	InvitationId interface{}
}

// The set of arguments for constructing a InvitationAcceptor resource.
type InvitationAcceptorArgs struct {
	// ID of the invitation to accept
	InvitationId interface{}
}
