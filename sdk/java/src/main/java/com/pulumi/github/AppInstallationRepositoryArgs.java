// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class AppInstallationRepositoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppInstallationRepositoryArgs Empty = new AppInstallationRepositoryArgs();

    /**
     * The GitHub app installation id.
     * 
     */
    @Import(name="installationId", required=true)
    private Output<String> installationId;

    /**
     * @return The GitHub app installation id.
     * 
     */
    public Output<String> installationId() {
        return this.installationId;
    }

    /**
     * The repository to install the app on.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The repository to install the app on.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private AppInstallationRepositoryArgs() {}

    private AppInstallationRepositoryArgs(AppInstallationRepositoryArgs $) {
        this.installationId = $.installationId;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppInstallationRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppInstallationRepositoryArgs $;

        public Builder() {
            $ = new AppInstallationRepositoryArgs();
        }

        public Builder(AppInstallationRepositoryArgs defaults) {
            $ = new AppInstallationRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param installationId The GitHub app installation id.
         * 
         * @return builder
         * 
         */
        public Builder installationId(Output<String> installationId) {
            $.installationId = installationId;
            return this;
        }

        /**
         * @param installationId The GitHub app installation id.
         * 
         * @return builder
         * 
         */
        public Builder installationId(String installationId) {
            return installationId(Output.of(installationId));
        }

        /**
         * @param repository The repository to install the app on.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The repository to install the app on.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public AppInstallationRepositoryArgs build() {
            if ($.installationId == null) {
                throw new MissingRequiredPropertyException("AppInstallationRepositoryArgs", "installationId");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("AppInstallationRepositoryArgs", "repository");
            }
            return $;
        }
    }

}
