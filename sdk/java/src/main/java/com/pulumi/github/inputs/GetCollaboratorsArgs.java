// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCollaboratorsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCollaboratorsArgs Empty = new GetCollaboratorsArgs();

    @Import(name="affiliation")
    private @Nullable Output<String> affiliation;

    public Optional<Output<String>> affiliation() {
        return Optional.ofNullable(this.affiliation);
    }

    @Import(name="owner", required=true)
    private Output<String> owner;

    public Output<String> owner() {
        return this.owner;
    }

    @Import(name="repository", required=true)
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }

    private GetCollaboratorsArgs() {}

    private GetCollaboratorsArgs(GetCollaboratorsArgs $) {
        this.affiliation = $.affiliation;
        this.owner = $.owner;
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

        public Builder affiliation(@Nullable Output<String> affiliation) {
            $.affiliation = affiliation;
            return this;
        }

        public Builder affiliation(String affiliation) {
            return affiliation(Output.of(affiliation));
        }

        public Builder owner(Output<String> owner) {
            $.owner = owner;
            return this;
        }

        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public GetCollaboratorsArgs build() {
            $.owner = Objects.requireNonNull($.owner, "expected parameter 'owner' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            return $;
        }
    }

}
