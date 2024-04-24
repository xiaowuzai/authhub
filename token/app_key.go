package token

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"errors"
)

const (
	APP_KEY_TIME_FORMAT = time.RFC3339
)

var (
	ErrAppKeyFormat   = errors.New("app key format error") // 格式错误
	ErrAppKeyNotExist = errors.New("app key not exist")    //  不存在
	// errAppKeyExpired  = errors.New("app key is expired")   // 过期了
	ErrTokenNotMatch = errors.New("token not match") // token 不匹配

)

type AppKeyValidator struct {
	expired time.Duration     // 过期时间
	keyMap  map[string]string // 存储 key secret
}

// 过期时间
func NewAppKeyValidator(expired time.Duration, keyMap map[string]string) (*AppKeyValidator, error) {
	akm := &AppKeyValidator{
		expired: expired,
		keyMap:  keyMap,
	}

	return akm, nil
}

// 创建token
func (akm *AppKeyValidator) CreateToken(appKey *AppKey) (string, error) {
	secret, has := akm.getSecret(appKey)
	if !has {
		return "", ErrAppKeyNotExist
	}
	return akm.createToken(appKey, secret)
}

func (akm *AppKeyValidator) createToken(appKey *AppKey, secret string) (string, error) {
	str := appKey.formatString(secret)
	log.Println("GenSignature str: ", str)
	token := genSignature(str)
	return token, nil
}

func (akm *AppKeyValidator) VerifyToken(appKey *AppKey, token string) error {
	// 获取 secret
	secret, has := akm.getSecret(appKey)
	if !has {
		return ErrAppKeyNotExist
	}

	// 验证 app key 的参数
	if err := appKey.validateCreated(); err != nil {
		return err
	}
	// 验证是否过期
	if err := akm.validateExpired(appKey); err != nil {
		return err
	}

	newToken, err := akm.createToken(appKey, secret)
	if err != nil {
		return err
	}

	if token != newToken {
		return ErrTokenNotMatch
	}

	return nil
}

// 返回 app secret
func (akm *AppKeyValidator) getSecret(appKey *AppKey) (string, bool) {
	secret, has := akm.keyMap[appKey.key]
	return secret, has
}

// 验证过期时间。使用 UTC 时间
func (akm *AppKeyValidator) validateExpired(appKey *AppKey) error {
	created, err := appKey.parseTime(appKey.created)
	if err != nil {
		return err
	}
	log.Println("created", created)

	since := time.Now().UTC().Sub(created)
	if since > akm.expired {
		log.Printf("since %v, expired %v\n", since, akm.expired)
		return ErrExpiredToken
	}
	return nil
}

func (akm *AppKeyValidator) newAppKeyByToken(token string) (*AppKey, error) {
	fields := splitToken(token)
	if len(fields) != 4 {
		return nil, ErrAppKeyFormat
	}

	appKey := NewAppKey(fields[0], fields[1], fields[2])
	return appKey, nil
}

type AppKey struct {
	requestId string // 请求头中的 id
	created   string // 创建时间
	key       string // 分配的 key
}

func NewAppKey(
	requestId string, // 请求头中的 id
	created string, // 创建时间
	key string, // 分配的 key
) *AppKey {
	return &AppKey{
		requestId: requestId,
		created:   created,
		key:       key,
	}
}

// formatString格式化签名字符串
func (ak *AppKey) formatString(appSecret string) string {
	return ak.requestId + ak.created + appSecret
}

func splitToken(token string) []string {
	return strings.Split(token, "_")
}

// 验证时间字符串
func (ak *AppKey) validateCreated() error {
	_, err := ak.parseTime(ak.created)
	if err != nil {
		return fmt.Errorf("invalid created time: %w", err)
	}
	return nil
}

// 要求时间格式 "2006-01-02T15:04:05Z07:00"
func (ak *AppKey) parseTime(timeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeStr)
}

// genSignature 生成签名
func genSignature(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hs := hash.Sum(nil)
	base64String := base64.StdEncoding.EncodeToString(hs)
	return url.QueryEscape(base64String)
}
