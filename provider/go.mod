module github.com/pulumi/pulumi-github/provider/v4

go 1.16

replace github.com/terraform-providers/terraform-provider-github => github.com/pulumi/terraform-provider-github v0.0.0-20220321192532-b3054b587e7f

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.27.0
	github.com/terraform-providers/terraform-provider-github v1.3.1-0.20201201134554-f6add3ed515b
)
