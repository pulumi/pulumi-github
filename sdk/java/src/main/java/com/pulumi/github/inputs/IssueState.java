// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IssueState extends com.pulumi.resources.ResourceArgs {

    public static final IssueState Empty = new IssueState();

    /**
     * List of Logins to assign the to the issue
     * 
     */
    @Import(name="assignees")
    private @Nullable Output<List<String>> assignees;

    /**
     * @return List of Logins to assign the to the issue
     * 
     */
    public Optional<Output<List<String>>> assignees() {
        return Optional.ofNullable(this.assignees);
    }

    /**
     * Body of the issue
     * 
     */
    @Import(name="body")
    private @Nullable Output<String> body;

    /**
     * @return Body of the issue
     * 
     */
    public Optional<Output<String>> body() {
        return Optional.ofNullable(this.body);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * (Computed) - The issue id
     * 
     */
    @Import(name="issueId")
    private @Nullable Output<Integer> issueId;

    /**
     * @return (Computed) - The issue id
     * 
     */
    public Optional<Output<Integer>> issueId() {
        return Optional.ofNullable(this.issueId);
    }

    /**
     * List of labels to attach to the issue
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<String>> labels;

    /**
     * @return List of labels to attach to the issue
     * 
     */
    public Optional<Output<List<String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Milestone number to assign to the issue
     * 
     */
    @Import(name="milestoneNumber")
    private @Nullable Output<Integer> milestoneNumber;

    /**
     * @return Milestone number to assign to the issue
     * 
     */
    public Optional<Output<Integer>> milestoneNumber() {
        return Optional.ofNullable(this.milestoneNumber);
    }

    /**
     * (Computed) - The issue number
     * 
     */
    @Import(name="number")
    private @Nullable Output<Integer> number;

    /**
     * @return (Computed) - The issue number
     * 
     */
    public Optional<Output<Integer>> number() {
        return Optional.ofNullable(this.number);
    }

    /**
     * The GitHub repository name
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return The GitHub repository name
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * Title of the issue
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return Title of the issue
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    private IssueState() {}

    private IssueState(IssueState $) {
        this.assignees = $.assignees;
        this.body = $.body;
        this.etag = $.etag;
        this.issueId = $.issueId;
        this.labels = $.labels;
        this.milestoneNumber = $.milestoneNumber;
        this.number = $.number;
        this.repository = $.repository;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IssueState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IssueState $;

        public Builder() {
            $ = new IssueState();
        }

        public Builder(IssueState defaults) {
            $ = new IssueState(Objects.requireNonNull(defaults));
        }

        /**
         * @param assignees List of Logins to assign the to the issue
         * 
         * @return builder
         * 
         */
        public Builder assignees(@Nullable Output<List<String>> assignees) {
            $.assignees = assignees;
            return this;
        }

        /**
         * @param assignees List of Logins to assign the to the issue
         * 
         * @return builder
         * 
         */
        public Builder assignees(List<String> assignees) {
            return assignees(Output.of(assignees));
        }

        /**
         * @param assignees List of Logins to assign the to the issue
         * 
         * @return builder
         * 
         */
        public Builder assignees(String... assignees) {
            return assignees(List.of(assignees));
        }

        /**
         * @param body Body of the issue
         * 
         * @return builder
         * 
         */
        public Builder body(@Nullable Output<String> body) {
            $.body = body;
            return this;
        }

        /**
         * @param body Body of the issue
         * 
         * @return builder
         * 
         */
        public Builder body(String body) {
            return body(Output.of(body));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param issueId (Computed) - The issue id
         * 
         * @return builder
         * 
         */
        public Builder issueId(@Nullable Output<Integer> issueId) {
            $.issueId = issueId;
            return this;
        }

        /**
         * @param issueId (Computed) - The issue id
         * 
         * @return builder
         * 
         */
        public Builder issueId(Integer issueId) {
            return issueId(Output.of(issueId));
        }

        /**
         * @param labels List of labels to attach to the issue
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels List of labels to attach to the issue
         * 
         * @return builder
         * 
         */
        public Builder labels(List<String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels List of labels to attach to the issue
         * 
         * @return builder
         * 
         */
        public Builder labels(String... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param milestoneNumber Milestone number to assign to the issue
         * 
         * @return builder
         * 
         */
        public Builder milestoneNumber(@Nullable Output<Integer> milestoneNumber) {
            $.milestoneNumber = milestoneNumber;
            return this;
        }

        /**
         * @param milestoneNumber Milestone number to assign to the issue
         * 
         * @return builder
         * 
         */
        public Builder milestoneNumber(Integer milestoneNumber) {
            return milestoneNumber(Output.of(milestoneNumber));
        }

        /**
         * @param number (Computed) - The issue number
         * 
         * @return builder
         * 
         */
        public Builder number(@Nullable Output<Integer> number) {
            $.number = number;
            return this;
        }

        /**
         * @param number (Computed) - The issue number
         * 
         * @return builder
         * 
         */
        public Builder number(Integer number) {
            return number(Output.of(number));
        }

        /**
         * @param repository The GitHub repository name
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The GitHub repository name
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param title Title of the issue
         * 
         * @return builder
         * 
         */
        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Title of the issue
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public IssueState build() {
            return $;
        }
    }

}