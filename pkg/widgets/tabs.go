package widgets

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

var (
	styleTabActive     = tcell.StyleDefault.Background(tcell.ColorSilver).Foreground(tcell.ColorWhite).Bold(true)
	styleTabInactive   = tcell.StyleDefault.Background(tcell.ColorGray).Foreground(tcell.ColorWhite)
	styleTabBackground = tcell.StyleDefault.Background(tcell.ColorWhite)
)

type tabsItem struct {
	name string
	text *views.Text
}

// Tabs is a View with multiple items in single line.
type Tabs struct {
	items    []tabsItem
	selected int

	views.BoxLayout
	views.WidgetWatchers
}

// NewTabs returns a new Tabs.
func NewTabs() *Tabs {
	w := &Tabs{
		selected: -1,
	}
	w.SetStyle(styleTabBackground)
	w.SetOrientation(views.Horizontal)
	return w
}

// AddTab adds a new item with the name.
func (w *Tabs) AddTab(name string) {
	text := &views.Text{}
	text.SetText(" " + name + " ")
	text.SetStyle(styleTabInactive)

	w.AddWidget(text, 0)
	w.items = append(w.items, tabsItem{
		name: name,
		text: text,
	})
}

// TabCount returns the count of the tabs.
func (w *Tabs) TabCount() int {
	return len(w.items)
}

// SelectNext selects next tab of the current
func (w *Tabs) SelectNext() {
	index := w.selected + 1
	if index >= len(w.items) {
		index = 0
	}
	w.SelectAt(index)
}

// SelectPrev selects previous tab of the current
func (w *Tabs) SelectPrev() {
	index := w.selected + 1
	if index >= len(w.items) {
		index = 0
	}
	w.SelectAt(index)
}

// Clear clears current tabs.
func (w *Tabs) Clear() {
	for _, item := range w.items {
		w.RemoveWidget(item.text)
	}
	w.items = nil
	w.selected = -1
}

// SelectAt selects the tab of the index in items.
func (w *Tabs) SelectAt(index int) {
	if index == w.selected {
		return
	}
	if w.selected >= 0 {
		item := w.items[w.selected]
		item.text.SetStyle(styleTabInactive)
	}
	if index < 0 || index >= len(w.items) {
		return
	}
	w.selected = index
	item := w.items[w.selected]
	item.text.SetStyle(styleTabActive)

	w.PostEventWidgetContent(w)

	ev := &EventItemSelected{
		Name:   item.name,
		Index:  index,
		widget: w,
	}
	ev.SetEventNow()

	w.PostEvent(ev)
}
