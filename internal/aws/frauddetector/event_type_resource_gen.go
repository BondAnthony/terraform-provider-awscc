// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package frauddetector

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_frauddetector_event_type", eventTypeResource)
}

// eventTypeResource returns the Terraform awscc_frauddetector_event_type resource.
// This Terraform resource corresponds to the CloudFormation AWS::FraudDetector::EventType resource.
func eventTypeResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the event type.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the event type.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the event type was created.",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the event type was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the event type.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the event type.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EntityTypes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      },
		//	      "CreatedTime": {
		//	        "description": "The time when the event type was created.",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "The description.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Inline": {
		//	        "type": "boolean"
		//	      },
		//	      "LastUpdatedTime": {
		//	        "description": "The time when the event type was last updated.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "Tags associated with this event type.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 200,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"entity_types": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was created.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Inline
					"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LastUpdatedTime
					"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was last updated.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Tags
					"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 128),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(0, 256),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Tags associated with this event type.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.SizeAtMost(200),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventVariables
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      },
		//	      "CreatedTime": {
		//	        "description": "The time when the event type was created.",
		//	        "type": "string"
		//	      },
		//	      "DataSource": {
		//	        "enum": [
		//	          "EVENT"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "DataType": {
		//	        "enum": [
		//	          "STRING",
		//	          "INTEGER",
		//	          "FLOAT",
		//	          "BOOLEAN"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "DefaultValue": {
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "The description.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Inline": {
		//	        "type": "boolean"
		//	      },
		//	      "LastUpdatedTime": {
		//	        "description": "The time when the event type was last updated.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "Tags associated with this event type.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 200,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "VariableType": {
		//	        "enum": [
		//	          "AUTH_CODE",
		//	          "AVS",
		//	          "BILLING_ADDRESS_L1",
		//	          "BILLING_ADDRESS_L2",
		//	          "BILLING_CITY",
		//	          "BILLING_COUNTRY",
		//	          "BILLING_NAME",
		//	          "BILLING_PHONE",
		//	          "BILLING_STATE",
		//	          "BILLING_ZIP",
		//	          "CARD_BIN",
		//	          "CATEGORICAL",
		//	          "CURRENCY_CODE",
		//	          "EMAIL_ADDRESS",
		//	          "FINGERPRINT",
		//	          "FRAUD_LABEL",
		//	          "FREE_FORM_TEXT",
		//	          "IP_ADDRESS",
		//	          "NUMERIC",
		//	          "ORDER_ID",
		//	          "PAYMENT_TYPE",
		//	          "PHONE_NUMBER",
		//	          "PRICE",
		//	          "PRODUCT_CATEGORY",
		//	          "SHIPPING_ADDRESS_L1",
		//	          "SHIPPING_ADDRESS_L2",
		//	          "SHIPPING_CITY",
		//	          "SHIPPING_COUNTRY",
		//	          "SHIPPING_NAME",
		//	          "SHIPPING_PHONE",
		//	          "SHIPPING_STATE",
		//	          "SHIPPING_ZIP",
		//	          "USERAGENT"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"event_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was created.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DataSource
					"data_source": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"EVENT",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DataType
					"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"STRING",
								"INTEGER",
								"FLOAT",
								"BOOLEAN",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DefaultValue
					"default_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Inline
					"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LastUpdatedTime
					"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was last updated.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Tags
					"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 128),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(0, 256),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Tags associated with this event type.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.SizeAtMost(200),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: VariableType
					"variable_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"AUTH_CODE",
								"AVS",
								"BILLING_ADDRESS_L1",
								"BILLING_ADDRESS_L2",
								"BILLING_CITY",
								"BILLING_COUNTRY",
								"BILLING_NAME",
								"BILLING_PHONE",
								"BILLING_STATE",
								"BILLING_ZIP",
								"CARD_BIN",
								"CATEGORICAL",
								"CURRENCY_CODE",
								"EMAIL_ADDRESS",
								"FINGERPRINT",
								"FRAUD_LABEL",
								"FREE_FORM_TEXT",
								"IP_ADDRESS",
								"NUMERIC",
								"ORDER_ID",
								"PAYMENT_TYPE",
								"PHONE_NUMBER",
								"PRICE",
								"PRODUCT_CATEGORY",
								"SHIPPING_ADDRESS_L1",
								"SHIPPING_ADDRESS_L2",
								"SHIPPING_CITY",
								"SHIPPING_COUNTRY",
								"SHIPPING_NAME",
								"SHIPPING_PHONE",
								"SHIPPING_STATE",
								"SHIPPING_ZIP",
								"USERAGENT",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Labels
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      },
		//	      "CreatedTime": {
		//	        "description": "The time when the event type was created.",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "The description.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Inline": {
		//	        "type": "boolean"
		//	      },
		//	      "LastUpdatedTime": {
		//	        "description": "The time when the event type was last updated.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "Tags associated with this event type.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 200,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 2,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"labels": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was created.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Inline
					"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LastUpdatedTime
					"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was last updated.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Tags
					"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 128),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(0, 256),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Tags associated with this event type.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.SizeAtMost(200),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(2),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the event type was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the event type was last updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the event type",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-z_-]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the event type",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-z_-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags associated with this event type.",
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
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags associated with this event type.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
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
		Description: "A resource schema for an EventType in Amazon Fraud Detector.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::EventType").WithTerraformTypeName("awscc_frauddetector_event_type")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"created_time":      "CreatedTime",
		"data_source":       "DataSource",
		"data_type":         "DataType",
		"default_value":     "DefaultValue",
		"description":       "Description",
		"entity_types":      "EntityTypes",
		"event_variables":   "EventVariables",
		"inline":            "Inline",
		"key":               "Key",
		"labels":            "Labels",
		"last_updated_time": "LastUpdatedTime",
		"name":              "Name",
		"tags":              "Tags",
		"value":             "Value",
		"variable_type":     "VariableType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
