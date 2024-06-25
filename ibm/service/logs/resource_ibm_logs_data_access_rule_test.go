// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsDataAccessRuleBasic(t *testing.T) {
	var conf logsv0.DataAccessRule
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	defaultExpression := fmt.Sprintf("tf_default_expression_%d", acctest.RandIntRange(10, 100))
	displayNameUpdate := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	defaultExpressionUpdate := fmt.Sprintf("tf_default_expression_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDataAccessRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfigBasic(displayName, defaultExpression),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDataAccessRuleExists("ibm_logs_data_access_rule.logs_data_access_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayName),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpression),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfigBasic(displayNameUpdate, defaultExpressionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayNameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpressionUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsDataAccessRuleAllArgs(t *testing.T) {
	var conf logsv0.DataAccessRule
	displayName := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	defaultExpression := fmt.Sprintf("tf_default_expression_%d", acctest.RandIntRange(10, 100))
	displayNameUpdate := fmt.Sprintf("tf_display_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	defaultExpressionUpdate := fmt.Sprintf("tf_default_expression_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDataAccessRuleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfig(displayName, description, defaultExpression),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDataAccessRuleExists("ibm_logs_data_access_rule.logs_data_access_rule_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayName),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpression),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDataAccessRuleConfig(displayNameUpdate, descriptionUpdate, defaultExpressionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "display_name", displayNameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_data_access_rule.logs_data_access_rule_instance", "default_expression", defaultExpressionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_data_access_rule.logs_data_access_rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsDataAccessRuleConfigBasic(displayName string, defaultExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
			display_name = "%s"
			filters {
				entity_type = "logs"
				expression = "true"
			}
			default_expression = "%s"
		}
	`, displayName, defaultExpression)
}

func testAccCheckIbmLogsDataAccessRuleConfig(displayName string, description string, defaultExpression string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
			display_name = "%s"
			description = "%s"
			filters {
				entity_type = "logs"
				expression = "true"
			}
			default_expression = "%s"
			instance_id = "%s"
			region      = "%s"
		}
	`, displayName, description, defaultExpression, acc.LogsInstanceId, acc.LogsInstanceRegion)
}

func testAccCheckIbmLogsDataAccessRuleExists(n string, obj logsv0.DataAccessRule) resource.TestCheckFunc {

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

		listDataAccessRulesOptions := &logsv0.ListDataAccessRulesOptions{}

		listDataAccessRulesOptions.SetID(resourceID)

		dataAccessRule, _, err := logsClient.ListDataAccessRules(listDataAccessRulesOptions)
		if err != nil {
			return err
		}

		obj = *dataAccessRule
		return nil
	}
}

func testAccCheckIbmLogsDataAccessRuleDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_data_access_rule" {
			continue
		}

		listDataAccessRulesOptions := &logsv0.ListDataAccessRulesOptions{}

		listDataAccessRulesOptions.SetID(rs.Primary.ID)

		// niranjan
		// Try to find the key
		_, response, err := logsClient.ListDataAccessRules(listDataAccessRulesOptions)

		if err == nil {
			return fmt.Errorf("logs_data_access_rule still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_data_access_rule (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmLogsDataAccessRuleDataAccessRuleFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["entity_type"] = "logs"
		model["expression"] = "true"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.DataAccessRuleFilter)
	model.EntityType = core.StringPtr("logs")
	model.Expression = core.StringPtr("true")

	result, err := logs.ResourceIbmLogsDataAccessRuleDataAccessRuleFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsDataAccessRuleMapToDataAccessRuleFilter(t *testing.T) {
	checkResult := func(result *logsv0.DataAccessRuleFilter) {
		model := new(logsv0.DataAccessRuleFilter)
		model.EntityType = core.StringPtr("logs")
		model.Expression = core.StringPtr("true")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["entity_type"] = "logs"
	model["expression"] = "true"

	result, err := logs.ResourceIbmLogsDataAccessRuleMapToDataAccessRuleFilter(model)
	assert.Nil(t, err)
	checkResult(result)
}
