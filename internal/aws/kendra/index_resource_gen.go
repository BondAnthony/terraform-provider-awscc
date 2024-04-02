// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package kendra

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_kendra_index", indexResource)
}

// indexResource returns the Terraform awscc_kendra_index resource.
// This Terraform resource corresponds to the CloudFormation AWS::Kendra::Index resource.
func indexResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CapacityUnits
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Capacity units",
		//	  "properties": {
		//	    "QueryCapacityUnits": {
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "StorageCapacityUnits": {
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "StorageCapacityUnits",
		//	    "QueryCapacityUnits"
		//	  ],
		//	  "type": "object"
		//	}
		"capacity_units": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: QueryCapacityUnits
				"query_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: StorageCapacityUnits
				"storage_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Capacity units",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the index",
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the index",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DocumentMetadataConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Document metadata configurations",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Name": {
		//	        "maxLength": 30,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Relevance": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Duration": {
		//	            "maxLength": 10,
		//	            "minLength": 1,
		//	            "pattern": "[0-9]+[s]",
		//	            "type": "string"
		//	          },
		//	          "Freshness": {
		//	            "type": "boolean"
		//	          },
		//	          "Importance": {
		//	            "maximum": 10,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "RankOrder": {
		//	            "enum": [
		//	              "ASCENDING",
		//	              "DESCENDING"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "ValueImportanceItems": {
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Key": {
		//	                  "maxLength": 50,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "maximum": 10,
		//	                  "minimum": 1,
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Search": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Displayable": {
		//	            "type": "boolean"
		//	          },
		//	          "Facetable": {
		//	            "type": "boolean"
		//	          },
		//	          "Searchable": {
		//	            "type": "boolean"
		//	          },
		//	          "Sortable": {
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "STRING_VALUE",
		//	          "STRING_LIST_VALUE",
		//	          "LONG_VALUE",
		//	          "DATE_VALUE"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 500,
		//	  "type": "array"
		//	}
		"document_metadata_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 30),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Relevance
					"relevance": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Duration
							"duration": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 10),
									stringvalidator.RegexMatches(regexp.MustCompile("[0-9]+[s]"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Freshness
							"freshness": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Importance
							"importance": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(1, 10),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: RankOrder
							"rank_order": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"ASCENDING",
										"DESCENDING",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ValueImportanceItems
							"value_importance_items": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Key
										"key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 50),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.Int64{ /*START VALIDATORS*/
												int64validator.Between(1, 10),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
												int64planmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Search
					"search": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Displayable
							"displayable": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Facetable
							"facetable": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Searchable
							"searchable": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Sortable
							"sortable": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
									boolplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"STRING_VALUE",
								"STRING_LIST_VALUE",
								"LONG_VALUE",
								"DATE_VALUE",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Document metadata configurations",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Edition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Edition of index",
		//	  "enum": [
		//	    "DEVELOPER_EDITION",
		//	    "ENTERPRISE_EDITION"
		//	  ],
		//	  "type": "string"
		//	}
		"edition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Edition of index",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DEVELOPER_EDITION",
					"ENTERPRISE_EDITION",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique ID of index",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "type": "string"
		//	}
		"index_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique ID of index",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of index",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of index",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1000),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Role Arn",
		//	  "maxLength": 1284,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Role Arn",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1284),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ServerSideEncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Server side encryption configuration",
		//	  "properties": {
		//	    "KmsKeyId": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"server_side_encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Server side encryption configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags for labeling the index",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A label for tagging Kendra resources",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A string used to identify this tag",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A string containing the value for the tag",
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
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify this tag",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string containing the value for the tag",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags for labeling the index",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserContextPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ATTRIBUTE_FILTER",
		//	    "USER_TOKEN"
		//	  ],
		//	  "type": "string"
		//	}
		"user_context_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ATTRIBUTE_FILTER",
					"USER_TOKEN",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserTokenConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "JsonTokenTypeConfiguration": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "GroupAttributeField": {
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "UserNameAttributeField": {
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "UserNameAttributeField",
		//	          "GroupAttributeField"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "JwtTokenTypeConfiguration": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ClaimRegex": {
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "GroupAttributeField": {
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Issuer": {
		//	            "maxLength": 65,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "KeyLocation": {
		//	            "enum": [
		//	              "URL",
		//	              "SECRET_MANAGER"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "SecretManagerArn": {
		//	            "description": "Role Arn",
		//	            "maxLength": 1284,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          },
		//	          "URL": {
		//	            "maxLength": 2048,
		//	            "minLength": 1,
		//	            "pattern": "^(https?|ftp|file):\\/\\/([^\\s]*)",
		//	            "type": "string"
		//	          },
		//	          "UserNameAttributeField": {
		//	            "maxLength": 100,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "KeyLocation"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "type": "array"
		//	}
		"user_token_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: JsonTokenTypeConfiguration
					"json_token_type_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: GroupAttributeField
							"group_attribute_field": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 100),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: UserNameAttributeField
							"user_name_attribute_field": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 100),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: JwtTokenTypeConfiguration
					"jwt_token_type_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ClaimRegex
							"claim_regex": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 100),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: GroupAttributeField
							"group_attribute_field": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 100),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Issuer
							"issuer": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 65),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: KeyLocation
							"key_location": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"URL",
										"SECRET_MANAGER",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: SecretManagerArn
							"secret_manager_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Role Arn",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 1284),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: URL
							"url": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 2048),
									stringvalidator.RegexMatches(regexp.MustCompile("^(https?|ftp|file):\\/\\/([^\\s]*)"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: UserNameAttributeField
							"user_name_attribute_field": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 100),
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
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(1),
			}, /*END VALIDATORS*/
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
		Description: "A Kendra index",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kendra::Index").WithTerraformTypeName("awscc_kendra_index")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"capacity_units":                       "CapacityUnits",
		"claim_regex":                          "ClaimRegex",
		"description":                          "Description",
		"displayable":                          "Displayable",
		"document_metadata_configurations":     "DocumentMetadataConfigurations",
		"duration":                             "Duration",
		"edition":                              "Edition",
		"facetable":                            "Facetable",
		"freshness":                            "Freshness",
		"group_attribute_field":                "GroupAttributeField",
		"importance":                           "Importance",
		"index_id":                             "Id",
		"issuer":                               "Issuer",
		"json_token_type_configuration":        "JsonTokenTypeConfiguration",
		"jwt_token_type_configuration":         "JwtTokenTypeConfiguration",
		"key":                                  "Key",
		"key_location":                         "KeyLocation",
		"kms_key_id":                           "KmsKeyId",
		"name":                                 "Name",
		"query_capacity_units":                 "QueryCapacityUnits",
		"rank_order":                           "RankOrder",
		"relevance":                            "Relevance",
		"role_arn":                             "RoleArn",
		"search":                               "Search",
		"searchable":                           "Searchable",
		"secret_manager_arn":                   "SecretManagerArn",
		"server_side_encryption_configuration": "ServerSideEncryptionConfiguration",
		"sortable":                             "Sortable",
		"storage_capacity_units":               "StorageCapacityUnits",
		"tags":                                 "Tags",
		"type":                                 "Type",
		"url":                                  "URL",
		"user_context_policy":                  "UserContextPolicy",
		"user_name_attribute_field":            "UserNameAttributeField",
		"user_token_configurations":            "UserTokenConfigurations",
		"value":                                "Value",
		"value_importance_items":               "ValueImportanceItems",
	})

	opts = opts.WithCreateTimeoutInMinutes(240).WithDeleteTimeoutInMinutes(720)

	opts = opts.WithUpdateTimeoutInMinutes(240)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
