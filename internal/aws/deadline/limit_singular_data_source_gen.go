// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package deadline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_deadline_limit", limitDataSource)
}

// limitDataSource returns the Terraform awscc_deadline_limit data source.
// This Terraform data source corresponds to the CloudFormation AWS::Deadline::Limit resource.
func limitDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AmountRequirementName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"amount_requirement_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CurrentCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 2147483647,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"current_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "",
		//	  "maxLength": 100,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FarmId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^farm-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"farm_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LimitId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^limit-[0-9a-f]{32}$",
		//	  "type": "string"
		//	}
		"limit_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MaxCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 2147483647,
		//	  "minimum": -1,
		//	  "type": "integer"
		//	}
		"max_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Deadline::Limit",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Deadline::Limit").WithTerraformTypeName("awscc_deadline_limit")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"amount_requirement_name": "AmountRequirementName",
		"current_count":           "CurrentCount",
		"description":             "Description",
		"display_name":            "DisplayName",
		"farm_id":                 "FarmId",
		"limit_id":                "LimitId",
		"max_count":               "MaxCount",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
