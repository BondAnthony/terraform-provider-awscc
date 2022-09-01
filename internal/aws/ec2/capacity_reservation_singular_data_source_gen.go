// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_capacity_reservation", capacityReservationDataSourceType)
}

// capacityReservationDataSourceType returns the Terraform awscc_ec2_capacity_reservation data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::CapacityReservation resource type.
func capacityReservationDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"available_instance_count": {
			// Property: AvailableInstanceCount
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"ebs_optimized": {
			// Property: EbsOptimized
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"end_date": {
			// Property: EndDate
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"end_date_type": {
			// Property: EndDateType
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"ephemeral_storage": {
			// Property: EphemeralStorage
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"instance_count": {
			// Property: InstanceCount
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"instance_match_criteria": {
			// Property: InstanceMatchCriteria
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"instance_platform": {
			// Property: InstancePlatform
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"instance_type": {
			// Property: InstanceType
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"out_post_arn": {
			// Property: OutPostArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"placement_group_arn": {
			// Property: PlacementGroupArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tag_specifications": {
			// Property: TagSpecifications
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ResourceType": {
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "type": "string"
			//             },
			//             "Value": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Value",
			//             "Key"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"resource_type": {
						// Property: ResourceType
						Type:     types.StringType,
						Computed: true,
					},
					"tags": {
						// Property: Tags
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:     types.StringType,
									Computed: true,
								},
								"value": {
									// Property: Value
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tenancy": {
			// Property: Tenancy
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"total_instance_count": {
			// Property: TotalInstanceCount
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::CapacityReservation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::CapacityReservation").WithTerraformTypeName("awscc_ec2_capacity_reservation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"availability_zone":        "AvailabilityZone",
		"available_instance_count": "AvailableInstanceCount",
		"ebs_optimized":            "EbsOptimized",
		"end_date":                 "EndDate",
		"end_date_type":            "EndDateType",
		"ephemeral_storage":        "EphemeralStorage",
		"id":                       "Id",
		"instance_count":           "InstanceCount",
		"instance_match_criteria":  "InstanceMatchCriteria",
		"instance_platform":        "InstancePlatform",
		"instance_type":            "InstanceType",
		"key":                      "Key",
		"out_post_arn":             "OutPostArn",
		"placement_group_arn":      "PlacementGroupArn",
		"resource_type":            "ResourceType",
		"tag_specifications":       "TagSpecifications",
		"tags":                     "Tags",
		"tenancy":                  "Tenancy",
		"total_instance_count":     "TotalInstanceCount",
		"value":                    "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}