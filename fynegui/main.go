// package main

// import (
// 	// "gofar/fynegui/clock"
// 	"gofar/fynegui/mergefile"
// )

// func main() {
// 	// clock.NewClock()
// 	mergefile.NewMerge()
// }

package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	a.Settings().SetTheme(theme.DarkTheme())

	obj := canvas.NewCircle(color.Black)
	obj.Resize(fyne.NewSize(100, 100))
	obj.Move(fyne.NewPos(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))

	red := color.NRGBA{R: 0xff, A: 0xff}
	green := color.NRGBA{G: 0xff, A: 0xff}
	blue := color.NRGBA{B: 0xff, A: 0xff}

	go func() {
		for i := 0; i < 7; {
			fmt.Printf("i = %d\n", i)
			if i == 6 {
				fmt.Printf("now i = 6\n")
				i -= 6
			}
			if i == 0 {
				r_b := canvas.NewColorRGBAAnimation(red, blue, 5*time.Second, func(c color.Color) {
					obj.FillColor = c
					canvas.Refresh(obj)
				})
				i += 2
				r_b.Start()
				fmt.Printf("r->b  i = %d\n", i)
				time.Sleep(5 * time.Second)
			}

			if i == 2 {
				b_r := canvas.NewColorRGBAAnimation(blue, green, 5*time.Second, func(c color.Color) {
					obj.FillColor = c
					canvas.Refresh(obj)
				})
				i += 2
				b_r.Start()
				fmt.Printf("b->g  i = %d\n", i)
				time.Sleep(5 * time.Second)
			}

			if i == 4 {
				b_r := canvas.NewColorRGBAAnimation(green, red, 5*time.Second, func(c color.Color) {
					obj.FillColor = c
					canvas.Refresh(obj)
				})
				i += 2
				b_r.Start()
				fmt.Printf("g->b  i = %d\n", i)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	w.Resize(fyne.NewSize(200, 200))
	w.SetPadded(false)
	w.ShowAndRun()

}
