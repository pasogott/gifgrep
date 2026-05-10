---
title: KLIPY provider
description: "Use KLIPY as gifgrep's GIF backend."
---

# KLIPY

KLIPY is the second supported provider, and the default fallback for [`auto`](auto.md) when `GIPHY_API_KEY` is not set.

> KLIPY replaced Tenor as gifgrep's secondary backend. `--source tenor` is kept as a legacy alias and routes through KLIPY using `KLIPY_API_KEY`.

## Enable

```bash
export KLIPY_API_KEY=your-key
gifgrep --source klipy cats
gifgrep --source tenor cats   # alias — same backend
```

## Get a key

Sign up at <https://klipy.com> and grab a developer key. Then:

```bash
echo 'export KLIPY_API_KEY=...' >> ~/.zshrc
```

## What gifgrep sends

A KLIPY v2 search call:

```text
GET https://api.klipy.com/v2/search
    ?q=<query>
    &key=<key>
    &limit=<--max>
    &contentfilter=low
    &media_filter=gif,tinygif,mediumgif,nanogif,preview
```

Responses are normalised into the [`Result`](../json.md#fields) shape:

- `url` ← preferred media format: `gif` → `mediumgif` → `tinygif` → `nanogif` → `preview`.
- `preview_url` ← preferred small format: `tinygif` → `nanogif` → `preview` → `mediumgif` → `gif`.
- `width` / `height` ← `dims` of the chosen media format (when reported).
- `tags` ← KLIPY's first-class `tags` array.

## Limits and content

- gifgrep requests `contentfilter=low`. Adjust by switching to GIPHY (which uses `rating=g`) or by patching the source if you self-host.
- KLIPY tends to ship richer `tags` arrays than GIPHY, which is handy for filtering in JSON pipelines.

## When to choose KLIPY explicitly

```bash
gifgrep --source klipy "office handshake"
```

- Free-tier friendly for quick prototyping.
- Better tag metadata for JSON / `jq` filtering.
- Different catalogue — useful when GIPHY misses a niche query.

## Tenor compatibility

`gifgrep` previously shipped a Tenor backend. That's been retired; `--source tenor` now resolves to KLIPY at lookup time:

```bash
gifgrep --source tenor cats   # works, talks to KLIPY
```

If `KLIPY_API_KEY` isn't set, `--source tenor` errors the same way `--source klipy` would.

## See also

- [`auto`](auto.md) — auto-resolution rules.
- [GIPHY](giphy.md) — the other provider.
- [Search](../search.md) — full CLI surface.
