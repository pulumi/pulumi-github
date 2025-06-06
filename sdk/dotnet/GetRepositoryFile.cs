// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryFile
    {
        /// <summary>
        /// This data source allows you to read files within a
        /// GitHub repository.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Github.GetRepositoryFile.Invoke(new()
        ///     {
        ///         Repository = fooGithubRepository.Name,
        ///         Branch = "main",
        ///         File = ".gitignore",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryFileResult> InvokeAsync(GetRepositoryFileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryFileResult>("github:index/getRepositoryFile:getRepositoryFile", args ?? new GetRepositoryFileArgs(), options.WithDefaults());

        /// <summary>
        /// This data source allows you to read files within a
        /// GitHub repository.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Github.GetRepositoryFile.Invoke(new()
        ///     {
        ///         Repository = fooGithubRepository.Name,
        ///         Branch = "main",
        ///         File = ".gitignore",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryFileResult> Invoke(GetRepositoryFileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryFileResult>("github:index/getRepositoryFile:getRepositoryFile", args ?? new GetRepositoryFileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source allows you to read files within a
        /// GitHub repository.
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Github.GetRepositoryFile.Invoke(new()
        ///     {
        ///         Repository = fooGithubRepository.Name,
        ///         Branch = "main",
        ///         File = ".gitignore",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryFileResult> Invoke(GetRepositoryFileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryFileResult>("github:index/getRepositoryFile:getRepositoryFile", args ?? new GetRepositoryFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryFileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Git branch. Defaults to the repository's default branch.
        /// </summary>
        [Input("branch")]
        public string? Branch { get; set; }

        /// <summary>
        /// The path of the file to read.
        /// </summary>
        [Input("file", required: true)]
        public string File { get; set; } = null!;

        /// <summary>
        /// The repository to read the file from. If an unqualified repo name (without an owner) is passed, the owner will be inferred from the owner of the token used to execute the plan. If a name of the type "owner/repo" (with a slash in the middle) is passed, the owner will be as specified and not the owner of the token.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryFileArgs()
        {
        }
        public static new GetRepositoryFileArgs Empty => new GetRepositoryFileArgs();
    }

    public sealed class GetRepositoryFileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Git branch. Defaults to the repository's default branch.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// The path of the file to read.
        /// </summary>
        [Input("file", required: true)]
        public Input<string> File { get; set; } = null!;

        /// <summary>
        /// The repository to read the file from. If an unqualified repo name (without an owner) is passed, the owner will be inferred from the owner of the token used to execute the plan. If a name of the type "owner/repo" (with a slash in the middle) is passed, the owner will be as specified and not the owner of the token.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryFileInvokeArgs()
        {
        }
        public static new GetRepositoryFileInvokeArgs Empty => new GetRepositoryFileInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryFileResult
    {
        public readonly string? Branch;
        /// <summary>
        /// Committer author name.
        /// </summary>
        public readonly string CommitAuthor;
        /// <summary>
        /// Committer email address.
        /// </summary>
        public readonly string CommitEmail;
        /// <summary>
        /// Commit message when file was last updated.
        /// </summary>
        public readonly string CommitMessage;
        /// <summary>
        /// The SHA of the commit that modified the file.
        /// </summary>
        public readonly string CommitSha;
        /// <summary>
        /// The file content.
        /// </summary>
        public readonly string Content;
        public readonly string File;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the commit/branch/tag.
        /// </summary>
        public readonly string Ref;
        public readonly string Repository;
        /// <summary>
        /// The SHA blob of the file.
        /// </summary>
        public readonly string Sha;

        [OutputConstructor]
        private GetRepositoryFileResult(
            string? branch,

            string commitAuthor,

            string commitEmail,

            string commitMessage,

            string commitSha,

            string content,

            string file,

            string id,

            string @ref,

            string repository,

            string sha)
        {
            Branch = branch;
            CommitAuthor = commitAuthor;
            CommitEmail = commitEmail;
            CommitMessage = commitMessage;
            CommitSha = commitSha;
            Content = content;
            File = file;
            Id = id;
            Ref = @ref;
            Repository = repository;
            Sha = sha;
        }
    }
}
