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

// This resource allows you to create and manage environments for a GitHub repository.
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
//			current, err := github.GetUser(ctx, &github.GetUserArgs{
//				Username: "",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleRepository, err := github.NewRepository(ctx, "exampleRepository", &github.RepositoryArgs{
//				Description: pulumi.String("My awesome codebase"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryEnvironment(ctx, "exampleRepositoryEnvironment", &github.RepositoryEnvironmentArgs{
//				Environment:       pulumi.String("example"),
//				Repository:        exampleRepository.Name,
//				PreventSelfReview: pulumi.Bool(true),
//				Reviewers: github.RepositoryEnvironmentReviewerArray{
//					&github.RepositoryEnvironmentReviewerArgs{
//						Users: pulumi.IntArray{
//							*pulumi.String(current.Id),
//						},
//					},
//				},
//				DeploymentBranchPolicy: &github.RepositoryEnvironmentDeploymentBranchPolicyArgs{
//					ProtectedBranches:    pulumi.Bool(true),
//					CustomBranchPolicies: pulumi.Bool(false),
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
// GitHub Repository Environment can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment, separated by a `:` character, e.g.
//
// ```sh
//
//	$ pulumi import github:index/repositoryEnvironment:RepositoryEnvironment daily terraform:daily
//
// ```
type RepositoryEnvironment struct {
	pulumi.CustomResourceState

	// Can repository admins bypass the environment protections.  Defaults to `true`.
	CanAdminsBypass pulumi.BoolPtrOutput `pulumi:"canAdminsBypass"`
	// The deployment branch policy configuration
	DeploymentBranchPolicy RepositoryEnvironmentDeploymentBranchPolicyPtrOutput `pulumi:"deploymentBranchPolicy"`
	// The name of the environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
	PreventSelfReview pulumi.BoolPtrOutput `pulumi:"preventSelfReview"`
	// The repository of the environment.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The environment reviewers configuration.
	Reviewers RepositoryEnvironmentReviewerArrayOutput `pulumi:"reviewers"`
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer pulumi.IntPtrOutput `pulumi:"waitTimer"`
}

// NewRepositoryEnvironment registers a new resource with the given unique name, arguments, and options.
func NewRepositoryEnvironment(ctx *pulumi.Context,
	name string, args *RepositoryEnvironmentArgs, opts ...pulumi.ResourceOption) (*RepositoryEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryEnvironment
	err := ctx.RegisterResource("github:index/repositoryEnvironment:RepositoryEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryEnvironment gets an existing RepositoryEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryEnvironmentState, opts ...pulumi.ResourceOption) (*RepositoryEnvironment, error) {
	var resource RepositoryEnvironment
	err := ctx.ReadResource("github:index/repositoryEnvironment:RepositoryEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryEnvironment resources.
type repositoryEnvironmentState struct {
	// Can repository admins bypass the environment protections.  Defaults to `true`.
	CanAdminsBypass *bool `pulumi:"canAdminsBypass"`
	// The deployment branch policy configuration
	DeploymentBranchPolicy *RepositoryEnvironmentDeploymentBranchPolicy `pulumi:"deploymentBranchPolicy"`
	// The name of the environment.
	Environment *string `pulumi:"environment"`
	// Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
	PreventSelfReview *bool `pulumi:"preventSelfReview"`
	// The repository of the environment.
	Repository *string `pulumi:"repository"`
	// The environment reviewers configuration.
	Reviewers []RepositoryEnvironmentReviewer `pulumi:"reviewers"`
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer *int `pulumi:"waitTimer"`
}

type RepositoryEnvironmentState struct {
	// Can repository admins bypass the environment protections.  Defaults to `true`.
	CanAdminsBypass pulumi.BoolPtrInput
	// The deployment branch policy configuration
	DeploymentBranchPolicy RepositoryEnvironmentDeploymentBranchPolicyPtrInput
	// The name of the environment.
	Environment pulumi.StringPtrInput
	// Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
	PreventSelfReview pulumi.BoolPtrInput
	// The repository of the environment.
	Repository pulumi.StringPtrInput
	// The environment reviewers configuration.
	Reviewers RepositoryEnvironmentReviewerArrayInput
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer pulumi.IntPtrInput
}

func (RepositoryEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryEnvironmentState)(nil)).Elem()
}

type repositoryEnvironmentArgs struct {
	// Can repository admins bypass the environment protections.  Defaults to `true`.
	CanAdminsBypass *bool `pulumi:"canAdminsBypass"`
	// The deployment branch policy configuration
	DeploymentBranchPolicy *RepositoryEnvironmentDeploymentBranchPolicy `pulumi:"deploymentBranchPolicy"`
	// The name of the environment.
	Environment string `pulumi:"environment"`
	// Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
	PreventSelfReview *bool `pulumi:"preventSelfReview"`
	// The repository of the environment.
	Repository string `pulumi:"repository"`
	// The environment reviewers configuration.
	Reviewers []RepositoryEnvironmentReviewer `pulumi:"reviewers"`
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer *int `pulumi:"waitTimer"`
}

// The set of arguments for constructing a RepositoryEnvironment resource.
type RepositoryEnvironmentArgs struct {
	// Can repository admins bypass the environment protections.  Defaults to `true`.
	CanAdminsBypass pulumi.BoolPtrInput
	// The deployment branch policy configuration
	DeploymentBranchPolicy RepositoryEnvironmentDeploymentBranchPolicyPtrInput
	// The name of the environment.
	Environment pulumi.StringInput
	// Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
	PreventSelfReview pulumi.BoolPtrInput
	// The repository of the environment.
	Repository pulumi.StringInput
	// The environment reviewers configuration.
	Reviewers RepositoryEnvironmentReviewerArrayInput
	// Amount of time to delay a job after the job is initially triggered.
	WaitTimer pulumi.IntPtrInput
}

func (RepositoryEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryEnvironmentArgs)(nil)).Elem()
}

