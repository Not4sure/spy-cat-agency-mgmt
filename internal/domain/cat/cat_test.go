package cat_test

import (
	"testing"
	"time"

	"github.com/not4sure/spy-cat-agency-mgmt/internal/domain/cat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCat(t *testing.T) {
	testCases := []struct {
		name       string
		catName    string
		experience uint
		breed      string
		salary     uint
	}{
		{
			name:       "OK",
			catName:    "Nancy",
			experience: 1,
			breed:      "Test",
			salary:     100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, err := cat.New(tc.name, tc.experience, tc.breed, tc.salary)

			require.NoError(t, err)

			assert.NotZero(t, c.ID())
			assert.WithinDuration(t, time.Now(), c.CreatedAt(), time.Second)
			assert.Equal(t, tc.name, c.Name())
			assert.Equal(t, tc.experience, c.YearsOfExperience())
			assert.Equal(t, tc.breed, c.Breed())
			assert.Equal(t, tc.salary, c.Salary())
		})
	}
}

func TestInvalidCat(t *testing.T) {
	testCases := []struct {
		testName   string
		name       string
		experience uint
		breed      string
		salary     uint
		err        error
	}{
		{
			testName:   "Empty name",
			name:       "",
			experience: 1,
			breed:      "Test",
			salary:     100,
			err:        cat.ErrEmptyName,
		},
		{
			testName:   "Empty breed",
			name:       "Test",
			experience: 1,
			breed:      "",
			salary:     100,
			err:        cat.ErrEmptyBreed,
		},
		{
			testName:   "Zero salary",
			name:       "Test",
			experience: 1,
			breed:      "Test",
			salary:     0,
			err:        cat.ErrZeroSalary,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			c, err := cat.New(tc.name, tc.experience, tc.breed, tc.salary)

			require.Zero(t, c)
			require.Error(t, err)
			require.Equal(t, tc.err, err)
		})
	}
}
