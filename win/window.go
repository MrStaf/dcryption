package win

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MrStaf/dcryption/util"
)

func DisplayFiles() {

}

func Init() fyne.Window {
	// Initialize the window
	a := app.New()
	w := a.NewWindow("dCryption")
	w.Resize(fyne.NewSize(1280, 720))
	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	path := widget.NewEntry()
	key := widget.NewEntry()
	output := widget.NewEntry()

	form1 := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Path", Widget: path},
			{Text: "Key", Widget: key},
			{Text: "Output", Widget: output},
		},
		OnSubmit: func() {
			s := util.LoadSettingsFile()
			s.SetEncryptedPath(path.Text)
			s.SetKey(key.Text)
			s.SetDecryptedPath(output.Text)
			util.SaveSettingsFile(s)
		},
		SubmitText: "Save",
	}
	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, form1)
	w.SetContent(content)
	return w
}

func Run(app *fyne.Window) {
	// Run the window
	app2 := *app
	app2.ShowAndRun()
}
