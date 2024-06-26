---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotwireless_multicast_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::IoTWireless::MulticastGroup
---

# awscc_iotwireless_multicast_group (Data Source)

Data Source schema for AWS::IoTWireless::MulticastGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) Multicast group arn. Returned after successful create.
- `associate_wireless_device` (String) Wireless device to associate. Only for update request.
- `description` (String) Multicast group description
- `disassociate_wireless_device` (String) Wireless device to disassociate. Only for update request.
- `lo_ra_wan` (Attributes) Multicast group LoRaWAN (see [below for nested schema](#nestedatt--lo_ra_wan))
- `multicast_group_id` (String) Multicast group id. Returned after successful create.
- `name` (String) Name of Multicast group
- `status` (String) Multicast group status. Returned after successful read.
- `tags` (Attributes Set) A list of key-value pairs that contain metadata for the Multicast group. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--lo_ra_wan"></a>
### Nested Schema for `lo_ra_wan`

Read-Only:

- `dl_class` (String) Multicast group LoRaWAN DL Class
- `number_of_devices_in_group` (Number) Multicast group number of devices in group. Returned after successful read.
- `number_of_devices_requested` (Number) Multicast group number of devices requested. Returned after successful read.
- `rf_region` (String) Multicast group LoRaWAN RF region


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
