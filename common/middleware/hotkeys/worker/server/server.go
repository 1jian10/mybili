package server

import (
	"encoding/binary"
	"github.com/panjf2000/gnet"
)

type Codec struct {
}

func Serve(Addr string) error {
	codec := gnet.NewLengthFieldBasedFrameCodec(
		gnet.EncoderConfig{
			ByteOrder:                       binary.BigEndian,
			LengthFieldLength:               4,
			LengthIncludesLengthFieldLength: false,
		},
		gnet.DecoderConfig{
			ByteOrder:           binary.BigEndian,
			LengthFieldLength:   4,
			InitialBytesToStrip: 4,
		})

	handler := &Handler{}
	return gnet.Serve(handler, Addr, gnet.WithMulticore(true), gnet.WithCodec(codec))
}
