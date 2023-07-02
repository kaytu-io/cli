package config

// Write the route as Key. The value is the name of the cli command.
// use - instead of / in the route and remove : or {} signs.
var UrlNames = map[string]string{
	// benchmarks
	"get-inventory-api-v1-accounts-resource-count":        "accounts-resource-count",
	"get-inventory-api-v1-resources-distribution":         "resources-distribution",
	"get-inventory-api-v1-resources-top-growing-accounts": "top-growing-accounts-by-resources",
	"get-inventory-api-v1-services-distribution":          "services-distribution",
	"get-inventory-api-v2-services-summary":               "services-summary",
	"get-inventory-api-v2-services-summary-service-name":  "service-summary",

	// benchmarks_assignment
	"delete-compliance-api-v1-assignments-benchmark-id-connection-connection-id": "remove-assignment",
	"get-compliance-api-v1-assignments-benchmark-benchmark-id":                   "benchmark-assignments",
	"get-compliance-api-v1-assignments-connection-connection-id":                 "connection-assignments",
	"get-compliance-api-v1-assignments":                                          "list-assignments",
	"post-compliance-api-v1-assignments-benchmark-id-connection-connection-id":   "add-assignment",

	// compliance
	"get-compliance-api-v1-benchmark-benchmark-id-summary":              "benchmark-summary",
	"get-compliance-api-v1-benchmark-benchmark-id-summary-result-trend": "benchmark-summary-trend",
	"get-compliance-api-v1-benchmark-benchmark-id-tree":                 "benchmark-tree",
	"get-compliance-api-v1-benchmarks-benchmark-id":                     "get-benchmark",
	"get-compliance-api-v1-benchmarks-benchmark-id-policies":            "benchmark-policies",
	"get-compliance-api-v1-benchmark-benchmark-id-policies-policy-id":   "get-policy",
	"get-compliance-api-v1-benchmarks":                                  "list-benchmarks",
	"get-compliance-api-v1-benchmarks-summary":                          "benchmarks-summary",
	"get-compliance-api-v1-findings-benchmark-id-field-top-count":       "benchmark-findings-top-count",
	"get-compliance-api-v1-findings-metrics":                            "findings-metrics",
	"get-compliance-api-v1-queries-query-id":                            "get-query",
	"get-schedule-api-v1-benchmark-evaluations":                         "benchmark-evaluations",
	"post-compliance-api-v1-alarms-top":                                 "top-alarms",
	"post-compliance-api-v1-findings":                                   "get-findings",
	"get-compliance-api-v1-benchmarks-policies-policy-id":               "get-policy",

	// connection
	"get-inventory-api-v2-connections-data-connection-id": "get-connection",
	"get-inventory-api-v2-connections-data":               "list-connections",

	// connections
	"get-onboard-api-v1-connections-summary-connection-id": "connection-summary",
	"get-onboard-api-v1-connections-summary":               "connections-summary",

	// cost
	"get-inventory-api-v1-cost-top-accounts": "top-accounts-by-cost",
	"get-inventory-api-v1-cost-top-services": "top-services-by-cost",

	// describe
	"get-schedule-api-v0-compliance-summarizer-trigger":  "trigger-compliance-summerizer",
	"get-schedule-api-v0-compliance-trigger":             "trigger-compliance",
	"get-schedule-api-v0-describe-trigger":               "describe-trigger",
	"get-schedule-api-v0-insight-trigger":                "trigger-insight",
	"get-schedule-api-v0-summarize-trigger":              "trigger-summerize",
	"get-schedule-api-v1-insight-job-job-id":             "get-insight-job",
	"post-schedule-api-v1-describe-resource":             "describe-resource",
	"post-schedule-api-v1-stacks-insight-trigger":        "trigger-stack-insight",
	"put-schedule-api-v1-benchmark-evaluation-trigger":   "trigger-benchmark-evaluation",
	"put-schedule-api-v1-compliance-trigger":             "trigger-compliance",
	"put-schedule-api-v1-describe-trigger-connection-id": "trigger-describe",
	"put-schedule-api-v1-insight-evaluation-trigger":     "trigger-insight",

	// insight
	"get-compliance-api-v2-insights-insight-id":                  "get-insight",
	"get-compliance-api-v2-insights-insight-id-trend":            "get-insight-trend",
	"get-compliance-api-v2-insights-job-job-id":                  "get-insight-job",
	"get-compliance-api-v2-insights":                             "list-insights",
	"get-compliance-api-v1-insight-group-insight-group-id":       "get-insight-group",
	"get-compliance-api-v1-insight-group-insight-group-id-trend": "get-insight-group-trend",
	"get-compliance-api-v1-insight-group":                        "list-insight-groups",
	"get-compliance-api-v1-metadata-insight-insight-id":          "get-insight-metadata",
	"get-inventory-api-v2-insights-insight-id-trend":             "old-get-insight-trend",
	"get-inventory-api-v2-insights":                              "old-list-insights",
	"get-inventory-api-v2-insights-job-job-id":                   "old-get-insight-job",
	"get-inventory-api-v2-insights-insight-id":                   "old-get-insight",

	// insights
	"get-compliance-api-v1-insight-insight-id":       "get-insight",
	"get-compliance-api-v1-insight-insight-id-trend": "get-insight-trend",
	"get-compliance-api-v1-insight":                  "list-insights",
	"get-compliance-api-v1-metadata-insight":         "list-insights-metadata",
	"get-compliance-api-v1-metadata-tag-insight":     "list-insights-metadata-tags",
	"get-compliance-api-v1-metadata-tag-insight-key": "list-insights-metadata-tag",

	// inventory
	"get-inventory-api-v1-resources-count":                "all-resources-count",
	"get-inventory-api-v1-resources-regions":              "regions-by-resources",
	"get-inventory-api-v1-resources-top-regions":          "top-regions-by-resources",
	"get-inventory-api-v2-resources-composition-key":      "resources-composition",
	"get-inventory-api-v2-resources-metric":               "resources-metrics",
	"get-inventory-api-v2-resources-metric-resource-type": "resource-metrics",
	"get-inventory-api-v2-resources-tag-key":              "get-resource-tag",
	"get-inventory-api-v2-resources-tag":                  "list-resource-tags",
	"get-inventory-api-v2-resources-trend":                "resource-trend",
	"get-inventory-api-v2-services-composition-key":       "services-composition",
	"get-inventory-api-v2-services-cost-trend":            "service-cost-trend",
	"get-inventory-api-v2-services-metric":                "service-metrics",
	"get-inventory-api-v2-services-metric-service-name":   "service-metric",
	"get-inventory-api-v2-services-tag-key":               "get-service-tag",
	"get-inventory-api-v2-services-tag":                   "list-service-tags",
	"post-inventory-api-v1-resources-aws":                 "list-aws-resources",
	"post-inventory-api-v1-resources-azure":               "list-azure-resources",
	"post-inventory-api-v1-resources":                     "list-all-resources",
	"post-inventory-api-v1-resources-filters":             "list-resources-by-filters",
	"get-inventory-api-v2-cost-trend":                     "cost-trend",
	"get-inventory-api-v2-cost-metric":                    "cost-metric",
	"get-inventory-api-v2-cost-composition":               "cost-composition",
	"get-inventory-api-v2-resources-regions-trend":        "get-resource-trend-by-region",
	"get-inventory-api-v2-resources-regions-composition":  "get-resource-composition-by-region",
	"get-inventory-api-v2-resources-regions-summary":      "get-resource-summary-by-region",

	// keys
	"delete-auth-api-v1-key-id-delete": "remove-key",
	"get-auth-api-v1-key-id":           "get-key",
	"get-auth-api-v1-keys":             "list-keys",
	"post-auth-api-v1-key-create":      "create-key",
	"post-auth-api-v1-key-id-activate": "activate-key",
	"post-auth-api-v1-key-id-suspend":  "suspend-key",
	"post-auth-api-v1-key-role":        "update-key-role",

	// locations
	"get-inventory-api-v1-locations-connector": "list-connector-locations",

	// metadata
	"get-inventory-api-v2-metadata-resourcetype":               "list-resource-types",
	"get-inventory-api-v2-metadata-resourcetype-resource-type": "get-resource-type",
	"get-inventory-api-v2-metadata-services":                   "list-services",
	"get-inventory-api-v2-metadata-services-service-name":      "get-service",
	"get-metadata-api-v1-metadata-key":                         "get-config-metadata",
	"post-metadata-api-v1-metadata":                            "set-config-metadata",

	// onboard
	"delete-onboard-api-v1-credential-credential-id":           "delete-credential",
	"delete-onboard-api-v1-source-source-id":                   "delete-source",
	"get-onboard-api-v1-catalog-connectors":                    "get-catalog-connectors",
	"get-onboard-api-v1-catalog-metrics":                       "get-catalog-metrics",
	"get-onboard-api-v1-connections-count":                     "get-connections-count",
	"get-onboard-api-v1-connector-connector-name":              "get-connector",
	"get-onboard-api-v1-connectors-connector-name":             "get-connector",
	"get-onboard-api-v1-connectors":                            "get-connectors",
	"get-onboard-api-v1-connector":                             "get-connectors",
	"get-onboard-api-v1-credential-credential-id-healthcheck":  "healthcheck-credential",
	"get-onboard-api-v1-credential-credential-id":              "get-credential",
	"get-onboard-api-v1-credential":                            "list-credentials",
	"get-onboard-api-v1-credential-sources-list":               "list-credentials-sources",
	"get-onboard-api-v1-providers":                             "get-providers",
	"get-onboard-api-v1-providers-types":                       "get-providers-types",
	"get-onboard-api-v1-source-account-account-id":             "get-account-source",
	"get-onboard-api-v1-source-source-id-credentials":          "source-credentials",
	"get-onboard-api-v1-source-source-id":                      "get-source",
	"get-onboard-api-v1-sources":                               "list-sources",
	"get-onboard-api-v1-sources-count":                         "sources-count",
	"post-onboard-api-v1-connections-connection-id-state":      "update-connection-state",
	"post-onboard-api-v1-credential-credential-id-autoonboard": "auto-onboard-credential",
	"post-onboard-api-v1-credential-credential-id-disable":     "disable-credential",
	"post-onboard-api-v1-credential-credential-id-enable":      "enable-credential",
	"post-onboard-api-v1-credential":                           "create-credential",
	"post-onboard-api-v1-source-aws":                           "create-aws-source",
	"post-onboard-api-v1-source-azure":                         "create-azure-source",
	"post-onboard-api-v1-source-source-id-disable":             "disable-source",
	"post-onboard-api-v1-source-source-id-enable":              "enable-source",
	"post-onboard-api-v1-source-source-id-healthcheck":         "healthchek-source",
	"post-onboard-api-v1-sources":                              "get-source",
	"put-onboard-api-v1-credential-credential-id":              "update-credential",
	"put-onboard-api-v1-source-source-id-credentials":          "put-source-credentials",

	// resource
	"post-inventory-api-v1-resource": "create-resource",

	// roles
	"get-auth-api-v1-role-role-name-keys":  "list-role-keys",
	"get-auth-api-v1-role-role-name-users": "list-role-users",
	"get-auth-api-v1-roles-role-name":      "get-role",
	"get-auth-api-v1-roles":                "list-roles",

	// schedule
	"get-schedule-api-v1-compliance-report-last-completed":           "last-compliance-jobs",
	"get-schedule-api-v1-describe-resource-jobs-pending":             "pending-describe-resource-jobs",
	"get-schedule-api-v1-describe-source-jobs-pending":               "pending-describe-source-jobs",
	"get-schedule-api-v1-insight-jobs-pending":                       "pending-insight-jobs",
	"get-schedule-api-v1-resource-type-provider":                     "provider-resource-types",
	"get-schedule-api-v1-sources":                                    "list-sources",
	"get-schedule-api-v1-sources-source-id-jobs-compliance":          "source-compliance-jobs",
	"get-schedule-api-v1-sources-source-id-jobs-describe":            "source-describe-jobs",
	"get-schedule-api-v1-sources-source-id":                          "get-source",
	"get-schedule-api-v1-summarize-jobs-pending":                     "pending-summarize-jobs",
	"post-schedule-api-v1-sources-source-id-jobs-compliance-refresh": "refresh-source-compliance-jobs",
	"post-schedule-api-v1-sources-source-id-jobs-describe-refresh":   "refresh-source-describe-jobs",

	// smart_query
	"get-inventory-api-v1-query-count":     "get-queries-count",
	"get-inventory-api-v1-query":           "get-queries",
	"post-inventory-api-v1-query-query-id": "get-query",

	// stack
	"delete-schedule-api-v1-stacks-stack-id":          "delete-stack",
	"get-schedule-api-v1-stacks-findings-job-id":      "stack-findings",
	"get-schedule-api-v1-stacks-stack-id":             "get-stack",
	"get-schedule-api-v1-stacks":                      "list-stacks",
	"get-schedule-api-v1-stacks-resource":             "resource-stacks",
	"get-schedule-api-v1-stacks-stack-id-insights":    "stack-insights",
	"post-schedule-api-v1-stacks-benchmark-trigger":   "trigger-stack-benchmark",
	"post-schedule-api-v1-stacks-create":              "create-stack",
	"post-schedule-api-v1-stacks-stack-id-findings":   "stack-findings",
	"get-schedule-api-v1-stacks-stack-id-insight":     "stack-insights",
	"get-schedule-api-v1-stacks-resource-resource-id": "resource-stacks",
	"post-schedule-api-v1-stacks-describer-trigger":   "trigger-stack-describer",

	// users
	"delete-auth-api-v1-user-invite":                    "delete-user-invite",
	"delete-auth-api-v1-user-role-binding":              "delete-user-role-binding",
	"get-auth-api-v1-user-role-bindings":                "list-user-role-bindings",
	"get-auth-api-v1-user-user-id":                      "get-user",
	"get-auth-api-v1-users":                             "list-users",
	"get-auth-api-v1-user-user-id-workspace-membership": "list-user-workspace-memberships",
	"post-auth-api-v1-user-invite":                      "invite-user",
	"get-auth-api-v1-workspace-role-bindings":           "list-workspace-role-bindings",
	"put-auth-api-v1-user-role-binding":                 "update-user-role-binding",

	// workspace
	"delete-workspace-api-v1-workspace-workspace-id":            "delete-workspace",
	"get-workspace-api-v1-workspace-workspace-id":               "get-workspace",
	"get-workspace-api-v1-workspaces":                           "list-workspaces",
	"post-workspace-api-v1-workspace":                           "create-workspace",
	"get-workspace-api-v1-workspaces-byid-workspace-id":         "get-workspace-by-id",
	"get-workspace-api-v1-workspaces-limits-byid-workspace-id":  "get-workspace-limits-by-id",
	"get-workspace-api-v1-workspaces-limits-workspace-name":     "get-workspace-limits",
	"post-workspace-api-v1-workspace-workspace-id-name":         "update-workspace-name",
	"post-workspace-api-v1-workspace-workspace-id-organization": "update-workspace-organization",
	"post-workspace-api-v1-workspace-workspace-id-owner":        "update-workspace-owner",
	"post-workspace-api-v1-workspace-workspace-id-resume":       "resume-workspace",
	"post-workspace-api-v1-workspace-workspace-id-suspend":      "suspend-workspace",
	"post-workspace-api-v1-workspace-workspace-id-tier":         "update-workspace-tier",
	"get-workspace-api-v1-workspaces-workspace-id":              "get-workspace",
	"get-workspace-api-v1-workspace-current":                    "get-current-workspace",
}