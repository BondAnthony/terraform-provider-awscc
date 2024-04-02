// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package devopsguru

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
	registry.AddResourceFactory("awscc_devopsguru_notification_channel", notificationChannelResource)
}

// notificationChannelResource returns the Terraform awscc_devopsguru_notification_channel resource.
// This Terraform resource corresponds to the CloudFormation AWS::DevOpsGuru::NotificationChannel resource.
func notificationChannelResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Config
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Information about notification channels you have configured with DevOps Guru.",
		//	  "properties": {
		//	    "Filters": {
		//	      "additionalProperties": false,
		//	      "description": "Information about filters of a notification channel configured in DevOpsGuru to filter for insights.",
		//	      "properties": {
		//	        "MessageTypes": {
		//	          "description": "DevOps Guru message types to filter for",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "DevOps Guru NotificationMessageType Enum",
		//	            "enum": [
		//	              "NEW_INSIGHT",
		//	              "CLOSED_INSIGHT",
		//	              "NEW_ASSOCIATION",
		//	              "SEVERITY_UPGRADED",
		//	              "NEW_RECOMMENDATION"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 5,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        },
		//	        "Severities": {
		//	          "description": "DevOps Guru insight severities to filter for",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "DevOps Guru Insight Severity Enum",
		//	            "enum": [
		//	              "LOW",
		//	              "MEDIUM",
		//	              "HIGH"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 3,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Sns": {
		//	      "additionalProperties": false,
		//	      "description": "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
		//	      "properties": {
		//	        "TopicArn": {
		//	          "maxLength": 1024,
		//	          "minLength": 36,
		//	          "pattern": "^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Filters
				"filters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MessageTypes
						"message_types": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "DevOps Guru message types to filter for",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 5),
								listvalidator.ValueStringsAre(
									stringvalidator.OneOf(
										"NEW_INSIGHT",
										"CLOSED_INSIGHT",
										"NEW_ASSOCIATION",
										"SEVERITY_UPGRADED",
										"NEW_RECOMMENDATION",
									),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Severities
						"severities": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "DevOps Guru insight severities to filter for",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 3),
								listvalidator.ValueStringsAre(
									stringvalidator.OneOf(
										"LOW",
										"MEDIUM",
										"HIGH",
									),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Information about filters of a notification channel configured in DevOpsGuru to filter for insights.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Sns
				"sns": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TopicArn
						"topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(36, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information about notification channels you have configured with DevOps Guru.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a notification channel.",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
		//	  "type": "string"
		//	}
		"notification_channel_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a notification channel.",
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
		Description: "This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::NotificationChannel").WithTerraformTypeName("awscc_devopsguru_notification_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"config":                  "Config",
		"filters":                 "Filters",
		"message_types":           "MessageTypes",
		"notification_channel_id": "Id",
		"severities":              "Severities",
		"sns":                     "Sns",
		"topic_arn":               "TopicArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
