# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['OrganizationWebhookArgs', 'OrganizationWebhook']

@pulumi.input_type
class OrganizationWebhookArgs:
    def __init__(__self__, *,
                 events: pulumi.Input[Sequence[pulumi.Input[str]]],
                 active: Optional[pulumi.Input[bool]] = None,
                 configuration: Optional[pulumi.Input['OrganizationWebhookConfigurationArgs']] = None):
        """
        The set of arguments for constructing a OrganizationWebhook resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events: A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
        :param pulumi.Input[bool] active: Indicate of the webhook should receive events. Defaults to `true`.
        :param pulumi.Input['OrganizationWebhookConfigurationArgs'] configuration: key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
        """
        pulumi.set(__self__, "events", events)
        if active is not None:
            pulumi.set(__self__, "active", active)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)

    @property
    @pulumi.getter
    def events(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
        """
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "events", value)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate of the webhook should receive events. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['OrganizationWebhookConfigurationArgs']]:
        """
        key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['OrganizationWebhookConfigurationArgs']]):
        pulumi.set(self, "configuration", value)


@pulumi.input_type
class _OrganizationWebhookState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 configuration: Optional[pulumi.Input['OrganizationWebhookConfigurationArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationWebhook resources.
        :param pulumi.Input[bool] active: Indicate of the webhook should receive events. Defaults to `true`.
        :param pulumi.Input['OrganizationWebhookConfigurationArgs'] configuration: key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events: A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
        :param pulumi.Input[str] url: URL of the webhook
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if events is not None:
            pulumi.set(__self__, "events", events)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate of the webhook should receive events. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def configuration(self) -> Optional[pulumi.Input['OrganizationWebhookConfigurationArgs']]:
        """
        key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
        """
        return pulumi.get(self, "configuration")

    @configuration.setter
    def configuration(self, value: Optional[pulumi.Input['OrganizationWebhookConfigurationArgs']]):
        pulumi.set(self, "configuration", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def events(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
        """
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "events", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the webhook
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class OrganizationWebhook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['OrganizationWebhookConfigurationArgs']]] = None,
                 events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage webhooks for GitHub organization.

        ## Import

        Organization webhooks can be imported using the `id` of the webhook.
        The `id` of the webhook can be found in the URL of the webhook. For example, `"https://github.com/organizations/foo-org/settings/hooks/123456789"`.

        ```sh
        $ pulumi import github:index/organizationWebhook:OrganizationWebhook terraform 123456789
        ```
        If secret is populated in the webhook's configuration, the value will be imported as "********".

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Indicate of the webhook should receive events. Defaults to `true`.
        :param pulumi.Input[pulumi.InputType['OrganizationWebhookConfigurationArgs']] configuration: key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events: A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationWebhookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage webhooks for GitHub organization.

        ## Import

        Organization webhooks can be imported using the `id` of the webhook.
        The `id` of the webhook can be found in the URL of the webhook. For example, `"https://github.com/organizations/foo-org/settings/hooks/123456789"`.

        ```sh
        $ pulumi import github:index/organizationWebhook:OrganizationWebhook terraform 123456789
        ```
        If secret is populated in the webhook's configuration, the value will be imported as "********".

        :param str resource_name: The name of the resource.
        :param OrganizationWebhookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationWebhookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['OrganizationWebhookConfigurationArgs']]] = None,
                 events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationWebhookArgs.__new__(OrganizationWebhookArgs)

            __props__.__dict__["active"] = active
            __props__.__dict__["configuration"] = configuration
            if events is None and not opts.urn:
                raise TypeError("Missing required property 'events'")
            __props__.__dict__["events"] = events
            __props__.__dict__["etag"] = None
            __props__.__dict__["url"] = None
        super(OrganizationWebhook, __self__).__init__(
            'github:index/organizationWebhook:OrganizationWebhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            configuration: Optional[pulumi.Input[pulumi.InputType['OrganizationWebhookConfigurationArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            events: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'OrganizationWebhook':
        """
        Get an existing OrganizationWebhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Indicate of the webhook should receive events. Defaults to `true`.
        :param pulumi.Input[pulumi.InputType['OrganizationWebhookConfigurationArgs']] configuration: key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events: A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
        :param pulumi.Input[str] url: URL of the webhook
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationWebhookState.__new__(_OrganizationWebhookState)

        __props__.__dict__["active"] = active
        __props__.__dict__["configuration"] = configuration
        __props__.__dict__["etag"] = etag
        __props__.__dict__["events"] = events
        __props__.__dict__["url"] = url
        return OrganizationWebhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate of the webhook should receive events. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Optional['outputs.OrganizationWebhookConfiguration']]:
        """
        key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def events(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the webhook
        """
        return pulumi.get(self, "url")

