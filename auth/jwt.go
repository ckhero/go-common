package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ckhero/go-common/util/idcreator/snowflake"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

func NewUserJwtToken(userId uint64, info map[string]interface{}, secretKey string) (string, int64, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["jti"] = strconv.FormatUint(snowflake.NextID(), 10)
	claims["iss"] = "wanxin"
	claims["sub"] = strconv.FormatUint(userId, 10)
	expireAt := time.Now().AddDate(0, 0, 30).Unix()
	claims["exp"] = expireAt
	if info != nil {
		infoBytes, err := json.Marshal(info)
		if err != nil {
			return "", 0, err
		}
		claims["info"] = string(infoBytes)
	}
	token.Claims = claims
	signedToken, err := token.SignedString([]byte(secretKey))
	return signedToken, expireAt, err
}

func ResolveJWTToken(tokenString string, secretKey string, log *logrus.Entry) (userId uint64, info map[string]interface{}, err error) {
	
	if tokenString == "" {
		return 0, nil, errors.New("token invalid")

	}

	log = log.WithField("token", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			log.WithField("signMethod", token.Method).
				Warn("invalid token sign method")
			return nil, fmt.Errorf("unexpected sign method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {

		validationErr, ok := err.(*jwt.ValidationError)
		if ok && (validationErr.Errors&jwt.ValidationErrorExpired != 0) &&
			(validationErr.Errors&jwt.ValidationErrorSignatureInvalid == 0) {

			return 0, nil, errors.New("token expired")
		}

		log.WithError(err).
			Warn("token is invalid")
		return 0, nil, errors.New("token invalid")
	}

	if !token.Valid {
		log.Warn("token is invalid")
		return 0, nil, errors.New("token invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Warn("convert claims failed")
		return 0, nil, errors.New("token invalid")
	}

	subject, ok := claims["sub"]

	if !ok {
		log.Warn("subject not exist")
		return 0, nil, errors.New("token invalid")
	}

	subjectStr, ok := subject.(string)

	if !ok {
		log.Warn("subject not a string")
		return 0, nil, errors.New("token invalid")
	}

	userId, err = strconv.ParseUint(subjectStr, 10, 64)

	if err != nil {
		log.WithError(err).Warn("parse userId error")
		return 0, nil, errors.New("token invalid")
	}

	if infoJSON, ok := claims["info"]; ok {

		if infoJSONString, ok := infoJSON.(string); ok {

			err := json.Unmarshal([]byte(infoJSONString), &info)
			if err != nil {
				log.WithField("infoJSON", infoJSONString).WithError(err).Warn("unmarshal info json error")
			}
		}
	}

	return userId, info, nil
}


