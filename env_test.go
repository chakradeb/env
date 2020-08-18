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
	_ = os.Setenv("IS_DEFAULT", "false")
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
		IsDefault bool `env:"IS_DEFAULT"`
	}
	config := args{}
	errs := Parse(&config)

	assert.Equal(t, len(errs), 0, "unexpected error while parsing")
	assert.Equal(t, config.Port, 5000, "expectation mismatch for port")
	assert.Equal(t, config.Name, "envParser", "expectation mismatch for name")
	assert.Equal(t, config.ID, int8(14), "expectation mismatch for id")
	assert.Equal(t, config.Row, int16(2), "expectation mismatch for row")
	assert.Equal(t, config.Total, int32(87), "expectation mismatch for total")
	assert.Equal(t, config.Ref, int64(6152), "expectation mismatch for reference")
	assert.Equal(t, config.Avg, float32(61.45), "expectation mismatch for average")
	assert.Equal(t, config.Pct, 86.74, "expectation mismatch for percent")
	assert.Equal(t, config.IsDefault, false, "expectation mismatch for is_default")
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
	errs := Parse(config)

	assert.Equal(t, len(errs), 1,"unexpected error while parsing")
	assert.Equal(t, errs[0], errors.New("env: expected ptr but got struct"), "wrong error message")
}

func TestEnvParserWithNonStructValue(t *testing.T) {
	config := "abcd"
	errs := Parse(&config)

	assert.Equal(t, len(errs), 1,"unexpected error while parsing")
	assert.Equal(t, errs[0], errors.New("env: expected struct but got string"), "wrong error message")
}

func TestEnvParserWithUnsupportedTypes(t *testing.T) {
	_ = os.Setenv("DEFAULT", "321")
	defer os.Clearenv()

	type args struct {
		IsDefault uint `env:"DEFAULT"`
	}
	config := args{}
	errs := Parse(&config)

	assert.Equal(t, len(errs), 1, "unexpected error while parsing")
	assert.Equal(t, errs[0], errors.New("env: uint is not a supported type"), "wrong error message")
}

func TestEnvParserWithWrongValues(t *testing.T) {
	_ = os.Setenv("PORT", "5a")
	_ = os.Setenv("AVERAGE", "5a.23")
	_ = os.Setenv("IS_DEFAULT", "not false")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT"`
		Avg float64 `env:"AVERAGE"`
		IsDefault bool `env:"IS_DEFAULT"`
	}
	config := args{}
	errs := Parse(&config)

	assert.Equal(t, len(errs), 3, "unexpected error while parsing")
	assert.Equal(t, errs[0], errors.New("env: strconv.ParseInt: parsing \"5a\": invalid syntax"), "wrong error message")
	assert.Equal(t, errs[1], errors.New("env: strconv.ParseFloat: parsing \"5a.23\": invalid syntax"), "wrong error message")
	assert.Equal(t, errs[2], errors.New("env: strconv.ParseBool: parsing \"not false\": invalid syntax"), "wrong error message")
}

func TestEnvParserWithoutEnvironmentVariable(t *testing.T) {
	_ = os.Setenv("PORT", "5000")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT"`
		Name string `env:"Name"`
		ID int8 `env:"ID"`
		Row int16 `env:"ROW"`
		Total int32 `env:"TOTAL"`
		Ref int64 `env:"REFERENCE"`
		Avg float32 `env:"AVERAGE"`
		Pct float64 `env:"PERCENT"`
		IsDefault bool `env:"IS_DEFAULT"`
	}
	config := args{}
	errs := Parse(&config)

	assert.Equal(t, len(errs), 0,"unexpected error while parsing")
	assert.Equal(t, config.Port, 5000, "expectation mismatch for port")
	assert.Equal(t, config.Name, "", "expectation mismatch for name")
	assert.Equal(t, config.ID, int8(0), "expectation mismatch for id")
	assert.Equal(t, config.Row, int16(0), "expectation mismatch for row")
	assert.Equal(t, config.Total, int32(0), "expectation mismatch for total")
	assert.Equal(t, config.Ref, int64(0), "expectation mismatch for reference")
	assert.Equal(t, config.Avg, float32(0), "expectation mismatch for average")
	assert.Equal(t, config.Pct, float64(0), "expectation mismatch for percent")
	assert.Equal(t, config.IsDefault, false, "expectation mismatch for is default")
}

func TestEnvParserWithDefaultValueProvided(t *testing.T) {
	_ = os.Setenv("ID", "14")
	defer os.Clearenv()

	type args struct {
		Port int `env:"PORT" default:"5000"`
		Name string `env:"Name"`
		ID int `env:"ID"`
		Avg int `env:"AVERAGE"`
		Place string `env:"PLACE" default:"Bangalore"`
	}
	config := args{}
	errs := Parse(&config)

	assert.Equal(t, len(errs), 0,"unexpected error while parsing")
	assert.Equal(t, config.Port, 5000, "expectation mismatch for port")
	assert.Equal(t, config.Name, "", "expectation mismatch for name")
	assert.Equal(t, config.ID, 14, "expectation mismatch for id")
	assert.Equal(t, config.Avg, 0, "expectation mismatch for average")
	assert.Equal(t, config.Place, "Bangalore", "expectation mismatch for place")
}
