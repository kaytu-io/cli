# kaytu list accounts summary



```
kaytu list accounts summary [flags]
```

### Options

```
      --connection-groups stringArray     Connection Groups
      --connection-id stringArray         Connection IDs
      --connector stringArray             Connector
      --end-time int                      end time in unix seconds
      --filter string                     Filter costs
      --health-state string               health state filter
  -h, --help                              help for summary
      --lifecycle-state string            lifecycle state filter
      --need-cost                         for quicker inquiry send this parameter as false, default: true
      --need-resource-count               for quicker inquiry send this parameter as false, default: true
      --page-number int                   page number - default is 1
      --page-size int                     page size - default is 20
      --resource-collection stringArray   Resource collection IDs to filter by
      --sort-by string                    column to sort by - default is cost
      --start-time int                    start time in unix seconds
```

### Options inherited from parent commands

```
      --output-type string      output type [summary, json, csv, list, table] (default &amp;quot;summary&amp;quot;)
      --workspace-name string   
```

### SEE ALSO

* [kaytu list accounts](kaytu_list_accounts)	 - 

###### Auto generated by spf13/cobra on 2-Dec-2023
