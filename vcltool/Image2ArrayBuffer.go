// 由res2go自动生成，不要编辑。
package vcltool

import (
    "github.com/ying32/govcl/vcl"
)

type TImage2ArrayBuffer struct {
    *vcl.TFrame
    DropDownPanel     *vcl.TPanel
    Label1            *vcl.TLabel

    //::private::
    TImage2ArrayBufferFields
}


func NewImage2ArrayBuffer(owner vcl.IComponent) (root *TImage2ArrayBuffer)  {
    vcl.CreateResFrame(owner, &root)
    return
}

var Image2ArrayBufferBytes = []byte("\x54\x50\x46\x30\x12\x54\x49\x6D\x61\x67\x65\x32\x41\x72\x72\x61\x79\x42\x75\x66\x66\x65\x72\x11\x49\x6D\x61\x67\x65\x32\x41\x72\x72\x61\x79\x42\x75\x66\x66\x65\x72\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xFA\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x88\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xFA\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x88\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x0A\x44\x65\x73\x69\x67\x6E\x4C\x65\x66\x74\x03\xDA\x01\x09\x44\x65\x73\x69\x67\x6E\x54\x6F\x70\x03\x14\x01\x00\x06\x54\x50\x61\x6E\x65\x6C\x0D\x44\x72\x6F\x70\x44\x6F\x77\x6E\x50\x61\x6E\x65\x6C\x04\x4C\x65\x66\x74\x02\x0F\x06\x48\x65\x69\x67\x68\x74\x03\xDC\x00\x03\x54\x6F\x70\x02\x0F\x05\x57\x69\x64\x74\x68\x03\x6A\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x41\x72\x6F\x75\x6E\x64\x02\x0F\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xDC\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x6A\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x16\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x4C\x65\x66\x74\x2E\x43\x6F\x6E\x74\x72\x6F\x6C\x07\x0D\x44\x72\x6F\x70\x44\x6F\x77\x6E\x50\x61\x6E\x65\x6C\x13\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x4C\x65\x66\x74\x2E\x53\x69\x64\x65\x07\x09\x61\x73\x72\x43\x65\x6E\x74\x65\x72\x15\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x54\x6F\x70\x2E\x43\x6F\x6E\x74\x72\x6F\x6C\x07\x0D\x44\x72\x6F\x70\x44\x6F\x77\x6E\x50\x61\x6E\x65\x6C\x12\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x54\x6F\x70\x2E\x53\x69\x64\x65\x07\x09\x61\x73\x72\x43\x65\x6E\x74\x65\x72\x17\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x52\x69\x67\x68\x74\x2E\x43\x6F\x6E\x74\x72\x6F\x6C\x07\x0D\x44\x72\x6F\x70\x44\x6F\x77\x6E\x50\x61\x6E\x65\x6C\x14\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x52\x69\x67\x68\x74\x2E\x53\x69\x64\x65\x07\x09\x61\x73\x72\x43\x65\x6E\x74\x65\x72\x18\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x42\x6F\x74\x74\x6F\x6D\x2E\x43\x6F\x6E\x74\x72\x6F\x6C\x07\x0D\x44\x72\x6F\x70\x44\x6F\x77\x6E\x50\x61\x6E\x65\x6C\x15\x41\x6E\x63\x68\x6F\x72\x53\x69\x64\x65\x42\x6F\x74\x74\x6F\x6D\x2E\x53\x69\x64\x65\x07\x09\x61\x73\x72\x43\x65\x6E\x74\x65\x72\x04\x4C\x65\x66\x74\x03\x87\x00\x06\x48\x65\x69\x67\x68\x74\x02\x10\x03\x54\x6F\x70\x02\x66\x05\x57\x69\x64\x74\x68\x02\x5D\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE6\x8B\x96\xE5\x8A\xA8\xE5\x9B\xBE\xE7\x89\x87\xE5\x88\xB0\xE8\xBF\x99\xE9\x87\x8C\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0B\x4C\x61\x62\x65\x6C\x31\x43\x6C\x69\x63\x6B\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(TImage2ArrayBuffer{}, &Image2ArrayBufferBytes)
