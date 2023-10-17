package manual_commands

import (
	"errors"
	"fmt"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/alerting"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var POSTAlertingApiV1RuleCreateCmd = &cobra.Command{
	Use:   "create-rule",
	Short: `create an rule by the specified input`,
	Long:  `create an rule by the specified input`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_alerting_api_v_1_rule_create] : %v", err)
		}

		var resEventType models.GithubComKaytuIoKaytuEnginePkgAlertingAPIEventType

		if insightId := flags.ReadInt64Flag(cmd, "insight-id"); insightId != 0 {
			resEventType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIEventType{
				InsightID: flags.ReadInt64Flag(cmd, "insight-id"),
			}
		} else {
			resEventType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIEventType{
				BenchmarkID: flags.ReadStringFlag(cmd, "benchmark-id"),
			}
		}

		var resScope models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope
		if connectionGroup := flags.ReadStringFlag(cmd, "connection-group"); connectionGroup != "" {
			if connector := flags.ReadStringFlag(cmd, "connector"); connector != "" {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionGroup: connectionGroup,
					Connector:       models.SourceType(flags.ReadStringFlag(cmd, "connector")),
				}
			} else {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionGroup: connectionGroup,
				}
			}
		} else {
			if connector := flags.ReadStringFlag(cmd, "connector"); connector != "" {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionID: flags.ReadStringFlag(cmd, "connection-id"),
					Connector:    models.SourceType(flags.ReadStringFlag(cmd, "connector")),
				}
			} else {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionID: flags.ReadStringFlag(cmd, "connection-id"),
				}
			}
		}

		var operatorRes *models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct
		operator := flags.ReadStringFlag(cmd, "operator")
		if operator != "" {
			operatorRes1 := HandelOperatorInput(operator)
			operatorRes = &operatorRes1
		} else {
			operatorRes = nil
		}

		req := alerting.NewPostAlertingAPIV1RuleCreateParams()
		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgAlertingAPICreateRuleRequest{
			ActionID:  flags.ReadInt64Flag(cmd, "action-id"),
			EventType: &resEventType,
			Scope:     &resScope,
			Operator:  operatorRes,
		})

		resp, err := client.Alerting.PostAlertingAPIV1RuleCreate(req, auth)
		if err != nil {
			return fmt.Errorf("[post_alerting_api_v_1_rule_create] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "post_alerting_api_v_1_rule_create", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_alerting_api_v_1_rule_create] : %v", err)
		}

		return nil
	},
}

var PUTAlertingApiV1RuleUpdateCmd = &cobra.Command{
	Use:   "update-rule",
	Short: `update an rule by the specified input`,
	Long:  `update an rule by the specified input`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[put_alerting_api_v_1_rule_update] : %v", err)
		}

		var resEventType models.GithubComKaytuIoKaytuEnginePkgAlertingAPIEventType
		resEventTypePo := &resEventType
		if benchmarkId := flags.ReadStringFlag(cmd, "benchmark-id"); benchmarkId != "" {
			resEventType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIEventType{
				BenchmarkID: benchmarkId,
			}
			resEventTypePo = &resEventType
		} else if insightId := flags.ReadInt64Flag(cmd, "insight-id"); insightId != 0 {
			resEventType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIEventType{
				InsightID: insightId,
			}
			resEventTypePo = &resEventType
		} else {
			resEventTypePo = nil
		}

		var resScope models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope
		resScopePo := &resScope
		if connectionGroup := flags.ReadStringFlag(cmd, "connection-group"); connectionGroup != "" {
			if connector := flags.ReadStringFlag(cmd, "connector"); connector != "" {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionGroup: connectionGroup,
					Connector:       models.SourceType(flags.ReadStringFlag(cmd, "connector")),
				}
			} else {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionGroup: connectionGroup,
				}
			}
			resScopePo = &resScope
		} else if connectionId := flags.ReadStringFlag(cmd, "connection-id"); connectionId != "" {
			if connector := flags.ReadStringFlag(cmd, "connector"); connector != "" {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionID: connectionId,
					Connector:    models.SourceType(flags.ReadStringFlag(cmd, "connector")),
				}
			} else {
				resScope = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIScope{
					ConnectionID: connectionId,
				}
			}
			resScopePo = &resScope
		} else {
			resScopePo = nil
		}

		var operatorRes *models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct
		operator := flags.ReadStringFlag(cmd, "operator")
		if operator != "" {
			operatorRes1 := HandelOperatorInput(operator)
			operatorRes = &operatorRes1
		} else {
			operatorRes = nil
		}

		req := alerting.NewPutAlertingAPIV1RuleUpdateRuleIDParams()
		req.SetRuleID(flags.ReadStringFlag(cmd, "rule-id"))
		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgAlertingAPIUpdateRuleRequest{
			ActionID:  flags.ReadInt64Flag(cmd, "action-id"),
			EventType: resEventTypePo,
			Scope:     resScopePo,
			Operator:  operatorRes,
		})

		resp, err := client.Alerting.PutAlertingAPIV1RuleUpdateRuleID(req, auth)
		if err != nil {
			return fmt.Errorf("[put_alerting_api_v_1_rule_update] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "put_alerting_api_v_1_rule_update", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_alerting_api_v_1_rule_update] : %v", err)
		}

		return nil
	},
}

