package miop

type Packet struct {
	Hdr Header
	Br  Body
}

type Header struct {
	Magic       string
	Version     string
	ByteOrder   bool
	MessageType int
	Size        int
}

type Body struct {
	ReqHeader RequestHeader
	ReqBody   RequestBody
	RepHeader ReplyHeader
	RepBody   ReplyBody
}

type RequestHeader struct {
	Context          string
	RequestID        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

type RequestBody struct {
	Body []interface{}
}

type ReplyHeader struct {
	Context   string
	RequestID int
	Status    int
}

type ReplyBody struct {
	OperationResult interface{}
}
