// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.91.0-d9755c53-20240605-153412
 */

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsDataAccessRulesDataSourceBasic(t *testing.T) {
	dataAccessRuleDisplayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	dataAccessRuleDefaultExpression := fmt.Sprintf("tf_default_expression_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRulesDataSourceConfigBasic(dataAccessRuleDisplayName, dataAccessRuleDefaultExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.#"),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.display_name", dataAccessRuleDisplayName),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.default_expression", dataAccessRuleDefaultExpression),
				),
			},
		},
	})
}

func TestAccIbmLogsDataAccessRulesDataSourceAllArgs(t *testing.T) {
	dataAccessRuleDisplayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	dataAccessRuleDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	dataAccessRuleDefaultExpression := fmt.Sprintf("tf_default_expression_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRulesDataSourceConfig(dataAccessRuleDisplayName, dataAccessRuleDescription, dataAccessRuleDefaultExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "logs_data_access_rules_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.display_name", dataAccessRuleDisplayName),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.description", dataAccessRuleDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_data_access_rules.logs_data_access_rules_instance", "data_access_rules.0.default_expression", dataAccessRuleDefaultExpression),
				),
			},
		},
	})
}

func testAccCheckIbmLogsDataAccessRulesDataSourceConfigBasic(dataAccessRuleDisplayName string, dataAccessRuleDefaultExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
			display_name = "%s"
			filters {
				entity_type = "logs"
				expression = "true"
			}
			default_expression = "%s"
		}

		data "ibm_logs_data_access_rules" "logs_data_access_rules_instance" {
			logs_data_access_rules_id = ["4f966911-4bda-407e-b069-477394effa59"]
		}
	`, dataAccessRuleDisplayName, dataAccessRuleDefaultExpression)
}

func testAccCheckIbmLogsDataAccessRulesDataSourceConfig(dataAccessRuleDisplayName string, dataAccessRuleDescription string, dataAccessRuleDefaultExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
			display_name = "%s"
			description = "%s"
			filters {
				entity_type = "logs"
				expression = "true"
			}
			default_expression = "%s"
		}

		data "ibm_logs_data_access_rules" "logs_data_access_rules_instance" {
			logs_data_access_rules_id = ["4f966911-4bda-407e-b069-477394effa59"]
		}
	`, dataAccessRuleDisplayName, dataAccessRuleDescription, dataAccessRuleDefaultExpression)
}

func TestDataSourceIbmLogsDataAccessRulesDataAccessRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataAccessRuleFilterModel := make(map[string]interface{})
		dataAccessRuleFilterModel["entity_type"] = "logs"
		dataAccessRuleFilterModel["expression"] = "<v1> foo == 'bar'"

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["display_name"] = "Data Access Rule for group 'users'"
		model["description"] = "Data Access Rule that defines restriction on 'users' group"
		model["filters"] = []map[string]interface{}{dataAccessRuleFilterModel}
		model["default_expression"] = "true"

		assert.Equal(t, result, model)
	}

	dataAccessRuleFilterModel := new(logsv0.DataAccessRuleFilter)
	dataAccessRuleFilterModel.EntityType = core.StringPtr("logs")
	dataAccessRuleFilterModel.Expression = core.StringPtr("<v1> foo == 'bar'")

	model := new(logsv0.DataAccessRule)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.DisplayName = core.StringPtr("Data Access Rule for group 'users'")
	model.Description = core.StringPtr("Data Access Rule that defines restriction on 'users' group")
	model.Filters = []logsv0.DataAccessRuleFilter{*dataAccessRuleFilterModel}
	model.DefaultExpression = core.StringPtr("true")

	result, err := logs.DataSourceIbmLogsDataAccessRulesDataAccessRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsDataAccessRulesDataAccessRuleFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["entity_type"] = "logs"
		model["expression"] = "true"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.DataAccessRuleFilter)
	model.EntityType = core.StringPtr("logs")
	model.Expression = core.StringPtr("true")

	result, err := logs.DataSourceIbmLogsDataAccessRulesDataAccessRuleFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
