// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class RepositoryDependabotSecurityUpdatesArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryDependabotSecurityUpdatesArgs Empty = new RepositoryDependabotSecurityUpdatesArgs();

    /**
     * The state of the automated security fixes.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return The state of the automated security fixes.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * The name of the GitHub repository.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The name of the GitHub repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private RepositoryDependabotSecurityUpdatesArgs() {}

    private RepositoryDependabotSecurityUpdatesArgs(RepositoryDependabotSecurityUpdatesArgs $) {
        this.enabled = $.enabled;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryDependabotSecurityUpdatesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryDependabotSecurityUpdatesArgs $;

        public Builder() {
            $ = new RepositoryDependabotSecurityUpdatesArgs();
        }

        public Builder(RepositoryDependabotSecurityUpdatesArgs defaults) {
            $ = new RepositoryDependabotSecurityUpdatesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled The state of the automated security fixes.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled The state of the automated security fixes.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param repository The name of the GitHub repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The name of the GitHub repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public RepositoryDependabotSecurityUpdatesArgs build() {
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("RepositoryDependabotSecurityUpdatesArgs", "enabled");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("RepositoryDependabotSecurityUpdatesArgs", "repository");
            }
            return $;
        }
    }

}
