package token

import (
	"testing"
	"time"

	"git.corp.zgcszkw.com/authhub/util"
	"github.com/stretchr/testify/assert"
)

func TestAppKeyCreate(t *testing.T) {
	expired := util.RandomDuration()
	key, secret := util.RandomString(10), util.RandomString(10)
	keyMap := map[string]string{
		key: secret,
	}

	akValidator, err := NewAppKeyValidator(expired, keyMap)
	assert.NoError(t, err)

	requestId := util.RandomString(10)
	created := time.Now().Format(APP_KEY_TIME_FORMAT)
	appKey := NewAppKey(requestId, created, key)

	signature, err := akValidator.CreateToken(appKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, signature)

	err = akValidator.VerifyToken(appKey, signature)
	assert.NoError(t, err)
}
