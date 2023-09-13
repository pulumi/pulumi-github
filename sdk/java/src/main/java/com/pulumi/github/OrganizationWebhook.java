// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.OrganizationWebhookArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.OrganizationWebhookState;
import com.pulumi.github.outputs.OrganizationWebhookConfiguration;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage webhooks for GitHub organization.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.OrganizationWebhook;
 * import com.pulumi.github.OrganizationWebhookArgs;
 * import com.pulumi.github.inputs.OrganizationWebhookConfigurationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var foo = new OrganizationWebhook(&#34;foo&#34;, OrganizationWebhookArgs.builder()        
 *             .active(false)
 *             .configuration(OrganizationWebhookConfigurationArgs.builder()
 *                 .contentType(&#34;form&#34;)
 *                 .insecureSsl(false)
 *                 .url(&#34;https://google.de/&#34;)
 *                 .build())
 *             .events(&#34;issues&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Organization webhooks can be imported using the `id` of the webhook. The `id` of the webhook can be found in the URL of the webhook. For example, `&#34;https://github.com/organizations/foo-org/settings/hooks/123456789&#34;`.
 * 
 * ```sh
 *  $ pulumi import github:index/organizationWebhook:OrganizationWebhook terraform 123456789
 * ```
 *  If secret is populated in the webhook&#39;s configuration, the value will be imported as &#34;********&#34;.
 * 
 */
@ResourceType(type="github:index/organizationWebhook:OrganizationWebhook")
public class OrganizationWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     * 
     */
    @Export(name="active", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> active;

    /**
     * @return Indicate of the webhook should receive events. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> active() {
        return Codegen.optional(this.active);
    }
    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
     * 
     */
    @Export(name="configuration", refs={OrganizationWebhookConfiguration.class}, tree="[0]")
    private Output</* @Nullable */ OrganizationWebhookConfiguration> configuration;

    /**
     * @return key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
     * 
     */
    public Output<Optional<OrganizationWebhookConfiguration>> configuration() {
        return Codegen.optional(this.configuration);
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
     * 
     */
    @Export(name="events", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> events;

    /**
     * @return A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
     * 
     */
    public Output<List<String>> events() {
        return this.events;
    }
    /**
     * URL of the webhook
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return URL of the webhook
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrganizationWebhook(String name) {
        this(name, OrganizationWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrganizationWebhook(String name, OrganizationWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrganizationWebhook(String name, OrganizationWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/organizationWebhook:OrganizationWebhook", name, args == null ? OrganizationWebhookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrganizationWebhook(String name, Output<String> id, @Nullable OrganizationWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/organizationWebhook:OrganizationWebhook", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static OrganizationWebhook get(String name, Output<String> id, @Nullable OrganizationWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrganizationWebhook(name, id, state, options);
    }
}
