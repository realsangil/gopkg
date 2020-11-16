package smtp

import (
	"fmt"
	"io"
	"mime"
	"strings"

	"github.com/pkg/errors"
)

const (
	ErrEmptyAddress     = Error("empty address")
	ErrMalformedAddress = Error("malformed address")
	ErrUndefinedCharset = Error("undefined charset")
)

type Mail struct {
	From        Address
	To          []Address
	CC          []Address
	BCC         []Address
	Body        string
	ContentType string
	Attachments []AttachmentFile
}

type Address struct {
	Name    string
	Address string
}

func (a Address) String() string {
	if a.Name == "" {
		return a.Address
	}
	return fmt.Sprintf("%s <%s>", mime.BEncoding.Encode("UTF-8", a.Name), a.Address)
}

func ParseAddress(str string) (Address, error) {
	if str == "" {
		return Address{}, ErrEmptyAddress
	}
	s := strings.Split(str, " ")
	switch len(s) {
	case 1:
		return Address{Address: str}, nil
	case 2:
		dec := new(mime.WordDecoder)
		name, err := dec.DecodeHeader(s[0])
		if err != nil {
			return Address{}, errors.Wrap(err, "failed to decode name")
		}
		return Address{Name: name, Address: strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(s[1]), "<"), ">")}, nil
	default:
		return Address{}, ErrMalformedAddress
	}
}

type ContentType int

const (
	ContentTypeText ContentType = iota + 1
	ContentTypeHTML
	ContentTypeAlternative
)

var validContentTypes = map[ContentType]struct{}{
	ContentTypeText:        {},
	ContentTypeHTML:        {},
	ContentTypeAlternative: {},
}

// Validate returns ContentType's validation result
func (t ContentType) Validate() error {
	if _, exist := validContentTypes[t]; !exist {
		return ErrUndefinedContentType
	}
	return nil
}

type AttachmentFile struct {
	Name string
	Body io.ReadCloser
}
