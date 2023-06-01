// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryAutolinkReferencesAutolinkReference {
    /**
     * @return True if alphanumeric.
     * 
     */
    private Boolean isAlphanumeric;
    /**
     * @return Key prefix.
     * 
     */
    private String keyPrefix;
    /**
     * @return Target url template.
     * 
     */
    private String targetUrlTemplate;

    private GetRepositoryAutolinkReferencesAutolinkReference() {}
    /**
     * @return True if alphanumeric.
     * 
     */
    public Boolean isAlphanumeric() {
        return this.isAlphanumeric;
    }
    /**
     * @return Key prefix.
     * 
     */
    public String keyPrefix() {
        return this.keyPrefix;
    }
    /**
     * @return Target url template.
     * 
     */
    public String targetUrlTemplate() {
        return this.targetUrlTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryAutolinkReferencesAutolinkReference defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean isAlphanumeric;
        private String keyPrefix;
        private String targetUrlTemplate;
        public Builder() {}
        public Builder(GetRepositoryAutolinkReferencesAutolinkReference defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.isAlphanumeric = defaults.isAlphanumeric;
    	      this.keyPrefix = defaults.keyPrefix;
    	      this.targetUrlTemplate = defaults.targetUrlTemplate;
        }

        @CustomType.Setter
        public Builder isAlphanumeric(Boolean isAlphanumeric) {
            this.isAlphanumeric = Objects.requireNonNull(isAlphanumeric);
            return this;
        }
        @CustomType.Setter
        public Builder keyPrefix(String keyPrefix) {
            this.keyPrefix = Objects.requireNonNull(keyPrefix);
            return this;
        }
        @CustomType.Setter
        public Builder targetUrlTemplate(String targetUrlTemplate) {
            this.targetUrlTemplate = Objects.requireNonNull(targetUrlTemplate);
            return this;
        }
        public GetRepositoryAutolinkReferencesAutolinkReference build() {
            final var o = new GetRepositoryAutolinkReferencesAutolinkReference();
            o.isAlphanumeric = isAlphanumeric;
            o.keyPrefix = keyPrefix;
            o.targetUrlTemplate = targetUrlTemplate;
            return o;
        }
    }
}
