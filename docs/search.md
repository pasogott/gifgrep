---
title: Search
description: "gifgrep search — pipe-friendly GIF search from your terminal."
---

# `gifgrep search`

Search a provider for GIFs and print results. This is what `gifgrep <query>` runs by default.

```text
gifgrep <query> ... [flags]
gifgrep search <query> ... [flags]
```

## At a glance

```bash
gifgrep cats                              # plain on TTY, URL-per-line in pipes
gifgrep cats --max 5                      # cap results
gifgrep cats --json | jq '.[0].url'       # structured
gifgrep cats --format md                  # markdown image links
gifgrep cats --thumbs                     # inline still thumbs (Kitty/iTerm2 TTY)
gifgrep cats --download --max 1           # save to ~/Downloads
gifgrep --source giphy "office handshake" # force GIPHY
```

## Output formats

| Mode               | What you get                                                  |
|--------------------|---------------------------------------------------------------|
| `--format auto`    | Plain readable list on TTY, URL-per-line on pipes (default).  |
| `--format url`     | One URL per line. Best for `xargs`, `wget`, or `pbcopy`.      |
| `--format plain`   | `title — url` per line, no decoration.                        |
| `--format tsv`     | `id<TAB>title<TAB>url<TAB>preview_url<TAB>w<TAB>h`.           |
| `--format md`      | `![title](url)` markdown image lines.                         |
| `--format comment` | URLs prefixed with `# title` for clipboard-friendly snippets. |
| `--format json`    | Same envelope as `--json`.                                    |
| `--json`           | Pretty-printed JSON array — see [JSON output](json.md).       |

`-n` / `--number` prefixes each line with its 1-based index.

## Common flags

```text
--source <auto|giphy|klipy|tenor>   choose a provider (default: auto)
--max, -m <N>                       max results (default 20)
--json                              pretty JSON array
--format <auto|...>                 see above
--number, -n                        prefix lines with 1-based index
--thumbs[=auto|on|off]              inline still thumbs (Kitty/iTerm2, TTY only)
--download                          save results to ~/Downloads
--reveal                            after --download, open the folder
--color <auto|always|never>         color output
--no-color                          shorthand for --color=never
-v / -vv                            stderr debug logs
-q                                  quiet
```

## Inline thumbnails (CLI)

`--thumbs` shows a single still frame next to each result, decoded locally and uploaded via the [Kitty graphics protocol](previews.md#kitty-graphics) (Kitty/Ghostty) or [OSC 1337](previews.md#iterm2-osc-1337) (iTerm2). Notes:

- TTY only — pipes never receive image bytes.
- One **still** frame per row. For animated previews, use the [TUI](tui.md).
- `auto` enables thumbs only when the terminal is detected as supporting inline images.

## Downloading

```bash
gifgrep cats --download --max 3
gifgrep cats --download --reveal --max 1   # also opens Finder/Explorer
```

Downloads land in `~/Downloads`. Filenames are derived from the result title (sanitized) plus the provider's id.

## Pipe recipes

```bash
# Copy first URL to clipboard (macOS)
gifgrep "shipped it" --format url --max 1 | pbcopy

# Drop top 5 URLs into a markdown file
gifgrep "ship it" --format md --max 5 >> notes.md

# Bulk download
gifgrep "deploy" --format url --max 10 | xargs -n1 curl -O

# JSON pipeline with jq
gifgrep "office handshake" --json --max 20 | jq '.[] | select(.width > 400) | .url'
```

## Provider selection

```bash
gifgrep --source auto cats         # auto-pick: GIPHY when keyed, else KLIPY
gifgrep --source giphy cats        # force GIPHY (needs GIPHY_API_KEY)
gifgrep --source klipy cats        # force KLIPY (needs KLIPY_API_KEY)
gifgrep --source tenor cats        # legacy alias for KLIPY
```

Full provider matrix: [Providers](providers/).

## Exit behavior

- `0` — at least one result was printed.
- non-zero — provider error, missing API key, or transport failure (logged to stderr).

Human progress and errors always go to stderr, so `--json` and `--format url` pipes stay parseable.
