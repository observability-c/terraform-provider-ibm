// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

// func TestAccIbmLogsAlertDefinitionBasic(t *testing.T) {
// 	var conf logsv0.AlertDefinition
// 	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
// 	typeVar := "logs_immediate_or_unspecified"
// 	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
// 	typeVarUpdate := "flow"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { acc.TestAccPreCheck(t) },
// 		Providers:    acc.TestAccProviders,
// 		CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(name, typeVar),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVar),
// 				),
// 			},
// 			resource.TestStep{
// 				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(nameUpdate, typeVarUpdate),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVarUpdate),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccIbmLogsAlertDefinitionAllArgs(t *testing.T) {
// 	var conf logsv0.AlertDefinition
// 	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
// 	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
// 	enabled := "false"
// 	priority := "p5_or_unspecified"
// 	typeVar := "logs_immediate_or_unspecified"
// 	phantomMode := "true"
// 	deleted := "false"
// 	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
// 	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
// 	enabledUpdate := "true"
// 	priorityUpdate := "p1"
// 	typeVarUpdate := "flow"
// 	phantomModeUpdate := "false"
// 	deletedUpdate := "true"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { acc.TestAccPreCheck(t) },
// 		Providers:    acc.TestAccProviders,
// 		CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccCheckIbmLogsAlertDefinitionConfig(name, description, enabled, priority, typeVar, phantomMode, deleted),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "description", description),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabled),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priority),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVar),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "phantom_mode", phantomMode),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "deleted", deleted),
// 				),
// 			},
// 			resource.TestStep{
// 				Config: testAccCheckIbmLogsAlertDefinitionConfig(nameUpdate, descriptionUpdate, enabledUpdate, priorityUpdate, typeVarUpdate, phantomModeUpdate, deletedUpdate),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "description", descriptionUpdate),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabledUpdate),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priorityUpdate),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVarUpdate),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "phantom_mode", phantomModeUpdate),
// 					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "deleted", deletedUpdate),
// 				),
// 			},
// 			resource.TestStep{
// 				ResourceName:      "ibm_logs_alert_definition.logs_alert_definition",
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

func testAccCheckIbmLogsAlertDefinitionConfigBasic(name string, typeVar string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			name = "%s"
			type = "%s"
			group_by_keys = "FIXME"
		}
	`, name, typeVar)
}

func testAccCheckIbmLogsAlertDefinitionConfig(name string, description string, enabled string, priority string, typeVar string, phantomMode string, deleted string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			name = "%s"
			description = "%s"
			enabled = %s
			priority = "%s"
			active_on {
				day_of_week = ["sunday"]
				start_time {
					hours = 14
					minutes = 30
				}
				end_time {
					hours = 14
					minutes = 30
				}
			}
			type = "%s"
			group_by_keys = "FIXME"
			incidents_settings {
				notify_on = "triggered_and_resolved"
				minutes = 30
			}
			notification_group {
				group_by_keys = ["key1","key2"]
				webhooks {
					notify_on = "triggered_and_resolved"
					integration {
						integration_id = 123
					}
					minutes = 15
				}
			}
			entity_labels = "FIXME"
			phantom_mode = %s
			deleted = %s
			logs_immediate {
				logs_filter {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				notification_payload_filter = ["obj.field"]
			}
			logs_threshold {
				logs_filter {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				rules {
					condition {
						threshold = 100.0
						time_window {
							logs_time_window_specific_value = "hours_36"
						}
						condition_type = "less_than"
					}
					override {
						priority = "p1"
					}
				}
				notification_payload_filter = ["obj.field"]
				evaluation_delay_ms = 60000
			}
			logs_ratio_threshold {
				numerator {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				numerator_alias = "numerator_alias"
				denominator {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				denominator_alias = "denominator_alias"
				rules {
					condition {
						threshold = 10.0
						time_window {
							logs_ratio_time_window_specific_value = "hours_36"
						}
						condition_type = "less_than"
					}
					override {
						priority = "p1"
					}
				}
				notification_payload_filter = ["obj.field"]
				group_by_for = "denumerator_only"
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				ignore_infinity = true
				evaluation_delay_ms = 60000
			}
			logs_time_relative_threshold {
				logs_filter {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						threshold = 100.0
						compared_to = "same_day_last_month"
						condition_type = "less_than"
					}
					override {
						priority = "p1"
					}
				}
				ignore_infinity = true
				notification_payload_filter = ["obj.field"]
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				evaluation_delay_ms = 60000
			}
			metric_threshold {
				metric_filter {
					promql = "avg_over_time(metric_name[5m]) > 10"
				}
				rules {
					condition {
						threshold = 100.0
						for_over_pct = 80
						of_the_last {
							metric_time_window_specific_value = "hours_36"
						}
						condition_type = "less_than_or_equals"
					}
					override {
						priority = "p1"
					}
				}
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				missing_values {
					replace_with_zero = true
				}
				evaluation_delay_ms = 60000
			}
			flow {
				stages {
					timeframe_ms = "60000"
					timeframe_type = "up_to"
					flow_stages_groups {
						groups {
							alert_defs {
								id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
								not = true
							}
							next_op = "or"
							alerts_op = "or"
						}
					}
				}
				enforce_suppression = true
			}
			logs_anomaly {
				logs_filter {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						minimum_threshold = 10.0
						time_window {
							logs_time_window_specific_value = "hours_36"
						}
						condition_type = "more_than_usual_or_unspecified"
					}
				}
				notification_payload_filter = ["obj.field"]
				evaluation_delay_ms = 60000
				anomaly_alert_settings {
					percentage_of_deviation = 10.0
				}
			}
			metric_anomaly {
				metric_filter {
					promql = "avg_over_time(metric_name[5m]) > 10"
				}
				rules {
					condition {
						threshold = 10.0
						for_over_pct = 20
						of_the_last {
							metric_time_window_specific_value = "hours_36"
						}
						min_non_null_values_pct = 10
						condition_type = "less_than_usual"
					}
				}
				evaluation_delay_ms = 60000
				anomaly_alert_settings {
					percentage_of_deviation = 10.0
				}
			}
			logs_new_value {
				logs_filter {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						keypath_to_track = "metadata.field"
						time_window {
							logs_new_value_time_window_specific_value = "months_3"
						}
					}
				}
				notification_payload_filter = ["obj.field"]
			}
			logs_unique_count {
				logs_filter {
					simple_filter {
						lucene_query = "text:\"error\""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						max_unique_count = "100"
						time_window {
							logs_unique_value_time_window_specific_value = "hours_36"
						}
					}
				}
				notification_payload_filter = ["obj.field"]
				max_unique_count_per_group_by_key = "100"
				unique_count_keypath = "obj.field"
			}
		}
	`, name, description, enabled, priority, typeVar, phantomMode, deleted)
}

// TODO
// func testAccCheckIbmLogsAlertDefinitionExists(n string, obj logsv0.AlertDefinition) resource.TestCheckFunc {

// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", n)
// 		}

// 		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
// 		if err != nil {
// 			return err
// 		}

// 		getAlertDefOptions := &logsv0.GetAlertDefOptions{}

// 		getAlertDefOptions.SetID(rs.Primary.ID)

// 		alertDefinitionIntf, _, err := logsClient.GetAlertDef(getAlertDefOptions)
// 		if err != nil {
// 			return err
// 		}

// 		alertDefinition := alertDefinitionIntf.(*logsv0.AlertDefinition)
// 		obj = *alertDefinition
// 		return nil
// 	}
// }

