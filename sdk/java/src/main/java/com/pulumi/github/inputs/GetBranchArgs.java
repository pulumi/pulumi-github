// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetBranchArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBranchArgs Empty = new GetBranchArgs();

    @Import(name="branch", required=true)
    private Output<String> branch;

    public Output<String> branch() {
        return this.branch;
    }

    @Import(name="repository", required=true)
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }

    private GetBranchArgs() {}

    private GetBranchArgs(GetBranchArgs $) {
        this.branch = $.branch;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBranchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBranchArgs $;

        public Builder() {
            $ = new GetBranchArgs();
        }

        public Builder(GetBranchArgs defaults) {
            $ = new GetBranchArgs(Objects.requireNonNull(defaults));
        }

        public Builder branch(Output<String> branch) {
            $.branch = branch;
            return this;
        }

        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public GetBranchArgs build() {
            $.branch = Objects.requireNonNull($.branch, "expected parameter 'branch' to be non-null");
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            return $;
        }
    }

}
