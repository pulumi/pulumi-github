# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AppAuth',
]

@pulumi.output_type
class AppAuth(dict):
    def __init__(__self__, *,
                 id: str,
                 installation_id: str,
                 pem_file: str):
        """
        :param str id: The GitHub App ID.
        :param str installation_id: The GitHub App installation instance ID.
        :param str pem_file: The GitHub App PEM file contents.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "installation_id", installation_id)
        pulumi.set(__self__, "pem_file", pem_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The GitHub App ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="installationId")
    def installation_id(self) -> str:
        """
        The GitHub App installation instance ID.
        """
        return pulumi.get(self, "installation_id")

    @property
    @pulumi.getter(name="pemFile")
    def pem_file(self) -> str:
        """
        The GitHub App PEM file contents.
        """
        return pulumi.get(self, "pem_file")


