[![Open in Dev Containers](https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/lreimer/tf2eksctl)

# tf2eksctl

A CLI to extract infrastructure details from Terraform state / output to generate an eksctl YAML manifest file.
The manifest can be output to stdout or a dedicated file. To fine-tune the generated manifest the CLI provides
several command line options and parameters.

## Usage

```bash
tf2eksctl
tf2eksctl version

tf2eksctl create demo-cluster --version 1.22 --region eu-central-1
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the Apache v2.0 open source license, read the `LICENSE`
file for details.
