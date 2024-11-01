## beamlit completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(beamlit completion bash)

To load completions for every new session, execute once:

#### Linux:

	beamlit completion bash > /etc/bash_completion.d/beamlit

#### macOS:

	beamlit completion bash > $(brew --prefix)/etc/bash_completion.d/beamlit

You will need to start a new shell for this setup to take effect.


```
beamlit completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -o, --output string      Output format. One of: pretty,yaml,json,table
  -w, --workspace string   Specify the workspace name
```

### SEE ALSO

* [beamlit completion](beamlit_completion.md)	 - Generate the autocompletion script for the specified shell

