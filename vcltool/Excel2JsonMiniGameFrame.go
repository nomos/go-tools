// 由res2go自动生成，不要编辑。
package vcltool

import (
    "github.com/ying32/govcl/vcl"
)

type TExcel2JsonMiniGameFrame struct {
    *vcl.TFrame
    OpenExcelDirButton      *vcl.TSpeedButton
    OpenDistDirButton       *vcl.TSpeedButton
    ExcelDirLabel           *vcl.TEdit
    DistDirLabel            *vcl.TEdit
    IndieFolderCheck        *vcl.TCheckBox
    Label1                  *vcl.TLabel
    Label2                  *vcl.TLabel
    GenerateButton          *vcl.TButton

    //::private::
    TExcel2JsonMiniGameFrameFields
}


func NewExcel2JsonMiniGameFrame(owner vcl.IComponent) (root *TExcel2JsonMiniGameFrame)  {
    vcl.CreateResFrame(owner, &root)
    return
}

var Excel2JsonMiniGameFrameBytes = []byte("\x54\x50\x46\x30\x18\x54\x45\x78\x63\x65\x6C\x32\x4A\x73\x6F\x6E\x4D\x69\x6E\x69\x47\x61\x6D\x65\x46\x72\x61\x6D\x65\x17\x45\x78\x63\x65\x6C\x32\x4A\x73\x6F\x6E\x4D\x69\x6E\x69\x47\x61\x6D\x65\x46\x72\x61\x6D\x65\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xD4\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x0B\x02\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xD4\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x0B\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x0A\x44\x65\x73\x69\x67\x6E\x4C\x65\x66\x74\x03\x4A\x02\x09\x44\x65\x73\x69\x67\x6E\x54\x6F\x70\x03\xC7\x00\x00\x0C\x54\x53\x70\x65\x65\x64\x42\x75\x74\x74\x6F\x6E\x12\x4F\x70\x65\x6E\x45\x78\x63\x65\x6C\x44\x69\x72\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x1E\x03\x54\x6F\x70\x02\x28\x05\x57\x69\x64\x74\x68\x02\x1E\x0A\x47\x6C\x79\x70\x68\x2E\x44\x61\x74\x61\x0A\x3A\x10\x00\x00\x36\x10\x00\x00\x42\x4D\x36\x10\x00\x00\x00\x00\x00\x00\x36\x00\x00\x00\x28\x00\x00\x00\x20\x00\x00\x00\x20\x00\x00\x00\x01\x00\x20\x00\x00\x00\x00\x00\x00\x10\x00\x00\x64\x00\x00\x00\x64\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2E\x88\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2E\x89\xD8\xFE\x20\x80\xDF\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x30\x8A\xD9\xFF\x58\xB5\xE4\xFF\x64\xC1\xE7\xFF\x63\xC0\xE7\xFF\x62\xC0\xE6\xFF\x61\xBF\xE6\xFF\x61\xBF\xE6\xFF\x60\xBE\xE6\xFF\x5F\xBE\xE5\xFF\x5E\xBD\xE5\xFF\x5E\xBD\xE5\xFF\x5D\xBD\xE5\xFF\x5C\xBC\xE4\xFF\x5C\xBC\xE4\xFF\x5B\xBB\xE4\xFF\x5A\xBB\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x33\x8E\xD9\xFF\x2B\x87\xD9\x35\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x30\x89\xD9\xFF\x52\xAC\xE3\xFF\x6B\xC5\xEA\xFF\x6A\xC5\xEA\xFF\x69\xC4\xE9\xFF\x69\xC4\xE9\xFF\x68\xC4\xE9\xFF\x67\xC3\xE8\xFF\x67\xC3\xE8\xFF\x66\xC2\xE8\xFF\x65\xC2\xE8\xFF\x64\xC1\xE7\xFF\x64\xC1\xE7\xFF\x63\xC0\xE7\xFF\x62\xC0\xE6\xFF\x61\xBF\xE6\xFF\x61\xBF\xE6\xFF\x60\xBE\xE6\xFF\x5F\xBE\xE5\xFF\x5E\xBD\xE5\xFF\x5E\xBD\xE5\xFF\x5D\xBD\xE5\xFF\x5C\xBC\xE4\xFF\x5C\xBC\xE4\xFF\x5B\xBB\xE4\xFF\x3D\x99\xDC\xFF\x2C\x87\xD8\x68\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x47\xA1\xE0\xFF\x72\xCA\xED\xFF\x72\xCA\xEC\xFF\x71\xC9\xEC\xFF\x70\xC9\xEC\xFF\x6F\xC8\xEB\xFF\x6F\xC8\xEB\xFF\x6E\xC7\xEB\xFF\x6D\xC7\xEB\xFF\x6C\xC6\xEA\xFF\x6C\xC6\xEA\xFF\x6B\xC5\xEA\xFF\x6A\xC5\xEA\xFF\x69\xC4\xE9\xFF\x69\xC4\xE9\xFF\x68\xC4\xE9\xFF\x67\xC3\xE8\xFF\x67\xC3\xE8\xFF\x66\xC2\xE8\xFF\x65\xC2\xE8\xFF\x64\xC1\xE7\xFF\x64\xC1\xE7\xFF\x63\xC0\xE7\xFF\x62\xC0\xE6\xFF\x4C\xA8\xE0\xFF\x30\x8A\xD8\xA5\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x3A\x94\xDB\xFF\x7A\xCF\xEF\xFF\x79\xCE\xEF\xFF\x78\xCE\xEF\xFF\x77\xCD\xEF\xFF\x77\xCD\xEE\xFF\x76\xCC\xEE\xFF\x75\xCC\xEE\xFF\x74\xCB\xED\xFF\x74\xCB\xED\xFF\x73\xCB\xED\xFF\x72\xCA\xED\xFF\x72\xCA\xEC\xFF\x71\xC9\xEC\xFF\x70\xC9\xEC\xFF\x6F\xC8\xEB\xFF\x6F\xC8\xEB\xFF\x6E\xC7\xEB\xFF\x6D\xC7\xEB\xFF\x6C\xC6\xEA\xFF\x6C\xC6\xEA\xFF\x6B\xC5\xEA\xFF\x6A\xC5\xEA\xFF\x69\xC4\xE9\xFF\x5C\xB7\xE6\xFF\x31\x8B\xD9\xDC\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x2F\x8A\xD9\xFF\x7E\xD0\xF1\xFF\x80\xD3\xF2\xFF\x7F\xD2\xF2\xFF\x7F\xD2\xF1\xFF\x7E\xD2\xF1\xFF\x7D\xD1\xF1\xFF\x7D\xD1\xF0\xFF\x7C\xD0\xF0\xFF\x7B\xD0\xF0\xFF\x7A\xCF\xF0\xFF\x7A\xCF\xEF\xFF\x79\xCE\xEF\xFF\x78\xCE\xEF\xFF\x77\xCD\xEF\xFF\x77\xCD\xEE\xFF\x76\xCC\xEE\xFF\x75\xCC\xEE\xFF\x74\xCB\xED\xFF\x74\xCB\xED\xFF\x73\xCB\xED\xFF\x72\xCA\xED\xFF\x72\xCA\xEC\xFF\x71\xC9\xEC\xFF\x6D\xC7\xEC\xFF\x2E\x89\xD9\xFA\x20\x80\xDF\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x38\x92\xDA\xFF\x75\xC7\xEF\xFF\x88\xD8\xF5\xFF\x87\xD7\xF4\xFF\x86\xD7\xF4\xFF\x85\xD6\xF4\xFF\x85\xD6\xF3\xFF\x84\xD5\xF3\xFF\x83\xD5\xF3\xFF\x82\xD4\xF3\xFF\x82\xD4\xF2\xFF\x81\xD3\xF2\xFF\x80\xD3\xF2\xFF\x7F\xD2\xF2\xFF\x7F\xD2\xF1\xFF\x7E\xD2\xF1\xFF\x7D\xD1\xF1\xFF\x7D\xD1\xF0\xFF\x7C\xD0\xF0\xFF\x7B\xD0\xF0\xFF\x7A\xCF\xF0\xFF\x7A\xCF\xEF\xFF\x79\xCE\xEF\xFF\x78\xCE\xEF\xFF\x77\xCD\xEF\xFF\x37\x90\xDB\xF1\x2B\x87\xD9\x35\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x40\x9C\xDC\xFF\x67\xBA\xEB\xFF\x8F\xDC\xF7\xFF\x8E\xDC\xF7\xFF\x8D\xDB\xF7\xFF\x8D\xDB\xF7\xFF\x8C\xDA\xF6\xFF\x8B\xDA\xF6\xFF\x8A\xD9\xF6\xFF\x8A\xD9\xF5\xFF\x89\xD8\xF5\xFF\x88\xD8\xF5\xFF\x88\xD8\xF5\xFF\x87\xD7\xF4\xFF\x86\xD7\xF4\xFF\x85\xD6\xF4\xFF\x85\xD6\xF3\xFF\x84\xD5\xF3\xFF\x83\xD5\xF3\xFF\x82\xD4\xF3\xFF\x82\xD4\xF2\xFF\x81\xD3\xF2\xFF\x80\xD3\xF2\xFF\x7F\xD2\xF2\xFF\x7F\xD2\xF1\xFF\x4A\xA2\xE1\xF5\x2C\x87\xD8\x68\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x4B\xA8\xDF\xFF\x54\xA9\xE4\xFF\x96\xE1\xFA\xFF\x95\xE0\xFA\xFF\x95\xE0\xFA\xFF\x94\xDF\xF9\xFF\x93\xDF\xF9\xFF\x93\xDF\xF9\xFF\x92\xDE\xF8\xFF\x91\xDE\xF8\xFF\x90\xDD\xF8\xFF\x90\xDD\xF8\xFF\x8F\xDC\xF7\xFF\x8E\xDC\xF7\xFF\x8D\xDB\xF7\xFF\x8D\xDB\xF7\xFF\x8C\xDA\xF6\xFF\x8B\xDA\xF6\xFF\x8A\xD9\xF6\xFF\x8A\xD9\xF5\xFF\x89\xD8\xF5\xFF\x88\xD8\xF5\xFF\x88\xD8\xF5\xFF\x87\xD7\xF4\xFF\x86\xD7\xF4\xFF\x61\xB5\xE9\xFF\x31\x8B\xD8\xA5\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x5B\xB6\xE3\xFF\x3F\x97\xDD\xFF\x9E\xE6\xFD\xFF\x9D\xE5\xFD\xFF\x9C\xE5\xFC\xFF\x9B\xE4\xFC\xFF\x9B\xE4\xFC\xFF\x9A\xE3\xFC\xFF\x99\xE3\xFB\xFF\x98\xE2\xFB\xFF\x98\xE2\xFB\xFF\x97\xE1\xFA\xFF\x96\xE1\xFA\xFF\x95\xE0\xFA\xFF\x95\xE0\xFA\xFF\x94\xDF\xF9\xFF\x93\xDF\xF9\xFF\x93\xDF\xF9\xFF\x92\xDE\xF8\xFF\x91\xDE\xF8\xFF\x90\xDD\xF8\xFF\x90\xDD\xF8\xFF\x8F\xDC\xF7\xFF\x8E\xDC\xF7\xFF\x8D\xDB\xF7\xFF\x79\xC9\xF1\xFF\x33\x8C\xDA\xDC\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x6D\xC6\xE9\xFF\x30\x8A\xDA\xFF\x9F\xE6\xFE\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA2\xE8\xFF\xFF\xA1\xE8\xFE\xFF\xA0\xE7\xFE\xFF\xA0\xE7\xFE\xFF\x9F\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9D\xE5\xFD\xFF\x9C\xE5\xFC\xFF\x9B\xE4\xFC\xFF\x9B\xE4\xFC\xFF\x9A\xE3\xFC\xFF\x99\xE3\xFB\xFF\x98\xE2\xFB\xFF\x98\xE2\xFB\xFF\x97\xE1\xFA\xFF\x96\xE1\xFA\xFF\x95\xE0\xFA\xFF\x95\xE0\xFA\xFF\x90\xDC\xF8\xFF\x2F\x89\xD9\xFA\x20\x80\xDF\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x76\xCC\xEC\xFF\x3D\x95\xDC\xFF\x8A\xD5\xF7\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA2\xE8\xFF\xFF\xA1\xE8\xFE\xFF\xA0\xE7\xFE\xFF\xA0\xE7\xFE\xFF\x9F\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9D\xE5\xFD\xFF\x9C\xE5\xFC\xFF\x9B\xE4\xFC\xFF\x3C\x93\xDC\xF1\x2B\x87\xD9\x35\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x7E\xD1\xEF\xFF\x4D\xA4\xE1\xFF\x73\xC1\xEF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\x57\xAA\xE6\xF5\x2C\x87\xD8\x68\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x85\xD6\xF2\xFF\x60\xB4\xE7\xFF\x5A\xAD\xE7\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\x72\xC0\xEF\xFF\x33\x8D\xDA\xA5\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x8D\xDA\xF5\xFF\x77\xC7\xEE\xFF\x42\x99\xDE\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\x8A\xD4\xF7\xFF\x35\x8E\xDA\xDD\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x94\xDF\xF8\xFF\x8F\xDB\xF6\xFF\x33\x8D\xDA\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2E\x87\xD8\xFB\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x9C\xE4\xFB\xFF\x9A\xE3\xFA\xFF\x99\xE2\xFA\xFF\x97\xE1\xF9\xFF\x96\xE0\xF8\xFF\x94\xDF\xF8\xFF\x93\xDE\xF7\xFF\x91\xDD\xF7\xFF\x90\xDC\xF6\xFF\x8E\xDB\xF5\xFF\x8D\xDA\xF5\xFF\x8B\xD9\xF4\xFF\x8A\xD8\xF4\xFF\x88\xD8\xF3\xFF\x87\xD7\xF2\xFF\x85\xD6\xF2\xFF\x84\xD5\xF1\xFF\x82\xD4\xF1\xFF\x81\xD3\xF0\xFF\x7F\xD2\xF0\xFF\x7E\xD1\xEF\xFF\x7C\xD0\xEE\xFF\x7B\xCF\xEE\xFF\x79\xCE\xED\xFF\x2D\x87\xD8\xFF\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA3\xE8\xFE\xFF\xA2\xE7\xFD\xFF\xA0\xE6\xFD\xFF\x9F\xE5\xFC\xFF\x9D\xE4\xFB\xFF\x9C\xE4\xFB\xFF\x9A\xE3\xFA\xFF\x99\xE2\xFA\xFF\x97\xE1\xF9\xFF\x96\xE0\xF8\xFF\x94\xDF\xF8\xFF\x93\xDE\xF7\xFF\x91\xDD\xF7\xFF\x90\xDC\xF6\xFF\x8E\xDB\xF5\xFF\x8D\xDA\xF5\xFF\x8B\xD9\xF4\xFF\x8A\xD8\xF4\xFF\x88\xD8\xF3\xFF\x87\xD7\xF2\xFF\x85\xD6\xF2\xFF\x84\xD5\xF1\xFF\x82\xD4\xF1\xFF\x81\xD3\xF0\xFF\x2D\x87\xD8\xFF\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA5\xE9\xFE\xFF\xA3\xE8\xFE\xFF\xA1\xE7\xFD\xFF\xA0\xE6\xFD\xFF\x9F\xE5\xFC\xFF\x9D\xE4\xFB\xFF\x9C\xE4\xFB\xFF\x9A\xE3\xFA\xFF\x98\xE2\xFA\xFF\x97\xE1\xF9\xFF\x96\xE0\xF8\xFF\x94\xDF\xF8\xFF\x93\xDE\xF7\xFF\x91\xDD\xF7\xFF\x90\xDC\xF6\xFF\x8E\xDB\xF5\xFF\x8D\xDA\xF5\xFF\x8B\xD9\xF4\xFF\x89\xD8\xF4\xFF\x88\xD8\xF3\xFF\x2D\x87\xD8\xFF\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\x84\xCF\xF4\xFF\x30\x89\xD9\xFB\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2F\x88\xD9\xF0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\x8A\xD3\xF6\xFF\x32\x8B\xD9\xF4\x2C\x87\xD7\x40\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2E\x88\xD8\xF3\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2F\x88\xD9\xF7\x2D\x88\xDA\x3E\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0C\x54\x53\x70\x65\x65\x64\x42\x75\x74\x74\x6F\x6E\x11\x4F\x70\x65\x6E\x44\x69\x73\x74\x44\x69\x72\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x1E\x03\x54\x6F\x70\x02\x68\x05\x57\x69\x64\x74\x68\x02\x1E\x0A\x47\x6C\x79\x70\x68\x2E\x44\x61\x74\x61\x0A\x3A\x10\x00\x00\x36\x10\x00\x00\x42\x4D\x36\x10\x00\x00\x00\x00\x00\x00\x36\x00\x00\x00\x28\x00\x00\x00\x20\x00\x00\x00\x20\x00\x00\x00\x01\x00\x20\x00\x00\x00\x00\x00\x00\x10\x00\x00\x64\x00\x00\x00\x64\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2E\x88\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2E\x89\xD8\xFE\x20\x80\xDF\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x30\x8A\xD9\xFF\x58\xB5\xE4\xFF\x64\xC1\xE7\xFF\x63\xC0\xE7\xFF\x62\xC0\xE6\xFF\x61\xBF\xE6\xFF\x61\xBF\xE6\xFF\x60\xBE\xE6\xFF\x5F\xBE\xE5\xFF\x5E\xBD\xE5\xFF\x5E\xBD\xE5\xFF\x5D\xBD\xE5\xFF\x5C\xBC\xE4\xFF\x5C\xBC\xE4\xFF\x5B\xBB\xE4\xFF\x5A\xBB\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x59\xBA\xE3\xFF\x33\x8E\xD9\xFF\x2B\x87\xD9\x35\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x30\x89\xD9\xFF\x52\xAC\xE3\xFF\x6B\xC5\xEA\xFF\x6A\xC5\xEA\xFF\x69\xC4\xE9\xFF\x69\xC4\xE9\xFF\x68\xC4\xE9\xFF\x67\xC3\xE8\xFF\x67\xC3\xE8\xFF\x66\xC2\xE8\xFF\x65\xC2\xE8\xFF\x64\xC1\xE7\xFF\x64\xC1\xE7\xFF\x63\xC0\xE7\xFF\x62\xC0\xE6\xFF\x61\xBF\xE6\xFF\x61\xBF\xE6\xFF\x60\xBE\xE6\xFF\x5F\xBE\xE5\xFF\x5E\xBD\xE5\xFF\x5E\xBD\xE5\xFF\x5D\xBD\xE5\xFF\x5C\xBC\xE4\xFF\x5C\xBC\xE4\xFF\x5B\xBB\xE4\xFF\x3D\x99\xDC\xFF\x2C\x87\xD8\x68\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x47\xA1\xE0\xFF\x72\xCA\xED\xFF\x72\xCA\xEC\xFF\x71\xC9\xEC\xFF\x70\xC9\xEC\xFF\x6F\xC8\xEB\xFF\x6F\xC8\xEB\xFF\x6E\xC7\xEB\xFF\x6D\xC7\xEB\xFF\x6C\xC6\xEA\xFF\x6C\xC6\xEA\xFF\x6B\xC5\xEA\xFF\x6A\xC5\xEA\xFF\x69\xC4\xE9\xFF\x69\xC4\xE9\xFF\x68\xC4\xE9\xFF\x67\xC3\xE8\xFF\x67\xC3\xE8\xFF\x66\xC2\xE8\xFF\x65\xC2\xE8\xFF\x64\xC1\xE7\xFF\x64\xC1\xE7\xFF\x63\xC0\xE7\xFF\x62\xC0\xE6\xFF\x4C\xA8\xE0\xFF\x30\x8A\xD8\xA5\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x3A\x94\xDB\xFF\x7A\xCF\xEF\xFF\x79\xCE\xEF\xFF\x78\xCE\xEF\xFF\x77\xCD\xEF\xFF\x77\xCD\xEE\xFF\x76\xCC\xEE\xFF\x75\xCC\xEE\xFF\x74\xCB\xED\xFF\x74\xCB\xED\xFF\x73\xCB\xED\xFF\x72\xCA\xED\xFF\x72\xCA\xEC\xFF\x71\xC9\xEC\xFF\x70\xC9\xEC\xFF\x6F\xC8\xEB\xFF\x6F\xC8\xEB\xFF\x6E\xC7\xEB\xFF\x6D\xC7\xEB\xFF\x6C\xC6\xEA\xFF\x6C\xC6\xEA\xFF\x6B\xC5\xEA\xFF\x6A\xC5\xEA\xFF\x69\xC4\xE9\xFF\x5C\xB7\xE6\xFF\x31\x8B\xD9\xDC\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x2F\x8A\xD9\xFF\x7E\xD0\xF1\xFF\x80\xD3\xF2\xFF\x7F\xD2\xF2\xFF\x7F\xD2\xF1\xFF\x7E\xD2\xF1\xFF\x7D\xD1\xF1\xFF\x7D\xD1\xF0\xFF\x7C\xD0\xF0\xFF\x7B\xD0\xF0\xFF\x7A\xCF\xF0\xFF\x7A\xCF\xEF\xFF\x79\xCE\xEF\xFF\x78\xCE\xEF\xFF\x77\xCD\xEF\xFF\x77\xCD\xEE\xFF\x76\xCC\xEE\xFF\x75\xCC\xEE\xFF\x74\xCB\xED\xFF\x74\xCB\xED\xFF\x73\xCB\xED\xFF\x72\xCA\xED\xFF\x72\xCA\xEC\xFF\x71\xC9\xEC\xFF\x6D\xC7\xEC\xFF\x2E\x89\xD9\xFA\x20\x80\xDF\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x38\x92\xDA\xFF\x75\xC7\xEF\xFF\x88\xD8\xF5\xFF\x87\xD7\xF4\xFF\x86\xD7\xF4\xFF\x85\xD6\xF4\xFF\x85\xD6\xF3\xFF\x84\xD5\xF3\xFF\x83\xD5\xF3\xFF\x82\xD4\xF3\xFF\x82\xD4\xF2\xFF\x81\xD3\xF2\xFF\x80\xD3\xF2\xFF\x7F\xD2\xF2\xFF\x7F\xD2\xF1\xFF\x7E\xD2\xF1\xFF\x7D\xD1\xF1\xFF\x7D\xD1\xF0\xFF\x7C\xD0\xF0\xFF\x7B\xD0\xF0\xFF\x7A\xCF\xF0\xFF\x7A\xCF\xEF\xFF\x79\xCE\xEF\xFF\x78\xCE\xEF\xFF\x77\xCD\xEF\xFF\x37\x90\xDB\xF1\x2B\x87\xD9\x35\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x40\x9C\xDC\xFF\x67\xBA\xEB\xFF\x8F\xDC\xF7\xFF\x8E\xDC\xF7\xFF\x8D\xDB\xF7\xFF\x8D\xDB\xF7\xFF\x8C\xDA\xF6\xFF\x8B\xDA\xF6\xFF\x8A\xD9\xF6\xFF\x8A\xD9\xF5\xFF\x89\xD8\xF5\xFF\x88\xD8\xF5\xFF\x88\xD8\xF5\xFF\x87\xD7\xF4\xFF\x86\xD7\xF4\xFF\x85\xD6\xF4\xFF\x85\xD6\xF3\xFF\x84\xD5\xF3\xFF\x83\xD5\xF3\xFF\x82\xD4\xF3\xFF\x82\xD4\xF2\xFF\x81\xD3\xF2\xFF\x80\xD3\xF2\xFF\x7F\xD2\xF2\xFF\x7F\xD2\xF1\xFF\x4A\xA2\xE1\xF5\x2C\x87\xD8\x68\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x4B\xA8\xDF\xFF\x54\xA9\xE4\xFF\x96\xE1\xFA\xFF\x95\xE0\xFA\xFF\x95\xE0\xFA\xFF\x94\xDF\xF9\xFF\x93\xDF\xF9\xFF\x93\xDF\xF9\xFF\x92\xDE\xF8\xFF\x91\xDE\xF8\xFF\x90\xDD\xF8\xFF\x90\xDD\xF8\xFF\x8F\xDC\xF7\xFF\x8E\xDC\xF7\xFF\x8D\xDB\xF7\xFF\x8D\xDB\xF7\xFF\x8C\xDA\xF6\xFF\x8B\xDA\xF6\xFF\x8A\xD9\xF6\xFF\x8A\xD9\xF5\xFF\x89\xD8\xF5\xFF\x88\xD8\xF5\xFF\x88\xD8\xF5\xFF\x87\xD7\xF4\xFF\x86\xD7\xF4\xFF\x61\xB5\xE9\xFF\x31\x8B\xD8\xA5\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x5B\xB6\xE3\xFF\x3F\x97\xDD\xFF\x9E\xE6\xFD\xFF\x9D\xE5\xFD\xFF\x9C\xE5\xFC\xFF\x9B\xE4\xFC\xFF\x9B\xE4\xFC\xFF\x9A\xE3\xFC\xFF\x99\xE3\xFB\xFF\x98\xE2\xFB\xFF\x98\xE2\xFB\xFF\x97\xE1\xFA\xFF\x96\xE1\xFA\xFF\x95\xE0\xFA\xFF\x95\xE0\xFA\xFF\x94\xDF\xF9\xFF\x93\xDF\xF9\xFF\x93\xDF\xF9\xFF\x92\xDE\xF8\xFF\x91\xDE\xF8\xFF\x90\xDD\xF8\xFF\x90\xDD\xF8\xFF\x8F\xDC\xF7\xFF\x8E\xDC\xF7\xFF\x8D\xDB\xF7\xFF\x79\xC9\xF1\xFF\x33\x8C\xDA\xDC\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x6D\xC6\xE9\xFF\x30\x8A\xDA\xFF\x9F\xE6\xFE\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA2\xE8\xFF\xFF\xA1\xE8\xFE\xFF\xA0\xE7\xFE\xFF\xA0\xE7\xFE\xFF\x9F\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9D\xE5\xFD\xFF\x9C\xE5\xFC\xFF\x9B\xE4\xFC\xFF\x9B\xE4\xFC\xFF\x9A\xE3\xFC\xFF\x99\xE3\xFB\xFF\x98\xE2\xFB\xFF\x98\xE2\xFB\xFF\x97\xE1\xFA\xFF\x96\xE1\xFA\xFF\x95\xE0\xFA\xFF\x95\xE0\xFA\xFF\x90\xDC\xF8\xFF\x2F\x89\xD9\xFA\x20\x80\xDF\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x76\xCC\xEC\xFF\x3D\x95\xDC\xFF\x8A\xD5\xF7\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA2\xE8\xFF\xFF\xA1\xE8\xFE\xFF\xA0\xE7\xFE\xFF\xA0\xE7\xFE\xFF\x9F\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9E\xE6\xFD\xFF\x9D\xE5\xFD\xFF\x9C\xE5\xFC\xFF\x9B\xE4\xFC\xFF\x3C\x93\xDC\xF1\x2B\x87\xD9\x35\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x7E\xD1\xEF\xFF\x4D\xA4\xE1\xFF\x73\xC1\xEF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\x57\xAA\xE6\xF5\x2C\x87\xD8\x68\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x85\xD6\xF2\xFF\x60\xB4\xE7\xFF\x5A\xAD\xE7\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\x72\xC0\xEF\xFF\x33\x8D\xDA\xA5\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x8D\xDA\xF5\xFF\x77\xC7\xEE\xFF\x42\x99\xDE\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\xA3\xE9\xFF\xFF\x8A\xD4\xF7\xFF\x35\x8E\xDA\xDD\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x94\xDF\xF8\xFF\x8F\xDB\xF6\xFF\x33\x8D\xDA\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2E\x87\xD8\xFB\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\x9C\xE4\xFB\xFF\x9A\xE3\xFA\xFF\x99\xE2\xFA\xFF\x97\xE1\xF9\xFF\x96\xE0\xF8\xFF\x94\xDF\xF8\xFF\x93\xDE\xF7\xFF\x91\xDD\xF7\xFF\x90\xDC\xF6\xFF\x8E\xDB\xF5\xFF\x8D\xDA\xF5\xFF\x8B\xD9\xF4\xFF\x8A\xD8\xF4\xFF\x88\xD8\xF3\xFF\x87\xD7\xF2\xFF\x85\xD6\xF2\xFF\x84\xD5\xF1\xFF\x82\xD4\xF1\xFF\x81\xD3\xF0\xFF\x7F\xD2\xF0\xFF\x7E\xD1\xEF\xFF\x7C\xD0\xEE\xFF\x7B\xCF\xEE\xFF\x79\xCE\xED\xFF\x2D\x87\xD8\xFF\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA3\xE8\xFE\xFF\xA2\xE7\xFD\xFF\xA0\xE6\xFD\xFF\x9F\xE5\xFC\xFF\x9D\xE4\xFB\xFF\x9C\xE4\xFB\xFF\x9A\xE3\xFA\xFF\x99\xE2\xFA\xFF\x97\xE1\xF9\xFF\x96\xE0\xF8\xFF\x94\xDF\xF8\xFF\x93\xDE\xF7\xFF\x91\xDD\xF7\xFF\x90\xDC\xF6\xFF\x8E\xDB\xF5\xFF\x8D\xDA\xF5\xFF\x8B\xD9\xF4\xFF\x8A\xD8\xF4\xFF\x88\xD8\xF3\xFF\x87\xD7\xF2\xFF\x85\xD6\xF2\xFF\x84\xD5\xF1\xFF\x82\xD4\xF1\xFF\x81\xD3\xF0\xFF\x2D\x87\xD8\xFF\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA5\xE9\xFE\xFF\xA3\xE8\xFE\xFF\xA1\xE7\xFD\xFF\xA0\xE6\xFD\xFF\x9F\xE5\xFC\xFF\x9D\xE4\xFB\xFF\x9C\xE4\xFB\xFF\x9A\xE3\xFA\xFF\x98\xE2\xFA\xFF\x97\xE1\xF9\xFF\x96\xE0\xF8\xFF\x94\xDF\xF8\xFF\x93\xDE\xF7\xFF\x91\xDD\xF7\xFF\x90\xDC\xF6\xFF\x8E\xDB\xF5\xFF\x8D\xDA\xF5\xFF\x8B\xD9\xF4\xFF\x89\xD8\xF4\xFF\x88\xD8\xF3\xFF\x2D\x87\xD8\xFF\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\x84\xCF\xF4\xFF\x30\x89\xD9\xFB\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2F\x88\xD9\xF0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2D\x87\xD8\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\xA6\xEA\xFF\xFF\x8A\xD3\xF6\xFF\x32\x8B\xD9\xF4\x2C\x87\xD7\x40\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x2E\x88\xD8\xF3\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2D\x87\xD8\xFF\x2F\x88\xD9\xF7\x2D\x88\xDA\x3E\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x54\x45\x64\x69\x74\x0D\x45\x78\x63\x65\x6C\x44\x69\x72\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x30\x05\x57\x69\x64\x74\x68\x03\x98\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x05\x54\x45\x64\x69\x74\x0C\x44\x69\x73\x74\x44\x69\x72\x4C\x61\x62\x65\x6C\x04\x4C\x65\x66\x74\x02\x50\x06\x48\x65\x69\x67\x68\x74\x02\x16\x03\x54\x6F\x70\x02\x70\x05\x57\x69\x64\x74\x68\x03\x98\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x10\x49\x6E\x64\x69\x65\x46\x6F\x6C\x64\x65\x72\x43\x68\x65\x63\x6B\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x12\x03\x54\x6F\x70\x03\xA0\x00\x05\x57\x69\x64\x74\x68\x02\x6C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE5\xAF\xBC\xE5\x87\xBA\xE5\x88\xB0\xE7\x8B\xAC\xE7\xAB\x8B\xE6\x96\x87\xE4\xBB\xB6\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x10\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x02\x55\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x11\x45\x78\x63\x65\x6C\xE6\x96\x87\xE4\xBB\xB6\xE8\xB7\xAF\xE5\xBE\x84\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x20\x06\x48\x65\x69\x67\x68\x74\x02\x10\x03\x54\x6F\x70\x02\x58\x05\x57\x69\x64\x74\x68\x02\x35\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBE\x93\xE5\x87\xBA\xE8\xB7\xAF\xE5\xBE\x84\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x0E\x47\x65\x6E\x65\x72\x61\x74\x65\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x03\xC8\x00\x06\x48\x65\x69\x67\x68\x74\x02\x29\x03\x54\x6F\x70\x03\x98\x00\x05\x57\x69\x64\x74\x68\x02\x79\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\x94\x9F\xE6\x88\x90\xE9\x85\x8D\xE7\xBD\xAE\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(TExcel2JsonMiniGameFrame{}, &Excel2JsonMiniGameFrameBytes)
