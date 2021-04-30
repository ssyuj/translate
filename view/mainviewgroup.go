package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func FileSelectView(callback func(fyne.URIReadCloser, error), view fyne.Window) *dialog.FileDialog {
	fileselectdialog := dialog.NewFileOpen(callback, view)
	fileselectdialog.Resize(fyne.Size{Width: 700, Height: 550})
	loadfileuri, err := storage.ParseURI("file://e:/")
	if err != nil {
		fmt.Println(err)
	}
	fileuris, err := storage.ListerForURI(loadfileuri)
	if err != nil {
		fmt.Println(err)
	}
	fileselectdialog.SetLocation(fileuris)
	return fileselectdialog
}
