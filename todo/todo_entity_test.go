package todo

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

// Can sort the data by Title or Date or Status fields
// Can search the data by Title or Description fields
// The TODO application can UPDATE a task with the following requirements
// Can update a task by ID field
// Can update Title, Description, Date, Image, and Status fields corresponding to the requirements from the CREATE feature

func TestTodoEntity(t *testing.T) {
	t.Run("Id must be required", func(t *testing.T) {
		want := "id cannot be empty"
		_, err := NewTodo(
			"",
			"__TEST_TITLE__",
			"__TEST_DESCRIPTION__",
			time.Now().UTC().Format("2006-01-02 15:04:05 MST"),
			"__TEST_IMAGE_URL",
			IN_PROGRESS,
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Id must be UUID format", func(t *testing.T) {
		want := "id must be UUID format"
		_, err := NewTodo(
			"__INVALID_UUID_FORMAT__",
			"__TEST_TITLE__",
			"__TEST_DESCRIPTION__",
			time.Now().UTC().Format("2006-01-02 15:04:05 MST"),
			"__TEST_IMAGE_URL",
			IN_PROGRESS,
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("CreatedAt must be required", func(t *testing.T) {
		want := "created_at cannot be empty"
		_, err := NewTodo(
			uuid.New().String(),
			"__TEST_TITLE__",
			"__TEST_DESCRIPTION__",
			"",
			"__TEST_IMAGE_URL",
			IN_PROGRESS,
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("CreatedAt must be RFC3339 format", func(t *testing.T) {
		want := "created_at must be RFC3339 format"
		_, err := NewTodo(
			uuid.New().String(),
			"__TEST_TITLE__",
			"__TEST_DESCRIPTION__",
			"2021-01-01 00:00:00 UTC",
			"__TEST_IMAGE_URL",
			IN_PROGRESS,
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Todo title cannot be empty", func(t *testing.T) {
		want := "title cannot be empty"

		_, err := NewTodo(
			uuid.New().String(),
			"",
			"__TEST_DESCRIPTION__",
			time.Now().UTC().Format("2006-01-02 15:04:05 MST"),
			"__TEST_IMAGE_URL",
			IN_PROGRESS,
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Todo status cannot be empty", func(t *testing.T) {
		want := "status cannot be empty"

		_, err := NewTodo(
			uuid.New().String(),
			"__TEST_TITLE__",
			"__TEST_DESCRIPTION__",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_IMAGE_URL",
			"",
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Todo status must be IN_PROGRESS or COMPLETE", func(t *testing.T) {
		want := "status must be IN_PROGRESS or COMPLETE"

		_, err := NewTodo(
			uuid.New().String(),
			"__TEST_TITLE__",
			"__TEST_DESCRIPTION__",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_IMAGE_URL",
			"__INVALID_STATUS__",
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Todo title must less than 100 characters", func(t *testing.T) {
		want := "title must not over 100 characters"

		_, err := NewTodo(
			uuid.New().String(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ac tincidunt nulla. Nulla facilisi. Donec nec",
			"__TEST_DESCRIPTION__",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_IMAGE_URL",
			IN_PROGRESS,
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("Image must be Base64 Encode format", func(t *testing.T) {
		want := "image must be Base64 Encode format"

		_, err := NewTodo(
			uuid.New().String(),
			"__TEST_TITLE__",
			"__TEST_DESCRIPTION__",
			time.Now().UTC().Format(time.RFC3339),
			"__INVALID_IMAGE_URL__",
			IN_PROGRESS,
		)

		if err == nil {
			t.Error("expected error, got nil")
		}
		got := err.Error()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
