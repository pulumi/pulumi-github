// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReleaseArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseArgs Empty = new ReleaseArgs();

    @Import(name="body")
    private @Nullable Output<String> body;

    public Optional<Output<String>> body() {
        return Optional.ofNullable(this.body);
    }

    @Import(name="discussionCategoryName")
    private @Nullable Output<String> discussionCategoryName;

    public Optional<Output<String>> discussionCategoryName() {
        return Optional.ofNullable(this.discussionCategoryName);
    }

    @Import(name="draft")
    private @Nullable Output<Boolean> draft;

    public Optional<Output<Boolean>> draft() {
        return Optional.ofNullable(this.draft);
    }

    @Import(name="generateReleaseNotes")
    private @Nullable Output<Boolean> generateReleaseNotes;

    public Optional<Output<Boolean>> generateReleaseNotes() {
        return Optional.ofNullable(this.generateReleaseNotes);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="prerelease")
    private @Nullable Output<Boolean> prerelease;

    public Optional<Output<Boolean>> prerelease() {
        return Optional.ofNullable(this.prerelease);
    }

    @Import(name="repository", required=true)
    private Output<String> repository;

    public Output<String> repository() {
        return this.repository;
    }

    @Import(name="tagName", required=true)
    private Output<String> tagName;

    public Output<String> tagName() {
        return this.tagName;
    }

    @Import(name="targetCommitish")
    private @Nullable Output<String> targetCommitish;

    public Optional<Output<String>> targetCommitish() {
        return Optional.ofNullable(this.targetCommitish);
    }

    private ReleaseArgs() {}

    private ReleaseArgs(ReleaseArgs $) {
        this.body = $.body;
        this.discussionCategoryName = $.discussionCategoryName;
        this.draft = $.draft;
        this.generateReleaseNotes = $.generateReleaseNotes;
        this.name = $.name;
        this.prerelease = $.prerelease;
        this.repository = $.repository;
        this.tagName = $.tagName;
        this.targetCommitish = $.targetCommitish;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReleaseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseArgs $;

        public Builder() {
            $ = new ReleaseArgs();
        }

        public Builder(ReleaseArgs defaults) {
            $ = new ReleaseArgs(Objects.requireNonNull(defaults));
        }

        public Builder body(@Nullable Output<String> body) {
            $.body = body;
            return this;
        }

        public Builder body(String body) {
            return body(Output.of(body));
        }

        public Builder discussionCategoryName(@Nullable Output<String> discussionCategoryName) {
            $.discussionCategoryName = discussionCategoryName;
            return this;
        }

        public Builder discussionCategoryName(String discussionCategoryName) {
            return discussionCategoryName(Output.of(discussionCategoryName));
        }

        public Builder draft(@Nullable Output<Boolean> draft) {
            $.draft = draft;
            return this;
        }

        public Builder draft(Boolean draft) {
            return draft(Output.of(draft));
        }

        public Builder generateReleaseNotes(@Nullable Output<Boolean> generateReleaseNotes) {
            $.generateReleaseNotes = generateReleaseNotes;
            return this;
        }

        public Builder generateReleaseNotes(Boolean generateReleaseNotes) {
            return generateReleaseNotes(Output.of(generateReleaseNotes));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder prerelease(@Nullable Output<Boolean> prerelease) {
            $.prerelease = prerelease;
            return this;
        }

        public Builder prerelease(Boolean prerelease) {
            return prerelease(Output.of(prerelease));
        }

        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public Builder tagName(Output<String> tagName) {
            $.tagName = tagName;
            return this;
        }

        public Builder tagName(String tagName) {
            return tagName(Output.of(tagName));
        }

        public Builder targetCommitish(@Nullable Output<String> targetCommitish) {
            $.targetCommitish = targetCommitish;
            return this;
        }

        public Builder targetCommitish(String targetCommitish) {
            return targetCommitish(Output.of(targetCommitish));
        }

        public ReleaseArgs build() {
            $.repository = Objects.requireNonNull($.repository, "expected parameter 'repository' to be non-null");
            $.tagName = Objects.requireNonNull($.tagName, "expected parameter 'tagName' to be non-null");
            return $;
        }
    }

}
