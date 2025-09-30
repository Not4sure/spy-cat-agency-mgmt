package mission

import "errors"

type Target struct {
	name    string
	country string
	notes   string

	isCompleted bool
}

func NewTarget(
	name string,
	country string,
	notes string,
) *Target {
	return &Target{
		name:    name,
		country: country,
		notes:   notes,
	}
}

func (t *Target) SetNotes(notes string) error {
	if t.IsCompleted() {
		return errors.New("cannot update notes on completed target")
	}

	t.notes = notes
	return nil
}

func (t *Target) MarkCompleted() {
	t.isCompleted = true
}

func (t *Target) IsCompleted() bool {
	return t.isCompleted
}

func (t *Target) Country() string {
	return t.country
}

func (t *Target) Name() string {
	return t.name
}

func (t *Target) Notes() string {
	return t.notes
}
