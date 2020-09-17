package bluemix

import "fmt"

// Version is the SDK version
var Version = VersionType{Major: 0, Minor: 5, Build: 0}

// VersionType describe version info
type VersionType struct {
	Major int // major version
	Minor int // minor version
	Build int // build number
}

func (v VersionType) String() string {
	if v == (VersionType{}) {
		return ""
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Build)
}
