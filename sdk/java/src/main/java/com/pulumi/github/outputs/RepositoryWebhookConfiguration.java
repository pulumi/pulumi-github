// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RepositoryWebhookConfiguration {
    /**
     * @return The content type for the payload. Valid values are either `form` or `json`.
     * 
     */
    private @Nullable String contentType;
    /**
     * @return Insecure SSL boolean toggle. Defaults to `false`.
     * 
     */
    private @Nullable Boolean insecureSsl;
    /**
     * @return The shared secret for the webhook. [See API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).
     * 
     */
    private @Nullable String secret;
    /**
     * @return The URL of the webhook.
     * 
     */
    private String url;

    private RepositoryWebhookConfiguration() {}
    /**
     * @return The content type for the payload. Valid values are either `form` or `json`.
     * 
     */
    public Optional<String> contentType() {
        return Optional.ofNullable(this.contentType);
    }
    /**
     * @return Insecure SSL boolean toggle. Defaults to `false`.
     * 
     */
    public Optional<Boolean> insecureSsl() {
        return Optional.ofNullable(this.insecureSsl);
    }
    /**
     * @return The shared secret for the webhook. [See API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).
     * 
     */
    public Optional<String> secret() {
        return Optional.ofNullable(this.secret);
    }
    /**
     * @return The URL of the webhook.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryWebhookConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String contentType;
        private @Nullable Boolean insecureSsl;
        private @Nullable String secret;
        private String url;
        public Builder() {}
        public Builder(RepositoryWebhookConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contentType = defaults.contentType;
    	      this.insecureSsl = defaults.insecureSsl;
    	      this.secret = defaults.secret;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder contentType(@Nullable String contentType) {

            this.contentType = contentType;
            return this;
        }
        @CustomType.Setter
        public Builder insecureSsl(@Nullable Boolean insecureSsl) {

            this.insecureSsl = insecureSsl;
            return this;
        }
        @CustomType.Setter
        public Builder secret(@Nullable String secret) {

            this.secret = secret;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("RepositoryWebhookConfiguration", "url");
            }
            this.url = url;
            return this;
        }
        public RepositoryWebhookConfiguration build() {
            final var _resultValue = new RepositoryWebhookConfiguration();
            _resultValue.contentType = contentType;
            _resultValue.insecureSsl = insecureSsl;
            _resultValue.secret = secret;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}
