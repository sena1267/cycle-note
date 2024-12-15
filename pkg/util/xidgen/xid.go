package xidgen

import "github.com/rs/xid"

type XIDGenerator struct{}

func (g XIDGenerator) Generate() string {
	return xid.New().String()
}
