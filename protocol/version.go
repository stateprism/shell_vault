package protocol

import "fmt"

func (v *Version) Display() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
