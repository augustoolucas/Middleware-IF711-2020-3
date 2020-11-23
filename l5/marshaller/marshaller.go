package marshaller

import (
	"Middleware-IF711-2020-3/l5/miop"
	"encoding/json"
	"fmt"
)

type Marshaller struct{}

func (Marshaller) Marshall(msg miop.Packet) []byte {
	r, err := json.Marshal(msg)

	if err != nil {
		fmt.Println("Erro Marshall", err)
	}

	return r
}

func (Marshaller) Unmarshall(msg []byte) miop.Packet {
	r := miop.Packet{}
	err := json.Unmarshal(msg, &r)

	if err != nil {
		fmt.Println("Erro Unmarshall", err)
	}

	return r
}
