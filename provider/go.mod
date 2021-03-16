module github.com/pulumi/pulumi-github/provider/v3

go 1.16

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/terraform-providers/terraform-provider-github => github.com/pulumi/terraform-provider-github v0.0.0-20210305152442-c08bc53228f2
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.21.0
	github.com/pulumi/pulumi/sdk/v2 v2.22.1-0.20210310211618-1f16423ede4c
	github.com/terraform-providers/terraform-provider-github v1.3.1-0.20201201134554-f6add3ed515b
)
