// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package opensearchserverless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_opensearchserverless_collection", collectionDataSource)
}

// collectionDataSource returns the Terraform awscc_opensearchserverless_collection data source.
// This Terraform data source corresponds to the CloudFormation AWS::OpenSearchServerless::Collection resource.
func collectionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the collection.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the collection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CollectionEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The endpoint for the collection.",
		//	  "type": "string"
		//	}
		"collection_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The endpoint for the collection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OpenSearch Dashboards endpoint for the collection.",
		//	  "type": "string"
		//	}
		"dashboard_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OpenSearch Dashboards endpoint for the collection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the collection",
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the collection",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the collection",
		//	  "maxLength": 40,
		//	  "minLength": 3,
		//	  "type": "string"
		//	}
		"collection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the collection",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the collection.\n\nThe name must meet the following criteria:\nUnique to your account and AWS Region\nStarts with a lowercase letter\nContains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)\nContains between 3 and 32 characters\n",
		//	  "maxLength": 32,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z][a-z0-9-]{2,31}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the collection.\n\nThe name must meet the following criteria:\nUnique to your account and AWS Region\nStarts with a lowercase letter\nContains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)\nContains between 3 and 32 characters\n",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StandbyReplicas
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The possible standby replicas for the collection",
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"standby_replicas": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The possible standby replicas for the collection",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of tags to be added to the resource",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair metadata associated with resource",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key in the key-value pair",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value in the key-value pair",
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
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key in the key-value pair",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value in the key-value pair",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of tags to be added to the resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The possible types for the collection",
		//	  "enum": [
		//	    "SEARCH",
		//	    "TIMESERIES",
		//	    "VECTORSEARCH"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The possible types for the collection",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::OpenSearchServerless::Collection",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::OpenSearchServerless::Collection").WithTerraformTypeName("awscc_opensearchserverless_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"collection_endpoint": "CollectionEndpoint",
		"collection_id":       "Id",
		"dashboard_endpoint":  "DashboardEndpoint",
		"description":         "Description",
		"key":                 "Key",
		"name":                "Name",
		"standby_replicas":    "StandbyReplicas",
		"tags":                "Tags",
		"type":                "Type",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
