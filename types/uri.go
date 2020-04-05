package types

import (
	"fmt"
	"regexp"
)

type URI string

func (uri URI) Validate() Fault {
	var causes []Fault

	match, _ := regexp.MatchString(`^(\w+:(/?/?))?[^\s]+$`, string(uri))
	if !match {
		fault := NewFault("invalid_expression", fmt.Sprintf("invalid expression: '%s'", string(uri)))
		causes = append(causes, fault)
	}

	if len(causes) > 0 {
		return NewFault("invalid_uri", fmt.Sprintf("invalid URI: '%s'", uri), causes...)
	}
	return nil
}
