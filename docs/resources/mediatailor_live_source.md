---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_mediatailor_live_source Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::MediaTailor::LiveSource Resource Type
---

# awscc_mediatailor_live_source (Resource)

Definition of AWS::MediaTailor::LiveSource Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `http_package_configurations` (Attributes List) <p>A list of HTTP package configuration parameters for this live source.</p> (see [below for nested schema](#nestedatt--http_package_configurations))
- `live_source_name` (String)
- `source_location_name` (String)

### Optional

- `tags` (Attributes Set) The tags to assign to the live source. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) <p>The ARN of the live source.</p>
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--http_package_configurations"></a>
### Nested Schema for `http_package_configurations`

Required:

- `path` (String) <p>The relative path to the URL for this VOD source. This is combined with <code>SourceLocation::HttpConfiguration::BaseUrl</code> to form a valid URL.</p>
- `source_group` (String) <p>The name of the source group. This has to match one of the <code>Channel::Outputs::SourceGroup</code>.</p>
- `type` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_mediatailor_live_source.example "live_source_name|source_location_name"
```
