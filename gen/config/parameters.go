package config

import (
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks_assignment"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/compliance"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/describe"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/inventory"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/keys"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/location"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/metadata"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/onboard"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/roles"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/schedule"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/services"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/smart_query"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/stack"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/users"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
)

var ParamModels = map[string]interface{}{
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
	"PostComplianceAPIV1AlarmsTop":                                      compliance.PostComplianceAPIV1AlarmsTopParams{},
	"PostComplianceAPIV1Findings":                                       compliance.PostComplianceAPIV1FindingsParams{},
	"GetOnboardAPIV1ConnectionsSummaryConnectionID":                     connections.GetOnboardAPIV1ConnectionsSummaryConnectionIDParams{},
	"GetOnboardAPIV1ConnectionsSummary":                                 connections.GetOnboardAPIV1ConnectionsSummaryParams{},
	"GetScheduleAPIV0ComplianceSummarizerTrigger":                       describe.GetScheduleAPIV0ComplianceSummarizerTriggerParams{},
	"GetScheduleAPIV0ComplianceTrigger":                                 describe.GetScheduleAPIV0ComplianceTriggerParams{},
	"GetScheduleAPIV0DescribeTrigger":                                   describe.GetScheduleAPIV0DescribeTriggerParams{},
	"GetScheduleAPIV0InsightTrigger":                                    describe.GetScheduleAPIV0InsightTriggerParams{},
	"GetScheduleAPIV0SummarizeTrigger":                                  describe.GetScheduleAPIV0SummarizeTriggerParams{},
	"GetScheduleAPIV1BenchmarkEvaluations":                              describe.GetScheduleAPIV1BenchmarkEvaluationsParams{},
	"GetScheduleAPIV1InsightJobJobID":                                   describe.GetScheduleAPIV1InsightJobJobIDParams{},
	"PostScheduleAPIV1DescribeResource":                                 describe.PostScheduleAPIV1DescribeResourceParams{},
	"PostScheduleAPIV1StacksInsightTrigger":                             describe.PostScheduleAPIV1StacksInsightTriggerParams{},
	"PutScheduleAPIV1BenchmarkEvaluationTrigger":                        describe.PutScheduleAPIV1BenchmarkEvaluationTriggerParams{},
	"PutScheduleAPIV1DescribeTriggerConnectionID":                       describe.PutScheduleAPIV1DescribeTriggerConnectionIDParams{},
	"PutScheduleAPIV1InsightEvaluationTrigger":                          describe.PutScheduleAPIV1InsightEvaluationTriggerParams{},
	"GetComplianceAPIV1InsightGroupInsightGroupID":                      insights.GetComplianceAPIV1InsightGroupInsightGroupIDParams{},
	"GetComplianceAPIV1InsightGroupInsightGroupIDTrend":                 insights.GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams{},
	"GetComplianceAPIV1InsightGroup":                                    insights.GetComplianceAPIV1InsightGroupParams{},
	"GetComplianceAPIV1InsightInsightID":                                insights.GetComplianceAPIV1InsightInsightIDParams{},
	"GetComplianceAPIV1InsightInsightIDTrend":                           insights.GetComplianceAPIV1InsightInsightIDTrendParams{},
	"GetComplianceAPIV1Insight":                                         insights.GetComplianceAPIV1InsightParams{},
	"GetComplianceAPIV1MetadataInsightInsightID":                        insights.GetComplianceAPIV1MetadataInsightInsightIDParams{},
	"GetComplianceAPIV1MetadataInsight":                                 insights.GetComplianceAPIV1MetadataInsightParams{},
	"GetComplianceAPIV1MetadataTagInsightKey":                           insights.GetComplianceAPIV1MetadataTagInsightKeyParams{},
	"GetComplianceAPIV1MetadataTagInsight":                              insights.GetComplianceAPIV1MetadataTagInsightParams{},
	"GetInventoryAPIV2CostComposition":                                  inventory.GetInventoryAPIV2CostCompositionParams{},
	"GetInventoryAPIV2CostMetric":                                       inventory.GetInventoryAPIV2CostMetricParams{},
	"GetInventoryAPIV2CostTrend":                                        inventory.GetInventoryAPIV2CostTrendParams{},
	"GetInventoryAPIV2ResourcesCompositionKey":                          inventory.GetInventoryAPIV2ResourcesCompositionKeyParams{},
	"GetInventoryAPIV2ResourcesMetric":                                  inventory.GetInventoryAPIV2ResourcesMetricParams{},
	"GetInventoryAPIV2ResourcesMetricResourceType":                      inventory.GetInventoryAPIV2ResourcesMetricResourceTypeParams{},
	"GetInventoryAPIV2ResourcesTagKey":                                  inventory.GetInventoryAPIV2ResourcesTagKeyParams{},
	"GetInventoryAPIV2ResourcesTag":                                     inventory.GetInventoryAPIV2ResourcesTagParams{},
	"GetInventoryAPIV2ResourcesTrend":                                   inventory.GetInventoryAPIV2ResourcesTrendParams{},
	"GetInventoryAPIV2ServicesCostTrend":                                inventory.GetInventoryAPIV2ServicesCostTrendParams{},
	"GetInventoryAPIV2ServicesMetric":                                   inventory.GetInventoryAPIV2ServicesMetricParams{},
	"GetInventoryAPIV2ServicesMetricServiceName":                        inventory.GetInventoryAPIV2ServicesMetricServiceNameParams{},
	"GetInventoryAPIV2ServicesTagKey":                                   inventory.GetInventoryAPIV2ServicesTagKeyParams{},
	"GetInventoryAPIV2ServicesTag":                                      inventory.GetInventoryAPIV2ServicesTagParams{},
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
	"GetInventoryAPIV1ResourcesRegions":                                 resource.GetInventoryAPIV1ResourcesRegionsParams{},
	"GetInventoryAPIV1ResourcesTopRegions":                              resource.GetInventoryAPIV1ResourcesTopRegionsParams{},
	"GetInventoryAPIV2ResourcesCount":                                   resource.GetInventoryAPIV2ResourcesCountParams{},
	"GetInventoryAPIV2ResourcesRegionsComposition":                      resource.GetInventoryAPIV2ResourcesRegionsCompositionParams{},
	"GetInventoryAPIV2ResourcesRegionsSummary":                          resource.GetInventoryAPIV2ResourcesRegionsSummaryParams{},
	"GetInventoryAPIV2ResourcesRegionsTrend":                            resource.GetInventoryAPIV2ResourcesRegionsTrendParams{},
	"PostInventoryAPIV1Resource":                                        resource.PostInventoryAPIV1ResourceParams{},
	"PostInventoryAPIV1ResourcesAws":                                    resource.PostInventoryAPIV1ResourcesAwsParams{},
	"PostInventoryAPIV1ResourcesAzure":                                  resource.PostInventoryAPIV1ResourcesAzureParams{},
	"PostInventoryAPIV1ResourcesFilters":                                resource.PostInventoryAPIV1ResourcesFiltersParams{},
	"PostInventoryAPIV1Resources":                                       resource.PostInventoryAPIV1ResourcesParams{},
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
	"GetInventoryAPIV2ServicesSummary":                                  services.GetInventoryAPIV2ServicesSummaryParams{},
	"GetInventoryAPIV2ServicesSummaryServiceName":                       services.GetInventoryAPIV2ServicesSummaryServiceNameParams{},
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
	"PostScheduleAPIV1StacksDescriberTrigger":                           stack.PostScheduleAPIV1StacksDescriberTriggerParams{},
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
	"GetWorkspaceAPIV1WorkspaceCurrent":                                 workspace.GetWorkspaceAPIV1WorkspaceCurrentParams{},
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
