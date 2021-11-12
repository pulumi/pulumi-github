# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

import types

__config__ = pulumi.Config('github')


class _ExportableConfig(types.ModuleType):
    @property
    def app_auth(self) -> Optional[str]:
        """
        The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
        and `app_auth` are not set.
        """
        return __config__.get('appAuth')

    @property
    def base_url(self) -> str:
        """
        The GitHub Base API URL
        """
        return __config__.get('baseUrl') or (_utilities.get_env('GITHUB_BASE_URL') or 'https://api.github.com/')

    @property
    def insecure(self) -> Optional[bool]:
        """
        Enable `insecure` mode for testing purposes
        """
        return __config__.get_bool('insecure')

    @property
    def organization(self) -> Optional[str]:
        """
        The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
        """
        return __config__.get('organization')

    @property
    def owner(self) -> Optional[str]:
        """
        The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
        """
        return __config__.get('owner')

    @property
    def token(self) -> Optional[str]:
        """
        The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `app_auth` are not set.
        """
        return __config__.get('token')

    @property
    def write_delay_ms(self) -> Optional[int]:
        """
        Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.
        """
        return __config__.get_int('writeDelayMs')

