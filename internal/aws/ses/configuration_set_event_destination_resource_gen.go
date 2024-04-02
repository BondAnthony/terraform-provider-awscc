// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
	registry.AddResourceFactory("awscc_ses_configuration_set_event_destination", configurationSetEventDestinationResource)
}

// configurationSetEventDestinationResource returns the Terraform awscc_ses_configuration_set_event_destination resource.
// This Terraform resource corresponds to the CloudFormation AWS::SES::ConfigurationSetEventDestination resource.
func configurationSetEventDestinationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConfigurationSetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the configuration set that contains the event destination.",
		//	  "type": "string"
		//	}
		"configuration_set_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the configuration set that contains the event destination.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventDestination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The event destination object.",
		//	  "properties": {
		//	    "CloudWatchDestination": {
		//	      "additionalProperties": false,
		//	      "description": "An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.",
		//	      "properties": {
		//	        "DimensionConfigurations": {
		//	          "description": "A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.",
		//	            "properties": {
		//	              "DefaultDimensionValue": {
		//	                "description": "The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.",
		//	                "maxLength": 256,
		//	                "minLength": 1,
		//	                "pattern": "^[a-zA-Z0-9_-]{1,256}$",
		//	                "type": "string"
		//	              },
		//	              "DimensionName": {
		//	                "description": "The name of an Amazon CloudWatch dimension associated with an email sending metric.",
		//	                "maxLength": 256,
		//	                "minLength": 1,
		//	                "pattern": "^[a-zA-Z0-9_:-]{1,256}$",
		//	                "type": "string"
		//	              },
		//	              "DimensionValueSource": {
		//	                "description": "The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch. To use the message tags that you specify using an X-SES-MESSAGE-TAGS header or a parameter to the SendEmail/SendRawEmail API, specify messageTag. To use your own email headers, specify emailHeader. To put a custom tag on any link included in your email, specify linkTag.",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "DimensionValueSource",
		//	              "DefaultDimensionValue",
		//	              "DimensionName"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Enabled": {
		//	      "description": "Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set. Set to true to enable publishing to this destination; set to false to prevent publishing to this destination. The default value is false.   ",
		//	      "type": "boolean"
		//	    },
		//	    "KinesisFirehoseDestination": {
		//	      "additionalProperties": false,
		//	      "description": "An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.",
		//	      "properties": {
		//	        "DeliveryStreamARN": {
		//	          "description": "The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.",
		//	          "type": "string"
		//	        },
		//	        "IAMRoleARN": {
		//	          "description": "The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "IAMRoleARN",
		//	        "DeliveryStreamARN"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "MatchingEventTypes": {
		//	      "description": "The type of email sending events, send, reject, bounce, complaint, delivery, open, click, renderingFailure, deliveryDelay, and subscription.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "Name": {
		//	      "description": "The name of the event destination set.",
		//	      "pattern": "^[a-zA-Z0-9_-]{0,64}$",
		//	      "type": "string"
		//	    },
		//	    "SnsDestination": {
		//	      "additionalProperties": false,
		//	      "description": "An object that contains SNS topic ARN associated event destination.",
		//	      "properties": {
		//	        "TopicARN": {
		//	          "maxLength": 1024,
		//	          "minLength": 36,
		//	          "pattern": "^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TopicARN"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "MatchingEventTypes"
		//	  ],
		//	  "type": "object"
		//	}
		"event_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudWatchDestination
				"cloudwatch_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DimensionConfigurations
						"dimension_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DefaultDimensionValue
									"default_dimension_value": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The default value of the dimension that is published to Amazon CloudWatch if you do not provide the value of the dimension when you send an email.",
										Required:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 256),
											stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{1,256}$"), ""),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: DimensionName
									"dimension_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of an Amazon CloudWatch dimension associated with an email sending metric.",
										Required:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 256),
											stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_:-]{1,256}$"), ""),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: DimensionValueSource
									"dimension_value_source": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The place where Amazon SES finds the value of a dimension to publish to Amazon CloudWatch. To use the message tags that you specify using an X-SES-MESSAGE-TAGS header or a parameter to the SendEmail/SendRawEmail API, specify messageTag. To use your own email headers, specify emailHeader. To put a custom tag on any link included in your email, specify linkTag.",
										Required:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set. Set to true to enable publishing to this destination; set to false to prevent publishing to this destination. The default value is false.   ",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KinesisFirehoseDestination
				"kinesis_firehose_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DeliveryStreamARN
						"delivery_stream_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.",
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: IAMRoleARN
						"iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MatchingEventTypes
				"matching_event_types": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The type of email sending events, send, reject, bounce, complaint, delivery, open, click, renderingFailure, deliveryDelay, and subscription.",
					Required:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the event destination set.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{0,64}$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SnsDestination
				"sns_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TopicARN
						"topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(36, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[a-z0-9-]*:sns:[a-z0-9-]+:\\d{12}:[^:]+$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that contains SNS topic ARN associated event destination.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The event destination object.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"configuration_set_event_destination_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
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
		Description: "Resource Type definition for AWS::SES::ConfigurationSetEventDestination",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::ConfigurationSetEventDestination").WithTerraformTypeName("awscc_ses_configuration_set_event_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudwatch_destination":                 "CloudWatchDestination",
		"configuration_set_event_destination_id": "Id",
		"configuration_set_name":                 "ConfigurationSetName",
		"default_dimension_value":                "DefaultDimensionValue",
		"delivery_stream_arn":                    "DeliveryStreamARN",
		"dimension_configurations":               "DimensionConfigurations",
		"dimension_name":                         "DimensionName",
		"dimension_value_source":                 "DimensionValueSource",
		"enabled":                                "Enabled",
		"event_destination":                      "EventDestination",
		"iam_role_arn":                           "IAMRoleARN",
		"kinesis_firehose_destination":           "KinesisFirehoseDestination",
		"matching_event_types":                   "MatchingEventTypes",
		"name":                                   "Name",
		"sns_destination":                        "SnsDestination",
		"topic_arn":                              "TopicARN",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
