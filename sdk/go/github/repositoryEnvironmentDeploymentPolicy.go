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

// This resource allows you to create and manage environment deployment branch policies for a GitHub repository.
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
//			current, err := github.GetUser(ctx, &github.GetUserArgs{
//				Username: "",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			testRepository, err := github.NewRepository(ctx, "testRepository", nil)
//			if err != nil {
//				return err
//			}
//			testRepositoryEnvironment, err := github.NewRepositoryEnvironment(ctx, "testRepositoryEnvironment", &github.RepositoryEnvironmentArgs{
//				Repository:  testRepository.Name,
//				Environment: pulumi.String("environment/test"),
//				WaitTimer:   pulumi.Int(10000),
//				Reviewers: github.RepositoryEnvironmentReviewerArray{
//					&github.RepositoryEnvironmentReviewerArgs{
//						Users: pulumi.IntArray{
//							*pulumi.String(current.Id),
//						},
//					},
//				},
//				DeploymentBranchPolicy: &github.RepositoryEnvironmentDeploymentBranchPolicyArgs{
//					ProtectedBranches:    pulumi.Bool(false),
//					CustomBranchPolicies: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewRepositoryEnvironmentDeploymentPolicy(ctx, "testRepositoryEnvironmentDeploymentPolicy", &github.RepositoryEnvironmentDeploymentPolicyArgs{
//				Repository:    testRepository.Name,
//				Environment:   testRepositoryEnvironment.Environment,
//				BranchPattern: pulumi.String("releases/*"),
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
// GitHub Repository Environment Deployment Policy can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment with the `Id` of the deployment policy, separated by a `:` character, e.g.
//
// ```sh
//
//	$ pulumi import github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy daily terraform:daily:123456
//
// ```
type RepositoryEnvironmentDeploymentPolicy struct {
	pulumi.CustomResourceState

	// The name pattern that branches must match in order to deploy to the environment.
	BranchPattern pulumi.StringOutput `pulumi:"branchPattern"`
	// The name of the environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The repository of the environment.
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewRepositoryEnvironmentDeploymentPolicy registers a new resource with the given unique name, arguments, and options.
func NewRepositoryEnvironmentDeploymentPolicy(ctx *pulumi.Context,
	name string, args *RepositoryEnvironmentDeploymentPolicyArgs, opts ...pulumi.ResourceOption) (*RepositoryEnvironmentDeploymentPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BranchPattern == nil {
		return nil, errors.New("invalid value for required argument 'BranchPattern'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryEnvironmentDeploymentPolicy
	err := ctx.RegisterResource("github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryEnvironmentDeploymentPolicy gets an existing RepositoryEnvironmentDeploymentPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryEnvironmentDeploymentPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryEnvironmentDeploymentPolicyState, opts ...pulumi.ResourceOption) (*RepositoryEnvironmentDeploymentPolicy, error) {
	var resource RepositoryEnvironmentDeploymentPolicy
	err := ctx.ReadResource("github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryEnvironmentDeploymentPolicy resources.
type repositoryEnvironmentDeploymentPolicyState struct {
	// The name pattern that branches must match in order to deploy to the environment.
	BranchPattern *string `pulumi:"branchPattern"`
	// The name of the environment.
	Environment *string `pulumi:"environment"`
	// The repository of the environment.
	Repository *string `pulumi:"repository"`
}

type RepositoryEnvironmentDeploymentPolicyState struct {
	// The name pattern that branches must match in order to deploy to the environment.
	BranchPattern pulumi.StringPtrInput
	// The name of the environment.
	Environment pulumi.StringPtrInput
	// The repository of the environment.
	Repository pulumi.StringPtrInput
}

func (RepositoryEnvironmentDeploymentPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryEnvironmentDeploymentPolicyState)(nil)).Elem()
}

type repositoryEnvironmentDeploymentPolicyArgs struct {
	// The name pattern that branches must match in order to deploy to the environment.
	BranchPattern string `pulumi:"branchPattern"`
	// The name of the environment.
	Environment string `pulumi:"environment"`
	// The repository of the environment.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryEnvironmentDeploymentPolicy resource.
type RepositoryEnvironmentDeploymentPolicyArgs struct {
	// The name pattern that branches must match in order to deploy to the environment.
	BranchPattern pulumi.StringInput
	// The name of the environment.
	Environment pulumi.StringInput
	// The repository of the environment.
	Repository pulumi.StringInput
}

func (RepositoryEnvironmentDeploymentPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryEnvironmentDeploymentPolicyArgs)(nil)).Elem()
}

type RepositoryEnvironmentDeploymentPolicyInput interface {
	pulumi.Input

	ToRepositoryEnvironmentDeploymentPolicyOutput() RepositoryEnvironmentDeploymentPolicyOutput
	ToRepositoryEnvironmentDeploymentPolicyOutputWithContext(ctx context.Context) RepositoryEnvironmentDeploymentPolicyOutput
}

func (*RepositoryEnvironmentDeploymentPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryEnvironmentDeploymentPolicy)(nil)).Elem()
}

func (i *RepositoryEnvironmentDeploymentPolicy) ToRepositoryEnvironmentDeploymentPolicyOutput() RepositoryEnvironmentDeploymentPolicyOutput {
	return i.ToRepositoryEnvironmentDeploymentPolicyOutputWithContext(context.Background())
}

func (i *RepositoryEnvironmentDeploymentPolicy) ToRepositoryEnvironmentDeploymentPolicyOutputWithContext(ctx context.Context) RepositoryEnvironmentDeploymentPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryEnvironmentDeploymentPolicyOutput)
}

