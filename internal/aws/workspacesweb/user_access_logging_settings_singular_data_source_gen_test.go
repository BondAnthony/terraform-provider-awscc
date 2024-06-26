// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package workspacesweb_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSWorkSpacesWebUserAccessLoggingSettingsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::WorkSpacesWeb::UserAccessLoggingSettings", "awscc_workspacesweb_user_access_logging_settings", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSWorkSpacesWebUserAccessLoggingSettingsDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::WorkSpacesWeb::UserAccessLoggingSettings", "awscc_workspacesweb_user_access_logging_settings", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
