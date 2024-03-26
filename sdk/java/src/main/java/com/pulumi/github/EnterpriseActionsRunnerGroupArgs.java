// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnterpriseActionsRunnerGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final EnterpriseActionsRunnerGroupArgs Empty = new EnterpriseActionsRunnerGroupArgs();

    /**
     * Whether public repositories can be added to the runner group. Defaults to false.
     * 
     */
    @Import(name="allowsPublicRepositories")
    private @Nullable Output<Boolean> allowsPublicRepositories;

    /**
     * @return Whether public repositories can be added to the runner group. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> allowsPublicRepositories() {
        return Optional.ofNullable(this.allowsPublicRepositories);
    }

    /**
     * The slug of the enterprise.
     * 
     */
    @Import(name="enterpriseSlug", required=true)
    private Output<String> enterpriseSlug;

    /**
     * @return The slug of the enterprise.
     * 
     */
    public Output<String> enterpriseSlug() {
        return this.enterpriseSlug;
    }

    /**
     * Name of the runner group
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the runner group
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
     * 
     */
    @Import(name="restrictedToWorkflows")
    private @Nullable Output<Boolean> restrictedToWorkflows;

    /**
     * @return If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> restrictedToWorkflows() {
        return Optional.ofNullable(this.restrictedToWorkflows);
    }

    /**
     * IDs of the organizations which should be added to the runner group
     * 
     */
    @Import(name="selectedOrganizationIds")
    private @Nullable Output<List<Integer>> selectedOrganizationIds;

    /**
     * @return IDs of the organizations which should be added to the runner group
     * 
     */
    public Optional<Output<List<Integer>>> selectedOrganizationIds() {
        return Optional.ofNullable(this.selectedOrganizationIds);
    }

    /**
     * List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
     * 
     */
    @Import(name="selectedWorkflows")
    private @Nullable Output<List<String>> selectedWorkflows;

    /**
     * @return List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
     * 
     */
    public Optional<Output<List<String>>> selectedWorkflows() {
        return Optional.ofNullable(this.selectedWorkflows);
    }

    /**
     * Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
     * 
     */
    @Import(name="visibility", required=true)
    private Output<String> visibility;

    /**
     * @return Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }

    private EnterpriseActionsRunnerGroupArgs() {}

    private EnterpriseActionsRunnerGroupArgs(EnterpriseActionsRunnerGroupArgs $) {
        this.allowsPublicRepositories = $.allowsPublicRepositories;
        this.enterpriseSlug = $.enterpriseSlug;
        this.name = $.name;
        this.restrictedToWorkflows = $.restrictedToWorkflows;
        this.selectedOrganizationIds = $.selectedOrganizationIds;
        this.selectedWorkflows = $.selectedWorkflows;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnterpriseActionsRunnerGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnterpriseActionsRunnerGroupArgs $;

        public Builder() {
            $ = new EnterpriseActionsRunnerGroupArgs();
        }

        public Builder(EnterpriseActionsRunnerGroupArgs defaults) {
            $ = new EnterpriseActionsRunnerGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowsPublicRepositories Whether public repositories can be added to the runner group. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder allowsPublicRepositories(@Nullable Output<Boolean> allowsPublicRepositories) {
            $.allowsPublicRepositories = allowsPublicRepositories;
            return this;
        }

        /**
         * @param allowsPublicRepositories Whether public repositories can be added to the runner group. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder allowsPublicRepositories(Boolean allowsPublicRepositories) {
            return allowsPublicRepositories(Output.of(allowsPublicRepositories));
        }

        /**
         * @param enterpriseSlug The slug of the enterprise.
         * 
         * @return builder
         * 
         */
        public Builder enterpriseSlug(Output<String> enterpriseSlug) {
            $.enterpriseSlug = enterpriseSlug;
            return this;
        }

        /**
         * @param enterpriseSlug The slug of the enterprise.
         * 
         * @return builder
         * 
         */
        public Builder enterpriseSlug(String enterpriseSlug) {
            return enterpriseSlug(Output.of(enterpriseSlug));
        }

        /**
         * @param name Name of the runner group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the runner group
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param restrictedToWorkflows If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder restrictedToWorkflows(@Nullable Output<Boolean> restrictedToWorkflows) {
            $.restrictedToWorkflows = restrictedToWorkflows;
            return this;
        }

        /**
         * @param restrictedToWorkflows If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder restrictedToWorkflows(Boolean restrictedToWorkflows) {
            return restrictedToWorkflows(Output.of(restrictedToWorkflows));
        }

        /**
         * @param selectedOrganizationIds IDs of the organizations which should be added to the runner group
         * 
         * @return builder
         * 
         */
        public Builder selectedOrganizationIds(@Nullable Output<List<Integer>> selectedOrganizationIds) {
            $.selectedOrganizationIds = selectedOrganizationIds;
            return this;
        }

        /**
         * @param selectedOrganizationIds IDs of the organizations which should be added to the runner group
         * 
         * @return builder
         * 
         */
        public Builder selectedOrganizationIds(List<Integer> selectedOrganizationIds) {
            return selectedOrganizationIds(Output.of(selectedOrganizationIds));
        }

        /**
         * @param selectedOrganizationIds IDs of the organizations which should be added to the runner group
         * 
         * @return builder
         * 
         */
        public Builder selectedOrganizationIds(Integer... selectedOrganizationIds) {
            return selectedOrganizationIds(List.of(selectedOrganizationIds));
        }

        /**
         * @param selectedWorkflows List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
         * 
         * @return builder
         * 
         */
        public Builder selectedWorkflows(@Nullable Output<List<String>> selectedWorkflows) {
            $.selectedWorkflows = selectedWorkflows;
            return this;
        }

        /**
         * @param selectedWorkflows List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
         * 
         * @return builder
         * 
         */
        public Builder selectedWorkflows(List<String> selectedWorkflows) {
            return selectedWorkflows(Output.of(selectedWorkflows));
        }

        /**
         * @param selectedWorkflows List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
         * 
         * @return builder
         * 
         */
        public Builder selectedWorkflows(String... selectedWorkflows) {
            return selectedWorkflows(List.of(selectedWorkflows));
        }

        /**
         * @param visibility Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
         * 
         * @return builder
         * 
         */
        public Builder visibility(Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public EnterpriseActionsRunnerGroupArgs build() {
            if ($.enterpriseSlug == null) {
                throw new MissingRequiredPropertyException("EnterpriseActionsRunnerGroupArgs", "enterpriseSlug");
            }
            if ($.visibility == null) {
                throw new MissingRequiredPropertyException("EnterpriseActionsRunnerGroupArgs", "visibility");
            }
            return $;
        }
    }

}