// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to retrieve information about a GitHub's IP addresses.
func LookupIpRanges(ctx *pulumi.Context) (*GetIpRangesResult, error) {
	outputs, err := ctx.Invoke("github:index/getIpRanges:getIpRanges", nil)
	if err != nil {
		return nil, err
	}
	return &GetIpRangesResult{
		Gits: outputs["gits"],
		Hooks: outputs["hooks"],
		Importers: outputs["importers"],
		Pages: outputs["pages"],
		Id: outputs["id"],
	}, nil
}

// A collection of values returned by getIpRanges.
type GetIpRangesResult struct {
	// An Array of IP addresses in CIDR format specifying the Git servers.
	Gits interface{}
	// An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
	Hooks interface{}
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
	Importers interface{}
	// An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
	Pages interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
