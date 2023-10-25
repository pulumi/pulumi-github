// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the identity provider (IdP) groups for an organization.
 */
export function getOrganizationTeamSyncGroups(opts?: pulumi.InvokeOptions): Promise<GetOrganizationTeamSyncGroupsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getOrganizationTeamSyncGroups:getOrganizationTeamSyncGroups", {
    }, opts);
}

/**
 * A collection of values returned by getOrganizationTeamSyncGroups.
 */
export interface GetOrganizationTeamSyncGroupsResult {
    /**
     * An Array of GitHub Identity Provider Groups.  Each `group` block consists of the fields documented below.
     */
    readonly groups: outputs.GetOrganizationTeamSyncGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * Use this data source to retrieve the identity provider (IdP) groups for an organization.
 */
export function getOrganizationTeamSyncGroupsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationTeamSyncGroupsResult> {
    return pulumi.output(getOrganizationTeamSyncGroups(opts))
}
