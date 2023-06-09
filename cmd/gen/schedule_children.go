package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/schedule"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var GetScheduleApiV1ResourceTypeProviderCmd = &cobra.Command{
	Use: "resource-type-provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_resource_type_provider] : %v", err)
		}

        req := schedule.NewGetScheduleAPIV1ResourceTypeProviderParams()

        req.SetProvider(flags.ReadStringFlag("Provider"))


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

var GetScheduleApiV1SourcesCmd = &cobra.Command{
	Use: "sources",
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

var GetScheduleApiV1SourcesSourceIdCmd = &cobra.Command{
	Use: "sources-source-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id] : %v", err)
		}

        req := schedule.NewGetScheduleAPIV1SourcesSourceIDParams()

        req.SetSourceID(flags.ReadStringFlag("SourceID"))


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

var PostScheduleApiV1SourcesSourceIdJobsDescribeRefreshCmd = &cobra.Command{
	Use: "sources-source-id-jobs-describe-refresh",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_describe_refresh] : %v", err)
		}

        req := schedule.NewPostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshParams()

        req.SetSourceID(flags.ReadStringFlag("SourceID"))


		_, err = client.Schedule.PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_describe_refresh] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1DescribeResourceJobsPendingCmd = &cobra.Command{
	Use: "describe-resource-jobs-pending",
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
	Use: "describe-source-jobs-pending",
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
	Use: "insight-jobs-pending",
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

var GetScheduleApiV1SummarizeJobsPendingCmd = &cobra.Command{
	Use: "summarize-jobs-pending",
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

var PostScheduleApiV1SourcesSourceIdJobsComplianceRefreshCmd = &cobra.Command{
	Use: "sources-source-id-jobs-compliance-refresh",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh] : %v", err)
		}

        req := schedule.NewPostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshParams()

        req.SetSourceID(flags.ReadStringFlag("SourceID"))


		_, err = client.Schedule.PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_sources_source_id_jobs_compliance_refresh] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1ComplianceReportLastCompletedCmd = &cobra.Command{
	Use: "compliance-report-last-completed",
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

var GetScheduleApiV1SourcesSourceIdJobsComplianceCmd = &cobra.Command{
	Use: "sources-source-id-jobs-compliance",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_compliance] : %v", err)
		}

        req := schedule.NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParams()

        req.SetFrom(flags.ReadInt64OptionalFlag("From"))
req.SetSourceID(flags.ReadStringFlag("SourceID"))
req.SetTo(flags.ReadInt64OptionalFlag("To"))


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
	Use: "sources-source-id-jobs-describe",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_sources_source_id_jobs_describe] : %v", err)
		}

        req := schedule.NewGetScheduleAPIV1SourcesSourceIDJobsDescribeParams()

        req.SetSourceID(flags.ReadStringFlag("SourceID"))


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
