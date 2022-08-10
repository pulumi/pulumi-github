// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserGpgKeyState extends com.pulumi.resources.ResourceArgs {

    public static final UserGpgKeyState Empty = new UserGpgKeyState();

    /**
     * Your public GPG key, generated in ASCII-armored format.
     * See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
     * 
     */
    @Import(name="armoredPublicKey")
    private @Nullable Output<String> armoredPublicKey;

    /**
     * @return Your public GPG key, generated in ASCII-armored format.
     * See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
     * 
     */
    public Optional<Output<String>> armoredPublicKey() {
        return Optional.ofNullable(this.armoredPublicKey);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The key ID of the GPG key, e.g. `3262EFF25BA0D270`
     * 
     */
    @Import(name="keyId")
    private @Nullable Output<String> keyId;

    /**
     * @return The key ID of the GPG key, e.g. `3262EFF25BA0D270`
     * 
     */
    public Optional<Output<String>> keyId() {
        return Optional.ofNullable(this.keyId);
    }

    private UserGpgKeyState() {}

    private UserGpgKeyState(UserGpgKeyState $) {
        this.armoredPublicKey = $.armoredPublicKey;
        this.etag = $.etag;
        this.keyId = $.keyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserGpgKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserGpgKeyState $;

        public Builder() {
            $ = new UserGpgKeyState();
        }

        public Builder(UserGpgKeyState defaults) {
            $ = new UserGpgKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param armoredPublicKey Your public GPG key, generated in ASCII-armored format.
         * See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
         * 
         * @return builder
         * 
         */
        public Builder armoredPublicKey(@Nullable Output<String> armoredPublicKey) {
            $.armoredPublicKey = armoredPublicKey;
            return this;
        }

        /**
         * @param armoredPublicKey Your public GPG key, generated in ASCII-armored format.
         * See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
         * 
         * @return builder
         * 
         */
        public Builder armoredPublicKey(String armoredPublicKey) {
            return armoredPublicKey(Output.of(armoredPublicKey));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param keyId The key ID of the GPG key, e.g. `3262EFF25BA0D270`
         * 
         * @return builder
         * 
         */
        public Builder keyId(@Nullable Output<String> keyId) {
            $.keyId = keyId;
            return this;
        }

        /**
         * @param keyId The key ID of the GPG key, e.g. `3262EFF25BA0D270`
         * 
         * @return builder
         * 
         */
        public Builder keyId(String keyId) {
            return keyId(Output.of(keyId));
        }

        public UserGpgKeyState build() {
            return $;
        }
    }

}