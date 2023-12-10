package formdata

import "net/url"

type FormData struct {
	url.Values
}

func (fd *FormData) ToByteArray() []byte {
	return []byte(fd.Encode())
}
