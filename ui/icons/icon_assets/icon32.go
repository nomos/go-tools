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
	addIcon(m,"boolean_icon",boolean_icon)
	addIcon(m,"save",save)
	addIcon(m,"sort_up",sort_up)
	addIcon(m,"sort_down",sort_down)
	addIcon(m,"cancel",cancel)
	addIcon(m,"green_circle",green_circle)
	addIcon(m,"red_circle",red_circle)
	addIcon(m,"search_client",search_client)
	addIcon(m,"search",search)
	addIcon(m,"user",user)
	addIcon(m,"microsoft_excel",microsoft_excel)
	addIcon(m,"toggle_on",toggle_on)
	addIcon(m,"toggle_off",toggle_off)
}

func addIcon(m map[string][]byte,s string,data []byte){
	m[s] = data
}
//go:embed icons/icons8-toggle_off.png
var toggle_off []byte

//go:embed icons/icons8-toggle_on.png
var toggle_on []byte

//go:embed icons/icons8-microsoft_excel.png
var microsoft_excel []byte

//go:embed icons/icons8-user.png
var user []byte

//go:embed icons/icons8-search.png
var search []byte

//go:embed icons/icons8-search_client.png
var search_client []byte

//go:embed icons/icons8-green_circle.png
var green_circle []byte

//go:embed icons/icons8-red_circle.png
var red_circle []byte

//go:embed icons/icons8-cancel.png
var cancel []byte

//go:embed icons/icons8-sort_up.png
var sort_up []byte

//go:embed icons/icons8-sort_down.png
var sort_down []byte

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

//go:embed icons/icons8-boolean.png
var boolean_icon []byte

//go:embed icons/icons8-number.png
var number_icon []byte

//go:embed icons/icons8-string.png
var string_icon []byte

//go:embed icons/icons8-save.png
var save []byte
