package player

import "errors"

type PlayerStatus struct {
	value string
}

type Status int

const (
	Attend Status = iota
	Break
	Absence
)

func (s Status) String() string {
	switch s {
	case Attend:
		return "attend"
	case Break:
		return "break"
	default:
		return "absence"
	}
}

func NewPlayerStatus(status Status) (PlayerStatus, error) {
	return PlayerStatus{status.String()}, nil
}

func PlayerStatusFromStr(status string) (PlayerStatus, error) {
	switch status {
	case Attend.String():
		status, _ := NewPlayerStatus(Attend)
		return status, nil
	case Break.String():
		status, _ := NewPlayerStatus(Break)
		return status, nil
	case Absence.String():
		status, _ := NewPlayerStatus(Absence)
		return status, nil
	default:
		return PlayerStatus{}, errors.New("invalid player status")
	}
}

func (p *PlayerStatus) Value() string {
	return p.value
}
