# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ActionsEnvironmentSecretArgs', 'ActionsEnvironmentSecret']

@pulumi.input_type
class ActionsEnvironmentSecretArgs:
    def __init__(__self__, *,
                 environment: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 secret_name: pulumi.Input[str],
                 encrypted_value: Optional[pulumi.Input[str]] = None,
                 plaintext_value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ActionsEnvironmentSecret resource.
        """
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "repository", repository)
        pulumi.set(__self__, "secret_name", secret_name)
        if encrypted_value is not None:
            pulumi.set(__self__, "encrypted_value", encrypted_value)
        if plaintext_value is not None:
            pulumi.set(__self__, "plaintext_value", plaintext_value)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter(name="encryptedValue")
    def encrypted_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "encrypted_value")

    @encrypted_value.setter
    def encrypted_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encrypted_value", value)

    @property
    @pulumi.getter(name="plaintextValue")
    def plaintext_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "plaintext_value")

    @plaintext_value.setter
    def plaintext_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plaintext_value", value)


@pulumi.input_type
class _ActionsEnvironmentSecretState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 encrypted_value: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 plaintext_value: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ActionsEnvironmentSecret resources.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if encrypted_value is not None:
            pulumi.set(__self__, "encrypted_value", encrypted_value)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if plaintext_value is not None:
            pulumi.set(__self__, "plaintext_value", plaintext_value)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if secret_name is not None:
            pulumi.set(__self__, "secret_name", secret_name)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="encryptedValue")
    def encrypted_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "encrypted_value")

    @encrypted_value.setter
    def encrypted_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encrypted_value", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="plaintextValue")
    def plaintext_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "plaintext_value")

    @plaintext_value.setter
    def plaintext_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plaintext_value", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class ActionsEnvironmentSecret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encrypted_value: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 plaintext_value: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ActionsEnvironmentSecret resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionsEnvironmentSecretArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ActionsEnvironmentSecret resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ActionsEnvironmentSecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionsEnvironmentSecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encrypted_value: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 plaintext_value: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActionsEnvironmentSecretArgs.__new__(ActionsEnvironmentSecretArgs)

            __props__.__dict__["encrypted_value"] = None if encrypted_value is None else pulumi.Output.secret(encrypted_value)
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            __props__.__dict__["plaintext_value"] = None if plaintext_value is None else pulumi.Output.secret(plaintext_value)
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            if secret_name is None and not opts.urn:
                raise TypeError("Missing required property 'secret_name'")
            __props__.__dict__["secret_name"] = secret_name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["encryptedValue", "plaintextValue"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ActionsEnvironmentSecret, __self__).__init__(
            'github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            encrypted_value: Optional[pulumi.Input[str]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            plaintext_value: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            secret_name: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'ActionsEnvironmentSecret':
        """
        Get an existing ActionsEnvironmentSecret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionsEnvironmentSecretState.__new__(_ActionsEnvironmentSecretState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["encrypted_value"] = encrypted_value
        __props__.__dict__["environment"] = environment
        __props__.__dict__["plaintext_value"] = plaintext_value
        __props__.__dict__["repository"] = repository
        __props__.__dict__["secret_name"] = secret_name
        __props__.__dict__["updated_at"] = updated_at
        return ActionsEnvironmentSecret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="encryptedValue")
    def encrypted_value(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "encrypted_value")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="plaintextValue")
    def plaintext_value(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "plaintext_value")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

