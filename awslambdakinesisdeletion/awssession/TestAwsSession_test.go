package awssession

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAwsSession(t *testing.T) {
	_, err := CreateSession()
	assert.NoError(t, err, "error occurs at create session")
}
