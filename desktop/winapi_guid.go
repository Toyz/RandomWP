// build +windows

package desktop

import (
	"bytes"
	"encoding/hex"
	"regexp"
)

type GUID struct {
	data [16]byte
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// format: "374DE290-123F-4565-9164-39C4925E467B"
func GUIDNew(s string) GUID {
	var m GUID

	p := regexp.MustCompile("(\\w+)-(\\w+)-(\\w+)-(\\w+)-(\\w+)")
	ret := p.FindStringSubmatch(s)

	var buf *bytes.Buffer = &bytes.Buffer{}

	bb, err := hex.DecodeString(ret[1])
	if err != nil {
		panic(err)
	}
	reverse(bb)
	buf.Write(bb)

	bb, err = hex.DecodeString(ret[2])
	if err != nil {
		panic(err)
	}
	reverse(bb)
	buf.Write(bb)

	bb, err = hex.DecodeString(ret[3])
	if err != nil {
		panic(err)
	}
	reverse(bb)
	buf.Write(bb)

	bb, err = hex.DecodeString(ret[4])
	if err != nil {
		panic(err)
	}
	buf.Write(bb)

	bb, err = hex.DecodeString(ret[5])
	if err != nil {
		panic(err)
	}
	buf.Write(bb)

	bb = buf.Bytes()

	for i := range bb {
		m.data[i] = bb[i]
	}

	return m
}

func (m GUID) String() string {
	b1 := m.data[0:4]
	reverse(b1)
	b2 := m.data[4:6]
	reverse(b2)
	b3 := m.data[6:8]
	reverse(b3)
	b4 := m.data[8:10]
	b5 := m.data[10:16]
	str := ""
	str += hex.EncodeToString(b1) + "-"
	str += hex.EncodeToString(b2) + "-"
	str += hex.EncodeToString(b3) + "-"
	str += hex.EncodeToString(b4) + "-"
	str += hex.EncodeToString(b5)
	return str
}
