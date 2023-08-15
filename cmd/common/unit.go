package common

import (
	"github.com/google/uuid"
)

type UnitStructType struct {
	Pid string
	Tid string
}

func ZeroPid() string {
	return "1122011070J0000"
}

func ZeroTid() string {
	return "00000000-0000-0000-0000-000000000000"
}

func GetPid(prefix string) string {
	if len(prefix) == 0 {
		id := uuid.New()
		return ZeroPid()[:11] + id.String()[:4]
	}
	return ZeroPid()[:11] + prefix[:4]
}

func GetTid(prefix string) string {
	if len(prefix) == 0 {
		id := uuid.New()
		return ZeroTid()[:24] + id.String()[24:len(ZeroTid())]
	}
	return ZeroTid()[:24] + prefix[:12]
}

func (u *UnitStructType) SetPid(prefix string) string {
	u.Tid = GetPid(prefix)

	return u.Pid
}

func (u *UnitStructType) SetTid(prefix string) string {
	u.Tid = GetTid(prefix)

	return u.Tid
}

func (u *UnitStructType) Zero() {
	u.Pid = ZeroPid()
	u.Tid = ZeroTid()
}
