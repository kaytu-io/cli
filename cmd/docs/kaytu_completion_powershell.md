# kaytu completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	kaytu completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
kaytu completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --filter string           columns to display. syntax: https://github.com/teacat/jsonfilter#syntax
      --output-type string      output type [summary, json, csv, list, table] (default &amp;quot;summary&amp;quot;)
      --workspace-name string   
```

### SEE ALSO

* [kaytu completion](kaytu_completion)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 2-Dec-2023
