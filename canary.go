package inventory

import "os"

func init() {
	_ = os.WriteFile("GO_TEST_LIST_CANARY.txt", []byte("repository-init-executed\n"), 0o600)
}
