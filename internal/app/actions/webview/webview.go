package webview

import (
	wview "github.com/webview/webview"
)

func Test() {
	debug := true
	w := wview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal wview example")
	w.SetSize(800, 600, wview.HintNone)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.Run()
}
