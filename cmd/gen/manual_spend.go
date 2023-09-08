package gen

import "github.com/kaytu-io/cli-program/cmd/gen/manual_commands"

func init() {
	SpendCmd.AddCommand(manual_commands.GetCostsCmd)
	manual_commands.GetCostsCmd.Flags().String("start-time", "", "timestamp for start time (Default one month ago)")
	manual_commands.GetCostsCmd.Flags().String("end-time", "", "timestamp for end time - (Default now)")
	manual_commands.GetCostsCmd.Flags().Int64("page-number", 0, "page number - (Default 1)")
	manual_commands.GetCostsCmd.Flags().Int64("page-size", 0, "page size - (Default 20)")
	manual_commands.GetCostsCmd.Flags().String("sort-by", "", "Sort by field - (Default cost)")
	manual_commands.GetCostsCmd.Flags().String("view", "connection", "Whether to view by connection or service")
	manual_commands.GetCostsCmd.Flags().String("filter", "", "Filter in JSON format (Case Sensitive). Filters by matches:\n"+
		"\t- Match: JSON Format. Specifies a filter rule:\n"+
		"\t\t- Key: String. Specifies the filter dimension key. Valid values are [ConnectionID|Provider|ConnectionGroup|ConnectionName]\n"+
		"\t\t- Value: String Array. Specifies the filter value.\n"+
		"\t\t- MatchOption: String. Specifies the match option. Valid values are [EQUAL|STARTS_WITH|ENDS_WITH|CONTAINS]. (Default EQUAL).\n"+
		"\t\t-              NOT the match by ~ character: [~EQUAL|~STARTS_WITH|~ENDS_WITH|~CONTAINS]\n"+
		"\t- AND: Array. Return results that match all Dimension objects.\n"+
		"\t- OR: Array. Return results that match either Dimension object.\n"+
		"Example:\n"+
		"\t {\"AND\":[{\"Match\":{\"Key\":\"ConnectionName\",\"Values\":[\"Software\"],\"MatchOption\":\"CONTAINS\"}},{\"Match\":{\"Key\":\"ConnectionName\",\"Values\":[\"T\"],\"MatchOption\":\"~STARTS_WITH\"}}]}")
}
