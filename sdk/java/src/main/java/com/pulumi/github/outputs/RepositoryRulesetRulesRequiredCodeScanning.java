// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool;
import java.util.List;
import java.util.Objects;

@CustomType
public final class RepositoryRulesetRulesRequiredCodeScanning {
    /**
     * @return Tools that must provide code scanning results for this rule to pass.
     * 
     */
    private List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool> requiredCodeScanningTools;

    private RepositoryRulesetRulesRequiredCodeScanning() {}
    /**
     * @return Tools that must provide code scanning results for this rule to pass.
     * 
     */
    public List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool> requiredCodeScanningTools() {
        return this.requiredCodeScanningTools;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryRulesetRulesRequiredCodeScanning defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool> requiredCodeScanningTools;
        public Builder() {}
        public Builder(RepositoryRulesetRulesRequiredCodeScanning defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.requiredCodeScanningTools = defaults.requiredCodeScanningTools;
        }

        @CustomType.Setter
        public Builder requiredCodeScanningTools(List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool> requiredCodeScanningTools) {
            if (requiredCodeScanningTools == null) {
              throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanning", "requiredCodeScanningTools");
            }
            this.requiredCodeScanningTools = requiredCodeScanningTools;
            return this;
        }
        public Builder requiredCodeScanningTools(RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool... requiredCodeScanningTools) {
            return requiredCodeScanningTools(List.of(requiredCodeScanningTools));
        }
        public RepositoryRulesetRulesRequiredCodeScanning build() {
            final var _resultValue = new RepositoryRulesetRulesRequiredCodeScanning();
            _resultValue.requiredCodeScanningTools = requiredCodeScanningTools;
            return _resultValue;
        }
    }
}