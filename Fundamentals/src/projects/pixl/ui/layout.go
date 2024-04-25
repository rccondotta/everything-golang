package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorpicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorpicker, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}
