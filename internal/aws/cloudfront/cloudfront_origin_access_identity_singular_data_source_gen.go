// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_cloudfront_origin_access_identity", cloudFrontOriginAccessIdentityDataSource)
}

// cloudFrontOriginAccessIdentityDataSource returns the Terraform awscc_cloudfront_cloudfront_origin_access_identity data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::CloudFrontOriginAccessIdentity resource.
func cloudFrontOriginAccessIdentityDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CloudFrontOriginAccessIdentityConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The current configuration information for the identity.",
		//	  "properties": {
		//	    "Comment": {
		//	      "description": "A comment to describe the origin access identity. The comment cannot be longer than 128 characters.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Comment"
		//	  ],
		//	  "type": "object"
		//	}
		"cloudfront_origin_access_identity_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A comment to describe the origin access identity. The comment cannot be longer than 128 characters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The current configuration information for the identity.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"cloudfront_origin_access_identity_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: S3CanonicalUserId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"s3_canonical_user_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::CloudFrontOriginAccessIdentity",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::CloudFrontOriginAccessIdentity").WithTerraformTypeName("awscc_cloudfront_cloudfront_origin_access_identity")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudfront_origin_access_identity_config": "CloudFrontOriginAccessIdentityConfig",
		"cloudfront_origin_access_identity_id":     "Id",
		"comment":                                  "Comment",
		"s3_canonical_user_id":                     "S3CanonicalUserId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
