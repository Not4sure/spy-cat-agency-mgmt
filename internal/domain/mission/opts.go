package mission

import "github.com/google/uuid"

type MissionOption func(*Mission)

func WithTarget(t Target) MissionOption {
	return func(m *Mission) {
		m.targets = append(m.targets, t)
	}
}

func WithAsigneeID(id uuid.UUID) MissionOption {
	return func(m *Mission) {
		m.asigneeID = &id
	}
}
