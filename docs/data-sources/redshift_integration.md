---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_redshift_integration Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Redshift::Integration
---

# awscc_redshift_integration (Data Source)

Data Source schema for AWS::Redshift::Integration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `additional_encryption_context` (Map of String) An optional set of non-secret key?value pairs that contains additional contextual information about the data.
- `create_time` (String) The time (UTC) when the integration was created.
- `integration_arn` (String) The Amazon Resource Name (ARN) of the integration.
- `integration_name` (String) The name of the integration.
- `kms_key_id` (String) An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used.
- `source_arn` (String) The Amazon Resource Name (ARN) of the database to use as the source for replication, for example, arn:aws:dynamodb:us-east-2:123412341234:table/dynamotable
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `target_arn` (String) The Amazon Resource Name (ARN) of the Redshift data warehouse to use as the target for replication, for example, arn:aws:redshift:us-east-2:123412341234:namespace:e43aab3e-10a3-4ec4-83d4-f227ff9bfbcf

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.