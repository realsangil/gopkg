package smtp

import "github.com/pkg/errors"

const (
	// ErrInvalidHost means invalid SMTP host
	ErrInvalidHost = Error("invalid smtp host")
	// ErrInvalidPort means invalid SMTP port
	ErrInvalidPort = Error("invalid smtp port")
	// ErrInvalidUsername means invalid SMTP username
	ErrInvalidUsername = Error("invalid smtp username")
	// ErrInvalidPassword means invalid SMTP password
	ErrInvalidPassword = Error("invalid smtp password")
	// ErrUndefinedSecureProtocol means undefined SMTP secure protocol
	ErrUndefinedSecureProtocol = Error("undefined smtp secure protocol")
	// ErrUndefinedContentType means undefined content type
	ErrUndefinedContentType = Error("undefined smtp content type")
)

// Client is  Wrapping SMTP Client
type Client struct {
	host      string
	port      int
	username  string
	password  string
	mechanism Mechanism
	logger    Logger
}

func (c *Client) Send(m *Mail) error {
	panic("implement me")
}

// NewClient creates Client
// host means smtp host
// port means smtp port
// secureProtocol means secure protocol for auth
// username means username for auth
// password means password for auth
func NewClient(host string, port int, secureProtocol Mechanism, username, password string) (*Client, error) {
	switch {
	case host == "":
		return nil, ErrInvalidHost
	case port == 0:
		return nil, ErrInvalidPort
	case username == "":
		return nil, ErrInvalidUsername
	case password == "":
		return nil, ErrInvalidPassword
	}
	if err := secureProtocol.Validate(); err != nil {
		return nil, errors.WithStack(err)
	}
	return &Client{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}, nil
}

// Mechanism means secure protocol for auth
type Mechanism int

const (
	// MechanismUndefined means undefined protocol
	MechanismUndefined Mechanism = iota - 1
	// MechanismNone means none protocol
	MechanismNone
	// MechanismSSL means SSL
	MechanismSSL
	// MechanismTLS means TLS
	MechanismTLS
)

// IsNone returns true If p and MechanismNone are the same.
func (m Mechanism) IsNone() bool {
	return m == MechanismNone
}

// IsSSL returns true If p and MechanismSSL are the same.
func (m Mechanism) IsSSL() bool {
	return m == MechanismSSL
}

// IsTLS returns true If p and MechanismTLS are the same.
func (m Mechanism) IsTLS() bool {
	return m == MechanismTLS
}

// Validate returns Mechanism's validation result
func (m Mechanism) Validate() error {
	for m < MechanismNone || m > MechanismTLS {
		return ErrUndefinedSecureProtocol
	}
	return nil
}

func (m Mechanism) String() string {
	switch m {
	case MechanismNone:
		return ""
	case MechanismSSL:
		return "SSL"
	case MechanismTLS:
		return "TLS"
	default:
		return "undefined"
	}
}
