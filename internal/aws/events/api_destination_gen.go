// Code generated by generators/resource/main.go; DO NOT EDIT.

package events

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_events_api_destination", apiDestinationResourceType)
}

// apiDestinationResourceType returns the Terraform aws_events_api_destination resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Events::ApiDestination resource type.
func apiDestinationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The arn of the api destination.",
			     "type": "string"
			   }
			*/
			Description: "The arn of the api destination.",
			Type:        types.StringType,
			Computed:    true,
		},
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The arn of the connection.",
			     "type": "string"
			   }
			*/
			Description: "The arn of the connection.",
			Type:        types.StringType,
			Required:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 512,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"http_method": {
			// Property: HttpMethod
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "GET",
			       "HEAD",
			       "POST",
			       "OPTIONS",
			       "PUT",
			       "DELETE",
			       "PATCH"
			     ],
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"invocation_endpoint": {
			// Property: InvocationEndpoint
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Url endpoint to invoke.",
			     "type": "string"
			   }
			*/
			Description: "Url endpoint to invoke.",
			Type:        types.StringType,
			Required:    true,
		},
		"invocation_rate_limit_per_second": {
			// Property: InvocationRateLimitPerSecond
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "integer"
			   }
			*/
			Type:     types.NumberType,
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Name of the apiDestination.",
			     "maxLength": 64,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Name of the apiDestination.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::Events::ApiDestination.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Events::ApiDestination").WithTerraformTypeName("aws_events_api_destination").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_events_api_destination", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}