# kaytu completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	kaytu completion fish | source

To load completions for every new session, execute once:

	kaytu completion fish &amp;gt; ~/.config/fish/completions/kaytu.fish

You will need to start a new shell for this setup to take effect.


```
kaytu completion fish [flags]
```

### Options

```
  -h, --help              help for fish
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
