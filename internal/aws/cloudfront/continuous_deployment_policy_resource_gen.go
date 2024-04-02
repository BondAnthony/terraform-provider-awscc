// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cloudfront_continuous_deployment_policy", continuousDeploymentPolicyResource)
}

// continuousDeploymentPolicyResource returns the Terraform awscc_cloudfront_continuous_deployment_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFront::ContinuousDeploymentPolicy resource.
func continuousDeploymentPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContinuousDeploymentPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Enabled": {
		//	      "type": "boolean"
		//	    },
		//	    "SingleHeaderPolicyConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Header": {
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Value": {
		//	          "maxLength": 1783,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Header",
		//	        "Value"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "SingleWeightPolicyConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SessionStickinessConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "IdleTTL": {
		//	              "maximum": 3600,
		//	              "minimum": 300,
		//	              "type": "integer"
		//	            },
		//	            "MaximumTTL": {
		//	              "maximum": 3600,
		//	              "minimum": 300,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "required": [
		//	            "IdleTTL",
		//	            "MaximumTTL"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Weight": {
		//	          "maximum": 1,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "Weight"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "StagingDistributionDnsNames": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TrafficConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SingleHeaderConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Header": {
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 1783,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Header",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "SingleWeightConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "SessionStickinessConfig": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "IdleTTL": {
		//	                  "maximum": 3600,
		//	                  "minimum": 300,
		//	                  "type": "integer"
		//	                },
		//	                "MaximumTTL": {
		//	                  "maximum": 3600,
		//	                  "minimum": 300,
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "required": [
		//	                "IdleTTL",
		//	                "MaximumTTL"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "Weight": {
		//	              "maximum": 1,
		//	              "minimum": 0,
		//	              "type": "number"
		//	            }
		//	          },
		//	          "required": [
		//	            "Weight"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Type": {
		//	          "enum": [
		//	            "SingleWeight",
		//	            "SingleHeader"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Type"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "enum": [
		//	        "SingleWeight",
		//	        "SingleHeader"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Enabled",
		//	    "StagingDistributionDnsNames"
		//	  ],
		//	  "type": "object"
		//	}
		"continuous_deployment_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: SingleHeaderPolicyConfig
				"single_header_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Header
						"header": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 256),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Value
						"value": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 1783),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SingleWeightPolicyConfig
				"single_weight_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SessionStickinessConfig
						"session_stickiness_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IdleTTL
								"idle_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.Between(300, 3600),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: MaximumTTL
								"maximum_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.Between(300, 3600),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Weight
						"weight": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 1.000000),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StagingDistributionDnsNames
				"staging_distribution_dns_names": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
						listvalidator.UniqueValues(),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: TrafficConfig
				"traffic_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SingleHeaderConfig
						"single_header_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Header
								"header": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 256),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 1783),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: SingleWeightConfig
						"single_weight_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: SessionStickinessConfig
								"session_stickiness_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: IdleTTL
										"idle_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Required: true,
											Validators: []validator.Int64{ /*START VALIDATORS*/
												int64validator.Between(300, 3600),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
										// Property: MaximumTTL
										"maximum_ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Required: true,
											Validators: []validator.Int64{ /*START VALIDATORS*/
												int64validator.Between(300, 3600),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Weight
								"weight": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.Between(0.000000, 1.000000),
									}, /*END VALIDATORS*/
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
									"SingleWeight",
									"SingleHeader",
								),
							}, /*END VALIDATORS*/
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
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SingleWeight",
							"SingleHeader",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"continuous_deployment_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		Description: "Resource Type definition for AWS::CloudFront::ContinuousDeploymentPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::ContinuousDeploymentPolicy").WithTerraformTypeName("awscc_cloudfront_continuous_deployment_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"continuous_deployment_policy_config": "ContinuousDeploymentPolicyConfig",
		"continuous_deployment_policy_id":     "Id",
		"enabled":                             "Enabled",
		"header":                              "Header",
		"idle_ttl":                            "IdleTTL",
		"last_modified_time":                  "LastModifiedTime",
		"maximum_ttl":                         "MaximumTTL",
		"session_stickiness_config":           "SessionStickinessConfig",
		"single_header_config":                "SingleHeaderConfig",
		"single_header_policy_config":         "SingleHeaderPolicyConfig",
		"single_weight_config":                "SingleWeightConfig",
		"single_weight_policy_config":         "SingleWeightPolicyConfig",
		"staging_distribution_dns_names":      "StagingDistributionDnsNames",
		"traffic_config":                      "TrafficConfig",
		"type":                                "Type",
		"value":                               "Value",
		"weight":                              "Weight",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
