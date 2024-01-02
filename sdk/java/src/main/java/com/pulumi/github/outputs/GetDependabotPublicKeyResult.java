// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDependabotPublicKeyResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String key;
    private String keyId;
    private String repository;

    private GetDependabotPublicKeyResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String key() {
        return this.key;
    }
    public String keyId() {
        return this.keyId;
    }
    public String repository() {
        return this.repository;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDependabotPublicKeyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String key;
        private String keyId;
        private String repository;
        public Builder() {}
        public Builder(GetDependabotPublicKeyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.key = defaults.key;
    	      this.keyId = defaults.keyId;
    	      this.repository = defaults.repository;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDependabotPublicKeyResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            if (key == null) {
              throw new MissingRequiredPropertyException("GetDependabotPublicKeyResult", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder keyId(String keyId) {
            if (keyId == null) {
              throw new MissingRequiredPropertyException("GetDependabotPublicKeyResult", "keyId");
            }
            this.keyId = keyId;
            return this;
        }
        @CustomType.Setter
        public Builder repository(String repository) {
            if (repository == null) {
              throw new MissingRequiredPropertyException("GetDependabotPublicKeyResult", "repository");
            }
            this.repository = repository;
            return this;
        }
        public GetDependabotPublicKeyResult build() {
            final var _resultValue = new GetDependabotPublicKeyResult();
            _resultValue.id = id;
            _resultValue.key = key;
            _resultValue.keyId = keyId;
            _resultValue.repository = repository;
            return _resultValue;
        }
    }
}
