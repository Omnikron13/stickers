package flexbox

import "github.com/charmbracelet/lipgloss"

type base struct {
	style lipgloss.Style
	width   int
	height  int
}

func (r base) getContentWidth() int {
	return r.getMaxWidth() - r.getExtraWidth()
}

func (r base) getContentHeight() int {
	return r.getMaxHeight() - r.getExtraHeight()
}

func (r base) getExtraWidth() int {
	return r.style.GetHorizontalFrameSize()
}

func (r base) getExtraHeight() int {
	return r.style.GetVerticalFrameSize()
}

func (r base) getMaxWidth() int {
	return r.width
}

func (r base) getMaxHeight() int {
	return r.height
}

