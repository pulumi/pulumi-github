# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'GetRestApiResult',
    'AwaitableGetRestApiResult',
    'get_rest_api',
    'get_rest_api_output',
]

@pulumi.output_type
class GetRestApiResult:
    """
    A collection of values returned by getRestApi.
    """
    def __init__(__self__, body=None, code=None, endpoint=None, headers=None, id=None, status=None):
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if code and not isinstance(code, int):
            raise TypeError("Expected argument 'code' to be a int")
        pulumi.set(__self__, "code", code)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if headers and not isinstance(headers, str):
            raise TypeError("Expected argument 'headers' to be a str")
        pulumi.set(__self__, "headers", headers)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def body(self) -> str:
        """
        A JSON string containing response body.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def code(self) -> int:
        """
        A response status code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def headers(self) -> str:
        """
        A JSON string containing response headers.
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        A response status string.
        """
        return pulumi.get(self, "status")


class AwaitableGetRestApiResult(GetRestApiResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRestApiResult(
            body=self.body,
            code=self.code,
            endpoint=self.endpoint,
            headers=self.headers,
            id=self.id,
            status=self.status)


def get_rest_api(endpoint: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRestApiResult:
    """
    Use this data source to retrieve information about a GitHub resource through REST API.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_rest_api(endpoint="repos/example_repo/git/refs/heads/main")
    ```


    :param str endpoint: REST API endpoint to send the GET request to.
    """
    __args__ = dict()
    __args__['endpoint'] = endpoint
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRestApi:getRestApi', __args__, opts=opts, typ=GetRestApiResult).value

    return AwaitableGetRestApiResult(
        body=pulumi.get(__ret__, 'body'),
        code=pulumi.get(__ret__, 'code'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        headers=pulumi.get(__ret__, 'headers'),
        id=pulumi.get(__ret__, 'id'),
        status=pulumi.get(__ret__, 'status'))
def get_rest_api_output(endpoint: Optional[pulumi.Input[str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRestApiResult]:
    """
    Use this data source to retrieve information about a GitHub resource through REST API.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_rest_api(endpoint="repos/example_repo/git/refs/heads/main")
    ```


    :param str endpoint: REST API endpoint to send the GET request to.
    """
    __args__ = dict()
    __args__['endpoint'] = endpoint
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getRestApi:getRestApi', __args__, opts=opts, typ=GetRestApiResult)
    return __ret__.apply(lambda __response__: GetRestApiResult(
        body=pulumi.get(__response__, 'body'),
        code=pulumi.get(__response__, 'code'),
        endpoint=pulumi.get(__response__, 'endpoint'),
        headers=pulumi.get(__response__, 'headers'),
        id=pulumi.get(__response__, 'id'),
        status=pulumi.get(__response__, 'status')))
