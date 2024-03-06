// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package resourceexplorer2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_resourceexplorer2_default_view_association", defaultViewAssociationResource)
}

// defaultViewAssociationResource returns the Terraform awscc_resourceexplorer2_default_view_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::ResourceExplorer2::DefaultViewAssociation resource.
func defaultViewAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociatedAwsPrincipal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS principal that the default view is associated with, used as the unique identifier for this resource.",
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"associated_aws_principal": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS principal that the default view is associated with, used as the unique identifier for this resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ViewArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"view_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
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
		Description: "Definition of AWS::ResourceExplorer2::DefaultViewAssociation Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResourceExplorer2::DefaultViewAssociation").WithTerraformTypeName("awscc_resourceexplorer2_default_view_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"associated_aws_principal": "AssociatedAwsPrincipal",
		"view_arn":                 "ViewArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
