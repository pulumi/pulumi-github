// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchProtectionRequiredStatusCheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchProtectionRequiredStatusCheckArgs Empty = new BranchProtectionRequiredStatusCheckArgs();

    /**
     * The list of status checks to require in order to merge into this branch. No status checks are required by default.
     * 
     * &gt; Note: This attribute can contain multiple string patterns.
     * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
     * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
     * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
     * 
     */
    @Import(name="contexts")
    private @Nullable Output<List<String>> contexts;

    /**
     * @return The list of status checks to require in order to merge into this branch. No status checks are required by default.
     * 
     * &gt; Note: This attribute can contain multiple string patterns.
     * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
     * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
     * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
     * 
     */
    public Optional<Output<List<String>>> contexts() {
        return Optional.ofNullable(this.contexts);
    }

    /**
     * Require branches to be up to date before merging. Defaults to `false`.
     * 
     */
    @Import(name="strict")
    private @Nullable Output<Boolean> strict;

    /**
     * @return Require branches to be up to date before merging. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> strict() {
        return Optional.ofNullable(this.strict);
    }

    private BranchProtectionRequiredStatusCheckArgs() {}

    private BranchProtectionRequiredStatusCheckArgs(BranchProtectionRequiredStatusCheckArgs $) {
        this.contexts = $.contexts;
        this.strict = $.strict;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchProtectionRequiredStatusCheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchProtectionRequiredStatusCheckArgs $;

        public Builder() {
            $ = new BranchProtectionRequiredStatusCheckArgs();
        }

        public Builder(BranchProtectionRequiredStatusCheckArgs defaults) {
            $ = new BranchProtectionRequiredStatusCheckArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param contexts The list of status checks to require in order to merge into this branch. No status checks are required by default.
         * 
         * &gt; Note: This attribute can contain multiple string patterns.
         * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
         * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
         * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
         * 
         * @return builder
         * 
         */
        public Builder contexts(@Nullable Output<List<String>> contexts) {
            $.contexts = contexts;
            return this;
        }

        /**
         * @param contexts The list of status checks to require in order to merge into this branch. No status checks are required by default.
         * 
         * &gt; Note: This attribute can contain multiple string patterns.
         * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
         * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
         * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
         * 
         * @return builder
         * 
         */
        public Builder contexts(List<String> contexts) {
            return contexts(Output.of(contexts));
        }

        /**
         * @param contexts The list of status checks to require in order to merge into this branch. No status checks are required by default.
         * 
         * &gt; Note: This attribute can contain multiple string patterns.
         * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
         * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
         * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
         * 
         * @return builder
         * 
         */
        public Builder contexts(String... contexts) {
            return contexts(List.of(contexts));
        }

        /**
         * @param strict Require branches to be up to date before merging. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder strict(@Nullable Output<Boolean> strict) {
            $.strict = strict;
            return this;
        }

        /**
         * @param strict Require branches to be up to date before merging. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder strict(Boolean strict) {
            return strict(Output.of(strict));
        }

        public BranchProtectionRequiredStatusCheckArgs build() {
            return $;
        }
    }

}