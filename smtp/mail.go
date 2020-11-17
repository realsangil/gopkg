package smtp

import (
	"fmt"
	"io"
	"mime"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

const (
	ErrEmptyAddress     = Error("empty address")
	ErrMalformedAddress = Error("malformed address")
	ErrUndefinedCharset = Error("undefined charset")
)

var regexpAddress = regexp.MustCompile(`(.*)\<(.+)\>`)

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

func ParseAddress(addresses ...string) ([]Address, error) {
	var result []Address
	for _, addr := range addresses {
		if addr == "" {
			return nil, ErrEmptyAddress
		}
		s := regexpAddress.FindStringSubmatch(addr)
		switch len(s) {
		case 3:
			dec := new(mime.WordDecoder)
			name, err := dec.DecodeHeader(strings.TrimSpace(s[1]))
			if err != nil {
				return nil, errors.Wrap(err, "failed to decode header")
			}
			result = append(result, Address{Name: name, Address: s[2]})
		default:
			result = append(result, Address{Address: addr})
		}
	}
	return result, nil
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
