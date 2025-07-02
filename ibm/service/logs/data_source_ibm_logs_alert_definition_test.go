// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
*/

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsAlertDefinitionDataSourceBasic(t *testing.T) {
	alertDefinitionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDefinitionType := "logs_immediate_or_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionDataSourceConfigBasic(alertDefinitionName, alertDefinitionType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_alert_definition_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "group_by_keys.#"),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertDefinitionDataSourceAllArgs(t *testing.T) {
	alertDefinitionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDefinitionDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	alertDefinitionEnabled := "false"
	alertDefinitionPriority := "p5_or_unspecified"
	alertDefinitionType := "logs_immediate_or_unspecified"
	alertDefinitionPhantomMode := "true"
	alertDefinitionDeleted := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionDataSourceConfig(alertDefinitionName, alertDefinitionDescription, alertDefinitionEnabled, alertDefinitionPriority, alertDefinitionType, alertDefinitionPhantomMode, alertDefinitionDeleted),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_alert_definition_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "alert_version_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "priority"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "active_on.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "group_by_keys.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "incidents_settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "notification_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "entity_labels.%"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "phantom_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "deleted"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_immediate.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_threshold.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_ratio_threshold.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_time_relative_threshold.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "metric_threshold.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "flow.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_anomaly.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "metric_anomaly.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_new_value.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_unique_count.#"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsAlertDefinitionDataSourceConfigBasic(alertDefinitionName string, alertDefinitionType string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			name = "%s"
			type = "%s"
			group_by_keys = "FIXME"
		}

		data "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			logs_alert_definition_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		}
	`, alertDefinitionName, alertDefinitionType)
}

func testAccCheckIbmLogsAlertDefinitionDataSourceConfig(alertDefinitionName string, alertDefinitionDescription string, alertDefinitionEnabled string, alertDefinitionPriority string, alertDefinitionType string, alertDefinitionPhantomMode string, alertDefinitionDeleted string) string {
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
						lucene_query = "text:"error""
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
						lucene_query = "text:"error""
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
						lucene_query = "text:"error""
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
						lucene_query = "text:"error""
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
						lucene_query = "text:"error""
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
						lucene_query = "text:"error""
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
						lucene_query = "text:"error""
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
						lucene_query = "text:"error""
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

		data "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			logs_alert_definition_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		}
	`, alertDefinitionName, alertDefinitionDescription, alertDefinitionEnabled, alertDefinitionPriority, alertDefinitionType, alertDefinitionPhantomMode, alertDefinitionDeleted)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionActivityScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hours"] = int(14)
		model["minutes"] = int(30)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionTimeOfDay)
	model.Hours = core.Int64Ptr(int64(14))
	model.Minutes = core.Int64Ptr(int64(30))

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionTimeOfDayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["notify_on"] = "triggered_and_resolved"
		model["minutes"] = int(30)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	model.NotifyOn = core.StringPtr("triggered_and_resolved")
	model.Minutes = core.Int64Ptr(int64(30))

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefIncidentSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefNotificationGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefWebhooksSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["integration_id"] = int(123)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionIntegrationType)
	model.IntegrationID = core.Int64Ptr(int64(123))

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["integration_id"] = int(123)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	model.IntegrationID = core.Int64Ptr(int64(123))

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	model := new(logsv0.ApisAlertDefinitionLogsImmediateType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.NotificationPayloadFilter = []string{"obj.field"}

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsImmediateTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	model := new(logsv0.ApisAlertDefinitionLogsFilter)
	model.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		model := make(map[string]interface{})
		model["lucene_query"] = "text:"error""
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
	model.LuceneQuery = core.StringPtr("text:"error"")
	model.LabelFilters = apisAlertDefinitionLabelFiltersModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsSimpleFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = "my-app"
		model["operation"] = "starts_with"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLabelFilterType)
	model.Value = core.StringPtr("my-app")
	model.Operation = core.StringPtr("starts_with")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLabelFilterTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["trigger_undetected_values"] = true
		model["auto_retire_timeframe"] = "hours_24"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	model.TriggerUndetectedValues = core.BoolPtr(true)
	model.AutoRetireTimeframe = core.StringPtr("hours_24")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionUndetectedValuesManagementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsThresholdConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	model.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["priority"] = "p1"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	model.Priority = core.StringPtr("p1")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAlertDefOverrideToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioRulesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_ratio_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	model.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsRatioTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsTimeRelativeConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["promql"] = "avg_over_time(metric_name[5m]) > 10"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricFilter)
	model.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricThresholdConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_specific_value"] = "hours_36"
		model["metric_time_window_dynamic_duration"] = "1h30m"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindow)
	model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")
	model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_dynamic_duration"] = "1h30m"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration)
	model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["replace_with_zero"] = true
		model["min_non_null_values_pct"] = int(80)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValues)
	model.ReplaceWithZero = core.BoolPtr(true)
	model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["replace_with_zero"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
	model.ReplaceWithZero = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["min_non_null_values_pct"] = int(80)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct)
	model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["not"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Not = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsAnomalyConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["percentage_of_deviation"] = float64(10.0)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	model.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionAnomalyAlertSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionMetricAnomalyConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_new_value_time_window_specific_value"] = "months_3"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	model.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsNewValueTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
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
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(t *testing.T) {
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

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueCountConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_unique_value_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	model.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
