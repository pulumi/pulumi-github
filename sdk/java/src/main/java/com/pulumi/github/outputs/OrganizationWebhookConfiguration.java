// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrganizationWebhookConfiguration {
    private @Nullable String contentType;
    private @Nullable Boolean insecureSsl;
    private @Nullable String secret;
    /**
     * @return URL of the webhook
     * 
     */
    private String url;

    private OrganizationWebhookConfiguration() {}
    public Optional<String> contentType() {
        return Optional.ofNullable(this.contentType);
    }
    public Optional<Boolean> insecureSsl() {
        return Optional.ofNullable(this.insecureSsl);
    }
    public Optional<String> secret() {
        return Optional.ofNullable(this.secret);
    }
    /**
     * @return URL of the webhook
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationWebhookConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String contentType;
        private @Nullable Boolean insecureSsl;
        private @Nullable String secret;
        private String url;
        public Builder() {}
        public Builder(OrganizationWebhookConfiguration defaults) {
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
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public OrganizationWebhookConfiguration build() {
            final var _resultValue = new OrganizationWebhookConfiguration();
            _resultValue.contentType = contentType;
            _resultValue.insecureSsl = insecureSsl;
            _resultValue.secret = secret;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}
