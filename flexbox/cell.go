package flexbox

import (
	"github.com/charmbracelet/lipgloss"
)

// Cell is a building block object of the FlexBox, it represents a single cell within a box
// A FlexBox stacks cells horizonally.
// A HorizontalFlexBox stacks cells vertically. (controverse, isn't it?)
type Cell struct {
	base;

	// id of the cell, if not set it will default to the index in the row
	id string

	// TODO: all ratios and sizes should be uint
	// ratioX width ratio of the cell
	ratioX int
	// ratioY height ratio of the cell
	ratioY int
	// minWidth minimal width of the cell
	minWidth int
	// minHeigth minimal heigth of the cell
	minHeigth int

	content string
}

// NewCell initialize FlexBoxCell object with defaults
func NewCell(ratioX, ratioY int) *Cell {
   cell := &Cell{
		ratioX:   ratioX,
		ratioY:   ratioY,
		minWidth: 0,
	}

   cell.style = lipgloss.NewStyle()
   cell.height = 0
   cell.width = 0

   return cell
}

// SetID sets the cells ID
func (r *Cell) SetID(id string) *Cell {
	r.id = id
	return r
}

// SetContent sets the cells content
func (r *Cell) SetContent(content string) *Cell {
	r.content = content
	return r
}

// GetContent returns the cells raw content
func (r *Cell) GetContent() string {
	return r.content
}

// SetMinWidth sets the cells minimum width, this will not disable responsivness.
// This has only an effect to cells of a normal FlexBox, not a HorizontalFlexBox.
func (r *Cell) SetMinWidth(value int) *Cell {
	r.minWidth = value
	return r
}

// SetMinHeigth sets the cells minimum height, this will not disable responsivness.
// This has only an effect to cells of a HorizontalFlexBox.
func (r *Cell) SetMinHeigth(value int) *Cell {
	r.minHeigth = value
	return r
}

// SetStyle replaces the style, it unsets width/height related keys
func (r *Cell) SetStyle(style lipgloss.Style) *Cell {
	r.style = style.
		UnsetWidth().
		UnsetMaxWidth().
		UnsetHeight().
		UnsetMaxHeight()
	return r
}

// GetStyle returns the copy of the cells current style
func (r *Cell) GetStyle() lipgloss.Style {
	return r.style.Copy()
}

// GetWidth returns real width of the cell
func (r *Cell) GetWidth() int {
	return r.getMaxWidth()
}

// GetHeight returns real height of the cell
func (r *Cell) GetHeight() int {
	return r.getMaxHeight()
}

// render the cell into string
func (r *Cell) render(inherited ...lipgloss.Style) string {
	for _, style := range inherited {
		r.style = r.style.Inherit(style)
	}

	s := r.GetStyle().
		Width(r.getContentWidth()).MaxWidth(r.getMaxWidth()).
		Height(r.getContentHeight()).MaxHeight(r.getMaxHeight())
	return s.Render(r.content)
}

func (r *Cell) copy() Cell {
	cellCopy := *r
	cellCopy.style = r.GetStyle()
	return cellCopy
}
