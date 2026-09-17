package cells

import "github.com/charmbracelet/x/ansi"

type Style uint8

var MaxStylesANSILen = len(ansi.NewStyle(
	ansi.AttrBold,
	ansi.AttrFaint,
	ansi.AttrItalic,
	ansi.AttrUnderline,
	ansi.AttrBlink,
	ansi.AttrReverse,
	ansi.AttrConceal,
	ansi.AttrStrikethrough,
).String())

func (s Style) ANSI() string {
	style := ansi.NewStyle()

	if s&Bold != 0 {
		style = style.Bold()
	}
	if s&Faint != 0 {
		style = style.Faint()
	}
	if s&Italic != 0 {
		style = style.Italic(true)
	}
	if s&Underlined != 0 {
		style = style.Underline(true)
	}
	if s&Blink != 0 {
		style = style.Blink(true)
	}
	if s&Reversed != 0 {
		style = style.Reverse(true)
	}
	if s&Concealed != 0 {
		style = style.Conceal(true)
	}
	if s&Strikeout != 0 {
		style = style.Strikethrough(true)
	}

	return style.String()
}

const (
	Bold Style = 1 << iota
	Faint
	Italic
	Underlined
	Blink
	Reversed
	Concealed
	Strikeout
)
