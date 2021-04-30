package diytheme

import (
	"image/color"

	"fyne.io/fyne"
)

type themedefault struct{}

var _ fyne.Theme = (*themedefault)(nil)

func (mf *themedefault) BackgroundColor() color.Color {
	return color.RGBA{0x71, 0x80, 0x93, 0xaa}
}

func (mf *themedefault) ButtonColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) DisabledButtonColor() color.Color {
	panic("not implemented") // TODO: Implement
}

// Deprecated: Hyperlinks now use the primary color for consistency.
func (mf *themedefault) HyperlinkColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) TextColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) DisabledTextColor() color.Color {
	panic("not implemented") // TODO: Implement
}

// Deprecated: Icons now use the text colour for consistency.
func (mf *themedefault) IconColor() color.Color {
	panic("not implemented") // TODO: Implement
}

// Deprecated: Disabled icons match disabled text color for consistency.
func (mf *themedefault) DisabledIconColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) PlaceHolderColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) PrimaryColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) HoverColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) FocusColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) ScrollBarColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) ShadowColor() color.Color {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) TextSize() int {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) TextFont() fyne.Resource {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) TextBoldFont() fyne.Resource {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) TextItalicFont() fyne.Resource {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) TextBoldItalicFont() fyne.Resource {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) TextMonospaceFont() fyne.Resource {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) Padding() int {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) IconInlineSize() int {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) ScrollBarSize() int {
	panic("not implemented") // TODO: Implement
}

func (mf *themedefault) ScrollBarSmallSize() int {
	panic("not implemented") // TODO: Implement
}
