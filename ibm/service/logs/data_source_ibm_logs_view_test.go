// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsViewDataSourceBasic(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewDataSourceConfigBasic(viewName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "logs_view_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "time_selection.#"),
				),
			},
		},
	})
}

func TestAccIbmLogsViewDataSourceAllArgs(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	viewTier := "priority_insights"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewDataSourceConfig(viewName, viewTier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "logs_view_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "search_query.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "time_selection.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "folder_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view.logs_view_instance", "tier"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsViewDataSourceConfigBasic(viewName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_view" "logs_view_instance" {
			instance_id = "%s"
			region      = "%s"
			name = "%s"
			time_selection {
				quick_selection {
					caption = "Last 1 hour"
					seconds = 3600
				}
			}
			tier = "priority_insights"
			search_query {
				query       = ""
				syntax_type = "lucene"
			}
		}

		data "ibm_logs_view" "logs_view_instance" {
			instance_id  = ibm_logs_view.logs_view_instance.instance_id
			region       = ibm_logs_view.logs_view_instance.region
			logs_view_id = ibm_logs_view.logs_view_instance.view_id
		}
`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName)
	}

func testAccCheckIbmLogsViewDataSourceConfig(viewName string, viewTier string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_view" "logs_view_instance" {
			instance_id = "%s"
			region      = "%s"
			name = "%s"
			search_query {
				query = "error"
				syntax_type = "dataprime"
			}
			time_selection {
				quick_selection {
					caption = "Last 1 hour"
					seconds = 3600
				}
			}
			filters {
				filters {
					name = "applicationName"
					selected_values = {"cs-rest-test1":true,"demo":true}
				}
			}
			
			tier = "%s"
		}

		data "ibm_logs_view" "logs_view_instance" {
			instance_id  = ibm_logs_view.logs_view_instance.instance_id
			region       = ibm_logs_view.logs_view_instance.region
			logs_view_id = ibm_logs_view.logs_view_instance.view_id
		}
`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName, viewTier)
	}