---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_transfer_connector Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Transfer::Connector
---

# awscc_transfer_connector (Resource)

Resource Type definition for AWS::Transfer::Connector



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `access_role` (String) Specifies the access role for the connector.
- `url` (String) URL for Connector

### Optional

- `as_2_config` (Attributes) Configuration for an AS2 connector. (see [below for nested schema](#nestedatt--as_2_config))
- `logging_role` (String) Specifies the logging role for the connector.
- `security_policy_name` (String) Security policy for SFTP Connector
- `sftp_config` (Attributes) Configuration for an SFTP connector. (see [below for nested schema](#nestedatt--sftp_config))
- `tags` (Attributes Set) Key-value pairs that can be used to group and search for connectors. Tags are metadata attached to connectors for any purpose. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) Specifies the unique Amazon Resource Name (ARN) for the connector.
- `connector_id` (String) A unique identifier for the connector.
- `id` (String) Uniquely identifies the resource.
- `service_managed_egress_ip_addresses` (List of String) The list of egress IP addresses of this connector. These IP addresses are assigned automatically when you create the connector.

<a id="nestedatt--as_2_config"></a>
### Nested Schema for `as_2_config`

Optional:

- `basic_auth_secret_id` (String) ARN or name of the secret in AWS Secrets Manager which contains the credentials for Basic authentication. If empty, Basic authentication is disabled for the AS2 connector
- `compression` (String) Compression setting for this AS2 connector configuration.
- `encryption_algorithm` (String) Encryption algorithm for this AS2 connector configuration.
- `local_profile_id` (String) A unique identifier for the local profile.
- `mdn_response` (String) MDN Response setting for this AS2 connector configuration.
- `mdn_signing_algorithm` (String) MDN Signing algorithm for this AS2 connector configuration.
- `message_subject` (String) The message subject for this AS2 connector configuration.
- `partner_profile_id` (String) A unique identifier for the partner profile.
- `signing_algorithm` (String) Signing algorithm for this AS2 connector configuration.


<a id="nestedatt--sftp_config"></a>
### Nested Schema for `sftp_config`

Optional:

- `trusted_host_keys` (List of String) List of public host keys, for the external server to which you are connecting.
- `user_secret_id` (String) ARN or name of the secret in AWS Secrets Manager which contains the SFTP user's private keys or passwords.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The name assigned to the tag that you create.
- `value` (String) Contains one or more values that you assigned to the key name you create.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_transfer_connector.example "connector_id"
```
