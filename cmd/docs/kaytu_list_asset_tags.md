# kaytu list asset tags



```
kaytu list asset tags [flags]
```

### Options

```
      --connection-group stringArray      Connection group to filter by - mutually exclusive with connectionId
      --connection-id stringArray         Connection IDs to filter by - mutually exclusive with connectionGroup
      --connector stringArray             Connector type to filter by
      --end-time int                      End time in unix timestamp format, default now
  -h, --help                              help for tags
      --metric-type string                Metric type, default: assets
      --min-count int                     Minimum number of resources/spend with this tag value, default 1
      --resource-collection stringArray   Resource collection IDs to filter by
      --start-time int                    Start time in unix timestamp format, default now - 1 month
```

### Options inherited from parent commands

```
      --filter string           columns to display. syntax: https://github.com/teacat/jsonfilter#syntax
      --output-type string      output type [summary, json, csv, list, table] (default &amp;quot;summary&amp;quot;)
      --workspace-name string   
```

### SEE ALSO

* [kaytu list asset](kaytu_list_asset)	 - 

###### Auto generated by spf13/cobra on 2-Dec-2023