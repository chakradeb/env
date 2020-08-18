package env

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvParser(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	_ = os.Setenv("NAME", "envParser")
	_ = os.Setenv("ID", "14")
	_ = os.Setenv("ROW", "2")
	_ = os.Setenv("TOTAL", "87")
	_ = os.Setenv("REFERENCE", "6152")
	_ = os.Setenv("AVERAGE", "61.45")
	_ = os.Setenv("PERCENT", "86.74")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT"`
		Name string `env:"NAME"`
		ID int8 `env:"ID"`
		Row int16 `env:"ROW"`
		Total int32 `env:"TOTAL"`
		Ref int64 `env:"REFERENCE"`
		Avg float32 `env:"AVERAGE"`
		Pct float64 `env:"PERCENT"`
	}
	config := args{}
	err := Parse(&config)

	assert.NoError(t, err, "unexpected error while parsing")
	assert.Equal(t, config.Port, 5000, "expectation mismatch for port")
	assert.Equal(t, config.Name, "envParser", "expectation mismatch for name")
	assert.Equal(t, config.ID, int8(14), "expectation mismatch for id")
	assert.Equal(t, config.Row, int16(2), "expectation mismatch for row")
	assert.Equal(t, config.Total, int32(87), "expectation mismatch for total")
	assert.Equal(t, config.Ref, int64(6152), "expectation mismatch for reference")
	assert.Equal(t, config.Avg, float32(61.45), "expectation mismatch for average")
	assert.Equal(t, config.Pct, 86.74, "expectation mismatch for percent")
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
	_ = os.Setenv("IS_DEFAULT", "false")
	defer os.Clearenv()

	type args struct {
		IsDefault bool `env:"IS_DEFAULT"`
	}
	config := args{}
	err := Parse(&config)

	assert.Error(t, err, "unexpected error while parsing")
	assert.Equal(t, err, errors.New("env: bool is not a supported type"), "wrong error message")
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
