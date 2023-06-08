package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/schedule"
	"github.com/spf13/cobra"
)
var GetScheduleApiV1DescribeSourceJobsPendingCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_describe_source_jobs_pending",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_source_jobs_pending] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1DescribeSourceJobsPending(schedule.NewGetScheduleAPIV1DescribeSourceJobsPendingParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_source_jobs_pending] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_source_jobs_pending] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1SourcesCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_sources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1Sources(schedule.NewGetScheduleAPIV1SourcesParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1SummarizeJobsPendingCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_summarize_jobs_pending",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_summarize_jobs_pending] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1SummarizeJobsPending(schedule.NewGetScheduleAPIV1SummarizeJobsPendingParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_summarize_jobs_pending] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_summarize_jobs_pending] : %v", err)
		}

		return nil
	},
}
var PostScheduleApiV1SourcesSourceIdJobsComplianceRefreshCmd = &cobra.Command{
	Use: "post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh] : %v", err)
		}

		resp, err := client.Schedule.PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh(schedule.NewPostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1SourcesSourceIdJobsDescribeCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_sources_source_id_jobs_describe",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_describe] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1SourcesSourceIDJobsDescribe(schedule.NewGetScheduleAPIV1SourcesSourceIDJobsDescribeParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_describe] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_describe] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1SourcesSourceIdCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_sources_source_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1SourcesSourceID(schedule.NewGetScheduleAPIV1SourcesSourceIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id] : %v", err)
		}

		return nil
	},
}
var PostScheduleApiV1SourcesSourceIdJobsDescribeRefreshCmd = &cobra.Command{
	Use: "post_schedule_api_v_1_sources_source_id_jobs_describe_refresh",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_describe_refresh] : %v", err)
		}

		resp, err := client.Schedule.PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh(schedule.NewPostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_describe_refresh] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_describe_refresh] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1ComplianceReportLastCompletedCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_compliance_report_last_completed",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_compliance_report_last_completed] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1ComplianceReportLastCompleted(schedule.NewGetScheduleAPIV1ComplianceReportLastCompletedParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_compliance_report_last_completed] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_compliance_report_last_completed] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1DescribeResourceJobsPendingCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_describe_resource_jobs_pending",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_resource_jobs_pending] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1DescribeResourceJobsPending(schedule.NewGetScheduleAPIV1DescribeResourceJobsPendingParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_resource_jobs_pending] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_resource_jobs_pending] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1InsightJobsPendingCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_insight_jobs_pending",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_insight_jobs_pending] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1InsightJobsPending(schedule.NewGetScheduleAPIV1InsightJobsPendingParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_insight_jobs_pending] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_insight_jobs_pending] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1ResourceTypeProviderCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_resource_type_provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_resource_type_provider] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1ResourceTypeProvider(schedule.NewGetScheduleAPIV1ResourceTypeProviderParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_resource_type_provider] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_resource_type_provider] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1SourcesSourceIdJobsComplianceCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_sources_source_id_jobs_compliance",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_compliance] : %v", err)
		}

		resp, err := client.Schedule.GetScheduleAPIV1SourcesSourceIDJobsCompliance(schedule.NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_compliance] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_compliance] : %v", err)
		}

		return nil
	},
}


