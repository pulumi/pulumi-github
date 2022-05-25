module github.com/pulumi/pulumi-github/provider/v4

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/integrations/terraform-provider-github/v4 v4.25.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.21.0
	github.com/pulumi/pulumi/sdk/v3 v3.30.0
)

//replace github.com/terraform-providers/terraform-provider-github => github.com/pulumi/terraform-provider-github upstream-v4.25.0
