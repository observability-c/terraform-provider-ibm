// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestAccIbmLogsAlertDefinitionBasic(t *testing.T) {
	var conf logsv0.AlertDefinition
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	isEnabled := "false"
	priority := "p3"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	isEnabledUpdate := "true"
	priorityUpdate := "p2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(name, isEnabled, priority),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", isEnabled),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priority),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(nameUpdate, isEnabledUpdate, priorityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", isEnabledUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priorityUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertDefinitionAllArgs(t *testing.T) {
	var conf logsv0.AlertDefinition
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	enabled := "false"
	priority := "p5_or_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	enabledUpdate := "true"
	priorityUpdate := "p1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		// CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfig(name, enabled, priority),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priority),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfig(nameUpdate, enabledUpdate, priorityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priorityUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_alert_definition.logs_alert_definition_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsAlertDefinitionConfigBasic(name string, isEnabled string, priority string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			instance_id      = "%s"
			region           = "%s"
			name             = "%s"
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, isEnabled, priority)
}

func testAccCheckIbmLogsAlertDefinitionConfig(name string, isEnabled string, priority string) string {

	return fmt.Sprintf(`	
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			instance_id      = "%s"
			region           = "%s"
			name             = "%s"
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, isEnabled, priority)
}

func testAccCheckIbmLogsAlertDefinitionExists(n string, obj logsv0.AlertDefinition) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}

		getAlertDefOptions := &logsv0.GetAlertDefOptions{}

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		alertDefinitionIntf, _, err := logsClient.GetAlertDef(getAlertDefOptions)
		if err != nil {
			return err
		}

		alertDefinition := alertDefinitionIntf.(*logsv0.AlertDefinition)
		obj = *alertDefinition
		return nil
	}
}

func testAccCheckIbmLogsAlertDefinitionDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_alert_definition" {
			continue
		}

		getAlertDefOptions := &logsv0.GetAlertDefOptions{}

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		getAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		// Try to find the key
		_, response, err := logsClient.GetAlertDef(getAlertDefOptions)

		if err == nil {
			return fmt.Errorf("logs_alert_definition still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_alert_definition (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
