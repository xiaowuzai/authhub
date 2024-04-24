package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"git.corp.zgcszkw.com/authhub/token"
)

func addAppKeyAuthorization(
	t *testing.T,
	request *http.Request,
	tokenValidator *token.AppKeyValidator,
	authorizationType string,
	requestId string,
	created string,
	key string,
) {

	appKey := token.NewAppKey(requestId, created, key)
	// 生成token
	token, err := tokenValidator.CreateToken(appKey)
	require.NoError(t, err)

	// 添加Authorization头
	request.Header.Set(authorizationHeaderKey, fmt.Sprintf("%s %s", authorizationType, token))
	request.Header.Set(headerRequestId, requestId)
	request.Header.Set(headerCreated, created)
	request.Header.Set(headerAppKey, key)
}

func TestAuthMiddlewareAppKey(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenValidator *token.AppKeyValidator)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenValidator *token.AppKeyValidator) {
				created := time.Now().Format(time.RFC3339)
				addAppKeyAuthorization(t, request, tokenValidator, authorizationTypeAppKey,
					"requestIdOK", created, "test_key")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenValidator *token.AppKeyValidator) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "NoTypeAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenValidator *token.AppKeyValidator) {
				created := time.Now().UTC().Format(time.RFC3339)
				addAppKeyAuthorization(t, request, tokenValidator, "",
					"requestIdNoTypeAuthorization", created, "test_key")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenValidator *token.AppKeyValidator) {
				created := time.Now().UTC().Format(time.RFC3339)
				addAppKeyAuthorization(t, request, tokenValidator, "UnsupportedAuthorization",
					"requestIdUnsupportedAuthorization", created, "test_key")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request, tokenValidator *token.AppKeyValidator) {
				created := time.Now().UTC().AddDate(0, 0, -1).Format(time.RFC3339)
				addAppKeyAuthorization(t, request, tokenValidator, authorizationTypeAppKey,
					"requestIdExpiredToken", created, "test_key")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)

			// 注册一个测试路由
			authPath := "/v1/auth"
			router := gin.Default()
			router.GET(
				authPath,
				authMiddlewareAppKey(server.appKeyValidator),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)
			server.setRouter(router)

			recoder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.appKeyValidator)

			server.httpServer.Handler.ServeHTTP(recoder, request)
			tc.checkResponse(t, recoder)
		})
	}
}
