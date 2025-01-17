// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.RepositoryCollaboratorsIgnoreTeamArgs;
import com.pulumi.github.inputs.RepositoryCollaboratorsTeamArgs;
import com.pulumi.github.inputs.RepositoryCollaboratorsUserArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryCollaboratorsState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryCollaboratorsState Empty = new RepositoryCollaboratorsState();

    /**
     * List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
     * 
     */
    @Import(name="ignoreTeams")
    private @Nullable Output<List<RepositoryCollaboratorsIgnoreTeamArgs>> ignoreTeams;

    /**
     * @return List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
     * 
     */
    public Optional<Output<List<RepositoryCollaboratorsIgnoreTeamArgs>>> ignoreTeams() {
        return Optional.ofNullable(this.ignoreTeams);
    }

    /**
     * Map of usernames to invitation ID for any users added as part of creation of this resource to
     * be used in `github.UserInvitationAccepter`.
     * 
     */
    @Import(name="invitationIds")
    private @Nullable Output<Map<String,String>> invitationIds;

    /**
     * @return Map of usernames to invitation ID for any users added as part of creation of this resource to
     * be used in `github.UserInvitationAccepter`.
     * 
     */
    public Optional<Output<Map<String,String>>> invitationIds() {
        return Optional.ofNullable(this.invitationIds);
    }

    /**
     * The GitHub repository.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return The GitHub repository.
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * List of teams to grant access to the repository.
     * 
     */
    @Import(name="teams")
    private @Nullable Output<List<RepositoryCollaboratorsTeamArgs>> teams;

    /**
     * @return List of teams to grant access to the repository.
     * 
     */
    public Optional<Output<List<RepositoryCollaboratorsTeamArgs>>> teams() {
        return Optional.ofNullable(this.teams);
    }

    /**
     * List of users to grant access to the repository.
     * 
     */
    @Import(name="users")
    private @Nullable Output<List<RepositoryCollaboratorsUserArgs>> users;

    /**
     * @return List of users to grant access to the repository.
     * 
     */
    public Optional<Output<List<RepositoryCollaboratorsUserArgs>>> users() {
        return Optional.ofNullable(this.users);
    }

    private RepositoryCollaboratorsState() {}

    private RepositoryCollaboratorsState(RepositoryCollaboratorsState $) {
        this.ignoreTeams = $.ignoreTeams;
        this.invitationIds = $.invitationIds;
        this.repository = $.repository;
        this.teams = $.teams;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryCollaboratorsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryCollaboratorsState $;

        public Builder() {
            $ = new RepositoryCollaboratorsState();
        }

        public Builder(RepositoryCollaboratorsState defaults) {
            $ = new RepositoryCollaboratorsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ignoreTeams List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
         * 
         * @return builder
         * 
         */
        public Builder ignoreTeams(@Nullable Output<List<RepositoryCollaboratorsIgnoreTeamArgs>> ignoreTeams) {
            $.ignoreTeams = ignoreTeams;
            return this;
        }

        /**
         * @param ignoreTeams List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
         * 
         * @return builder
         * 
         */
        public Builder ignoreTeams(List<RepositoryCollaboratorsIgnoreTeamArgs> ignoreTeams) {
            return ignoreTeams(Output.of(ignoreTeams));
        }

        /**
         * @param ignoreTeams List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
         * 
         * @return builder
         * 
         */
        public Builder ignoreTeams(RepositoryCollaboratorsIgnoreTeamArgs... ignoreTeams) {
            return ignoreTeams(List.of(ignoreTeams));
        }

        /**
         * @param invitationIds Map of usernames to invitation ID for any users added as part of creation of this resource to
         * be used in `github.UserInvitationAccepter`.
         * 
         * @return builder
         * 
         */
        public Builder invitationIds(@Nullable Output<Map<String,String>> invitationIds) {
            $.invitationIds = invitationIds;
            return this;
        }

        /**
         * @param invitationIds Map of usernames to invitation ID for any users added as part of creation of this resource to
         * be used in `github.UserInvitationAccepter`.
         * 
         * @return builder
         * 
         */
        public Builder invitationIds(Map<String,String> invitationIds) {
            return invitationIds(Output.of(invitationIds));
        }

        /**
         * @param repository The GitHub repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The GitHub repository.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param teams List of teams to grant access to the repository.
         * 
         * @return builder
         * 
         */
        public Builder teams(@Nullable Output<List<RepositoryCollaboratorsTeamArgs>> teams) {
            $.teams = teams;
            return this;
        }

        /**
         * @param teams List of teams to grant access to the repository.
         * 
         * @return builder
         * 
         */
        public Builder teams(List<RepositoryCollaboratorsTeamArgs> teams) {
            return teams(Output.of(teams));
        }

        /**
         * @param teams List of teams to grant access to the repository.
         * 
         * @return builder
         * 
         */
        public Builder teams(RepositoryCollaboratorsTeamArgs... teams) {
            return teams(List.of(teams));
        }

        /**
         * @param users List of users to grant access to the repository.
         * 
         * @return builder
         * 
         */
        public Builder users(@Nullable Output<List<RepositoryCollaboratorsUserArgs>> users) {
            $.users = users;
            return this;
        }

        /**
         * @param users List of users to grant access to the repository.
         * 
         * @return builder
         * 
         */
        public Builder users(List<RepositoryCollaboratorsUserArgs> users) {
            return users(Output.of(users));
        }

        /**
         * @param users List of users to grant access to the repository.
         * 
         * @return builder
         * 
         */
        public Builder users(RepositoryCollaboratorsUserArgs... users) {
            return users(List.of(users));
        }

        public RepositoryCollaboratorsState build() {
            return $;
        }
    }

}
