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


public final class TeamArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeamArgs Empty = new TeamArgs();

    /**
     * Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
     * 
     */
    @Import(name="createDefaultMaintainer")
    private @Nullable Output<Boolean> createDefaultMaintainer;

    /**
     * @return Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
     * 
     */
    public Optional<Output<Boolean>> createDefaultMaintainer() {
        return Optional.ofNullable(this.createDefaultMaintainer);
    }

    /**
     * A description of the team.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the team.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
     * 
     */
    @Import(name="ldapDn")
    private @Nullable Output<String> ldapDn;

    /**
     * @return The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
     * 
     */
    public Optional<Output<String>> ldapDn() {
        return Optional.ofNullable(this.ldapDn);
    }

    /**
     * The name of the team.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the team.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID or slug of the parent team, if this is a nested team.
     * 
     */
    @Import(name="parentTeamId")
    private @Nullable Output<String> parentTeamId;

    /**
     * @return The ID or slug of the parent team, if this is a nested team.
     * 
     */
    public Optional<Output<String>> parentTeamId() {
        return Optional.ofNullable(this.parentTeamId);
    }

    /**
     * The id of the parent team read in Github.
     * 
     */
    @Import(name="parentTeamReadId")
    private @Nullable Output<String> parentTeamReadId;

    /**
     * @return The id of the parent team read in Github.
     * 
     */
    public Optional<Output<String>> parentTeamReadId() {
        return Optional.ofNullable(this.parentTeamReadId);
    }

    /**
     * The id of the parent team read in Github.
     * 
     */
    @Import(name="parentTeamReadSlug")
    private @Nullable Output<String> parentTeamReadSlug;

    /**
     * @return The id of the parent team read in Github.
     * 
     */
    public Optional<Output<String>> parentTeamReadSlug() {
        return Optional.ofNullable(this.parentTeamReadSlug);
    }

    /**
     * The level of privacy for the team. Must be one of `secret` or `closed`.
     * Defaults to `secret`.
     * 
     */
    @Import(name="privacy")
    private @Nullable Output<String> privacy;

    /**
     * @return The level of privacy for the team. Must be one of `secret` or `closed`.
     * Defaults to `secret`.
     * 
     */
    public Optional<Output<String>> privacy() {
        return Optional.ofNullable(this.privacy);
    }

    private TeamArgs() {}

    private TeamArgs(TeamArgs $) {
        this.createDefaultMaintainer = $.createDefaultMaintainer;
        this.description = $.description;
        this.ldapDn = $.ldapDn;
        this.name = $.name;
        this.parentTeamId = $.parentTeamId;
        this.parentTeamReadId = $.parentTeamReadId;
        this.parentTeamReadSlug = $.parentTeamReadSlug;
        this.privacy = $.privacy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamArgs $;

        public Builder() {
            $ = new TeamArgs();
        }

        public Builder(TeamArgs defaults) {
            $ = new TeamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param createDefaultMaintainer Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
         * 
         * @return builder
         * 
         */
        public Builder createDefaultMaintainer(@Nullable Output<Boolean> createDefaultMaintainer) {
            $.createDefaultMaintainer = createDefaultMaintainer;
            return this;
        }

        /**
         * @param createDefaultMaintainer Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
         * 
         * @return builder
         * 
         */
        public Builder createDefaultMaintainer(Boolean createDefaultMaintainer) {
            return createDefaultMaintainer(Output.of(createDefaultMaintainer));
        }

        /**
         * @param description A description of the team.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the team.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ldapDn The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
         * 
         * @return builder
         * 
         */
        public Builder ldapDn(@Nullable Output<String> ldapDn) {
            $.ldapDn = ldapDn;
            return this;
        }

        /**
         * @param ldapDn The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
         * 
         * @return builder
         * 
         */
        public Builder ldapDn(String ldapDn) {
            return ldapDn(Output.of(ldapDn));
        }

        /**
         * @param name The name of the team.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the team.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentTeamId The ID or slug of the parent team, if this is a nested team.
         * 
         * @return builder
         * 
         */
        public Builder parentTeamId(@Nullable Output<String> parentTeamId) {
            $.parentTeamId = parentTeamId;
            return this;
        }

        /**
         * @param parentTeamId The ID or slug of the parent team, if this is a nested team.
         * 
         * @return builder
         * 
         */
        public Builder parentTeamId(String parentTeamId) {
            return parentTeamId(Output.of(parentTeamId));
        }

        /**
         * @param parentTeamReadId The id of the parent team read in Github.
         * 
         * @return builder
         * 
         */
        public Builder parentTeamReadId(@Nullable Output<String> parentTeamReadId) {
            $.parentTeamReadId = parentTeamReadId;
            return this;
        }

        /**
         * @param parentTeamReadId The id of the parent team read in Github.
         * 
         * @return builder
         * 
         */
        public Builder parentTeamReadId(String parentTeamReadId) {
            return parentTeamReadId(Output.of(parentTeamReadId));
        }

        /**
         * @param parentTeamReadSlug The id of the parent team read in Github.
         * 
         * @return builder
         * 
         */
        public Builder parentTeamReadSlug(@Nullable Output<String> parentTeamReadSlug) {
            $.parentTeamReadSlug = parentTeamReadSlug;
            return this;
        }

        /**
         * @param parentTeamReadSlug The id of the parent team read in Github.
         * 
         * @return builder
         * 
         */
        public Builder parentTeamReadSlug(String parentTeamReadSlug) {
            return parentTeamReadSlug(Output.of(parentTeamReadSlug));
        }

        /**
         * @param privacy The level of privacy for the team. Must be one of `secret` or `closed`.
         * Defaults to `secret`.
         * 
         * @return builder
         * 
         */
        public Builder privacy(@Nullable Output<String> privacy) {
            $.privacy = privacy;
            return this;
        }

        /**
         * @param privacy The level of privacy for the team. Must be one of `secret` or `closed`.
         * Defaults to `secret`.
         * 
         * @return builder
         * 
         */
        public Builder privacy(String privacy) {
            return privacy(Output.of(privacy));
        }

        public TeamArgs build() {
            return $;
        }
    }

}
