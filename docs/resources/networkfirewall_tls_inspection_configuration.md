---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_networkfirewall_tls_inspection_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration
---

# awscc_networkfirewall_tls_inspection_configuration (Resource)

Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `tls_inspection_configuration` (Attributes) (see [below for nested schema](#nestedatt--tls_inspection_configuration))
- `tls_inspection_configuration_name` (String)

### Optional

- `description` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `tls_inspection_configuration_arn` (String) A resource ARN.
- `tls_inspection_configuration_id` (String)

<a id="nestedatt--tls_inspection_configuration"></a>
### Nested Schema for `tls_inspection_configuration`

Optional:

- `server_certificate_configurations` (Attributes List) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations))

<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations`

Optional:

- `certificate_authority_arn` (String) A resource ARN.
- `check_certificate_revocation_status` (Attributes) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations--check_certificate_revocation_status))
- `scopes` (Attributes List) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes))
- `server_certificates` (Attributes Set) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations--server_certificates))

<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations--check_certificate_revocation_status"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations.check_certificate_revocation_status`

Optional:

- `revoked_status_action` (String)
- `unknown_status_action` (String)


<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations.scopes`

Optional:

- `destination_ports` (Attributes List) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--destination_ports))
- `destinations` (Attributes List) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--destinations))
- `protocols` (List of Number)
- `source_ports` (Attributes List) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--source_ports))
- `sources` (Attributes List) (see [below for nested schema](#nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--sources))

<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--destination_ports"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations.scopes.destination_ports`

Required:

- `from_port` (Number)
- `to_port` (Number)


<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--destinations"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations.scopes.destinations`

Required:

- `address_definition` (String)


<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--source_ports"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations.scopes.source_ports`

Required:

- `from_port` (Number)
- `to_port` (Number)


<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations--scopes--sources"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations.scopes.sources`

Required:

- `address_definition` (String)



<a id="nestedatt--tls_inspection_configuration--server_certificate_configurations--server_certificates"></a>
### Nested Schema for `tls_inspection_configuration.server_certificate_configurations.server_certificates`

Optional:

- `resource_arn` (String) A resource ARN.




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_networkfirewall_tls_inspection_configuration.example "tls_inspection_configuration_arn"
```
