// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCollaboratorsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCollaboratorsArgs Empty = new GetCollaboratorsArgs();

    /**
     * Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
     * 
     */
    @Import(name="affiliation")
    private @Nullable Output<String> affiliation;

    /**
     * @return Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
     * 
     */
    public Optional<Output<String>> affiliation() {
        return Optional.ofNullable(this.affiliation);
    }

    /**
     * The organization that owns the repository.
     * 
     */
    @Import(name="owner", required=true)
    private Output<String> owner;

    /**
     * @return The organization that owns the repository.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }

    /**
     * Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
     * 
     */
    @Import(name="permission")
    private @Nullable Output<String> permission;

    /**
     * @return Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
     * 
     */
    public Optional<Output<String>> permission() {
        return Optional.ofNullable(this.permission);
    }

    /**
     * The name of the repository.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The name of the repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private GetCollaboratorsArgs() {}

    private GetCollaboratorsArgs(GetCollaboratorsArgs $) {
        this.affiliation = $.affiliation;
        this.owner = $.owner;
        this.permission = $.permission;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCollaboratorsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCollaboratorsArgs $;

        public Builder() {
            $ = new GetCollaboratorsArgs();
        }

        public Builder(GetCollaboratorsArgs defaults) {
            $ = new GetCollaboratorsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param affiliation Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
         * 
         * @return builder
         * 
         */
        public Builder affiliation(@Nullable Output<String> affiliation) {
            $.affiliation = affiliation;
            return this;
        }

        /**
         * @param affiliation Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
         * 
         * @return builder
         * 
         */
        public Builder affiliation(String affiliation) {
            return affiliation(Output.of(affiliation));
        }

        /**
         * @param owner The organization that owns the repository.
         * 
         * @return builder
         * 
         */
        public Builder owner(Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The organization that owns the repository.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param permission Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
         * 
         * @return builder
         * 
         */
        public Builder permission(@Nullable Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        /**
         * @param repository The name of the repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The name of the repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public GetCollaboratorsArgs build() {
            if ($.owner == null) {
                throw new MissingRequiredPropertyException("GetCollaboratorsArgs", "owner");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("GetCollaboratorsArgs", "repository");
            }
            return $;
        }
    }

}