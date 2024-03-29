// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetRepositoryEnvironmentsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryEnvironmentsPlainArgs Empty = new GetRepositoryEnvironmentsPlainArgs();

    /**
     * Name of the repository to retrieve the environments from.
     * 
     */
    @Import(name="repository", required=true)
    private String repository;

    /**
     * @return Name of the repository to retrieve the environments from.
     * 
     */
    public String repository() {
        return this.repository;
    }

    private GetRepositoryEnvironmentsPlainArgs() {}

    private GetRepositoryEnvironmentsPlainArgs(GetRepositoryEnvironmentsPlainArgs $) {
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryEnvironmentsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryEnvironmentsPlainArgs $;

        public Builder() {
            $ = new GetRepositoryEnvironmentsPlainArgs();
        }

        public Builder(GetRepositoryEnvironmentsPlainArgs defaults) {
            $ = new GetRepositoryEnvironmentsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param repository Name of the repository to retrieve the environments from.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            $.repository = repository;
            return this;
        }

        public GetRepositoryEnvironmentsPlainArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetRepositoryEnvironmentsPlainArgs", "repository");
            }
            return $;
        }
    }

}
