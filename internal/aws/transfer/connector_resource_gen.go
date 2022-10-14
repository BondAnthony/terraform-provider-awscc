// Code generated by generators/resource/main.go; DO NOT EDIT.

package transfer

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_transfer_connector", connectorResource)
}

// connectorResource returns the Terraform awscc_transfer_connector resource.
// This Terraform resource corresponds to the CloudFormation AWS::Transfer::Connector resource.
func connectorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_role": {
			// Property: AccessRole
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the access role for the connector.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:.*role/.*",
			//   "type": "string"
			// }
			Description: "Specifies the access role for the connector.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("arn:.*role/.*"), ""),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
			//   "maxLength": 1600,
			//   "minLength": 20,
			//   "pattern": "arn:.*",
			//   "type": "string"
			// }
			Description: "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"as_2_config": {
			// Property: As2Config
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration for an AS2 connector.",
			//   "properties": {
			//     "Compression": {
			//       "description": "Compression setting for this AS2 connector configuration.",
			//       "enum": [
			//         "ZLIB",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     },
			//     "EncryptionAlgorithm": {
			//       "description": "Encryption algorithm for this AS2 connector configuration.",
			//       "enum": [
			//         "AES128_CBC",
			//         "AES192_CBC",
			//         "AES256_CBC"
			//       ],
			//       "type": "string"
			//     },
			//     "LocalProfileId": {
			//       "description": "A unique identifier for the local profile.",
			//       "maxLength": 19,
			//       "minLength": 19,
			//       "pattern": "^p-([0-9a-f]{17})$",
			//       "type": "string"
			//     },
			//     "MdnResponse": {
			//       "description": "MDN Response setting for this AS2 connector configuration.",
			//       "enum": [
			//         "SYNC",
			//         "NONE"
			//       ],
			//       "type": "string"
			//     },
			//     "MdnSigningAlgorithm": {
			//       "description": "MDN Signing algorithm for this AS2 connector configuration.",
			//       "enum": [
			//         "SHA256",
			//         "SHA384",
			//         "SHA512",
			//         "SHA1",
			//         "NONE",
			//         "DEFAULT"
			//       ],
			//       "type": "string"
			//     },
			//     "MessageSubject": {
			//       "description": "The message subject for this AS2 connector configuration.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "PartnerProfileId": {
			//       "description": "A unique identifier for the partner profile.",
			//       "maxLength": 19,
			//       "minLength": 19,
			//       "pattern": "^p-([0-9a-f]{17})$",
			//       "type": "string"
			//     },
			//     "SigningAlgorithm": {
			//       "description": "Signing algorithm for this AS2 connector configuration.",
			//       "enum": [
			//         "SHA256",
			//         "SHA384",
			//         "SHA512",
			//         "SHA1",
			//         "NONE"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Configuration for an AS2 connector.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"compression": {
						// Property: Compression
						Description: "Compression setting for this AS2 connector configuration.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ZLIB",
								"DISABLED",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"encryption_algorithm": {
						// Property: EncryptionAlgorithm
						Description: "Encryption algorithm for this AS2 connector configuration.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"AES128_CBC",
								"AES192_CBC",
								"AES256_CBC",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"local_profile_id": {
						// Property: LocalProfileId
						Description: "A unique identifier for the local profile.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(19, 19),
							validate.StringMatch(regexp.MustCompile("^p-([0-9a-f]{17})$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"mdn_response": {
						// Property: MdnResponse
						Description: "MDN Response setting for this AS2 connector configuration.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SYNC",
								"NONE",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"mdn_signing_algorithm": {
						// Property: MdnSigningAlgorithm
						Description: "MDN Signing algorithm for this AS2 connector configuration.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SHA256",
								"SHA384",
								"SHA512",
								"SHA1",
								"NONE",
								"DEFAULT",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"message_subject": {
						// Property: MessageSubject
						Description: "The message subject for this AS2 connector configuration.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1024),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"partner_profile_id": {
						// Property: PartnerProfileId
						Description: "A unique identifier for the partner profile.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(19, 19),
							validate.StringMatch(regexp.MustCompile("^p-([0-9a-f]{17})$"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"signing_algorithm": {
						// Property: SigningAlgorithm
						Description: "Signing algorithm for this AS2 connector configuration.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"SHA256",
								"SHA384",
								"SHA512",
								"SHA1",
								"NONE",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"connector_id": {
			// Property: ConnectorId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the connector.",
			//   "maxLength": 19,
			//   "minLength": 19,
			//   "pattern": "^c-([0-9a-f]{17})$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the connector.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"logging_role": {
			// Property: LoggingRole
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the logging role for the connector.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:.*role/.*",
			//   "type": "string"
			// }
			Description: "Specifies the logging role for the connector.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("arn:.*role/.*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Creates a key-value pair for a specific resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The name assigned to the tag that you create.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Contains one or more values that you assigned to the key name you create.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The name assigned to the tag that you create.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "Contains one or more values that you assigned to the key name you create.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			// {
			//   "description": "URL for Connector",
			//   "maxLength": 255,
			//   "type": "string"
			// }
			Description: "URL for Connector",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(255),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Transfer::Connector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Connector").WithTerraformTypeName("awscc_transfer_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_role":           "AccessRole",
		"arn":                   "Arn",
		"as_2_config":           "As2Config",
		"compression":           "Compression",
		"connector_id":          "ConnectorId",
		"encryption_algorithm":  "EncryptionAlgorithm",
		"key":                   "Key",
		"local_profile_id":      "LocalProfileId",
		"logging_role":          "LoggingRole",
		"mdn_response":          "MdnResponse",
		"mdn_signing_algorithm": "MdnSigningAlgorithm",
		"message_subject":       "MessageSubject",
		"partner_profile_id":    "PartnerProfileId",
		"signing_algorithm":     "SigningAlgorithm",
		"tags":                  "Tags",
		"url":                   "Url",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}