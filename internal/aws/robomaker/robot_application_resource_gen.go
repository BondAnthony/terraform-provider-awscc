// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_robomaker_robot_application", robotApplicationResource)
}

// robotApplicationResource returns the Terraform awscc_robomaker_robot_application resource.
// This Terraform resource corresponds to the CloudFormation AWS::RoboMaker::RobotApplication resource.
func robotApplicationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CurrentRevisionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The revision ID of robot application.",
		//	  "maxLength": 40,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"current_revision_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The revision ID of robot application.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 40),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Environment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URI of the Docker image for the robot application.",
		//	  "type": "string"
		//	}
		"environment": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URI of the Docker image for the robot application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the robot application.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the robot application.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RobotSoftwareSuite
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The robot software suite used by the robot application.",
		//	  "properties": {
		//	    "Name": {
		//	      "description": "The name of robot software suite.",
		//	      "enum": [
		//	        "ROS",
		//	        "ROS2",
		//	        "General"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of robot software suite.",
		//	      "enum": [
		//	        "Kinetic",
		//	        "Melodic",
		//	        "Dashing"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name"
		//	  ],
		//	  "type": "object"
		//	}
		"robot_software_suite": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of robot software suite.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ROS",
							"ROS2",
							"General",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of robot software suite.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"Kinetic",
							"Melodic",
							"Dashing",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
					// Version is a write-only property.
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The robot software suite used by the robot application.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Sources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The sources of the robot application.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Architecture": {
		//	        "description": "The architecture of robot application.",
		//	        "enum": [
		//	          "X86_64",
		//	          "ARM64",
		//	          "ARMHF"
		//	        ],
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "S3Bucket": {
		//	        "description": "The Arn of the S3Bucket that stores the robot application source.",
		//	        "type": "string"
		//	      },
		//	      "S3Key": {
		//	        "description": "The s3 key of robot application source.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "S3Bucket",
		//	      "S3Key",
		//	      "Architecture"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"sources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Architecture
					"architecture": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The architecture of robot application.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 255),
							stringvalidator.OneOf(
								"X86_64",
								"ARM64",
								"ARMHF",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: S3Bucket
					"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Arn of the S3Bucket that stores the robot application source.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: S3Key
					"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The s3 key of robot application source.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The sources of the robot application.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Sources is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair to associate with a resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "This schema is for testing purpose only.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::RobotApplication").WithTerraformTypeName("awscc_robomaker_robot_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"architecture":         "Architecture",
		"arn":                  "Arn",
		"current_revision_id":  "CurrentRevisionId",
		"environment":          "Environment",
		"name":                 "Name",
		"robot_software_suite": "RobotSoftwareSuite",
		"s3_bucket":            "S3Bucket",
		"s3_key":               "S3Key",
		"sources":              "Sources",
		"tags":                 "Tags",
		"version":              "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/RobotSoftwareSuite/Version",
		"/properties/Sources",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
