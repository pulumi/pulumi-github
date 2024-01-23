// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class BranchProtectionRequiredStatusCheckGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("contexts")]
        private InputList<string>? _contexts;

        /// <summary>
        /// The list of status checks to require in order to merge into this branch. No status checks are required by default.
        /// 
        /// &gt; Note: This attribute can contain multiple string patterns.
        /// If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
        /// For workflows that use matrixes, append the matrix name to the value using the following pattern `(&lt;matrix_value&gt;[, &lt;matrix_value&gt;])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See GitHub Documentation for more information.
        /// For workflows that use reusable workflows, the pattern is `&lt;initial_workflow.jobs.job.[name/id]&gt; / &lt;reused-workflow.jobs.job.[name/id]&gt;`. This can extend multiple levels.
        /// </summary>
        public InputList<string> Contexts
        {
            get => _contexts ?? (_contexts = new InputList<string>());
            set => _contexts = value;
        }

        /// <summary>
        /// Require branches to be up to date before merging. Defaults to `false`.
        /// </summary>
        [Input("strict")]
        public Input<bool>? Strict { get; set; }

        public BranchProtectionRequiredStatusCheckGetArgs()
        {
        }
        public static new BranchProtectionRequiredStatusCheckGetArgs Empty => new BranchProtectionRequiredStatusCheckGetArgs();
    }
}
