// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package securitylake

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_securitylake_data_lake", dataLakeDataSource)
}

// dataLakeDataSource returns the Terraform awscc_securitylake_data_lake data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecurityLake::DataLake resource.
func dataLakeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) created by you to provide to the subscriber.",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) created by you to provide to the subscriber.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Provides encryption details of Amazon Security Lake object.",
		//	  "properties": {
		//	    "KmsKeyId": {
		//	      "description": "The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Provides encryption details of Amazon Security Lake object.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LifecycleConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Provides lifecycle details of Amazon Security Lake object.",
		//	  "properties": {
		//	    "Expiration": {
		//	      "additionalProperties": false,
		//	      "description": "Provides data expiration details of Amazon Security Lake object.",
		//	      "properties": {
		//	        "Days": {
		//	          "description": "Number of days before data expires in the Amazon Security Lake object.",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Transitions": {
		//	      "description": "Provides data storage transition details of Amazon Security Lake object.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Days": {
		//	            "description": "Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.",
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "StorageClass": {
		//	            "description": "The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lifecycle_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Expiration
				"expiration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Days
						"days": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Number of days before data expires in the Amazon Security Lake object.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Provides data expiration details of Amazon Security Lake object.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Transitions
				"transitions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Days
							"days": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: StorageClass
							"storage_class": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Provides data storage transition details of Amazon Security Lake object.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Provides lifecycle details of Amazon Security Lake object.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MetaStoreManagerRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) used to index AWS Glue table partitions that are generated by the ingestion and normalization of AWS log sources and custom sources.",
		//	  "pattern": "^arn:.*$",
		//	  "type": "string"
		//	}
		"meta_store_manager_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) used to index AWS Glue table partitions that are generated by the ingestion and normalization of AWS log sources and custom sources.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReplicationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Provides replication details of Amazon Security Lake object.",
		//	  "properties": {
		//	    "Regions": {
		//	      "description": "Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. Amazon S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different AWS Regions or within the same Region as the source bucket.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "pattern": "^(af|ap|ca|eu|me|sa|us)-(central|north|(north(?:east|west))|south|south(?:east|west)|east|west)-\\d+$",
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "RoleArn": {
		//	      "description": "Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.",
		//	      "pattern": "^arn:.*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"replication_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Regions
				"regions": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. Amazon S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different AWS Regions or within the same Region as the source bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Provides replication details of Amazon Security Lake object.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: S3BucketArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN for the Amazon Security Lake Amazon S3 bucket.",
		//	  "type": "string"
		//	}
		"s3_bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN for the Amazon Security Lake Amazon S3 bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 characters in length.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 characters in length.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecurityLake::DataLake",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecurityLake::DataLake").WithTerraformTypeName("awscc_securitylake_data_lake")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"days":                        "Days",
		"encryption_configuration":    "EncryptionConfiguration",
		"expiration":                  "Expiration",
		"key":                         "Key",
		"kms_key_id":                  "KmsKeyId",
		"lifecycle_configuration":     "LifecycleConfiguration",
		"meta_store_manager_role_arn": "MetaStoreManagerRoleArn",
		"regions":                     "Regions",
		"replication_configuration":   "ReplicationConfiguration",
		"role_arn":                    "RoleArn",
		"s3_bucket_arn":               "S3BucketArn",
		"storage_class":               "StorageClass",
		"tags":                        "Tags",
		"transitions":                 "Transitions",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
