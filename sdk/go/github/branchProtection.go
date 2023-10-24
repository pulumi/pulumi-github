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

// ## Import
//
// GitHub Branch Protection can be imported using an ID made up of `repository:pattern`, e.g.
//
// ```sh
//
//	$ pulumi import github:index/branchProtection:BranchProtection terraform terraform:main
//
// ```
type BranchProtection struct {
	pulumi.CustomResourceState

	// Boolean, setting this to `true` to allow the branch to be deleted.
	AllowsDeletions pulumi.BoolPtrOutput `pulumi:"allowsDeletions"`
	// Boolean, setting this to `true` to allow force pushes on the branch.
	AllowsForcePushes pulumi.BoolPtrOutput `pulumi:"allowsForcePushes"`
	// Boolean, setting this to `true` to block creating the branch.
	BlocksCreations pulumi.BoolPtrOutput `pulumi:"blocksCreations"`
	// Boolean, setting this to `true` enforces status checks for repository administrators.
	EnforceAdmins pulumi.BoolPtrOutput `pulumi:"enforceAdmins"`
	// The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	ForcePushBypassers pulumi.StringArrayOutput `pulumi:"forcePushBypassers"`
	// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
	LockBranch pulumi.BoolPtrOutput `pulumi:"lockBranch"`
	// Identifies the protection rule pattern.
	Pattern pulumi.StringOutput `pulumi:"pattern"`
	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PushRestrictions pulumi.StringArrayOutput `pulumi:"pushRestrictions"`
	// The name or node ID of the repository associated with this branch protection rule.
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
	RequireConversationResolution pulumi.BoolPtrOutput `pulumi:"requireConversationResolution"`
	// Boolean, setting this to `true` requires all commits to be signed with GPG.
	RequireSignedCommits pulumi.BoolPtrOutput `pulumi:"requireSignedCommits"`
	// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	RequiredLinearHistory pulumi.BoolPtrOutput `pulumi:"requiredLinearHistory"`
	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	RequiredPullRequestReviews BranchProtectionRequiredPullRequestReviewArrayOutput `pulumi:"requiredPullRequestReviews"`
	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	RequiredStatusChecks BranchProtectionRequiredStatusCheckArrayOutput `pulumi:"requiredStatusChecks"`
}

