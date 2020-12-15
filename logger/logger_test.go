package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInit to test the init method
func TestInit(t *testing.T) {
	assert.Nil(t, log)
	lc := Config{
		FileName:   "test",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     1,
		LogLevel:   InfoLevel,
	}
	Init(lc)
	assert.NotNil(t, log)
}

// TestField method to test the field method
func TestField(t *testing.T) {
	lc := Config{
		FileName:   "test",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     1,
		LogLevel:   InfoLevel,
	}
	Init(lc)
	f := Field("mykey", "myvalue")
	assert.NotNil(t, f)
	assert.EqualValues(t, "mykey", f.Key)
}
