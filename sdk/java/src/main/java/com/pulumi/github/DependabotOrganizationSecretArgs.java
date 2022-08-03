// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DependabotOrganizationSecretArgs extends com.pulumi.resources.ResourceArgs {

    public static final DependabotOrganizationSecretArgs Empty = new DependabotOrganizationSecretArgs();

    @Import(name="encryptedValue")
    private @Nullable Output<String> encryptedValue;

    public Optional<Output<String>> encryptedValue() {
        return Optional.ofNullable(this.encryptedValue);
    }

    @Import(name="plaintextValue")
    private @Nullable Output<String> plaintextValue;

    public Optional<Output<String>> plaintextValue() {
        return Optional.ofNullable(this.plaintextValue);
    }

    @Import(name="secretName", required=true)
    private Output<String> secretName;

    public Output<String> secretName() {
        return this.secretName;
    }

    @Import(name="selectedRepositoryIds")
    private @Nullable Output<List<Integer>> selectedRepositoryIds;

    public Optional<Output<List<Integer>>> selectedRepositoryIds() {
        return Optional.ofNullable(this.selectedRepositoryIds);
    }

    @Import(name="visibility", required=true)
    private Output<String> visibility;

    public Output<String> visibility() {
        return this.visibility;
    }

    private DependabotOrganizationSecretArgs() {}

    private DependabotOrganizationSecretArgs(DependabotOrganizationSecretArgs $) {
        this.encryptedValue = $.encryptedValue;
        this.plaintextValue = $.plaintextValue;
        this.secretName = $.secretName;
        this.selectedRepositoryIds = $.selectedRepositoryIds;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DependabotOrganizationSecretArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DependabotOrganizationSecretArgs $;

        public Builder() {
            $ = new DependabotOrganizationSecretArgs();
        }

        public Builder(DependabotOrganizationSecretArgs defaults) {
            $ = new DependabotOrganizationSecretArgs(Objects.requireNonNull(defaults));
        }

        public Builder encryptedValue(@Nullable Output<String> encryptedValue) {
            $.encryptedValue = encryptedValue;
            return this;
        }

        public Builder encryptedValue(String encryptedValue) {
            return encryptedValue(Output.of(encryptedValue));
        }

        public Builder plaintextValue(@Nullable Output<String> plaintextValue) {
            $.plaintextValue = plaintextValue;
            return this;
        }

        public Builder plaintextValue(String plaintextValue) {
            return plaintextValue(Output.of(plaintextValue));
        }

        public Builder secretName(Output<String> secretName) {
            $.secretName = secretName;
            return this;
        }

        public Builder secretName(String secretName) {
            return secretName(Output.of(secretName));
        }

        public Builder selectedRepositoryIds(@Nullable Output<List<Integer>> selectedRepositoryIds) {
            $.selectedRepositoryIds = selectedRepositoryIds;
            return this;
        }

        public Builder selectedRepositoryIds(List<Integer> selectedRepositoryIds) {
            return selectedRepositoryIds(Output.of(selectedRepositoryIds));
        }

        public Builder selectedRepositoryIds(Integer... selectedRepositoryIds) {
            return selectedRepositoryIds(List.of(selectedRepositoryIds));
        }

        public Builder visibility(Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public DependabotOrganizationSecretArgs build() {
            $.secretName = Objects.requireNonNull($.secretName, "expected parameter 'secretName' to be non-null");
            $.visibility = Objects.requireNonNull($.visibility, "expected parameter 'visibility' to be non-null");
            return $;
        }
    }

}
