package wxApi

import (
	"testing"
)

func TestVerify(t *testing.T) {
	u := new(UserToken)
	u.Verify()
}
