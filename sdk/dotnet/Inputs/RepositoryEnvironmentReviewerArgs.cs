// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryEnvironmentReviewerArgs : global::Pulumi.ResourceArgs
    {
        [Input("teams")]
        private InputList<int>? _teams;
        public InputList<int> Teams
        {
            get => _teams ?? (_teams = new InputList<int>());
            set => _teams = value;
        }

        [Input("users")]
        private InputList<int>? _users;
        public InputList<int> Users
        {
            get => _users ?? (_users = new InputList<int>());
            set => _users = value;
        }

        public RepositoryEnvironmentReviewerArgs()
        {
        }
        public static new RepositoryEnvironmentReviewerArgs Empty => new RepositoryEnvironmentReviewerArgs();
    }
}
