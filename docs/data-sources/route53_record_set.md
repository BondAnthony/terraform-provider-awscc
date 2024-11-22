---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53_record_set Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Route53::RecordSet
---

# awscc_route53_record_set (Data Source)

Data Source schema for AWS::Route53::RecordSet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `alias_target` (Attributes) (see [below for nested schema](#nestedatt--alias_target))
- `cidr_routing_config` (Attributes) (see [below for nested schema](#nestedatt--cidr_routing_config))
- `comment` (String)
- `failover` (String)
- `geo_location` (Attributes) (see [below for nested schema](#nestedatt--geo_location))
- `geo_proximity_location` (Attributes) (see [below for nested schema](#nestedatt--geo_proximity_location))
- `health_check_id` (String)
- `hosted_zone_id` (String)
- `hosted_zone_name` (String)
- `multi_value_answer` (Boolean)
- `name` (String)
- `record_set_id` (String)
- `region` (String)
- `resource_records` (List of String)
- `set_identifier` (String)
- `ttl` (String)
- `type` (String)
- `weight` (Number)

<a id="nestedatt--alias_target"></a>
### Nested Schema for `alias_target`

Read-Only:

- `dns_name` (String)
- `evaluate_target_health` (Boolean)
- `hosted_zone_id` (String)


<a id="nestedatt--cidr_routing_config"></a>
### Nested Schema for `cidr_routing_config`

Read-Only:

- `collection_id` (String)
- `location_name` (String)


<a id="nestedatt--geo_location"></a>
### Nested Schema for `geo_location`

Read-Only:

- `continent_code` (String)
- `country_code` (String)
- `subdivision_code` (String)


<a id="nestedatt--geo_proximity_location"></a>
### Nested Schema for `geo_proximity_location`

Read-Only:

- `aws_region` (String)
- `bias` (Number)
- `coordinates` (Attributes) (see [below for nested schema](#nestedatt--geo_proximity_location--coordinates))
- `local_zone_group` (String)

<a id="nestedatt--geo_proximity_location--coordinates"></a>
### Nested Schema for `geo_proximity_location.coordinates`

Read-Only:

- `latitude` (String)
- `longitude` (String)