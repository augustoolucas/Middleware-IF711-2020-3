package auxiliar

//Response em sha256
type Response struct {
	PwSha256 string
}

//Request pro hasher
type Request struct {
	Op     string
	Params []interface{}
}

type Invocation struct {
	Host    string
	Port    int
	Request Request
}
