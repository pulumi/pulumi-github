// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.inputs.RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs;
import java.util.List;
import java.util.Objects;


public final class RepositoryRulesetRulesRequiredCodeScanningArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryRulesetRulesRequiredCodeScanningArgs Empty = new RepositoryRulesetRulesRequiredCodeScanningArgs();

    /**
     * Tools that must provide code scanning results for this rule to pass.
     * 
     */
    @Import(name="requiredCodeScanningTools", required=true)
    private Output<List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs>> requiredCodeScanningTools;

    /**
     * @return Tools that must provide code scanning results for this rule to pass.
     * 
     */
    public Output<List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs>> requiredCodeScanningTools() {
        return this.requiredCodeScanningTools;
    }

    private RepositoryRulesetRulesRequiredCodeScanningArgs() {}

    private RepositoryRulesetRulesRequiredCodeScanningArgs(RepositoryRulesetRulesRequiredCodeScanningArgs $) {
        this.requiredCodeScanningTools = $.requiredCodeScanningTools;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryRulesetRulesRequiredCodeScanningArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryRulesetRulesRequiredCodeScanningArgs $;

        public Builder() {
            $ = new RepositoryRulesetRulesRequiredCodeScanningArgs();
        }

        public Builder(RepositoryRulesetRulesRequiredCodeScanningArgs defaults) {
            $ = new RepositoryRulesetRulesRequiredCodeScanningArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param requiredCodeScanningTools Tools that must provide code scanning results for this rule to pass.
         * 
         * @return builder
         * 
         */
        public Builder requiredCodeScanningTools(Output<List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs>> requiredCodeScanningTools) {
            $.requiredCodeScanningTools = requiredCodeScanningTools;
            return this;
        }

        /**
         * @param requiredCodeScanningTools Tools that must provide code scanning results for this rule to pass.
         * 
         * @return builder
         * 
         */
        public Builder requiredCodeScanningTools(List<RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs> requiredCodeScanningTools) {
            return requiredCodeScanningTools(Output.of(requiredCodeScanningTools));
        }

        /**
         * @param requiredCodeScanningTools Tools that must provide code scanning results for this rule to pass.
         * 
         * @return builder
         * 
         */
        public Builder requiredCodeScanningTools(RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs... requiredCodeScanningTools) {
            return requiredCodeScanningTools(List.of(requiredCodeScanningTools));
        }

        public RepositoryRulesetRulesRequiredCodeScanningArgs build() {
            if ($.requiredCodeScanningTools == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredCodeScanningArgs", "requiredCodeScanningTools");
            }
            return $;
        }
    }

}
