---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ivs_public_key Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IVS::PublicKey
---

# awscc_ivs_public_key (Resource)

Resource Type definition for AWS::IVS::PublicKey



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name` (String) Name of the public key to be imported. The value does not need to be unique.
- `public_key_material` (String) The public portion of a customer-generated key pair.
- `tags` (Attributes Set) A list of key-value pairs that contain metadata for the asset model. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) Key-pair identifier.
- `fingerprint` (String) Key-pair identifier.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ivs_public_key.example "arn"
```
