package config

// Write the route as Key. The value is the name of the cli command.
// use _ or _ instead of / in the route and remove : or {} signs.
var UrlNames = map[string]string{
	// stack
	"delete_schedule_api_v1_stacks_stack_id":          "delete-stack",
	"get_schedule_api_v1_stacks_findings_job_id":      "stack-findings",
	"get_schedule_api_v1_stacks_stack_id":             "get-stack",
	"get_schedule_api_v1_stacks":                      "list-stacks",
	"get_schedule_api_v1_stacks_resource":             "resource-stacks",
	"get_schedule_api_v1_stacks_stack_id_insights":    "list-stack-insights",
	"post_schedule_api_v1_stacks_benchmark_trigger":   "trigger-stack-benchmark",
	"post_schedule_api_v1_stacks_create":              "create-stack",
	"post_schedule_api_v1_stacks_stack_id_findings":   "list-stack-findings",
	"get_schedule_api_v1_stacks_stack_id_insight":     "get-stack-insight",
	"get_schedule_api_v1_stacks_resource_resource_id": "resource-stacks",
	"post_schedule_api_v1_stacks_describer_trigger":   "trigger-stack-describer",

	// benchmarks
	"get_inventory_api_v1_accounts_resource_count":        "accounts-resource-count",
	"get_inventory_api_v1_resources_distribution":         "resources-distribution",
	"get_inventory_api_v1_resources_top_growing_accounts": "top-growing-accounts-by-resources",
	"get_inventory_api_v1_services_distribution":          "services-distribution",
	"get_inventory_api_v2_services_summary":               "services-summary",
	"get_inventory_api_v2_services_summary_service_name":  "service-summary",

	// benchmarks_assignment
	"delete_compliance_api_v1_assignments_benchmark_id_connection_connection_id": "remove-assignment",
	"get_compliance_api_v1_assignments_benchmark_benchmark_id":                   "benchmark-assignments",
	"get_compliance_api_v1_assignments_connection_connection_id":                 "connection-assignments",
	"get_compliance_api_v1_assignments":                                          "list-assignments",
	"post_compliance_api_v1_assignments_benchmark_id_connection_connection_id":   "add-assignment",
	"delete_compliance_api_v1_assignments_benchmark_id_connection":               "remove-benchmark-assignment",
	"post_compliance_api_v1_assignments_benchmark_id_connection":                 "create-benchmark-assignment",

	// compliance
	"get_compliance_api_v1_benchmark_benchmark_id_summary":               "benchmark-summary",
	"get_compliance_api_v1_benchmark_benchmark_id_summary_result_trend":  "benchmark-summary-trend",
	"get_compliance_api_v1_benchmark_benchmark_id_tree":                  "benchmark-tree",
	"get_compliance_api_v1_benchmarks_benchmark_id":                      "get-benchmark",
	"get_compliance_api_v1_benchmarks_benchmark_id_policies":             "benchmark-policies",
	"get_compliance_api_v1_benchmark_benchmark_id_policies_policy_id":    "get-policy",
	"get_compliance_api_v1_benchmarks":                                   "list-benchmarks",
	"get_compliance_api_v1_benchmarks_summary":                           "benchmarks-summary",
	"get_compliance_api_v1_findings_benchmark_id_field_top_count":        "benchmark-findings-top-count",
	"get_compliance_api_v1_findings_metrics":                             "findings-metrics",
	"get_compliance_api_v1_queries_query_id":                             "get-query",
	"get_schedule_api_v1_benchmark_evaluations":                          "benchmark-evaluations",
	"post_compliance_api_v1_alarms_top":                                  "top-alarms",
	"post_compliance_api_v1_findings":                                    "get-findings",
	"get_compliance_api_v1_benchmarks_policies_policy_id":                "get-policy",
	"get_compliance_api_v1_queries_sync":                                 "sync-queries",
	"get_compliance_api_v1_benchmarks_benchmark_id_summary":              "get-benchmark-summary",
	"get_compliance_api_v1_benchmarks_benchmark_id_summary_result_trend": "get-benchmark-summary-trend",
	"get_compliance_api_v1_benchmarks_benchmark_id_tree":                 "get-benchmark-tree",
	"get_compliance_api_v1_benchmarks_benchmark_id_trend":                "get-benchmark-trend",

	// gpt
	"post_ai_api_v1_gpt_run": "_",

	// connection
	"get_inventory_api_v2_connections_data_connection_id": "get-connection",
	"get_inventory_api_v2_connections_data":               "list-connections",

	// connections
	"get_onboard_api_v1_connections_summary_connection_id": "connection-summary",
	"get_onboard_api_v1_connections_summary":               "connections-summary",

	// cost
	"get_inventory_api_v1_cost_top_accounts": "top-accounts-by-cost",
	"get_inventory_api_v1_cost_top_services": "top-services-by-cost",

	// describe
	"get_schedule_api_v0_compliance_summarizer_trigger":  "-",
	"get_schedule_api_v0_compliance_trigger":             "-",
	"get_schedule_api_v0_describe_trigger":               "-",
	"get_schedule_api_v0_insight_trigger":                "-",
	"get_schedule_api_v0_summarize_trigger":              "-",
	"get_schedule_api_v1_insight_job_job_id":             "get-insight-job",
	"post_schedule_api_v1_describe_resource":             "describe-resource",
	"post_schedule_api_v1_stacks_insight_trigger":        "run-stack-insight",
	"put_schedule_api_v1_benchmark_evaluation_trigger":   "run-benchmark-evaluation",
	"put_schedule_api_v1_compliance_trigger":             "run-compliance",
	"put_schedule_api_v1_describe_trigger_connection_id": "run-describe",
	"put_schedule_api_v1_insight_evaluation_trigger":     "run-insight",

	// insight
	"get_compliance_api_v2_insights_insight_id":                  "get-insight",
	"get_compliance_api_v2_insights_insight_id_trend":            "get-insight-trend",
	"get_compliance_api_v2_insights_job_job_id":                  "get-insight-job",
	"get_compliance_api_v1_insight_group_insight_group_id":       "get-insight-group",
	"get_compliance_api_v1_insight_group_insight_group_id_trend": "get-insight-group-trend",
	"get_compliance_api_v1_insight_group":                        "list-insight-groups",
	"get_compliance_api_v1_metadata_insight_insight_id":          "get-insight-metadata",
	"get_inventory_api_v2_insights_insight_id_trend":             "old-get-insight-trend",
	"get_inventory_api_v2_insights_job_job_id":                   "old-get-insight-job",
	"get_inventory_api_v2_insights_insight_id":                   "old-get-insight",

	// insights
	"get_compliance_api_v1_insight_insight_id":       "get-insight",
	"get_compliance_api_v1_insight_insight_id_trend": "get-insight-trend",
	"get_compliance_api_v1_insight":                  "list-insights",
	"get_compliance_api_v1_metadata_insight":         "list-insights-metadata",
	"get_compliance_api_v1_metadata_tag_insight":     "list-insights-metadata-tags",
	"get_compliance_api_v1_metadata_tag_insight_key": "list-insights-metadata-tag",

	// inventory
	"get_inventory_api_v1_resources_count":                "all-resources-count",
	"get_inventory_api_v1_resources_regions":              "regions-by-resources",
	"get_inventory_api_v1_resources_top_regions":          "top-regions-by-resources",
	"get_inventory_api_v2_resources_composition_key":      "resources-composition",
	"get_inventory_api_v2_resources_metric":               "resources-metrics",
	"get_inventory_api_v2_resources_metric_resource_type": "resource-metrics",
	"get_inventory_api_v2_resources_tag_key":              "get-resource-tag",
	"get_inventory_api_v2_resources_tag":                  "list-resource-tags",
	"get_inventory_api_v2_resources_trend":                "resource-trend",
	"get_inventory_api_v2_services_composition_key":       "services-composition",
	"get_inventory_api_v2_services_cost_trend":            "service-cost-trend",
	"get_inventory_api_v2_services_metric":                "service-metrics",
	"get_inventory_api_v2_services_metric_service_name":   "service-metric",
	"get_inventory_api_v2_services_tag_key":               "get-service-tag",
	"get_inventory_api_v2_services_tag":                   "list-service-tags",
	"post_inventory_api_v1_resources_aws":                 "list-aws-resources",
	"post_inventory_api_v1_resources_azure":               "list-azure-resources",
	"post_inventory_api_v1_resources":                     "list-all-resources",
	"post_inventory_api_v1_resources_filters":             "list-resources-by-filters",
	"get_inventory_api_v2_cost_trend":                     "cost-trend",
	"get_inventory_api_v2_cost_metric":                    "cost-metric",
	"get_inventory_api_v2_cost_composition":               "cost-composition",
	"get_inventory_api_v2_resources_regions_trend":        "get-resource-trend-by-region",
	"get_inventory_api_v2_resources_regions_composition":  "get-resource-composition-by-region",
	"get_inventory_api_v2_resources_regions_summary":      "get-resource-summary-by-region",
	"get_inventory_api_v2_resources_count":                "get-resource-count",
	"get_inventory_api_v2_analytics_metrics_list":         "list-metrics",
	"get_inventory_api_v2_analytics_spend_table":          "list-spend-table",
	"get_inventory_api_v2_analytics_table":                "list-costs-table",

	// keys
	"delete_auth_api_v1_key_id_delete": "remove-key",
	"get_auth_api_v1_key_id":           "get-key",
	"get_auth_api_v1_keys":             "list-keys",
	"post_auth_api_v1_key_create":      "create-key",
	"post_auth_api_v1_key_id_activate": "activate-key",
	"post_auth_api_v1_key_id_suspend":  "suspend-key",
	"post_auth_api_v1_key_role":        "update-key-role",

	// locations
	"get_inventory_api_v1_locations_connector": "list-connector-locations",

	// metadata
	"get_inventory_api_v2_metadata_resourcetype":               "list-resource-types",
	"get_inventory_api_v2_metadata_resourcetype_resource_type": "get-resource-type",
	"get_inventory_api_v2_metadata_services":                   "list-services",
	"get_inventory_api_v2_metadata_services_service_name":      "get-service",
	"get_metadata_api_v1_metadata_key":                         "get-config-metadata",
	"post_metadata_api_v1_metadata":                            "set-config-metadata",
	"get_metadata_api_v1_filter":                               "list-filters",
	"post_metadata_api_v1_filter":                              "add-filter",

	// onboard
	"delete_onboard_api_v1_credential_credential_id":             "delete-credential",
	"delete_onboard_api_v1_source_source_id":                     "delete-source",
	"get_onboard_api_v1_catalog_connectors":                      "get-catalog-connectors",
	"get_onboard_api_v1_catalog_metrics":                         "get-catalog-metrics",
	"get_onboard_api_v1_connections_count":                       "get-connections-count",
	"get_onboard_api_v1_connector_connector_name":                "get-connector",
	"get_onboard_api_v1_connectors_connector_name":               "get-connector",
	"get_onboard_api_v1_connectors":                              "get-connectors",
	"get_onboard_api_v1_connector":                               "get-connectors",
	"get_onboard_api_v1_credential_credential_id_healthcheck":    "healthcheck-credential",
	"get_onboard_api_v1_credential_credential_id":                "get-credential",
	"get_onboard_api_v1_credential":                              "list-credentials",
	"get_onboard_api_v1_credential_sources_list":                 "list-credentials-sources",
	"get_onboard_api_v1_providers":                               "get-providers",
	"get_onboard_api_v1_providers_types":                         "get-providers-types",
	"get_onboard_api_v1_source_account_account_id":               "get-account-source",
	"get_onboard_api_v1_source_source_id_credentials":            "list-source-credentials",
	"get_onboard_api_v1_source_source_id":                        "get-source",
	"get_onboard_api_v1_sources":                                 "list-sources",
	"get_onboard_api_v1_sources_count":                           "get-sources-count",
	"post_onboard_api_v1_connections_connection_id_state":        "update-connection-state",
	"post_onboard_api_v1_credential_credential_id_autoonboard":   "auto-onboard-credential",
	"post_onboard_api_v1_credential_credential_id_disable":       "disable-credential",
	"post_onboard_api_v1_credential_credential_id_enable":        "enable-credential",
	"post_onboard_api_v1_credential":                             "create-credential",
	"post_onboard_api_v1_source_aws":                             "create-aws-source",
	"post_onboard_api_v1_source_azure":                           "create-azure-source",
	"post_onboard_api_v1_source_source_id_disable":               "disable-source",
	"post_onboard_api_v1_source_source_id_enable":                "enable-source",
	"post_onboard_api_v1_source_source_id_healthcheck":           "healthchek-source",
	"post_onboard_api_v1_sources":                                "get-source",
	"put_onboard_api_v1_credential_credential_id":                "update-credential",
	"put_onboard_api_v1_source_source_id_credentials":            "put-source-credentials",
	"get_onboard_api_v1_connection_groups_connection_group_name": "get-connection-group",
	"get_onboard_api_v1_source_source_id_healthcheck":            "healthcheck-source",
	"get_onboard_api_v1_connection_groups":                       "list-connection-groups",

	// resource
	"post_inventory_api_v1_resource": "create-resource",

	// roles
	"get_auth_api_v1_role_role_name_keys":  "list-role-keys",
	"get_auth_api_v1_role_role_name_users": "list-role-users",
	"get_auth_api_v1_roles_role_name":      "get-role",
	"get_auth_api_v1_roles":                "list-roles",

	// schedule
	"get_schedule_api_v1_compliance_report_last_completed":           "list-last-compliance-jobs",
	"get_schedule_api_v1_describe_resource_jobs_pending":             "list-pending-describe-resource-jobs",
	"get_schedule_api_v1_describe_source_jobs_pending":               "list-pending-describe-source-jobs",
	"get_schedule_api_v1_insight_jobs_pending":                       "list-pending-insight-jobs",
	"get_schedule_api_v1_resource_type_provider":                     "list-provider-resource-types",
	"get_schedule_api_v1_sources":                                    "list-sources",
	"get_schedule_api_v1_sources_source_id_jobs_compliance":          "source-compliance-jobs",
	"get_schedule_api_v1_sources_source_id_jobs_describe":            "source-describe-jobs",
	"get_schedule_api_v1_sources_source_id":                          "get-source",
	"get_schedule_api_v1_summarize_jobs_pending":                     "list-pending-summarize-jobs",
	"post_schedule_api_v1_sources_source_id_jobs_compliance_refresh": "refresh-source-compliance-jobs",
	"post_schedule_api_v1_sources_source_id_jobs_describe_refresh":   "refresh-source-describe-jobs",

	// smart_query
	"get_inventory_api_v1_query_count":       "get-queries-count",
	"get_inventory_api_v1_query":             "list-queries",
	"post_inventory_api_v1_query_query_id":   "get-query",
	"get_inventory_api_v1_query_run_history": "get-query-run-history",
	"post_inventory_api_v1_query_run":        "run-query",

	// users
	"delete_auth_api_v1_user_invite":                    "delete-user-invite",
	"delete_auth_api_v1_user_role_binding":              "delete-user-role-binding",
	"get_auth_api_v1_user_role_bindings":                "list-user-role-bindings",
	"get_auth_api_v1_user_user_id":                      "get-user",
	"get_auth_api_v1_users":                             "list-users",
	"get_auth_api_v1_user_user_id_workspace_membership": "list-user-workspace-memberships",
	"post_auth_api_v1_user_invite":                      "invite-user",
	"get_auth_api_v1_workspace_role_bindings":           "list-workspace-role-bindings",
	"put_auth_api_v1_user_role_binding":                 "update-user-role-binding",

	// workspace
	"delete_workspace_api_v1_workspace_workspace_id":            "delete-workspace",
	"get_workspace_api_v1_workspace_workspace_id":               "get-workspace",
	"get_workspace_api_v1_workspaces":                           "list-workspaces",
	"post_workspace_api_v1_workspace":                           "create-workspace",
	"get_workspace_api_v1_workspaces_byid_workspace_id":         "get-workspace-by-id",
	"get_workspace_api_v1_workspaces_limits_byid_workspace_id":  "get-workspace-limit-by-id",
	"get_workspace_api_v1_workspaces_limits_workspace_name":     "list-workspace-limits",
	"post_workspace_api_v1_workspace_workspace_id_name":         "update-workspace-name",
	"post_workspace_api_v1_workspace_workspace_id_organization": "update-workspace-organization",
	"post_workspace_api_v1_workspace_workspace_id_owner":        "update-workspace-owner",
	"post_workspace_api_v1_workspace_workspace_id_resume":       "resume-workspace",
	"post_workspace_api_v1_workspace_workspace_id_suspend":      "suspend-workspace",
	"post_workspace_api_v1_workspace_workspace_id_tier":         "update-workspace-tier",
	"get_workspace_api_v1_workspaces_workspace_id":              "get-workspace",
	"get_workspace_api_v1_workspace_current":                    "get-current-workspace",
	"delete_workspace_api_v1_organization_organization_id":      "delete-organization",
	"post_workspace_api_v1_organization":                        "create-organization",

	// analytics
	"get_inventory_api_v2_analytics_trend":               "get-assets-trend",
	"get_inventory_api_v2_analytics_categories":          "list-categories",
	"get_inventory_api_v2_analytics_metric":              "get-assets-by-metric",
	"get_inventory_api_v2_analytics_spend_composition":   "get-spend-composition",
	"get_inventory_api_v2_analytics_regions_summary":     "get-assets-by-regions",
	"get_inventory_api_v2_analytics_spend_metrics":       "list-analytics-spend-metrics",
	"get_inventory_api_v2_analytics_spend_metric":        "get-spend-by-metric",
	"get_inventory_api_v2_analytics_composition_keys":    "list-analytics-composition-keys",
	"get_inventory_api_v2_analytics_composition_key":     "get-assets-composition",
	"get_inventory_api_v2_analytics_spend_trend":         "get-costs-trend",
	"get_inventory_api_v2_analytics_tag":                 "get-tags",
	"get_inventory_api_v2_analytics_spend_metrics_trend": "get-spend-trend-by-metrics",
}

var PreferredService = map[string][]string{
	"get_inventory_api_v2_analytics_spend_trend":         {"Spend"},
	"get_inventory_api_v2_analytics_composition_key":     {"Assets"},
	"get_inventory_api_v2_analytics_metric":              {"Assets"},
	"get_inventory_api_v2_analytics_regions_summary":     {"Assets"},
	"get_inventory_api_v2_analytics_spend_composition":   {"Spend"},
	"get_inventory_api_v2_analytics_spend_metric":        {"Spend"},
	"get_inventory_api_v2_analytics_spend_metrics_trend": {"Spend"},
	"get_inventory_api_v2_analytics_tag":                 {"Spend", "Assets"},
	"get_inventory_api_v2_analytics_trend":               {"Assets"},
	"get_inventory_api_v2_analytics_categories":          {"Assets", "Spend"},
	"get_inventory_api_v2_analytics_metrics_list":        {"Assets", "Spend"},
	"get_inventory_api_v2_analytics_table":               {"Assets"},
	"get_inventory_api_v2_analytics_spend_table":         {"Spend"},
}
