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


public final class ProjectCardState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectCardState Empty = new ProjectCardState();

    @Import(name="cardId")
    private @Nullable Output<Integer> cardId;

    public Optional<Output<Integer>> cardId() {
        return Optional.ofNullable(this.cardId);
    }

    @Import(name="columnId")
    private @Nullable Output<String> columnId;

    public Optional<Output<String>> columnId() {
        return Optional.ofNullable(this.columnId);
    }

    @Import(name="contentId")
    private @Nullable Output<Integer> contentId;

    public Optional<Output<Integer>> contentId() {
        return Optional.ofNullable(this.contentId);
    }

    @Import(name="contentType")
    private @Nullable Output<String> contentType;

    public Optional<Output<String>> contentType() {
        return Optional.ofNullable(this.contentType);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    @Import(name="note")
    private @Nullable Output<String> note;

    public Optional<Output<String>> note() {
        return Optional.ofNullable(this.note);
    }

    private ProjectCardState() {}

    private ProjectCardState(ProjectCardState $) {
        this.cardId = $.cardId;
        this.columnId = $.columnId;
        this.contentId = $.contentId;
        this.contentType = $.contentType;
        this.etag = $.etag;
        this.note = $.note;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectCardState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectCardState $;

        public Builder() {
            $ = new ProjectCardState();
        }

        public Builder(ProjectCardState defaults) {
            $ = new ProjectCardState(Objects.requireNonNull(defaults));
        }

        public Builder cardId(@Nullable Output<Integer> cardId) {
            $.cardId = cardId;
            return this;
        }

        public Builder cardId(Integer cardId) {
            return cardId(Output.of(cardId));
        }

        public Builder columnId(@Nullable Output<String> columnId) {
            $.columnId = columnId;
            return this;
        }

        public Builder columnId(String columnId) {
            return columnId(Output.of(columnId));
        }

        public Builder contentId(@Nullable Output<Integer> contentId) {
            $.contentId = contentId;
            return this;
        }

        public Builder contentId(Integer contentId) {
            return contentId(Output.of(contentId));
        }

        public Builder contentType(@Nullable Output<String> contentType) {
            $.contentType = contentType;
            return this;
        }

        public Builder contentType(String contentType) {
            return contentType(Output.of(contentType));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        public Builder note(@Nullable Output<String> note) {
            $.note = note;
            return this;
        }

        public Builder note(String note) {
            return note(Output.of(note));
        }

        public ProjectCardState build() {
            return $;
        }
    }

}
