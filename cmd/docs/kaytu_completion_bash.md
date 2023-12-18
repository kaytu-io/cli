# kaytu completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the &amp;apos;bash-completion&amp;apos; package.
If it is not installed already, you can install it via your OS&amp;apos;s package manager.

To load completions in your current shell session:

	source &amp;lt;(kaytu completion bash)

To load completions for every new session, execute once:

#### Linux:

	kaytu completion bash &amp;gt; /etc/bash_completion.d/kaytu

#### macOS:

	kaytu completion bash &amp;gt; $(brew --prefix)/etc/bash_completion.d/kaytu

You will need to start a new shell for this setup to take effect.


```
kaytu completion bash
```

### Options

```
  -h, --help              help for bash
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