package main

import (
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks_assignment"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/compliance"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connection"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/cost"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/describe"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insight"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/inventory"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/keys"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/location"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/metadata"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/onboard"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/roles"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/schedule"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/smart_query"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/stack"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/users"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
)

var paramModels = map[string]interface{}{
	"GetInventoryAPIV1AccountsResourceCount":                            benchmarks.GetInventoryAPIV1AccountsResourceCountParams{},
	"GetInventoryAPIV1ResourcesDistribution":                            benchmarks.GetInventoryAPIV1ResourcesDistributionParams{},
	"GetInventoryAPIV1ResourcesTopGrowingAccounts":                      benchmarks.GetInventoryAPIV1ResourcesTopGrowingAccountsParams{},
	"GetInventoryAPIV1ServicesDistribution":                             benchmarks.GetInventoryAPIV1ServicesDistributionParams{},
	"GetInventoryAPIV2ServicesSummary":                                  benchmarks.GetInventoryAPIV2ServicesSummaryParams{},
	"GetInventoryAPIV2ServicesSummaryServiceName":                       benchmarks.GetInventoryAPIV2ServicesSummaryServiceNameParams{},
	"DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID": benchmarks_assignment.DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams{},
	"GetComplianceAPIV1AssignmentsBenchmarkBenchmarkID":                 benchmarks_assignment.GetComplianceAPIV1AssignmentsBenchmarkBenchmarkIDParams{},
	"GetComplianceAPIV1AssignmentsConnectionConnectionID":               benchmarks_assignment.GetComplianceAPIV1AssignmentsConnectionConnectionIDParams{},
	"GetComplianceAPIV1Assignments":                                     benchmarks_assignment.GetComplianceAPIV1AssignmentsParams{},
	"PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID":   benchmarks_assignment.PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams{},
	"GetComplianceAPIV1BenchmarkBenchmarkIDSummary":                     compliance.GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams{},
	"GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrend":          compliance.GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams{},
	"GetComplianceAPIV1BenchmarkBenchmarkIDTree":                        compliance.GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams{},
	"GetComplianceAPIV1BenchmarksBenchmarkID":                           compliance.GetComplianceAPIV1BenchmarksBenchmarkIDParams{},
	"GetComplianceAPIV1BenchmarksBenchmarkIDPolicies":                   compliance.GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesParams{},
	"GetComplianceAPIV1Benchmarks":                                      compliance.GetComplianceAPIV1BenchmarksParams{},
	"GetComplianceAPIV1BenchmarksPoliciesPolicyID":                      compliance.GetComplianceAPIV1BenchmarksPoliciesPolicyIDParams{},
	"GetComplianceAPIV1BenchmarksSummary":                               compliance.GetComplianceAPIV1BenchmarksSummaryParams{},
	"GetComplianceAPIV1FindingsBenchmarkIDFieldTopCount":                compliance.GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams{},
	"GetComplianceAPIV1FindingsMetrics":                                 compliance.GetComplianceAPIV1FindingsMetricsParams{},
	"GetComplianceAPIV1QueriesQueryID":                                  compliance.GetComplianceAPIV1QueriesQueryIDParams{},
	"GetScheduleAPIV1BenchmarkEvaluations":                              compliance.GetScheduleAPIV1BenchmarkEvaluationsParams{},
	"PostComplianceAPIV1AlarmsTop":                                      compliance.PostComplianceAPIV1AlarmsTopParams{},
	"PostComplianceAPIV1Findings":                                       compliance.PostComplianceAPIV1FindingsParams{},
	"GetInventoryAPIV2ConnectionsDataConnectionID":                      connection.GetInventoryAPIV2ConnectionsDataConnectionIDParams{},
	"GetInventoryAPIV2ConnectionsData":                                  connection.GetInventoryAPIV2ConnectionsDataParams{},
	"GetOnboardAPIV1ConnectionsSummaryConnectionID":                     connections.GetOnboardAPIV1ConnectionsSummaryConnectionIDParams{},
	"GetOnboardAPIV1ConnectionsSummary":                                 connections.GetOnboardAPIV1ConnectionsSummaryParams{},
	"GetInventoryAPIV1CostTopAccounts":                                  cost.GetInventoryAPIV1CostTopAccountsParams{},
	"GetInventoryAPIV1CostTopServices":                                  cost.GetInventoryAPIV1CostTopServicesParams{},
	"GetScheduleAPIV0ComplianceSummarizerTrigger":                       describe.GetScheduleAPIV0ComplianceSummarizerTriggerParams{},
	"GetScheduleAPIV0ComplianceTrigger":                                 describe.GetScheduleAPIV0ComplianceTriggerParams{},
	"GetScheduleAPIV0DescribeTrigger":                                   describe.GetScheduleAPIV0DescribeTriggerParams{},
	"GetScheduleAPIV0InsightTrigger":                                    describe.GetScheduleAPIV0InsightTriggerParams{},
	"GetScheduleAPIV0SummarizeTrigger":                                  describe.GetScheduleAPIV0SummarizeTriggerParams{},
	"GetScheduleAPIV1InsightJobJobID":                                   describe.GetScheduleAPIV1InsightJobJobIDParams{},
	"PostScheduleAPIV1DescribeResource":                                 describe.PostScheduleAPIV1DescribeResourceParams{},
	"PostScheduleAPIV1StacksInsightTrigger":                             describe.PostScheduleAPIV1StacksInsightTriggerParams{},
	"PutScheduleAPIV1BenchmarkEvaluationTrigger":                        describe.PutScheduleAPIV1BenchmarkEvaluationTriggerParams{},
	"PutScheduleAPIV1DescribeTriggerConnectionID":                       describe.PutScheduleAPIV1DescribeTriggerConnectionIDParams{},
	"PutScheduleAPIV1InsightEvaluationTrigger":                          describe.PutScheduleAPIV1InsightEvaluationTriggerParams{},
	"GetInventoryAPIV2InsightsInsightID":                                insight.GetInventoryAPIV2InsightsInsightIDParams{},
	"GetInventoryAPIV2InsightsInsightIDTrend":                           insight.GetInventoryAPIV2InsightsInsightIDTrendParams{},
	"GetInventoryAPIV2InsightsJobJobID":                                 insight.GetInventoryAPIV2InsightsJobJobIDParams{},
	"GetInventoryAPIV2Insights":                                         insight.GetInventoryAPIV2InsightsParams{},
	"GetComplianceAPIV1InsightInsightID":                                insights.GetComplianceAPIV1InsightInsightIDParams{},
	"GetComplianceAPIV1InsightInsightIDTrend":                           insights.GetComplianceAPIV1InsightInsightIDTrendParams{},
	"GetComplianceAPIV1Insight":                                         insights.GetComplianceAPIV1InsightParams{},
	"GetComplianceAPIV1MetadataInsightInsightID":                        insights.GetComplianceAPIV1MetadataInsightInsightIDParams{},
	"GetComplianceAPIV1MetadataInsight":                                 insights.GetComplianceAPIV1MetadataInsightParams{},
	"GetComplianceAPIV1MetadataTagInsightKey":                           insights.GetComplianceAPIV1MetadataTagInsightKeyParams{},
	"GetComplianceAPIV1MetadataTagInsight":                              insights.GetComplianceAPIV1MetadataTagInsightParams{},
	"GetInventoryAPIV1ResourcesCount":                                   inventory.GetInventoryAPIV1ResourcesCountParams{},
	"GetInventoryAPIV1ResourcesRegions":                                 inventory.GetInventoryAPIV1ResourcesRegionsParams{},
	"GetInventoryAPIV1ResourcesTopRegions":                              inventory.GetInventoryAPIV1ResourcesTopRegionsParams{},
	"GetInventoryAPIV2CostComposition":                                  inventory.GetInventoryAPIV2CostCompositionParams{},
	"GetInventoryAPIV2CostMetric":                                       inventory.GetInventoryAPIV2CostMetricParams{},
	"GetInventoryAPIV2CostTrend":                                        inventory.GetInventoryAPIV2CostTrendParams{},
	"GetInventoryAPIV2ResourcesCompositionKey":                          inventory.GetInventoryAPIV2ResourcesCompositionKeyParams{},
	"GetInventoryAPIV2ResourcesMetric":                                  inventory.GetInventoryAPIV2ResourcesMetricParams{},
	"GetInventoryAPIV2ResourcesMetricResourceType":                      inventory.GetInventoryAPIV2ResourcesMetricResourceTypeParams{},
	"GetInventoryAPIV2ResourcesTagKey":                                  inventory.GetInventoryAPIV2ResourcesTagKeyParams{},
	"GetInventoryAPIV2ResourcesTag":                                     inventory.GetInventoryAPIV2ResourcesTagParams{},
	"GetInventoryAPIV2ResourcesTrend":                                   inventory.GetInventoryAPIV2ResourcesTrendParams{},
	"GetInventoryAPIV2ServicesMetric":                                   inventory.GetInventoryAPIV2ServicesMetricParams{},
	"GetInventoryAPIV2ServicesMetricServiceName":                        inventory.GetInventoryAPIV2ServicesMetricServiceNameParams{},
	"GetInventoryAPIV2ServicesTagKey":                                   inventory.GetInventoryAPIV2ServicesTagKeyParams{},
	"GetInventoryAPIV2ServicesTag":                                      inventory.GetInventoryAPIV2ServicesTagParams{},
	"PostInventoryAPIV1ResourcesAws":                                    inventory.PostInventoryAPIV1ResourcesAwsParams{},
	"PostInventoryAPIV1ResourcesAzure":                                  inventory.PostInventoryAPIV1ResourcesAzureParams{},
	"PostInventoryAPIV1ResourcesFilters":                                inventory.PostInventoryAPIV1ResourcesFiltersParams{},
	"PostInventoryAPIV1Resources":                                       inventory.PostInventoryAPIV1ResourcesParams{},
	"DeleteAuthAPIV1KeyIDDelete":                                        keys.DeleteAuthAPIV1KeyIDDeleteParams{},
	"GetAuthAPIV1KeyID":                                                 keys.GetAuthAPIV1KeyIDParams{},
	"GetAuthAPIV1Keys":                                                  keys.GetAuthAPIV1KeysParams{},
	"PostAuthAPIV1KeyCreate":                                            keys.PostAuthAPIV1KeyCreateParams{},
	"PostAuthAPIV1KeyIDActivate":                                        keys.PostAuthAPIV1KeyIDActivateParams{},
	"PostAuthAPIV1KeyIDSuspend":                                         keys.PostAuthAPIV1KeyIDSuspendParams{},
	"PostAuthAPIV1KeyRole":                                              keys.PostAuthAPIV1KeyRoleParams{},
	"GetInventoryAPIV1LocationsConnector":                               location.GetInventoryAPIV1LocationsConnectorParams{},
	"GetInventoryAPIV2MetadataResourcetype":                             metadata.GetInventoryAPIV2MetadataResourcetypeParams{},
	"GetInventoryAPIV2MetadataResourcetypeResourceType":                 metadata.GetInventoryAPIV2MetadataResourcetypeResourceTypeParams{},
	"GetInventoryAPIV2MetadataServices":                                 metadata.GetInventoryAPIV2MetadataServicesParams{},
	"GetInventoryAPIV2MetadataServicesServiceName":                      metadata.GetInventoryAPIV2MetadataServicesServiceNameParams{},
	"GetMetadataAPIV1MetadataKey":                                       metadata.GetMetadataAPIV1MetadataKeyParams{},
	"PostMetadataAPIV1Metadata":                                         metadata.PostMetadataAPIV1MetadataParams{},
	"DeleteOnboardAPIV1CredentialCredentialID":                          onboard.DeleteOnboardAPIV1CredentialCredentialIDParams{},
	"DeleteOnboardAPIV1SourceSourceID":                                  onboard.DeleteOnboardAPIV1SourceSourceIDParams{},
	"GetOnboardAPIV1CatalogConnectors":                                  onboard.GetOnboardAPIV1CatalogConnectorsParams{},
	"GetOnboardAPIV1CatalogMetrics":                                     onboard.GetOnboardAPIV1CatalogMetricsParams{},
	"GetOnboardAPIV1ConnectionsCount":                                   onboard.GetOnboardAPIV1ConnectionsCountParams{},
	"GetOnboardAPIV1ConnectorConnectorName":                             onboard.GetOnboardAPIV1ConnectorConnectorNameParams{},
	"GetOnboardAPIV1Connector":                                          onboard.GetOnboardAPIV1ConnectorParams{},
	"GetOnboardAPIV1CredentialCredentialIDHealthcheck":                  onboard.GetOnboardAPIV1CredentialCredentialIDHealthcheckParams{},
	"GetOnboardAPIV1CredentialCredentialID":                             onboard.GetOnboardAPIV1CredentialCredentialIDParams{},
	"GetOnboardAPIV1Credential":                                         onboard.GetOnboardAPIV1CredentialParams{},
	"GetOnboardAPIV1CredentialSourcesList":                              onboard.GetOnboardAPIV1CredentialSourcesListParams{},
	"GetOnboardAPIV1SourceAccountAccountID":                             onboard.GetOnboardAPIV1SourceAccountAccountIDParams{},
	"GetOnboardAPIV1SourceSourceIDCredentials":                          onboard.GetOnboardAPIV1SourceSourceIDCredentialsParams{},
	"GetOnboardAPIV1SourceSourceID":                                     onboard.GetOnboardAPIV1SourceSourceIDParams{},
	"GetOnboardAPIV1SourcesCount":                                       onboard.GetOnboardAPIV1SourcesCountParams{},
	"GetOnboardAPIV1Sources":                                            onboard.GetOnboardAPIV1SourcesParams{},
	"PostOnboardAPIV1ConnectionsConnectionIDState":                      onboard.PostOnboardAPIV1ConnectionsConnectionIDStateParams{},
	"PostOnboardAPIV1CredentialCredentialIDAutoonboard":                 onboard.PostOnboardAPIV1CredentialCredentialIDAutoonboardParams{},
	"PostOnboardAPIV1CredentialCredentialIDDisable":                     onboard.PostOnboardAPIV1CredentialCredentialIDDisableParams{},
	"PostOnboardAPIV1CredentialCredentialIDEnable":                      onboard.PostOnboardAPIV1CredentialCredentialIDEnableParams{},
	"PostOnboardAPIV1Credential":                                        onboard.PostOnboardAPIV1CredentialParams{},
	"PostOnboardAPIV1SourceAws":                                         onboard.PostOnboardAPIV1SourceAwsParams{},
	"PostOnboardAPIV1SourceAzure":                                       onboard.PostOnboardAPIV1SourceAzureParams{},
	"PostOnboardAPIV1SourceSourceIDHealthcheck":                         onboard.PostOnboardAPIV1SourceSourceIDHealthcheckParams{},
	"PostOnboardAPIV1Sources":                                           onboard.PostOnboardAPIV1SourcesParams{},
	"PutOnboardAPIV1CredentialCredentialID":                             onboard.PutOnboardAPIV1CredentialCredentialIDParams{},
	"PutOnboardAPIV1SourceSourceIDCredentials":                          onboard.PutOnboardAPIV1SourceSourceIDCredentialsParams{},
	"PostInventoryAPIV1Resource":                                        resource.PostInventoryAPIV1ResourceParams{},
	"GetAuthAPIV1RoleRoleNameKeys":                                      roles.GetAuthAPIV1RoleRoleNameKeysParams{},
	"GetAuthAPIV1RoleRoleNameUsers":                                     roles.GetAuthAPIV1RoleRoleNameUsersParams{},
	"GetAuthAPIV1Roles":                                                 roles.GetAuthAPIV1RolesParams{},
	"GetAuthAPIV1RolesRoleName":                                         roles.GetAuthAPIV1RolesRoleNameParams{},
	"GetScheduleAPIV1ComplianceReportLastCompleted":                     schedule.GetScheduleAPIV1ComplianceReportLastCompletedParams{},
	"GetScheduleAPIV1DescribeResourceJobsPending":                       schedule.GetScheduleAPIV1DescribeResourceJobsPendingParams{},
	"GetScheduleAPIV1DescribeSourceJobsPending":                         schedule.GetScheduleAPIV1DescribeSourceJobsPendingParams{},
	"GetScheduleAPIV1InsightJobsPending":                                schedule.GetScheduleAPIV1InsightJobsPendingParams{},
	"GetScheduleAPIV1ResourceTypeProvider":                              schedule.GetScheduleAPIV1ResourceTypeProviderParams{},
	"GetScheduleAPIV1Sources":                                           schedule.GetScheduleAPIV1SourcesParams{},
	"GetScheduleAPIV1SourcesSourceIDJobsCompliance":                     schedule.GetScheduleAPIV1SourcesSourceIDJobsComplianceParams{},
	"GetScheduleAPIV1SourcesSourceIDJobsDescribe":                       schedule.GetScheduleAPIV1SourcesSourceIDJobsDescribeParams{},
	"GetScheduleAPIV1SourcesSourceID":                                   schedule.GetScheduleAPIV1SourcesSourceIDParams{},
	"GetScheduleAPIV1SummarizeJobsPending":                              schedule.GetScheduleAPIV1SummarizeJobsPendingParams{},
	"PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh":             schedule.PostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshParams{},
	"PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh":               schedule.PostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshParams{},
	"GetInventoryAPIV1QueryCount":                                       smart_query.GetInventoryAPIV1QueryCountParams{},
	"GetInventoryAPIV1Query":                                            smart_query.GetInventoryAPIV1QueryParams{},
	"PostInventoryAPIV1QueryQueryID":                                    smart_query.PostInventoryAPIV1QueryQueryIDParams{},
	"DeleteScheduleAPIV1StacksStackID":                                  stack.DeleteScheduleAPIV1StacksStackIDParams{},
	"GetScheduleAPIV1Stacks":                                            stack.GetScheduleAPIV1StacksParams{},
	"GetScheduleAPIV1StacksResource":                                    stack.GetScheduleAPIV1StacksResourceParams{},
	"GetScheduleAPIV1StacksStackIDInsight":                              stack.GetScheduleAPIV1StacksStackIDInsightParams{},
	"GetScheduleAPIV1StacksStackID":                                     stack.GetScheduleAPIV1StacksStackIDParams{},
	"PostScheduleAPIV1StacksBenchmarkTrigger":                           stack.PostScheduleAPIV1StacksBenchmarkTriggerParams{},
	"PostScheduleAPIV1StacksCreate":                                     stack.PostScheduleAPIV1StacksCreateParams{},
	"PostScheduleAPIV1StacksStackIDFindings":                            stack.PostScheduleAPIV1StacksStackIDFindingsParams{},
	"DeleteAuthAPIV1UserInvite":                                         users.DeleteAuthAPIV1UserInviteParams{},
	"DeleteAuthAPIV1UserRoleBinding":                                    users.DeleteAuthAPIV1UserRoleBindingParams{},
	"GetAuthAPIV1UserRoleBindings":                                      users.GetAuthAPIV1UserRoleBindingsParams{},
	"GetAuthAPIV1UserUserID":                                            users.GetAuthAPIV1UserUserIDParams{},
	"GetAuthAPIV1UserUserIDWorkspaceMembership":                         users.GetAuthAPIV1UserUserIDWorkspaceMembershipParams{},
	"GetAuthAPIV1Users":                                                 users.GetAuthAPIV1UsersParams{},
	"GetAuthAPIV1WorkspaceRoleBindings":                                 users.GetAuthAPIV1WorkspaceRoleBindingsParams{},
	"PostAuthAPIV1UserInvite":                                           users.PostAuthAPIV1UserInviteParams{},
	"PutAuthAPIV1UserRoleBinding":                                       users.PutAuthAPIV1UserRoleBindingParams{},
	"DeleteWorkspaceAPIV1WorkspaceWorkspaceID":                          workspace.DeleteWorkspaceAPIV1WorkspaceWorkspaceIDParams{},
	"GetWorkspaceAPIV1WorkspacesByidWorkspaceID":                        workspace.GetWorkspaceAPIV1WorkspacesByidWorkspaceIDParams{},
	"GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceID":                  workspace.GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams{},
	"GetWorkspaceAPIV1WorkspacesLimitsWorkspaceName":                    workspace.GetWorkspaceAPIV1WorkspacesLimitsWorkspaceNameParams{},
	"GetWorkspaceAPIV1Workspaces":                                       workspace.GetWorkspaceAPIV1WorkspacesParams{},
	"GetWorkspaceAPIV1WorkspacesWorkspaceID":                            workspace.GetWorkspaceAPIV1WorkspacesWorkspaceIDParams{},
	"PostWorkspaceAPIV1Workspace":                                       workspace.PostWorkspaceAPIV1WorkspaceParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDName":                        workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDNameParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDOrganization":                workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOrganizationParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDOwner":                       workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDResume":                      workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspend":                     workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDTier":                        workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDTierParams{},
}