type RepositoryEnvironmentInput interface {
	pulumi.Input

	ToRepositoryEnvironmentOutput() RepositoryEnvironmentOutput
	ToRepositoryEnvironmentOutputWithContext(ctx context.Context) RepositoryEnvironmentOutput
}

func (*RepositoryEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryEnvironment)(nil)).Elem()
}

func (i *RepositoryEnvironment) ToRepositoryEnvironmentOutput() RepositoryEnvironmentOutput {
	return i.ToRepositoryEnvironmentOutputWithContext(context.Background())
}

func (i *RepositoryEnvironment) ToRepositoryEnvironmentOutputWithContext(ctx context.Context) RepositoryEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryEnvironmentOutput)
}

// RepositoryEnvironmentArrayInput is an input type that accepts RepositoryEnvironmentArray and RepositoryEnvironmentArrayOutput values.
// You can construct a concrete instance of `RepositoryEnvironmentArrayInput` via:
//
//	RepositoryEnvironmentArray{ RepositoryEnvironmentArgs{...} }
type RepositoryEnvironmentArrayInput interface {
	pulumi.Input

	ToRepositoryEnvironmentArrayOutput() RepositoryEnvironmentArrayOutput
	ToRepositoryEnvironmentArrayOutputWithContext(context.Context) RepositoryEnvironmentArrayOutput
}

type RepositoryEnvironmentArray []RepositoryEnvironmentInput

func (RepositoryEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryEnvironment)(nil)).Elem()
}

func (i RepositoryEnvironmentArray) ToRepositoryEnvironmentArrayOutput() RepositoryEnvironmentArrayOutput {
	return i.ToRepositoryEnvironmentArrayOutputWithContext(context.Background())
}

func (i RepositoryEnvironmentArray) ToRepositoryEnvironmentArrayOutputWithContext(ctx context.Context) RepositoryEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryEnvironmentArrayOutput)
}

// RepositoryEnvironmentMapInput is an input type that accepts RepositoryEnvironmentMap and RepositoryEnvironmentMapOutput values.
// You can construct a concrete instance of `RepositoryEnvironmentMapInput` via:
//
//	RepositoryEnvironmentMap{ "key": RepositoryEnvironmentArgs{...} }
type RepositoryEnvironmentMapInput interface {
	pulumi.Input

	ToRepositoryEnvironmentMapOutput() RepositoryEnvironmentMapOutput
	ToRepositoryEnvironmentMapOutputWithContext(context.Context) RepositoryEnvironmentMapOutput
}

type RepositoryEnvironmentMap map[string]RepositoryEnvironmentInput

func (RepositoryEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryEnvironment)(nil)).Elem()
}

func (i RepositoryEnvironmentMap) ToRepositoryEnvironmentMapOutput() RepositoryEnvironmentMapOutput {
	return i.ToRepositoryEnvironmentMapOutputWithContext(context.Background())
}

