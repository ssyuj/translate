package diytheme

import (
	"fmt"

	"fyne.io/fyne"
)

func GetIcon() fyne.Resource {
	icon, err := fyne.LoadResourceFromPath("./icon/favico.png")
	if err != nil {
		fmt.Println(err)
	}
	if icon == nil {
		icon, err := fyne.LoadResourceFromPath("./resource/image/favico.png")
		if err != nil {
			fmt.Println(err)
		}
		return icon
	}
	return icon
}
