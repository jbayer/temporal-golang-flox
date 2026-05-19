# Temporal Go Getting Started with Flox

A simple Temporal greeting workflow in Go, packaged with a [Flox](https://flox.dev) environment that provides everything you need to get started. The Flox environment works across macOS and Linux (including Windows WLS2) without modification -- no manual dependency installation, no platform-specific setup such as with the [Temporal Docs Golang Tutorial](https://docs.temporal.io/develop/go/set-up-your-local-go).

## Prerequisites

Install Flox 1.12.0 or newer [flox.dev/download](https://flox.dev/download)

## Quick Start

**Desktop**

```bash
git clone https://github.com/jbayer/temporal-golang-flox.git
cd temporal-golang-flox
flox activate
╭─────────────────────────────────────╮
│                                     │
│  Temporal Golang Getting Started    │
│                                     │
│  Check service status:              │
│    flox services status             │
│                                     │
│  View service logs:                 │
│    flox services logs temporal      │
│    flox services logs worker        │
│                                     │
│  Send a hello workflow:             │
│    go run ./start/main.go YourName  │
│                                     │
│  Open the Temporal Web UI:          │
│    http://127.0.0.1:8233            │
│                                     │
╰─────────────────────────────────────╯

✔ You are now using the environment 'temporal'
To stop using this environment, type 'exit'
```

**Devcontainer**

Open in VSCode or other tooling such as `devcontainer` that has devcontainer support to run in a container with Flox.

```bash
devcontainer up --workspace-folder . && \
INSIDE_DEVCONTAINER=1 devcontainer exec --workspace-folder . bash
```

That's it. When you activate, Flox installs all dependencies (Go, Temporal CLI, etc.), starts the Temporal dev server and a worker process as background services, and prints instructions for what to do next.

## Services

The environment defines two services that start automatically on activation:

- **temporal** -- a Temporal dev server with SQLite storage and the built-in Web UI
- **worker** -- a Go worker that listens on `my-task-queue` and executes the greeting workflow

## Usage

Check service status:

```bash
flox services status
```

View service logs:

```bash
flox services logs temporal
flox services logs worker
```

Send a hello workflow:

```bash
go run ./start/main.go YourName
```

Open the Temporal Web UI:

```
http://127.0.0.1:8233
```

## How It Works

The entire environment is defined in `.flox/env/manifest.toml`. This single file declares the packages to install, environment variables, startup hooks, and service definitions. Consumers of this repo don't need to install Go, Temporal Server, Temporal CLI, or anything else by hand -- `flox activate` handles it all and gives them a shell with everything they need to be productive, ready to go.

## Cleanup

Exiting the shell where Flox was activated will shut down the temporal server and worker.

## Uninstall

You an uninstall Flox and all of the related software with [this uninstall documentation](https://flox.dev/docs/install-flox/uninstall/?h=uninstall).
