// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppInstallationRepositoryState extends com.pulumi.resources.ResourceArgs {

    public static final AppInstallationRepositoryState Empty = new AppInstallationRepositoryState();

    /**
     * The GitHub app installation id.
     * 
     */
    @Import(name="installationId")
    private @Nullable Output<String> installationId;

    /**
     * @return The GitHub app installation id.
     * 
     */
    public Optional<Output<String>> installationId() {
        return Optional.ofNullable(this.installationId);
    }

    @Import(name="repoId")
    private @Nullable Output<Integer> repoId;

    public Optional<Output<Integer>> repoId() {
        return Optional.ofNullable(this.repoId);
    }

    /**
     * The repository to install the app on.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return The repository to install the app on.
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    private AppInstallationRepositoryState() {}

    private AppInstallationRepositoryState(AppInstallationRepositoryState $) {
        this.installationId = $.installationId;
        this.repoId = $.repoId;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppInstallationRepositoryState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppInstallationRepositoryState $;

        public Builder() {
            $ = new AppInstallationRepositoryState();
        }

        public Builder(AppInstallationRepositoryState defaults) {
            $ = new AppInstallationRepositoryState(Objects.requireNonNull(defaults));
        }

        /**
         * @param installationId The GitHub app installation id.
         * 
         * @return builder
         * 
         */
        public Builder installationId(@Nullable Output<String> installationId) {
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

        public Builder repoId(@Nullable Output<Integer> repoId) {
            $.repoId = repoId;
            return this;
        }

        public Builder repoId(Integer repoId) {
            return repoId(Output.of(repoId));
        }

        /**
         * @param repository The repository to install the app on.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
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

        public AppInstallationRepositoryState build() {
            return $;
        }
    }

}
