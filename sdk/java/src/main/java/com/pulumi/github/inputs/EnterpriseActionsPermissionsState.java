// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.EnterpriseActionsPermissionsAllowedActionsConfigArgs;
import com.pulumi.github.inputs.EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnterpriseActionsPermissionsState extends com.pulumi.resources.ResourceArgs {

    public static final EnterpriseActionsPermissionsState Empty = new EnterpriseActionsPermissionsState();

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
     * Sets the actions that are allowed in an enterprise. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    @Import(name="allowedActionsConfig")
    private @Nullable Output<EnterpriseActionsPermissionsAllowedActionsConfigArgs> allowedActionsConfig;

    /**
     * @return Sets the actions that are allowed in an enterprise. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    public Optional<Output<EnterpriseActionsPermissionsAllowedActionsConfigArgs>> allowedActionsConfig() {
        return Optional.ofNullable(this.allowedActionsConfig);
    }

    /**
     * The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     * 
     */
    @Import(name="enabledOrganizations")
    private @Nullable Output<String> enabledOrganizations;

    /**
     * @return The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     * 
     */
    public Optional<Output<String>> enabledOrganizations() {
        return Optional.ofNullable(this.enabledOrganizations);
    }

    /**
     * Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabled_organizations` = `selected`. See Enabled Organizations Config below for details.
     * 
     */
    @Import(name="enabledOrganizationsConfig")
    private @Nullable Output<EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs> enabledOrganizationsConfig;

    /**
     * @return Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabled_organizations` = `selected`. See Enabled Organizations Config below for details.
     * 
     */
    public Optional<Output<EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs>> enabledOrganizationsConfig() {
        return Optional.ofNullable(this.enabledOrganizationsConfig);
    }

    /**
     * The ID of the enterprise.
     * 
     */
    @Import(name="enterpriseId")
    private @Nullable Output<String> enterpriseId;

    /**
     * @return The ID of the enterprise.
     * 
     */
    public Optional<Output<String>> enterpriseId() {
        return Optional.ofNullable(this.enterpriseId);
    }

    private EnterpriseActionsPermissionsState() {}

    private EnterpriseActionsPermissionsState(EnterpriseActionsPermissionsState $) {
        this.allowedActions = $.allowedActions;
        this.allowedActionsConfig = $.allowedActionsConfig;
        this.enabledOrganizations = $.enabledOrganizations;
        this.enabledOrganizationsConfig = $.enabledOrganizationsConfig;
        this.enterpriseId = $.enterpriseId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnterpriseActionsPermissionsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnterpriseActionsPermissionsState $;

        public Builder() {
            $ = new EnterpriseActionsPermissionsState();
        }

        public Builder(EnterpriseActionsPermissionsState defaults) {
            $ = new EnterpriseActionsPermissionsState(Objects.requireNonNull(defaults));
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
         * @param allowedActionsConfig Sets the actions that are allowed in an enterprise. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
         * 
         * @return builder
         * 
         */
        public Builder allowedActionsConfig(@Nullable Output<EnterpriseActionsPermissionsAllowedActionsConfigArgs> allowedActionsConfig) {
            $.allowedActionsConfig = allowedActionsConfig;
            return this;
        }

        /**
         * @param allowedActionsConfig Sets the actions that are allowed in an enterprise. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
         * 
         * @return builder
         * 
         */
        public Builder allowedActionsConfig(EnterpriseActionsPermissionsAllowedActionsConfigArgs allowedActionsConfig) {
            return allowedActionsConfig(Output.of(allowedActionsConfig));
        }

        /**
         * @param enabledOrganizations The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
         * 
         * @return builder
         * 
         */
        public Builder enabledOrganizations(@Nullable Output<String> enabledOrganizations) {
            $.enabledOrganizations = enabledOrganizations;
            return this;
        }

        /**
         * @param enabledOrganizations The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
         * 
         * @return builder
         * 
         */
        public Builder enabledOrganizations(String enabledOrganizations) {
            return enabledOrganizations(Output.of(enabledOrganizations));
        }

        /**
         * @param enabledOrganizationsConfig Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabled_organizations` = `selected`. See Enabled Organizations Config below for details.
         * 
         * @return builder
         * 
         */
        public Builder enabledOrganizationsConfig(@Nullable Output<EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs> enabledOrganizationsConfig) {
            $.enabledOrganizationsConfig = enabledOrganizationsConfig;
            return this;
        }

        /**
         * @param enabledOrganizationsConfig Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabled_organizations` = `selected`. See Enabled Organizations Config below for details.
         * 
         * @return builder
         * 
         */
        public Builder enabledOrganizationsConfig(EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs enabledOrganizationsConfig) {
            return enabledOrganizationsConfig(Output.of(enabledOrganizationsConfig));
        }

        /**
         * @param enterpriseId The ID of the enterprise.
         * 
         * @return builder
         * 
         */
        public Builder enterpriseId(@Nullable Output<String> enterpriseId) {
            $.enterpriseId = enterpriseId;
            return this;
        }

        /**
         * @param enterpriseId The ID of the enterprise.
         * 
         * @return builder
         * 
         */
        public Builder enterpriseId(String enterpriseId) {
            return enterpriseId(Output.of(enterpriseId));
        }

        public EnterpriseActionsPermissionsState build() {
            return $;
        }
    }

}
