// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetReleasePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetReleasePlainArgs Empty = new GetReleasePlainArgs();

    /**
     * Owner of the repository.
     * 
     */
    @Import(name="owner", required=true)
    private String owner;

    /**
     * @return Owner of the repository.
     * 
     */
    public String owner() {
        return this.owner;
    }

    /**
     * ID of the release to retrieve. Must be specified when `retrieve_by` = `id`.
     * 
     */
    @Import(name="releaseId")
    private @Nullable Integer releaseId;

    /**
     * @return ID of the release to retrieve. Must be specified when `retrieve_by` = `id`.
     * 
     */
    public Optional<Integer> releaseId() {
        return Optional.ofNullable(this.releaseId);
    }

    /**
     * Tag of the release to retrieve. Must be specified when `retrieve_by` = `tag`.
     * 
     */
    @Import(name="releaseTag")
    private @Nullable String releaseTag;

    /**
     * @return Tag of the release to retrieve. Must be specified when `retrieve_by` = `tag`.
     * 
     */
    public Optional<String> releaseTag() {
        return Optional.ofNullable(this.releaseTag);
    }

    /**
     * Name of the repository to retrieve the release from.
     * 
     */
    @Import(name="repository", required=true)
    private String repository;

    /**
     * @return Name of the repository to retrieve the release from.
     * 
     */
    public String repository() {
        return this.repository;
    }

    /**
     * Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
     * 
     */
    @Import(name="retrieveBy", required=true)
    private String retrieveBy;

    /**
     * @return Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
     * 
     */
    public String retrieveBy() {
        return this.retrieveBy;
    }

    private GetReleasePlainArgs() {}

    private GetReleasePlainArgs(GetReleasePlainArgs $) {
        this.owner = $.owner;
        this.releaseId = $.releaseId;
        this.releaseTag = $.releaseTag;
        this.repository = $.repository;
        this.retrieveBy = $.retrieveBy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReleasePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReleasePlainArgs $;

        public Builder() {
            $ = new GetReleasePlainArgs();
        }

        public Builder(GetReleasePlainArgs defaults) {
            $ = new GetReleasePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param owner Owner of the repository.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param releaseId ID of the release to retrieve. Must be specified when `retrieve_by` = `id`.
         * 
         * @return builder
         * 
         */
        public Builder releaseId(@Nullable Integer releaseId) {
            $.releaseId = releaseId;
            return this;
        }

        /**
         * @param releaseTag Tag of the release to retrieve. Must be specified when `retrieve_by` = `tag`.
         * 
         * @return builder
         * 
         */
        public Builder releaseTag(@Nullable String releaseTag) {
            $.releaseTag = releaseTag;
            return this;
        }

        /**
         * @param repository Name of the repository to retrieve the release from.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param retrieveBy Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
         * 
         * @return builder
         * 
         */
        public Builder retrieveBy(String retrieveBy) {
            $.retrieveBy = retrieveBy;
            return this;
        }

        public GetReleasePlainArgs build() {
            if ($.owner == null) {
                throw new MissingRequiredPropertyException("GetReleasePlainArgs", "owner");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetReleasePlainArgs", "repository");
            }
            if ($.retrieveBy == null) {
                throw new MissingRequiredPropertyException("GetReleasePlainArgs", "retrieveBy");
            }
            return $;
        }
    }

}