func (i *RepositoryEnvironmentDeploymentPolicy) ToOutput(ctx context.Context) pulumix.Output[*RepositoryEnvironmentDeploymentPolicy] {
	return pulumix.Output[*RepositoryEnvironmentDeploymentPolicy]{
		OutputState: i.ToRepositoryEnvironmentDeploymentPolicyOutputWithContext(ctx).OutputState,
	}
}

// RepositoryEnvironmentDeploymentPolicyArrayInput is an input type that accepts RepositoryEnvironmentDeploymentPolicyArray and RepositoryEnvironmentDeploymentPolicyArrayOutput values.
// You can construct a concrete instance of `RepositoryEnvironmentDeploymentPolicyArrayInput` via:
//
//	RepositoryEnvironmentDeploymentPolicyArray{ RepositoryEnvironmentDeploymentPolicyArgs{...} }
type RepositoryEnvironmentDeploymentPolicyArrayInput interface {
	pulumi.Input

	ToRepositoryEnvironmentDeploymentPolicyArrayOutput() RepositoryEnvironmentDeploymentPolicyArrayOutput
	ToRepositoryEnvironmentDeploymentPolicyArrayOutputWithContext(context.Context) RepositoryEnvironmentDeploymentPolicyArrayOutput
}

type RepositoryEnvironmentDeploymentPolicyArray []RepositoryEnvironmentDeploymentPolicyInput

func (RepositoryEnvironmentDeploymentPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryEnvironmentDeploymentPolicy)(nil)).Elem()
}

func (i RepositoryEnvironmentDeploymentPolicyArray) ToRepositoryEnvironmentDeploymentPolicyArrayOutput() RepositoryEnvironmentDeploymentPolicyArrayOutput {
	return i.ToRepositoryEnvironmentDeploymentPolicyArrayOutputWithContext(context.Background())
}

func (i RepositoryEnvironmentDeploymentPolicyArray) ToRepositoryEnvironmentDeploymentPolicyArrayOutputWithContext(ctx context.Context) RepositoryEnvironmentDeploymentPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryEnvironmentDeploymentPolicyArrayOutput)
}

func (i RepositoryEnvironmentDeploymentPolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryEnvironmentDeploymentPolicy] {
	return pulumix.Output[[]*RepositoryEnvironmentDeploymentPolicy]{
		OutputState: i.ToRepositoryEnvironmentDeploymentPolicyArrayOutputWithContext(ctx).OutputState,
	}
}

// RepositoryEnvironmentDeploymentPolicyMapInput is an input type that accepts RepositoryEnvironmentDeploymentPolicyMap and RepositoryEnvironmentDeploymentPolicyMapOutput values.
// You can construct a concrete instance of `RepositoryEnvironmentDeploymentPolicyMapInput` via:
//
//	RepositoryEnvironmentDeploymentPolicyMap{ "key": RepositoryEnvironmentDeploymentPolicyArgs{...} }
type RepositoryEnvironmentDeploymentPolicyMapInput interface {
	pulumi.Input

	ToRepositoryEnvironmentDeploymentPolicyMapOutput() RepositoryEnvironmentDeploymentPolicyMapOutput
	ToRepositoryEnvironmentDeploymentPolicyMapOutputWithContext(context.Context) RepositoryEnvironmentDeploymentPolicyMapOutput
}

type RepositoryEnvironmentDeploymentPolicyMap map[string]RepositoryEnvironmentDeploymentPolicyInput

func (RepositoryEnvironmentDeploymentPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryEnvironmentDeploymentPolicy)(nil)).Elem()
}

func (i RepositoryEnvironmentDeploymentPolicyMap) ToRepositoryEnvironmentDeploymentPolicyMapOutput() RepositoryEnvironmentDeploymentPolicyMapOutput {
	return i.ToRepositoryEnvironmentDeploymentPolicyMapOutputWithContext(context.Background())
}

func (i RepositoryEnvironmentDeploymentPolicyMap) ToRepositoryEnvironmentDeploymentPolicyMapOutputWithContext(ctx context.Context) RepositoryEnvironmentDeploymentPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryEnvironmentDeploymentPolicyMapOutput)
}

func (i RepositoryEnvironmentDeploymentPolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryEnvironmentDeploymentPolicy] {
	return pulumix.Output[map[string]*RepositoryEnvironmentDeploymentPolicy]{
		OutputState: i.ToRepositoryEnvironmentDeploymentPolicyMapOutputWithContext(ctx).OutputState,
	}
}

type RepositoryEnvironmentDeploymentPolicyOutput struct{ *pulumi.OutputState }

func (RepositoryEnvironmentDeploymentPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryEnvironmentDeploymentPolicy)(nil)).Elem()
}

func (o RepositoryEnvironmentDeploymentPolicyOutput) ToRepositoryEnvironmentDeploymentPolicyOutput() RepositoryEnvironmentDeploymentPolicyOutput {
	return o
}

func (o RepositoryEnvironmentDeploymentPolicyOutput) ToRepositoryEnvironmentDeploymentPolicyOutputWithContext(ctx context.Context) RepositoryEnvironmentDeploymentPolicyOutput {
	return o
}

func (o RepositoryEnvironmentDeploymentPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*RepositoryEnvironmentDeploymentPolicy] {
	return pulumix.Output[*RepositoryEnvironmentDeploymentPolicy]{
		OutputState: o.OutputState,
	}
}

// The name pattern that branches must match in order to deploy to the environment.
func (o RepositoryEnvironmentDeploymentPolicyOutput) BranchPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryEnvironmentDeploymentPolicy) pulumi.StringOutput { return v.BranchPattern }).(pulumi.StringOutput)
}

// The name of the environment.
func (o RepositoryEnvironmentDeploymentPolicyOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryEnvironmentDeploymentPolicy) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The repository of the environment.
func (o RepositoryEnvironmentDeploymentPolicyOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryEnvironmentDeploymentPolicy) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

type RepositoryEnvironmentDeploymentPolicyArrayOutput struct{ *pulumi.OutputState }

func (RepositoryEnvironmentDeploymentPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryEnvironmentDeploymentPolicy)(nil)).Elem()
}

func (o RepositoryEnvironmentDeploymentPolicyArrayOutput) ToRepositoryEnvironmentDeploymentPolicyArrayOutput() RepositoryEnvironmentDeploymentPolicyArrayOutput {
	return o
}

func (o RepositoryEnvironmentDeploymentPolicyArrayOutput) ToRepositoryEnvironmentDeploymentPolicyArrayOutputWithContext(ctx context.Context) RepositoryEnvironmentDeploymentPolicyArrayOutput {
	return o
}

func (o RepositoryEnvironmentDeploymentPolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryEnvironmentDeploymentPolicy] {
	return pulumix.Output[[]*RepositoryEnvironmentDeploymentPolicy]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryEnvironmentDeploymentPolicyArrayOutput) Index(i pulumi.IntInput) RepositoryEnvironmentDeploymentPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryEnvironmentDeploymentPolicy {
		return vs[0].([]*RepositoryEnvironmentDeploymentPolicy)[vs[1].(int)]
	}).(RepositoryEnvironmentDeploymentPolicyOutput)
}

type RepositoryEnvironmentDeploymentPolicyMapOutput struct{ *pulumi.OutputState }

func (RepositoryEnvironmentDeploymentPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryEnvironmentDeploymentPolicy)(nil)).Elem()
}

func (o RepositoryEnvironmentDeploymentPolicyMapOutput) ToRepositoryEnvironmentDeploymentPolicyMapOutput() RepositoryEnvironmentDeploymentPolicyMapOutput {
	return o
}

func (o RepositoryEnvironmentDeploymentPolicyMapOutput) ToRepositoryEnvironmentDeploymentPolicyMapOutputWithContext(ctx context.Context) RepositoryEnvironmentDeploymentPolicyMapOutput {
	return o
}

func (o RepositoryEnvironmentDeploymentPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryEnvironmentDeploymentPolicy] {
	return pulumix.Output[map[string]*RepositoryEnvironmentDeploymentPolicy]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryEnvironmentDeploymentPolicyMapOutput) MapIndex(k pulumi.StringInput) RepositoryEnvironmentDeploymentPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryEnvironmentDeploymentPolicy {
		return vs[0].(map[string]*RepositoryEnvironmentDeploymentPolicy)[vs[1].(string)]
	}).(RepositoryEnvironmentDeploymentPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryEnvironmentDeploymentPolicyInput)(nil)).Elem(), &RepositoryEnvironmentDeploymentPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryEnvironmentDeploymentPolicyArrayInput)(nil)).Elem(), RepositoryEnvironmentDeploymentPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryEnvironmentDeploymentPolicyMapInput)(nil)).Elem(), RepositoryEnvironmentDeploymentPolicyMap{})
	pulumi.RegisterOutputType(RepositoryEnvironmentDeploymentPolicyOutput{})
	pulumi.RegisterOutputType(RepositoryEnvironmentDeploymentPolicyArrayOutput{})
	pulumi.RegisterOutputType(RepositoryEnvironmentDeploymentPolicyMapOutput{})
}
