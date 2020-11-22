package clientproxy

type ClientProxy struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

func (proxy ClientProxy) Hash(rawString string) string {
	params := make([]interface{}, 1)
	params[0] = rawString
	return ""
}
