// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iam_group_policy", groupPolicyDataSource)
}

// groupPolicyDataSource returns the Terraform awscc_iam_group_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::IAM::GroupPolicy resource.
func groupPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: GroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the group to associate the policy with.",
		//	  "type": "string"
		//	}
		"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the group to associate the policy with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The policy document.",
		//	  "type": "object"
		//	}
		"policy_document": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The policy document.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the policy document.",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the policy document.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IAM::GroupPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::GroupPolicy").WithTerraformTypeName("awscc_iam_group_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"group_name":      "GroupName",
		"policy_document": "PolicyDocument",
		"policy_name":     "PolicyName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
