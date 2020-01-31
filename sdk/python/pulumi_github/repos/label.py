# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Label(pulumi.CustomResource):
    color: pulumi.Output[str]
    """
    A 6 character hex code, **without the leading #**, identifying the color of the label.
    """
    description: pulumi.Output[str]
    """
    A short description of the label.
    """
    etag: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    The name of the label.
    """
    repository: pulumi.Output[str]
    """
    The GitHub repository
    """
    url: pulumi.Output[str]
    """
    The URL to the issue label
    """
    def __init__(__self__, resource_name, opts=None, color=None, description=None, name=None, repository=None, __name__=None, __opts__=None):
        """
        Provides a GitHub issue label resource.
        
        This resource allows you to create and manage issue labels within your
        GitHub organization.
        
        Issue labels are keyed off of their "name", so pre-existing issue labels result
        in a 422 HTTP error if they exist outside of Terraform. Normally this would not
        be an issue, except new repositories are created with a "default" set of labels,
        and those labels easily conflict with custom ones.
        
        This resource will first check if the label exists, and then issue an update,
        otherwise it will create.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: A 6 character hex code, **without the leading #**, identifying the color of the label.
        :param pulumi.Input[str] description: A short description of the label.
        :param pulumi.Input[str] name: The name of the label.
        :param pulumi.Input[str] repository: The GitHub repository
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if color is None:
            raise TypeError("Missing required property 'color'")
        __props__['color'] = color

        __props__['description'] = description

        __props__['name'] = name

        if repository is None:
            raise TypeError("Missing required property 'repository'")
        __props__['repository'] = repository

        __props__['etag'] = None
        __props__['url'] = None

        super(Label, __self__).__init__(
            'github:repos/label:Label',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
