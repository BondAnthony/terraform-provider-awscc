// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cur

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/defaults"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cur_report_definition", reportDefinitionResource)
}

// reportDefinitionResource returns the Terraform awscc_cur_report_definition resource.
// This Terraform resource corresponds to the CloudFormation AWS::CUR::ReportDefinition resource.
func reportDefinitionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdditionalArtifacts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "description": "A list of manifests that you want Amazon Web Services to create for this report.",
		//	  "items": {
		//	    "description": "The types of manifest that you want AWS to create for this report.",
		//	    "enum": [
		//	      "REDSHIFT",
		//	      "QUICKSIGHT",
		//	      "ATHENA"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"additional_artifacts": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of manifests that you want Amazon Web Services to create for this report.",
			Optional:    true,
			Computed:    true,
			Default:     defaults.StaticListOfString(),
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"REDSHIFT",
						"QUICKSIGHT",
						"ATHENA",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AdditionalSchemaElements
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": [],
		//	  "description": "A list of strings that indicate additional content that Amazon Web Services includes in the report, such as individual resource IDs.",
		//	  "items": {
		//	    "description": "Whether or not AWS includes resource IDs in the report.",
		//	    "enum": [
		//	      "RESOURCES",
		//	      "SPLIT_COST_ALLOCATION_DATA",
		//	      "MANUAL_DISCOUNT_COMPATIBILITY"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"additional_schema_elements": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of strings that indicate additional content that Amazon Web Services includes in the report, such as individual resource IDs.",
			Optional:    true,
			Computed:    true,
			Default:     defaults.StaticListOfString(),
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"RESOURCES",
						"SPLIT_COST_ALLOCATION_DATA",
						"MANUAL_DISCOUNT_COMPATIBILITY",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BillingViewArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon resource name of the billing view. You can get this value by using the billing view service public APIs.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "(arn:aws(-cn)?:billing::[0-9]{12}:billingview/)?[a-zA-Z0-9_\\+=\\.\\-@].{1,30}",
		//	  "type": "string"
		//	}
		"billing_view_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon resource name of the billing view. You can get this value by using the billing view service public APIs.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("(arn:aws(-cn)?:billing::[0-9]{12}:billingview/)?[a-zA-Z0-9_\\+=\\.\\-@].{1,30}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Compression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The compression format that AWS uses for the report.",
		//	  "enum": [
		//	    "ZIP",
		//	    "GZIP",
		//	    "Parquet"
		//	  ],
		//	  "type": "string"
		//	}
		"compression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The compression format that AWS uses for the report.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ZIP",
					"GZIP",
					"Parquet",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Format
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The format that AWS saves the report in.",
		//	  "enum": [
		//	    "textORcsv",
		//	    "Parquet"
		//	  ],
		//	  "type": "string"
		//	}
		"format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The format that AWS saves the report in.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"textORcsv",
					"Parquet",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RefreshClosedReports
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees.",
		//	  "type": "boolean"
		//	}
		"refresh_closed_reports": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReportName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[0-9A-Za-z!\\-_.*\\'()]+",
		//	  "type": "string"
		//	}
		"report_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9A-Za-z!\\-_.*\\'()]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReportVersioning
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether you want Amazon Web Services to overwrite the previous version of each report or to deliver the report in addition to the previous versions.",
		//	  "enum": [
		//	    "CREATE_NEW_REPORT",
		//	    "OVERWRITE_REPORT"
		//	  ],
		//	  "type": "string"
		//	}
		"report_versioning": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Whether you want Amazon Web Services to overwrite the previous version of each report or to deliver the report in addition to the previous versions.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CREATE_NEW_REPORT",
					"OVERWRITE_REPORT",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3Bucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The S3 bucket where AWS delivers the report.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[A-Za-z0-9_\\.\\-]+",
		//	  "type": "string"
		//	}
		"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The S3 bucket where AWS delivers the report.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("[A-Za-z0-9_\\.\\-]+"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: S3Prefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[0-9A-Za-z!\\-_.*\\'()/]*",
		//	  "type": "string"
		//	}
		"s3_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9A-Za-z!\\-_.*\\'()/]*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: S3Region
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The region of the S3 bucket that AWS delivers the report into.",
		//	  "type": "string"
		//	}
		"s3_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The region of the S3 bucket that AWS delivers the report into.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: TimeUnit
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The granularity of the line items in the report.",
		//	  "enum": [
		//	    "HOURLY",
		//	    "DAILY",
		//	    "MONTHLY"
		//	  ],
		//	  "type": "string"
		//	}
		"time_unit": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The granularity of the line items in the report.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"HOURLY",
					"DAILY",
					"MONTHLY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "The AWS::CUR::ReportDefinition resource creates a Cost & Usage Report with user-defined settings. You can use this resource to define settings like time granularity (hourly, daily, monthly), file format (Parquet, CSV), and S3 bucket for delivery of these reports.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CUR::ReportDefinition").WithTerraformTypeName("awscc_cur_report_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_artifacts":       "AdditionalArtifacts",
		"additional_schema_elements": "AdditionalSchemaElements",
		"billing_view_arn":           "BillingViewArn",
		"compression":                "Compression",
		"format":                     "Format",
		"refresh_closed_reports":     "RefreshClosedReports",
		"report_name":                "ReportName",
		"report_versioning":          "ReportVersioning",
		"s3_bucket":                  "S3Bucket",
		"s3_prefix":                  "S3Prefix",
		"s3_region":                  "S3Region",
		"time_unit":                  "TimeUnit",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
