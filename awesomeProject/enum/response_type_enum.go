package enum

type ResponseType int

const (
	OperateOk   ResponseType = 200
	OperateFail ResponseType = 500
)

func (p ResponseType) String() string {
	switch p {
	case OperateOk:
		return "OK"
	case OperateFail:
		return "Fail"
	default:
		return "UNKNOWN"
	}
}
