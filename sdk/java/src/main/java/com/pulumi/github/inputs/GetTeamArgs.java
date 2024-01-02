// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTeamArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTeamArgs Empty = new GetTeamArgs();

    /**
     * Type of membershp to be requested to fill the list of members. Can be either &#34;all&#34; or &#34;immediate&#34;. Default: &#34;all&#34;
     * 
     */
    @Import(name="membershipType")
    private @Nullable Output<String> membershipType;

    /**
     * @return Type of membershp to be requested to fill the list of members. Can be either &#34;all&#34; or &#34;immediate&#34;. Default: &#34;all&#34;
     * 
     */
    public Optional<Output<String>> membershipType() {
        return Optional.ofNullable(this.membershipType);
    }

    /**
     * Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
     * 
     */
    @Import(name="resultsPerPage")
    private @Nullable Output<Integer> resultsPerPage;

    /**
     * @return Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
     * 
     */
    public Optional<Output<Integer>> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }

    /**
     * The team slug.
     * 
     */
    @Import(name="slug", required=true)
    private Output<String> slug;

    /**
     * @return The team slug.
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }

    /**
     * Exclude the members and repositories of the team from the returned result. Defaults to `false`.
     * 
     */
    @Import(name="summaryOnly")
    private @Nullable Output<Boolean> summaryOnly;

    /**
     * @return Exclude the members and repositories of the team from the returned result. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> summaryOnly() {
        return Optional.ofNullable(this.summaryOnly);
    }

    private GetTeamArgs() {}

    private GetTeamArgs(GetTeamArgs $) {
        this.membershipType = $.membershipType;
        this.resultsPerPage = $.resultsPerPage;
        this.slug = $.slug;
        this.summaryOnly = $.summaryOnly;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTeamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTeamArgs $;

        public Builder() {
            $ = new GetTeamArgs();
        }

        public Builder(GetTeamArgs defaults) {
            $ = new GetTeamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param membershipType Type of membershp to be requested to fill the list of members. Can be either &#34;all&#34; or &#34;immediate&#34;. Default: &#34;all&#34;
         * 
         * @return builder
         * 
         */
        public Builder membershipType(@Nullable Output<String> membershipType) {
            $.membershipType = membershipType;
            return this;
        }

        /**
         * @param membershipType Type of membershp to be requested to fill the list of members. Can be either &#34;all&#34; or &#34;immediate&#34;. Default: &#34;all&#34;
         * 
         * @return builder
         * 
         */
        public Builder membershipType(String membershipType) {
            return membershipType(Output.of(membershipType));
        }

        /**
         * @param resultsPerPage Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
         * 
         * @return builder
         * 
         */
        public Builder resultsPerPage(@Nullable Output<Integer> resultsPerPage) {
            $.resultsPerPage = resultsPerPage;
            return this;
        }

        /**
         * @param resultsPerPage Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
         * 
         * @return builder
         * 
         */
        public Builder resultsPerPage(Integer resultsPerPage) {
            return resultsPerPage(Output.of(resultsPerPage));
        }

        /**
         * @param slug The team slug.
         * 
         * @return builder
         * 
         */
        public Builder slug(Output<String> slug) {
            $.slug = slug;
            return this;
        }

        /**
         * @param slug The team slug.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            return slug(Output.of(slug));
        }

        /**
         * @param summaryOnly Exclude the members and repositories of the team from the returned result. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder summaryOnly(@Nullable Output<Boolean> summaryOnly) {
            $.summaryOnly = summaryOnly;
            return this;
        }

        /**
         * @param summaryOnly Exclude the members and repositories of the team from the returned result. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder summaryOnly(Boolean summaryOnly) {
            return summaryOnly(Output.of(summaryOnly));
        }

        public GetTeamArgs build() {
            if ($.slug == null) {
                throw new MissingRequiredPropertyException("GetTeamArgs", "slug");
            }
            return $;
        }
    }

}
