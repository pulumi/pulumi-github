// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.EmuGroupMappingArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.EmuGroupMappingState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * This resource manages mappings between external groups for enterprise managed users and GitHub teams. It wraps the API detailed [here](https://docs.github.com/en/rest/reference/teams#external-groups). Note that this is a distinct resource from `github.TeamSyncGroupMapping`. `github.EmuGroupMapping` is special to the Enterprise Managed User (EMU) external group feature, whereas `github.TeamSyncGroupMapping` is specific to Identity Provider Groups.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.EmuGroupMapping;
 * import com.pulumi.github.EmuGroupMappingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleEmuGroupMapping = new EmuGroupMapping("exampleEmuGroupMapping", EmuGroupMappingArgs.builder()
 *             .teamSlug("emu-test-team")
 *             .groupId(28836)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitHub EMU External Group Mappings can be imported using the external `group_id`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/emuGroupMapping:EmuGroupMapping example_emu_group_mapping 28836
 * ```
 * 
 */
@ResourceType(type="github:index/emuGroupMapping:EmuGroupMapping")
public class EmuGroupMapping extends com.pulumi.resources.CustomResource {
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Integer corresponding to the external group ID to be linked
     * 
     */
    @Export(name="groupId", refs={Integer.class}, tree="[0]")
    private Output<Integer> groupId;

    /**
     * @return Integer corresponding to the external group ID to be linked
     * 
     */
    public Output<Integer> groupId() {
        return this.groupId;
    }
    /**
     * Slug of the GitHub team
     * 
     */
    @Export(name="teamSlug", refs={String.class}, tree="[0]")
    private Output<String> teamSlug;

    /**
     * @return Slug of the GitHub team
     * 
     */
    public Output<String> teamSlug() {
        return this.teamSlug;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EmuGroupMapping(java.lang.String name) {
        this(name, EmuGroupMappingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EmuGroupMapping(java.lang.String name, EmuGroupMappingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EmuGroupMapping(java.lang.String name, EmuGroupMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/emuGroupMapping:EmuGroupMapping", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EmuGroupMapping(java.lang.String name, Output<java.lang.String> id, @Nullable EmuGroupMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/emuGroupMapping:EmuGroupMapping", name, state, makeResourceOptions(options, id), false);
    }

    private static EmuGroupMappingArgs makeArgs(EmuGroupMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EmuGroupMappingArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static EmuGroupMapping get(java.lang.String name, Output<java.lang.String> id, @Nullable EmuGroupMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EmuGroupMapping(name, id, state, options);
    }
}
