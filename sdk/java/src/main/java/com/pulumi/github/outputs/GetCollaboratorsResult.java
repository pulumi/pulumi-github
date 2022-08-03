// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetCollaboratorsCollaborator;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCollaboratorsResult {
    private final @Nullable String affiliation;
    /**
     * @return An Array of GitHub collaborators.  Each `collaborator` block consists of the fields documented below.
     * 
     */
    private final List<GetCollaboratorsCollaborator> collaborators;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final String owner;
    private final String repository;

    @CustomType.Constructor
    private GetCollaboratorsResult(
        @CustomType.Parameter("affiliation") @Nullable String affiliation,
        @CustomType.Parameter("collaborators") List<GetCollaboratorsCollaborator> collaborators,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("owner") String owner,
        @CustomType.Parameter("repository") String repository) {
        this.affiliation = affiliation;
        this.collaborators = collaborators;
        this.id = id;
        this.owner = owner;
        this.repository = repository;
    }

    public Optional<String> affiliation() {
        return Optional.ofNullable(this.affiliation);
    }
    /**
     * @return An Array of GitHub collaborators.  Each `collaborator` block consists of the fields documented below.
     * 
     */
    public List<GetCollaboratorsCollaborator> collaborators() {
        return this.collaborators;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String owner() {
        return this.owner;
    }
    public String repository() {
        return this.repository;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCollaboratorsResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String affiliation;
        private List<GetCollaboratorsCollaborator> collaborators;
        private String id;
        private String owner;
        private String repository;

        public Builder() {
    	      // Empty
        }

        public Builder(GetCollaboratorsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.affiliation = defaults.affiliation;
    	      this.collaborators = defaults.collaborators;
    	      this.id = defaults.id;
    	      this.owner = defaults.owner;
    	      this.repository = defaults.repository;
        }

        public Builder affiliation(@Nullable String affiliation) {
            this.affiliation = affiliation;
            return this;
        }
        public Builder collaborators(List<GetCollaboratorsCollaborator> collaborators) {
            this.collaborators = Objects.requireNonNull(collaborators);
            return this;
        }
        public Builder collaborators(GetCollaboratorsCollaborator... collaborators) {
            return collaborators(List.of(collaborators));
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        public Builder repository(String repository) {
            this.repository = Objects.requireNonNull(repository);
            return this;
        }        public GetCollaboratorsResult build() {
            return new GetCollaboratorsResult(affiliation, collaborators, id, owner, repository);
        }
    }
}
