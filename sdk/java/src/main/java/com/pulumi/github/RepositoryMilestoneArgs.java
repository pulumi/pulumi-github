// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryMilestoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryMilestoneArgs Empty = new RepositoryMilestoneArgs();

    /**
     * A description of the milestone.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the milestone.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The milestone due date. In `yyyy-mm-dd` format.
     * 
     */
    @Import(name="dueDate")
    private @Nullable Output<String> dueDate;

    /**
     * @return The milestone due date. In `yyyy-mm-dd` format.
     * 
     */
    public Optional<Output<String>> dueDate() {
        return Optional.ofNullable(this.dueDate);
    }

    /**
     * The owner of the GitHub Repository.
     * 
     */
    @Import(name="owner", required=true)
    private Output<String> owner;

    /**
     * @return The owner of the GitHub Repository.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }

    /**
     * The name of the GitHub Repository.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The name of the GitHub Repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     * The state of the milestone. Either `open` or `closed`. Default: `open`
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return The state of the milestone. Either `open` or `closed`. Default: `open`
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    /**
     * The title of the milestone.
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return The title of the milestone.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    private RepositoryMilestoneArgs() {}

    private RepositoryMilestoneArgs(RepositoryMilestoneArgs $) {
        this.description = $.description;
        this.dueDate = $.dueDate;
        this.owner = $.owner;
        this.repository = $.repository;
        this.state = $.state;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryMilestoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryMilestoneArgs $;

        public Builder() {
            $ = new RepositoryMilestoneArgs();
        }

        public Builder(RepositoryMilestoneArgs defaults) {
            $ = new RepositoryMilestoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description of the milestone.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the milestone.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dueDate The milestone due date. In `yyyy-mm-dd` format.
         * 
         * @return builder
         * 
         */
        public Builder dueDate(@Nullable Output<String> dueDate) {
            $.dueDate = dueDate;
            return this;
        }

        /**
         * @param dueDate The milestone due date. In `yyyy-mm-dd` format.
         * 
         * @return builder
         * 
         */
        public Builder dueDate(String dueDate) {
            return dueDate(Output.of(dueDate));
        }

        /**
         * @param owner The owner of the GitHub Repository.
         * 
         * @return builder
         * 
         */
        public Builder owner(Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The owner of the GitHub Repository.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param repository The name of the GitHub Repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The name of the GitHub Repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param state The state of the milestone. Either `open` or `closed`. Default: `open`
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state The state of the milestone. Either `open` or `closed`. Default: `open`
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        /**
         * @param title The title of the milestone.
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title The title of the milestone.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public RepositoryMilestoneArgs build() {
            if ($.owner == null) {
                throw new MissingRequiredPropertyException("RepositoryMilestoneArgs", "owner");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("RepositoryMilestoneArgs", "repository");
            }
            if ($.title == null) {
                throw new MissingRequiredPropertyException("RepositoryMilestoneArgs", "title");
            }
            return $;
        }
    }

}
