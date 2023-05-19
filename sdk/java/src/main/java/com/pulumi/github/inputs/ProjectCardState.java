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

    /**
     * The ID of the card.
     * 
     */
    @Import(name="cardId")
    private @Nullable Output<Integer> cardId;

    /**
     * @return The ID of the card.
     * 
     */
    public Optional<Output<Integer>> cardId() {
        return Optional.ofNullable(this.cardId);
    }

    /**
     * The ID of the card.
     * 
     */
    @Import(name="columnId")
    private @Nullable Output<String> columnId;

    /**
     * @return The ID of the card.
     * 
     */
    public Optional<Output<String>> columnId() {
        return Optional.ofNullable(this.columnId);
    }

    /**
     * `github_issue.issue_id`.
     * 
     */
    @Import(name="contentId")
    private @Nullable Output<Integer> contentId;

    /**
     * @return `github_issue.issue_id`.
     * 
     */
    public Optional<Output<Integer>> contentId() {
        return Optional.ofNullable(this.contentId);
    }

    /**
     * Must be either `Issue` or `PullRequest`
     * 
     * **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
     * See note example or issue example for more information.
     * 
     */
    @Import(name="contentType")
    private @Nullable Output<String> contentType;

    /**
     * @return Must be either `Issue` or `PullRequest`
     * 
     * **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
     * See note example or issue example for more information.
     * 
     */
    public Optional<Output<String>> contentType() {
        return Optional.ofNullable(this.contentType);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The note contents of the card. Markdown supported.
     * 
     */
    @Import(name="note")
    private @Nullable Output<String> note;

    /**
     * @return The note contents of the card. Markdown supported.
     * 
     */
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

        /**
         * @param cardId The ID of the card.
         * 
         * @return builder
         * 
         */
        public Builder cardId(@Nullable Output<Integer> cardId) {
            $.cardId = cardId;
            return this;
        }

        /**
         * @param cardId The ID of the card.
         * 
         * @return builder
         * 
         */
        public Builder cardId(Integer cardId) {
            return cardId(Output.of(cardId));
        }

        /**
         * @param columnId The ID of the card.
         * 
         * @return builder
         * 
         */
        public Builder columnId(@Nullable Output<String> columnId) {
            $.columnId = columnId;
            return this;
        }

        /**
         * @param columnId The ID of the card.
         * 
         * @return builder
         * 
         */
        public Builder columnId(String columnId) {
            return columnId(Output.of(columnId));
        }

        /**
         * @param contentId `github_issue.issue_id`.
         * 
         * @return builder
         * 
         */
        public Builder contentId(@Nullable Output<Integer> contentId) {
            $.contentId = contentId;
            return this;
        }

        /**
         * @param contentId `github_issue.issue_id`.
         * 
         * @return builder
         * 
         */
        public Builder contentId(Integer contentId) {
            return contentId(Output.of(contentId));
        }

        /**
         * @param contentType Must be either `Issue` or `PullRequest`
         * 
         * **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
         * See note example or issue example for more information.
         * 
         * @return builder
         * 
         */
        public Builder contentType(@Nullable Output<String> contentType) {
            $.contentType = contentType;
            return this;
        }

        /**
         * @param contentType Must be either `Issue` or `PullRequest`
         * 
         * **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
         * See note example or issue example for more information.
         * 
         * @return builder
         * 
         */
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

        /**
         * @param note The note contents of the card. Markdown supported.
         * 
         * @return builder
         * 
         */
        public Builder note(@Nullable Output<String> note) {
            $.note = note;
            return this;
        }

        /**
         * @param note The note contents of the card. Markdown supported.
         * 
         * @return builder
         * 
         */
        public Builder note(String note) {
            return note(Output.of(note));
        }

        public ProjectCardState build() {
            return $;
        }
    }

}
