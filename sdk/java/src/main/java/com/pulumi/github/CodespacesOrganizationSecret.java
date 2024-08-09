// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.CodespacesOrganizationSecretArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.CodespacesOrganizationSecretState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ## Import
 * 
 * This resource can be imported using an ID made up of the secret name:
 * 
 * ```sh
 * $ pulumi import github:index/codespacesOrganizationSecret:CodespacesOrganizationSecret test_secret test_secret_name
 * ```
 * 
 * NOTE: the implementation is limited in that it won&#39;t fetch the value of the
 * 
 * `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
 * 
 */
@ResourceType(type="github:index/codespacesOrganizationSecret:CodespacesOrganizationSecret")
public class CodespacesOrganizationSecret extends com.pulumi.resources.CustomResource {
    /**
     * Date of codespaces_secret creation.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Date of codespaces_secret creation.
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
     * An array of repository ids that can access the organization secret.
     * 
     */
    @Export(name="selectedRepositoryIds", refs={List.class,Integer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Integer>> selectedRepositoryIds;

    /**
     * @return An array of repository ids that can access the organization secret.
     * 
     */
    public Output<Optional<List<Integer>>> selectedRepositoryIds() {
        return Codegen.optional(this.selectedRepositoryIds);
    }
    /**
     * Date of codespaces_secret update.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Date of codespaces_secret update.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CodespacesOrganizationSecret(java.lang.String name) {
        this(name, CodespacesOrganizationSecretArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CodespacesOrganizationSecret(java.lang.String name, CodespacesOrganizationSecretArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CodespacesOrganizationSecret(java.lang.String name, CodespacesOrganizationSecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/codespacesOrganizationSecret:CodespacesOrganizationSecret", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CodespacesOrganizationSecret(java.lang.String name, Output<java.lang.String> id, @Nullable CodespacesOrganizationSecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/codespacesOrganizationSecret:CodespacesOrganizationSecret", name, state, makeResourceOptions(options, id), false);
    }

    private static CodespacesOrganizationSecretArgs makeArgs(CodespacesOrganizationSecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CodespacesOrganizationSecretArgs.Empty : args;
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
    public static CodespacesOrganizationSecret get(java.lang.String name, Output<java.lang.String> id, @Nullable CodespacesOrganizationSecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CodespacesOrganizationSecret(name, id, state, options);
    }
}
