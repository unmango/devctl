# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`devctl` — Go CLI (cobra-based) for development productivity: listing source files by language, and a convention-based dependency-version system backed by plaintext files in `.versions/`.

## Commands

Build/test go through `make`, which drives itself using `devctl` (dogfooding — `bin/devctl list --go` discovers Go source files for its own build deps).

```shell
make build          # builds bin/devctl
make test           # ginkgo run, non-E2E specs only (uses .make/test as a stamp file)
make test_all        # ginkgo run -r ./ , including E2E specs
make format          # dprint fmt (README.md, .github/renovate.json, .vscode/extensions.json, .dprint.json)
make check           # nix flake check --all-systems
make tidy            # regenerates go.sum and gomod2nix.toml from go.mod
```

Run a single spec/package directly with ginkgo (installed as a go tool):

```shell
go tool ginkgo run --label-filter '!E2E' ./pkg/version/...
go tool ginkgo run -r --focus "RelPath" ./pkg/version/
```

E2E specs are tagged with the Ginkgo label `E2E` (see `test/e2e/`) and build the real binary via `gexec.Build` in `BeforeSuite` — they're excluded from `make test` (`TEST_FLAGS := --label-filter !E2E`) unless `CI` is set, but included in `make test_all`.

Tests use Ginkgo v2 + Gomega exclusively (no stdlib `testing.T` assertions). New test files/packages follow the Ginkgo bootstrap/generate convention:

```shell
cd pkg/foo && go tool ginkgo bootstrap   # creates foo_suite_test.go
cd pkg/foo && go tool ginkgo generate bar  # creates bar_test.go for bar.go
```
(Also wired as Makefile pattern rules: `%_suite_test.go` and `%_test.go`.)

No standalone lint command; formatting is via `dprint` (`.dprint.json`) for non-Go files, `gofmt`/`go vet` conventions for Go.

## Environment

Nix flake + `direnv` (`.envrc`) is the primary dev environment (`devShells.default` in `flake.nix` provides `go`, `gomod2nix`, `dprint`, `nixfmt`, etc). Building the Nix package itself runs `ginkgo run --label-filter=!E2E -r .` as its `checkPhase`.

`gomod2nix.toml` must stay in sync with `go.mod`/`go.sum` — regenerate with `go tool gomod2nix` after dependency changes (CI's `gomod2nix` job fails the PR if it drifts).

Go module tools are declared in the `tool (...)` block of `go.mod` (ginkgo, gomod2nix) and invoked via `go tool <name>` rather than being globally installed.

## Architecture

### Command layer (`cmd/`)
Cobra commands. `cmd/root.go` wires subcommands: `initialize`/`init` (`cmd/initialize/`), `config` (`cmd/config/`), `install`, `list`, `localbin`, `version`. Each subpackage under `cmd/` defines a package-level `Cmd`/`*Cmd` var plus a `New*()` constructor so commands are testable/composable independent of the global `root`.

Commands consistently resolve their working directory through `pkg/work` (either `work.Load(ctx)` — tries git root, then cwd — or an explicit `--chdir/-C` flag via `work.ChdirOptions`/`work.ChdirFlag`), rather than assuming cwd.

### `pkg/work` — the workspace abstraction
`work.Directory` is a string-typed path with `.Fs()` (returns an `afero.Fs` base-pathed at that directory) and `.Join()` helpers. Nearly every other package takes a `work.Directory` rather than a raw path, and filesystem access goes through the returned `afero.Fs` (enables afero's in-memory FS for tests, e.g. `work.WithRoot(afero.NewMemMapFs())`).

### `pkg/version` — the `.versions/` convention
Core convention: a repo's root `.versions/<name>` file holds a plaintext version string (optionally `v`-prefixed). `version.RelPath`, `version.Regex`, `version.Clean`/`Prefixed` are the low-level string helpers; `version.Cat`/`Init`/`WriteMakefile` do the file I/O and Makefile-snippet generation. `version.Source` is the interface for resolving a version from elsewhere (`version.String` is a literal; `version.GitHub` — via `github.com/unmango/aferox/github` — is stubbed/unimplemented (`panic("unimplemented")` for `Latest`/`Name`)). `version.GuessSource` picks a `Source` by pattern-matching the input (looks like a semver → `String`, contains "github" → `GitHub`).

### `pkg/config` — devctl's own config file
`devctl.yml` (or `.yaml`) in the workspace root, loaded via viper (`pkg/config.Viper`/`FromDirectory`/`Init`). Schema is just `Config{ Tools map[string]tool.Config }`.

### `pkg/tool` — tool installation
`tool.Config` (`url`, `version`, `script`) + `tool.Tool` describes an installable binary; `Tool.Install` downloads (`.tar.gz` auto-extracted via `Untar`) into a `work.Directory`'s `bin/`, chmod'ing it executable if it has no file extension.

### `pkg/list` — source file discovery
Language-aware source listing (`--go`, `--ts`, `--proto`, `--cs`, `--fs`, `--dotnet`, with `--exclude-tests`/`--absolute`) used both as a CLI feature and by the Makefile to compute Go build dependencies (`bin/devctl: $(shell $(DEVCTL) list --go --exclude-tests)`).

### `pkg/renovate` — generated schema
`pkg/renovate/zz_generated.schema.go` is generated from `.make/renovate-schema.json` (downloaded from docs.renovatebot.com, filtered by `hack/renovate/*.jq`, then run through `go-jsonschema`). Regenerate via `make pkg/renovate/zz_generated.schema.go`; don't hand-edit the `zz_generated` file. CI's `clean` job fails if this file is out of sync with its inputs.

## Conventions

- Command errors: use `github.com/unmango/go/cli.Fail(err)` in `Run` funcs (prints and exits non-zero) rather than returning errors from cobra's `RunE`.
- Logging: `github.com/charmbracelet/log`, with `ReportTimestamp: false`.
- External deps favor the `unmango/*` and `unmango/aferox/*` ecosystem (companion libraries by the same author) over reinventing filesystem/git/CLI helpers.
