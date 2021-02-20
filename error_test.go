package xerror_test

import (
	"errors"
	"os"
	"testing"

	"github.com/gofor-little/xerror"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	workingDirectory, err := os.Getwd()
	require.NoError(t, err)

	expected := &xerror.Error{
		Err:          errors.New("the thing that failed"),
		FunctionName: "github.com/gofor-little/xerror_test.TestNew",
		FileName:     workingDirectory + "/error_test.go",
		LineNumber:   24,
		Message:      "failed to do something",
	}

	require.Equal(t, expected, xerror.New("failed to do something", errors.New("the thing that failed")))
}
