# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, avatar_url=None, bio=None, blog=None, company=None, created_at=None, email=None, followers=None, following=None, gpg_keys=None, gravatar_id=None, id=None, location=None, login=None, name=None, node_id=None, public_gists=None, public_repos=None, site_admin=None, ssh_keys=None, suspended_at=None, updated_at=None, username=None):
        if avatar_url and not isinstance(avatar_url, str):
            raise TypeError("Expected argument 'avatar_url' to be a str")
        pulumi.set(__self__, "avatar_url", avatar_url)
        if bio and not isinstance(bio, str):
            raise TypeError("Expected argument 'bio' to be a str")
        pulumi.set(__self__, "bio", bio)
        if blog and not isinstance(blog, str):
            raise TypeError("Expected argument 'blog' to be a str")
        pulumi.set(__self__, "blog", blog)
        if company and not isinstance(company, str):
            raise TypeError("Expected argument 'company' to be a str")
        pulumi.set(__self__, "company", company)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if followers and not isinstance(followers, int):
            raise TypeError("Expected argument 'followers' to be a int")
        pulumi.set(__self__, "followers", followers)
        if following and not isinstance(following, int):
            raise TypeError("Expected argument 'following' to be a int")
        pulumi.set(__self__, "following", following)
        if gpg_keys and not isinstance(gpg_keys, list):
            raise TypeError("Expected argument 'gpg_keys' to be a list")
        pulumi.set(__self__, "gpg_keys", gpg_keys)
        if gravatar_id and not isinstance(gravatar_id, str):
            raise TypeError("Expected argument 'gravatar_id' to be a str")
        pulumi.set(__self__, "gravatar_id", gravatar_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if login and not isinstance(login, str):
            raise TypeError("Expected argument 'login' to be a str")
        pulumi.set(__self__, "login", login)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_id and not isinstance(node_id, str):
            raise TypeError("Expected argument 'node_id' to be a str")
        pulumi.set(__self__, "node_id", node_id)
        if public_gists and not isinstance(public_gists, int):
            raise TypeError("Expected argument 'public_gists' to be a int")
        pulumi.set(__self__, "public_gists", public_gists)
        if public_repos and not isinstance(public_repos, int):
            raise TypeError("Expected argument 'public_repos' to be a int")
        pulumi.set(__self__, "public_repos", public_repos)
        if site_admin and not isinstance(site_admin, bool):
            raise TypeError("Expected argument 'site_admin' to be a bool")
        pulumi.set(__self__, "site_admin", site_admin)
        if ssh_keys and not isinstance(ssh_keys, list):
            raise TypeError("Expected argument 'ssh_keys' to be a list")
        pulumi.set(__self__, "ssh_keys", ssh_keys)
        if suspended_at and not isinstance(suspended_at, str):
            raise TypeError("Expected argument 'suspended_at' to be a str")
        pulumi.set(__self__, "suspended_at", suspended_at)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="avatarUrl")
    def avatar_url(self) -> str:
        return pulumi.get(self, "avatar_url")

    @property
    @pulumi.getter
    def bio(self) -> str:
        return pulumi.get(self, "bio")

    @property
    @pulumi.getter
    def blog(self) -> str:
        return pulumi.get(self, "blog")

    @property
    @pulumi.getter
    def company(self) -> str:
        return pulumi.get(self, "company")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def followers(self) -> int:
        return pulumi.get(self, "followers")

    @property
    @pulumi.getter
    def following(self) -> int:
        return pulumi.get(self, "following")

    @property
    @pulumi.getter(name="gpgKeys")
    def gpg_keys(self) -> Sequence[str]:
        return pulumi.get(self, "gpg_keys")

    @property
    @pulumi.getter(name="gravatarId")
    def gravatar_id(self) -> str:
        return pulumi.get(self, "gravatar_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def login(self) -> str:
        return pulumi.get(self, "login")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> str:
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter(name="publicGists")
    def public_gists(self) -> int:
        return pulumi.get(self, "public_gists")

    @property
    @pulumi.getter(name="publicRepos")
    def public_repos(self) -> int:
        return pulumi.get(self, "public_repos")

    @property
    @pulumi.getter(name="siteAdmin")
    def site_admin(self) -> bool:
        return pulumi.get(self, "site_admin")

    @property
    @pulumi.getter(name="sshKeys")
    def ssh_keys(self) -> Sequence[str]:
        return pulumi.get(self, "ssh_keys")

    @property
    @pulumi.getter(name="suspendedAt")
    def suspended_at(self) -> str:
        return pulumi.get(self, "suspended_at")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            avatar_url=self.avatar_url,
            bio=self.bio,
            blog=self.blog,
            company=self.company,
            created_at=self.created_at,
            email=self.email,
            followers=self.followers,
            following=self.following,
            gpg_keys=self.gpg_keys,
            gravatar_id=self.gravatar_id,
            id=self.id,
            location=self.location,
            login=self.login,
            name=self.name,
            node_id=self.node_id,
            public_gists=self.public_gists,
            public_repos=self.public_repos,
            site_admin=self.site_admin,
            ssh_keys=self.ssh_keys,
            suspended_at=self.suspended_at,
            updated_at=self.updated_at,
            username=self.username)


def get_user(username: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        avatar_url=__ret__.avatar_url,
        bio=__ret__.bio,
        blog=__ret__.blog,
        company=__ret__.company,
        created_at=__ret__.created_at,
        email=__ret__.email,
        followers=__ret__.followers,
        following=__ret__.following,
        gpg_keys=__ret__.gpg_keys,
        gravatar_id=__ret__.gravatar_id,
        id=__ret__.id,
        location=__ret__.location,
        login=__ret__.login,
        name=__ret__.name,
        node_id=__ret__.node_id,
        public_gists=__ret__.public_gists,
        public_repos=__ret__.public_repos,
        site_admin=__ret__.site_admin,
        ssh_keys=__ret__.ssh_keys,
        suspended_at=__ret__.suspended_at,
        updated_at=__ret__.updated_at,
        username=__ret__.username)


@_utilities.lift_output_func(get_user)
def get_user_output(username: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
