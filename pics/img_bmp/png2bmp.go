package img_bmp

import (
	"bytes"
	"github.com/nomos/go-lokas/log"
	"golang.org/x/image/bmp"
	"image/png"
)

func Png2Bmp (data []byte)([]byte,error){
	img,err:=png.Decode(bytes.NewBuffer(data))
	if err != nil {
		log.Error(err.Error())
		return nil,err
	}
	dst :=bytes.NewBuffer(make([]byte,0,1024*1024*80))
	err=bmp.Encode(dst,img)
	if err != nil {
		log.Error(err.Error())
		return nil,err
	}
	ret:= dst.Bytes()
	return ret,nil
}
