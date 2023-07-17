package config

// Write the route as Key. The value is the name of the cli command.
// use - instead of / in the route and remove : or {} signs.
var UrlSummary = map[string]string{
	// stack
	//"delete-schedule-api-v1-stacks-stack-id":          "delete-stack",
	//"get-schedule-api-v1-stacks-findings-job-id":      "stack-findings",
	"get-schedule-api-v1-stacks-stack-id": "accountIds,status,tags,sourceType,createdAt,updatedAt",
	"get-schedule-api-v1-stacks":          "stackId,tags,status,updatedAt",
	//"get-schedule-api-v1-stacks-resource":             "resource-stacks",
	"get-schedule-api-v1-stacks-stack-id-insights": "id,shortTitle,totalResultValue,tags,result(details,executedAt,results)",
	//"post-schedule-api-v1-stacks-benchmark-trigger":   "trigger-stack-benchmark",
	//"post-schedule-api-v1-stacks-create":              "create-stack",
	//"post-schedule-api-v1-stacks-stack-id-findings":   "list-stack-findings",
	//"get-schedule-api-v1-stacks-stack-id-insight":     "get-stack-insight",
	//"get-schedule-api-v1-stacks-resource-resource-id": "resource-stacks",
	//"post-schedule-api-v1-stacks-describer-trigger":   "trigger-stack-describer",
}
