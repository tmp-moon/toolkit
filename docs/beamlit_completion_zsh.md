## beamlit completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(beamlit completion zsh)

To load completions for every new session, execute once:

#### Linux:

	beamlit completion zsh > "${fpath[1]}/_beamlit"

#### macOS:

	beamlit completion zsh > $(brew --prefix)/share/zsh/site-functions/_beamlit

You will need to start a new shell for this setup to take effect.


```
beamlit completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -o, --output string      Output format. One of: pretty,yaml,json,table
  -w, --workspace string   Specify the workspace name
```

### SEE ALSO

* [beamlit completion](beamlit_completion.md)	 - Generate the autocompletion script for the specified shell

