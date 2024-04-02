// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_connect_instance_storage_config", instanceStorageConfigResource)
}

// instanceStorageConfigResource returns the Terraform awscc_connect_instance_storage_config resource.
// This Terraform resource corresponds to the CloudFormation AWS::Connect::InstanceStorageConfig resource.
func instanceStorageConfigResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An associationID is automatically generated when a storage config is associated with an instance",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[-a-z0-9]*$",
		//	  "type": "string"
		//	}
		"association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An associationID is automatically generated when a storage config is associated with an instance",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Connect Instance ID with which the storage config will be associated",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Connect Instance ID with which the storage config will be associated",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KinesisFirehoseConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "FirehoseArn": {
		//	      "description": "An ARN is a unique AWS resource identifier.",
		//	      "pattern": "^arn:aws[-a-z0-9]*:firehose:[-a-z0-9]*:[0-9]{12}:deliverystream/[-a-zA-Z0-9_.]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "FirehoseArn"
		//	  ],
		//	  "type": "object"
		//	}
		"kinesis_firehose_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FirehoseArn
				"firehose_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "An ARN is a unique AWS resource identifier.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[-a-z0-9]*:firehose:[-a-z0-9]*:[0-9]{12}:deliverystream/[-a-zA-Z0-9_.]*$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KinesisStreamConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "StreamArn": {
		//	      "description": "An ARN is a unique AWS resource identifier.",
		//	      "pattern": "^arn:aws[-a-z0-9]*:kinesis:[-a-z0-9]*:[0-9]{12}:stream/[-a-zA-Z0-9_.]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "StreamArn"
		//	  ],
		//	  "type": "object"
		//	}
		"kinesis_stream_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: StreamArn
				"stream_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "An ARN is a unique AWS resource identifier.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[-a-z0-9]*:kinesis:[-a-z0-9]*:[0-9]{12}:stream/[-a-zA-Z0-9_.]*$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KinesisVideoStreamConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EncryptionConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EncryptionType": {
		//	          "description": "Specifies default encryption using AWS KMS-Managed Keys",
		//	          "enum": [
		//	            "KMS"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "KeyId": {
		//	          "description": "Specifies the encryption key id",
		//	          "maxLength": 128,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "EncryptionType",
		//	        "KeyId"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Prefix": {
		//	      "description": "Prefixes are used to infer logical hierarchy",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "RetentionPeriodHours": {
		//	      "description": "Number of hours",
		//	      "type": "number"
		//	    }
		//	  },
		//	  "required": [
		//	    "Prefix",
		//	    "RetentionPeriodHours",
		//	    "EncryptionConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"kinesis_video_stream_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EncryptionConfig
				"encryption_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EncryptionType
						"encryption_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies default encryption using AWS KMS-Managed Keys",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"KMS",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: KeyId
						"key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies the encryption key id",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 128),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: Prefix
				"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Prefixes are used to infer logical hierarchy",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 128),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: RetentionPeriodHours
				"retention_period_hours": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "Number of hours",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the type of storage resource available for the instance",
		//	  "enum": [
		//	    "CHAT_TRANSCRIPTS",
		//	    "CALL_RECORDINGS",
		//	    "SCHEDULED_REPORTS",
		//	    "MEDIA_STREAMS",
		//	    "CONTACT_TRACE_RECORDS",
		//	    "AGENT_EVENTS"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the type of storage resource available for the instance",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CHAT_TRANSCRIPTS",
					"CALL_RECORDINGS",
					"SCHEDULED_REPORTS",
					"MEDIA_STREAMS",
					"CONTACT_TRACE_RECORDS",
					"AGENT_EVENTS",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3Config
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BucketName": {
		//	      "description": "A name for the S3 Bucket",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "BucketPrefix": {
		//	      "description": "Prefixes are used to infer logical hierarchy",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "EncryptionConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EncryptionType": {
		//	          "description": "Specifies default encryption using AWS KMS-Managed Keys",
		//	          "enum": [
		//	            "KMS"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "KeyId": {
		//	          "description": "Specifies the encryption key id",
		//	          "maxLength": 128,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "EncryptionType",
		//	        "KeyId"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "BucketName",
		//	    "BucketPrefix"
		//	  ],
		//	  "type": "object"
		//	}
		"s3_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BucketName
				"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A name for the S3 Bucket",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 128),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: BucketPrefix
				"bucket_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Prefixes are used to infer logical hierarchy",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 128),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: EncryptionConfig
				"encryption_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EncryptionType
						"encryption_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies default encryption using AWS KMS-Managed Keys",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"KMS",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: KeyId
						"key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies the encryption key id",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 128),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StorageType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the storage type to be associated with the instance",
		//	  "enum": [
		//	    "S3",
		//	    "KINESIS_VIDEO_STREAM",
		//	    "KINESIS_STREAM",
		//	    "KINESIS_FIREHOSE"
		//	  ],
		//	  "type": "string"
		//	}
		"storage_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the storage type to be associated with the instance",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"S3",
					"KINESIS_VIDEO_STREAM",
					"KINESIS_STREAM",
					"KINESIS_FIREHOSE",
				),
			}, /*END VALIDATORS*/
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
		Description: "Resource Type definition for AWS::Connect::InstanceStorageConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::InstanceStorageConfig").WithTerraformTypeName("awscc_connect_instance_storage_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id":              "AssociationId",
		"bucket_name":                 "BucketName",
		"bucket_prefix":               "BucketPrefix",
		"encryption_config":           "EncryptionConfig",
		"encryption_type":             "EncryptionType",
		"firehose_arn":                "FirehoseArn",
		"instance_arn":                "InstanceArn",
		"key_id":                      "KeyId",
		"kinesis_firehose_config":     "KinesisFirehoseConfig",
		"kinesis_stream_config":       "KinesisStreamConfig",
		"kinesis_video_stream_config": "KinesisVideoStreamConfig",
		"prefix":                      "Prefix",
		"resource_type":               "ResourceType",
		"retention_period_hours":      "RetentionPeriodHours",
		"s3_config":                   "S3Config",
		"storage_type":                "StorageType",
		"stream_arn":                  "StreamArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
