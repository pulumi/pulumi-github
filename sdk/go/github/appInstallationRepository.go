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

// > **Note**: This resource is not compatible with the GitHub App Installation authentication method.
//
// This resource manages relationships between app installations and repositories
// in your GitHub organization.
//
// Creating this resource installs a particular app on a particular repository.
//
// The app installation and the repository must both belong to the same
// organization on GitHub. Note: you can review your organization's installations
// by the following the instructions at this
// [link](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/reviewing-your-organizations-installed-integrations).
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
//			someRepo, err := github.NewRepository(ctx, "someRepo", nil)
//			if err != nil {
//				return err
//			}
//			_, err = github.NewAppInstallationRepository(ctx, "someAppRepo", &github.AppInstallationRepositoryArgs{
//				InstallationId: pulumi.String("1234567"),
//				Repository:     someRepo.Name,
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
// GitHub App Installation Repository can be imported using an ID made up of `installation_id:repository`, e.g.
//
// ```sh
//
//	$ pulumi import github:index/appInstallationRepository:AppInstallationRepository terraform_repo 1234567:terraform
//
// ```
type AppInstallationRepository struct {
	pulumi.CustomResourceState

	// The GitHub app installation id.
	InstallationId pulumi.StringOutput `pulumi:"installationId"`
	RepoId         pulumi.IntOutput    `pulumi:"repoId"`
	// The repository to install the app on.
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewAppInstallationRepository registers a new resource with the given unique name, arguments, and options.
func NewAppInstallationRepository(ctx *pulumi.Context,
	name string, args *AppInstallationRepositoryArgs, opts ...pulumi.ResourceOption) (*AppInstallationRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstallationId == nil {
		return nil, errors.New("invalid value for required argument 'InstallationId'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppInstallationRepository
	err := ctx.RegisterResource("github:index/appInstallationRepository:AppInstallationRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppInstallationRepository gets an existing AppInstallationRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppInstallationRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppInstallationRepositoryState, opts ...pulumi.ResourceOption) (*AppInstallationRepository, error) {
	var resource AppInstallationRepository
	err := ctx.ReadResource("github:index/appInstallationRepository:AppInstallationRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppInstallationRepository resources.
type appInstallationRepositoryState struct {
	// The GitHub app installation id.
	InstallationId *string `pulumi:"installationId"`
	RepoId         *int    `pulumi:"repoId"`
	// The repository to install the app on.
	Repository *string `pulumi:"repository"`
}

type AppInstallationRepositoryState struct {
	// The GitHub app installation id.
	InstallationId pulumi.StringPtrInput
	RepoId         pulumi.IntPtrInput
	// The repository to install the app on.
	Repository pulumi.StringPtrInput
}

func (AppInstallationRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*appInstallationRepositoryState)(nil)).Elem()
}

type appInstallationRepositoryArgs struct {
	// The GitHub app installation id.
	InstallationId string `pulumi:"installationId"`
	// The repository to install the app on.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a AppInstallationRepository resource.
type AppInstallationRepositoryArgs struct {
	// The GitHub app installation id.
	InstallationId pulumi.StringInput
	// The repository to install the app on.
	Repository pulumi.StringInput
}

func (AppInstallationRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appInstallationRepositoryArgs)(nil)).Elem()
}

type AppInstallationRepositoryInput interface {
	pulumi.Input

	ToAppInstallationRepositoryOutput() AppInstallationRepositoryOutput
	ToAppInstallationRepositoryOutputWithContext(ctx context.Context) AppInstallationRepositoryOutput
}

func (*AppInstallationRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**AppInstallationRepository)(nil)).Elem()
}

func (i *AppInstallationRepository) ToAppInstallationRepositoryOutput() AppInstallationRepositoryOutput {
	return i.ToAppInstallationRepositoryOutputWithContext(context.Background())
}

func (i *AppInstallationRepository) ToAppInstallationRepositoryOutputWithContext(ctx context.Context) AppInstallationRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInstallationRepositoryOutput)
}

// AppInstallationRepositoryArrayInput is an input type that accepts AppInstallationRepositoryArray and AppInstallationRepositoryArrayOutput values.
// You can construct a concrete instance of `AppInstallationRepositoryArrayInput` via:
//
//	AppInstallationRepositoryArray{ AppInstallationRepositoryArgs{...} }
type AppInstallationRepositoryArrayInput interface {
	pulumi.Input

	ToAppInstallationRepositoryArrayOutput() AppInstallationRepositoryArrayOutput
	ToAppInstallationRepositoryArrayOutputWithContext(context.Context) AppInstallationRepositoryArrayOutput
}

type AppInstallationRepositoryArray []AppInstallationRepositoryInput

