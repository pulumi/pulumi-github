// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetIssueLabelsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIssueLabelsPlainArgs Empty = new GetIssueLabelsPlainArgs();

    /**
     * The name of the repository.
     * 
     */
    @Import(name="repository", required=true)
    private String repository;

    /**
     * @return The name of the repository.
     * 
     */
    public String repository() {
        return this.repository;
    }

    private GetIssueLabelsPlainArgs() {}

    private GetIssueLabelsPlainArgs(GetIssueLabelsPlainArgs $) {
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIssueLabelsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIssueLabelsPlainArgs $;

        public Builder() {
            $ = new GetIssueLabelsPlainArgs();
        }

        public Builder(GetIssueLabelsPlainArgs defaults) {
            $ = new GetIssueLabelsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param repository The name of the repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            $.repository = repository;
            return this;
        }

        public GetIssueLabelsPlainArgs build() {
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            return $;
        }
    }

}
