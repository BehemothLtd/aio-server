package auths

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/constants"
	jsonwebtoken "aio-server/pkg/jsonWebToken"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", exceptions.NewUnauthorizedError("Bad header value given")

	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", exceptions.NewUnauthorizedError("Incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (uid int32, err error) {
	var userClaim models.UserClaims

	decodedErr := jsonwebtoken.DecodeJwtToken(jwtToken, &userClaim)

	if decodedErr != nil {
		return 0, exceptions.NewUnauthorizedError("Bad jwt token")
	}

	return userClaim.Sub, nil
}

func JwtTokenCheck(c *gin.Context) {
	jwtToken, tokenErr := extractBearerToken(c.GetHeader(constants.AuthorizationHeader))

	if tokenErr != nil {
		return
	}

	uid, parseError := parseToken(jwtToken)

	if parseError == nil {
		var user models.User

		result := database.Db.Table("users").First(&user, uid)

		if result.Error == nil {
			c.Set(constants.ContextCurrentUser, user)
		}
	}

	c.Next()
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constants.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
