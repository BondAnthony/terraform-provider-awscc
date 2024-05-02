---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ecr_repository_creation_template Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ECR::RepositoryCreationTemplate
---

# awscc_ecr_repository_creation_template (Data Source)

Data Source schema for AWS::ECR::RepositoryCreationTemplate



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `applied_for` (Set of String) A list of enumerable Strings representing the repository creation scenarios that the template will apply towards.
- `created_at` (String) Create timestamp of the template.
- `description` (String) The description of the template.
- `encryption_configuration` (Attributes) The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest. By default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.

For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html (see [below for nested schema](#nestedatt--encryption_configuration))
- `image_tag_mutability` (String) The image tag mutability setting for the repository.
- `lifecycle_policy` (String) The JSON lifecycle policy text to apply to the repository. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
- `prefix` (String) The prefix use to match the repository name and apply the template.
- `repository_policy` (String) The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html
- `resource_tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--resource_tags))
- `updated_at` (String) Update timestamp of the template.

<a id="nestedatt--encryption_configuration"></a>
### Nested Schema for `encryption_configuration`

Read-Only:

- `encryption_type` (String) The encryption type to use.
- `kms_key` (String) If you use the KMS encryption type, specify the CMK to use for encryption. The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.


<a id="nestedatt--resource_tags"></a>
### Nested Schema for `resource_tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.