func (i RepositoryEnvironmentMap) ToRepositoryEnvironmentMapOutputWithContext(ctx context.Context) RepositoryEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryEnvironmentMapOutput)
}

type RepositoryEnvironmentOutput struct{ *pulumi.OutputState }

func (RepositoryEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryEnvironment)(nil)).Elem()
}

func (o RepositoryEnvironmentOutput) ToRepositoryEnvironmentOutput() RepositoryEnvironmentOutput {
	return o
}

func (o RepositoryEnvironmentOutput) ToRepositoryEnvironmentOutputWithContext(ctx context.Context) RepositoryEnvironmentOutput {
	return o
}

// Can repository admins bypass the environment protections.  Defaults to `true`.
func (o RepositoryEnvironmentOutput) CanAdminsBypass() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryEnvironment) pulumi.BoolPtrOutput { return v.CanAdminsBypass }).(pulumi.BoolPtrOutput)
}

// The deployment branch policy configuration
func (o RepositoryEnvironmentOutput) DeploymentBranchPolicy() RepositoryEnvironmentDeploymentBranchPolicyPtrOutput {
	return o.ApplyT(func(v *RepositoryEnvironment) RepositoryEnvironmentDeploymentBranchPolicyPtrOutput {
		return v.DeploymentBranchPolicy
	}).(RepositoryEnvironmentDeploymentBranchPolicyPtrOutput)
}

// The name of the environment.
func (o RepositoryEnvironmentOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryEnvironment) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
func (o RepositoryEnvironmentOutput) PreventSelfReview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryEnvironment) pulumi.BoolPtrOutput { return v.PreventSelfReview }).(pulumi.BoolPtrOutput)
}

// The repository of the environment.
func (o RepositoryEnvironmentOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryEnvironment) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The environment reviewers configuration.
func (o RepositoryEnvironmentOutput) Reviewers() RepositoryEnvironmentReviewerArrayOutput {
	return o.ApplyT(func(v *RepositoryEnvironment) RepositoryEnvironmentReviewerArrayOutput { return v.Reviewers }).(RepositoryEnvironmentReviewerArrayOutput)
}

// Amount of time to delay a job after the job is initially triggered.
func (o RepositoryEnvironmentOutput) WaitTimer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RepositoryEnvironment) pulumi.IntPtrOutput { return v.WaitTimer }).(pulumi.IntPtrOutput)
}

type RepositoryEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (RepositoryEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryEnvironment)(nil)).Elem()
}

func (o RepositoryEnvironmentArrayOutput) ToRepositoryEnvironmentArrayOutput() RepositoryEnvironmentArrayOutput {
	return o
}

func (o RepositoryEnvironmentArrayOutput) ToRepositoryEnvironmentArrayOutputWithContext(ctx context.Context) RepositoryEnvironmentArrayOutput {
	return o
}

func (o RepositoryEnvironmentArrayOutput) Index(i pulumi.IntInput) RepositoryEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryEnvironment {
		return vs[0].([]*RepositoryEnvironment)[vs[1].(int)]
	}).(RepositoryEnvironmentOutput)
}

type RepositoryEnvironmentMapOutput struct{ *pulumi.OutputState }

func (RepositoryEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryEnvironment)(nil)).Elem()
}

func (o RepositoryEnvironmentMapOutput) ToRepositoryEnvironmentMapOutput() RepositoryEnvironmentMapOutput {
	return o
}

func (o RepositoryEnvironmentMapOutput) ToRepositoryEnvironmentMapOutputWithContext(ctx context.Context) RepositoryEnvironmentMapOutput {
	return o
}

func (o RepositoryEnvironmentMapOutput) MapIndex(k pulumi.StringInput) RepositoryEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryEnvironment {
		return vs[0].(map[string]*RepositoryEnvironment)[vs[1].(string)]
	}).(RepositoryEnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryEnvironmentInput)(nil)).Elem(), &RepositoryEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryEnvironmentArrayInput)(nil)).Elem(), RepositoryEnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryEnvironmentMapInput)(nil)).Elem(), RepositoryEnvironmentMap{})
	pulumi.RegisterOutputType(RepositoryEnvironmentOutput{})
	pulumi.RegisterOutputType(RepositoryEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(RepositoryEnvironmentMapOutput{})
}
