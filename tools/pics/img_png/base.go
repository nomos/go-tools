package img_png

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func ifelse(c bool, a, b int) int {
	if c {
		return a
	} else {
		return b
	}
}

type Frame struct {
	Rect         image.Rectangle
	Offset       image.Point
	OriginalSize image.Point
	Rotated      int
}

func LoadImage(path string) (img image.Image, err error) {
	return imaging.Open(path)
}

func SaveImage(path string, img image.Image) (err error) {

	if filepath.Ext(path) == "" {
		path = path + ".png"
	}

	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	imgfile, err := os.Create(path)
	defer imgfile.Close()
	return png.Encode(imgfile, img)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func GetFiles(dir string, allow []string) []string {

	allowMap := map[string]bool{}
	if allow != nil {
		for _, v := range allow {
			allowMap[v] = true
		}
	}

	ret := []string{}
	filepath.Walk(dir, func(fpath string, f os.FileInfo, err error) error {
		if f == nil || f.IsDir() {
			return nil
		}

		ext := path.Ext(fpath)
		if allowMap[ext] {
			ret = append(ret, filepath.ToSlash(fpath))
		}

		return nil
	})

	return ret
}

type AtlasPart struct {
	ImageFile string
	Frames    map[string]*Frame
}

type DumpContext struct {
	FileName    string
	FileContent []byte
	Atlases     []*AtlasPart
}

func (dc *DumpContext) AppendPart() *AtlasPart {
	part := &AtlasPart{}
	part.Frames = map[string]*Frame{}
	dc.Atlases = append(dc.Atlases, part)
	return part
}

func (dc *DumpContext) dumpFrames(part *AtlasPart) error {
	textureFileName := filepath.Join(filepath.Dir(dc.FileName), part.ImageFile)

	textureImage, err := LoadImage(textureFileName)
	if err != nil {
		return fmt.Errorf("open image error:" + textureFileName)
	}

	outdir := filepath.Join(dc.FileName + ".out")
	if !IsDir(outdir) {
		err = os.MkdirAll(outdir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	for k, v := range part.Frames {
		fmt.Println(k)

		var subImage image.Image

		w, h := v.Rect.Size().X, v.Rect.Size().Y
		ox, oy := v.Offset.X, v.Offset.Y
		ow, oh := v.OriginalSize.X, v.OriginalSize.Y
		x, y := v.Rect.Min.X, v.Rect.Min.Y

		if v.Rotated == 90 {
			subImage = imaging.Crop(textureImage, image.Rect(x, y, x+h, y+w))
			subImage = imaging.Rotate90(subImage)
		} else if v.Rotated == 270 {
			subImage = imaging.Crop(textureImage, image.Rect(x, y, x+h, y+w))
			subImage = imaging.Rotate270(subImage)
		} else {
			subImage = imaging.Crop(textureImage, image.Rect(x, y, x+w, y+h))
		}

		destImage := image.NewRGBA(image.Rect(0, 0, ow, oh))
		newImage := imaging.Paste(destImage, subImage, image.Point{(ow-w)/2 + ox, (oh-h)/2 - oy})

		savepath := path.Join(outdir, k)
		if path.Ext(savepath) == "" {
			savepath += ".png"
		}

		SaveImage(savepath, newImage)
	}

	return nil
}

func dumpByFileName(filename string) {

	c := DumpContext{
		FileName: filename,
		Atlases:  []*AtlasPart{},
	}

	data, _ := ioutil.ReadFile(c.FileName)
	c.FileContent = data

	var err error

	ext := path.Ext(filename)
	switch ext {
	case ".plist":
		err = dumpPlist(&c)
	case ".json":
		err = dumpJson(&c)
	case ".fnt":
		err = dumpFnt(&c)
	case ".atlas":
		err = dumpSpine(&c)
	default:
		return
	}

	if err != nil {
		panic(err)
	}

	for _, part := range c.Atlases {
		err = c.dumpFrames(part)
		if err != nil {
			panic(err)
		}
	}
}

//func main() {
//	var args struct {
//		Input string `arg:"positional"`
//		Ext   string `arg:"-e,--ext" help:"dump ext json,plist,fnt,atlas "`
//	}
//
//	arg.MustParse(&args)
//
//	var ext []string
//
//	if args.Ext == "" {
//		ext = []string{".json", ".plist", ".fnt", ".atlas"}
//	} else {
//		ext = []string{}
//		arr := strings.Split(args.Ext, ",")
//		for _, v := range arr {
//			ext = append(ext, "."+v)
//		}
//	}
//
//	if args.Input == "" {
//		args.Input = "./"
//	}
//
//	fmt.Println(ext)
//
//	allfiles := []string{}
//	if IsDir(args.Input) {
//		files := GetFiles(args.Input, ext)
//		allfiles = append(allfiles, files...)
//	} else {
//		allfiles = append(allfiles, args.Input)
//	}
//
//	fmt.Println(fmt.Sprintf("开始导出：共（%d）个", len(allfiles)))
//	for i, v := range allfiles {
//		fmt.Println(fmt.Sprintf("导出 %d/%d %s", i+1, len(allfiles), v))
//		dumpByFileName(v)
//	}
//
//	fmt.Printf("\n")
//	fmt.Printf("好用请给个Star，谢谢.\n")
//	fmt.Printf("https://github.com/qcdong2016/PlistDumper.git\n")
//}
