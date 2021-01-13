# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['UserGpgKey']


class UserGpgKey(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 armored_public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a GitHub user's GPG key resource.

        This resource allows you to add/remove GPG keys from your user account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example = github.UserGpgKey("example", armored_public_key=\"\"\"-----BEGIN PGP PUBLIC KEY BLOCK-----
        ...
        -----END PGP PUBLIC KEY BLOCK-----
        \"\"\")
        ```

        ## Import

        GPG keys are not importable due to the fact that [API](https://developer.github.com/v3/users/gpg_keys/#gpg-keys) does not return previously uploaded GPG key.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] armored_public_key: Your public GPG key, generated in ASCII-armored format.
               See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if armored_public_key is None and not opts.urn:
                raise TypeError("Missing required property 'armored_public_key'")
            __props__['armored_public_key'] = armored_public_key
            __props__['etag'] = None
            __props__['key_id'] = None
        super(UserGpgKey, __self__).__init__(
            'github:index/userGpgKey:UserGpgKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            armored_public_key: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None) -> 'UserGpgKey':
        """
        Get an existing UserGpgKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] armored_public_key: Your public GPG key, generated in ASCII-armored format.
               See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
        :param pulumi.Input[str] key_id: The key ID of the GPG key, e.g. `3262EFF25BA0D270`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["armored_public_key"] = armored_public_key
        __props__["etag"] = etag
        __props__["key_id"] = key_id
        return UserGpgKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="armoredPublicKey")
    def armored_public_key(self) -> pulumi.Output[str]:
        """
        Your public GPG key, generated in ASCII-armored format.
        See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
        """
        return pulumi.get(self, "armored_public_key")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        The key ID of the GPG key, e.g. `3262EFF25BA0D270`
        """
        return pulumi.get(self, "key_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

