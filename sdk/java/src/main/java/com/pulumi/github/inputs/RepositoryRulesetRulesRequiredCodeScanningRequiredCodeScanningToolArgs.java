// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs Empty = new RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs();

    /**
     * The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
     * 
     */
    @Import(name="alertsThreshold", required=true)
    private Output<String> alertsThreshold;

    /**
     * @return The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
     * 
     */
    public Output<String> alertsThreshold() {
        return this.alertsThreshold;
    }

    /**
     * The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
     * 
     */
    @Import(name="securityAlertsThreshold", required=true)
    private Output<String> securityAlertsThreshold;

    /**
     * @return The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
     * 
     */
    public Output<String> securityAlertsThreshold() {
        return this.securityAlertsThreshold;
    }

    /**
     * The name of a code scanning tool
     * 
     */
    @Import(name="tool", required=true)
    private Output<String> tool;

    /**
     * @return The name of a code scanning tool
     * 
     */
    public Output<String> tool() {
        return this.tool;
    }

    private RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs() {}

    private RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs(RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs $) {
        this.alertsThreshold = $.alertsThreshold;
        this.securityAlertsThreshold = $.securityAlertsThreshold;
        this.tool = $.tool;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs $;

        public Builder() {
            $ = new RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs();
        }

        public Builder(RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs defaults) {
            $ = new RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertsThreshold The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
         * 
         * @return builder
         * 
         */
        public Builder alertsThreshold(Output<String> alertsThreshold) {
            $.alertsThreshold = alertsThreshold;
            return this;
        }

        /**
         * @param alertsThreshold The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
         * 
         * @return builder
         * 
         */
        public Builder alertsThreshold(String alertsThreshold) {
            return alertsThreshold(Output.of(alertsThreshold));
        }

        /**
         * @param securityAlertsThreshold The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
         * 
         * @return builder
         * 
         */
        public Builder securityAlertsThreshold(Output<String> securityAlertsThreshold) {
            $.securityAlertsThreshold = securityAlertsThreshold;
            return this;
        }

        /**
         * @param securityAlertsThreshold The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
         * 
         * @return builder
         * 
         */
        public Builder securityAlertsThreshold(String securityAlertsThreshold) {
            return securityAlertsThreshold(Output.of(securityAlertsThreshold));
        }

        /**
         * @param tool The name of a code scanning tool
         * 
         * @return builder
         * 
         */
        public Builder tool(Output<String> tool) {
            $.tool = tool;
            return this;
        }

        /**
         * @param tool The name of a code scanning tool
         * 
         * @return builder
         * 
         */
        public Builder tool(String tool) {
            return tool(Output.of(tool));
        }

        public RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs build() {
            if ($.alertsThreshold == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs", "alertsThreshold");
            }
            if ($.securityAlertsThreshold == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs", "securityAlertsThreshold");
            }
            if ($.tool == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs", "tool");
            }
            return $;
        }
    }

}
