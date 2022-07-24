package neisgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeisClient_GetSchoolInfo(t *testing.T) {
	client := NewNeisClient("")
	schools, err := client.GetSchoolInfo("", "", "선린인터넷고등학교", "", "", "")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, schools[0].Name, "선린인터넷고등학교")
}
