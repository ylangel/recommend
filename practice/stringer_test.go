package practice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ylangel/recommend/practice"
)

func TestStatus(t *testing.T) {
	assert.Equal(t, practice.Success.String(), "成功")
}