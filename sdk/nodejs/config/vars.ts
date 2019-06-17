// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("github");

/**
 * The GitHub Base API URL
 */
export let baseUrl: string | undefined = __config.get("baseUrl");
/**
 * Whether server should be accessed without verifying the TLS certificate.
 */
export let insecure: boolean | undefined = __config.getObject<boolean>("insecure");
/**
 * The GitHub organization name to manage.
 */
export let organization: string | undefined = __config.get("organization");
/**
 * The OAuth token used to connect to GitHub.
 */
export let token: string | undefined = __config.get("token");
