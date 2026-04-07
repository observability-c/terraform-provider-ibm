// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsViewBasic(t *testing.T) {
	var conf logsv0.View
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsViewDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsViewExists("ibm_logs_view.logs_view_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsViewConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsViewAllArgs(t *testing.T) {
	var conf logsv0.View
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	tier := "priority_insights"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	tierUpdate := "all_logs_templates"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsViewDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewConfig(name, tier),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsViewExists("ibm_logs_view.logs_view_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "tier", tier),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsViewConfig(nameUpdate, tierUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_view.logs_view_instance", "tier", tierUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_view.logs_view",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsViewConfigBasic(name string) string {
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
`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
	}

func testAccCheckIbmLogsViewConfig(name string, tier string) string {
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
`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, tier)
	}

func testAccCheckIbmLogsViewExists(n string, obj logsv0.View) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		viewId, err := strconv.ParseInt(resourceID[2], 10, 64)
		if err != nil {
			return fmt.Errorf("Error parsing view ID %s: %s", resourceID[2], err)
		}

		getViewOptions := &logsv0.GetViewOptions{}
		getViewOptions.SetID(viewId)

		view, _, err := logsClient.GetView(getViewOptions)
		if err != nil {
			return err
		}

		obj = *view
		return nil
	}
}

func testAccCheckIbmLogsViewDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_view" {
			continue
		}

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		viewId, err := strconv.ParseInt(resourceID[2], 10, 64)
		if err != nil {
			return fmt.Errorf("Error parsing view ID %s: %s", resourceID[2], err)
		}

		getViewOptions := &logsv0.GetViewOptions{}
		getViewOptions.SetID(viewId)

		// Try to find the key
		_, response, err := logsClient.GetView(getViewOptions)

		if err == nil {
			return fmt.Errorf("logs_view still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_view (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
