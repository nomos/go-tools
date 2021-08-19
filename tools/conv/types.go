package conv

import "github.com/nomos/go-lokas/util/convert"

func init(){
	registerConvertFunc(convert.STRING,convert.RUNE,convert.String2Rune)
	registerConvertFunc(convert.RUNE,convert.STRING,convert.Rune2String)
	registerConvertFunc(convert.STRING,convert.UNICODE,convert.String2Unicode)
	registerConvertFunc(convert.UNICODE,convert.STRING,convert.Unicode2String)
	registerConvertFunc(convert.UNICODE,convert.RUNE,convert.Unicode2Rune)
	registerConvertFunc(convert.RUNE,convert.UNICODE,convert.Rune2Unicode)
	registerConvertFunc(convert.DECIMAL,convert.BINARY,convert.Decimal2Binary)
	registerConvertFunc(convert.DECIMAL,convert.OCTAL,convert.Decimal2Octal)
	registerConvertFunc(convert.DECIMAL,convert.HEX,convert.Decimal2Hex)
	registerConvertFunc(convert.BINARY,convert.DECIMAL,convert.Binary2Decimal)
	registerConvertFunc(convert.OCTAL,convert.DECIMAL,convert.Octal2Decimal)
	registerConvertFunc(convert.HEX,convert.DECIMAL,convert.Hex2Decimal)
}
