---
title: gifgrep
permalink: /
description: "gifgrep is a tiny Go CLI/TUI that searches animated GIFs from GIPHY or KLIPY, pipes URLs/JSON into your scripts, and previews results inline in Kitty, Ghostty, or iTerm2."
---

![gifgrep TUI with inline animated previews](assets/gifgrep-tui.mp4)

```console
$ gifgrep cats -m 3
https://media.giphy.com/.../cat-typing.gif
https://media.giphy.com/.../cat-yes.gif
https://media.giphy.com/.../cat-stare.gif

$ gifgrep tui "office handshake"
   ┌─ gifgrep ───────────────────────┐
   │ /office handshake               │
   │ ▸ Pam & Jim handshake.gif       │
   │   The Office handshake meme.gif │
   │   Schrute approves.gif          │
   │   [animated preview here]       │
   └─────────────────────────────────┘
```

## Try it

```bash
brew install steipete/tap/gifgrep
gifgrep cats           # plain on TTY, URLs in pipes
gifgrep tui cats       # interactive browser with inline previews
```

`--json` prints a stable envelope on stdout. Human progress and warnings go to stderr so pipes stay clean.

## What gifgrep does

- **One binary, two ergonomics.** `gifgrep <query>` for shell pipelines; `gifgrep tui` for arrow-key browsing with animated inline previews.
- **Multi-provider.** [GIPHY](providers/giphy.md) (preferred when keyed), [KLIPY](providers/klipy.md), or [`auto`](providers/auto.md) which picks the best available.
- **Inline previews that actually animate.** Kitty graphics for Kitty/Ghostty, OSC 1337 for iTerm2 — see [Previews](previews.md).
- **Local frame extraction.** [`still`](still.md) pulls one frame; [`sheet`](sheet.md) lays out a contact sheet of N frames, no provider required.
- **Pipe-shaped output.** Plain on TTY, URL-per-line in pipes, plus `--format md|tsv|comment|url|json` for the rest.
- **Quiet by default, loud on demand.** `-v`, `-vv`, `-q`, `--no-color` if you really want.

## Pick your path

- **Trying it out.** [Install](install.md) → [Quickstart](quickstart.md). Two minutes from `brew install` to your first GIF in stdout.
- **Wiring it into shell scripts.** [Search](search.md) for the CLI surface, [JSON output](json.md) for structured data, [Providers](providers/) to choose a backend.
- **Browsing visually.** [TUI](tui.md) for the keyboard-driven browser, [Previews](previews.md) for the terminal protocol details.
- **Cutting frames out of a GIF.** [`still`](still.md) for one frame, [`sheet`](sheet.md) for a contact sheet.
- **Looking up a flag.** Every command has its own page — start at [Commands](commands.md).

## Project

It's Go. It's tiny. It cleans up after itself. (Or at least it tries.)

- Source: [github.com/steipete/gifgrep](https://github.com/steipete/gifgrep)
- Changelog: [CHANGELOG.md](https://github.com/steipete/gifgrep/blob/main/CHANGELOG.md)
- License: [MIT](https://github.com/steipete/gifgrep/blob/main/LICENSE)
- GIFs courtesy of [GIPHY](https://giphy.com) and [KLIPY](https://klipy.com). Not affiliated with either.

If you somehow manage to grep the wrong GIF: that's on you. ❤️
