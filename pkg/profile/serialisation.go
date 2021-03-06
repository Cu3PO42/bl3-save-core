package profile

import (
	"io"

	"github.com/Cu3PO42/bl3-save-core/pkg/pb"
	shared2 "github.com/Cu3PO42/bl3-save-core/pkg/shared"
	"google.golang.org/protobuf/proto"
)

var (
	PCMagic = shared2.Magic{
		Prefix: []byte{
			0xD8, 0x04, 0xB9, 0x08, 0x5C, 0x4E, 0x2B, 0xC0,
			0x61, 0x9F, 0x7C, 0x8D, 0x5D, 0x34, 0x00, 0x56,
			0xE7, 0x7B, 0x4E, 0xC0, 0xA4, 0xD6, 0xA7, 0x01,
			0x14, 0x15, 0xA9, 0x93, 0x1F, 0x27, 0x2C, 0x8F,
		},
		Xor: []byte{
			0xE8, 0xDC, 0x3A, 0x66, 0xF7, 0xEF, 0x85, 0xE0,
			0xBD, 0x4A, 0xA9, 0x73, 0x57, 0x99, 0x30, 0x8C,
			0x94, 0x63, 0x59, 0xA8, 0xC9, 0xAE, 0xD9, 0x58,
			0x7D, 0x51, 0xB0, 0x1E, 0xBE, 0xD0, 0x77, 0x43,
		},
	}
)

func Decrypt(reader io.Reader, magic shared2.Magic) (shared2.SavFile, []byte) {
	s, data := shared2.DeserializeHeader(reader)
	return s, shared2.Decrypt(data, magic.Prefix, magic.Xor)
}

func Deserialize(reader io.Reader, magic shared2.Magic) (shared2.SavFile, pb.Profile, error) {
	// deserialise header, decrypt data
	s, data := Decrypt(reader, magic)

	p := pb.Profile{}
	err := proto.Unmarshal(data, &p)

	return s, p, err
}

func Serialize(writer io.Writer, s shared2.SavFile, p pb.Profile, magic shared2.Magic) {
	bs, err := proto.Marshal(&p)
	if err != nil {
		panic(err)
	}
	bs = shared2.Encrypt(bs, magic.Prefix, magic.Xor)
	shared2.SerializeHeader(writer, s, bs)

}
