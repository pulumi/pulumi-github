// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs Empty = new RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs();

    /**
     * (String) The status check context name that must be present on the commit.
     * 
     */
    @Import(name="context", required=true)
    private Output<String> context;

    /**
     * @return (String) The status check context name that must be present on the commit.
     * 
     */
    public Output<String> context() {
        return this.context;
    }

    /**
     * (Number) The optional integration ID that this status check must originate from.
     * 
     */
    @Import(name="integrationId")
    private @Nullable Output<Integer> integrationId;

    /**
     * @return (Number) The optional integration ID that this status check must originate from.
     * 
     */
    public Optional<Output<Integer>> integrationId() {
        return Optional.ofNullable(this.integrationId);
    }

    private RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs() {}

    private RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs(RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs $) {
        this.context = $.context;
        this.integrationId = $.integrationId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs $;

        public Builder() {
            $ = new RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs();
        }

        public Builder(RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs defaults) {
            $ = new RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param context (String) The status check context name that must be present on the commit.
         * 
         * @return builder
         * 
         */
        public Builder context(Output<String> context) {
            $.context = context;
            return this;
        }

        /**
         * @param context (String) The status check context name that must be present on the commit.
         * 
         * @return builder
         * 
         */
        public Builder context(String context) {
            return context(Output.of(context));
        }

        /**
         * @param integrationId (Number) The optional integration ID that this status check must originate from.
         * 
         * @return builder
         * 
         */
        public Builder integrationId(@Nullable Output<Integer> integrationId) {
            $.integrationId = integrationId;
            return this;
        }

        /**
         * @param integrationId (Number) The optional integration ID that this status check must originate from.
         * 
         * @return builder
         * 
         */
        public Builder integrationId(Integer integrationId) {
            return integrationId(Output.of(integrationId));
        }

        public RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs build() {
            if ($.context == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs", "context");
            }
            return $;
        }
    }

}
