package restpc

import "testing"
import "github.com/stretchr/testify/assert"

func TestFromEmpty(t *testing.T) {
	{
		value, err := FromEmpty.GetString(nil, "foo")
		assert.NoError(t, err)
		assert.Equal(t, "", *value)
	}
	{
		value, err := FromEmpty.GetStringList(nil, "foo")
		assert.NoError(t, err)
		assert.Equal(t, []string{}, value)
	}
	{
		value, err := FromEmpty.GetInt(nil, "foo")
		assert.NoError(t, err)
		assert.Equal(t, 0, *value)
	}
	{
		value, err := FromEmpty.GetFloat(nil, "foo")
		assert.NoError(t, err)
		assert.Equal(t, 0.0, *value)
	}
	{
		value, err := FromEmpty.GetBool(nil, "foo")
		assert.NoError(t, err)
		assert.Equal(t, false, *value)
	}
	{
		value, err := FromEmpty.GetTime(nil, "foo")
		assert.NoError(t, err)
		assert.Equal(t, int64(-62135596800), value.Unix())
	}
}
