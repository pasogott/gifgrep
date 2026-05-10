package termcaps

import "testing"

func TestDetectInlineOverride(t *testing.T) {
	getenv := func(k string) string {
		switch k {
		case "GIFGREP_INLINE":
			return "iterm"
		default:
			return ""
		}
	}
	if got := DetectInline(getenv); got != InlineIterm {
		t.Fatalf("expected iterm, got %v", got)
	}
}

func TestDetectInlineSixelOverride(t *testing.T) {
	getenv := func(k string) string {
		switch k {
		case "GIFGREP_INLINE":
			return "sixel"
		default:
			return ""
		}
	}
	if got := DetectInline(getenv); got != InlineSixel {
		t.Fatalf("expected sixel, got %v", got)
	}
}

func TestDetectInlineANSIOverride(t *testing.T) {
	getenv := func(k string) string {
		if k == "GIFGREP_INLINE" {
			return "ansi"
		}
		return ""
	}
	if got := DetectInline(getenv); got != InlineANSI {
		t.Fatalf("expected ansi, got %v", got)
	}
}

func TestDetectInlineKittyEnv(t *testing.T) {
	getenv := func(k string) string {
		switch k {
		case "KITTY_WINDOW_ID":
			return "123"
		case "TERM_PROGRAM":
			return "iTerm.app"
		default:
			return ""
		}
	}
	if got := DetectInline(getenv); got != InlineKitty {
		t.Fatalf("expected kitty, got %v", got)
	}
}

func TestDetectInlineWindowsTerminalUsesSixel(t *testing.T) {
	getenv := func(k string) string {
		switch k {
		case "WT_SESSION":
			return "abc"
		default:
			return ""
		}
	}
	if got := DetectInline(getenv); got != InlineSixel {
		t.Fatalf("expected sixel, got %v", got)
	}
}

func TestDetectInlineItermEnv(t *testing.T) {
	getenv := func(k string) string {
		switch k {
		case "TERM_PROGRAM":
			return "iTerm.app"
		default:
			return ""
		}
	}
	if got := DetectInline(getenv); got != InlineIterm {
		t.Fatalf("expected iterm, got %v", got)
	}
}

func TestDetectInlineRobustProbeNegative(t *testing.T) {
	getenv := func(k string) string {
		switch k {
		case "TERM":
			return "xterm-kitty"
		default:
			return ""
		}
	}
	got := detectInlineRobust(getenv, func() kittyProbeResult { return kittyProbeNotSupported })
	if got != InlineNone {
		t.Fatalf("expected none, got %v", got)
	}
}

func TestDetectInlineRobustProbeUnknownKeepsKitty(t *testing.T) {
	getenv := func(k string) string {
		switch k {
		case "TERM":
			return "xterm-kitty"
		default:
			return ""
		}
	}
	got := detectInlineRobust(getenv, func() kittyProbeResult { return kittyProbeUnknown })
	if got != InlineKitty {
		t.Fatalf("expected kitty, got %v", got)
	}
}
