package img_png

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/ecs"
	"github.com/nomos/go-lokas/protocol"
	"reflect"
	"time"
)

var _ lokas.IComponent = (*OK)(nil)

type OK struct {
	ecs.Component `json:"-" bson:"-"`

}

func (this *OK) OnAdd(e lokas.IEntity, r lokas.IRuntime) {

}

func (this *OK) OnRemove(e lokas.IEntity, r lokas.IRuntime) {

}

func (this *OK) OnCreate(r lokas.IRuntime) {

}

func (this *OK) OnDestroy(r lokas.IRuntime) {

}

func (this *OK) GetId()(protocol.BINARY_TAG,error){
	return protocol.GetTypeRegistry().GetTagByType(reflect.TypeOf(this).Elem())
}

func (this *OK) Serializable()protocol.ISerializable {
	return this
}

var _ lokas.IComponent = (*PixImg)(nil)

type PixImg struct {
	ecs.Component `json:"-" bson:"-"`
	Path string
	Atlas string
	ModTime time.Time
	OriginX int32
	OriginY int32
	Width int32
	Height int32
	Data []byte
}

func (this *PixImg) ToFrameV0()*FrameV0{
	return &FrameV0{
		Height:         int(this.Height),
		Width:          int(this.Width),
		X:              int(this.OriginX),
		Y:              int(this.OriginY),
		OriginalWidth:  int(this.Width),
		OriginalHeight: int(this.Height),
		OffsetX:        0,
		OffsetY:        0,
	}
}

func (this *PixImg) OnAdd(e lokas.IEntity, r lokas.IRuntime) {

}

func (this *PixImg) OnRemove(e lokas.IEntity, r lokas.IRuntime) {

}

func (this *PixImg) OnCreate(r lokas.IRuntime) {

}

func (this *PixImg) OnDestroy(r lokas.IRuntime) {

}

func (this *PixImg) GetId()(protocol.BINARY_TAG,error){
	return protocol.GetTypeRegistry().GetTagByType(reflect.TypeOf(this).Elem())
}

func (this *PixImg) Serializable()protocol.ISerializable {
	return this
}

var _ lokas.IComponent = (*PixAtlasFile)(nil)

type PixAtlasFile struct {
	ecs.Component `json:"-" bson:"-"`
	Path string
	Atlas *PixImg
	Space int32
	Bleeding bool
	AntiAlias bool
	Images map[string]*PixImg
}

func (this *PixAtlasFile) ToPlistV0()*PlistV0{
	ret:=&PlistV0{Frames: map[string]*FrameV0{}}
	for k,v:=range this.Images {
		ret.Frames[k] = v.ToFrameV0()
	}
	return ret
}

func (this *PixAtlasFile) OnAdd(e lokas.IEntity, r lokas.IRuntime) {

}

func (this *PixAtlasFile) OnRemove(e lokas.IEntity, r lokas.IRuntime) {

}

func (this *PixAtlasFile) OnCreate(r lokas.IRuntime) {

}

func (this *PixAtlasFile) OnDestroy(r lokas.IRuntime) {

}

func (this *PixAtlasFile) GetId()(protocol.BINARY_TAG,error){
	return protocol.GetTypeRegistry().GetTagByType(reflect.TypeOf(this).Elem())
}

func (this *PixAtlasFile) Serializable()protocol.ISerializable {
	return this
}

const (
	TAG_OK  protocol.BINARY_TAG = 6000
	TAG_PIX_ATLAS_FILE  protocol.BINARY_TAG = 6100
	TAG_PIX_IMG  protocol.BINARY_TAG = 6101
)

func init(){
	protocol.GetTypeRegistry().RegistryType(TAG_OK,reflect.TypeOf((*OK)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(TAG_PIX_ATLAS_FILE,reflect.TypeOf((*PixAtlasFile)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(TAG_PIX_IMG,reflect.TypeOf((*PixImg)(nil)).Elem())
}
