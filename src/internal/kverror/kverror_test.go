package kverror_test

import (
	"errors"
	"testing"

	"github.com/SenRecep/redisclone/src/internal/kverror"
)

var ErrInner = errors.New("inner")

func TestError(t *testing.T) {
	err := kverror.New("some error", true)
	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	shouldEqual := "some error"
	if kvErr.Message != shouldEqual {
		t.Errorf("error message does not match, want: %s, got: %s", shouldEqual, kvErr.Message)
	}

	if !kvErr.Loggable {
		t.Errorf("error should be loggable, want: %t, got: %t", !kvErr.Loggable, kvErr.Loggable)
	}
}

func TestWrap(t *testing.T) {
	err := kverror.New("some error", false)
	wrappedErr := err.Wrap(ErrInner)

	var kvErr *kverror.Error

	if !errors.As(wrappedErr, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	if kvErr.Err == nil {
		t.Errorf("wrapped error can not be nil, want: %v, got: nil", kvErr.Err)
	}

	shouldEqual := "inner, some error"
	if err.Error() != shouldEqual {
		t.Errorf("wrapped error does not match, want: %s, got: %s", shouldEqual, err.Error())
	}
}

func TestUnwrap(t *testing.T) {
	err := kverror.New("some error", false)
	wrappedErr := err.Wrap(ErrInner)

	var kvErr *kverror.Error

	if !errors.As(wrappedErr, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	shouldEqual := "inner"
	unwrappedErr := kvErr.Unwrap()
	if unwrappedErr.Error() != shouldEqual {
		t.Errorf("unwrapped error does not match, want: %s, got: %s", shouldEqual, unwrappedErr.Error())
	}
}

func TestAddDataDestroyData(t *testing.T) {
	err := kverror.New("some error", false).AddData("hello")

	var kvErr *kverror.Error

	if !errors.As(err, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	if kvErr.Data == nil {
		t.Errorf("data should not be nil, want: %v, got: nil", kvErr.Data)
	}

	shouldEqual := "hello"
	data, ok := kvErr.Data.(string)
	if !ok {
		t.Error("data should be assertable to string")
	}

	if data != shouldEqual {
		t.Errorf("data does not match, want: %s, got: %s", shouldEqual, data)
	}

	shouldEqual = "some error"
	if err.Error() != shouldEqual {
		t.Errorf("error does not match, want: %s, got: %s", shouldEqual, err.Error())
	}

	err = err.DestoryData()
	if !errors.As(err, &kvErr) {
		t.Errorf("error does not match the target type, want: %T, got: %v", kvErr, err)
	}

	if kvErr.Data != nil {
		t.Errorf("data should be nil, want: nil, got: %v", kvErr.Data)
	}
}
