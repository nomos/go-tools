package pix_icon_assets

import (
	_ "embed"
)

func Assign(m map[string][]byte){
	addPixIcon(m,"folder",folder)
	addPixIcon(m,"folder_open",folder_open)
	addPixIcon(m,"clock",clock)
	addPixIcon(m,"start",start)
}

func addPixIcon(m map[string][]byte,s string,data []byte){
	m["pix_"+s] = data
}

//go:embed icons/folder.png
var folder []byte

//go:embed icons/folder_open.png
var folder_open []byte

//go:embed icons/clock.png
var clock []byte

//go:embed icons/start.png
var start []byte
