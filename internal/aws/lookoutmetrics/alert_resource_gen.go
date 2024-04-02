// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lookoutmetrics_alert", alertResource)
}

// alertResource returns the Terraform awscc_lookoutmetrics_alert resource.
// This Terraform resource corresponds to the CloudFormation AWS::LookoutMetrics::Alert resource.
func alertResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The action to be taken by the alert when an anomaly is detected.",
		//	  "properties": {
		//	    "LambdaConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Configuration options for a Lambda alert action.",
		//	      "properties": {
		//	        "LambdaArn": {
		//	          "description": "ARN of a Lambda to send alert notifications to.",
		//	          "maxLength": 256,
		//	          "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
		//	          "type": "string"
		//	        },
		//	        "RoleArn": {
		//	          "description": "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
		//	          "maxLength": 256,
		//	          "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "RoleArn",
		//	        "LambdaArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "SNSConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Configuration options for an SNS alert action.",
		//	      "properties": {
		//	        "RoleArn": {
		//	          "description": "ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.",
		//	          "maxLength": 256,
		//	          "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
		//	          "type": "string"
		//	        },
		//	        "SnsTopicArn": {
		//	          "description": "ARN of an SNS topic to send alert notifications to.",
		//	          "maxLength": 256,
		//	          "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "RoleArn",
		//	        "SnsTopicArn"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"action": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LambdaConfiguration
				"lambda_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LambdaArn
						"lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "ARN of a Lambda to send alert notifications to.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(256),
								stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: RoleArn
						"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(256),
								stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Configuration options for a Lambda alert action.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SNSConfiguration
				"sns_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: RoleArn
						"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(256),
								stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: SnsTopicArn
						"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "ARN of an SNS topic to send alert notifications to.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(256),
								stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Configuration options for an SNS alert action.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The action to be taken by the alert when an anomaly is detected.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AlertDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the alert.",
		//	  "maxLength": 256,
		//	  "pattern": ".*\\S.*",
		//	  "type": "string"
		//	}
		"alert_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the alert.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile(".*\\S.*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AlertName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the alert. If not provided, a name is generated automatically.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	  "type": "string"
		//	}
		"alert_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the alert. If not provided, a name is generated automatically.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AlertSensitivityThreshold
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
		//	  "maximum": 100,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"alert_sensitivity_threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
			Required:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(0, 100),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AnomalyDetectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
		//	  "type": "string"
		//	}
		"anomaly_detector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):.*:.*:.*:.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN assigned to the alert.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):.*:.*:.*:.+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN assigned to the alert.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::LookoutMetrics::Alert",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutMetrics::Alert").WithTerraformTypeName("awscc_lookoutmetrics_alert")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                      "Action",
		"alert_description":           "AlertDescription",
		"alert_name":                  "AlertName",
		"alert_sensitivity_threshold": "AlertSensitivityThreshold",
		"anomaly_detector_arn":        "AnomalyDetectorArn",
		"arn":                         "Arn",
		"lambda_arn":                  "LambdaArn",
		"lambda_configuration":        "LambdaConfiguration",
		"role_arn":                    "RoleArn",
		"sns_configuration":           "SNSConfiguration",
		"sns_topic_arn":               "SnsTopicArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