// NewBranchProtection registers a new resource with the given unique name, arguments, and options.
func NewBranchProtection(ctx *pulumi.Context,
	name string, args *BranchProtectionArgs, opts ...pulumi.ResourceOption) (*BranchProtection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Pattern == nil {
		return nil, errors.New("invalid value for required argument 'Pattern'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BranchProtection
	err := ctx.RegisterResource("github:index/branchProtection:BranchProtection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranchProtection gets an existing BranchProtection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranchProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchProtectionState, opts ...pulumi.ResourceOption) (*BranchProtection, error) {
	var resource BranchProtection
	err := ctx.ReadResource("github:index/branchProtection:BranchProtection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BranchProtection resources.
type branchProtectionState struct {
	// Boolean, setting this to `true` to allow the branch to be deleted.
	AllowsDeletions *bool `pulumi:"allowsDeletions"`
	// Boolean, setting this to `true` to allow force pushes on the branch.
	AllowsForcePushes *bool `pulumi:"allowsForcePushes"`
	// Boolean, setting this to `true` to block creating the branch.
	BlocksCreations *bool `pulumi:"blocksCreations"`
	// Boolean, setting this to `true` enforces status checks for repository administrators.
	EnforceAdmins *bool `pulumi:"enforceAdmins"`
	// The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	ForcePushBypassers []string `pulumi:"forcePushBypassers"`
	// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
	LockBranch *bool `pulumi:"lockBranch"`
	// Identifies the protection rule pattern.
	Pattern *string `pulumi:"pattern"`
	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PushRestrictions []string `pulumi:"pushRestrictions"`
	// The name or node ID of the repository associated with this branch protection rule.
	RepositoryId *string `pulumi:"repositoryId"`
	// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
	RequireConversationResolution *bool `pulumi:"requireConversationResolution"`
	// Boolean, setting this to `true` requires all commits to be signed with GPG.
	RequireSignedCommits *bool `pulumi:"requireSignedCommits"`
	// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	RequiredLinearHistory *bool `pulumi:"requiredLinearHistory"`
	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	RequiredPullRequestReviews []BranchProtectionRequiredPullRequestReview `pulumi:"requiredPullRequestReviews"`
	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	RequiredStatusChecks []BranchProtectionRequiredStatusCheck `pulumi:"requiredStatusChecks"`
}

type BranchProtectionState struct {
	// Boolean, setting this to `true` to allow the branch to be deleted.
	AllowsDeletions pulumi.BoolPtrInput
	// Boolean, setting this to `true` to allow force pushes on the branch.
	AllowsForcePushes pulumi.BoolPtrInput
	// Boolean, setting this to `true` to block creating the branch.
	BlocksCreations pulumi.BoolPtrInput
	// Boolean, setting this to `true` enforces status checks for repository administrators.
	EnforceAdmins pulumi.BoolPtrInput
	// The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	ForcePushBypassers pulumi.StringArrayInput
	// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
	LockBranch pulumi.BoolPtrInput
	// Identifies the protection rule pattern.
	Pattern pulumi.StringPtrInput
	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PushRestrictions pulumi.StringArrayInput
	// The name or node ID of the repository associated with this branch protection rule.
	RepositoryId pulumi.StringPtrInput
	// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
	RequireConversationResolution pulumi.BoolPtrInput
	// Boolean, setting this to `true` requires all commits to be signed with GPG.
	RequireSignedCommits pulumi.BoolPtrInput
	// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	RequiredLinearHistory pulumi.BoolPtrInput
	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	RequiredPullRequestReviews BranchProtectionRequiredPullRequestReviewArrayInput
	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	RequiredStatusChecks BranchProtectionRequiredStatusCheckArrayInput
}

func (BranchProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionState)(nil)).Elem()
}

type branchProtectionArgs struct {
	// Boolean, setting this to `true` to allow the branch to be deleted.
	AllowsDeletions *bool `pulumi:"allowsDeletions"`
	// Boolean, setting this to `true` to allow force pushes on the branch.
	AllowsForcePushes *bool `pulumi:"allowsForcePushes"`
	// Boolean, setting this to `true` to block creating the branch.
	BlocksCreations *bool `pulumi:"blocksCreations"`
	// Boolean, setting this to `true` enforces status checks for repository administrators.
	EnforceAdmins *bool `pulumi:"enforceAdmins"`
	// The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	ForcePushBypassers []string `pulumi:"forcePushBypassers"`
	// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
	LockBranch *bool `pulumi:"lockBranch"`
	// Identifies the protection rule pattern.
	Pattern string `pulumi:"pattern"`
	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PushRestrictions []string `pulumi:"pushRestrictions"`
	// The name or node ID of the repository associated with this branch protection rule.
	RepositoryId string `pulumi:"repositoryId"`
	// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
	RequireConversationResolution *bool `pulumi:"requireConversationResolution"`
	// Boolean, setting this to `true` requires all commits to be signed with GPG.
	RequireSignedCommits *bool `pulumi:"requireSignedCommits"`
	// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	RequiredLinearHistory *bool `pulumi:"requiredLinearHistory"`
	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	RequiredPullRequestReviews []BranchProtectionRequiredPullRequestReview `pulumi:"requiredPullRequestReviews"`
	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	RequiredStatusChecks []BranchProtectionRequiredStatusCheck `pulumi:"requiredStatusChecks"`
}

// The set of arguments for constructing a BranchProtection resource.
type BranchProtectionArgs struct {
	// Boolean, setting this to `true` to allow the branch to be deleted.
	AllowsDeletions pulumi.BoolPtrInput
	// Boolean, setting this to `true` to allow force pushes on the branch.
	AllowsForcePushes pulumi.BoolPtrInput
	// Boolean, setting this to `true` to block creating the branch.
	BlocksCreations pulumi.BoolPtrInput
	// Boolean, setting this to `true` enforces status checks for repository administrators.
	EnforceAdmins pulumi.BoolPtrInput
	// The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	ForcePushBypassers pulumi.StringArrayInput
	// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
	LockBranch pulumi.BoolPtrInput
	// Identifies the protection rule pattern.
	Pattern pulumi.StringInput
	// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
	PushRestrictions pulumi.StringArrayInput
	// The name or node ID of the repository associated with this branch protection rule.
	RepositoryId pulumi.StringInput
	// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
	RequireConversationResolution pulumi.BoolPtrInput
	// Boolean, setting this to `true` requires all commits to be signed with GPG.
	RequireSignedCommits pulumi.BoolPtrInput
	// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
	RequiredLinearHistory pulumi.BoolPtrInput
	// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
	RequiredPullRequestReviews BranchProtectionRequiredPullRequestReviewArrayInput
	// Enforce restrictions for required status checks. See Required Status Checks below for details.
	RequiredStatusChecks BranchProtectionRequiredStatusCheckArrayInput
}

func (BranchProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchProtectionArgs)(nil)).Elem()
}

type BranchProtectionInput interface {
	pulumi.Input

	ToBranchProtectionOutput() BranchProtectionOutput
	ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput
}

func (*BranchProtection) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtection)(nil)).Elem()
}

