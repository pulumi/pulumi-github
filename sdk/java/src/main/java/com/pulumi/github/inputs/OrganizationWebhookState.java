// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.OrganizationWebhookConfigurationArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrganizationWebhookState extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationWebhookState Empty = new OrganizationWebhookState();

    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return Indicate of the webhook should receive events. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
     * 
     */
    @Import(name="configuration")
    private @Nullable Output<OrganizationWebhookConfigurationArgs> configuration;

    /**
     * @return key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
     * 
     */
    public Optional<Output<OrganizationWebhookConfigurationArgs>> configuration() {
        return Optional.ofNullable(this.configuration);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
     * 
     */
    @Import(name="events")
    private @Nullable Output<List<String>> events;

    /**
     * @return A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
     * 
     */
    public Optional<Output<List<String>>> events() {
        return Optional.ofNullable(this.events);
    }

    /**
     * URL of the webhook
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return URL of the webhook
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private OrganizationWebhookState() {}

    private OrganizationWebhookState(OrganizationWebhookState $) {
        this.active = $.active;
        this.configuration = $.configuration;
        this.etag = $.etag;
        this.events = $.events;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationWebhookState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationWebhookState $;

        public Builder() {
            $ = new OrganizationWebhookState();
        }

        public Builder(OrganizationWebhookState defaults) {
            $ = new OrganizationWebhookState(Objects.requireNonNull(defaults));
        }

        /**
         * @param active Indicate of the webhook should receive events. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active Indicate of the webhook should receive events. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param configuration key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
         * 
         * @return builder
         * 
         */
        public Builder configuration(@Nullable Output<OrganizationWebhookConfigurationArgs> configuration) {
            $.configuration = configuration;
            return this;
        }

        /**
         * @param configuration key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
         * 
         * @return builder
         * 
         */
        public Builder configuration(OrganizationWebhookConfigurationArgs configuration) {
            return configuration(Output.of(configuration));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param events A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
         * 
         * @return builder
         * 
         */
        public Builder events(@Nullable Output<List<String>> events) {
            $.events = events;
            return this;
        }

        /**
         * @param events A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
         * 
         * @return builder
         * 
         */
        public Builder events(List<String> events) {
            return events(Output.of(events));
        }

        /**
         * @param events A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
         * 
         * @return builder
         * 
         */
        public Builder events(String... events) {
            return events(List.of(events));
        }

        /**
         * @param url URL of the webhook
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL of the webhook
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public OrganizationWebhookState build() {
            return $;
        }
    }

}
