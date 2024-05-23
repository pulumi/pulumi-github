// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class OrganizationRulesetBypassActorArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationRulesetBypassActorArgs Empty = new OrganizationRulesetBypassActorArgs();

    /**
     * (Number) The ID of the actor that can bypass a ruleset.
     * 
     */
    @Import(name="actorId", required=true)
    private Output<Integer> actorId;

    /**
     * @return (Number) The ID of the actor that can bypass a ruleset.
     * 
     */
    public Output<Integer> actorId() {
        return this.actorId;
    }

    /**
     * The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
     * 
     */
    @Import(name="actorType", required=true)
    private Output<String> actorType;

    /**
     * @return The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
     * 
     */
    public Output<String> actorType() {
        return this.actorType;
    }

    /**
     * (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
     * 
     * ~&gt;Note: at the time of writing this, the following actor types correspond to the following actor IDs:
     * 
     * * `OrganizationAdmin` &gt; `1`
     * * `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
     * 
     */
    @Import(name="bypassMode", required=true)
    private Output<String> bypassMode;

    /**
     * @return (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
     * 
     * ~&gt;Note: at the time of writing this, the following actor types correspond to the following actor IDs:
     * 
     * * `OrganizationAdmin` &gt; `1`
     * * `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
     * 
     */
    public Output<String> bypassMode() {
        return this.bypassMode;
    }

    private OrganizationRulesetBypassActorArgs() {}

    private OrganizationRulesetBypassActorArgs(OrganizationRulesetBypassActorArgs $) {
        this.actorId = $.actorId;
        this.actorType = $.actorType;
        this.bypassMode = $.bypassMode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationRulesetBypassActorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationRulesetBypassActorArgs $;

        public Builder() {
            $ = new OrganizationRulesetBypassActorArgs();
        }

        public Builder(OrganizationRulesetBypassActorArgs defaults) {
            $ = new OrganizationRulesetBypassActorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param actorId (Number) The ID of the actor that can bypass a ruleset.
         * 
         * @return builder
         * 
         */
        public Builder actorId(Output<Integer> actorId) {
            $.actorId = actorId;
            return this;
        }

        /**
         * @param actorId (Number) The ID of the actor that can bypass a ruleset.
         * 
         * @return builder
         * 
         */
        public Builder actorId(Integer actorId) {
            return actorId(Output.of(actorId));
        }

        /**
         * @param actorType The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
         * 
         * @return builder
         * 
         */
        public Builder actorType(Output<String> actorType) {
            $.actorType = actorType;
            return this;
        }

        /**
         * @param actorType The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
         * 
         * @return builder
         * 
         */
        public Builder actorType(String actorType) {
            return actorType(Output.of(actorType));
        }

        /**
         * @param bypassMode (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
         * 
         * ~&gt;Note: at the time of writing this, the following actor types correspond to the following actor IDs:
         * 
         * * `OrganizationAdmin` &gt; `1`
         * * `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
         * 
         * @return builder
         * 
         */
        public Builder bypassMode(Output<String> bypassMode) {
            $.bypassMode = bypassMode;
            return this;
        }

        /**
         * @param bypassMode (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
         * 
         * ~&gt;Note: at the time of writing this, the following actor types correspond to the following actor IDs:
         * 
         * * `OrganizationAdmin` &gt; `1`
         * * `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
         * 
         * @return builder
         * 
         */
        public Builder bypassMode(String bypassMode) {
            return bypassMode(Output.of(bypassMode));
        }

        public OrganizationRulesetBypassActorArgs build() {
            if ($.actorId == null) {
                throw new MissingRequiredPropertyException("OrganizationRulesetBypassActorArgs", "actorId");
            }
            if ($.actorType == null) {
                throw new MissingRequiredPropertyException("OrganizationRulesetBypassActorArgs", "actorType");
            }
            if ($.bypassMode == null) {
                throw new MissingRequiredPropertyException("OrganizationRulesetBypassActorArgs", "bypassMode");
            }
            return $;
        }
    }

}
