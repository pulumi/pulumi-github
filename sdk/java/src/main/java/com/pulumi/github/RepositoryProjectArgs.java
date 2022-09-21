// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryProjectArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryProjectArgs Empty = new RepositoryProjectArgs();

    @Import(name="body")
    private @Nullable Output<String> body;

    public Optional<Output<String>> body() {
        return Optional.ofNullable(this.body);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="repository", required=true)
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }

    private RepositoryProjectArgs() {}

    private RepositoryProjectArgs(RepositoryProjectArgs $) {
        this.body = $.body;
        this.name = $.name;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryProjectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryProjectArgs $;

        public Builder() {
            $ = new RepositoryProjectArgs();
        }

        public Builder(RepositoryProjectArgs defaults) {
            $ = new RepositoryProjectArgs(Objects.requireNonNull(defaults));
        }

        public Builder body(@Nullable Output<String> body) {
            $.body = body;
            return this;
        }

        public Builder body(String body) {
            return body(Output.of(body));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public RepositoryProjectArgs build() {
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            return $;
        }
    }

}
