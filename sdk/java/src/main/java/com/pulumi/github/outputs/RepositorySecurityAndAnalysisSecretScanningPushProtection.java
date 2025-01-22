// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class RepositorySecurityAndAnalysisSecretScanningPushProtection {
    /**
     * @return The GitHub Pages site&#39;s build status e.g. `building` or `built`.
     * 
     */
    private String status;

    private RepositorySecurityAndAnalysisSecretScanningPushProtection() {}
    /**
     * @return The GitHub Pages site&#39;s build status e.g. `building` or `built`.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositorySecurityAndAnalysisSecretScanningPushProtection defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String status;
        public Builder() {}
        public Builder(RepositorySecurityAndAnalysisSecretScanningPushProtection defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("RepositorySecurityAndAnalysisSecretScanningPushProtection", "status");
            }
            this.status = status;
            return this;
        }
        public RepositorySecurityAndAnalysisSecretScanningPushProtection build() {
            final var _resultValue = new RepositorySecurityAndAnalysisSecretScanningPushProtection();
            _resultValue.status = status;
            return _resultValue;
        }
    }
}