package mw256enc

import (
	"bytes"
	"compress/zlib"
	"github.com/ehmry/encoding/base256"
	"io/ioutil"
)

var E = base256.Braille

func Decode(s string) ([]byte, error) {
	dst, err := E.DecodeString(s)
	if err != nil {
		return nil, err
	}
	b := bytes.NewReader(dst)
	buf, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	res, err := ioutil.ReadAll(buf)
	return res, err
}

func Encode(src []byte) string {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(src)
	w.Close()
	buf := b.Bytes()
	return E.EncodeToString(buf)
}
