---
page_title: "awscc_datazone_environment_blueprint_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type
---

# awscc_datazone_environment_blueprint_configuration (Resource)

Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type

## Example Usage

```terraform
resource "awscc_datazone_environment_blueprint_configuration" "example" {
  domain_identifier                = awscc_datazone_domain.example.domain_id
  enabled_regions                  = ["us-east-1"]
  environment_blueprint_identifier = "DefaultDataLake"
  manage_access_role_arn           = awscc_iam_role.access.arn
  provisioning_role_arn            = awscc_iam_role.provisioning.arn
  regional_parameters = [{
    parameters = {
      "S3Location" : "s3://replace_with_bucket_name"
    }
    region = "us-east-1" # region is required for the datalake environment
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_identifier` (String)
- `enabled_regions` (List of String)
- `environment_blueprint_identifier` (String)

### Optional

- `manage_access_role_arn` (String)
- `provisioning_role_arn` (String)
- `regional_parameters` (Attributes Set) (see [below for nested schema](#nestedatt--regional_parameters))

### Read-Only

- `created_at` (String)
- `domain_id` (String)
- `environment_blueprint_id` (String)
- `id` (String) Uniquely identifies the resource.
- `updated_at` (String)

<a id="nestedatt--regional_parameters"></a>
### Nested Schema for `regional_parameters`

Optional:

- `parameters` (Map of String)
- `region` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_datazone_environment_blueprint_configuration.example "domain_id|environment_blueprint_id"
```