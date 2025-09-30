package mission

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Mission struct {
	id          uuid.UUID
	createdAt   time.Time
	isCompleted bool
	isDeleted   bool

	asigneeID *uuid.UUID

	targets []Target
}

func NewMission(id uuid.UUID, opts ...MissionOption) (*Mission, error) {
	m := &Mission{
		id:        id,
		createdAt: time.Now(),
		targets:   make([]Target, 0),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m, nil
}

func (m *Mission) MarkCompleted() {
	for _, t := range m.targets {
		t.MarkCompleted()
	}

	m.isCompleted = true
}

func (m *Mission) MarkDeleted() error {
	if m.IsAssigned() {
		return errors.New("cannot delete assigned mission")
	}

	m.isDeleted = true
	return nil
}

func (m *Mission) IsModifiable() bool {
	return !m.IsDeleted() && !m.IsCompleted()
}

func (m *Mission) IsDeleted() bool {
	return m.isDeleted
}

func (m *Mission) IsCompleted() bool {
	return m.isCompleted
}

func (m *Mission) IsAssigned() bool {
	return m.asigneeID != nil
}
