// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cloudformation_type_activation", typeActivationResource)
}

// typeActivationResource returns the Terraform awscc_cloudformation_type_activation resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFormation::TypeActivation resource.
func typeActivationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the extension.",
		//	  "pattern": "arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the extension.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoUpdate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.",
		//	  "type": "boolean"
		//	}
		"auto_update": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// AutoUpdate is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ExecutionRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
		//	  "type": "string"
		//	}
		"execution_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// ExecutionRoleArn is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: LoggingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies logging configuration information for a type.",
		//	  "properties": {
		//	    "LogGroupName": {
		//	      "description": "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
		//	      "maxLength": 512,
		//	      "minLength": 1,
		//	      "pattern": "^[\\.\\-_/#A-Za-z0-9]+$",
		//	      "type": "string"
		//	    },
		//	    "LogRoleArn": {
		//	      "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"logging_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogGroupName
				"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 512),
						stringvalidator.RegexMatches(regexp.MustCompile("^[\\.\\-_/#A-Za-z0-9]+$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LogRoleArn
				"log_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 256),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies logging configuration information for a type.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// LoggingConfig is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: MajorVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Major Version of the type you want to enable",
		//	  "maxLength": 100000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"major_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Major Version of the type you want to enable",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// MajorVersion is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: PublicTypeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
		//	  "maxLength": 1024,
		//	  "pattern": "arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+",
		//	  "type": "string"
		//	}
		"public_type_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) assigned to the public extension upon publication",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PublisherId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The publisher id assigned by CloudFormation for publishing in this region.",
		//	  "maxLength": 40,
		//	  "minLength": 1,
		//	  "pattern": "[0-9a-zA-Z]{40}",
		//	  "type": "string"
		//	}
		"publisher_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The publisher id assigned by CloudFormation for publishing in this region.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 40),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9a-zA-Z]{40}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The kind of extension",
		//	  "enum": [
		//	    "RESOURCE",
		//	    "MODULE",
		//	    "HOOK"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The kind of extension",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"RESOURCE",
					"MODULE",
					"HOOK",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Type is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: TypeName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
		//	  "pattern": "[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}",
		//	  "type": "string"
		//	}
		"type_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TypeNameAlias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.",
		//	  "maxLength": 204,
		//	  "minLength": 10,
		//	  "pattern": "[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}",
		//	  "type": "string"
		//	}
		"type_name_alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(10, 204),
				stringvalidator.RegexMatches(regexp.MustCompile("[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VersionBump
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled",
		//	  "enum": [
		//	    "MAJOR",
		//	    "MINOR"
		//	  ],
		//	  "type": "string"
		//	}
		"version_bump": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"MAJOR",
					"MINOR",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// VersionBump is a write-only property.
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Enable a resource that has been published in the CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::TypeActivation").WithTerraformTypeName("awscc_cloudformation_type_activation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"auto_update":        "AutoUpdate",
		"execution_role_arn": "ExecutionRoleArn",
		"log_group_name":     "LogGroupName",
		"log_role_arn":       "LogRoleArn",
		"logging_config":     "LoggingConfig",
		"major_version":      "MajorVersion",
		"public_type_arn":    "PublicTypeArn",
		"publisher_id":       "PublisherId",
		"type":               "Type",
		"type_name":          "TypeName",
		"type_name_alias":    "TypeNameAlias",
		"version_bump":       "VersionBump",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ExecutionRoleArn",
		"/properties/Type",
		"/properties/LoggingConfig",
		"/properties/VersionBump",
		"/properties/AutoUpdate",
		"/properties/MajorVersion",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
