// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.config.inputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AppAuth {
    /**
     * @return The GitHub App ID.
     * 
     */
    private String id;
    /**
     * @return The GitHub App installation instance ID.
     * 
     */
    private String installationId;
    /**
     * @return The GitHub App PEM file contents.
     * 
     */
    private String pemFile;

    private AppAuth() {}
    /**
     * @return The GitHub App ID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The GitHub App installation instance ID.
     * 
     */
    public String installationId() {
        return this.installationId;
    }
    /**
     * @return The GitHub App PEM file contents.
     * 
     */
    public String pemFile() {
        return this.pemFile;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppAuth defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String installationId;
        private String pemFile;
        public Builder() {}
        public Builder(AppAuth defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.installationId = defaults.installationId;
    	      this.pemFile = defaults.pemFile;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("AppAuth", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder installationId(String installationId) {
            if (installationId == null) {
              throw new MissingRequiredPropertyException("AppAuth", "installationId");
            }
            this.installationId = installationId;
            return this;
        }
        @CustomType.Setter
        public Builder pemFile(String pemFile) {
            if (pemFile == null) {
              throw new MissingRequiredPropertyException("AppAuth", "pemFile");
            }
            this.pemFile = pemFile;
            return this;
        }
        public AppAuth build() {
            final var _resultValue = new AppAuth();
            _resultValue.id = id;
            _resultValue.installationId = installationId;
            _resultValue.pemFile = pemFile;
            return _resultValue;
        }
    }
}