func (i *BranchProtection) ToBranchProtectionOutput() BranchProtectionOutput {
	return i.ToBranchProtectionOutputWithContext(context.Background())
}

func (i *BranchProtection) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionOutput)
}

func (i *BranchProtection) ToOutput(ctx context.Context) pulumix.Output[*BranchProtection] {
	return pulumix.Output[*BranchProtection]{
		OutputState: i.ToBranchProtectionOutputWithContext(ctx).OutputState,
	}
}

// BranchProtectionArrayInput is an input type that accepts BranchProtectionArray and BranchProtectionArrayOutput values.
// You can construct a concrete instance of `BranchProtectionArrayInput` via:
//
//	BranchProtectionArray{ BranchProtectionArgs{...} }
type BranchProtectionArrayInput interface {
	pulumi.Input

	ToBranchProtectionArrayOutput() BranchProtectionArrayOutput
	ToBranchProtectionArrayOutputWithContext(context.Context) BranchProtectionArrayOutput
}

type BranchProtectionArray []BranchProtectionInput

func (BranchProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtection)(nil)).Elem()
}

func (i BranchProtectionArray) ToBranchProtectionArrayOutput() BranchProtectionArrayOutput {
	return i.ToBranchProtectionArrayOutputWithContext(context.Background())
}

func (i BranchProtectionArray) ToBranchProtectionArrayOutputWithContext(ctx context.Context) BranchProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionArrayOutput)
}

func (i BranchProtectionArray) ToOutput(ctx context.Context) pulumix.Output[[]*BranchProtection] {
	return pulumix.Output[[]*BranchProtection]{
		OutputState: i.ToBranchProtectionArrayOutputWithContext(ctx).OutputState,
	}
}

// BranchProtectionMapInput is an input type that accepts BranchProtectionMap and BranchProtectionMapOutput values.
// You can construct a concrete instance of `BranchProtectionMapInput` via:
//
//	BranchProtectionMap{ "key": BranchProtectionArgs{...} }
type BranchProtectionMapInput interface {
	pulumi.Input

	ToBranchProtectionMapOutput() BranchProtectionMapOutput
	ToBranchProtectionMapOutputWithContext(context.Context) BranchProtectionMapOutput
}

type BranchProtectionMap map[string]BranchProtectionInput

func (BranchProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtection)(nil)).Elem()
}

func (i BranchProtectionMap) ToBranchProtectionMapOutput() BranchProtectionMapOutput {
	return i.ToBranchProtectionMapOutputWithContext(context.Background())
}

func (i BranchProtectionMap) ToBranchProtectionMapOutputWithContext(ctx context.Context) BranchProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchProtectionMapOutput)
}

func (i BranchProtectionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BranchProtection] {
	return pulumix.Output[map[string]*BranchProtection]{
		OutputState: i.ToBranchProtectionMapOutputWithContext(ctx).OutputState,
	}
}

type BranchProtectionOutput struct{ *pulumi.OutputState }

func (BranchProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BranchProtection)(nil)).Elem()
}

func (o BranchProtectionOutput) ToBranchProtectionOutput() BranchProtectionOutput {
	return o
}

func (o BranchProtectionOutput) ToBranchProtectionOutputWithContext(ctx context.Context) BranchProtectionOutput {
	return o
}

func (o BranchProtectionOutput) ToOutput(ctx context.Context) pulumix.Output[*BranchProtection] {
	return pulumix.Output[*BranchProtection]{
		OutputState: o.OutputState,
	}
}

// Boolean, setting this to `true` to allow the branch to be deleted.
func (o BranchProtectionOutput) AllowsDeletions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.AllowsDeletions }).(pulumi.BoolPtrOutput)
}

// Boolean, setting this to `true` to allow force pushes on the branch.
func (o BranchProtectionOutput) AllowsForcePushes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.AllowsForcePushes }).(pulumi.BoolPtrOutput)
}

