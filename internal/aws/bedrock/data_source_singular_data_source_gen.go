// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package bedrock

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_bedrock_data_source", dataSourceDataSource)
}

// dataSourceDataSource returns the Terraform awscc_bedrock_data_source data source.
// This Terraform data source corresponds to the CloudFormation AWS::Bedrock::DataSource resource.
func dataSourceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time at which the data source was created.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time at which the data source was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataSourceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies a raw data source location to ingest.",
		//	  "properties": {
		//	    "S3Configuration": {
		//	      "additionalProperties": false,
		//	      "description": "Contains information about the S3 configuration of the data source.",
		//	      "properties": {
		//	        "BucketArn": {
		//	          "description": "The ARN of the bucket that contains the data source.",
		//	          "maxLength": 2048,
		//	          "minLength": 1,
		//	          "pattern": "^arn:aws(|-cn|-us-gov):s3:::[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$",
		//	          "type": "string"
		//	        },
		//	        "InclusionPrefixes": {
		//	          "description": "A list of S3 prefixes that define the object containing the data sources.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "Prefix for s3 object.",
		//	            "maxLength": 300,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "BucketArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "description": "The type of the data source location.",
		//	      "enum": [
		//	        "S3"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type",
		//	    "S3Configuration"
		//	  ],
		//	  "type": "object"
		//	}
		"data_source_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Configuration
				"s3_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketArn
						"bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the bucket that contains the data source.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: InclusionPrefixes
						"inclusion_prefixes": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "A list of S3 prefixes that define the object containing the data sources.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains information about the S3 configuration of the data source.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of the data source location.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies a raw data source location to ingest.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataSourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifier for a resource.",
		//	  "pattern": "^[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"data_source_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifier for a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataSourceStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of a data source.",
		//	  "enum": [
		//	    "AVAILABLE",
		//	    "DELETING"
		//	  ],
		//	  "type": "string"
		//	}
		"data_source_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of a data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the Resource.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the Resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBaseId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of the knowledge base to which to add the data source.",
		//	  "pattern": "^[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"knowledge_base_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of the knowledge base to which to add the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the data source.",
		//	  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerSideEncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains details about the server-side encryption for the data source.",
		//	  "properties": {
		//	    "KmsKeyArn": {
		//	      "description": "The ARN of the AWS KMS key used to encrypt the resource.",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "pattern": "^arn:aws(|-cn|-us-gov):kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"server_side_encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the AWS KMS key used to encrypt the resource.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains details about the server-side encryption for the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time at which the knowledge base was last updated.",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time at which the knowledge base was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VectorIngestionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
		//	  "properties": {
		//	    "ChunkingConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
		//	      "properties": {
		//	        "ChunkingStrategy": {
		//	          "description": "Knowledge base can split your source data into chunks. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for NONE, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.",
		//	          "enum": [
		//	            "FIXED_SIZE",
		//	            "NONE"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "FixedSizeChunkingConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "Configurations for when you choose fixed-size chunking. If you set the chunkingStrategy as NONE, exclude this field.",
		//	          "properties": {
		//	            "MaxTokens": {
		//	              "description": "The maximum number of tokens to include in a chunk.",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            },
		//	            "OverlapPercentage": {
		//	              "description": "The percentage of overlap between adjacent chunks of a data source.",
		//	              "maximum": 99,
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "required": [
		//	            "MaxTokens",
		//	            "OverlapPercentage"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ChunkingStrategy"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vector_ingestion_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ChunkingConfiguration
				"chunking_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ChunkingStrategy
						"chunking_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Knowledge base can split your source data into chunks. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for NONE, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: FixedSizeChunkingConfiguration
						"fixed_size_chunking_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MaxTokens
								"max_tokens": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The maximum number of tokens to include in a chunk.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: OverlapPercentage
								"overlap_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The percentage of overlap between adjacent chunks of a data source.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Configurations for when you choose fixed-size chunking. If you set the chunkingStrategy as NONE, exclude this field.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Bedrock::DataSource",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Bedrock::DataSource").WithTerraformTypeName("awscc_bedrock_data_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_arn":                           "BucketArn",
		"chunking_configuration":               "ChunkingConfiguration",
		"chunking_strategy":                    "ChunkingStrategy",
		"created_at":                           "CreatedAt",
		"data_source_configuration":            "DataSourceConfiguration",
		"data_source_id":                       "DataSourceId",
		"data_source_status":                   "DataSourceStatus",
		"description":                          "Description",
		"fixed_size_chunking_configuration":    "FixedSizeChunkingConfiguration",
		"inclusion_prefixes":                   "InclusionPrefixes",
		"kms_key_arn":                          "KmsKeyArn",
		"knowledge_base_id":                    "KnowledgeBaseId",
		"max_tokens":                           "MaxTokens",
		"name":                                 "Name",
		"overlap_percentage":                   "OverlapPercentage",
		"s3_configuration":                     "S3Configuration",
		"server_side_encryption_configuration": "ServerSideEncryptionConfiguration",
		"type":                                 "Type",
		"updated_at":                           "UpdatedAt",
		"vector_ingestion_configuration":       "VectorIngestionConfiguration",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
