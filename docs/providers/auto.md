---
title: auto provider
description: "How gifgrep's --source auto picks a backend."
---

# `--source auto`

The default. Picks the best available backend from your environment, so you don't have to think about it.

```bash
gifgrep cats                  # implicitly --source auto
gifgrep --source auto cats    # same thing
```

## Resolution

```text
GIPHY_API_KEY set?
├── yes → try GIPHY
│         ├── ok          → return GIPHY results
│         └── error  ┐
│                    ├── KLIPY_API_KEY set? → fall back to KLIPY
│                    └── otherwise         → propagate error
└── no  → KLIPY_API_KEY set?
          ├── yes → use KLIPY
          └── no  → error: missing key
```

In short:

1. **GIPHY when keyed.** If `GIPHY_API_KEY` is set, GIPHY wins.
2. **KLIPY as fallback.** If GIPHY fails (rate limit, network, key revoked) and `KLIPY_API_KEY` is also set, gifgrep retries with KLIPY without telling you to.
3. **KLIPY when GIPHY isn't configured.** No GIPHY key? `auto` resolves to KLIPY.
4. **Hard error otherwise.** Neither key set → exit non-zero with a message on stderr.

## Why this default

- It removes a flag from your muscle memory: just run `gifgrep cats`.
- It survives one provider going down without changing your shell history.
- Output format is identical across providers (see [JSON](../json.md)), so scripts don't care which one served a result.

## Forcing a provider

Bypass `auto` when you need determinism:

```bash
gifgrep --source giphy cats     # always GIPHY (errors if key missing)
gifgrep --source klipy cats     # always KLIPY (errors if key missing)
gifgrep --source tenor cats     # alias for klipy
```

For CI or agent use, prefer an explicit `--source` to keep results reproducible.

## See also

- [GIPHY](giphy.md)
- [KLIPY](klipy.md)
- [Search](../search.md)