// Boolean, setting this to `true` to block creating the branch.
func (o BranchProtectionOutput) BlocksCreations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.BlocksCreations }).(pulumi.BoolPtrOutput)
}

// Boolean, setting this to `true` enforces status checks for repository administrators.
func (o BranchProtectionOutput) EnforceAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.EnforceAdmins }).(pulumi.BoolPtrOutput)
}

// The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
func (o BranchProtectionOutput) ForcePushBypassers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringArrayOutput { return v.ForcePushBypassers }).(pulumi.StringArrayOutput)
}

// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
func (o BranchProtectionOutput) LockBranch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.LockBranch }).(pulumi.BoolPtrOutput)
}

// Identifies the protection rule pattern.
func (o BranchProtectionOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.Pattern }).(pulumi.StringOutput)
}

// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
func (o BranchProtectionOutput) PushRestrictions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringArrayOutput { return v.PushRestrictions }).(pulumi.StringArrayOutput)
}

// The name or node ID of the repository associated with this branch protection rule.
func (o BranchProtectionOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
func (o BranchProtectionOutput) RequireConversationResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.RequireConversationResolution }).(pulumi.BoolPtrOutput)
}

// Boolean, setting this to `true` requires all commits to be signed with GPG.
func (o BranchProtectionOutput) RequireSignedCommits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.RequireSignedCommits }).(pulumi.BoolPtrOutput)
}

// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
func (o BranchProtectionOutput) RequiredLinearHistory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BranchProtection) pulumi.BoolPtrOutput { return v.RequiredLinearHistory }).(pulumi.BoolPtrOutput)
}

// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
func (o BranchProtectionOutput) RequiredPullRequestReviews() BranchProtectionRequiredPullRequestReviewArrayOutput {
	return o.ApplyT(func(v *BranchProtection) BranchProtectionRequiredPullRequestReviewArrayOutput {
		return v.RequiredPullRequestReviews
	}).(BranchProtectionRequiredPullRequestReviewArrayOutput)
}

// Enforce restrictions for required status checks. See Required Status Checks below for details.
func (o BranchProtectionOutput) RequiredStatusChecks() BranchProtectionRequiredStatusCheckArrayOutput {
	return o.ApplyT(func(v *BranchProtection) BranchProtectionRequiredStatusCheckArrayOutput {
		return v.RequiredStatusChecks
	}).(BranchProtectionRequiredStatusCheckArrayOutput)
}

type BranchProtectionArrayOutput struct{ *pulumi.OutputState }

func (BranchProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BranchProtection)(nil)).Elem()
}

func (o BranchProtectionArrayOutput) ToBranchProtectionArrayOutput() BranchProtectionArrayOutput {
	return o
}

func (o BranchProtectionArrayOutput) ToBranchProtectionArrayOutputWithContext(ctx context.Context) BranchProtectionArrayOutput {
	return o
}

func (o BranchProtectionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BranchProtection] {
	return pulumix.Output[[]*BranchProtection]{
		OutputState: o.OutputState,
	}
}

func (o BranchProtectionArrayOutput) Index(i pulumi.IntInput) BranchProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BranchProtection {
		return vs[0].([]*BranchProtection)[vs[1].(int)]
	}).(BranchProtectionOutput)
}

type BranchProtectionMapOutput struct{ *pulumi.OutputState }

func (BranchProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BranchProtection)(nil)).Elem()
}

func (o BranchProtectionMapOutput) ToBranchProtectionMapOutput() BranchProtectionMapOutput {
	return o
}

func (o BranchProtectionMapOutput) ToBranchProtectionMapOutputWithContext(ctx context.Context) BranchProtectionMapOutput {
	return o
}

func (o BranchProtectionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BranchProtection] {
	return pulumix.Output[map[string]*BranchProtection]{
		OutputState: o.OutputState,
	}
}

func (o BranchProtectionMapOutput) MapIndex(k pulumi.StringInput) BranchProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BranchProtection {
		return vs[0].(map[string]*BranchProtection)[vs[1].(string)]
	}).(BranchProtectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionInput)(nil)).Elem(), &BranchProtection{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionArrayInput)(nil)).Elem(), BranchProtectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BranchProtectionMapInput)(nil)).Elem(), BranchProtectionMap{})
	pulumi.RegisterOutputType(BranchProtectionOutput{})
	pulumi.RegisterOutputType(BranchProtectionArrayOutput{})
	pulumi.RegisterOutputType(BranchProtectionMapOutput{})
}
