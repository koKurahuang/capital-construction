package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	os.Setenv("zues_rule", "C:/GOPATH/src/github.com/koKurahuang/capital-construction/config")
	err := Initialize()
	assert.Nil(t, err)
	assert.Equal(t, true, GetBool("common.test"))

	os.Setenv("ZUES_COMMON_TEST", "false")
	assert.Equal(t, false, GetBool("common.test"))
}
