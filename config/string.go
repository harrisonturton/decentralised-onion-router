package config

import (
	"fmt"
	"strconv"
)

func (c Config) String() string {
	return fmt.Sprintf("Config:%s%s%s", c.Server, c.ExitNode, c.RelayNode)
}

func (c ServerConfig) String() string {
	builder := NewStringBuilder("")
	builder.addIndentLine(1, "Server Config")
	builder.addIndentLine(2, "Host: "+c.Host)
	builder.addIndentLine(2, "Port: "+c.Port)
	builder.addIndentLine(2, "IsTLS: "+strconv.FormatBool(c.IsTls))
	if c.TlsKeyFilename == nil {
		builder.addIndentLine(2, "TLS Key File: nil")
	} else {
		builder.addIndentLine(2, "TLS Key File: "+*c.TlsKeyFilename)
	}
	if c.TlsCrtFilename == nil {
		builder.addIndentLine(2, "TLS Crt File: nil")
	} else {
		builder.addIndentLine(2, "TLS Crt File: "+*c.TlsCrtFilename)
	}
	return builder.String
}

func (c RelayConfig) String() string {
	builder := NewStringBuilder("")
	builder.addIndentLine(1, "Relay Config:")
	builder.addIndentLine(2, "Timeout: "+strconv.Itoa(c.Timeout))
	return builder.String
}

func (c ExitConfig) String() string {
	builder := NewStringBuilder("")
	builder.addIndentLine(1, "Exit Config:")
	builder.addIndentLine(2, "Force HTTPS: "+strconv.FormatBool(c.ForceHttps))
	builder.addIndentLine(2, "Timeout: "+strconv.Itoa(c.Timeout))
	return builder.String
}

/* Helper functions that makes the config printing easier */

type StringBuilder struct {
	String string
}

func NewStringBuilder(str string) *StringBuilder {
	return &StringBuilder{
		String: str,
	}
}

func (s *StringBuilder) addIndentLine(indent int, str string) *StringBuilder {
	s.String += "\n"
	for i := 0; i < indent; i++ {
		s.String += "  "
	}
	s.String += str
	return s
}
