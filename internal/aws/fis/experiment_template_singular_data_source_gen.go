// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package fis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_fis_experiment_template", experimentTemplateDataSource)
}

// experimentTemplateDataSource returns the Terraform awscc_fis_experiment_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::FIS::ExperimentTemplate resource.
func experimentTemplateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Actions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The actions for the experiment.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies an action for the experiment template.",
		//	      "properties": {
		//	        "ActionId": {
		//	          "description": "The ID of the action.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "Description": {
		//	          "description": "A description for the action.",
		//	          "maxLength": 512,
		//	          "type": "string"
		//	        },
		//	        "Parameters": {
		//	          "additionalProperties": false,
		//	          "description": "The parameters for the action, if applicable.",
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 1024,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "StartAfter": {
		//	          "description": "The names of the actions that must be completed before the current action starts.",
		//	          "items": {
		//	            "maxLength": 64,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "Targets": {
		//	          "additionalProperties": false,
		//	          "description": "One or more targets for the action.",
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 64,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ActionId"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"actions":                 // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ActionId
					"action_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the action.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A description for the action.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Parameters
					"parameters":        // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The parameters for the action, if applicable.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: StartAfter
					"start_after": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The names of the actions that must be completed before the current action starts.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Targets
					"targets":           // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "One or more targets for the action.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The actions for the experiment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the experiment template.",
		//	  "maxLength": 512,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the experiment template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExperimentOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AccountTargeting": {
		//	      "description": "The account targeting setting for the experiment template.",
		//	      "enum": [
		//	        "multi-account",
		//	        "single-account"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "EmptyTargetResolutionMode": {
		//	      "description": "The target resolution failure mode for the experiment template.",
		//	      "enum": [
		//	        "fail",
		//	        "skip"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"experiment_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccountTargeting
				"account_targeting": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The account targeting setting for the experiment template.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EmptyTargetResolutionMode
				"empty_target_resolution_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The target resolution failure mode for the experiment template.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ExperimentReportConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DataSources": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CloudWatchDashboards": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "DashboardIdentifier": {
		//	                "maxLength": 512,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "DashboardIdentifier"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Outputs": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ExperimentReportS3Configuration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "BucketName": {
		//	              "maxLength": 63,
		//	              "minLength": 3,
		//	              "type": "string"
		//	            },
		//	            "Prefix": {
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "BucketName"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ExperimentReportS3Configuration"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "PostExperimentDuration": {
		//	      "type": "string"
		//	    },
		//	    "PreExperimentDuration": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Outputs"
		//	  ],
		//	  "type": "object"
		//	}
		"experiment_report_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataSources
				"data_sources": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchDashboards
						"cloudwatch_dashboards": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DashboardIdentifier
									"dashboard_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Outputs
				"outputs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ExperimentReportS3Configuration
						"experiment_report_s3_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BucketName
								"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Prefix
								"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PostExperimentDuration
				"post_experiment_duration": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PreExperimentDuration
				"pre_experiment_duration": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"experiment_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LogConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CloudWatchLogsConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "LogGroupArn": {
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "LogGroupArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LogSchemaVersion": {
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "S3Configuration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "BucketName": {
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "type": "string"
		//	        },
		//	        "Prefix": {
		//	          "maxLength": 700,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "BucketName"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "LogSchemaVersion"
		//	  ],
		//	  "type": "object"
		//	}
		"log_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudWatchLogsConfiguration
				"cloudwatch_logs_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LogGroupArn
						"log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LogSchemaVersion
				"log_schema_version": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: S3Configuration
				"s3_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketName
						"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Prefix
						"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.",
		//	  "maxLength": 1224,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StopConditions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more stop conditions.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Source": {
		//	        "maxLength": 64,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 2048,
		//	        "minLength": 20,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Source"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"stop_conditions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Source
					"source": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more stop conditions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Targets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The targets for the experiment.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies a target for an experiment.",
		//	      "properties": {
		//	        "Filters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "Describes a filter used for the target resource input in an experiment template.",
		//	            "properties": {
		//	              "Path": {
		//	                "description": "The attribute path for the filter.",
		//	                "maxLength": 256,
		//	                "type": "string"
		//	              },
		//	              "Values": {
		//	                "description": "The attribute values for the filter.",
		//	                "items": {
		//	                  "maxLength": 128,
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "required": [
		//	              "Path",
		//	              "Values"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "Parameters": {
		//	          "additionalProperties": false,
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 1024,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ResourceArns": {
		//	          "description": "The Amazon Resource Names (ARNs) of the target resources.",
		//	          "items": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "ResourceTags": {
		//	          "additionalProperties": false,
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ResourceType": {
		//	          "description": "The AWS resource type. The resource type must be supported for the specified action.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "SelectionMode": {
		//	          "description": "Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ResourceType",
		//	        "SelectionMode"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"targets":                 // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Filters
					"filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Path
								"path": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The attribute path for the filter.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Values
								"values": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "The attribute values for the filter.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Parameters
					"parameters":        // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourceArns
					"resource_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The Amazon Resource Names (ARNs) of the target resources.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourceTags
					"resource_tags":     // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourceType
					"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The AWS resource type. The resource type must be supported for the specified action.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SelectionMode
					"selection_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The targets for the experiment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::FIS::ExperimentTemplate",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FIS::ExperimentTemplate").WithTerraformTypeName("awscc_fis_experiment_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_targeting":                  "AccountTargeting",
		"action_id":                          "ActionId",
		"actions":                            "Actions",
		"bucket_name":                        "BucketName",
		"cloudwatch_dashboards":              "CloudWatchDashboards",
		"cloudwatch_logs_configuration":      "CloudWatchLogsConfiguration",
		"dashboard_identifier":               "DashboardIdentifier",
		"data_sources":                       "DataSources",
		"description":                        "Description",
		"empty_target_resolution_mode":       "EmptyTargetResolutionMode",
		"experiment_options":                 "ExperimentOptions",
		"experiment_report_configuration":    "ExperimentReportConfiguration",
		"experiment_report_s3_configuration": "ExperimentReportS3Configuration",
		"experiment_template_id":             "Id",
		"filters":                            "Filters",
		"log_configuration":                  "LogConfiguration",
		"log_group_arn":                      "LogGroupArn",
		"log_schema_version":                 "LogSchemaVersion",
		"outputs":                            "Outputs",
		"parameters":                         "Parameters",
		"path":                               "Path",
		"post_experiment_duration":           "PostExperimentDuration",
		"pre_experiment_duration":            "PreExperimentDuration",
		"prefix":                             "Prefix",
		"resource_arns":                      "ResourceArns",
		"resource_tags":                      "ResourceTags",
		"resource_type":                      "ResourceType",
		"role_arn":                           "RoleArn",
		"s3_configuration":                   "S3Configuration",
		"selection_mode":                     "SelectionMode",
		"source":                             "Source",
		"start_after":                        "StartAfter",
		"stop_conditions":                    "StopConditions",
		"tags":                               "Tags",
		"targets":                            "Targets",
		"value":                              "Value",
		"values":                             "Values",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
