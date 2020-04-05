package types

import (
	"fmt"
	"regexp"
)

type GUID string

func (guid GUID) Validate() Fault {
	var causes []Fault
	match, _ := regexp.MatchString(`^[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$`, string(guid))
	if !match {
		fault := NewFault("invalid_expression", fmt.Sprintf("invalid expression: '%s'", string(guid)))
		causes = append(causes, fault)
	}

	if len(causes) > 0 {
		return NewFault("invalid_guid", fmt.Sprintf("invalid GUID: '%s'", guid), causes...)
	}
	return nil
}
