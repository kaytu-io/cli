package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/schedule"
	"github.com/spf13/cobra"
)

var GetScheduleApiV1ComplianceReportLastCompletedCmd = &cobra.Command{
	Use: "last-compliance-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_compliance_report_last_completed] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1ComplianceReportLastCompletedParams()

		resp, err := client.Schedule.GetScheduleAPIV1ComplianceReportLastCompleted(req, auth)
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
	Use: "pending-describe-resource-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_resource_jobs_pending] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1DescribeResourceJobsPendingParams()

		resp, err := client.Schedule.GetScheduleAPIV1DescribeResourceJobsPending(req, auth)
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

var GetScheduleApiV1DescribeSourceJobsPendingCmd = &cobra.Command{
	Use: "pending-describe-source-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_describe_source_jobs_pending] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1DescribeSourceJobsPendingParams()

		resp, err := client.Schedule.GetScheduleAPIV1DescribeSourceJobsPending(req, auth)
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

var GetScheduleApiV1InsightJobsPendingCmd = &cobra.Command{
	Use: "pending-insight-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_insight_jobs_pending] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1InsightJobsPendingParams()

		resp, err := client.Schedule.GetScheduleAPIV1InsightJobsPending(req, auth)
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
	Use: "provider-resource-types",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_resource_type_provider] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1ResourceTypeProviderParams()

		req.SetProvider(flags.ReadStringFlag(cmd, "Provider"))

		resp, err := client.Schedule.GetScheduleAPIV1ResourceTypeProvider(req, auth)
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

var GetScheduleApiV1SourcesSourceIdCmd = &cobra.Command{
	Use: "get-source",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1SourcesSourceIDParams()

		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		resp, err := client.Schedule.GetScheduleAPIV1SourcesSourceID(req, auth)
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

var GetScheduleApiV1SummarizeJobsPendingCmd = &cobra.Command{
	Use: "pending-summarize-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_summarize_jobs_pending] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1SummarizeJobsPendingParams()

		resp, err := client.Schedule.GetScheduleAPIV1SummarizeJobsPending(req, auth)
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

var GetScheduleApiV1SourcesCmd = &cobra.Command{
	Use: "list-sources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1SourcesParams()

		resp, err := client.Schedule.GetScheduleAPIV1Sources(req, auth)
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

var GetScheduleApiV1SourcesSourceIdJobsComplianceCmd = &cobra.Command{
	Use: "source-compliance-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_compliance] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParams()

		req.SetFrom(flags.ReadInt64OptionalFlag(cmd, "From"))
		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))
		req.SetTo(flags.ReadInt64OptionalFlag(cmd, "To"))

		resp, err := client.Schedule.GetScheduleAPIV1SourcesSourceIDJobsCompliance(req, auth)
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

var GetScheduleApiV1SourcesSourceIdJobsDescribeCmd = &cobra.Command{
	Use: "source-describe-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_describe] : %v", err)
		}

		req := schedule.NewGetScheduleAPIV1SourcesSourceIDJobsDescribeParams()

		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		resp, err := client.Schedule.GetScheduleAPIV1SourcesSourceIDJobsDescribe(req, auth)
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

var PostScheduleApiV1SourcesSourceIdJobsComplianceRefreshCmd = &cobra.Command{
	Use: "refresh-source-compliance-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh] : %v", err)
		}

		req := schedule.NewPostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshParams()

		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		_, err = client.Schedule.PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh] : %v", err)
		}

		return nil
	},
}
var PostScheduleApiV1SourcesSourceIdJobsDescribeRefreshCmd = &cobra.Command{
	Use: "refresh-source-describe-jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_describe_refresh] : %v", err)
		}

		req := schedule.NewPostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshParams()

		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		_, err = client.Schedule.PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_describe_refresh] : %v", err)
		}

		return nil
	},
}
