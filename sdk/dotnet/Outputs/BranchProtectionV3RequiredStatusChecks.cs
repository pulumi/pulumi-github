// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class BranchProtectionV3RequiredStatusChecks
    {
        /// <summary>
        /// The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the context and app_id like so "context:app_id".
        /// </summary>
        public readonly ImmutableArray<string> Checks;
        /// <summary>
        /// [**DEPRECATED**] (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.
        /// 
        /// &gt; Note: This attribute can contain multiple string patterns.
        /// If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
        /// For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
        /// For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
        /// </summary>
        public readonly ImmutableArray<string> Contexts;
        public readonly bool? IncludeAdmins;
        /// <summary>
        /// Require branches to be up to date before merging. Defaults to `false`.
        /// </summary>
        public readonly bool? Strict;

        [OutputConstructor]
        private BranchProtectionV3RequiredStatusChecks(
            ImmutableArray<string> checks,

            ImmutableArray<string> contexts,

            bool? includeAdmins,

            bool? strict)
        {
            Checks = checks;
            Contexts = contexts;
            IncludeAdmins = includeAdmins;
            Strict = strict;
        }
    }
}