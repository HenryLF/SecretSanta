package main

import (
	"log"
	"os"
	"path/filepath"
	"secretsanta/drafter"

	webview "github.com/webview/webview_go"
)

func main() {
	var url string
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		url = "./static/index.html"
	}
	W := webview.New(true)
	defer W.Destroy()
	W.SetTitle("")
	W.SetSize(1080, 1080, webview.HintNone)
	SetBinding(W)
	url, err := filepath.Abs(url)
	if err != nil {
		log.Print("No Asset in ", url)
		os.Exit(2)
	}
	url = filepath.Join("file://", url)
	log.Println(url)
	W.Navigate(url)

	W.Run()
}

func SetBinding(W webview.WebView) {
	W.Bind("draft", drafter.Draft)
}
