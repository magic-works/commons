package types

import (
	"fmt"
	"regexp"
)

type Identifier string

func (identifier Identifier) Validate() Fault {
	var causes []Fault

	idStr := string(identifier)
	match, _ := regexp.MatchString(`^[a-z]+(_?[a-z0-9]+)*$`, idStr)
	if !match {
		fault := NewFault("invalid_expression", fmt.Sprintf("invalid expression: '%s'", idStr))
		causes = append(causes, fault)
	}

	const minLength = 2
	if len(idStr) < minLength {
		fault := NewFault("length_too_small", fmt.Sprintf("length too small: length must be greater or equals than %d: '%s'", minLength, idStr))
		causes = append(causes, fault)
	}

	if len(causes) > 0 {
		return NewFault("invalid_identifier", fmt.Sprintf("invalid identifier: '%s'", identifier), causes...)
	}

	return nil
}
