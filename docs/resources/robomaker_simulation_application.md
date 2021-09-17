---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_robomaker_simulation_application Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_robomaker_simulation_application (Resource)

An example resource schema demonstrating some basic constructs and validation rules.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **robot_software_suite** (Attributes) Information about a robot software suite (ROS distribution). (see [below for nested schema](#nestedatt--robot_software_suite))
- **simulation_software_suite** (Attributes) Information about a simulation software suite. (see [below for nested schema](#nestedatt--simulation_software_suite))

### Optional

- **current_revision_id** (String) The current revision id.
- **environment** (String) The URI of the Docker image for the robot application.
- **name** (String) The name of the simulation application.
- **rendering_engine** (Attributes) Information about a rendering engine. (see [below for nested schema](#nestedatt--rendering_engine))
- **sources** (Attributes List) The sources of the simulation application. (see [below for nested schema](#nestedatt--sources))
- **tags** (Map of String) A key-value pair to associate with a resource.

### Read-Only

- **arn** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--robot_software_suite"></a>
### Nested Schema for `robot_software_suite`

Required:

- **name** (String) The name of the robot software suite (ROS distribution).
- **version** (String) The version of the robot software suite (ROS distribution).


<a id="nestedatt--simulation_software_suite"></a>
### Nested Schema for `simulation_software_suite`

Required:

- **name** (String) The name of the simulation software suite.
- **version** (String) The version of the simulation software suite.


<a id="nestedatt--rendering_engine"></a>
### Nested Schema for `rendering_engine`

Optional:

- **name** (String) The name of the rendering engine.
- **version** (String) The version of the rendering engine.


<a id="nestedatt--sources"></a>
### Nested Schema for `sources`

Optional:

- **architecture** (String) The target processor architecture for the application.
- **s3_bucket** (String) The Amazon S3 bucket name.
- **s3_key** (String) The s3 object key.

