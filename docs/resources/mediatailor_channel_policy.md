---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_mediatailor_channel_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::MediaTailor::ChannelPolicy Resource Type
---

# awscc_mediatailor_channel_policy (Resource)

Definition of AWS::MediaTailor::ChannelPolicy Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `channel_name` (String)
- `policy` (String) <p>The IAM policy for the channel. IAM policies are used to control access to your channel.</p>

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_mediatailor_channel_policy.example "channel_name"
```
