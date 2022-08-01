package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var mode *walk.ComboBox
	var check *walk.CheckBox

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Size:    Size{600, 400},
		Layout:  HBox{},
		Children: []Widget{
			Label{
				Text: "Input",
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					CheckBox{
						AssignTo: &check,
						Text:     "Unspecified",
					},
					CheckBox{
						Text: "Not",
					},
				},
			},
			ComboBox{
				AssignTo: &mode,
				Model:    []string{"Unspecified", "Not"},
			},
		},
	}.Run()
}
