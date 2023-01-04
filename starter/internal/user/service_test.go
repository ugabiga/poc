package user_test

import (
	"context"
	"log"
	"starter/internal/common"
	"starter/internal/user"
	"testing"
)

func initService() *user.Service {
	entClient := common.NewEntClient(common.DBInfo())
	s := user.NewService(entClient)
	return s
}

func TestService_CreateUser(t *testing.T) {
	ctx := context.Background()
	s := initService()
	createUser, err := s.CreateUser(ctx)
	if err != nil {
		return
	}

	log.Println("createUser", *createUser)
}
