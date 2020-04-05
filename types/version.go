package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func NewVersion(parts ...uint32) Version {
	switch len(parts) {
	case 2:
		return Version{major: parts[0], minor: parts[1], size: 2}
	case 3:
		return Version{major: parts[0], minor: parts[1], patch: parts[2], size: 3}
	case 4:
		return Version{major: parts[0], minor: parts[1], patch: parts[2], build: parts[3], size: 4}
	default:
		return Version{size: 0}
	}
}

type Version struct {
	major uint32
	minor uint32
	patch uint32
	build uint32
	size  uint8 //2 = major and minor, 3 = major, minor and patch, 4 = major, minor, patch and build.
}

func (v *Version) Major() uint32 {
	return v.major
}

func (v *Version) Minor() uint32 {
	return v.minor
}

func (v *Version) Patch() uint32 {
	return v.patch
}

func (v *Version) Build() uint32 {
	return v.build
}

func (v *Version) String() string {
	switch v.size {
	case 2:
		return fmt.Sprintf("%d.%d", v.major, v.minor)
	case 3:
		return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
	default:
		return fmt.Sprintf("%d.%d.%d.%d", v.major, v.minor, v.patch, v.build)
	}
}

func ParseVersion(str string) Version {
	validStr, _ := regexp.MatchString(`^\d+\.\d+(\.\d+)?(\.\d+)?$`, str)

	if validStr {
		var size uint8 = 2
		versionSlices := strings.Split(str, ".")

		major, err := strconv.ParseInt(versionSlices[0], 10, 32)
		if err != nil {
			return Version{}
		}

		minor, err := strconv.ParseInt(versionSlices[1], 10, 32)
		if err != nil {
			return Version{}
		}

		var patch int64
		if len(versionSlices) > 2 {
			patch, err = strconv.ParseInt(versionSlices[2], 10, 32)
			if err != nil {
				return Version{}
			}
			size++
		}

		var build int64
		if len(versionSlices) > 3 {
			build, err = strconv.ParseInt(versionSlices[3], 10, 32)
			if err != nil {
				return Version{}
			}
			size++
		}

		return Version{major: uint32(major), minor: uint32(minor), patch: uint32(patch), build: uint32(build), size: size}
	}

	return Version{}
}

func (v *Version) Validate() Fault {
	if v.major == 0 && v.minor == 0 && v.patch == 0 && v.build == 0 {
		return NewFault("invalid_version", fmt.Sprintf("invalid version: %s", v.String()))
	}
	return nil
}

func (v *Version) Compare(operator VersionOperator, version Version) bool {
	switch operator {
	case VersionOperatorGreaterThan:
		return v.GreaterThan(version)
	case VersionOperatorLessThan:
		return v.LessThan(version)
	case VersionOperatorGreaterThanOrEqual:
		return v.GreaterThanOrEqual(version)
	case VersionOperatorLessThanOrEqual:
		return v.LessThanOrEqual(version)
	case VersionOperatorPessimistic:
		return v.PessimisticComparison(version)
	default:
		return false
	}
}

func (v *Version) GreaterThan(version Version) bool {
	if v.major > version.major {
		return true
	}
	if v.major == version.major && v.minor > version.minor {
		return true
	}
	if v.major == version.major && v.minor == version.minor && v.patch > version.patch {
		return true
	}
	if v.major == version.major && v.minor == version.minor && v.patch == version.patch && v.build > version.build {
		return true
	}
	return false
}

func (v *Version) LessThan(version Version) bool {
	if v.major < version.major {
		return true
	}
	if v.major == version.major && v.minor < version.minor {
		return true
	}
	if v.major == version.major && v.minor == version.minor && v.patch < version.patch {
		return true
	}
	if v.major == version.major && v.minor == version.minor && v.patch == version.patch && v.build < version.build {
		return true
	}
	return false
}

func (v *Version) GreaterThanOrEqual(version Version) bool {
	return v.major >= version.major && v.minor >= version.minor && v.patch >= version.patch && v.build >= version.build
}

func (v *Version) LessThanOrEqual(version Version) bool {
	return v.major <= version.major && v.minor <= version.minor && v.patch <= version.patch && v.build <= version.build
}

func (v *Version) PessimisticComparison(version Version) bool {
	if v.LessThan(version) {
		return false
	}
	switch version.size {
	case 2:
		version.major++
		version.minor = 0
		version.patch = 0
		version.build = 0
	case 3:
		version.minor++
		version.patch = 0
		version.build = 0
	case 4:
		version.patch++
		version.build = 0
	default:
		return false
	}
	return v.LessThan(version)
}
