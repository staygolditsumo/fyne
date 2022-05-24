package mergefile

import (
	"fmt"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewMerge() {
	myApp := app.New()
	w := myApp.NewWindow("MergeFile")

	contentSet(w)

	w.Resize(fyne.Size{900, 600})
	w.ShowAndRun()

}

func contentSet(w fyne.Window) {

	label1 := widget.NewLabel("File 1 : ")
	label1.Resize(fyne.Size{100, 40})
	label1.Move(fyne.NewPos(0, 30))

	url1 := binding.NewString()
	selectedFile1 := widget.NewEntryWithData(url1)
	selectedFile1.PlaceHolder = "Please select a file..."
	selectedFile1.TextStyle = fyne.TextStyle{Bold: true}
	selectedFile1.Disable()
	selectedFile1.Resize(fyne.Size{600, 40})
	selectedFile1.Move(fyne.NewPos(100, 30))

	fileSelBtn1 := widget.NewButton("Select", func() {
		fileOpen1 := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println("FileOpen1 err :" + err.Error())
			} else {
				if uc.URI() != nil {
					url1.Set(uc.URI().Path())
				}
			}
		}, w)
		fileOpen1.Resize(fyne.Size{800, 500})
		fileOpen1.Show()
	})
	fileSelBtn1.Resize(fyne.Size{80, 40})
	fileSelBtn1.Move(fyne.NewPos(720, 30))

	label2 := widget.NewLabel("File 2 : ")
	label2.Resize(fyne.Size{100, 40})
	label2.Move(fyne.NewPos(0, 120))

	url2 := binding.NewString()
	selectedFile2 := widget.NewEntryWithData(url2)
	selectedFile2.PlaceHolder = "Please select a file..."
	selectedFile2.TextStyle = fyne.TextStyle{Bold: true}
	selectedFile2.Disable()
	selectedFile2.Resize(fyne.Size{600, 40})
	selectedFile2.Move(fyne.NewPos(100, 120))

	fileSelBtn2 := widget.NewButton("Select", func() {
		fileOpen2 := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println("FileOpen2 err :" + err.Error())
			} else {
				if uc.URI() != nil {
					url2.Set(uc.URI().Path())
				}
			}
		}, w)
		fileOpen2.Resize(fyne.Size{800, 500})
		fileOpen2.Show()
	})
	fileSelBtn2.Resize(fyne.Size{80, 40})
	fileSelBtn2.Move(fyne.NewPos(720, 120))

	label3 := widget.NewLabel("Export Path: ")
	label3.Resize(fyne.Size{80, 40})
	label3.Move(fyne.NewPos(0, 210))

	dir := binding.NewString()
	exportPath := widget.NewEntryWithData(dir)
	exportPath.PlaceHolder = "Please select a directory..."
	exportPath.TextStyle = fyne.TextStyle{Bold: true}
	exportPath.Disable()
	exportPath.Resize(fyne.Size{600, 40})
	exportPath.Move(fyne.NewPos(100, 210))

	dirSelBtn := widget.NewButton("Select", func() {
		dirOpen := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
			if err != nil {
				fmt.Println("dirOpen err :" + err.Error())
			} else {
				if lu != nil {
					dir.Set(lu.Path())
				}
			}
		}, w)
		dirOpen.Resize(fyne.Size{800, 500})
		dirOpen.Show()
	})
	dirSelBtn.Resize(fyne.Size{80, 40})
	dirSelBtn.Move(fyne.NewPos(720, 210))

	mergeBtn := widget.NewButton("Merge", func() {
		url1Str, _ := url1.Get()
		url2Str, _ := url2.Get()
		dirStr, _ := dir.Get()
		if url1Str == "" {
			dialog.NewConfirm("Confirm", "Please select File 1", func(b bool) {
				w.Canvas().Focus(selectedFile1)
			}, w).Show()
			return
		} else if url2Str == "" {
			dialog.NewConfirm("Confirm", "Please select File 2", func(b bool) {
				w.Canvas().Focus(selectedFile2)
			}, w).Show()
			return
		} else if dirStr == "" {
			dialog.NewConfirm("Confirm", "Please select export path!", func(b bool) {
				w.Canvas().Focus(exportPath)
			}, w).Show()
			return
		}

		// fmt.Printf("path1 is %s , path2 is %s\n", path1, path2)
		err := merge(url1Str, url2Str, dirStr)
		if err == nil {
			dialog.NewInformation("Information", "Merge succeed!", w).Show()
		}
	})
	mergeBtn.Resize(fyne.Size{100, 50})
	mergeBtn.Move(fyne.NewPos(150, 300))

	exitBtn := widget.NewButton("Exit", func() {
		w.Close()
	})
	exitBtn.Resize(fyne.Size{100, 50})
	exitBtn.Move(fyne.NewPos(550, 300))

	w.SetContent(container.NewWithoutLayout(label1, selectedFile1, fileSelBtn1, label2, selectedFile2, fileSelBtn2, label3, exportPath, dirSelBtn, mergeBtn, exitBtn))
}

func merge(url1Str string, url2Str string, dirStr string) error {
	// pathChan := make(chan string)
	// for _, path := range args {
	// 	pathChan <- path
	// }
	// path1 := <-pathChan
	// path2 := <-pathChan
	fmt.Printf("url1Str : %s , url2Str : %s , dirStr : %s", url1Str, url2Str, dirStr)

	data1, err := ioutil.ReadFile(url1Str)
	if err != nil {
		fmt.Println("File 1 reading error: ", err)
		return err
	}
	data2, err := ioutil.ReadFile(url2Str)
	if err != nil {
		fmt.Println("File 1 reading error: ", err)
		return err
	}
	// fmt.Println(data)         // 读取到文件内容的字符串ASC码构成的切片[97 115 100 115 97 100 97 115 100 97 115 100]
	// fmt.Println(string(data1) + "\n" + string(data2)) // 将字符串切片转成字符串

	dataAll := append(data1, data2...)
	errW := ioutil.WriteFile(dirStr+"/mergefile.txt", dataAll, 0644)
	if errW != nil {
		log.Fatal(errW)
	}
	return errW
}
