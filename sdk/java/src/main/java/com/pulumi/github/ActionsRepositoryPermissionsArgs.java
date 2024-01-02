// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.inputs.ActionsRepositoryPermissionsAllowedActionsConfigArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ActionsRepositoryPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActionsRepositoryPermissionsArgs Empty = new ActionsRepositoryPermissionsArgs();

    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
     * 
     */
    @Import(name="allowedActions")
    private @Nullable Output<String> allowedActions;

    /**
     * @return The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
     * 
     */
    public Optional<Output<String>> allowedActions() {
        return Optional.ofNullable(this.allowedActions);
    }

    /**
     * Sets the actions that are allowed in an repository. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    @Import(name="allowedActionsConfig")
    private @Nullable Output<ActionsRepositoryPermissionsAllowedActionsConfigArgs> allowedActionsConfig;

    /**
     * @return Sets the actions that are allowed in an repository. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    public Optional<Output<ActionsRepositoryPermissionsAllowedActionsConfigArgs>> allowedActionsConfig() {
        return Optional.ofNullable(this.allowedActionsConfig);
    }

    /**
     * Should GitHub actions be enabled on this repository?
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Should GitHub actions be enabled on this repository?
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The GitHub repository
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The GitHub repository
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private ActionsRepositoryPermissionsArgs() {}

    private ActionsRepositoryPermissionsArgs(ActionsRepositoryPermissionsArgs $) {
        this.allowedActions = $.allowedActions;
        this.allowedActionsConfig = $.allowedActionsConfig;
        this.enabled = $.enabled;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActionsRepositoryPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActionsRepositoryPermissionsArgs $;

        public Builder() {
            $ = new ActionsRepositoryPermissionsArgs();
        }

        public Builder(ActionsRepositoryPermissionsArgs defaults) {
            $ = new ActionsRepositoryPermissionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedActions The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
         * 
         * @return builder
         * 
         */
        public Builder allowedActions(@Nullable Output<String> allowedActions) {
            $.allowedActions = allowedActions;
            return this;
        }

        /**
         * @param allowedActions The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
         * 
         * @return builder
         * 
         */
        public Builder allowedActions(String allowedActions) {
            return allowedActions(Output.of(allowedActions));
        }

        /**
         * @param allowedActionsConfig Sets the actions that are allowed in an repository. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
         * 
         * @return builder
         * 
         */
        public Builder allowedActionsConfig(@Nullable Output<ActionsRepositoryPermissionsAllowedActionsConfigArgs> allowedActionsConfig) {
            $.allowedActionsConfig = allowedActionsConfig;
            return this;
        }

        /**
         * @param allowedActionsConfig Sets the actions that are allowed in an repository. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
         * 
         * @return builder
         * 
         */
        public Builder allowedActionsConfig(ActionsRepositoryPermissionsAllowedActionsConfigArgs allowedActionsConfig) {
            return allowedActionsConfig(Output.of(allowedActionsConfig));
        }

        /**
         * @param enabled Should GitHub actions be enabled on this repository?
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Should GitHub actions be enabled on this repository?
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param repository The GitHub repository
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The GitHub repository
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public ActionsRepositoryPermissionsArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("ActionsRepositoryPermissionsArgs", "repository");
            }
            return $;
        }
    }

}
