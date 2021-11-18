package conv

import "github.com/nomos/go-lokas/util/convert"

func init(){
	RegisterConvertFunc(convert.STRING,convert.RUNE,convert.String2Rune)
	RegisterConvertFunc(convert.RUNE,convert.STRING,convert.Rune2String)
	RegisterConvertFunc(convert.STRING,convert.UNICODE,convert.String2Unicode)
	RegisterConvertFunc(convert.UNICODE,convert.STRING,convert.Unicode2String)
	RegisterConvertFunc(convert.UNICODE,convert.RUNE,convert.Unicode2Rune)
	RegisterConvertFunc(convert.RUNE,convert.UNICODE,convert.Rune2Unicode)
	RegisterConvertFunc(convert.DECIMAL,convert.BINARY,convert.Decimal2Binary)
	RegisterConvertFunc(convert.DECIMAL,convert.OCTAL,convert.Decimal2Octal)
	RegisterConvertFunc(convert.DECIMAL,convert.HEX,convert.Decimal2Hex)
	RegisterConvertFunc(convert.BINARY,convert.DECIMAL,convert.Binary2Decimal)
	RegisterConvertFunc(convert.OCTAL,convert.DECIMAL,convert.Octal2Decimal)
	RegisterConvertFunc(convert.HEX,convert.DECIMAL,convert.Hex2Decimal)
}
