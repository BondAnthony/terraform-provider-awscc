// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_acmpca_certificate_authority", certificateAuthorityDataSource)
}

// certificateAuthorityDataSource returns the Terraform awscc_acmpca_certificate_authority data source.
// This Terraform data source corresponds to the CloudFormation AWS::ACMPCA::CertificateAuthority resource.
func certificateAuthorityDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the certificate authority.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the certificate authority.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateSigningRequest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.",
		//	  "type": "string"
		//	}
		"certificate_signing_request": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CsrExtensions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.",
		//	  "properties": {
		//	    "KeyUsage": {
		//	      "additionalProperties": false,
		//	      "description": "Structure that contains X.509 KeyUsage information.",
		//	      "properties": {
		//	        "CRLSign": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "DataEncipherment": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "DecipherOnly": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "DigitalSignature": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "EncipherOnly": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "KeyAgreement": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "KeyCertSign": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "KeyEncipherment": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        },
		//	        "NonRepudiation": {
		//	          "default": false,
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "SubjectInformationAccess": {
		//	      "description": "Array of X.509 AccessDescription.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Structure that contains X.509 AccessDescription information.",
		//	        "properties": {
		//	          "AccessLocation": {
		//	            "additionalProperties": false,
		//	            "description": "Structure that contains X.509 GeneralName information. Assign one and ONLY one field.",
		//	            "properties": {
		//	              "DirectoryName": {
		//	                "additionalProperties": false,
		//	                "description": "Structure that contains X.500 distinguished name information for your CA.",
		//	                "properties": {
		//	                  "CommonName": {
		//	                    "type": "string"
		//	                  },
		//	                  "Country": {
		//	                    "type": "string"
		//	                  },
		//	                  "CustomAttributes": {
		//	                    "description": "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
		//	                    "items": {
		//	                      "additionalProperties": false,
		//	                      "description": "Structure that contains X.500 attribute type and value.",
		//	                      "properties": {
		//	                        "ObjectIdentifier": {
		//	                          "description": "String that contains X.509 ObjectIdentifier information.",
		//	                          "type": "string"
		//	                        },
		//	                        "Value": {
		//	                          "type": "string"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "ObjectIdentifier",
		//	                        "Value"
		//	                      ],
		//	                      "type": "object"
		//	                    },
		//	                    "type": "array"
		//	                  },
		//	                  "DistinguishedNameQualifier": {
		//	                    "type": "string"
		//	                  },
		//	                  "GenerationQualifier": {
		//	                    "type": "string"
		//	                  },
		//	                  "GivenName": {
		//	                    "type": "string"
		//	                  },
		//	                  "Initials": {
		//	                    "type": "string"
		//	                  },
		//	                  "Locality": {
		//	                    "type": "string"
		//	                  },
		//	                  "Organization": {
		//	                    "type": "string"
		//	                  },
		//	                  "OrganizationalUnit": {
		//	                    "type": "string"
		//	                  },
		//	                  "Pseudonym": {
		//	                    "type": "string"
		//	                  },
		//	                  "SerialNumber": {
		//	                    "type": "string"
		//	                  },
		//	                  "State": {
		//	                    "type": "string"
		//	                  },
		//	                  "Surname": {
		//	                    "type": "string"
		//	                  },
		//	                  "Title": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "DnsName": {
		//	                "description": "String that contains X.509 DnsName information.",
		//	                "type": "string"
		//	              },
		//	              "EdiPartyName": {
		//	                "additionalProperties": false,
		//	                "description": "Structure that contains X.509 EdiPartyName information.",
		//	                "properties": {
		//	                  "NameAssigner": {
		//	                    "type": "string"
		//	                  },
		//	                  "PartyName": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "PartyName"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "IpAddress": {
		//	                "description": "String that contains X.509 IpAddress information.",
		//	                "type": "string"
		//	              },
		//	              "OtherName": {
		//	                "additionalProperties": false,
		//	                "description": "Structure that contains X.509 OtherName information.",
		//	                "properties": {
		//	                  "TypeId": {
		//	                    "description": "String that contains X.509 ObjectIdentifier information.",
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "TypeId",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "RegisteredId": {
		//	                "description": "String that contains X.509 ObjectIdentifier information.",
		//	                "type": "string"
		//	              },
		//	              "Rfc822Name": {
		//	                "description": "String that contains X.509 Rfc822Name information.",
		//	                "type": "string"
		//	              },
		//	              "UniformResourceIdentifier": {
		//	                "description": "String that contains X.509 UniformResourceIdentifier information.",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "AccessMethod": {
		//	            "additionalProperties": false,
		//	            "description": "Structure that contains X.509 AccessMethod information. Assign one and ONLY one field.",
		//	            "properties": {
		//	              "AccessMethodType": {
		//	                "description": "Pre-defined enum string for X.509 AccessMethod ObjectIdentifiers.",
		//	                "type": "string"
		//	              },
		//	              "CustomObjectIdentifier": {
		//	                "description": "String that contains X.509 ObjectIdentifier information.",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "AccessMethod",
		//	          "AccessLocation"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"csr_extensions": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KeyUsage
				"key_usage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CRLSign
						"crl_sign": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DataEncipherment
						"data_encipherment": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DecipherOnly
						"decipher_only": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DigitalSignature
						"digital_signature": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: EncipherOnly
						"encipher_only": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: KeyAgreement
						"key_agreement": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: KeyCertSign
						"key_cert_sign": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: KeyEncipherment
						"key_encipherment": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: NonRepudiation
						"non_repudiation": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Structure that contains X.509 KeyUsage information.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SubjectInformationAccess
				"subject_information_access": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AccessLocation
							"access_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DirectoryName
									"directory_name": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: CommonName
											"common_name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Country
											"country": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: CustomAttributes
											"custom_attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
												NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: ObjectIdentifier
														"object_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
															Description: "String that contains X.509 ObjectIdentifier information.",
															Computed:    true,
														}, /*END ATTRIBUTE*/
														// Property: Value
														"value": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
												}, /*END NESTED OBJECT*/
												Description: "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: DistinguishedNameQualifier
											"distinguished_name_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: GenerationQualifier
											"generation_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: GivenName
											"given_name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Initials
											"initials": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Locality
											"locality": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Organization
											"organization": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: OrganizationalUnit
											"organizational_unit": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Pseudonym
											"pseudonym": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: SerialNumber
											"serial_number": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: State
											"state": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Surname
											"surname": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Title
											"title": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Structure that contains X.500 distinguished name information for your CA.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DnsName
									"dns_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "String that contains X.509 DnsName information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: EdiPartyName
									"edi_party_name": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: NameAssigner
											"name_assigner": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: PartyName
											"party_name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Structure that contains X.509 EdiPartyName information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: IpAddress
									"ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "String that contains X.509 IpAddress information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: OtherName
									"other_name": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: TypeId
											"type_id": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "String that contains X.509 ObjectIdentifier information.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: Value
											"value": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Structure that contains X.509 OtherName information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: RegisteredId
									"registered_id": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "String that contains X.509 ObjectIdentifier information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Rfc822Name
									"rfc_822_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "String that contains X.509 Rfc822Name information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: UniformResourceIdentifier
									"uniform_resource_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "String that contains X.509 UniformResourceIdentifier information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Structure that contains X.509 GeneralName information. Assign one and ONLY one field.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: AccessMethod
							"access_method": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: AccessMethodType
									"access_method_type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Pre-defined enum string for X.509 AccessMethod ObjectIdentifiers.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: CustomObjectIdentifier
									"custom_object_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "String that contains X.509 ObjectIdentifier information.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Structure that contains X.509 AccessMethod information. Assign one and ONLY one field.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Array of X.509 AccessDescription.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Structure that contains CSR pass through extension information used by the CreateCertificateAuthority action.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyAlgorithm
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.",
		//	  "type": "string"
		//	}
		"key_algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyStorageSecurityStandard
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.",
		//	  "type": "string"
		//	}
		"key_storage_security_standard": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RevocationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.",
		//	  "properties": {
		//	    "CrlConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.",
		//	      "properties": {
		//	        "CrlDistributionPointExtensionConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "Configures the default behavior of the CRL Distribution Point extension for certificates issued by your certificate authority",
		//	          "properties": {
		//	            "OmitExtension": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "OmitExtension"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "CrlType": {
		//	          "type": "string"
		//	        },
		//	        "CustomCname": {
		//	          "type": "string"
		//	        },
		//	        "CustomPath": {
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "type": "boolean"
		//	        },
		//	        "ExpirationInDays": {
		//	          "type": "integer"
		//	        },
		//	        "S3BucketName": {
		//	          "type": "string"
		//	        },
		//	        "S3ObjectAcl": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Enabled"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "OcspConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Helps to configure online certificate status protocol (OCSP) responder for your certificate authority",
		//	      "properties": {
		//	        "Enabled": {
		//	          "type": "boolean"
		//	        },
		//	        "OcspCustomCname": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Enabled"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"revocation_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CrlConfiguration
				"crl_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CrlDistributionPointExtensionConfiguration
						"crl_distribution_point_extension_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: OmitExtension
								"omit_extension": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Configures the default behavior of the CRL Distribution Point extension for certificates issued by your certificate authority",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: CrlType
						"crl_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: CustomCname
						"custom_cname": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: CustomPath
						"custom_path": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ExpirationInDays
						"expiration_in_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: S3BucketName
						"s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: S3ObjectAcl
						"s3_object_acl": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: OcspConfiguration
				"ocsp_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: OcspCustomCname
						"ocsp_custom_cname": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Helps to configure online certificate status protocol (OCSP) responder for your certificate authority",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Certificate revocation information used by the CreateCertificateAuthority and UpdateCertificateAuthority actions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SigningAlgorithm
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Algorithm your CA uses to sign certificate requests.",
		//	  "type": "string"
		//	}
		"signing_algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Algorithm your CA uses to sign certificate requests.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subject
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Structure that contains X.500 distinguished name information for your CA.",
		//	  "properties": {
		//	    "CommonName": {
		//	      "type": "string"
		//	    },
		//	    "Country": {
		//	      "type": "string"
		//	    },
		//	    "CustomAttributes": {
		//	      "description": "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Structure that contains X.500 attribute type and value.",
		//	        "properties": {
		//	          "ObjectIdentifier": {
		//	            "description": "String that contains X.509 ObjectIdentifier information.",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "ObjectIdentifier",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "DistinguishedNameQualifier": {
		//	      "type": "string"
		//	    },
		//	    "GenerationQualifier": {
		//	      "type": "string"
		//	    },
		//	    "GivenName": {
		//	      "type": "string"
		//	    },
		//	    "Initials": {
		//	      "type": "string"
		//	    },
		//	    "Locality": {
		//	      "type": "string"
		//	    },
		//	    "Organization": {
		//	      "type": "string"
		//	    },
		//	    "OrganizationalUnit": {
		//	      "type": "string"
		//	    },
		//	    "Pseudonym": {
		//	      "type": "string"
		//	    },
		//	    "SerialNumber": {
		//	      "type": "string"
		//	    },
		//	    "State": {
		//	      "type": "string"
		//	    },
		//	    "Surname": {
		//	      "type": "string"
		//	    },
		//	    "Title": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"subject": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CommonName
				"common_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Country
				"country": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: CustomAttributes
				"custom_attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ObjectIdentifier
							"object_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "String that contains X.509 ObjectIdentifier information.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Array of X.500 attribute type and value. CustomAttributes cannot be used along with pre-defined attributes.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DistinguishedNameQualifier
				"distinguished_name_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: GenerationQualifier
				"generation_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: GivenName
				"given_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Initials
				"initials": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Locality
				"locality": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Organization
				"organization": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OrganizationalUnit
				"organizational_unit": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Pseudonym
				"pseudonym": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SerialNumber
				"serial_number": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: State
				"state": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Surname
				"surname": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Title
				"title": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Structure that contains X.500 distinguished name information for your CA.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the certificate authority.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the certificate authority.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UsageMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Usage mode of the ceritificate authority.",
		//	  "type": "string"
		//	}
		"usage_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Usage mode of the ceritificate authority.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ACMPCA::CertificateAuthority",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ACMPCA::CertificateAuthority").WithTerraformTypeName("awscc_acmpca_certificate_authority")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_location":             "AccessLocation",
		"access_method":               "AccessMethod",
		"access_method_type":          "AccessMethodType",
		"arn":                         "Arn",
		"certificate_signing_request": "CertificateSigningRequest",
		"common_name":                 "CommonName",
		"country":                     "Country",
		"crl_configuration":           "CrlConfiguration",
		"crl_distribution_point_extension_configuration": "CrlDistributionPointExtensionConfiguration",
		"crl_sign":                      "CRLSign",
		"crl_type":                      "CrlType",
		"csr_extensions":                "CsrExtensions",
		"custom_attributes":             "CustomAttributes",
		"custom_cname":                  "CustomCname",
		"custom_object_identifier":      "CustomObjectIdentifier",
		"custom_path":                   "CustomPath",
		"data_encipherment":             "DataEncipherment",
		"decipher_only":                 "DecipherOnly",
		"digital_signature":             "DigitalSignature",
		"directory_name":                "DirectoryName",
		"distinguished_name_qualifier":  "DistinguishedNameQualifier",
		"dns_name":                      "DnsName",
		"edi_party_name":                "EdiPartyName",
		"enabled":                       "Enabled",
		"encipher_only":                 "EncipherOnly",
		"expiration_in_days":            "ExpirationInDays",
		"generation_qualifier":          "GenerationQualifier",
		"given_name":                    "GivenName",
		"initials":                      "Initials",
		"ip_address":                    "IpAddress",
		"key":                           "Key",
		"key_agreement":                 "KeyAgreement",
		"key_algorithm":                 "KeyAlgorithm",
		"key_cert_sign":                 "KeyCertSign",
		"key_encipherment":              "KeyEncipherment",
		"key_storage_security_standard": "KeyStorageSecurityStandard",
		"key_usage":                     "KeyUsage",
		"locality":                      "Locality",
		"name_assigner":                 "NameAssigner",
		"non_repudiation":               "NonRepudiation",
		"object_identifier":             "ObjectIdentifier",
		"ocsp_configuration":            "OcspConfiguration",
		"ocsp_custom_cname":             "OcspCustomCname",
		"omit_extension":                "OmitExtension",
		"organization":                  "Organization",
		"organizational_unit":           "OrganizationalUnit",
		"other_name":                    "OtherName",
		"party_name":                    "PartyName",
		"pseudonym":                     "Pseudonym",
		"registered_id":                 "RegisteredId",
		"revocation_configuration":      "RevocationConfiguration",
		"rfc_822_name":                  "Rfc822Name",
		"s3_bucket_name":                "S3BucketName",
		"s3_object_acl":                 "S3ObjectAcl",
		"serial_number":                 "SerialNumber",
		"signing_algorithm":             "SigningAlgorithm",
		"state":                         "State",
		"subject":                       "Subject",
		"subject_information_access":    "SubjectInformationAccess",
		"surname":                       "Surname",
		"tags":                          "Tags",
		"title":                         "Title",
		"type":                          "Type",
		"type_id":                       "TypeId",
		"uniform_resource_identifier":   "UniformResourceIdentifier",
		"usage_mode":                    "UsageMode",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
