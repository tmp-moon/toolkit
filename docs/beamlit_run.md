## beamlit run

Run inference

```
beamlit run [model] [environment] [flags]
```

### Options

```
      --data string          JSON body data for the inference request
      --env string           Environment to run the inference in (default "production")
      --header stringArray   Request headers in 'Key: Value' format. Can be specified multiple times
  -h, --help                 help for run
      --method string        HTTP method for the inference request (default "POST")
      --path string          path for the inference request
      --show-headers         Show response headers in output
```

### Options inherited from parent commands

```
  -o, --output string      Output format. One of: pretty,yaml,json,table
  -w, --workspace string   Specify the workspace name
```

### SEE ALSO

* [beamlit](beamlit.md)	 - Beamlit CLI is a command line tool to interact with Beamlit APIs.

