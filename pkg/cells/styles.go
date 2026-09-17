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

	if s&StyleBold != 0 {
		style = style.Bold()
	}
	if s&StyleFaint != 0 {
		style = style.Faint()
	}
	if s&StyleItalic != 0 {
		style = style.Italic(true)
	}
	if s&StyleUnderlined != 0 {
		style = style.Underline(true)
	}
	if s&StyleBlink != 0 {
		style = style.Blink(true)
	}
	if s&StyleReversed != 0 {
		style = style.Reverse(true)
	}
	if s&StyleConcealed != 0 {
		style = style.Conceal(true)
	}
	if s&StyleStrikeout != 0 {
		style = style.Strikethrough(true)
	}

	return style.String()
}

const (
	StyleBold Style = 1 << iota
	StyleFaint
	StyleItalic
	StyleUnderlined
	StyleBlink
	StyleReversed
	StyleConcealed
	StyleStrikeout
)
