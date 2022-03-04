package img_png

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/gzip"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"strconv"
)

//func main() {
//	json := Png2Json("/Users/wqs/Project2/art/D47hZt3WkAApK2U.png", "/Users/wqs/Project2/export.json")
//	log.Warn(json)
//	time.Sleep(time.Second)
//}

type ImgJsonFormat [][]uint32

func t2x(t uint32) string {
	result := strconv.FormatInt(int64(t), 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func rgb2hex(color []uint32) string {
	r := t2x(color[0])
	g := t2x(color[1])
	b := t2x(color[2])
	a := t2x(color[3])
	return r + g + b + a
}

func Png2ByteArray(path string) (*ImageRGBAMap, error) {
	img, err := ReadImageFile(path)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	ret := DecodeRGBA(img)
	return ret, nil
}

func Png2Base64(path string) (int, int, string, error) {
	img, err := ReadImageFile(path)
	if err != nil {
		log.Error(err.Error())
		return 0, 0, "", err
	}
	width, height, arr := DecodeImgData(img)
	ret := base64.StdEncoding.EncodeToString(arr)
	return width, height, ret, nil
}

func Png2CompressedBase64(path string) (int, int, string, error) {
	img, err := ReadImageFile(path)
	if err != nil {
		log.Error(err.Error())
		return 0, 0, "", err
	}
	width, height, arr := DecodeImgData(img)
	if width > 255 || height > 255 {
		return 0, 0, "", errors.New("width or height must <=255")
	}
	arr2 := make([]byte, len(arr)+2)
	arr2[0] = uint8(width)
	arr2[1] = uint8(height)
	arr3 := arr2[2:]
	copy(arr3, arr)
	ret, _ := gzip.CompressBytes2Base64(arr2)
	return width, height, ret, nil
}

func Png2ImageMap(path string) (int, int, []byte, error) {
	img, err := ReadImageFile(path)
	if err != nil {
		log.Error(err.Error())
		return 0, 0, nil, err
	}
	width, height, arr := DecodeImgData(img)
	return width, height, arr, nil
}

func ReadPngFile(path string) (img image.Image, err error) {
	fmt.Println("读取图片中....")
	fileByte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件失败....")
		return img, err
	}
	img, err = png.Decode(bytes.NewBuffer(fileByte))
	if err != nil {
		log.Error(err.Error())
		fmt.Println("图片解码失败....")
		return img, err
	}
	fmt.Println("读取图片完成....")
	return img, err
}

func ReadImageFile(path string) (img image.Image, err error) {
	fmt.Println("读取图片中....")
	fileByte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件失败....")
		return img, err
	}
	img, _, err = image.Decode(bytes.NewBuffer(fileByte))
	if err != nil {
		log.Error(err.Error())
		fmt.Println("图片解码失败....")
		return img, err
	}
	fmt.Println("读取图片完成....")
	return img, err
}

type ColorPoint struct {
	X int
	Y int
	R uint32
	G uint32
	B uint32
	A uint32
}

type ImageRGBAMap struct {
	Width  int
	Height int
	Data   []ColorPoint
}

func (this *ImageRGBAMap) GetData() []byte {
	ret := make([]byte, this.Width*this.Height*4)
	for y := 0; y < this.Height; y++ {
		for x := 0; x < this.Width; x++ {
			ret[x*4+y*this.Width*4+0] = byte(this.At(x, y).R)
			ret[x*4+y*this.Width*4+1] = byte(this.At(x, y).G)
			ret[x*4+y*this.Width*4+2] = byte(this.At(x, y).B)
			ret[x*4+y*this.Width*4+3] = byte(this.At(x, y).A)
		}
	}
	return ret
}

func (this *ImageRGBAMap) At(x int, y int) ColorPoint {
	return this.Data[x+y*this.Width]
}

func DecodeImgData(img image.Image) (int, int, []byte) {
	log.Info("读取图片数据中....")
	rectangle := img.Bounds()
	width := rectangle.Max.X - rectangle.Min.X
	height := rectangle.Max.Y - rectangle.Min.Y
	arr := make([]byte, height*width*4)
	for yindex := rectangle.Min.Y; yindex < rectangle.Max.Y; yindex++ {
		for xindex := rectangle.Min.X; xindex < rectangle.Max.X; xindex++ {
			x := xindex - rectangle.Min.X
			y := yindex - rectangle.Min.Y
			r, g, b, a := img.At(xindex, yindex).RGBA()
			arr[y*width*4+x*4] = byte(r >> 8)
			arr[y*width*4+x*4+1] = byte(g >> 8)
			arr[y*width*4+x*4+2] = byte(b >> 8)
			arr[y*width*4+x*4+3] = byte(a >> 8)
		}
	}
	log.Info("读取图片数据完成")
	return width, height, arr
}

func DecodeRGBA(img image.Image) *ImageRGBAMap {
	log.Info("读取图片RGBA中....")
	imageMap := &ImageRGBAMap{}
	rectangle := img.Bounds()
	imageMap.Width = rectangle.Max.X
	imageMap.Height = rectangle.Max.Y
	for yindex := rectangle.Min.Y; yindex < rectangle.Max.Y; yindex++ {
		for xindex := rectangle.Min.X; xindex < rectangle.Max.X; xindex++ {
			r, g, b, a := img.At(xindex, yindex).RGBA()
			imageMap.Data = append(imageMap.Data, ColorPoint{
				X: xindex,
				Y: yindex,
				R: r >> 8,
				G: g >> 8,
				B: b >> 8,
				A: a >> 8,
			})
		}
	}
	log.Info("读取图片RGBA完成")
	return imageMap
}
