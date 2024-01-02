// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class OrganizationRulesetBypassActor {
    /**
     * @return (Number) The ID of the actor that can bypass a ruleset.
     * 
     */
    private Integer actorId;
    /**
     * @return The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
     * 
     */
    private String actorType;
    /**
     * @return (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
     * 
     * ~&gt;Note: at the time of writing this, the following actor types correspond to the following actor IDs:
     * 
     */
    private String bypassMode;

    private OrganizationRulesetBypassActor() {}
    /**
     * @return (Number) The ID of the actor that can bypass a ruleset.
     * 
     */
    public Integer actorId() {
        return this.actorId;
    }
    /**
     * @return The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
     * 
     */
    public String actorType() {
        return this.actorType;
    }
    /**
     * @return (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
     * 
     * ~&gt;Note: at the time of writing this, the following actor types correspond to the following actor IDs:
     * 
     */
    public String bypassMode() {
        return this.bypassMode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationRulesetBypassActor defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer actorId;
        private String actorType;
        private String bypassMode;
        public Builder() {}
        public Builder(OrganizationRulesetBypassActor defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actorId = defaults.actorId;
    	      this.actorType = defaults.actorType;
    	      this.bypassMode = defaults.bypassMode;
        }

        @CustomType.Setter
        public Builder actorId(Integer actorId) {
            if (actorId == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetBypassActor", "actorId");
            }
            this.actorId = actorId;
            return this;
        }
        @CustomType.Setter
        public Builder actorType(String actorType) {
            if (actorType == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetBypassActor", "actorType");
            }
            this.actorType = actorType;
            return this;
        }
        @CustomType.Setter
        public Builder bypassMode(String bypassMode) {
            if (bypassMode == null) {
              throw new MissingRequiredPropertyException("OrganizationRulesetBypassActor", "bypassMode");
            }
            this.bypassMode = bypassMode;
            return this;
        }
        public OrganizationRulesetBypassActor build() {
            final var _resultValue = new OrganizationRulesetBypassActor();
            _resultValue.actorId = actorId;
            _resultValue.actorType = actorType;
            _resultValue.bypassMode = bypassMode;
            return _resultValue;
        }
    }
}
