---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_stepfunctions_state_machine_alias Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for StateMachineAlias
---

# awscc_stepfunctions_state_machine_alias (Resource)

Resource schema for StateMachineAlias



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `deployment_preference` (Attributes) The settings to enable gradual state machine deployments. (see [below for nested schema](#nestedatt--deployment_preference))
- `description` (String) An optional description of the alias.
- `name` (String) The alias name.
- `routing_configuration` (Attributes Set) The routing configuration of the alias. One or two versions can be mapped to an alias to split StartExecution requests of the same state machine. (see [below for nested schema](#nestedatt--routing_configuration))

### Read-Only

- `arn` (String) The ARN of the alias.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--deployment_preference"></a>
### Nested Schema for `deployment_preference`

Required:

- `state_machine_version_arn` (String)
- `type` (String) The type of deployment to perform.

Optional:

- `alarms` (Set of String) A list of CloudWatch alarm names that will be monitored during the deployment. The deployment will fail and rollback if any alarms go into ALARM state.
- `interval` (Number) The time in minutes between each traffic shifting increment.
- `percentage` (Number) The percentage of traffic to shift to the new version in each increment.


<a id="nestedatt--routing_configuration"></a>
### Nested Schema for `routing_configuration`

Required:

- `state_machine_version_arn` (String) The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.
- `weight` (Number) The percentage of traffic you want to route to the state machine version. The sum of the weights in the routing configuration must be equal to 100.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_stepfunctions_state_machine_alias.example "arn"
```
