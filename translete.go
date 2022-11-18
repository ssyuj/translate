package main

import (
	"fmt"
	"os"
	"strings"
	"translate/collter"
	"translate/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

var dictpath = "./resource/dict/dict.txt"

//go build -ldflags="-H windowsgui" 隐藏默认的cmd窗口
// var key []byte = []byte(string([]int32{64, 71, 97, 119, 94, 50, 116, 71, 81, 54, 49, 45, 206, 210, 110, 105}))

func init() {
	fpaths := findfont.List()
	for _, path := range fpaths {
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	 
	// 加载字典表
	collter.Fileload(dictpath)
}

func main() {
	var path string
	// 创建一个应用
	transapp := app.New()
	// 设置主题
	transapp.Settings().SetTheme(theme.LightTheme())
	// 创建一个窗体，标题为"翻译"
	transform := transapp.NewWindow("翻译")
	transform.SetFixedSize(true)
	// 创建歌词打印
	sourcetxt := widget.NewLabel("")
	// sourcetxt := widget.NewTextGrid()
	// sourcetxt := canvas.NewText("", color.Black)
	sourcetxt.Resize(fyne.NewSize(900, 600))
	sourcetxt.Alignment = fyne.TextAlignCenter
	sourcetxt.Hidden = true
	// 创建滚动条容器
	scrolltxt := container.NewVScroll(sourcetxt)
	scrolltxt.SetMinSize(fyne.NewSize(900, 600))
	scrolltxt.Resize(fyne.NewSize(900, 1200))

	// 打开文件
	selectfile := widget.NewButton("打开文件", func() {
		view.FileSelectView(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println(err)
			}
			if uc != nil && uc.URI() != nil {
				path = uc.URI().Path()
				content, err := collter.GetFileContent(path)
				if err != nil {
					sourcetxt.Text = err.Error()
					// sourcetxt.SetText(err.Error())
				}
				sourcetxt.Text = content
				// sourcetxt.SetText(content)
				sourcetxt.Hidden = false
				uc.Close()
			}
		}, transform).Show()
	})
	// 保存
	savefile := widget.NewButton("保存", func() {
		if strings.TrimSpace(path) == "" || len(path) <= 0 {
			infowin := dialog.NewInformation("提示", "请先打开文件", transform)
			infowin.Show()
		} else {
			content := sourcetxt.Text
			saveboo := collter.Savephonetic(path, content)
			if saveboo {
				infowin := dialog.NewInformation("提示", "保存成功，请在原文件路径下查看*_tmp.txt的文件", transform)
				infowin.Show()
			} else {
				infowin := dialog.NewInformation("提示", "文件保存失败", transform)
				infowin.Show()
			}
		}
	})
	// 加载音标
	phoneitcbtn := widget.NewButton("音标", func() {
		if strings.TrimSpace(path) == "" || len(path) <= 0 {
			infowin := dialog.NewInformation("提示", "请选择文件路径", transform)
			infowin.Show()
		} else {
			content := collter.Loadphonetic(path)
			sourcetxt.Refresh()
			sourcetxt.Text = content
		}
	})
	// 设置头部按钮
	// toolbtn := container.NewCenter(selectfile, phoneitcbtn, savefile)
	// container.NewHSplit()
	toolbtn := container.NewGridWithRows(1, selectfile, phoneitcbtn, savefile)

	transform.SetContent(container.NewVBox(toolbtn, scrolltxt))

	// // 事件注册器，键盘点击。预功能 `复制/粘贴`
	// if deskCanvas, ok := transform.Canvas().(desktop.Canvas); ok {
	// 	deskCanvas.SetOnKeyDown(func(ev *fyne.KeyEvent) {
	// 		fmt.Println("KeyDown: " + string(ev.Name))
	// 	})
	// 	deskCanvas.SetOnKeyUp(func(ev *fyne.KeyEvent) {
	// 		fmt.Println("- KeyDown: " + string(ev.Name))
	// 	})
	// }

	// transform.Canvas().SetOnTypedRune(func(r rune) {
	// 	fmt.Println("r:" + string(r))
	// })
	// transform.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
	// 	fmt.Println("key:" + string(ke.Name))
	// })

	// transform.SetIcon(diytheme.GetIcon())718093
	// txt := widget.NewLabel("这是测试")
	// txt.Hidden = true
	// transform.Canvas().SetContent(txt)
	transform.Resize(fyne.NewSize(900, 600))
	transform.ShowAndRun()

	os.Unsetenv("FYNE_FONT")
}

// func main() {
// 	// s := "  ad，sddd这是测试粤语音标時,,"
// 	// fmt.Println(collter.MarkWord(s))
// 	// fmt.Println(s)
// 	path := "./resource/tmp/对玄机.krc"
// 	// spath := "./resource/tmp/tmp.krcs"
// 	buff, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	sbuff := make([]byte, len(buff)-4)
// 	copy(sbuff, buff[4:])
// 	for i := 0; i < len(sbuff); i++ {
// 		sbuff[i] = (sbuff[i] ^ key[i%len(key)])
// 	}
// 	fmt.Println(string(buff))

// 	flateReader := flate.NewReader(bytes.NewBuffer(sbuff[2:]))
// 	defer flateReader.Close()

// 	io.Copy(os.Stdout, flateReader)

// 	// b := bytes.NewReader(sbuff)
// 	// r := bzip2.NewReader(b)
// 	// buf := make([]byte, 1024)
// 	// for {
// 	// 	n, err := r.Read(buf)
// 	// 	if n == 0 {
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		fmt.Println("bzip2:", err)
// 	// 	}
// 	// 	fmt.Println(buf)
// 	// }

// 	// rbuff, err := yamsutils.UnZip(sbuff)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(rbuff)
// }
