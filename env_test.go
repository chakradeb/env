package env

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvParser(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	_ = os.Setenv("Name", "envParser")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT"`
		Name string `env:"Name"`
	}
	config := args{}
	err := Parse(&config)

	assert.NoError(t, err, "unexpected error while parsing")
	assert.Equal(t, config.Port, 5000, "expectation mismatch for port")
	assert.Equal(t, config.Name, "envParser", "expectation mismatch for name")
}

func TestEnvParserWithoutPointer(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	_ = os.Setenv("Name", "envParser")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT"`
		Name string `env:"Name"`
	}
	config := args{}
	err := Parse(config)

	assert.Error(t, err, "unexpected error while parsing")
	assert.Equal(t, err, errors.New("env: expected ptr but got struct"), "wrong error message")
}

func TestEnvParserWithNonStructValue(t *testing.T) {
	config := "abcd"
	err := Parse(&config)

	assert.Error(t, err, "unexpected error while parsing")
	assert.Equal(t, err, errors.New("env: expected struct but got string"), "wrong error message")
}

func TestEnvParserWithUnsupportedTypes(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	defer os.Clearenv()

	type args struct {
		Port int16 `env:"PORT"`
	}
	config := args{}
	err := Parse(&config)

	assert.Error(t, err, "unexpected error while parsing")
	assert.Equal(t, err, errors.New("env: int16 is not a supported type"), "wrong error message")
}

func TestEnvParserWithWrongValues(t *testing.T) {
	_ = os.Setenv("PORT", "5a")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT"`
	}
	config := args{}
	err := Parse(&config)

	assert.Error(t, err, "unexpected error while parsing")
	assert.Equal(t, err, errors.New("env: strconv.ParseInt: parsing \"5a\": invalid syntax"), "wrong error message")
}

func TestEnvParserWithoutEnvironmentVariable(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT"`
		Name string `env:"Name"`
	}
	config := args{}
	err := Parse(&config)

	assert.NoError(t, err, "unexpected error while parsing")
	assert.Equal(t, config.Port, 5000, "expectation mismatch for port")
	assert.Equal(t, config.Name, "", "expectation mismatch for name")
}
