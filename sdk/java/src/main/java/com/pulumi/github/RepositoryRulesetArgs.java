// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.inputs.RepositoryRulesetBypassActorArgs;
import com.pulumi.github.inputs.RepositoryRulesetConditionsArgs;
import com.pulumi.github.inputs.RepositoryRulesetRulesArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryRulesetArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryRulesetArgs Empty = new RepositoryRulesetArgs();

    /**
     * (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
     * 
     */
    @Import(name="bypassActors")
    private @Nullable Output<List<RepositoryRulesetBypassActorArgs>> bypassActors;

    /**
     * @return (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
     * 
     */
    public Optional<Output<List<RepositoryRulesetBypassActorArgs>>> bypassActors() {
        return Optional.ofNullable(this.bypassActors);
    }

    /**
     * (Block List, Max: 1) Parameters for a repository ruleset ref name condition. (see below for nested schema)
     * 
     */
    @Import(name="conditions")
    private @Nullable Output<RepositoryRulesetConditionsArgs> conditions;

    /**
     * @return (Block List, Max: 1) Parameters for a repository ruleset ref name condition. (see below for nested schema)
     * 
     */
    public Optional<Output<RepositoryRulesetConditionsArgs>> conditions() {
        return Optional.ofNullable(this.conditions);
    }

    /**
     * (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
     * 
     */
    @Import(name="enforcement", required=true)
    private Output<String> enforcement;

    /**
     * @return (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
     * 
     */
    public Output<String> enforcement() {
        return this.enforcement;
    }

    /**
     * (String) The name of the ruleset.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return (String) The name of the ruleset.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * (String) Name of the repository to apply rulset to.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return (String) Name of the repository to apply rulset to.
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
     * 
     */
    @Import(name="rules", required=true)
    private Output<RepositoryRulesetRulesArgs> rules;

    /**
     * @return (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
     * 
     */
    public Output<RepositoryRulesetRulesArgs> rules() {
        return this.rules;
    }

    /**
     * (String) Possible values are `branch` and `tag`.
     * 
     */
    @Import(name="target", required=true)
    private Output<String> target;

    /**
     * @return (String) Possible values are `branch` and `tag`.
     * 
     */
    public Output<String> target() {
        return this.target;
    }

    private RepositoryRulesetArgs() {}

    private RepositoryRulesetArgs(RepositoryRulesetArgs $) {
        this.bypassActors = $.bypassActors;
        this.conditions = $.conditions;
        this.enforcement = $.enforcement;
        this.name = $.name;
        this.repository = $.repository;
        this.rules = $.rules;
        this.target = $.target;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryRulesetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryRulesetArgs $;

        public Builder() {
            $ = new RepositoryRulesetArgs();
        }

        public Builder(RepositoryRulesetArgs defaults) {
            $ = new RepositoryRulesetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bypassActors (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder bypassActors(@Nullable Output<List<RepositoryRulesetBypassActorArgs>> bypassActors) {
            $.bypassActors = bypassActors;
            return this;
        }

        /**
         * @param bypassActors (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder bypassActors(List<RepositoryRulesetBypassActorArgs> bypassActors) {
            return bypassActors(Output.of(bypassActors));
        }

        /**
         * @param bypassActors (Block List) The actors that can bypass the rules in this ruleset. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder bypassActors(RepositoryRulesetBypassActorArgs... bypassActors) {
            return bypassActors(List.of(bypassActors));
        }

        /**
         * @param conditions (Block List, Max: 1) Parameters for a repository ruleset ref name condition. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder conditions(@Nullable Output<RepositoryRulesetConditionsArgs> conditions) {
            $.conditions = conditions;
            return this;
        }

        /**
         * @param conditions (Block List, Max: 1) Parameters for a repository ruleset ref name condition. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder conditions(RepositoryRulesetConditionsArgs conditions) {
            return conditions(Output.of(conditions));
        }

        /**
         * @param enforcement (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
         * 
         * @return builder
         * 
         */
        public Builder enforcement(Output<String> enforcement) {
            $.enforcement = enforcement;
            return this;
        }

        /**
         * @param enforcement (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
         * 
         * @return builder
         * 
         */
        public Builder enforcement(String enforcement) {
            return enforcement(Output.of(enforcement));
        }

        /**
         * @param name (String) The name of the ruleset.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name (String) The name of the ruleset.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param repository (String) Name of the repository to apply rulset to.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository (String) Name of the repository to apply rulset to.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param rules (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder rules(Output<RepositoryRulesetRulesArgs> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules (Block List, Min: 1, Max: 1) Rules within the ruleset. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder rules(RepositoryRulesetRulesArgs rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param target (String) Possible values are `branch` and `tag`.
         * 
         * @return builder
         * 
         */
        public Builder target(Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target (String) Possible values are `branch` and `tag`.
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        public RepositoryRulesetArgs build() {
            if ($.enforcement == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetArgs", "enforcement");
            }
            if ($.rules == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetArgs", "rules");
            }
            if ($.target == null) {
                throw new MissingRequiredPropertyException("RepositoryRulesetArgs", "target");
            }
            return $;
        }
    }

}
