package auth_test

import (
	"fmt"
	"testing"

	"github.com/muchrief/go_pijar/src/auth"
	"github.com/muchrief/go_pijar/src/model"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessKey(t *testing.T) {
	data := &model.Auth{
		Id:       "qweqweasdasdasd",
		Username: "jhkgadsgjhadsg@gmail.com",
	}

	token, err := auth.GenerateAccessToken(data)
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)

	fmt.Println(token)
}
