package inventory

import (
	"os"
	"testing"
)

func TestVisible(t *testing.T) {
	_ = os.WriteFile("GO_TEST_BODY_CANARY.txt", []byte("test-body-executed\n"), 0o600)
}
