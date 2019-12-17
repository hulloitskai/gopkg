package mapstructx_test

import (
	"testing"

	"github.com/cockroachdb/errors"
	"go.stevenxie.me/gopkg/mapstructx"
	"go.stevenxie.me/gopkg/zero"
)

type wrapper struct {
	internal string
}

var _ mapstructx.Unmarshaler = (*wrapper)(nil)

func (m *wrapper) UnmarshalMap(v zero.Interface) error {
	s, ok := v.(string)
	if !ok {
		return errors.New("wrapper must be of type string")
	}
	m.internal = s
	return nil
}

func TestUnmarshaler(t *testing.T) {
	var (
		data struct {
			Wrapped *wrapper `mapstructure:"wrapped"`
		}
		input = map[string]string{
			"wrapped": "hello",
		}
	)
	if err := mapstructx.Unmarshal(input, &data); err != nil {
		t.Fatalf("Failed to unmarshal: %+v", err)
	}
	{
		const expected = "hello"
		if value := data.Wrapped.internal; value != expected {
			t.Errorf(
				"Expected data.Wrapped.internal to be '%s', but got '%s'",
				expected, value,
			)
		}
	}
	t.Logf("data = %+v", &data)
}
