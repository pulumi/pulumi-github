// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.inputs.ProviderAppAuthArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
     * and `app_auth` are not set.
     * 
     */
    @Import(name="appAuth", json=true)
    private @Nullable Output<ProviderAppAuthArgs> appAuth;

    /**
     * @return The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
     * and `app_auth` are not set.
     * 
     */
    public Optional<Output<ProviderAppAuthArgs>> appAuth() {
        return Optional.ofNullable(this.appAuth);
    }

    /**
     * The GitHub Base API URL
     * 
     */
    @Import(name="baseUrl")
    private @Nullable Output<String> baseUrl;

    /**
     * @return The GitHub Base API URL
     * 
     */
    public Optional<Output<String>> baseUrl() {
        return Optional.ofNullable(this.baseUrl);
    }

    /**
     * Enable `insecure` mode for testing purposes
     * 
     */
    @Import(name="insecure", json=true)
    private @Nullable Output<Boolean> insecure;

    /**
     * @return Enable `insecure` mode for testing purposes
     * 
     */
    public Optional<Output<Boolean>> insecure() {
        return Optional.ofNullable(this.insecure);
    }

    /**
     * Number of times to retry a request after receiving an error status codeDefaults to 3
     * 
     */
    @Import(name="maxRetries", json=true)
    private @Nullable Output<Integer> maxRetries;

    /**
     * @return Number of times to retry a request after receiving an error status codeDefaults to 3
     * 
     */
    public Optional<Output<Integer>> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }

    /**
     * The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
     * 
     * @deprecated
     * Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)
     * 
     */
    @Deprecated /* Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION) */
    @Import(name="organization")
    private @Nullable Output<String> organization;

    /**
     * @return The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
     * 
     * @deprecated
     * Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)
     * 
     */
    @Deprecated /* Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION) */
    public Optional<Output<String>> organization() {
        return Optional.ofNullable(this.organization);
    }

    /**
     * The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * Allow the provider to make parallel API calls to GitHub. You may want to set it to true when you have a private Github
     * Enterprise without strict rate limits. Although, it is not possible to enable this setting on github.com because we
     * enforce the respect of github.com&#39;s best practices to avoid hitting abuse rate limitsDefaults to false if not set
     * 
     */
    @Import(name="parallelRequests", json=true)
    private @Nullable Output<Boolean> parallelRequests;

    /**
     * @return Allow the provider to make parallel API calls to GitHub. You may want to set it to true when you have a private Github
     * Enterprise without strict rate limits. Although, it is not possible to enable this setting on github.com because we
     * enforce the respect of github.com&#39;s best practices to avoid hitting abuse rate limitsDefaults to false if not set
     * 
     */
    public Optional<Output<Boolean>> parallelRequests() {
        return Optional.ofNullable(this.parallelRequests);
    }

    /**
     * Amount of time in milliseconds to sleep in between non-write requests to GitHub API. Defaults to 0ms if not set.
     * 
     */
    @Import(name="readDelayMs", json=true)
    private @Nullable Output<Integer> readDelayMs;

    /**
     * @return Amount of time in milliseconds to sleep in between non-write requests to GitHub API. Defaults to 0ms if not set.
     * 
     */
    public Optional<Output<Integer>> readDelayMs() {
        return Optional.ofNullable(this.readDelayMs);
    }

    /**
     * Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or
     * 1s if not set, the max_retries must be set to greater than zero.
     * 
     */
    @Import(name="retryDelayMs", json=true)
    private @Nullable Output<Integer> retryDelayMs;

    /**
     * @return Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or
     * 1s if not set, the max_retries must be set to greater than zero.
     * 
     */
    public Optional<Output<Integer>> retryDelayMs() {
        return Optional.ofNullable(this.retryDelayMs);
    }

    /**
     * Allow the provider to retry after receiving an error status code, the max_retries should be set for this to workDefaults
     * to [500, 502, 503, 504]
     * 
     */
    @Import(name="retryableErrors", json=true)
    private @Nullable Output<List<Integer>> retryableErrors;

    /**
     * @return Allow the provider to retry after receiving an error status code, the max_retries should be set for this to workDefaults
     * to [500, 502, 503, 504]
     * 
     */
    public Optional<Output<List<Integer>>> retryableErrors() {
        return Optional.ofNullable(this.retryableErrors);
    }

    /**
     * The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.
     * 
     */
    @Import(name="writeDelayMs", json=true)
    private @Nullable Output<Integer> writeDelayMs;

    /**
     * @return Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.
     * 
     */
    public Optional<Output<Integer>> writeDelayMs() {
        return Optional.ofNullable(this.writeDelayMs);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.appAuth = $.appAuth;
        this.baseUrl = $.baseUrl;
        this.insecure = $.insecure;
        this.maxRetries = $.maxRetries;
        this.organization = $.organization;
        this.owner = $.owner;
        this.parallelRequests = $.parallelRequests;
        this.readDelayMs = $.readDelayMs;
        this.retryDelayMs = $.retryDelayMs;
        this.retryableErrors = $.retryableErrors;
        this.token = $.token;
        this.writeDelayMs = $.writeDelayMs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param appAuth The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
         * and `app_auth` are not set.
         * 
         * @return builder
         * 
         */
        public Builder appAuth(@Nullable Output<ProviderAppAuthArgs> appAuth) {
            $.appAuth = appAuth;
            return this;
        }

        /**
         * @param appAuth The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
         * and `app_auth` are not set.
         * 
         * @return builder
         * 
         */
        public Builder appAuth(ProviderAppAuthArgs appAuth) {
            return appAuth(Output.of(appAuth));
        }

        /**
         * @param baseUrl The GitHub Base API URL
         * 
         * @return builder
         * 
         */
        public Builder baseUrl(@Nullable Output<String> baseUrl) {
            $.baseUrl = baseUrl;
            return this;
        }

        /**
         * @param baseUrl The GitHub Base API URL
         * 
         * @return builder
         * 
         */
        public Builder baseUrl(String baseUrl) {
            return baseUrl(Output.of(baseUrl));
        }

        /**
         * @param insecure Enable `insecure` mode for testing purposes
         * 
         * @return builder
         * 
         */
        public Builder insecure(@Nullable Output<Boolean> insecure) {
            $.insecure = insecure;
            return this;
        }

        /**
         * @param insecure Enable `insecure` mode for testing purposes
         * 
         * @return builder
         * 
         */
        public Builder insecure(Boolean insecure) {
            return insecure(Output.of(insecure));
        }

        /**
         * @param maxRetries Number of times to retry a request after receiving an error status codeDefaults to 3
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(@Nullable Output<Integer> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        /**
         * @param maxRetries Number of times to retry a request after receiving an error status codeDefaults to 3
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Integer maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        /**
         * @param organization The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
         * 
         * @return builder
         * 
         * @deprecated
         * Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)
         * 
         */
        @Deprecated /* Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION) */
        public Builder organization(@Nullable Output<String> organization) {
            $.organization = organization;
            return this;
        }

        /**
         * @param organization The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
         * 
         * @return builder
         * 
         * @deprecated
         * Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)
         * 
         */
        @Deprecated /* Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION) */
        public Builder organization(String organization) {
            return organization(Output.of(organization));
        }

        /**
         * @param owner The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param parallelRequests Allow the provider to make parallel API calls to GitHub. You may want to set it to true when you have a private Github
         * Enterprise without strict rate limits. Although, it is not possible to enable this setting on github.com because we
         * enforce the respect of github.com&#39;s best practices to avoid hitting abuse rate limitsDefaults to false if not set
         * 
         * @return builder
         * 
         */
        public Builder parallelRequests(@Nullable Output<Boolean> parallelRequests) {
            $.parallelRequests = parallelRequests;
            return this;
        }

        /**
         * @param parallelRequests Allow the provider to make parallel API calls to GitHub. You may want to set it to true when you have a private Github
         * Enterprise without strict rate limits. Although, it is not possible to enable this setting on github.com because we
         * enforce the respect of github.com&#39;s best practices to avoid hitting abuse rate limitsDefaults to false if not set
         * 
         * @return builder
         * 
         */
        public Builder parallelRequests(Boolean parallelRequests) {
            return parallelRequests(Output.of(parallelRequests));
        }

        /**
         * @param readDelayMs Amount of time in milliseconds to sleep in between non-write requests to GitHub API. Defaults to 0ms if not set.
         * 
         * @return builder
         * 
         */
        public Builder readDelayMs(@Nullable Output<Integer> readDelayMs) {
            $.readDelayMs = readDelayMs;
            return this;
        }

        /**
         * @param readDelayMs Amount of time in milliseconds to sleep in between non-write requests to GitHub API. Defaults to 0ms if not set.
         * 
         * @return builder
         * 
         */
        public Builder readDelayMs(Integer readDelayMs) {
            return readDelayMs(Output.of(readDelayMs));
        }

        /**
         * @param retryDelayMs Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or
         * 1s if not set, the max_retries must be set to greater than zero.
         * 
         * @return builder
         * 
         */
        public Builder retryDelayMs(@Nullable Output<Integer> retryDelayMs) {
            $.retryDelayMs = retryDelayMs;
            return this;
        }

        /**
         * @param retryDelayMs Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or
         * 1s if not set, the max_retries must be set to greater than zero.
         * 
         * @return builder
         * 
         */
        public Builder retryDelayMs(Integer retryDelayMs) {
            return retryDelayMs(Output.of(retryDelayMs));
        }

        /**
         * @param retryableErrors Allow the provider to retry after receiving an error status code, the max_retries should be set for this to workDefaults
         * to [500, 502, 503, 504]
         * 
         * @return builder
         * 
         */
        public Builder retryableErrors(@Nullable Output<List<Integer>> retryableErrors) {
            $.retryableErrors = retryableErrors;
            return this;
        }

        /**
         * @param retryableErrors Allow the provider to retry after receiving an error status code, the max_retries should be set for this to workDefaults
         * to [500, 502, 503, 504]
         * 
         * @return builder
         * 
         */
        public Builder retryableErrors(List<Integer> retryableErrors) {
            return retryableErrors(Output.of(retryableErrors));
        }

        /**
         * @param retryableErrors Allow the provider to retry after receiving an error status code, the max_retries should be set for this to workDefaults
         * to [500, 502, 503, 504]
         * 
         * @return builder
         * 
         */
        public Builder retryableErrors(Integer... retryableErrors) {
            return retryableErrors(List.of(retryableErrors));
        }

        /**
         * @param token The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param writeDelayMs Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.
         * 
         * @return builder
         * 
         */
        public Builder writeDelayMs(@Nullable Output<Integer> writeDelayMs) {
            $.writeDelayMs = writeDelayMs;
            return this;
        }

        /**
         * @param writeDelayMs Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.
         * 
         * @return builder
         * 
         */
        public Builder writeDelayMs(Integer writeDelayMs) {
            return writeDelayMs(Output.of(writeDelayMs));
        }

        public ProviderArgs build() {
            $.baseUrl = Codegen.stringProp("baseUrl").output().arg($.baseUrl).env("GITHUB_BASE_URL").def("https://api.github.com/").getNullable();
            return $;
        }
    }

}
