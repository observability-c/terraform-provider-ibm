// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsAlertDefinitionsDataSourceBasic(t *testing.T) {
	alertName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionsDataSourceConfigBasic(alertName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertDefinitionsDataSourceAllArgs(t *testing.T) {
	alertDefinitionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDefinitionDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	alertDefinitionEnabled := "false"
	alertDefinitionPriority := "p3"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionsDataSourceConfig(alertDefinitionName, alertDefinitionDescription, alertDefinitionEnabled, alertDefinitionPriority),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.alert_version_id"),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.name", alertDefinitionName),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.description", alertDefinitionDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.enabled", alertDefinitionEnabled),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.priority", alertDefinitionPriority),
				),
			},
		},
	})
}

func testAccCheckIbmLogsAlertDefinitionsDataSourceConfigBasic(alertName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			instance_id      = "%s"
			region           = "%s"
			name             = "%s"
			deleted          = false
			enabled          = false
			entity_labels    = {}
			group_by_keys    = []
			phantom_mode     = false
			priority         = "p3"
			type             = "logs_threshold"
			incidents_settings {
				minutes   = 1
				notify_on = "triggered_only_unspecified"
			}

			logs_threshold {
				logs_filter {
					simple_filter {
						lucene_query = "\"push\""

						label_filters {
							severities = []

							application_name {
								operation = "is_or_unspecified"
								value     = "sev4"
							}

							subsystem_name {
								operation = "is_or_unspecified"
								value     = "sev4-logs"
							}
						}
					}
				}

				notification_payload_filter = []

				rules {
					condition {
						condition_type = "more_than_or_unspecified"
						threshold      = 1

						time_window {
							logs_time_window_specific_value = "minutes_10"
						}
					}
					override {
						priority = "p3"
					}
				}
				rules {
					condition {
						condition_type = "more_than_or_unspecified"
						threshold      = 1

						time_window {
							logs_time_window_specific_value = "minutes_5_or_unspecified"
						}
					}
					override {
						priority = "p4"
					}
				}
			}
        }

		data "ibm_logs_alert_definitions" "logs_alert_definitions_instance" {
			depends_on = [
				ibm_logs_alert_definition.logs_alert_definition_instance
			]
			instance_id = ibm_logs_alert_definition.logs_alert_definition_instance.instance_id
			region 		= ibm_logs_alert_definition.logs_alert_definition_instance.region
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, alertName)
}

func testAccCheckIbmLogsAlertDefinitionsDataSourceConfig(alertDefinitionName string, alertDefinitionDescription string, alertDefinitionEnabled string, alertDefinitionPriority string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			instance_id      = "%s"
			region           = "%s"
			name             = "%s"
			description     = "%s"
			deleted          = false
			enabled          = %s
			entity_labels    = {}
			group_by_keys    = []
			phantom_mode     = false
			priority         = "%s"
			type             = "logs_threshold"
			incidents_settings {
				minutes   = 1
				notify_on = "triggered_only_unspecified"
			}

			logs_threshold {
				logs_filter {
					simple_filter {
						lucene_query = "\"push\""

						label_filters {
							severities = []

							application_name {
								operation = "is_or_unspecified"
								value     = "sev4"
							}

							subsystem_name {
								operation = "is_or_unspecified"
								value     = "sev4-logs"
							}
						}
					}
				}

				notification_payload_filter = []

				rules {
					condition {
						condition_type = "more_than_or_unspecified"
						threshold      = 1

						time_window {
							logs_time_window_specific_value = "minutes_10"
						}
					}
					override {
						priority = "p3"
					}
				}
				rules {
					condition {
						condition_type = "more_than_or_unspecified"
						threshold      = 1

						time_window {
							logs_time_window_specific_value = "minutes_5_or_unspecified"
						}
					}
					override {
						priority = "p4"
					}
				}
			}
        }

		data "ibm_logs_alert_definitions" "logs_alert_definitions_instance" {
			depends_on = [
				ibm_logs_alert_definition.logs_alert_definition_instance
			]
			instance_id = ibm_logs_alert_definition.logs_alert_definition_instance.instance_id
			region 		= ibm_logs_alert_definition.logs_alert_definition_instance.region
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, alertDefinitionName, alertDefinitionDescription, alertDefinitionEnabled, alertDefinitionPriority)
}