// func testAccCheckIbmLogsAlertDefinitionDestroy(s *terraform.State) error {
// 	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
// 	if err != nil {
// 		return err
// 	}
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "ibm_logs_alert_definition" {
// 			continue
// 		}

// 		getAlertDefOptions := &logsv0.GetAlertDefOptions{}

// 		getAlertDefOptions.SetID(rs.Primary.ID)
// 		// Try to find the key
// 		_, response, err := logsClient.GetAlertDef(getAlertDefOptions)

// 		if err == nil {
// 			return fmt.Errorf("logs_alert_definition still exists: %s", rs.Primary.ID)
// 		} else if response.StatusCode != 404 {
// 			return fmt.Errorf("Error checking for logs_alert_definition (%s) has been destroyed: %s", rs.Primary.ID, err)
// 		}
// 	}

// 	return nil
// }

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}
		model["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		model["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	model := new(logsv0.ApisAlertDefinitionActivitySchedule)
	model.DayOfWeek = []string{"sunday"}
	model.StartTime = apisAlertDefinitionTimeOfDayModel
	model.EndTime = apisAlertDefinitionTimeOfDayModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hours"] = int(14)
		model["minutes"] = int(30)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionTimeOfDay)
	model.Hours = core.Int64Ptr(int64(14))
	model.Minutes = core.Int64Ptr(int64(30))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["notify_on"] = "triggered_and_resolved"
		model["minutes"] = int(30)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	model.NotifyOn = core.StringPtr("triggered_and_resolved")
	model.Minutes = core.Int64Ptr(int64(30))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		model := make(map[string]interface{})
		model["group_by_keys"] = []string{"key1", "key2"}
		model["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	model := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	model.GroupByKeys = []string{"key1", "key2"}
	model.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		model := make(map[string]interface{})
		model["notify_on"] = "triggered_and_resolved"
		model["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		model["minutes"] = int(15)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	model := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	model.NotifyOn = core.StringPtr("triggered_and_resolved")
	model.Integration = apisAlertDefinitionIntegrationTypeModel
	model.Minutes = core.Int64Ptr(int64(15))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["integration_id"] = int(123)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionIntegrationType)
	model.IntegrationID = core.Int64Ptr(int64(123))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["integration_id"] = int(123)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	model.IntegrationID = core.Int64Ptr(int64(123))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["notification_payload_filter"] = []string{"obj.field"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	model := new(logsv0.ApisAlertDefinitionLogsImmediateType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.NotificationPayloadFilter = []string{"obj.field"}

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		model := make(map[string]interface{})
		model["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	model := new(logsv0.ApisAlertDefinitionLogsFilter)
	model.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		model := make(map[string]interface{})
		model["lucene_query"] = "text:\"error\""
		model["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	model := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	model.LuceneQuery = core.StringPtr("text:\"error\"")
	model.LabelFilters = apisAlertDefinitionLabelFiltersModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		model := make(map[string]interface{})
		model["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		model["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		model["severities"] = []string{"critical"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	model := new(logsv0.ApisAlertDefinitionLabelFilters)
	model.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	model.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	model.Severities = []string{"critical"}

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = "my-app"
		model["operation"] = "starts_with"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLabelFilterType)
	model.Value = core.StringPtr("my-app")
	model.Operation = core.StringPtr("starts_with")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}
		apisAlertDefinitionLogsThresholdConditionModel["condition_type"] = "less_than"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsThresholdRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdConditionModel}
		apisAlertDefinitionLogsThresholdRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdRuleModel}
		model["notification_payload_filter"] = []string{"obj.field"}
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
	apisAlertDefinitionLogsThresholdConditionModel.ConditionType = core.StringPtr("less_than")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsThresholdRuleModel := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
	apisAlertDefinitionLogsThresholdRuleModel.Condition = apisAlertDefinitionLogsThresholdConditionModel
	apisAlertDefinitionLogsThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	model := new(logsv0.ApisAlertDefinitionLogsThresholdType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsThresholdRule{*apisAlertDefinitionLogsThresholdRuleModel}
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["trigger_undetected_values"] = true
		model["auto_retire_timeframe"] = "hours_24"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	model.TriggerUndetectedValues = core.BoolPtr(true)
	model.AutoRetireTimeframe = core.StringPtr("hours_24")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}
		apisAlertDefinitionLogsThresholdConditionModel["condition_type"] = "less_than"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
	apisAlertDefinitionLogsThresholdConditionModel.ConditionType = core.StringPtr("less_than")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
	model.Condition = apisAlertDefinitionLogsThresholdConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(100.0)
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}
		model["condition_type"] = "less_than"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	model.Threshold = core.Float64Ptr(float64(100.0))
	model.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
	model.ConditionType = core.StringPtr("less_than")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	model.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["priority"] = "p1"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	model.Priority = core.StringPtr("p1")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionLogsRatioConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}
		apisAlertDefinitionLogsRatioConditionModel["condition_type"] = "less_than"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsRatioRulesModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioRulesModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsRatioConditionModel}
		apisAlertDefinitionLogsRatioRulesModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		model := make(map[string]interface{})
		model["numerator"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["numerator_alias"] = "numerator_alias"
		model["denominator"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["denominator_alias"] = "denominator_alias"
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsRatioRulesModel}
		model["notification_payload_filter"] = []string{"obj.field"}
		model["group_by_for"] = "denumerator_only"
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["ignore_infinity"] = true
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
	apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel
	apisAlertDefinitionLogsRatioConditionModel.ConditionType = core.StringPtr("less_than")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsRatioRulesModel := new(logsv0.ApisAlertDefinitionLogsRatioRules)
	apisAlertDefinitionLogsRatioRulesModel.Condition = apisAlertDefinitionLogsRatioConditionModel
	apisAlertDefinitionLogsRatioRulesModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	model := new(logsv0.ApisAlertDefinitionLogsRatioThresholdType)
	model.Numerator = apisAlertDefinitionLogsFilterModel
	model.NumeratorAlias = core.StringPtr("numerator_alias")
	model.Denominator = apisAlertDefinitionLogsFilterModel
	model.DenominatorAlias = core.StringPtr("denominator_alias")
	model.Rules = []logsv0.ApisAlertDefinitionLogsRatioRules{*apisAlertDefinitionLogsRatioRulesModel}
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.GroupByFor = core.StringPtr("denumerator_only")
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.IgnoreInfinity = core.BoolPtr(true)
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionLogsRatioConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}
		apisAlertDefinitionLogsRatioConditionModel["condition_type"] = "less_than"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsRatioConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
	apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel
	apisAlertDefinitionLogsRatioConditionModel.ConditionType = core.StringPtr("less_than")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionLogsRatioRules)
	model.Condition = apisAlertDefinitionLogsRatioConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(10.0)
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}
		model["condition_type"] = "less_than"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
	model.Threshold = core.Float64Ptr(float64(10.0))
	model.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel
	model.ConditionType = core.StringPtr("less_than")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_ratio_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	model.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"
		apisAlertDefinitionLogsTimeRelativeConditionModel["condition_type"] = "less_than"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsTimeRelativeRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
		apisAlertDefinitionLogsTimeRelativeRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeRuleModel}
		model["ignore_infinity"] = true
		model["notification_payload_filter"] = []string{"obj.field"}
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
	apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")
	apisAlertDefinitionLogsTimeRelativeConditionModel.ConditionType = core.StringPtr("less_than")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsTimeRelativeRuleModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
	apisAlertDefinitionLogsTimeRelativeRuleModel.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
	apisAlertDefinitionLogsTimeRelativeRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsTimeRelativeRule{*apisAlertDefinitionLogsTimeRelativeRuleModel}
	model.IgnoreInfinity = core.BoolPtr(true)
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"
		apisAlertDefinitionLogsTimeRelativeConditionModel["condition_type"] = "less_than"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
	apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")
	apisAlertDefinitionLogsTimeRelativeConditionModel.ConditionType = core.StringPtr("less_than")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
	model.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["threshold"] = float64(100.0)
		model["compared_to"] = "same_day_last_month"
		model["condition_type"] = "less_than"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
	model.Threshold = core.Float64Ptr(float64(100.0))
	model.ComparedTo = core.StringPtr("same_day_last_month")
	model.ConditionType = core.StringPtr("less_than")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
		apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
		apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		apisAlertDefinitionMetricThresholdConditionModel["condition_type"] = "less_than_or_equals"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionMetricThresholdRuleModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdConditionModel}
		apisAlertDefinitionMetricThresholdRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionMetricMissingValuesModel := make(map[string]interface{})
		apisAlertDefinitionMetricMissingValuesModel["replace_with_zero"] = true

		model := make(map[string]interface{})
		model["metric_filter"] = []map[string]interface{}{apisAlertDefinitionMetricFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdRuleModel}
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["missing_values"] = []map[string]interface{}{apisAlertDefinitionMetricMissingValuesModel}
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
	apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
	apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
	apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	apisAlertDefinitionMetricThresholdConditionModel.ConditionType = core.StringPtr("less_than_or_equals")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionMetricThresholdRuleModel := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
	apisAlertDefinitionMetricThresholdRuleModel.Condition = apisAlertDefinitionMetricThresholdConditionModel
	apisAlertDefinitionMetricThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionMetricMissingValuesModel := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
	apisAlertDefinitionMetricMissingValuesModel.ReplaceWithZero = core.BoolPtr(true)

	model := new(logsv0.ApisAlertDefinitionMetricThresholdType)
	model.MetricFilter = apisAlertDefinitionMetricFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionMetricThresholdRule{*apisAlertDefinitionMetricThresholdRuleModel}
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.MissingValues = apisAlertDefinitionMetricMissingValuesModel
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["promql"] = "avg_over_time(metric_name[5m]) > 10"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricFilter)
	model.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
		apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		apisAlertDefinitionMetricThresholdConditionModel["condition_type"] = "less_than_or_equals"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
	apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
	apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	apisAlertDefinitionMetricThresholdConditionModel.ConditionType = core.StringPtr("less_than_or_equals")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
	model.Condition = apisAlertDefinitionMetricThresholdConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(100.0)
		model["for_over_pct"] = int(80)
		model["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		model["condition_type"] = "less_than_or_equals"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
	model.Threshold = core.Float64Ptr(float64(100.0))
	model.ForOverPct = core.Int64Ptr(int64(80))
	model.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	model.ConditionType = core.StringPtr("less_than_or_equals")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_specific_value"] = "hours_36"
		model["metric_time_window_dynamic_duration"] = "1h30m"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindow)
	model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")
	model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_dynamic_duration"] = "1h30m"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration)
	model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["replace_with_zero"] = true
		model["min_non_null_values_pct"] = int(80)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValues)
	model.ReplaceWithZero = core.BoolPtr(true)
	model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["replace_with_zero"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
	model.ReplaceWithZero = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["min_non_null_values_pct"] = int(80)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct)
	model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
		apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

		apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsModel["groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupModel}

		apisAlertDefinitionFlowStagesModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesModel["timeframe_ms"] = "60000"
		apisAlertDefinitionFlowStagesModel["timeframe_type"] = "up_to"
		apisAlertDefinitionFlowStagesModel["flow_stages_groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsModel}

		model := make(map[string]interface{})
		model["stages"] = []map[string]interface{}{apisAlertDefinitionFlowStagesModel}
		model["enforce_suppression"] = true

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
	apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

	apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
	apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

	apisAlertDefinitionFlowStagesModel := new(logsv0.ApisAlertDefinitionFlowStages)
	apisAlertDefinitionFlowStagesModel.TimeframeMs = core.StringPtr("60000")
	apisAlertDefinitionFlowStagesModel.TimeframeType = core.StringPtr("up_to")
	apisAlertDefinitionFlowStagesModel.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

	model := new(logsv0.ApisAlertDefinitionFlowType)
	model.Stages = []logsv0.ApisAlertDefinitionFlowStages{*apisAlertDefinitionFlowStagesModel}
	model.EnforceSuppression = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
		apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

		apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsModel["groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupModel}

		model := make(map[string]interface{})
		model["timeframe_ms"] = "60000"
		model["timeframe_type"] = "up_to"
		model["flow_stages_groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
	apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

	apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
	apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

	model := new(logsv0.ApisAlertDefinitionFlowStages)
	model.TimeframeMs = core.StringPtr("60000")
	model.TimeframeType = core.StringPtr("up_to")
	model.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
		apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

		model := make(map[string]interface{})
		model["groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
	apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

	model := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
	model.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		model := make(map[string]interface{})
		model["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		model["next_op"] = "or"
		model["alerts_op"] = "or"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	model := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	model.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	model.NextOp = core.StringPtr("or")
	model.AlertsOp = core.StringPtr("or")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["not"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Not = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
		apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}
		apisAlertDefinitionLogsAnomalyConditionModel["condition_type"] = "more_than_usual_or_unspecified"

		apisAlertDefinitionLogsAnomalyRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

		apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyRuleModel}
		model["notification_payload_filter"] = []string{"obj.field"}
		model["evaluation_delay_ms"] = int(60000)
		model["anomaly_alert_settings"] = []map[string]interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
	apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
	apisAlertDefinitionLogsAnomalyConditionModel.ConditionType = core.StringPtr("more_than_usual_or_unspecified")

	apisAlertDefinitionLogsAnomalyRuleModel := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
	apisAlertDefinitionLogsAnomalyRuleModel.Condition = apisAlertDefinitionLogsAnomalyConditionModel

	apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	model := new(logsv0.ApisAlertDefinitionLogsAnomalyType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsAnomalyRule{*apisAlertDefinitionLogsAnomalyRuleModel}
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))
	model.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
		apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}
		apisAlertDefinitionLogsAnomalyConditionModel["condition_type"] = "more_than_usual_or_unspecified"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
	apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
	apisAlertDefinitionLogsAnomalyConditionModel.ConditionType = core.StringPtr("more_than_usual_or_unspecified")

	model := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
	model.Condition = apisAlertDefinitionLogsAnomalyConditionModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["minimum_threshold"] = float64(10.0)
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}
		model["condition_type"] = "more_than_usual_or_unspecified"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
	model.MinimumThreshold = core.Float64Ptr(float64(10.0))
	model.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
	model.ConditionType = core.StringPtr("more_than_usual_or_unspecified")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["percentage_of_deviation"] = float64(10.0)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	model.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
		apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
		apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)
		apisAlertDefinitionMetricAnomalyConditionModel["condition_type"] = "less_than_usual"

		apisAlertDefinitionMetricAnomalyRuleModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

		apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

		model := make(map[string]interface{})
		model["metric_filter"] = []map[string]interface{}{apisAlertDefinitionMetricFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyRuleModel}
		model["evaluation_delay_ms"] = int(60000)
		model["anomaly_alert_settings"] = []map[string]interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
	apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
	apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
	apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))
	apisAlertDefinitionMetricAnomalyConditionModel.ConditionType = core.StringPtr("less_than_usual")

	apisAlertDefinitionMetricAnomalyRuleModel := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
	apisAlertDefinitionMetricAnomalyRuleModel.Condition = apisAlertDefinitionMetricAnomalyConditionModel

	apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	model := new(logsv0.ApisAlertDefinitionMetricAnomalyType)
	model.MetricFilter = apisAlertDefinitionMetricFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionMetricAnomalyRule{*apisAlertDefinitionMetricAnomalyRuleModel}
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))
	model.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
		apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)
		apisAlertDefinitionMetricAnomalyConditionModel["condition_type"] = "less_than_usual"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
	apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
	apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))
	apisAlertDefinitionMetricAnomalyConditionModel.ConditionType = core.StringPtr("less_than_usual")

	model := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
	model.Condition = apisAlertDefinitionMetricAnomalyConditionModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(10.0)
		model["for_over_pct"] = int(20)
		model["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		model["min_non_null_values_pct"] = int(10)
		model["condition_type"] = "less_than_usual"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
	model.Threshold = core.Float64Ptr(float64(10.0))
	model.ForOverPct = core.Int64Ptr(int64(20))
	model.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	model.MinNonNullValuesPct = core.Int64Ptr(int64(10))
	model.ConditionType = core.StringPtr("less_than_usual")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

		apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
		apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

		apisAlertDefinitionLogsNewValueRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueConditionModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueRuleModel}
		model["notification_payload_filter"] = []string{"obj.field"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
	apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
	apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

	apisAlertDefinitionLogsNewValueRuleModel := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
	apisAlertDefinitionLogsNewValueRuleModel.Condition = apisAlertDefinitionLogsNewValueConditionModel

	model := new(logsv0.ApisAlertDefinitionLogsNewValueType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsNewValueRule{*apisAlertDefinitionLogsNewValueRuleModel}
	model.NotificationPayloadFilter = []string{"obj.field"}

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

		apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
		apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
	apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
	apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

	model := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
	model.Condition = apisAlertDefinitionLogsNewValueConditionModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

		model := make(map[string]interface{})
		model["keypath_to_track"] = "metadata.field"
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	model := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
	model.KeypathToTrack = core.StringPtr("metadata.field")
	model.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_new_value_time_window_specific_value"] = "months_3"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	model.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
		apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

		apisAlertDefinitionLogsUniqueCountRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountRuleModel}
		model["notification_payload_filter"] = []string{"obj.field"}
		model["max_unique_count_per_group_by_key"] = "100"
		model["unique_count_keypath"] = "obj.field"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
	apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
	apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

	apisAlertDefinitionLogsUniqueCountRuleModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
	apisAlertDefinitionLogsUniqueCountRuleModel.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

	model := new(logsv0.ApisAlertDefinitionLogsUniqueCountType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsUniqueCountRule{*apisAlertDefinitionLogsUniqueCountRuleModel}
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.MaxUniqueCountPerGroupByKey = core.StringPtr("100")
	model.UniqueCountKeypath = core.StringPtr("obj.field")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
		apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
	apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
	apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

	model := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
	model.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["max_unique_count"] = "100"
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
	model.MaxUniqueCount = core.StringPtr("100")
	model.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_unique_value_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	model.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.ResourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionActivitySchedule) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		model := new(logsv0.ApisAlertDefinitionActivitySchedule)
		model.DayOfWeek = []string{"sunday"}
		model.StartTime = apisAlertDefinitionTimeOfDayModel
		model.EndTime = apisAlertDefinitionTimeOfDayModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	model := make(map[string]interface{})
	model["day_of_week"] = []interface{}{"sunday"}
	model["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	model["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionActivitySchedule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionTimeOfDay(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionTimeOfDay) {
		model := new(logsv0.ApisAlertDefinitionTimeOfDay)
		model.Hours = core.Int64Ptr(int64(14))
		model.Minutes = core.Int64Ptr(int64(30))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hours"] = int(14)
	model["minutes"] = int(30)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionTimeOfDay(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionAlertDefIncidentSettings) {
		model := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		model.NotifyOn = core.StringPtr("triggered_and_resolved")
		model.Minutes = core.Int64Ptr(int64(30))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["notify_on"] = "triggered_and_resolved"
	model["minutes"] = int(30)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefIncidentSettings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionAlertDefNotificationGroup) {
		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		model := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		model.GroupByKeys = []string{"key1", "key2"}
		model.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	model := make(map[string]interface{})
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefNotificationGroup(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefWebhooksSettings(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionAlertDefWebhooksSettings) {
		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		model := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		model.NotifyOn = core.StringPtr("triggered_and_resolved")
		model.Integration = apisAlertDefinitionIntegrationTypeModel
		model.Minutes = core.Int64Ptr(int64(15))

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	model := make(map[string]interface{})
	model["notify_on"] = "triggered_and_resolved"
	model["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	model["minutes"] = int(15)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefWebhooksSettings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionIntegrationType(t *testing.T) {
	checkResult := func(result logsv0.ApisAlertDefinitionIntegrationTypeIntf) {
		model := new(logsv0.ApisAlertDefinitionIntegrationType)
		model.IntegrationID = core.Int64Ptr(int64(123))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["integration_id"] = int(123)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionIntegrationType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID) {
		model := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		model.IntegrationID = core.Int64Ptr(int64(123))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["integration_id"] = int(123)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsImmediateType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsImmediateType) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		model := new(logsv0.ApisAlertDefinitionLogsImmediateType)
		model.LogsFilter = apisAlertDefinitionLogsFilterModel
		model.NotificationPayloadFilter = []string{"obj.field"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	model := make(map[string]interface{})
	model["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["notification_payload_filter"] = []interface{}{"obj.field"}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsImmediateType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsFilter) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		model := new(logsv0.ApisAlertDefinitionLogsFilter)
		model.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	model := make(map[string]interface{})
	model["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsFilter(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsSimpleFilter(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsSimpleFilter) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		model := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		model.LuceneQuery = core.StringPtr("text:\"error\"")
		model.LabelFilters = apisAlertDefinitionLabelFiltersModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	model := make(map[string]interface{})
	model["lucene_query"] = "text:\"error\""
	model["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsSimpleFilter(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilters(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLabelFilters) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		model := new(logsv0.ApisAlertDefinitionLabelFilters)
		model.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		model.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		model.Severities = []string{"critical"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	model := make(map[string]interface{})
	model["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	model["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	model["severities"] = []interface{}{"critical"}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilterType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLabelFilterType) {
		model := new(logsv0.ApisAlertDefinitionLabelFilterType)
		model.Value = core.StringPtr("my-app")
		model.Operation = core.StringPtr("starts_with")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["value"] = "my-app"
	model["operation"] = "starts_with"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLabelFilterType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsThresholdType) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
		apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		apisAlertDefinitionLogsThresholdConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionLogsThresholdRuleModel := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
		apisAlertDefinitionLogsThresholdRuleModel.Condition = apisAlertDefinitionLogsThresholdConditionModel
		apisAlertDefinitionLogsThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

		model := new(logsv0.ApisAlertDefinitionLogsThresholdType)
		model.LogsFilter = apisAlertDefinitionLogsFilterModel
		model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		model.Rules = []logsv0.ApisAlertDefinitionLogsThresholdRule{*apisAlertDefinitionLogsThresholdRuleModel}
		model.NotificationPayloadFilter = []string{"obj.field"}
		model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	apisAlertDefinitionLogsThresholdConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionLogsThresholdRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsThresholdRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsThresholdConditionModel}
	apisAlertDefinitionLogsThresholdRuleModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	model := make(map[string]interface{})
	model["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	model["rules"] = []interface{}{apisAlertDefinitionLogsThresholdRuleModel}
	model["notification_payload_filter"] = []interface{}{"obj.field"}
	model["evaluation_delay_ms"] = int(60000)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionUndetectedValuesManagement(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionUndetectedValuesManagement) {
		model := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		model.TriggerUndetectedValues = core.BoolPtr(true)
		model.AutoRetireTimeframe = core.StringPtr("hours_24")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["trigger_undetected_values"] = true
	model["auto_retire_timeframe"] = "hours_24"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionUndetectedValuesManagement(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdRule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsThresholdRule) {
		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
		apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		apisAlertDefinitionLogsThresholdConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		model := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
		model.Condition = apisAlertDefinitionLogsThresholdConditionModel
		model.Override = apisAlertDefinitionAlertDefOverrideModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	apisAlertDefinitionLogsThresholdConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionLogsThresholdConditionModel}
	model["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsThresholdCondition) {
		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		model := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
		model.Threshold = core.Float64Ptr(float64(100.0))
		model.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		model.ConditionType = core.StringPtr("less_than")

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	model := make(map[string]interface{})
	model["threshold"] = float64(100.0)
	model["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	model["condition_type"] = "less_than"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsThresholdCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeWindow(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsTimeWindow) {
		model := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		model.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["logs_time_window_specific_value"] = "hours_36"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeWindow(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefOverride(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionAlertDefOverride) {
		model := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		model.Priority = core.StringPtr("p1")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["priority"] = "p1"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAlertDefOverride(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioThresholdType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsRatioThresholdType) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
		apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
		apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel
		apisAlertDefinitionLogsRatioConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionLogsRatioRulesModel := new(logsv0.ApisAlertDefinitionLogsRatioRules)
		apisAlertDefinitionLogsRatioRulesModel.Condition = apisAlertDefinitionLogsRatioConditionModel
		apisAlertDefinitionLogsRatioRulesModel.Override = apisAlertDefinitionAlertDefOverrideModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		model := new(logsv0.ApisAlertDefinitionLogsRatioThresholdType)
		model.Numerator = apisAlertDefinitionLogsFilterModel
		model.NumeratorAlias = core.StringPtr("numerator_alias")
		model.Denominator = apisAlertDefinitionLogsFilterModel
		model.DenominatorAlias = core.StringPtr("denominator_alias")
		model.Rules = []logsv0.ApisAlertDefinitionLogsRatioRules{*apisAlertDefinitionLogsRatioRulesModel}
		model.NotificationPayloadFilter = []string{"obj.field"}
		model.GroupByFor = core.StringPtr("denumerator_only")
		model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		model.IgnoreInfinity = core.BoolPtr(true)
		model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
	apisAlertDefinitionLogsRatioConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}
	apisAlertDefinitionLogsRatioConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionLogsRatioRulesModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioRulesModel["condition"] = []interface{}{apisAlertDefinitionLogsRatioConditionModel}
	apisAlertDefinitionLogsRatioRulesModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	model := make(map[string]interface{})
	model["numerator"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["numerator_alias"] = "numerator_alias"
	model["denominator"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["denominator_alias"] = "denominator_alias"
	model["rules"] = []interface{}{apisAlertDefinitionLogsRatioRulesModel}
	model["notification_payload_filter"] = []interface{}{"obj.field"}
	model["group_by_for"] = "denumerator_only"
	model["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	model["ignore_infinity"] = true
	model["evaluation_delay_ms"] = int(60000)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioThresholdType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioRules(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsRatioRules) {
		apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
		apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
		apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel
		apisAlertDefinitionLogsRatioConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		model := new(logsv0.ApisAlertDefinitionLogsRatioRules)
		model.Condition = apisAlertDefinitionLogsRatioConditionModel
		model.Override = apisAlertDefinitionAlertDefOverrideModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
	apisAlertDefinitionLogsRatioConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}
	apisAlertDefinitionLogsRatioConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionLogsRatioConditionModel}
	model["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioRules(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsRatioCondition) {
		apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
		apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

		model := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
		model.Threshold = core.Float64Ptr(float64(10.0))
		model.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel
		model.ConditionType = core.StringPtr("less_than")

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

	model := make(map[string]interface{})
	model["threshold"] = float64(10.0)
	model["time_window"] = []interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}
	model["condition_type"] = "less_than"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioTimeWindow(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsRatioTimeWindow) {
		model := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
		model.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["logs_ratio_time_window_specific_value"] = "hours_36"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsRatioTimeWindow(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeThresholdType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
		apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")
		apisAlertDefinitionLogsTimeRelativeConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionLogsTimeRelativeRuleModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
		apisAlertDefinitionLogsTimeRelativeRuleModel.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
		apisAlertDefinitionLogsTimeRelativeRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType)
		model.LogsFilter = apisAlertDefinitionLogsFilterModel
		model.Rules = []logsv0.ApisAlertDefinitionLogsTimeRelativeRule{*apisAlertDefinitionLogsTimeRelativeRuleModel}
		model.IgnoreInfinity = core.BoolPtr(true)
		model.NotificationPayloadFilter = []string{"obj.field"}
		model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"
	apisAlertDefinitionLogsTimeRelativeConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionLogsTimeRelativeRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeRelativeRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
	apisAlertDefinitionLogsTimeRelativeRuleModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	model := make(map[string]interface{})
	model["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["rules"] = []interface{}{apisAlertDefinitionLogsTimeRelativeRuleModel}
	model["ignore_infinity"] = true
	model["notification_payload_filter"] = []interface{}{"obj.field"}
	model["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	model["evaluation_delay_ms"] = int(60000)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeThresholdType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeRule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsTimeRelativeRule) {
		apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
		apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")
		apisAlertDefinitionLogsTimeRelativeConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
		model.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
		model.Override = apisAlertDefinitionAlertDefOverrideModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"
	apisAlertDefinitionLogsTimeRelativeConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
	model["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsTimeRelativeCondition) {
		model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
		model.Threshold = core.Float64Ptr(float64(100.0))
		model.ComparedTo = core.StringPtr("same_day_last_month")
		model.ConditionType = core.StringPtr("less_than")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["threshold"] = float64(100.0)
	model["compared_to"] = "same_day_last_month"
	model["condition_type"] = "less_than"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsTimeRelativeCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricThresholdType) {
		apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
		apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
		apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
		apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		apisAlertDefinitionMetricThresholdConditionModel.ConditionType = core.StringPtr("less_than_or_equals")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionMetricThresholdRuleModel := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
		apisAlertDefinitionMetricThresholdRuleModel.Condition = apisAlertDefinitionMetricThresholdConditionModel
		apisAlertDefinitionMetricThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		apisAlertDefinitionMetricMissingValuesModel := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
		apisAlertDefinitionMetricMissingValuesModel.ReplaceWithZero = core.BoolPtr(true)

		model := new(logsv0.ApisAlertDefinitionMetricThresholdType)
		model.MetricFilter = apisAlertDefinitionMetricFilterModel
		model.Rules = []logsv0.ApisAlertDefinitionMetricThresholdRule{*apisAlertDefinitionMetricThresholdRuleModel}
		model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		model.MissingValues = apisAlertDefinitionMetricMissingValuesModel
		model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
	apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
	apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
	apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	apisAlertDefinitionMetricThresholdConditionModel["condition_type"] = "less_than_or_equals"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionMetricThresholdRuleModel := make(map[string]interface{})
	apisAlertDefinitionMetricThresholdRuleModel["condition"] = []interface{}{apisAlertDefinitionMetricThresholdConditionModel}
	apisAlertDefinitionMetricThresholdRuleModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	apisAlertDefinitionMetricMissingValuesModel := make(map[string]interface{})
	apisAlertDefinitionMetricMissingValuesModel["replace_with_zero"] = true

	model := make(map[string]interface{})
	model["metric_filter"] = []interface{}{apisAlertDefinitionMetricFilterModel}
	model["rules"] = []interface{}{apisAlertDefinitionMetricThresholdRuleModel}
	model["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	model["missing_values"] = []interface{}{apisAlertDefinitionMetricMissingValuesModel}
	model["evaluation_delay_ms"] = int(60000)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricFilter(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricFilter) {
		model := new(logsv0.ApisAlertDefinitionMetricFilter)
		model.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["promql"] = "avg_over_time(metric_name[5m]) > 10"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricFilter(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdRule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricThresholdRule) {
		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
		apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
		apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		apisAlertDefinitionMetricThresholdConditionModel.ConditionType = core.StringPtr("less_than_or_equals")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		model := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
		model.Condition = apisAlertDefinitionMetricThresholdConditionModel
		model.Override = apisAlertDefinitionAlertDefOverrideModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
	apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
	apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	apisAlertDefinitionMetricThresholdConditionModel["condition_type"] = "less_than_or_equals"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionMetricThresholdConditionModel}
	model["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricThresholdCondition) {
		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		model := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
		model.Threshold = core.Float64Ptr(float64(100.0))
		model.ForOverPct = core.Int64Ptr(int64(80))
		model.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		model.ConditionType = core.StringPtr("less_than_or_equals")

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	model := make(map[string]interface{})
	model["threshold"] = float64(100.0)
	model["for_over_pct"] = int(80)
	model["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	model["condition_type"] = "less_than_or_equals"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricThresholdCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindow(t *testing.T) {
	checkResult := func(result logsv0.ApisAlertDefinitionMetricTimeWindowIntf) {
		model := new(logsv0.ApisAlertDefinitionMetricTimeWindow)
		model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")
		model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["metric_time_window_specific_value"] = "hours_36"
	model["metric_time_window_dynamic_duration"] = "1h30m"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindow(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue) {
		model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["metric_time_window_specific_value"] = "hours_36"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration) {
		model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration)
		model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["metric_time_window_dynamic_duration"] = "1h30m"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValues(t *testing.T) {
	checkResult := func(result logsv0.ApisAlertDefinitionMetricMissingValuesIntf) {
		model := new(logsv0.ApisAlertDefinitionMetricMissingValues)
		model.ReplaceWithZero = core.BoolPtr(true)
		model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["replace_with_zero"] = true
	model["min_non_null_values_pct"] = int(80)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValues(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero) {
		model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
		model.ReplaceWithZero = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["replace_with_zero"] = true

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct) {
		model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct)
		model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["min_non_null_values_pct"] = int(80)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionFlowType) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

		apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
		apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
		apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

		apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
		apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

		apisAlertDefinitionFlowStagesModel := new(logsv0.ApisAlertDefinitionFlowStages)
		apisAlertDefinitionFlowStagesModel.TimeframeMs = core.StringPtr("60000")
		apisAlertDefinitionFlowStagesModel.TimeframeType = core.StringPtr("up_to")
		apisAlertDefinitionFlowStagesModel.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

		model := new(logsv0.ApisAlertDefinitionFlowType)
		model.Stages = []logsv0.ApisAlertDefinitionFlowStages{*apisAlertDefinitionFlowStagesModel}
		model.EnforceSuppression = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

	apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
	apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

	apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsModel["groups"] = []interface{}{apisAlertDefinitionFlowStagesGroupModel}

	apisAlertDefinitionFlowStagesModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesModel["timeframe_ms"] = "60000"
	apisAlertDefinitionFlowStagesModel["timeframe_type"] = "up_to"
	apisAlertDefinitionFlowStagesModel["flow_stages_groups"] = []interface{}{apisAlertDefinitionFlowStagesGroupsModel}

	model := make(map[string]interface{})
	model["stages"] = []interface{}{apisAlertDefinitionFlowStagesModel}
	model["enforce_suppression"] = true

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStages(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionFlowStages) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

		apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
		apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
		apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

		apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
		apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

		model := new(logsv0.ApisAlertDefinitionFlowStages)
		model.TimeframeMs = core.StringPtr("60000")
		model.TimeframeType = core.StringPtr("up_to")
		model.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

	apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
	apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

	apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsModel["groups"] = []interface{}{apisAlertDefinitionFlowStagesGroupModel}

	model := make(map[string]interface{})
	model["timeframe_ms"] = "60000"
	model["timeframe_type"] = "up_to"
	model["flow_stages_groups"] = []interface{}{apisAlertDefinitionFlowStagesGroupsModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStages(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroups(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionFlowStagesGroups) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

		apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
		apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
		apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

		model := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
		model.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

	apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
	apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

	model := make(map[string]interface{})
	model["groups"] = []interface{}{apisAlertDefinitionFlowStagesGroupModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroups(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroup(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionFlowStagesGroup) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

		model := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
		model.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		model.NextOp = core.StringPtr("or")
		model.AlertsOp = core.StringPtr("or")

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

	model := make(map[string]interface{})
	model["alert_defs"] = []interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	model["next_op"] = "or"
	model["alerts_op"] = "or"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroup(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroupsAlertDefs(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs) {
		model := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
		model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
		model.Not = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
	model["not"] = true

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionFlowStagesGroupsAlertDefs(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsAnomalyType) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
		apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		apisAlertDefinitionLogsAnomalyConditionModel.ConditionType = core.StringPtr("more_than_usual_or_unspecified")

		apisAlertDefinitionLogsAnomalyRuleModel := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
		apisAlertDefinitionLogsAnomalyRuleModel.Condition = apisAlertDefinitionLogsAnomalyConditionModel

		apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
		apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

		model := new(logsv0.ApisAlertDefinitionLogsAnomalyType)
		model.LogsFilter = apisAlertDefinitionLogsFilterModel
		model.Rules = []logsv0.ApisAlertDefinitionLogsAnomalyRule{*apisAlertDefinitionLogsAnomalyRuleModel}
		model.NotificationPayloadFilter = []string{"obj.field"}
		model.EvaluationDelayMs = core.Int64Ptr(int64(60000))
		model.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
	apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	apisAlertDefinitionLogsAnomalyConditionModel["condition_type"] = "more_than_usual_or_unspecified"

	apisAlertDefinitionLogsAnomalyRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsAnomalyRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

	apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

	model := make(map[string]interface{})
	model["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["rules"] = []interface{}{apisAlertDefinitionLogsAnomalyRuleModel}
	model["notification_payload_filter"] = []interface{}{"obj.field"}
	model["evaluation_delay_ms"] = int(60000)
	model["anomaly_alert_settings"] = []interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyRule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsAnomalyRule) {
		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
		apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		apisAlertDefinitionLogsAnomalyConditionModel.ConditionType = core.StringPtr("more_than_usual_or_unspecified")

		model := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
		model.Condition = apisAlertDefinitionLogsAnomalyConditionModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
	apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	apisAlertDefinitionLogsAnomalyConditionModel["condition_type"] = "more_than_usual_or_unspecified"

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsAnomalyCondition) {
		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		model := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
		model.MinimumThreshold = core.Float64Ptr(float64(10.0))
		model.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		model.ConditionType = core.StringPtr("more_than_usual_or_unspecified")

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	model := make(map[string]interface{})
	model["minimum_threshold"] = float64(10.0)
	model["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	model["condition_type"] = "more_than_usual_or_unspecified"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsAnomalyCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAnomalyAlertSettings(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionAnomalyAlertSettings) {
		model := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
		model.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["percentage_of_deviation"] = float64(10.0)

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionAnomalyAlertSettings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricAnomalyType) {
		apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
		apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
		apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
		apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))
		apisAlertDefinitionMetricAnomalyConditionModel.ConditionType = core.StringPtr("less_than_usual")

		apisAlertDefinitionMetricAnomalyRuleModel := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
		apisAlertDefinitionMetricAnomalyRuleModel.Condition = apisAlertDefinitionMetricAnomalyConditionModel

		apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
		apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

		model := new(logsv0.ApisAlertDefinitionMetricAnomalyType)
		model.MetricFilter = apisAlertDefinitionMetricFilterModel
		model.Rules = []logsv0.ApisAlertDefinitionMetricAnomalyRule{*apisAlertDefinitionMetricAnomalyRuleModel}
		model.EvaluationDelayMs = core.Int64Ptr(int64(60000))
		model.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
	apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
	apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
	apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
	apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)
	apisAlertDefinitionMetricAnomalyConditionModel["condition_type"] = "less_than_usual"

	apisAlertDefinitionMetricAnomalyRuleModel := make(map[string]interface{})
	apisAlertDefinitionMetricAnomalyRuleModel["condition"] = []interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

	apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

	model := make(map[string]interface{})
	model["metric_filter"] = []interface{}{apisAlertDefinitionMetricFilterModel}
	model["rules"] = []interface{}{apisAlertDefinitionMetricAnomalyRuleModel}
	model["evaluation_delay_ms"] = int(60000)
	model["anomaly_alert_settings"] = []interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyRule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricAnomalyRule) {
		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
		apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
		apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))
		apisAlertDefinitionMetricAnomalyConditionModel.ConditionType = core.StringPtr("less_than_usual")

		model := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
		model.Condition = apisAlertDefinitionMetricAnomalyConditionModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
	apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
	apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
	apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)
	apisAlertDefinitionMetricAnomalyConditionModel["condition_type"] = "less_than_usual"

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionMetricAnomalyCondition) {
		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		model := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
		model.Threshold = core.Float64Ptr(float64(10.0))
		model.ForOverPct = core.Int64Ptr(int64(20))
		model.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		model.MinNonNullValuesPct = core.Int64Ptr(int64(10))
		model.ConditionType = core.StringPtr("less_than_usual")

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	model := make(map[string]interface{})
	model["threshold"] = float64(10.0)
	model["for_over_pct"] = int(20)
	model["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	model["min_non_null_values_pct"] = int(10)
	model["condition_type"] = "less_than_usual"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionMetricAnomalyCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsNewValueType) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
		apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

		apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
		apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
		apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

		apisAlertDefinitionLogsNewValueRuleModel := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
		apisAlertDefinitionLogsNewValueRuleModel.Condition = apisAlertDefinitionLogsNewValueConditionModel

		model := new(logsv0.ApisAlertDefinitionLogsNewValueType)
		model.LogsFilter = apisAlertDefinitionLogsFilterModel
		model.Rules = []logsv0.ApisAlertDefinitionLogsNewValueRule{*apisAlertDefinitionLogsNewValueRuleModel}
		model.NotificationPayloadFilter = []string{"obj.field"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

	apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
	apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

	apisAlertDefinitionLogsNewValueRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsNewValueConditionModel}

	model := make(map[string]interface{})
	model["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["rules"] = []interface{}{apisAlertDefinitionLogsNewValueRuleModel}
	model["notification_payload_filter"] = []interface{}{"obj.field"}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueRule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsNewValueRule) {
		apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
		apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

		apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
		apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
		apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

		model := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
		model.Condition = apisAlertDefinitionLogsNewValueConditionModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

	apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
	apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionLogsNewValueConditionModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsNewValueCondition) {
		apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
		apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

		model := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
		model.KeypathToTrack = core.StringPtr("metadata.field")
		model.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

	model := make(map[string]interface{})
	model["keypath_to_track"] = "metadata.field"
	model["time_window"] = []interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueTimeWindow(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsNewValueTimeWindow) {
		model := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
		model.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["logs_new_value_time_window_specific_value"] = "months_3"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsNewValueTimeWindow(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountType(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsUniqueCountType) {
		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
		apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
		apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
		apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

		apisAlertDefinitionLogsUniqueCountRuleModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
		apisAlertDefinitionLogsUniqueCountRuleModel.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

		model := new(logsv0.ApisAlertDefinitionLogsUniqueCountType)
		model.LogsFilter = apisAlertDefinitionLogsFilterModel
		model.Rules = []logsv0.ApisAlertDefinitionLogsUniqueCountRule{*apisAlertDefinitionLogsUniqueCountRuleModel}
		model.NotificationPayloadFilter = []string{"obj.field"}
		model.MaxUniqueCountPerGroupByKey = core.StringPtr("100")
		model.UniqueCountKeypath = core.StringPtr("obj.field")

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
	apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

	apisAlertDefinitionLogsUniqueCountRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueCountRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

	model := make(map[string]interface{})
	model["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	model["rules"] = []interface{}{apisAlertDefinitionLogsUniqueCountRuleModel}
	model["notification_payload_filter"] = []interface{}{"obj.field"}
	model["max_unique_count_per_group_by_key"] = "100"
	model["unique_count_keypath"] = "obj.field"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountType(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountRule(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsUniqueCountRule) {
		apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
		apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
		apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
		apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

		model := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
		model.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
	apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

	model := make(map[string]interface{})
	model["condition"] = []interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountCondition(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsUniqueCountCondition) {
		apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
		apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

		model := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
		model.MaxUniqueCount = core.StringPtr("100")
		model.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

	model := make(map[string]interface{})
	model["max_unique_count"] = "100"
	model["time_window"] = []interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueCountCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueValueTimeWindow(t *testing.T) {
	checkResult := func(result *logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow) {
		model := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
		model.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["logs_unique_value_time_window_specific_value"] = "hours_36"

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToApisAlertDefinitionLogsUniqueValueTimeWindow(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototype(t *testing.T) {
	checkResult := func(result logsv0.AlertDefinitionPrototypeIntf) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsImmediateTypeModel := new(logsv0.ApisAlertDefinitionLogsImmediateType)
		apisAlertDefinitionLogsImmediateTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsImmediateTypeModel.NotificationPayloadFilter = []string{"obj.field"}

		model := new(logsv0.AlertDefinitionPrototype)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsImmediate = apisAlertDefinitionLogsImmediateTypeModel
		// model.LogsThreshold = apisAlertDefinitionLogsThresholdTypeModel
		// model.LogsRatioThreshold = apisAlertDefinitionLogsRatioThresholdTypeModel
		// model.LogsTimeRelativeThreshold = apisAlertDefinitionLogsTimeRelativeThresholdTypeModel
		// model.MetricThreshold = apisAlertDefinitionMetricThresholdTypeModel
		// model.Flow = apisAlertDefinitionFlowTypeModel
		// model.LogsAnomaly = apisAlertDefinitionLogsAnomalyTypeModel
		// model.MetricAnomaly = apisAlertDefinitionMetricAnomalyTypeModel
		// model.LogsNewValue = apisAlertDefinitionLogsNewValueTypeModel
		// model.LogsUniqueCount = apisAlertDefinitionLogsUniqueCountTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsImmediateTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsImmediateTypeModel["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsImmediateTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_immediate"] = []interface{}{apisAlertDefinitionLogsImmediateTypeModel}
	// model["logs_threshold"] = []interface{}{apisAlertDefinitionLogsThresholdTypeModel}
	// model["logs_ratio_threshold"] = []interface{}{apisAlertDefinitionLogsRatioThresholdTypeModel}
	// model["logs_time_relative_threshold"] = []interface{}{apisAlertDefinitionLogsTimeRelativeThresholdTypeModel}
	// model["metric_threshold"] = []interface{}{apisAlertDefinitionMetricThresholdTypeModel}
	// model["flow"] = []interface{}{apisAlertDefinitionFlowTypeModel}
	// model["logs_anomaly"] = []interface{}{apisAlertDefinitionLogsAnomalyTypeModel}
	// model["metric_anomaly"] = []interface{}{apisAlertDefinitionMetricAnomalyTypeModel}
	// model["logs_new_value"] = []interface{}{apisAlertDefinitionLogsNewValueTypeModel}
	// model["logs_unique_count"] = []interface{}{apisAlertDefinitionLogsUniqueCountTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsImmediateTypeModel := new(logsv0.ApisAlertDefinitionLogsImmediateType)
		apisAlertDefinitionLogsImmediateTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsImmediateTypeModel.NotificationPayloadFilter = []string{"obj.field"}

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsImmediate = apisAlertDefinitionLogsImmediateTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsImmediateTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsImmediateTypeModel["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsImmediateTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_immediate"] = []interface{}{apisAlertDefinitionLogsImmediateTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
		apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		apisAlertDefinitionLogsThresholdConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionLogsThresholdRuleModel := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
		apisAlertDefinitionLogsThresholdRuleModel.Condition = apisAlertDefinitionLogsThresholdConditionModel
		apisAlertDefinitionLogsThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

		apisAlertDefinitionLogsThresholdTypeModel := new(logsv0.ApisAlertDefinitionLogsThresholdType)
		apisAlertDefinitionLogsThresholdTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		apisAlertDefinitionLogsThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsThresholdRule{*apisAlertDefinitionLogsThresholdRuleModel}
		apisAlertDefinitionLogsThresholdTypeModel.NotificationPayloadFilter = []string{"obj.field"}
		apisAlertDefinitionLogsThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsThreshold = apisAlertDefinitionLogsThresholdTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	apisAlertDefinitionLogsThresholdConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionLogsThresholdRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsThresholdRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsThresholdConditionModel}
	apisAlertDefinitionLogsThresholdRuleModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	apisAlertDefinitionLogsThresholdTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsThresholdTypeModel["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsThresholdTypeModel["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	apisAlertDefinitionLogsThresholdTypeModel["rules"] = []interface{}{apisAlertDefinitionLogsThresholdRuleModel}
	apisAlertDefinitionLogsThresholdTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}
	apisAlertDefinitionLogsThresholdTypeModel["evaluation_delay_ms"] = int(60000)

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_threshold"] = []interface{}{apisAlertDefinitionLogsThresholdTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
		apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
		apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel
		apisAlertDefinitionLogsRatioConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionLogsRatioRulesModel := new(logsv0.ApisAlertDefinitionLogsRatioRules)
		apisAlertDefinitionLogsRatioRulesModel.Condition = apisAlertDefinitionLogsRatioConditionModel
		apisAlertDefinitionLogsRatioRulesModel.Override = apisAlertDefinitionAlertDefOverrideModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		apisAlertDefinitionLogsRatioThresholdTypeModel := new(logsv0.ApisAlertDefinitionLogsRatioThresholdType)
		apisAlertDefinitionLogsRatioThresholdTypeModel.Numerator = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsRatioThresholdTypeModel.NumeratorAlias = core.StringPtr("numerator_alias")
		apisAlertDefinitionLogsRatioThresholdTypeModel.Denominator = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsRatioThresholdTypeModel.DenominatorAlias = core.StringPtr("denominator_alias")
		apisAlertDefinitionLogsRatioThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsRatioRules{*apisAlertDefinitionLogsRatioRulesModel}
		apisAlertDefinitionLogsRatioThresholdTypeModel.NotificationPayloadFilter = []string{"obj.field"}
		apisAlertDefinitionLogsRatioThresholdTypeModel.GroupByFor = core.StringPtr("denumerator_only")
		apisAlertDefinitionLogsRatioThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		apisAlertDefinitionLogsRatioThresholdTypeModel.IgnoreInfinity = core.BoolPtr(true)
		apisAlertDefinitionLogsRatioThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsRatioThreshold = apisAlertDefinitionLogsRatioThresholdTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
	apisAlertDefinitionLogsRatioConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}
	apisAlertDefinitionLogsRatioConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionLogsRatioRulesModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioRulesModel["condition"] = []interface{}{apisAlertDefinitionLogsRatioConditionModel}
	apisAlertDefinitionLogsRatioRulesModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	apisAlertDefinitionLogsRatioThresholdTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsRatioThresholdTypeModel["numerator"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsRatioThresholdTypeModel["numerator_alias"] = "numerator_alias"
	apisAlertDefinitionLogsRatioThresholdTypeModel["denominator"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsRatioThresholdTypeModel["denominator_alias"] = "denominator_alias"
	apisAlertDefinitionLogsRatioThresholdTypeModel["rules"] = []interface{}{apisAlertDefinitionLogsRatioRulesModel}
	apisAlertDefinitionLogsRatioThresholdTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}
	apisAlertDefinitionLogsRatioThresholdTypeModel["group_by_for"] = "denumerator_only"
	apisAlertDefinitionLogsRatioThresholdTypeModel["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	apisAlertDefinitionLogsRatioThresholdTypeModel["ignore_infinity"] = true
	apisAlertDefinitionLogsRatioThresholdTypeModel["evaluation_delay_ms"] = int(60000)

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_ratio_threshold"] = []interface{}{apisAlertDefinitionLogsRatioThresholdTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
		apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")
		apisAlertDefinitionLogsTimeRelativeConditionModel.ConditionType = core.StringPtr("less_than")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionLogsTimeRelativeRuleModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
		apisAlertDefinitionLogsTimeRelativeRuleModel.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
		apisAlertDefinitionLogsTimeRelativeRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType)
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsTimeRelativeRule{*apisAlertDefinitionLogsTimeRelativeRuleModel}
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.IgnoreInfinity = core.BoolPtr(true)
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.NotificationPayloadFilter = []string{"obj.field"}
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsTimeRelativeThreshold = apisAlertDefinitionLogsTimeRelativeThresholdTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"
	apisAlertDefinitionLogsTimeRelativeConditionModel["condition_type"] = "less_than"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionLogsTimeRelativeRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeRelativeRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
	apisAlertDefinitionLogsTimeRelativeRuleModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["rules"] = []interface{}{apisAlertDefinitionLogsTimeRelativeRuleModel}
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["ignore_infinity"] = true
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["evaluation_delay_ms"] = int(60000)

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_time_relative_threshold"] = []interface{}{apisAlertDefinitionLogsTimeRelativeThresholdTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
		apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
		apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
		apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
		apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		apisAlertDefinitionMetricThresholdConditionModel.ConditionType = core.StringPtr("less_than_or_equals")

		apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
		apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

		apisAlertDefinitionMetricThresholdRuleModel := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
		apisAlertDefinitionMetricThresholdRuleModel.Condition = apisAlertDefinitionMetricThresholdConditionModel
		apisAlertDefinitionMetricThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

		apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
		apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
		apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

		apisAlertDefinitionMetricMissingValuesModel := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
		apisAlertDefinitionMetricMissingValuesModel.ReplaceWithZero = core.BoolPtr(true)

		apisAlertDefinitionMetricThresholdTypeModel := new(logsv0.ApisAlertDefinitionMetricThresholdType)
		apisAlertDefinitionMetricThresholdTypeModel.MetricFilter = apisAlertDefinitionMetricFilterModel
		apisAlertDefinitionMetricThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionMetricThresholdRule{*apisAlertDefinitionMetricThresholdRuleModel}
		apisAlertDefinitionMetricThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
		apisAlertDefinitionMetricThresholdTypeModel.MissingValues = apisAlertDefinitionMetricMissingValuesModel
		apisAlertDefinitionMetricThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.MetricThreshold = apisAlertDefinitionMetricThresholdTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
	apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
	apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
	apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
	apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	apisAlertDefinitionMetricThresholdConditionModel["condition_type"] = "less_than_or_equals"

	apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

	apisAlertDefinitionMetricThresholdRuleModel := make(map[string]interface{})
	apisAlertDefinitionMetricThresholdRuleModel["condition"] = []interface{}{apisAlertDefinitionMetricThresholdConditionModel}
	apisAlertDefinitionMetricThresholdRuleModel["override"] = []interface{}{apisAlertDefinitionAlertDefOverrideModel}

	apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
	apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
	apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

	apisAlertDefinitionMetricMissingValuesModel := make(map[string]interface{})
	apisAlertDefinitionMetricMissingValuesModel["replace_with_zero"] = true

	apisAlertDefinitionMetricThresholdTypeModel := make(map[string]interface{})
	apisAlertDefinitionMetricThresholdTypeModel["metric_filter"] = []interface{}{apisAlertDefinitionMetricFilterModel}
	apisAlertDefinitionMetricThresholdTypeModel["rules"] = []interface{}{apisAlertDefinitionMetricThresholdRuleModel}
	apisAlertDefinitionMetricThresholdTypeModel["undetected_values_management"] = []interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
	apisAlertDefinitionMetricThresholdTypeModel["missing_values"] = []interface{}{apisAlertDefinitionMetricMissingValuesModel}
	apisAlertDefinitionMetricThresholdTypeModel["evaluation_delay_ms"] = int(60000)

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["metric_threshold"] = []interface{}{apisAlertDefinitionMetricThresholdTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

		apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
		apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
		apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

		apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
		apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

		apisAlertDefinitionFlowStagesModel := new(logsv0.ApisAlertDefinitionFlowStages)
		apisAlertDefinitionFlowStagesModel.TimeframeMs = core.StringPtr("60000")
		apisAlertDefinitionFlowStagesModel.TimeframeType = core.StringPtr("up_to")
		apisAlertDefinitionFlowStagesModel.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

		apisAlertDefinitionFlowTypeModel := new(logsv0.ApisAlertDefinitionFlowType)
		apisAlertDefinitionFlowTypeModel.Stages = []logsv0.ApisAlertDefinitionFlowStages{*apisAlertDefinitionFlowStagesModel}
		apisAlertDefinitionFlowTypeModel.EnforceSuppression = core.BoolPtr(true)

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.Flow = apisAlertDefinitionFlowTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

	apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
	apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

	apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesGroupsModel["groups"] = []interface{}{apisAlertDefinitionFlowStagesGroupModel}

	apisAlertDefinitionFlowStagesModel := make(map[string]interface{})
	apisAlertDefinitionFlowStagesModel["timeframe_ms"] = "60000"
	apisAlertDefinitionFlowStagesModel["timeframe_type"] = "up_to"
	apisAlertDefinitionFlowStagesModel["flow_stages_groups"] = []interface{}{apisAlertDefinitionFlowStagesGroupsModel}

	apisAlertDefinitionFlowTypeModel := make(map[string]interface{})
	apisAlertDefinitionFlowTypeModel["stages"] = []interface{}{apisAlertDefinitionFlowStagesModel}
	apisAlertDefinitionFlowTypeModel["enforce_suppression"] = true

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["flow"] = []interface{}{apisAlertDefinitionFlowTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
		apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
		apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel
		apisAlertDefinitionLogsAnomalyConditionModel.ConditionType = core.StringPtr("more_than_usual_or_unspecified")

		apisAlertDefinitionLogsAnomalyRuleModel := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
		apisAlertDefinitionLogsAnomalyRuleModel.Condition = apisAlertDefinitionLogsAnomalyConditionModel

		apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
		apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

		apisAlertDefinitionLogsAnomalyTypeModel := new(logsv0.ApisAlertDefinitionLogsAnomalyType)
		apisAlertDefinitionLogsAnomalyTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsAnomalyTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsAnomalyRule{*apisAlertDefinitionLogsAnomalyRuleModel}
		apisAlertDefinitionLogsAnomalyTypeModel.NotificationPayloadFilter = []string{"obj.field"}
		apisAlertDefinitionLogsAnomalyTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))
		apisAlertDefinitionLogsAnomalyTypeModel.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsAnomaly = apisAlertDefinitionLogsAnomalyTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
	apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsTimeWindowModel}
	apisAlertDefinitionLogsAnomalyConditionModel["condition_type"] = "more_than_usual_or_unspecified"

	apisAlertDefinitionLogsAnomalyRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsAnomalyRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

	apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

	apisAlertDefinitionLogsAnomalyTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsAnomalyTypeModel["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsAnomalyTypeModel["rules"] = []interface{}{apisAlertDefinitionLogsAnomalyRuleModel}
	apisAlertDefinitionLogsAnomalyTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}
	apisAlertDefinitionLogsAnomalyTypeModel["evaluation_delay_ms"] = int(60000)
	apisAlertDefinitionLogsAnomalyTypeModel["anomaly_alert_settings"] = []interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_anomaly"] = []interface{}{apisAlertDefinitionLogsAnomalyTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
		apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

		apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
		apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
		apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
		apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
		apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
		apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))
		apisAlertDefinitionMetricAnomalyConditionModel.ConditionType = core.StringPtr("less_than_usual")

		apisAlertDefinitionMetricAnomalyRuleModel := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
		apisAlertDefinitionMetricAnomalyRuleModel.Condition = apisAlertDefinitionMetricAnomalyConditionModel

		apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
		apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

		apisAlertDefinitionMetricAnomalyTypeModel := new(logsv0.ApisAlertDefinitionMetricAnomalyType)
		apisAlertDefinitionMetricAnomalyTypeModel.MetricFilter = apisAlertDefinitionMetricFilterModel
		apisAlertDefinitionMetricAnomalyTypeModel.Rules = []logsv0.ApisAlertDefinitionMetricAnomalyRule{*apisAlertDefinitionMetricAnomalyRuleModel}
		apisAlertDefinitionMetricAnomalyTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))
		apisAlertDefinitionMetricAnomalyTypeModel.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.MetricAnomaly = apisAlertDefinitionMetricAnomalyTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
	apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

	apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
	apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
	apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
	apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []interface{}{apisAlertDefinitionMetricTimeWindowModel}
	apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)
	apisAlertDefinitionMetricAnomalyConditionModel["condition_type"] = "less_than_usual"

	apisAlertDefinitionMetricAnomalyRuleModel := make(map[string]interface{})
	apisAlertDefinitionMetricAnomalyRuleModel["condition"] = []interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

	apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

	apisAlertDefinitionMetricAnomalyTypeModel := make(map[string]interface{})
	apisAlertDefinitionMetricAnomalyTypeModel["metric_filter"] = []interface{}{apisAlertDefinitionMetricFilterModel}
	apisAlertDefinitionMetricAnomalyTypeModel["rules"] = []interface{}{apisAlertDefinitionMetricAnomalyRuleModel}
	apisAlertDefinitionMetricAnomalyTypeModel["evaluation_delay_ms"] = int(60000)
	apisAlertDefinitionMetricAnomalyTypeModel["anomaly_alert_settings"] = []interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["metric_anomaly"] = []interface{}{apisAlertDefinitionMetricAnomalyTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
		apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

		apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
		apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
		apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

		apisAlertDefinitionLogsNewValueRuleModel := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
		apisAlertDefinitionLogsNewValueRuleModel.Condition = apisAlertDefinitionLogsNewValueConditionModel

		apisAlertDefinitionLogsNewValueTypeModel := new(logsv0.ApisAlertDefinitionLogsNewValueType)
		apisAlertDefinitionLogsNewValueTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsNewValueTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsNewValueRule{*apisAlertDefinitionLogsNewValueRuleModel}
		apisAlertDefinitionLogsNewValueTypeModel.NotificationPayloadFilter = []string{"obj.field"}

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsNewValue = apisAlertDefinitionLogsNewValueTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

	apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
	apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

	apisAlertDefinitionLogsNewValueRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsNewValueConditionModel}

	apisAlertDefinitionLogsNewValueTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsNewValueTypeModel["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsNewValueTypeModel["rules"] = []interface{}{apisAlertDefinitionLogsNewValueRuleModel}
	apisAlertDefinitionLogsNewValueTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_new_value"] = []interface{}{apisAlertDefinitionLogsNewValueTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount(t *testing.T) {
	checkResult := func(result *logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount) {
		apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
		apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
		apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
		apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
		apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

		apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
		apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

		apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
		apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

		apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
		apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
		apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

		apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
		apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
		apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
		apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

		apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
		apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
		apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:\"error\"")
		apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

		apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
		apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

		apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
		apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

		apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
		apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
		apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

		apisAlertDefinitionLogsUniqueCountRuleModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
		apisAlertDefinitionLogsUniqueCountRuleModel.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

		apisAlertDefinitionLogsUniqueCountTypeModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountType)
		apisAlertDefinitionLogsUniqueCountTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
		apisAlertDefinitionLogsUniqueCountTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsUniqueCountRule{*apisAlertDefinitionLogsUniqueCountRuleModel}
		apisAlertDefinitionLogsUniqueCountTypeModel.NotificationPayloadFilter = []string{"obj.field"}
		apisAlertDefinitionLogsUniqueCountTypeModel.MaxUniqueCountPerGroupByKey = core.StringPtr("100")
		apisAlertDefinitionLogsUniqueCountTypeModel.UniqueCountKeypath = core.StringPtr("obj.field")

		model := new(logsv0.AlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount)
		model.Name = core.StringPtr("Unique count alert")
		model.Description = core.StringPtr("Example of unique count alert from terraform")
		model.Enabled = core.BoolPtr(true)
		model.Priority = core.StringPtr("p1")
		model.ActiveOn = apisAlertDefinitionActivityScheduleModel
		model.Type = core.StringPtr("flow")
		model.GroupByKeys = []string{"key1", "key2"}
		model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
		model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
		model.EntityLabels = map[string]string{"key1": "testString"}
		model.PhantomMode = core.BoolPtr(false)
		model.Deleted = core.BoolPtr(false)
		model.LogsUniqueCount = apisAlertDefinitionLogsUniqueCountTypeModel

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
	apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
	apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

	apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
	apisAlertDefinitionActivityScheduleModel["day_of_week"] = []interface{}{"sunday"}
	apisAlertDefinitionActivityScheduleModel["start_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}
	apisAlertDefinitionActivityScheduleModel["end_time"] = []interface{}{apisAlertDefinitionTimeOfDayModel}

	apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

	apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
	apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

	apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
	apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []interface{}{apisAlertDefinitionIntegrationTypeModel}
	apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

	apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
	apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []interface{}{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
	apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
	apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

	apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
	apisAlertDefinitionLabelFiltersModel["application_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []interface{}{apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel["severities"] = []interface{}{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:\"error\""
	apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []interface{}{apisAlertDefinitionLabelFiltersModel}

	apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
	apisAlertDefinitionLogsFilterModel["simple_filter"] = []interface{}{apisAlertDefinitionLogsSimpleFilterModel}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

	apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
	apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

	apisAlertDefinitionLogsUniqueCountRuleModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueCountRuleModel["condition"] = []interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

	apisAlertDefinitionLogsUniqueCountTypeModel := make(map[string]interface{})
	apisAlertDefinitionLogsUniqueCountTypeModel["logs_filter"] = []interface{}{apisAlertDefinitionLogsFilterModel}
	apisAlertDefinitionLogsUniqueCountTypeModel["rules"] = []interface{}{apisAlertDefinitionLogsUniqueCountRuleModel}
	apisAlertDefinitionLogsUniqueCountTypeModel["notification_payload_filter"] = []interface{}{"obj.field"}
	apisAlertDefinitionLogsUniqueCountTypeModel["max_unique_count_per_group_by_key"] = "100"
	apisAlertDefinitionLogsUniqueCountTypeModel["unique_count_keypath"] = "obj.field"

	model := make(map[string]interface{})
	model["name"] = "Unique count alert"
	model["description"] = "Example of unique count alert from terraform"
	model["enabled"] = true
	model["priority"] = "p1"
	model["active_on"] = []interface{}{apisAlertDefinitionActivityScheduleModel}
	model["type"] = "flow"
	model["group_by_keys"] = []interface{}{"key1", "key2"}
	model["incidents_settings"] = []interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
	model["notification_group"] = []interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
	model["entity_labels"] = map[string]interface{}{"key1": "testString"}
	model["phantom_mode"] = false
	model["deleted"] = false
	model["logs_unique_count"] = []interface{}{apisAlertDefinitionLogsUniqueCountTypeModel}

	result, err := logs.ResourceIbmLogsAlertDefinitionMapToAlertDefinitionPrototypeApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount(model)
	assert.Nil(t, err)
	checkResult(result)
}
