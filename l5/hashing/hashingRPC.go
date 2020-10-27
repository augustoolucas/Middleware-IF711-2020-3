package hashing

import (
	"crypto/sha256"
	"errors"
	"shared"
)

//Response em sha256
type Response struct {
	pwSha256 [32]byte
}

//Request pro hasher
type Request struct {
	pwRaw string
}

type Handler struct{}

func (h *Handler) hashPw(req Request, res *Response) (err error) {
	if req.pwRaw == "" {
		err = errors.New("não houve password informado")
		shared.ChecaErro(err, "não houve password informado")
		return
	}
	res.pwSha256 = sha256.Sum256([]byte(req.pwRaw))
	return
}
