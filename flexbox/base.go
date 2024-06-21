package flexbox

import "github.com/charmbracelet/lipgloss"

type base struct {
	style lipgloss.Style
}

func (r *Cell) getContentWidth() int {
	return r.getMaxWidth() - r.getExtraWidth()
}

func (r *Cell) getContentHeight() int {
	return r.getMaxHeight() - r.getExtraHeight()
}

func (r *Cell) getExtraWidth() int {
	return r.style.GetHorizontalFrameSize()
}

func (r *Cell) getExtraHeight() int {
	return r.style.GetVerticalFrameSize()
}

