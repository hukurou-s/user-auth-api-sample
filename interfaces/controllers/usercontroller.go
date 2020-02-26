package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/hukurou-s/user-auth-api-sample/crypt"
	"github.com/hukurou-s/user-auth-api-sample/domain"
	"github.com/hukurou-s/user-auth-api-sample/interfaces/database"
	"github.com/hukurou-s/user-auth-api-sample/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) UserProfile(c Context) error {
	// 送られてきたJWTからnameの値を取得
	jwtUser := c.Get("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	userName := claims["name"].(string)

	user, err := controller.Interactor.UserByName(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Login(c Context) error {
	// Postされたパラメータを取得
	loginParams := new(domain.LoginParams)
	if err := c.Bind(loginParams); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := controller.Interactor.UserByName(loginParams.Name)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	if !compareHashedPassword(user.Password, loginParams.Password) {
		return c.JSON(http.StatusUnauthorized, struct {
			Status string `json:"status"`
		}{
			Status: "fail",
		})

	}
	//
	//// 秘密鍵を読み込み
	//keyData, err := ioutil.ReadFile("./rsa/id_rsa")
	//if err != nil {
	//	panic(err)
	//}
	//key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	//if err != nil {
	//	panic(err)
	//}
	//
	key := crypt.NewPrivateKey()
	token := jwt.New(jwt.SigningMethodRS256)
	// claimの設定
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}{
		Status: "success",
		Token:  t,
	})
}

func toHashPassword(pass string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(converted)
}

func compareHashedPassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}
	return false
}
