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
		Role:     model.PUBLIC,
	}

	token, err := auth.GenerateAccessToken(data)
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)

	fmt.Println(token)
}
