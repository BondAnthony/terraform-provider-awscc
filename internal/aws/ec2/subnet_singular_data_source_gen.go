// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_subnet", subnetDataSource)
}

// subnetDataSource returns the Terraform awscc_ec2_subnet data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::Subnet resource.
func subnetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssignIpv6AddressOnCreation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
		//	  "type": "boolean"
		//	}
		"assign_ipv_6_address_on_creation": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone of the subnet.\n If you update this property, you must also update the ``CidrBlock`` property.",
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone of the subnet.\n If you update this property, you must also update the ``CidrBlock`` property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZoneId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AZ ID of the subnet.",
		//	  "type": "string"
		//	}
		"availability_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AZ ID of the subnet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv4 CIDR block assigned to the subnet.\n If you update this property, we create a new subnet, and then delete the existing one.",
		//	  "type": "string"
		//	}
		"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv4 CIDR block assigned to the subnet.\n If you update this property, we create a new subnet, and then delete the existing one.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableDns64
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. For more information, see [DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-nat64-dns64) in the *User Guide*.",
		//	  "type": "boolean"
		//	}
		"enable_dns_64": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations. For more information, see [DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-nat64-dns64) in the *User Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableLniAtDeviceIndex
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).",
		//	  "type": "integer"
		//	}
		"enable_lni_at_device_index": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv4IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv4 IPAM pool ID for the subnet.",
		//	  "type": "string"
		//	}
		"ipv_4_ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An IPv4 IPAM pool ID for the subnet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv4NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv4 netmask length for the subnet.",
		//	  "type": "integer"
		//	}
		"ipv_4_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "An IPv4 netmask length for the subnet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv6 CIDR block.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
		//	  "type": "string"
		//	}
		"ipv_6_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv6 CIDR block.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlocks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv6 network ranges for the subnet, in CIDR notation.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"ipv_6_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IPv6 network ranges for the subnet, in CIDR notation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv6 IPAM pool ID for the subnet.",
		//	  "type": "string"
		//	}
		"ipv_6_ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An IPv6 IPAM pool ID for the subnet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6Native
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.",
		//	  "type": "boolean"
		//	}
		"ipv_6_native": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Ipv6NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv6 netmask length for the subnet.",
		//	  "type": "integer"
		//	}
		"ipv_6_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "An IPv6 netmask length for the subnet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MapPublicIpOnLaunch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``.\n  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).",
		//	  "type": "boolean"
		//	}
		"map_public_ip_on_launch": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``.\n  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkAclAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"network_acl_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Outpost.",
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Outpost.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PrivateDnsNameOptionsOnLaunch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.\n Available options:\n  +  EnableResourceNameDnsAAAARecord (true | false)\n  +  EnableResourceNameDnsARecord (true | false)\n  +  HostnameType (ip-name | resource-name)",
		//	  "properties": {
		//	    "EnableResourceNameDnsAAAARecord": {
		//	      "type": "boolean"
		//	    },
		//	    "EnableResourceNameDnsARecord": {
		//	      "type": "boolean"
		//	    },
		//	    "HostnameType": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"private_dns_name_options_on_launch": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableResourceNameDnsAAAARecord
				"enable_resource_name_dns_aaaa_record": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EnableResourceNameDnsARecord
				"enable_resource_name_dns_a_record": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: HostnameType
				"hostname_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.\n Available options:\n  +  EnableResourceNameDnsAAAARecord (true | false)\n  +  EnableResourceNameDnsARecord (true | false)\n  +  HostnameType (ip-name | resource-name)",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the subnet.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a tag. For more information, see [Add tags to a resource](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#cloudformation-add-tag-specifications).",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Description: "The tag key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the subnet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC the subnet is in.\n If you update this property, you must also update the ``CidrBlock`` property.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC the subnet is in.\n If you update this property, you must also update the ``CidrBlock`` property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::Subnet",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Subnet").WithTerraformTypeName("awscc_ec2_subnet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assign_ipv_6_address_on_creation":     "AssignIpv6AddressOnCreation",
		"availability_zone":                    "AvailabilityZone",
		"availability_zone_id":                 "AvailabilityZoneId",
		"cidr_block":                           "CidrBlock",
		"enable_dns_64":                        "EnableDns64",
		"enable_lni_at_device_index":           "EnableLniAtDeviceIndex",
		"enable_resource_name_dns_a_record":    "EnableResourceNameDnsARecord",
		"enable_resource_name_dns_aaaa_record": "EnableResourceNameDnsAAAARecord",
		"hostname_type":                        "HostnameType",
		"ipv_4_ipam_pool_id":                   "Ipv4IpamPoolId",
		"ipv_4_netmask_length":                 "Ipv4NetmaskLength",
		"ipv_6_cidr_block":                     "Ipv6CidrBlock",
		"ipv_6_cidr_blocks":                    "Ipv6CidrBlocks",
		"ipv_6_ipam_pool_id":                   "Ipv6IpamPoolId",
		"ipv_6_native":                         "Ipv6Native",
		"ipv_6_netmask_length":                 "Ipv6NetmaskLength",
		"key":                                  "Key",
		"map_public_ip_on_launch":              "MapPublicIpOnLaunch",
		"network_acl_association_id":           "NetworkAclAssociationId",
		"outpost_arn":                          "OutpostArn",
		"private_dns_name_options_on_launch":   "PrivateDnsNameOptionsOnLaunch",
		"subnet_id":                            "SubnetId",
		"tags":                                 "Tags",
		"value":                                "Value",
		"vpc_id":                               "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
