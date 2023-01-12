module github.com/pulumi/pulumi-aws/examples/repo/go/repo-go

go 1.16

require (
	github.com/pulumi/pulumi-github/sdk/v5 v5.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.37.2
)

replace github.com/pulumi/pulumi-github/sdk/v5 => ../../../sdk
