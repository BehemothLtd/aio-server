package main

import (
	"aio-server/database"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/initializers"
	jsonwebtoken "aio-server/pkg/jsonWebToken"
	"aio-server/pkg/logger"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (uid int32, err error) {
	var userClaim models.UserClaims

	decodedErr := jsonwebtoken.DecodeJwtToken(jwtToken, &userClaim)

	if decodedErr != nil {
		return 0, errors.New("bad jwt token")
	}

	return userClaim.Sub, nil
}

func jwtTokenCheck(c *gin.Context) {
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

func main() {
	initializers.LoadEnv()

	// Load DB
	db := database.InitDb()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(logger.Logger(logrus.New()), gin.Recovery())

	r.Use(jwtTokenCheck, GinContextToContextMiddleware()).POST("/graphql", initializers.GqlHandler(db))

	r.Run()
}
