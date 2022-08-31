// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetIpRanges
    {
        /// <summary>
        /// Use this data source to retrieve information about GitHub's IP addresses.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Github.GetIpRanges.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIpRangesResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIpRangesResult>("github:index/getIpRanges:getIpRanges", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIpRangesResult
    {
        /// <summary>
        /// An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// A subset of the `actions` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> ActionsIpv4s;
        /// <summary>
        /// A subset of the `actions` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> ActionsIpv6s;
        /// <summary>
        /// A subset of the `api` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> ApiIpv4s;
        /// <summary>
        /// A subset of the `api` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> ApiIpv6s;
        /// <summary>
        /// An Array of IP addresses in CIDR format for the GitHub API.
        /// </summary>
        public readonly ImmutableArray<string> Apis;
        /// <summary>
        /// A subset of the `dependabot` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> DependabotIpv4s;
        /// <summary>
        /// A subset of the `dependabot` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> DependabotIpv6s;
        /// <summary>
        /// An array of IP addresses in CIDR format specifying the A records for dependabot.
        /// </summary>
        public readonly ImmutableArray<string> Dependabots;
        /// <summary>
        /// A subset of the `git` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> GitIpv4s;
        /// <summary>
        /// A subset of the `git` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> GitIpv6s;
        /// <summary>
        /// An Array of IP addresses in CIDR format specifying the Git servers.
        /// </summary>
        public readonly ImmutableArray<string> Gits;
        /// <summary>
        /// An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
        /// </summary>
        public readonly ImmutableArray<string> Hooks;
        /// <summary>
        /// A subset of the `hooks` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> HooksIpv4s;
        /// <summary>
        /// A subset of the `hooks` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> HooksIpv6s;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A subset of the `importer` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> ImporterIpv4s;
        /// <summary>
        /// A subset of the `importer` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> ImporterIpv6s;
        /// <summary>
        /// An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
        /// </summary>
        public readonly ImmutableArray<string> Importers;
        /// <summary>
        /// An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
        /// </summary>
        public readonly ImmutableArray<string> Pages;
        /// <summary>
        /// A subset of the `pages` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> PagesIpv4s;
        /// <summary>
        /// A subset of the `pages` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> PagesIpv6s;
        /// <summary>
        /// A subset of the `web` array that contains IP addresses in IPv4 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> WebIpv4s;
        /// <summary>
        /// A subset of the `web` array that contains IP addresses in IPv6 CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> WebIpv6s;
        /// <summary>
        /// An Array of IP addresses in CIDR format for GitHub Web.
        /// </summary>
        public readonly ImmutableArray<string> Webs;

        [OutputConstructor]
        private GetIpRangesResult(
            ImmutableArray<string> actions,

            ImmutableArray<string> actionsIpv4s,

            ImmutableArray<string> actionsIpv6s,

            ImmutableArray<string> apiIpv4s,

            ImmutableArray<string> apiIpv6s,

            ImmutableArray<string> apis,

            ImmutableArray<string> dependabotIpv4s,

            ImmutableArray<string> dependabotIpv6s,

            ImmutableArray<string> dependabots,

            ImmutableArray<string> gitIpv4s,

            ImmutableArray<string> gitIpv6s,

            ImmutableArray<string> gits,

            ImmutableArray<string> hooks,

            ImmutableArray<string> hooksIpv4s,

            ImmutableArray<string> hooksIpv6s,

            string id,

            ImmutableArray<string> importerIpv4s,

            ImmutableArray<string> importerIpv6s,

            ImmutableArray<string> importers,

            ImmutableArray<string> pages,

            ImmutableArray<string> pagesIpv4s,

            ImmutableArray<string> pagesIpv6s,

            ImmutableArray<string> webIpv4s,

            ImmutableArray<string> webIpv6s,

            ImmutableArray<string> webs)
        {
            Actions = actions;
            ActionsIpv4s = actionsIpv4s;
            ActionsIpv6s = actionsIpv6s;
            ApiIpv4s = apiIpv4s;
            ApiIpv6s = apiIpv6s;
            Apis = apis;
            DependabotIpv4s = dependabotIpv4s;
            DependabotIpv6s = dependabotIpv6s;
            Dependabots = dependabots;
            GitIpv4s = gitIpv4s;
            GitIpv6s = gitIpv6s;
            Gits = gits;
            Hooks = hooks;
            HooksIpv4s = hooksIpv4s;
            HooksIpv6s = hooksIpv6s;
            Id = id;
            ImporterIpv4s = importerIpv4s;
            ImporterIpv6s = importerIpv6s;
            Importers = importers;
            Pages = pages;
            PagesIpv4s = pagesIpv4s;
            PagesIpv6s = pagesIpv6s;
            WebIpv4s = webIpv4s;
            WebIpv6s = webIpv6s;
            Webs = webs;
        }
    }
}
