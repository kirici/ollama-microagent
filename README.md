# Micro-agent on local models using ollama

Proof-of-concept repository.

### Requirements

- ollama (tested with v0.2.5)
- podman

### Usage

Build a container with `micro-agent` in it

```bash
./bake build
```

Start your ollama instance

```bash
./bake ollama-serve
```

Place your target file, test file and your prompt under `./mission`.

Edit `./bake` as necessary

Run a non-interactive session

```bash
./bake exec
```

See other commands

```bash
./bake --usage
```
