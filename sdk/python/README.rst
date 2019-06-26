|Build Status|

GitHub provider
===============

The GitHub resource provider for Pulumi lets you use GitHub resources in
your cloud programs. To use this package, please `install the Pulumi CLI
first <https://pulumi.io/>`__.

Installing
----------

This package is available in many languages in the standard packaging
formats.

Node.js (Java/TypeScript)
~~~~~~~~~~~~~~~~~~~~~~~~~

To use from JavaScript or TypeScript in Node.js, install using either
``npm``:

::

   $ npm install @pulumi/github

or ``yarn``:

::

   $ yarn add @pulumi/github

Python
~~~~~~

To use from Python, install using ``pip``:

::

   $ pip install pulumi_github

Go
~~

To use from Go, use ``go get`` to grab the latest version of the library

::

   $ go get github.com/pulumi/pulumi-github/sdk/go/...

Configuration
-------------

The following configuration points are available:

-  token (Required) - GitHub personal access token. This can also be set
   via the ``GITHUB_TOKEN`` environment variable.
-  organization (Required) - This is the target GitHub organization to
   manage. The account corresponding to the token will need "owner"
   privileges for this organization. This can also be set via the
   ``GITHUB_ORGANIZATION`` environment variable.
-  baseUrl (Optional) - This is the target GitHub base API endpoint.
   Providing a value is a requirement when working with GitHub
   Enterprise. It can be sourced from ``GITHUB_BASE_URL`` environment
   variable and must end with ``\``.
-  insecure (Optional) - Whether server should be accessed without
   verifying the TLS certificate. As the name suggests this is insecure
   and should not be used beyond experiments, accessing local
   (non-production) GHE instance etc. The default is ``false``.

Reference
---------

For detailed reference documentation, please visit `the API
docs <https://pulumi.io/reference/pkg/nodejs/@pulumi/github/index.html>`__.

Updating this provider
----------------------

The AWS Resource Provider for Pulumi is based on the Terraform Provider
for AWS. Instructions for keeping it up to date are available
`here <https://github.com/pulumi/pulumi-terraform/wiki/Updating-Pulumi-Providers-Backed-By-Terraform-Providers>`__.

.. |Build Status| image:: https://travis-ci.com/pulumi/pulumi-github.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master
   :target: https://travis-ci.com/pulumi/pulumi-github
