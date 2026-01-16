package utils

import (
	"fmt"
	"math/big"
)

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}

type Signature struct {
	R *big.Int
	S *big.Int
}
