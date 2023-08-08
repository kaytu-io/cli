package config

import (
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/analytics"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks_assignment"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/compliance"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connection_groups"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/describe"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/inventory"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/keys"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/metadata"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/onboard"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/smart_query"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/stack"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/users"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
)

var ParamModels = map[string]interface{}{
	"GetInventoryAPIV2AnalyticsCategories":                              analytics.GetInventoryAPIV2AnalyticsCategoriesParams{},
	"GetInventoryAPIV2AnalyticsCompositionKey":                          analytics.GetInventoryAPIV2AnalyticsCompositionKeyParams{},
	"GetInventoryAPIV2AnalyticsMetric":                                  analytics.GetInventoryAPIV2AnalyticsMetricParams{},
	"GetInventoryAPIV2AnalyticsRegionsSummary":                          analytics.GetInventoryAPIV2AnalyticsRegionsSummaryParams{},
	"GetInventoryAPIV2AnalyticsTag":                                     analytics.GetInventoryAPIV2AnalyticsTagParams{},
	"GetInventoryAPIV2AnalyticsTrend":                                   analytics.GetInventoryAPIV2AnalyticsTrendParams{},
	"DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID": benchmarks_assignment.DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams{},
	"GetComplianceAPIV1AssignmentsBenchmarkBenchmarkID":                 benchmarks_assignment.GetComplianceAPIV1AssignmentsBenchmarkBenchmarkIDParams{},
	"PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID":   benchmarks_assignment.PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams{},
	"GetComplianceAPIV1BenchmarksBenchmarkIDSummary":                    compliance.GetComplianceAPIV1BenchmarksBenchmarkIDSummaryParams{},
	"GetComplianceAPIV1BenchmarksBenchmarkIDTree":                       compliance.GetComplianceAPIV1BenchmarksBenchmarkIDTreeParams{},
	"GetComplianceAPIV1BenchmarksBenchmarkIDTrend":                      compliance.GetComplianceAPIV1BenchmarksBenchmarkIDTrendParams{},
	"GetComplianceAPIV1BenchmarksSummary":                               compliance.GetComplianceAPIV1BenchmarksSummaryParams{},
	"GetComplianceAPIV1FindingsBenchmarkIDFieldTopCount":                compliance.GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams{},
	"GetComplianceAPIV1QueriesSync":                                     compliance.GetComplianceAPIV1QueriesSyncParams{},
	"PostComplianceAPIV1Findings":                                       compliance.PostComplianceAPIV1FindingsParams{},
	"GetOnboardAPIV1ConnectionGroupsConnectionGroupName":                connection_groups.GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams{},
	"GetOnboardAPIV1ConnectionGroups":                                   connection_groups.GetOnboardAPIV1ConnectionGroupsParams{},
	"GetOnboardAPIV1ConnectionsSummary":                                 connections.GetOnboardAPIV1ConnectionsSummaryParams{},
	"PutScheduleAPIV1DescribeTriggerConnectionID":                       describe.PutScheduleAPIV1DescribeTriggerConnectionIDParams{},
	"GetComplianceAPIV1InsightGroup":                                    insights.GetComplianceAPIV1InsightGroupParams{},
	"GetComplianceAPIV1InsightInsightID":                                insights.GetComplianceAPIV1InsightInsightIDParams{},
	"GetComplianceAPIV1InsightInsightIDTrend":                           insights.GetComplianceAPIV1InsightInsightIDTrendParams{},
	"GetComplianceAPIV1Insight":                                         insights.GetComplianceAPIV1InsightParams{},
	"GetComplianceAPIV1MetadataInsightInsightID":                        insights.GetComplianceAPIV1MetadataInsightInsightIDParams{},
	"GetComplianceAPIV1MetadataTagInsight":                              insights.GetComplianceAPIV1MetadataTagInsightParams{},
	"GetInventoryAPIV2AnalyticsSpendComposition":                        inventory.GetInventoryAPIV2AnalyticsSpendCompositionParams{},
	"GetInventoryAPIV2AnalyticsSpendMetric":                             inventory.GetInventoryAPIV2AnalyticsSpendMetricParams{},
	"GetInventoryAPIV2AnalyticsSpendMetricsTrend":                       inventory.GetInventoryAPIV2AnalyticsSpendMetricsTrendParams{},
	"GetInventoryAPIV2AnalyticsSpendTrend":                              inventory.GetInventoryAPIV2AnalyticsSpendTrendParams{},
	"GetInventoryAPIV2CostComposition":                                  inventory.GetInventoryAPIV2CostCompositionParams{},
	"GetInventoryAPIV2CostMetric":                                       inventory.GetInventoryAPIV2CostMetricParams{},
	"GetInventoryAPIV2CostTrend":                                        inventory.GetInventoryAPIV2CostTrendParams{},
	"GetInventoryAPIV2ResourcesMetricResourceType":                      inventory.GetInventoryAPIV2ResourcesMetricResourceTypeParams{},
	"GetInventoryAPIV2ResourcesTag":                                     inventory.GetInventoryAPIV2ResourcesTagParams{},
	"GetInventoryAPIV2ServicesCostTrend":                                inventory.GetInventoryAPIV2ServicesCostTrendParams{},
	"DeleteAuthAPIV1KeyIDDelete":                                        keys.DeleteAuthAPIV1KeyIDDeleteParams{},
	"GetAuthAPIV1Keys":                                                  keys.GetAuthAPIV1KeysParams{},
	"PostAuthAPIV1KeyCreate":                                            keys.PostAuthAPIV1KeyCreateParams{},
	"GetMetadataAPIV1MetadataKey":                                       metadata.GetMetadataAPIV1MetadataKeyParams{},
	"PostMetadataAPIV1Metadata":                                         metadata.PostMetadataAPIV1MetadataParams{},
	"DeleteOnboardAPIV1CredentialCredentialID":                          onboard.DeleteOnboardAPIV1CredentialCredentialIDParams{},
	"DeleteOnboardAPIV1SourceSourceID":                                  onboard.DeleteOnboardAPIV1SourceSourceIDParams{},
	"GetOnboardAPIV1CatalogMetrics":                                     onboard.GetOnboardAPIV1CatalogMetricsParams{},
	"GetOnboardAPIV1Connector":                                          onboard.GetOnboardAPIV1ConnectorParams{},
	"GetOnboardAPIV1CredentialCredentialID":                             onboard.GetOnboardAPIV1CredentialCredentialIDParams{},
	"GetOnboardAPIV1Credential":                                         onboard.GetOnboardAPIV1CredentialParams{},
	"GetOnboardAPIV1SourceSourceIDHealthcheck":                          onboard.GetOnboardAPIV1SourceSourceIDHealthcheckParams{},
	"PostOnboardAPIV1CredentialCredentialIDAutoonboard":                 onboard.PostOnboardAPIV1CredentialCredentialIDAutoonboardParams{},
	"PostOnboardAPIV1Credential":                                        onboard.PostOnboardAPIV1CredentialParams{},
	"PostOnboardAPIV1SourceAws":                                         onboard.PostOnboardAPIV1SourceAwsParams{},
	"PostOnboardAPIV1SourceAzure":                                       onboard.PostOnboardAPIV1SourceAzureParams{},
	"PutOnboardAPIV1CredentialCredentialID":                             onboard.PutOnboardAPIV1CredentialCredentialIDParams{},
	"PostAiAPIV1GptRun":                                                 resource.PostAiAPIV1GptRunParams{},
	"GetInventoryAPIV1Query":                                            smart_query.GetInventoryAPIV1QueryParams{},
	"GetInventoryAPIV1QueryRunHistory":                                  smart_query.GetInventoryAPIV1QueryRunHistoryParams{},
	"PostInventoryAPIV1QueryRun":                                        smart_query.PostInventoryAPIV1QueryRunParams{},
	"DeleteScheduleAPIV1StacksStackID":                                  stack.DeleteScheduleAPIV1StacksStackIDParams{},
	"GetScheduleAPIV1Stacks":                                            stack.GetScheduleAPIV1StacksParams{},
	"GetScheduleAPIV1StacksResource":                                    stack.GetScheduleAPIV1StacksResourceParams{},
	"GetScheduleAPIV1StacksStackIDInsight":                              stack.GetScheduleAPIV1StacksStackIDInsightParams{},
	"GetScheduleAPIV1StacksStackIDInsights":                             stack.GetScheduleAPIV1StacksStackIDInsightsParams{},
	"GetScheduleAPIV1StacksStackID":                                     stack.GetScheduleAPIV1StacksStackIDParams{},
	"PostScheduleAPIV1StacksCreate":                                     stack.PostScheduleAPIV1StacksCreateParams{},
	"PostScheduleAPIV1StacksStackIDFindings":                            stack.PostScheduleAPIV1StacksStackIDFindingsParams{},
	"DeleteAuthAPIV1UserRoleBinding":                                    users.DeleteAuthAPIV1UserRoleBindingParams{},
	"GetAuthAPIV1UserRoleBindings":                                      users.GetAuthAPIV1UserRoleBindingsParams{},
	"GetAuthAPIV1UserUserID":                                            users.GetAuthAPIV1UserUserIDParams{},
	"GetAuthAPIV1Users":                                                 users.GetAuthAPIV1UsersParams{},
	"GetAuthAPIV1WorkspaceRoleBindings":                                 users.GetAuthAPIV1WorkspaceRoleBindingsParams{},
	"PostAuthAPIV1UserInvite":                                           users.PostAuthAPIV1UserInviteParams{},
	"PutAuthAPIV1UserRoleBinding":                                       users.PutAuthAPIV1UserRoleBindingParams{},
	"DeleteWorkspaceAPIV1OrganizationOrganizationID":                    workspace.DeleteWorkspaceAPIV1OrganizationOrganizationIDParams{},
	"DeleteWorkspaceAPIV1WorkspaceWorkspaceID":                          workspace.DeleteWorkspaceAPIV1WorkspaceWorkspaceIDParams{},
	"GetWorkspaceAPIV1WorkspaceCurrent":                                 workspace.GetWorkspaceAPIV1WorkspaceCurrentParams{},
	"GetWorkspaceAPIV1WorkspacesByidWorkspaceID":                        workspace.GetWorkspaceAPIV1WorkspacesByidWorkspaceIDParams{},
	"GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceID":                  workspace.GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams{},
	"GetWorkspaceAPIV1WorkspacesLimitsWorkspaceName":                    workspace.GetWorkspaceAPIV1WorkspacesLimitsWorkspaceNameParams{},
	"GetWorkspaceAPIV1Workspaces":                                       workspace.GetWorkspaceAPIV1WorkspacesParams{},
	"GetWorkspaceAPIV1WorkspacesWorkspaceID":                            workspace.GetWorkspaceAPIV1WorkspacesWorkspaceIDParams{},
	"PostWorkspaceAPIV1Organization":                                    workspace.PostWorkspaceAPIV1OrganizationParams{},
	"PostWorkspaceAPIV1Workspace":                                       workspace.PostWorkspaceAPIV1WorkspaceParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDName":                        workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDNameParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDOrganization":                workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOrganizationParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDOwner":                       workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDResume":                      workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspend":                     workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams{},
	"PostWorkspaceAPIV1WorkspaceWorkspaceIDTier":                        workspace.PostWorkspaceAPIV1WorkspaceWorkspaceIDTierParams{},
}
