// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetActionsSecretsSecret {
    /**
     * @return Timestamp of the secret creation
     * 
     */
    private String createdAt;
    /**
     * @return The name of the repository.
     * 
     */
    private String name;
    /**
     * @return Timestamp of the secret last update
     * 
     */
    private String updatedAt;

    private GetActionsSecretsSecret() {}
    /**
     * @return Timestamp of the secret creation
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The name of the repository.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Timestamp of the secret last update
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetActionsSecretsSecret defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String name;
        private String updatedAt;
        public Builder() {}
        public Builder(GetActionsSecretsSecret defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.name = defaults.name;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetActionsSecretsSecret", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetActionsSecretsSecret", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetActionsSecretsSecret", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        public GetActionsSecretsSecret build() {
            final var _resultValue = new GetActionsSecretsSecret();
            _resultValue.createdAt = createdAt;
            _resultValue.name = name;
            _resultValue.updatedAt = updatedAt;
            return _resultValue;
        }
    }
}
