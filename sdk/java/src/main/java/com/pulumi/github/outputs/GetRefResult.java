// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRefResult {
    /**
     * @return An etag representing the ref.
     * 
     */
    private String etag;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String owner;
    private String ref;
    private String repository;
    /**
     * @return A string storing the reference&#39;s `HEAD` commit&#39;s SHA1.
     * 
     */
    private String sha;

    private GetRefResult() {}
    /**
     * @return An etag representing the ref.
     * 
     */
    public String etag() {
        return this.etag;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> owner() {
        return Optional.ofNullable(this.owner);
    }
    public String ref() {
        return this.ref;
    }
    public String repository() {
        return this.repository;
    }
    /**
     * @return A string storing the reference&#39;s `HEAD` commit&#39;s SHA1.
     * 
     */
    public String sha() {
        return this.sha;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRefResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String etag;
        private String id;
        private @Nullable String owner;
        private String ref;
        private String repository;
        private String sha;
        public Builder() {}
        public Builder(GetRefResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.etag = defaults.etag;
    	      this.id = defaults.id;
    	      this.owner = defaults.owner;
    	      this.ref = defaults.ref;
    	      this.repository = defaults.repository;
    	      this.sha = defaults.sha;
        }

        @CustomType.Setter
        public Builder etag(String etag) {
            if (etag == null) {
              throw new MissingRequiredPropertyException("GetRefResult", "etag");
            }
            this.etag = etag;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRefResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder owner(@Nullable String owner) {

            this.owner = owner;
            return this;
        }
        @CustomType.Setter
        public Builder ref(String ref) {
            if (ref == null) {
              throw new MissingRequiredPropertyException("GetRefResult", "ref");
            }
            this.ref = ref;
            return this;
        }
        @CustomType.Setter
        public Builder repository(String repository) {
            if (repository == null) {
              throw new MissingRequiredPropertyException("GetRefResult", "repository");
            }
            this.repository = repository;
            return this;
        }
        @CustomType.Setter
        public Builder sha(String sha) {
            if (sha == null) {
              throw new MissingRequiredPropertyException("GetRefResult", "sha");
            }
            this.sha = sha;
            return this;
        }
        public GetRefResult build() {
            final var _resultValue = new GetRefResult();
            _resultValue.etag = etag;
            _resultValue.id = id;
            _resultValue.owner = owner;
            _resultValue.ref = ref;
            _resultValue.repository = repository;
            _resultValue.sha = sha;
            return _resultValue;
        }
    }
}
