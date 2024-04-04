// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package securitylake

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_securitylake_subscriber", subscriberResource)
}

// subscriberResource returns the Terraform awscc_securitylake_subscriber resource.
// This Terraform resource corresponds to the CloudFormation AWS::SecurityLake::Subscriber resource.
func subscriberResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessTypes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon S3 or AWS Lake Formation access type.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "enum": [
		//	      "LAKEFORMATION",
		//	      "S3"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"access_types": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon S3 or AWS Lake Formation access type.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
				listvalidator.UniqueValues(),
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"LAKEFORMATION",
						"S3",
					),
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: DataLakeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN for the data lake.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"data_lake_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN for the data lake.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceShareArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"resource_share_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceShareName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"resource_share_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3BucketArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"s3_bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Sources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The supported AWS services from which logs and events are collected.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "properties": {
		//	      "AwsLogSource": {
		//	        "additionalProperties": false,
		//	        "description": "Amazon Security Lake supports log and event collection for natively supported AWS services.",
		//	        "properties": {
		//	          "SourceName": {
		//	            "description": "The name for a AWS source. This must be a Regionally unique value.",
		//	            "type": "string"
		//	          },
		//	          "SourceVersion": {
		//	            "description": "The version for a AWS source. This must be a Regionally unique value.",
		//	            "pattern": "^(latest|[0-9]\\.[0-9])$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "CustomLogSource": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "SourceName": {
		//	            "description": "The name for a third-party custom source. This must be a Regionally unique value.",
		//	            "maxLength": 64,
		//	            "minLength": 1,
		//	            "pattern": "^[\\\\\\w\\-_:/.]*$",
		//	            "type": "string"
		//	          },
		//	          "SourceVersion": {
		//	            "description": "The version for a third-party custom source. This must be a Regionally unique value.",
		//	            "maxLength": 32,
		//	            "minLength": 1,
		//	            "pattern": "^[A-Za-z0-9\\-\\.\\_]*$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"sources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AwsLogSource
					"aws_log_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SourceName
							"source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name for a AWS source. This must be a Regionally unique value.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: SourceVersion
							"source_version": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The version for a AWS source. This must be a Regionally unique value.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^(latest|[0-9]\\.[0-9])$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Amazon Security Lake supports log and event collection for natively supported AWS services.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CustomLogSource
					"custom_log_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SourceName
							"source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name for a third-party custom source. This must be a Regionally unique value.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 64),
									stringvalidator.RegexMatches(regexp.MustCompile("^[\\\\\\w\\-_:/.]*$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: SourceVersion
							"source_version": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The version for a third-party custom source. This must be a Regionally unique value.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 32),
									stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9\\-\\.\\_]*$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The supported AWS services from which logs and events are collected.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubscriberArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"subscriber_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriberDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description for your subscriber account in Security Lake.",
		//	  "type": "string"
		//	}
		"subscriber_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description for your subscriber account in Security Lake.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriberIdentity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The AWS identity used to access your data.",
		//	  "properties": {
		//	    "ExternalId": {
		//	      "description": "The external ID used to establish trust relationship with the AWS identity.",
		//	      "maxLength": 1224,
		//	      "minLength": 2,
		//	      "pattern": "^[\\w+=,.@:/-]*$",
		//	      "type": "string"
		//	    },
		//	    "Principal": {
		//	      "description": "The AWS identity principal.",
		//	      "pattern": "^([0-9]{12}|[a-z0-9\\.\\-]*\\.(amazonaws|amazon)\\.com)$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ExternalId",
		//	    "Principal"
		//	  ],
		//	  "type": "object"
		//	}
		"subscriber_identity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExternalId
				"external_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The external ID used to establish trust relationship with the AWS identity.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(2, 1224),
						stringvalidator.RegexMatches(regexp.MustCompile("^[\\w+=,.@:/-]*$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Principal
				"principal": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS identity principal.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^([0-9]{12}|[a-z0-9\\.\\-]*\\.(amazonaws|amazon)\\.com)$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The AWS identity used to access your data.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubscriberName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of your Security Lake subscriber account.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[\\\\\\w\\s\\-_:/,.@=+]*$",
		//	  "type": "string"
		//	}
		"subscriber_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of your Security Lake subscriber account.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\\\\\w\\s\\-_:/,.@=+]*$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriberRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"subscriber_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of objects, one for each tag to associate with the subscriber. For each tag, you must specify both a tag key and a tag value. A tag value cannot be null, but it can be an empty string.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The name of the tag. This is a general label that acts as a category for a more specific tag value (value).",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value that is associated with the specified tag key (key). This value acts as a descriptor for the tag key. A tag value cannot be null, but it can be an empty string.",
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
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the tag. This is a general label that acts as a category for a more specific tag value (value).",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value that is associated with the specified tag key (key). This value acts as a descriptor for the tag key. A tag value cannot be null, but it can be an empty string.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of objects, one for each tag to associate with the subscriber. For each tag, you must specify both a tag key and a tag value. A tag value cannot be null, but it can be an empty string.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SecurityLake::Subscriber",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecurityLake::Subscriber").WithTerraformTypeName("awscc_securitylake_subscriber")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_types":           "AccessTypes",
		"aws_log_source":         "AwsLogSource",
		"custom_log_source":      "CustomLogSource",
		"data_lake_arn":          "DataLakeArn",
		"external_id":            "ExternalId",
		"key":                    "Key",
		"principal":              "Principal",
		"resource_share_arn":     "ResourceShareArn",
		"resource_share_name":    "ResourceShareName",
		"s3_bucket_arn":          "S3BucketArn",
		"source_name":            "SourceName",
		"source_version":         "SourceVersion",
		"sources":                "Sources",
		"subscriber_arn":         "SubscriberArn",
		"subscriber_description": "SubscriberDescription",
		"subscriber_identity":    "SubscriberIdentity",
		"subscriber_name":        "SubscriberName",
		"subscriber_role_arn":    "SubscriberRoleArn",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
