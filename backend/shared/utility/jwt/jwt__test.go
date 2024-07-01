package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var userData = map[string]interface{}{
	"userId": "8298983-23233-32323-32323",
}

func TestGenerateJWT(t *testing.T) {
	secretKey = "aa8b73-9fa0430c-53b486d7-c8c438048c"
	audience = "TaskManagementAudience"
	issuer = "TaskManagementIssuer"
	firstToken, _ := GenerateJWT(userData)
	time.Sleep(1 * time.Second)
	secondToken, _ := GenerateJWT(userData)
	t.Run("must generate a different token on every call", func(t *testing.T) {
		assert.NotEqual(t, firstToken, secondToken)
	})
}

func TestParse(t *testing.T) {
	secretKey = "aa8b73-9fa0430c-53b486d7-c8c438048c"
	audience = "TaskManagementAudience"
	issuer = "TaskManagementIssuer"
	token, _ := GenerateJWT(userData)
	data, _ := Parse(token)
	t.Run("must parse a previous generated token", func(t *testing.T) {
		assert.Equal(t, userData["userId"], data["uid"])
	})
}
