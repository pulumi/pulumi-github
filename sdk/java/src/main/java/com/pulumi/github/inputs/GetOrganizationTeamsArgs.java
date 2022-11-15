// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOrganizationTeamsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOrganizationTeamsArgs Empty = new GetOrganizationTeamsArgs();

    @Import(name="resultsPerPage")
    private @Nullable Output<Integer> resultsPerPage;

    public Optional<Output<Integer>> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }

    @Import(name="rootTeamsOnly")
    private @Nullable Output<Boolean> rootTeamsOnly;

    public Optional<Output<Boolean>> rootTeamsOnly() {
        return Optional.ofNullable(this.rootTeamsOnly);
    }

    @Import(name="summaryOnly")
    private @Nullable Output<Boolean> summaryOnly;

    public Optional<Output<Boolean>> summaryOnly() {
        return Optional.ofNullable(this.summaryOnly);
    }

    private GetOrganizationTeamsArgs() {}

    private GetOrganizationTeamsArgs(GetOrganizationTeamsArgs $) {
        this.resultsPerPage = $.resultsPerPage;
        this.rootTeamsOnly = $.rootTeamsOnly;
        this.summaryOnly = $.summaryOnly;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOrganizationTeamsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOrganizationTeamsArgs $;

        public Builder() {
            $ = new GetOrganizationTeamsArgs();
        }

        public Builder(GetOrganizationTeamsArgs defaults) {
            $ = new GetOrganizationTeamsArgs(Objects.requireNonNull(defaults));
        }

        public Builder resultsPerPage(@Nullable Output<Integer> resultsPerPage) {
            $.resultsPerPage = resultsPerPage;
            return this;
        }

        public Builder resultsPerPage(Integer resultsPerPage) {
            return resultsPerPage(Output.of(resultsPerPage));
        }

        public Builder rootTeamsOnly(@Nullable Output<Boolean> rootTeamsOnly) {
            $.rootTeamsOnly = rootTeamsOnly;
            return this;
        }

        public Builder rootTeamsOnly(Boolean rootTeamsOnly) {
            return rootTeamsOnly(Output.of(rootTeamsOnly));
        }

        public Builder summaryOnly(@Nullable Output<Boolean> summaryOnly) {
            $.summaryOnly = summaryOnly;
            return this;
        }

        public Builder summaryOnly(Boolean summaryOnly) {
            return summaryOnly(Output.of(summaryOnly));
        }

        public GetOrganizationTeamsArgs build() {
            return $;
        }
    }

}
