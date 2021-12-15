package interception

type InterceptionKeyState uint16
type InterceptionFilter uint16

const (
	INTERCEPTION_KEY_DOWN             InterceptionKeyState = 0x00
	INTERCEPTION_KEY_UP               InterceptionKeyState = 0x01
	INTERCEPTION_KEY_E0               InterceptionKeyState = 0x02
	INTERCEPTION_KEY_E1               InterceptionKeyState = 0x04
	INTERCEPTION_KEY_TERMSRV_SET_LED  InterceptionKeyState = 0x08
	INTERCEPTION_KEY_TERMSRV_SHADOW   InterceptionKeyState = 0x10
	INTERCEPTION_KEY_TERMSRV_VKPACKET InterceptionKeyState = 0x2

	INTERCEPTION_FILTER_KEY_NONE             InterceptionFilter = 0x0000
	INTERCEPTION_FILTER_KEY_ALL              InterceptionFilter = 0xFFFF
	INTERCEPTION_FILTER_KEY_DOWN             InterceptionFilter = 0x01
	INTERCEPTION_FILTER_KEY_UP               InterceptionFilter = 0x01 << 1
	INTERCEPTION_FILTER_KEY_E0               InterceptionFilter = 0x02 << 1
	INTERCEPTION_FILTER_KEY_E1               InterceptionFilter = 0x04 << 1
	INTERCEPTION_FILTER_KEY_TERMSRV_SET_LED  InterceptionFilter = 0x08 << 1
	INTERCEPTION_FILTER_KEY_TERMSRV_SHADOW   InterceptionFilter = 0x10 << 1
	INTERCEPTION_FILTER_KEY_TERMSRV_VKPACKET InterceptionFilter = 0x2 << 1
)

type InterceptionKeyStroke struct {
	code        uint16
	state       uint16
	information uint32
}
