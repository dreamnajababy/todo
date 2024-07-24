package todo

import (
	"encoding/base64"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func assertGotWant(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	got := err.Error()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestTodoEntity(t *testing.T) {
	t.Run("Create new todo successfully", func(t *testing.T) {
		createdAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
		imageUrl := base64.StdEncoding.EncodeToString([]byte("__TEST_IMAGE_URL__"))
		want := Todo{
			Id:          uuid.New(),
			Title:       "__TEST_TITLE__",
			Description: "__TEST_DESCRIPTION__",
			CreatedAt:   createdAt,
			Image:       imageUrl,
			Status:      IN_PROGRESS,
		}
		got, _ := NewTodo(
			want.Id.String(),
			want.Title,
			want.Description,
			want.CreatedAt.Format(time.RFC3339),
			want.Image,
			want.Status,
		)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
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

		assertGotWant(t, err, want)
	})
}
