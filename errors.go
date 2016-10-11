package cgroups

import "errors"

var (
	ErrInvalidPid          = errors.New("cgroups: pid must be greater than 0")
	ErrMountPointNotExist  = errors.New("cgroups: cgroup mountpoint does not exist")
	ErrInvalidFormat       = errors.New("cgroups: parsing file with invalid format failed")
	ErrFreezerNotSupported = errors.New("cgroups: freezer cgroup not supported on this system")
)
