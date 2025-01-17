// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.inputs.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryRulesetRulesRequiredStatusChecksArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryRulesetRulesRequiredStatusChecksArgs Empty = new RepositoryRulesetRulesRequiredStatusChecksArgs();

    /**
     * Allow repositories and branches to be created if a check would otherwise prohibit it.
     * 
     */
    @Import(name="doNotEnforceOnCreate")
    private @Nullable Output<Boolean> doNotEnforceOnCreate;

    /**
     * @return Allow repositories and branches to be created if a check would otherwise prohibit it.
     * 
     */
    public Optional<Output<Boolean>> doNotEnforceOnCreate() {
        return Optional.ofNullable(this.doNotEnforceOnCreate);
    }

    /**
     * Status checks that are required. Several can be defined.
     * 
     */
    @Import(name="requiredChecks", required=true)
    private Output<List<RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs>> requiredChecks;

    /**
     * @return Status checks that are required. Several can be defined.
     * 
     */
    public Output<List<RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs>> requiredChecks() {
        return this.requiredChecks;
    }

    /**
     * Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
     * 
     */
    @Import(name="strictRequiredStatusChecksPolicy")
    private @Nullable Output<Boolean> strictRequiredStatusChecksPolicy;

    /**
     * @return Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> strictRequiredStatusChecksPolicy() {
        return Optional.ofNullable(this.strictRequiredStatusChecksPolicy);
    }

    private RepositoryRulesetRulesRequiredStatusChecksArgs() {}

    private RepositoryRulesetRulesRequiredStatusChecksArgs(RepositoryRulesetRulesRequiredStatusChecksArgs $) {
        this.doNotEnforceOnCreate = $.doNotEnforceOnCreate;
        this.requiredChecks = $.requiredChecks;
        this.strictRequiredStatusChecksPolicy = $.strictRequiredStatusChecksPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryRulesetRulesRequiredStatusChecksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryRulesetRulesRequiredStatusChecksArgs $;

        public Builder() {
            $ = new RepositoryRulesetRulesRequiredStatusChecksArgs();
        }

        public Builder(RepositoryRulesetRulesRequiredStatusChecksArgs defaults) {
            $ = new RepositoryRulesetRulesRequiredStatusChecksArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param doNotEnforceOnCreate Allow repositories and branches to be created if a check would otherwise prohibit it.
         * 
         * @return builder
         * 
         */
        public Builder doNotEnforceOnCreate(@Nullable Output<Boolean> doNotEnforceOnCreate) {
            $.doNotEnforceOnCreate = doNotEnforceOnCreate;
            return this;
        }

        /**
         * @param doNotEnforceOnCreate Allow repositories and branches to be created if a check would otherwise prohibit it.
         * 
         * @return builder
         * 
         */
        public Builder doNotEnforceOnCreate(Boolean doNotEnforceOnCreate) {
            return doNotEnforceOnCreate(Output.of(doNotEnforceOnCreate));
        }

        /**
         * @param requiredChecks Status checks that are required. Several can be defined.
         * 
         * @return builder
         * 
         */
        public Builder requiredChecks(Output<List<RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs>> requiredChecks) {
            $.requiredChecks = requiredChecks;
            return this;
        }

        /**
         * @param requiredChecks Status checks that are required. Several can be defined.
         * 
         * @return builder
         * 
         */
        public Builder requiredChecks(List<RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs> requiredChecks) {
            return requiredChecks(Output.of(requiredChecks));
        }

        /**
         * @param requiredChecks Status checks that are required. Several can be defined.
         * 
         * @return builder
         * 
         */
        public Builder requiredChecks(RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs... requiredChecks) {
            return requiredChecks(List.of(requiredChecks));
        }

        /**
         * @param strictRequiredStatusChecksPolicy Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder strictRequiredStatusChecksPolicy(@Nullable Output<Boolean> strictRequiredStatusChecksPolicy) {
            $.strictRequiredStatusChecksPolicy = strictRequiredStatusChecksPolicy;
            return this;
        }

        /**
         * @param strictRequiredStatusChecksPolicy Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder strictRequiredStatusChecksPolicy(Boolean strictRequiredStatusChecksPolicy) {
            return strictRequiredStatusChecksPolicy(Output.of(strictRequiredStatusChecksPolicy));
        }

        public RepositoryRulesetRulesRequiredStatusChecksArgs build() {
            if ($.requiredChecks == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredStatusChecksArgs", "requiredChecks");
            }
            return $;
        }
    }

}
