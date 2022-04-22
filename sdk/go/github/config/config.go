// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
// and `app_auth` are not set.
func GetAppAuth(ctx *pulumi.Context) string {
	return config.Get(ctx, "github:appAuth")
}

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
//
// Deprecated: Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)
func GetOrganization(ctx *pulumi.Context) string {
	return config.Get(ctx, "github:organization")
}

// The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
func GetOwner(ctx *pulumi.Context) string {
	return config.Get(ctx, "github:owner")
}

// Amount of time in milliseconds to sleep in between non-write requests to GitHub API. Defaults to 0ms if not set.
func GetReadDelayMs(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "github:readDelayMs")
}

// The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "github:token")
}

// Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.
func GetWriteDelayMs(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "github:writeDelayMs")
}
