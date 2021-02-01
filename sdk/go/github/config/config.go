// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// The GitHub Base API URL
func GetBaseUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "github:baseUrl")
	if err == nil {
		return v
	}
	return getEnvOrDefault("https://api.github.com/", nil, "GITHUB_BASE_URL").(string)
}

// Enable `insecure` mode for testing purposes
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "github:insecure")
}

// The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
func GetOrganization(ctx *pulumi.Context) string {
	return config.Get(ctx, "github:organization")
}

// The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
func GetOwner(ctx *pulumi.Context) string {
	return config.Get(ctx, "github:owner")
}

// The OAuth token used to connect to GitHub. `anonymous` mode is enabled if `token` is not configured.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "github:token")
}
