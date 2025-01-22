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


public final class BranchProtectionV3RequiredStatusChecksArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchProtectionV3RequiredStatusChecksArgs Empty = new BranchProtectionV3RequiredStatusChecksArgs();

    /**
     * The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the context and app_id like so &#34;context:app_id&#34;.
     * 
     */
    @Import(name="checks")
    private @Nullable Output<List<String>> checks;

    /**
     * @return The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the context and app_id like so &#34;context:app_id&#34;.
     * 
     */
    public Optional<Output<List<String>>> checks() {
        return Optional.ofNullable(this.checks);
    }

    /**
     * [**DEPRECATED**] (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.
     * 
     * &gt; Note: This attribute can contain multiple string patterns.
     * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
     * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
     * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
     * 
     * @deprecated
     * GitHub is deprecating the use of `contexts`. Use a `checks` array instead.
     * 
     */
    @Deprecated /* GitHub is deprecating the use of `contexts`. Use a `checks` array instead. */
    @Import(name="contexts")
    private @Nullable Output<List<String>> contexts;

    /**
     * @return [**DEPRECATED**] (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.
     * 
     * &gt; Note: This attribute can contain multiple string patterns.
     * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
     * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
     * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
     * 
     * @deprecated
     * GitHub is deprecating the use of `contexts`. Use a `checks` array instead.
     * 
     */
    @Deprecated /* GitHub is deprecating the use of `contexts`. Use a `checks` array instead. */
    public Optional<Output<List<String>>> contexts() {
        return Optional.ofNullable(this.contexts);
    }

    /**
     * @deprecated
     * Use enforce_admins instead
     * 
     */
    @Deprecated /* Use enforce_admins instead */
    @Import(name="includeAdmins")
    private @Nullable Output<Boolean> includeAdmins;

    /**
     * @deprecated
     * Use enforce_admins instead
     * 
     */
    @Deprecated /* Use enforce_admins instead */
    public Optional<Output<Boolean>> includeAdmins() {
        return Optional.ofNullable(this.includeAdmins);
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

    private BranchProtectionV3RequiredStatusChecksArgs() {}

    private BranchProtectionV3RequiredStatusChecksArgs(BranchProtectionV3RequiredStatusChecksArgs $) {
        this.checks = $.checks;
        this.contexts = $.contexts;
        this.includeAdmins = $.includeAdmins;
        this.strict = $.strict;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchProtectionV3RequiredStatusChecksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchProtectionV3RequiredStatusChecksArgs $;

        public Builder() {
            $ = new BranchProtectionV3RequiredStatusChecksArgs();
        }

        public Builder(BranchProtectionV3RequiredStatusChecksArgs defaults) {
            $ = new BranchProtectionV3RequiredStatusChecksArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param checks The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the context and app_id like so &#34;context:app_id&#34;.
         * 
         * @return builder
         * 
         */
        public Builder checks(@Nullable Output<List<String>> checks) {
            $.checks = checks;
            return this;
        }

        /**
         * @param checks The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the context and app_id like so &#34;context:app_id&#34;.
         * 
         * @return builder
         * 
         */
        public Builder checks(List<String> checks) {
            return checks(Output.of(checks));
        }

        /**
         * @param checks The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the context and app_id like so &#34;context:app_id&#34;.
         * 
         * @return builder
         * 
         */
        public Builder checks(String... checks) {
            return checks(List.of(checks));
        }

        /**
         * @param contexts [**DEPRECATED**] (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.
         * 
         * &gt; Note: This attribute can contain multiple string patterns.
         * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
         * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
         * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
         * 
         * @return builder
         * 
         * @deprecated
         * GitHub is deprecating the use of `contexts`. Use a `checks` array instead.
         * 
         */
        @Deprecated /* GitHub is deprecating the use of `contexts`. Use a `checks` array instead. */
        public Builder contexts(@Nullable Output<List<String>> contexts) {
            $.contexts = contexts;
            return this;
        }

        /**
         * @param contexts [**DEPRECATED**] (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.
         * 
         * &gt; Note: This attribute can contain multiple string patterns.
         * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
         * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
         * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
         * 
         * @return builder
         * 
         * @deprecated
         * GitHub is deprecating the use of `contexts`. Use a `checks` array instead.
         * 
         */
        @Deprecated /* GitHub is deprecating the use of `contexts`. Use a `checks` array instead. */
        public Builder contexts(List<String> contexts) {
            return contexts(Output.of(contexts));
        }

        /**
         * @param contexts [**DEPRECATED**] (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.
         * 
         * &gt; Note: This attribute can contain multiple string patterns.
         * If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
         * For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
         * For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
         * 
         * @return builder
         * 
         * @deprecated
         * GitHub is deprecating the use of `contexts`. Use a `checks` array instead.
         * 
         */
        @Deprecated /* GitHub is deprecating the use of `contexts`. Use a `checks` array instead. */
        public Builder contexts(String... contexts) {
            return contexts(List.of(contexts));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Use enforce_admins instead
         * 
         */
        @Deprecated /* Use enforce_admins instead */
        public Builder includeAdmins(@Nullable Output<Boolean> includeAdmins) {
            $.includeAdmins = includeAdmins;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Use enforce_admins instead
         * 
         */
        @Deprecated /* Use enforce_admins instead */
        public Builder includeAdmins(Boolean includeAdmins) {
            return includeAdmins(Output.of(includeAdmins));
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

        public BranchProtectionV3RequiredStatusChecksArgs build() {
            return $;
        }
    }

}