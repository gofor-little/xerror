package xerror_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gofor-little/xerror"
)

func TestNew(t *testing.T) {
	workingDirectory, err := os.Getwd()
	require.NoError(t, err)

	expected := &xerror.Error{
		Err:          errors.New("the thing that failed"),
		FunctionName: "github.com/gofor-little/xerror_test.TestNew",
		FileName:     workingDirectory + "/error_test.go",
		LineNumber:   25,
		Message:      "",
	}

	require.Equal(t, expected, xerror.New("the thing that failed"))
}

func TestNewf(t *testing.T) {
	workingDirectory, err := os.Getwd()
	require.NoError(t, err)

	expected := &xerror.Error{
		Err:          errors.New("the thing that failed"),
		FunctionName: "github.com/gofor-little/xerror.Newf",
		FileName:     workingDirectory + "/error.go",
		LineNumber:   33,
		Message:      "",
	}

	require.Equal(t, expected, xerror.Newf("the thing that %s", "failed"))
}

func TestWrap(t *testing.T) {
	workingDirectory, err := os.Getwd()
	require.NoError(t, err)

	expected := &xerror.Error{
		Err:          errors.New("the thing that failed"),
		FunctionName: "github.com/gofor-little/xerror_test.TestWrap",
		FileName:     workingDirectory + "/error_test.go",
		LineNumber:   55,
		Message:      "failed to do something",
	}

	require.Equal(t, expected, xerror.Wrap("failed to do something", errors.New("the thing that failed")))
}

func TestWrapf(t *testing.T) {
	workingDirectory, err := os.Getwd()
	require.NoError(t, err)

	expected := &xerror.Error{
		Err:          errors.New("the thing that failed"),
		FunctionName: "github.com/gofor-little/xerror.Wrapf",
		FileName:     workingDirectory + "/error.go",
		LineNumber:   51,
		Message:      "failed to do something: was this",
	}

	require.Equal(t, expected, xerror.Wrapf("failed to do something: %s", errors.New("the thing that failed"), "was this"))
}
