package clock

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewClock() {
	myApp := app.New()
	w := myApp.NewWindow("Clock")

	contentSet(w)

	w.Resize(fyne.Size{600, 300})
	w.ShowAndRun()
}

func contentSet(w fyne.Window) {
	str := binding.NewString()
	starttime := time.Now()
	starttimeTxt := starttime.Format("2006-01-02 15:04:05")
	str.Set(starttimeTxt)
	passed := binding.NewFloat()
	passed.Set(0.00)
	go func(starttime time.Time, str binding.String, passed binding.Float) {
		tiktok := time.NewTicker(time.Second)
		for range tiktok.C {
			passedF, _ := passed.Get()
			// fmt.Printf("start: %v   now: %v \n", starttime, time.Now())
			if time.Now().After(starttime.Add(5 * time.Second)) {
				str.Set("over")
				passed.Set(passedF + 0.20)
				return
			} else {
				str.Set(time.Now().Format("2006-01-02 15:04:05"))
				passed.Set(passedF + 0.20)
			}
		}
	}(starttime, str, passed)

	// 开始时间
	StaClock := widget.NewLabel(starttimeTxt)
	StaClock.Resize(fyne.Size{300, 50})
	StaClock.Move(fyne.NewPos(0, 0))

	// 现在时间
	curClock := widget.NewLabelWithData(str)
	curClock.Resize(fyne.Size{300, 50})
	curClock.Move(fyne.NewPos(201, 0))

	// 进度（已经过时间）
	passDura := widget.NewProgressBarWithData(passed)
	passDura.Resize(fyne.Size{600, 20})
	passDura.Move(fyne.NewPos(0, 150))

	w.SetContent(container.NewWithoutLayout(StaClock, curClock, passDura))
}
