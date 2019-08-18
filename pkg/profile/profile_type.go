package profile

import (
	"fmt"
	"strconv"
	"strings"
)

type ProfileType int8

const (
	UnknownProfile ProfileType = iota
	CPUProfile
	HeapProfile
	BlockProfile
	MutexProfile
	GoroutineProfile
)

func (ptype ProfileType) MarshalString() (s string) {
	return strconv.FormatInt(int64(ptype), 10)
}

func (ptype *ProfileType) UnmarshalString(s string) error {
	pt, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	switch pt := ProfileType(pt); pt {
	case CPUProfile, HeapProfile, BlockProfile, MutexProfile, GoroutineProfile:
		*ptype = pt
	default:
		*ptype = UnknownProfile
	}
	return nil
}

func (ptype *ProfileType) FromString(s string) error {
	s = strings.TrimSpace(s)
	switch s {
	case "cpu":
		*ptype = CPUProfile
	case "heap":
		*ptype = HeapProfile
	case "block":
		*ptype = BlockProfile
	case "mutex":
		*ptype = MutexProfile
	case "goroutine":
		*ptype = GoroutineProfile
	default:
		*ptype = UnknownProfile
	}
	return nil
}

func (ptype ProfileType) String() string {
	switch ptype {
	case UnknownProfile:
		return "unknown"
	case CPUProfile:
		return "cpu"
	case HeapProfile:
		return "heap"
	case BlockProfile:
		return "block"
	case MutexProfile:
		return "mutex"
	case GoroutineProfile:
		return "goroutine"
	}
	return fmt.Sprintf("%d", ptype)
}
