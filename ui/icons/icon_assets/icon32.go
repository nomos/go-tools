package icon_assets

import (
	_ "embed"
)

func Assign(m map[string][]byte){
	addIcon(m,"folder",folder)
	addIcon(m,"help",help)
	addIcon(m,"object_box",object_box)
	addIcon(m,"array_box",array_box)
	addIcon(m,"null_icon",null_icon)
	addIcon(m,"number_icon",number_icon)
	addIcon(m,"string_icon",string_icon)
}

func addIcon(m map[string][]byte,s string,data []byte){
	m[s] = data
}

//go:embed icons/icons8-folder.png
var folder []byte

//go:embed icons/icons8-help.png
var help []byte

//go:embed icons/icons8-object_box.png
var object_box []byte

//go:embed icons/icons8-array_box.png
var array_box []byte

//go:embed icons/icons8-null.png
var null_icon []byte

//go:embed icons/icons8-number.png
var number_icon []byte

//go:embed icons/icons8-string.png
var string_icon []byte