func HandelOperatorInput(operator string) models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct {
	var operatorStr models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct
	var conditionType string
	operatorArray := strings.Fields(operator)
	if operatorStr.Condition == nil {
		operatorStr.Condition = &models.GithubComKaytuIoKaytuEnginePkgAlertingAPIConditionStruct{}
	}
	if operatorStr.Condition.Operator == nil {
		operatorStr.Condition.Operator = make([]*models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct, 0)
	}

	for i := 0; i < len(operatorArray); i++ {
		if operatorArray[i] == "(" {
			// value > 100 AND ( value < 110 OR value > 120)

			siOpen := strings.Index(operator, "(")
			isClose := strings.Index(operator, ")")

			operator2 := operator[siOpen+1 : isClose]
			operator3 := strings.Fields(operator2)

			operator = operator[isClose+1:]
			operatorArray = strings.Fields(operator)
			i = -1

			if conditionType != "" {
				operatorStr.Condition.ConditionType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIConditionType(conditionType)
			}

			newOperator := compereParantes(operator3, operator2)
			operatorStr.Condition.Operator = append(operatorStr.Condition.Operator, &newOperator)
			continue
		}
		if operatorArray[i] == "value" {
			if len(operatorArray) < 4 {
				// value > 10
				value, _ := strconv.Atoi(operatorArray[i+2])
				operatorType := FixingOperatorType(operatorArray[i+1])
				operatorStr = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct{
					OperatorInfo: &models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation{
						OperatorType: models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorType(operatorType),
						Value:        int64(value),
					},
				}

				continue
			} else {
				// value > 10 AND value < 122323

				if operatorArray[i] == "value" {
					value, _ := strconv.Atoi(operatorArray[i+2])
					operatorType := FixingOperatorType(operatorArray[i+1])

					op := models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct{
						OperatorInfo: &models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation{
							OperatorType: models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorType(operatorType),
							Value:        int64(value),
						},
					}
					if conditionType != "" {
						operatorStr.Condition.ConditionType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIConditionType(conditionType)
					}
					operatorStr.Condition.Operator = append(operatorStr.Condition.Operator, &op)
				}
			}
		}
		if operatorArray[i] == "and" || operatorArray[i] == "AND" {
			conditionType = "AND"
		}
		if operatorArray[i] == "or" || operatorArray[i] == "OR" {
			conditionType = "OR"
		}
	}
	return operatorStr
}

func compereParantes(textArray []string, textString string) models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct {
	var operatorStr models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct

	if operatorStr.Condition == nil {
		operatorStr.Condition = &models.GithubComKaytuIoKaytuEnginePkgAlertingAPIConditionStruct{}
	}

	var conditionType string
	for i := 0; i < len(textArray); i++ {

		if textArray[i] == "(" {
			siOpen := strings.Index(textString, "(")
			isClose := strings.Index(textString, ")")

			operator2 := textString[siOpen+1 : isClose]
			operator3 := strings.Fields(operator2)

			textString = textString[isClose+1:]
			textArray = strings.Fields(textString)
			i = -1

			if conditionType != "" {
				operatorStr.Condition.ConditionType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIConditionType(conditionType)
			}

			newOperator := compereParantes(operator3, operator2)
			operatorStr.Condition.Operator = append(operatorStr.Condition.Operator, &newOperator)

			continue
		}

		if textArray[i] == "value" {
			// value > 10 AND (value < 100 OR value < 10000)
			// value > 10 AND value < 122323
			value, _ := strconv.Atoi(textArray[i+2])
			operatorType := FixingOperatorType(textArray[i+1])

			op := models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorStruct{
				OperatorInfo: &models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation{
					OperatorType: models.GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorType(operatorType),
					Value:        int64(value),
				},
			}
			operatorStr.Condition.Operator = append(operatorStr.Condition.Operator, &op)

			if conditionType != "" {
				operatorStr.Condition.ConditionType = models.GithubComKaytuIoKaytuEnginePkgAlertingAPIConditionType(conditionType)
			}
		}

		if textArray[i] == "and" || textArray[i] == "AND" {
			conditionType = "AND"
		}
		if textArray[i] == "or" || textArray[i] == "OR" {
			conditionType = "OR"
		}
	}
	return operatorStr
}

func FixingOperatorType(operatorType string) string {
	switch operatorType {
	case ">":
		return "GreaterThan"
	case "<":
		return "LessThan"
	case ">=":
		return "GreaterThanOrEqual"
	case "<=":
		return "LessThanOrEqual"
	case "=":
		return "Equal"
	case "!=":
		return "DoesNotEqual"
	default:
		return ""
	}
}
