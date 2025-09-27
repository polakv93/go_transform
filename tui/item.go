package tui

type item struct {
	fileName string
	path     string
}

func (i item) FilterValue() string {
	return i.fileName
}
