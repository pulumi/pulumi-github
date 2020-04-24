module github.com/pulumi/pulumi-github/provider

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/pulumi/pulumi-terraform-bridge/v2 => ../../pulumi-terraform-bridge
)

require (
	github.com/hashicorp/go-getter v1.4.2-0.20200106182914-9813cbd4eb02 // indirect
	github.com/hashicorp/terraform-config-inspect v0.0.0-20191212124732-c6ae6269b9d7 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.1-0.20200422224231-b67c774fccac
	github.com/terraform-providers/terraform-provider-github v1.3.1-0.20200407134903-b514ddb395b8
)
