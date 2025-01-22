// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool {
    /**
     * @return The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
     * 
     */
    private String alertsThreshold;
    /**
     * @return The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
     * 
     */
    private String securityAlertsThreshold;
    /**
     * @return The name of a code scanning tool
     * 
     */
    private String tool;

    private RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool() {}
    /**
     * @return The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
     * 
     */
    public String alertsThreshold() {
        return this.alertsThreshold;
    }
    /**
     * @return The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
     * 
     */
    public String securityAlertsThreshold() {
        return this.securityAlertsThreshold;
    }
    /**
     * @return The name of a code scanning tool
     * 
     */
    public String tool() {
        return this.tool;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String alertsThreshold;
        private String securityAlertsThreshold;
        private String tool;
        public Builder() {}
        public Builder(RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alertsThreshold = defaults.alertsThreshold;
    	      this.securityAlertsThreshold = defaults.securityAlertsThreshold;
    	      this.tool = defaults.tool;
        }

        @CustomType.Setter
        public Builder alertsThreshold(String alertsThreshold) {
            if (alertsThreshold == null) {
              throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool", "alertsThreshold");
            }
            this.alertsThreshold = alertsThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder securityAlertsThreshold(String securityAlertsThreshold) {
            if (securityAlertsThreshold == null) {
              throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool", "securityAlertsThreshold");
            }
            this.securityAlertsThreshold = securityAlertsThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder tool(String tool) {
            if (tool == null) {
              throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool", "tool");
            }
            this.tool = tool;
            return this;
        }
        public RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool build() {
            final var _resultValue = new RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool();
            _resultValue.alertsThreshold = alertsThreshold;
            _resultValue.securityAlertsThreshold = securityAlertsThreshold;
            _resultValue.tool = tool;
            return _resultValue;
        }
    }
}