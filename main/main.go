package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.SetIcon(theme.SettingsIcon())
	w := a.NewWindow("ICAR Toolbox")
	servoData1 := binding.NewString()
	servoData2 := binding.NewString()
	servoData3 := binding.NewString()
	servoData4 := binding.NewString()
	servoData5 := binding.NewString()
	servoData6 := binding.NewString()
	
	//Home页：放一些连接设置及基本信息
	servoSetForm := widget.NewForm(
		widget.NewFormItem("servo1 Angle: ", widget.NewEntryWithData(servoData1)),
		widget.NewFormItem("servo2 Angle: ", widget.NewEntryWithData(servoData2)),
		widget.NewFormItem("servo3 Angle: ", widget.NewEntryWithData(servoData3)),
		widget.NewFormItem("servo4 Angle: ", widget.NewEntryWithData(servoData4)),
		widget.NewFormItem("servo5 Angle: ", widget.NewEntryWithData(servoData5)),
		widget.NewFormItem("servo6 Angle: ", widget.NewEntryWithData(servoData6)),
	)
	servoSetForm.SubmitText = "SET"
	servoSetForm.OnSubmit = func() {

	}
	servoSetForm.CancelText = "RESET"
	servoSetForm.OnCancel = func() {
		servoData1.Set("")
		servoData2.Set("")
		servoData3.Set("")
		servoData4.Set("")
		servoData5.Set("")
		servoData6.Set("")
	}

	servoTab := container.NewVBox(
		widget.NewLabel("Servo Angle Set:"),
		servoSetForm,
	)

	//Tab选项卡
	tabs := container.NewAppTabs(
		container.NewTabItem("Home", widget.NewLabel("Home Page")),
		container.NewTabItem("Arm", servoTab),
		container.NewTabItem("Chassis", widget.NewEntry()),
		container.NewTabItem("About", widget.NewEntry()),
	)
	//设置Tabs的位置
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(700, 400))
	w.ShowAndRun()

}
