package types

import (
	"fmt"
	"regexp"
)

type VersionString string

func (versionString VersionString) ToVersion() Version {
	return ParseVersion(string(versionString))
}

func (versionString VersionString) Validate() Fault {
	match, _ := regexp.MatchString(`^\d+\.\d+(\.\d+)?(\.\d+)?$`, string(versionString))
	if !match {
		return NewFault("invalid_version_string", fmt.Sprintf("invalid version string: '%s'", versionString))
	}
	return nil
}
