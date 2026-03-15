# DevPod Provider for Sprites.dev

A [DevPod](https://devpod.sh) provider that creates and manages development environments on [Sprites.dev](https://sprites.dev).

## Installation

### From GitHub Release

```sh
devpod provider add github.com/team-alembic/devpod-sprites
```

### From Local Source

Build the binary and add the local provider:

```sh
go build -o devpod-provider-sprites
devpod provider add .
```

## Configuration

| Option | Required | Description |
|---|---|---|
| `SPRITE_TOKEN` | Yes | Your Sprites.dev API token |
| `AGENT_PATH` | No | Path for the DevPod agent on the sprite (default: `/tmp/devpod`) |

Set your token when adding the provider:

```sh
devpod provider add github.com/team-alembic/devpod-sprites
devpod provider set-options sprites -o SPRITE_TOKEN=your-token-here
```

## Usage

```sh
devpod up my-workspace --provider sprites
```

## Releasing

Push a version tag to trigger the GitHub Actions release workflow:

```sh
git tag v0.1.0
git push origin v0.1.0
```

This builds binaries for linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, and windows/amd64, then publishes them as a GitHub release with a generated `provider.yaml`.

## Licence

Copyright 2026 Alembic Pty Ltd

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.
