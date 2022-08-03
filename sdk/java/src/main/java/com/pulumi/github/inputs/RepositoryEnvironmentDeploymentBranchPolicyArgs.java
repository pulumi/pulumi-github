// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;


public final class RepositoryEnvironmentDeploymentBranchPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryEnvironmentDeploymentBranchPolicyArgs Empty = new RepositoryEnvironmentDeploymentBranchPolicyArgs();

    /**
     * Whether only branches that match the specified name patterns can deploy to this environment.
     * 
     */
    @Import(name="customBranchPolicies", required=true)
    private Output<Boolean> customBranchPolicies;

    /**
     * @return Whether only branches that match the specified name patterns can deploy to this environment.
     * 
     */
    public Output<Boolean> customBranchPolicies() {
        return this.customBranchPolicies;
    }

    /**
     * Whether only branches with branch protection rules can deploy to this environment.
     * 
     */
    @Import(name="protectedBranches", required=true)
    private Output<Boolean> protectedBranches;

    /**
     * @return Whether only branches with branch protection rules can deploy to this environment.
     * 
     */
    public Output<Boolean> protectedBranches() {
        return this.protectedBranches;
    }

    private RepositoryEnvironmentDeploymentBranchPolicyArgs() {}

    private RepositoryEnvironmentDeploymentBranchPolicyArgs(RepositoryEnvironmentDeploymentBranchPolicyArgs $) {
        this.customBranchPolicies = $.customBranchPolicies;
        this.protectedBranches = $.protectedBranches;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryEnvironmentDeploymentBranchPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryEnvironmentDeploymentBranchPolicyArgs $;

        public Builder() {
            $ = new RepositoryEnvironmentDeploymentBranchPolicyArgs();
        }

        public Builder(RepositoryEnvironmentDeploymentBranchPolicyArgs defaults) {
            $ = new RepositoryEnvironmentDeploymentBranchPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customBranchPolicies Whether only branches that match the specified name patterns can deploy to this environment.
         * 
         * @return builder
         * 
         */
        public Builder customBranchPolicies(Output<Boolean> customBranchPolicies) {
            $.customBranchPolicies = customBranchPolicies;
            return this;
        }

        /**
         * @param customBranchPolicies Whether only branches that match the specified name patterns can deploy to this environment.
         * 
         * @return builder
         * 
         */
        public Builder customBranchPolicies(Boolean customBranchPolicies) {
            return customBranchPolicies(Output.of(customBranchPolicies));
        }

        /**
         * @param protectedBranches Whether only branches with branch protection rules can deploy to this environment.
         * 
         * @return builder
         * 
         */
        public Builder protectedBranches(Output<Boolean> protectedBranches) {
            $.protectedBranches = protectedBranches;
            return this;
        }

        /**
         * @param protectedBranches Whether only branches with branch protection rules can deploy to this environment.
         * 
         * @return builder
         * 
         */
        public Builder protectedBranches(Boolean protectedBranches) {
            return protectedBranches(Output.of(protectedBranches));
        }

        public RepositoryEnvironmentDeploymentBranchPolicyArgs build() {
            $.customBranchPolicies = Objects.requireNonNull($.customBranchPolicies, "expected parameter 'customBranchPolicies' to be non-null");
            $.protectedBranches = Objects.requireNonNull($.protectedBranches, "expected parameter 'protectedBranches' to be non-null");
            return $;
        }
    }

}
