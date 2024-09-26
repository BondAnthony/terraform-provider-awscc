---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datazone_environment_actions Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::DataZone::EnvironmentActions Resource Type
---

# awscc_datazone_environment_actions (Resource)

Definition of AWS::DataZone::EnvironmentActions Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the environment action.

### Optional

- `description` (String) The description of the Amazon DataZone environment action.
- `domain_identifier` (String) The identifier of the Amazon DataZone domain in which the environment would be created.
- `environment_identifier` (String) The identifier of the Amazon DataZone environment in which the action is taking place
- `identifier` (String) The ID of the Amazon DataZone environment action.
- `parameters` (Attributes) The parameters of the environment action. (see [below for nested schema](#nestedatt--parameters))

### Read-Only

- `domain_id` (String) The identifier of the Amazon DataZone domain in which the environment is created.
- `environment_actions_id` (String) The ID of the Amazon DataZone environment action.
- `environment_id` (String) The identifier of the Amazon DataZone environment in which the action is taking place
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `uri` (String) The URI of the console link specified as part of the environment action.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_datazone_environment_actions.example "domain_id|environment_id|id"
```