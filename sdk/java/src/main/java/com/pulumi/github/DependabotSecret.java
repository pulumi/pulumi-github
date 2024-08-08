// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.DependabotSecretArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.DependabotSecretState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ## Import
 * 
 * This resource can be imported using an ID made up of the `repository` and `secret_name`:
 * 
 * ```sh
 * $ pulumi import github:index/dependabotSecret:DependabotSecret example_secret example_repository/example_secret
 * ```
 * NOTE: the implementation is limited in that it won&#39;t fetch the value of the
 * `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
 * 
 */
@ResourceType(type="github:index/dependabotSecret:DependabotSecret")
public class DependabotSecret extends com.pulumi.resources.CustomResource {
    /**
     * Date of dependabot_secret creation.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Date of dependabot_secret creation.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Encrypted value of the secret using the GitHub public key in Base64 format.
     * 
     */
    @Export(name="encryptedValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encryptedValue;

    /**
     * @return Encrypted value of the secret using the GitHub public key in Base64 format.
     * 
     */
    public Output<Optional<String>> encryptedValue() {
        return Codegen.optional(this.encryptedValue);
    }
    /**
     * Plaintext value of the secret to be encrypted
     * 
     */
    @Export(name="plaintextValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> plaintextValue;

    /**
     * @return Plaintext value of the secret to be encrypted
     * 
     */
    public Output<Optional<String>> plaintextValue() {
        return Codegen.optional(this.plaintextValue);
    }
    /**
     * Name of the repository
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return Name of the repository
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * Name of the secret
     * 
     */
    @Export(name="secretName", refs={String.class}, tree="[0]")
    private Output<String> secretName;

    /**
     * @return Name of the secret
     * 
     */
    public Output<String> secretName() {
        return this.secretName;
    }
    /**
     * Date of dependabot_secret update.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Date of dependabot_secret update.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DependabotSecret(java.lang.String name) {
        this(name, DependabotSecretArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DependabotSecret(java.lang.String name, DependabotSecretArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DependabotSecret(java.lang.String name, DependabotSecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/dependabotSecret:DependabotSecret", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DependabotSecret(java.lang.String name, Output<java.lang.String> id, @Nullable DependabotSecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/dependabotSecret:DependabotSecret", name, state, makeResourceOptions(options, id), false);
    }

    private static DependabotSecretArgs makeArgs(DependabotSecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DependabotSecretArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "encryptedValue",
                "plaintextValue"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DependabotSecret get(java.lang.String name, Output<java.lang.String> id, @Nullable DependabotSecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DependabotSecret(name, id, state, options);
    }
}
