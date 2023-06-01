// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetRepositoryEnvironmentsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryEnvironmentsArgs Empty = new GetRepositoryEnvironmentsArgs();

    /**
     * Name of the repository to retrieve the environments from.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return Name of the repository to retrieve the environments from.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private GetRepositoryEnvironmentsArgs() {}

    private GetRepositoryEnvironmentsArgs(GetRepositoryEnvironmentsArgs $) {
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryEnvironmentsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryEnvironmentsArgs $;

        public Builder() {
            $ = new GetRepositoryEnvironmentsArgs();
        }

        public Builder(GetRepositoryEnvironmentsArgs defaults) {
            $ = new GetRepositoryEnvironmentsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param repository Name of the repository to retrieve the environments from.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository Name of the repository to retrieve the environments from.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public GetRepositoryEnvironmentsArgs build() {
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            return $;
        }
    }

}
