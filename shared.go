package shared

type Request struct {
	Op string
	P1 int
	P2 int
}

type Reply struct {
	Result []interface{}
}

const CALCULATOR_PORT = 3300
const SAMPLE_SIZE = 10
