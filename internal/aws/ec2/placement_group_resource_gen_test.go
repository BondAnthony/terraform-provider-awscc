// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEC2PlacementGroup_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::PlacementGroup", "awscc_ec2_placement_group", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSEC2PlacementGroup_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EC2::PlacementGroup", "awscc_ec2_placement_group", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}