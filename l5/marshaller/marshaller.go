package marshaller

import (
	"../miop"
	"encoding/json"
	"fmt"
)

type Marshaller struct{}

func (Marshaller) Marshall(msg miop.Packet) []byte {
	r, err := json.Marshal(msg)

	if err != nil {
		fmt.Println("Erro Marshall")
	}

	return r
}

func (Marshaller) Unmarshall(msg []byte) miop.Packet {
	r := miop.Packet{}
	err := json.Unmarshal(msg, &r)

	if err != nil {
		fmt.Println("Erro Unmarshall")
	}

	return r
}
