package i2page

import (
	"context"
)

type Page struct {
	Content [65536]byte
	Size    uint32
}

func (p Page) Bytes() []byte {
	var u4 uint16 = uint16(p.Size)
	return p.Content[:u4]
}

type IntegerToPage func(context.Context, int64) (Page, error)
