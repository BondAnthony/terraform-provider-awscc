// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package athena

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_athena_capacity_reservation", capacityReservationDataSource)
}

// capacityReservationDataSource returns the Terraform awscc_athena_capacity_reservation data source.
// This Terraform data source corresponds to the CloudFormation AWS::Athena::CapacityReservation resource.
func capacityReservationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllocatedDpus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of DPUs Athena has provisioned and allocated for the reservation",
		//	  "format": "int64",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"allocated_dpus": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of DPUs Athena has provisioned and allocated for the reservation",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified capacity reservation",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified capacity reservation",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CapacityAssignmentConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Assignment configuration to assign workgroups to a reservation",
		//	  "properties": {
		//	    "CapacityAssignments": {
		//	      "description": "List of capacity assignments",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "WorkgroupNames": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "pattern": "[a-zA-Z0-9._-]{1,128}",
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "WorkgroupNames"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "CapacityAssignments"
		//	  ],
		//	  "type": "object"
		//	}
		"capacity_assignment_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CapacityAssignments
				"capacity_assignments": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: WorkgroupNames
							"workgroup_names": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of capacity assignments",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Assignment configuration to assign workgroups to a reservation",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time the reservation was created.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time the reservation was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastSuccessfulAllocationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp when the last successful allocated was made",
		//	  "type": "string"
		//	}
		"last_successful_allocation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp when the last successful allocated was made",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The reservation name.",
		//	  "pattern": "[a-zA-Z0-9._-]{1,128}",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The reservation name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the reservation.",
		//	  "enum": [
		//	    "PENDING",
		//	    "ACTIVE",
		//	    "CANCELLING",
		//	    "CANCELLED",
		//	    "FAILED",
		//	    "UPDATE_PENDING"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the reservation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetDpus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of DPUs to request to be allocated to the reservation.",
		//	  "format": "int64",
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"target_dpus": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of DPUs to request to be allocated to the reservation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Athena::CapacityReservation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::CapacityReservation").WithTerraformTypeName("awscc_athena_capacity_reservation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocated_dpus":                    "AllocatedDpus",
		"arn":                               "Arn",
		"capacity_assignment_configuration": "CapacityAssignmentConfiguration",
		"capacity_assignments":              "CapacityAssignments",
		"creation_time":                     "CreationTime",
		"key":                               "Key",
		"last_successful_allocation_time":   "LastSuccessfulAllocationTime",
		"name":                              "Name",
		"status":                            "Status",
		"tags":                              "Tags",
		"target_dpus":                       "TargetDpus",
		"value":                             "Value",
		"workgroup_names":                   "WorkgroupNames",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
