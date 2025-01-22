// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoriesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoriesPlainArgs Empty = new GetRepositoriesPlainArgs();

    /**
     * Returns a list of found repository IDs
     * 
     */
    @Import(name="includeRepoId")
    private @Nullable Boolean includeRepoId;

    /**
     * @return Returns a list of found repository IDs
     * 
     */
    public Optional<Boolean> includeRepoId() {
        return Optional.ofNullable(this.includeRepoId);
    }

    /**
     * Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
     * 
     */
    @Import(name="query", required=true)
    private String query;

    /**
     * @return Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
     * 
     */
    public String query() {
        return this.query;
    }

    /**
     * Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
     * 
     */
    @Import(name="resultsPerPage")
    private @Nullable Integer resultsPerPage;

    /**
     * @return Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
     * 
     */
    public Optional<Integer> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }

    /**
     * Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
     * 
     */
    @Import(name="sort")
    private @Nullable String sort;

    /**
     * @return Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
     * 
     */
    public Optional<String> sort() {
        return Optional.ofNullable(this.sort);
    }

    private GetRepositoriesPlainArgs() {}

    private GetRepositoriesPlainArgs(GetRepositoriesPlainArgs $) {
        this.includeRepoId = $.includeRepoId;
        this.query = $.query;
        this.resultsPerPage = $.resultsPerPage;
        this.sort = $.sort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoriesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoriesPlainArgs $;

        public Builder() {
            $ = new GetRepositoriesPlainArgs();
        }

        public Builder(GetRepositoriesPlainArgs defaults) {
            $ = new GetRepositoriesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param includeRepoId Returns a list of found repository IDs
         * 
         * @return builder
         * 
         */
        public Builder includeRepoId(@Nullable Boolean includeRepoId) {
            $.includeRepoId = includeRepoId;
            return this;
        }

        /**
         * @param query Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
         * 
         * @return builder
         * 
         */
        public Builder query(String query) {
            $.query = query;
            return this;
        }

        /**
         * @param resultsPerPage Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
         * 
         * @return builder
         * 
         */
        public Builder resultsPerPage(@Nullable Integer resultsPerPage) {
            $.resultsPerPage = resultsPerPage;
            return this;
        }

        /**
         * @param sort Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
         * 
         * @return builder
         * 
         */
        public Builder sort(@Nullable String sort) {
            $.sort = sort;
            return this;
        }

        public GetRepositoriesPlainArgs build() {
            if ($.query == null) {
                throw new MissingRequiredPropertyException("GetRepositoriesPlainArgs", "query");
            }
            return $;
        }
    }

}