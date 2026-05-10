---
title: Providers
description: "GIPHY vs KLIPY vs auto — picking a backend for gifgrep."
---

# Providers

`gifgrep` talks to multiple GIF providers. Pick one explicitly with `--source`, or let it auto-resolve.

| Source             | Backend             | API key required           |
|--------------------|---------------------|----------------------------|
| [`auto`](auto.md)  | First available     | At least one of the below  |
| [`giphy`](giphy.md)| GIPHY v1 search     | `GIPHY_API_KEY`            |
| [`klipy`](klipy.md)| KLIPY v2 search     | `KLIPY_API_KEY`            |
| `tenor`            | Alias for `klipy`   | `KLIPY_API_KEY`            |

```bash
gifgrep --source auto cats         # default
gifgrep --source giphy cats
gifgrep --source klipy cats
gifgrep --source tenor cats        # legacy alias for KLIPY
```

`tenor` is kept as an alias because Tenor used to be a first-class backend; it now resolves to KLIPY and uses `KLIPY_API_KEY`.

## How `auto` decides

1. If `GIPHY_API_KEY` is set → try GIPHY first.
2. If GIPHY fails *and* `KLIPY_API_KEY` is set → fall back to KLIPY.
3. If `GIPHY_API_KEY` is not set → KLIPY.
4. If neither is set → error on stderr, non-zero exit.

Full details: [`auto`](auto.md).

## Output is identical

All providers normalise into the same [`Result`](../json.md#fields) shape — same fields, same envelope. You can swap providers in a script without changing downstream parsing.

## Why two providers?

- **GIPHY** has the deepest catalogue and best coverage of meme culture. Recommended when you can get a key.
- **KLIPY** is friendlier for free-tier prototyping and works without sign-up gymnastics in some regions.

If you don't care: leave it on `auto` and set whichever key you already have.

## Future providers?

The provider interface is small (`Search(query, opts) → []Result`). PRs welcome — see `internal/search/` in the repo.
