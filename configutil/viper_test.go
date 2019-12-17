package configutil_test

import (
	"testing"

	"go.stevenxie.me/gopkg/configutil"
)

func TestUnmarshalViper(t *testing.T) {
	viper := configutil.NewViper("config", "namespace")
	viper.Set("test", "hello?")

	var config struct {
		Test string
	}
	if err := configutil.UnmarshalViper(viper, &config); err != nil {
		t.Fatalf("Failed to unmarshal Viper: %v", err)
	}

	{
		const expected = "hello?"
		if value := config.Test; value != "hello?" {
			t.Errorf(
				"Expected config.Test to equal '%s', but got '%s'",
				expected, value,
			)
		}
	}

	t.Logf("config = %+v\n", config)
}
