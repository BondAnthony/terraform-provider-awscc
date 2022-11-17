---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_organizations_organizational_unit Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  You can use organizational units (OUs) to group accounts together to administer as a single unit. This greatly simplifies the management of your accounts. For example, you can attach a policy-based control to an OU, and all accounts within the OU automatically inherit the policy. You can create multiple OUs within a single organization, and you can create OUs within other OUs. Each OU can contain multiple accounts, and you can move accounts from one OU to another. However, OU names must be unique within a parent OU or root.
---

# awscc_organizations_organizational_unit (Resource)

You can use organizational units (OUs) to group accounts together to administer as a single unit. This greatly simplifies the management of your accounts. For example, you can attach a policy-based control to an OU, and all accounts within the OU automatically inherit the policy. You can create multiple OUs within a single organization, and you can create OUs within other OUs. Each OU can contain multiple accounts, and you can move accounts from one OU to another. However, OU names must be unique within a parent OU or root.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The friendly name of this OU.
- `parent_id` (String) The unique identifier (ID) of the parent root or OU that you want to create the new OU in.

### Optional

- `tags` (Attributes Set) A list of tags that you want to attach to the newly created OU. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of this OU.
- `id` (String) The unique identifier (ID) associated with this OU.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key identifier, or name, of the tag.
- `value` (String) The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_organizations_organizational_unit.example <resource ID>
```