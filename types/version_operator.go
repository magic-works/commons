package types

import "fmt"

type VersionOperator string

const (
	VersionOperatorLessThan           VersionOperator = "<"
	VersionOperatorGreaterThan        VersionOperator = ">"
	VersionOperatorLessThanOrEqual    VersionOperator = "<="
	VersionOperatorGreaterThanOrEqual VersionOperator = ">="
	VersionOperatorPessimistic        VersionOperator = "~>"
)

func (versionOperator VersionOperator) Validate() Fault {
	switch versionOperator {
	case VersionOperatorLessThan,
		VersionOperatorGreaterThan,
		VersionOperatorLessThanOrEqual,
		VersionOperatorGreaterThanOrEqual,
		VersionOperatorPessimistic:
		return nil
	}
	return NewFault("invalid_version_operator", fmt.Sprintf("invalid version operator: %s", versionOperator))
}
