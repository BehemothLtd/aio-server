package gql

import (
	"aio-server/models"
	"aio-server/pkg/constants"
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resolver struct {
	Db *gorm.DB
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(constants.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func AuthUserFromCtx(ctx context.Context) (models.User, error) {
	gc, _ := GinContextFromContext(ctx)

	currentUser := gc.Value(constants.ContextCurrentUser)

	if currentUser == nil {
		return models.User{}, errors.New("unauthorized")
	}

	return currentUser.(models.User), nil
}