func (AppInstallationRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppInstallationRepository)(nil)).Elem()
}

func (i AppInstallationRepositoryArray) ToAppInstallationRepositoryArrayOutput() AppInstallationRepositoryArrayOutput {
	return i.ToAppInstallationRepositoryArrayOutputWithContext(context.Background())
}

func (i AppInstallationRepositoryArray) ToAppInstallationRepositoryArrayOutputWithContext(ctx context.Context) AppInstallationRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInstallationRepositoryArrayOutput)
}

// AppInstallationRepositoryMapInput is an input type that accepts AppInstallationRepositoryMap and AppInstallationRepositoryMapOutput values.
// You can construct a concrete instance of `AppInstallationRepositoryMapInput` via:
//
//	AppInstallationRepositoryMap{ "key": AppInstallationRepositoryArgs{...} }
type AppInstallationRepositoryMapInput interface {
	pulumi.Input

	ToAppInstallationRepositoryMapOutput() AppInstallationRepositoryMapOutput
	ToAppInstallationRepositoryMapOutputWithContext(context.Context) AppInstallationRepositoryMapOutput
}

type AppInstallationRepositoryMap map[string]AppInstallationRepositoryInput

func (AppInstallationRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppInstallationRepository)(nil)).Elem()
}

func (i AppInstallationRepositoryMap) ToAppInstallationRepositoryMapOutput() AppInstallationRepositoryMapOutput {
	return i.ToAppInstallationRepositoryMapOutputWithContext(context.Background())
}

func (i AppInstallationRepositoryMap) ToAppInstallationRepositoryMapOutputWithContext(ctx context.Context) AppInstallationRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInstallationRepositoryMapOutput)
}

type AppInstallationRepositoryOutput struct{ *pulumi.OutputState }

func (AppInstallationRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppInstallationRepository)(nil)).Elem()
}

func (o AppInstallationRepositoryOutput) ToAppInstallationRepositoryOutput() AppInstallationRepositoryOutput {
	return o
}

func (o AppInstallationRepositoryOutput) ToAppInstallationRepositoryOutputWithContext(ctx context.Context) AppInstallationRepositoryOutput {
	return o
}

// The GitHub app installation id.
func (o AppInstallationRepositoryOutput) InstallationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppInstallationRepository) pulumi.StringOutput { return v.InstallationId }).(pulumi.StringOutput)
}

func (o AppInstallationRepositoryOutput) RepoId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppInstallationRepository) pulumi.IntOutput { return v.RepoId }).(pulumi.IntOutput)
}

// The repository to install the app on.
func (o AppInstallationRepositoryOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *AppInstallationRepository) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

type AppInstallationRepositoryArrayOutput struct{ *pulumi.OutputState }

func (AppInstallationRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppInstallationRepository)(nil)).Elem()
}

func (o AppInstallationRepositoryArrayOutput) ToAppInstallationRepositoryArrayOutput() AppInstallationRepositoryArrayOutput {
	return o
}

func (o AppInstallationRepositoryArrayOutput) ToAppInstallationRepositoryArrayOutputWithContext(ctx context.Context) AppInstallationRepositoryArrayOutput {
	return o
}

func (o AppInstallationRepositoryArrayOutput) Index(i pulumi.IntInput) AppInstallationRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppInstallationRepository {
		return vs[0].([]*AppInstallationRepository)[vs[1].(int)]
	}).(AppInstallationRepositoryOutput)
}

type AppInstallationRepositoryMapOutput struct{ *pulumi.OutputState }

func (AppInstallationRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppInstallationRepository)(nil)).Elem()
}

func (o AppInstallationRepositoryMapOutput) ToAppInstallationRepositoryMapOutput() AppInstallationRepositoryMapOutput {
	return o
}

func (o AppInstallationRepositoryMapOutput) ToAppInstallationRepositoryMapOutputWithContext(ctx context.Context) AppInstallationRepositoryMapOutput {
	return o
}

func (o AppInstallationRepositoryMapOutput) MapIndex(k pulumi.StringInput) AppInstallationRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppInstallationRepository {
		return vs[0].(map[string]*AppInstallationRepository)[vs[1].(string)]
	}).(AppInstallationRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppInstallationRepositoryInput)(nil)).Elem(), &AppInstallationRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppInstallationRepositoryArrayInput)(nil)).Elem(), AppInstallationRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppInstallationRepositoryMapInput)(nil)).Elem(), AppInstallationRepositoryMap{})
	pulumi.RegisterOutputType(AppInstallationRepositoryOutput{})
	pulumi.RegisterOutputType(AppInstallationRepositoryArrayOutput{})
	pulumi.RegisterOutputType(AppInstallationRepositoryMapOutput{})
}
