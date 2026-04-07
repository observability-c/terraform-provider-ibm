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

func TestAccIbmLogsViewsDataSourceBasic(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewsDataSourceConfigBasic(viewName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmLogsViewsDataSourceAllArgs(t *testing.T) {
	viewName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	viewTier := "priority_insights"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewsDataSourceConfig(viewName, viewTier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "views.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "views.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_views.logs_views_instance", "views.0.name", viewName),
					resource.TestCheckResourceAttrSet("data.ibm_logs_views.logs_views_instance", "views.0.folder_id"),
					resource.TestCheckResourceAttr("data.ibm_logs_views.logs_views_instance", "views.0.tier", viewTier),
				),
			},
		},
	})
}

func testAccCheckIbmLogsViewsDataSourceConfigBasic(viewName string) string {
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

		data "ibm_logs_views" "logs_views_instance" {
			instance_id = "%s"
			region      = "%s"
			depends_on = [
				ibm_logs_view.logs_view_instance
			]
		}
`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName, acc.LogsInstanceId, acc.LogsInstanceRegion)
	}

func testAccCheckIbmLogsViewsDataSourceConfig(viewName string, viewTier string) string {
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

		data "ibm_logs_views" "logs_views_instance" {
			instance_id = "%s"
			region      = "%s"
			depends_on = [
				ibm_logs_view.logs_view_instance
			]
		}
`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewName, acc.LogsInstanceId, acc.LogsInstanceRegion, viewTier)
	}