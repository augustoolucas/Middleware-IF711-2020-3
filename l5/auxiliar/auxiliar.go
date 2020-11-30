package auxiliar

//Response em sha256
type Response struct {
	PwSha256 string
}

type Proxy struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

//Request pro hasher
type Request struct {
	Op     string
	Params []interface{}
}

type RequestNaming struct {
	Op    string
	Arg   string
	Proxy Proxy
}

type Invocation struct {
	Host    string
	Port    int
	Request Request
}
