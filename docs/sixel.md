# Sixel inline images protocol (gifgrep)

Sixel terminals render bitmap data from DEC Sixel escape sequences. gifgrep uses
Sixel for inline still thumbnails and TUI previews when detection selects it.

Known useful targets:

- Windows Terminal 1.22+
- WezTerm
- xterm with sixel graphics enabled
- mintty / MSYS2 terminals

Detection:

- `WT_SESSION` selects Sixel for Windows Terminal.
- `TERM_PROGRAM=WezTerm` selects Sixel.
- `TERM` values containing `sixel` select Sixel.
- `GIFGREP_INLINE=sixel` forces it.

Unlike Kitty graphics, Sixel has no gifgrep-managed image id to delete or native
animation timeline. gifgrep redraws decoded PNG frames into the same cell area
for animated TUI previews.

References:

- Windows Terminal Preview 1.22 release: `https://devblogs.microsoft.com/commandline/windows-terminal-preview-1-22-release/`
- PowerShell Sixel module: `https://www.powershellgallery.com/packages/Sixel`
