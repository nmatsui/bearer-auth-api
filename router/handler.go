package router

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/nmatsui/bearer-auth-api/token"
)

const AUTH_HEADER = "authorization"
const BEARER_RE = `(?i)^bearer (.+)$`

type Handler struct {
	Engine *gin.Engine
}

func NewHandler() *Handler {
	engine := gin.Default()
	holder := token.NewHolder()

	tokenRe := regexp.MustCompile(BEARER_RE)

	engine.NoRoute(func(context *gin.Context) {
		authHeader := context.Request.Header.Get(AUTH_HEADER)
		if len(authHeader) == 0 {
			authHeaderMissing(context)
		} else {
			matches := tokenRe.FindAllStringSubmatch(authHeader, -1)
			if len(matches) == 0 || !holder.HasToken(matches[0][1]) {
				tokenMissmatch(context)
			} else if !matchPath(context, holder.GetAllowedPaths(matches[0][1])) {
				pathNotAllowed(context)
			} else {
				statusOK(context)
			}
		}
	})

	router := &Handler{
		Engine: engine,
	}
	return router
}

func (router *Handler) Run(port string) {
	router.Engine.Run(port)
}

func matchPath(context *gin.Context, allowedPaths []*regexp.Regexp) bool {
	path := context.Request.URL.Path
	for _, allowedPath := range allowedPaths {
		if allowedPath.MatchString(path) {
			return true
		}
	}
	return false
}

func authHeaderMissing(context *gin.Context) {
	context.Writer.Header().Set("WWW-Authenticate", "Bearer realm=\"token_required\"")
	context.JSON(http.StatusUnauthorized, gin.H{
		"authorized": false,
		"error":      "missing Header: " + AUTH_HEADER,
	})
}

func tokenMissmatch(context *gin.Context) {
	context.Writer.Header().Set("WWW-Authenticate", "Bearer realm=\"token_required\" error=\"invalid_token\"")
	context.JSON(http.StatusUnauthorized, gin.H{
		"authorized": false,
		"error":      "token mismatch",
	})
}

func pathNotAllowed(context *gin.Context) {
	context.Writer.Header().Set("WWW-Authenticate", "Bearer realm=\"token_required\" error=\"not_allowed\"")
	context.JSON(http.StatusForbidden, gin.H{
		"authorized": false,
		"error":      "not allowd",
	})
}

func statusOK(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"authorized": true,
	})
}
