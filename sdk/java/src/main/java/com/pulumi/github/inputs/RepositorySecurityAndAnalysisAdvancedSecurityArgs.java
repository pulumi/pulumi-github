// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class RepositorySecurityAndAnalysisAdvancedSecurityArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositorySecurityAndAnalysisAdvancedSecurityArgs Empty = new RepositorySecurityAndAnalysisAdvancedSecurityArgs();

    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     * 
     */
    @Import(name="status", required=true)
    private Output<String> status;

    /**
     * @return Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    private RepositorySecurityAndAnalysisAdvancedSecurityArgs() {}

    private RepositorySecurityAndAnalysisAdvancedSecurityArgs(RepositorySecurityAndAnalysisAdvancedSecurityArgs $) {
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositorySecurityAndAnalysisAdvancedSecurityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositorySecurityAndAnalysisAdvancedSecurityArgs $;

        public Builder() {
            $ = new RepositorySecurityAndAnalysisAdvancedSecurityArgs();
        }

        public Builder(RepositorySecurityAndAnalysisAdvancedSecurityArgs defaults) {
            $ = new RepositorySecurityAndAnalysisAdvancedSecurityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param status Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
         * 
         * @return builder
         * 
         */
        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public RepositorySecurityAndAnalysisAdvancedSecurityArgs build() {
            if ($.status == null) {
                throw new MissingRequiredPropertyException("RepositorySecurityAndAnalysisAdvancedSecurityArgs", "status");
            }
            return $;
        }
    }

}
