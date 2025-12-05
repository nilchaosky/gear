package snowflake

import "testing"

func Test(t *testing.T) {
	Register(0, 0)
	t.Log(Builder.NextVal())
}
