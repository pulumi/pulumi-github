module github.com/pulumi/pulumi-github/provider/v4

go 1.16

replace github.com/terraform-providers/terraform-provider-github => github.com/pulumi/terraform-provider-github v0.0.0-20211019150943-8b8f01ed279f

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.3.2-0.20210526172205-85142462c7ed
	github.com/terraform-providers/terraform-provider-github v1.3.1-0.20201201134554-f6add3ed515b
)
