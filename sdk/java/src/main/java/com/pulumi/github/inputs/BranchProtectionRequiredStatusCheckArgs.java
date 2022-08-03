// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchProtectionRequiredStatusCheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchProtectionRequiredStatusCheckArgs Empty = new BranchProtectionRequiredStatusCheckArgs();

    @Import(name="contexts")
    private @Nullable Output<List<String>> contexts;

    public Optional<Output<List<String>>> contexts() {
        return Optional.ofNullable(this.contexts);
    }

    @Import(name="strict")
    private @Nullable Output<Boolean> strict;

    public Optional<Output<Boolean>> strict() {
        return Optional.ofNullable(this.strict);
    }

    private BranchProtectionRequiredStatusCheckArgs() {}

    private BranchProtectionRequiredStatusCheckArgs(BranchProtectionRequiredStatusCheckArgs $) {
        this.contexts = $.contexts;
        this.strict = $.strict;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchProtectionRequiredStatusCheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchProtectionRequiredStatusCheckArgs $;

        public Builder() {
            $ = new BranchProtectionRequiredStatusCheckArgs();
        }

        public Builder(BranchProtectionRequiredStatusCheckArgs defaults) {
            $ = new BranchProtectionRequiredStatusCheckArgs(Objects.requireNonNull(defaults));
        }

        public Builder contexts(@Nullable Output<List<String>> contexts) {
            $.contexts = contexts;
            return this;
        }

        public Builder contexts(List<String> contexts) {
            return contexts(Output.of(contexts));
        }

        public Builder contexts(String... contexts) {
            return contexts(List.of(contexts));
        }

        public Builder strict(@Nullable Output<Boolean> strict) {
            $.strict = strict;
            return this;
        }

        public Builder strict(Boolean strict) {
            return strict(Output.of(strict));
        }

        public BranchProtectionRequiredStatusCheckArgs build() {
            return $;
        }
    }

}
