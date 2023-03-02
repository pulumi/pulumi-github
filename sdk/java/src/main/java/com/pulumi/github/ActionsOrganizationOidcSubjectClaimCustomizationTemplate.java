// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.ActionsOrganizationOidcSubjectClaimCustomizationTemplateState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage an OpenID Connect subject claim customization template within a GitHub
 * organization.
 * 
 * More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
 * available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.ActionsOrganizationOidcSubjectClaimCustomizationTemplate;
 * import com.pulumi.github.ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs;
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
 *         var exampleTemplate = new ActionsOrganizationOidcSubjectClaimCustomizationTemplate(&#34;exampleTemplate&#34;, ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs.builder()        
 *             .includeClaimKeys(            
 *                 &#34;actor&#34;,
 *                 &#34;context&#34;,
 *                 &#34;repository_owner&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported using the organization&#39;s name.
 * 
 * ```sh
 *  $ pulumi import github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate test example_organization
 * ```
 * 
 */
@ResourceType(type="github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate")
public class ActionsOrganizationOidcSubjectClaimCustomizationTemplate extends com.pulumi.resources.CustomResource {
    /**
     * A list of OpenID Connect claims.
     * 
     */
    @Export(name="includeClaimKeys", type=List.class, parameters={String.class})
    private Output<List<String>> includeClaimKeys;

    /**
     * @return A list of OpenID Connect claims.
     * 
     */
    public Output<List<String>> includeClaimKeys() {
        return this.includeClaimKeys;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ActionsOrganizationOidcSubjectClaimCustomizationTemplate(String name) {
        this(name, ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ActionsOrganizationOidcSubjectClaimCustomizationTemplate(String name, ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ActionsOrganizationOidcSubjectClaimCustomizationTemplate(String name, ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate", name, args == null ? ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ActionsOrganizationOidcSubjectClaimCustomizationTemplate(String name, Output<String> id, @Nullable ActionsOrganizationOidcSubjectClaimCustomizationTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static ActionsOrganizationOidcSubjectClaimCustomizationTemplate get(String name, Output<String> id, @Nullable ActionsOrganizationOidcSubjectClaimCustomizationTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ActionsOrganizationOidcSubjectClaimCustomizationTemplate(name, id, state, options);
    }
}