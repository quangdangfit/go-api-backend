package utils

import (
	"encoding/json"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/quangdangfit/go-backend/pkg/errors"
)

// ContextKey type
type ContextKey string

// Constants
const (
	Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// TimeToString convert time to string with RFC3339Nano format
func TimeToString(t time.Time) string {
	return t.Local().Format(time.RFC3339Nano)
}

// ToBson transform interface to a bson Model
func ToBson(v interface{}) (mod *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, errors.New(errors.ECMarshal, err.Error(), "utils.ToBson")
	}

	err = bson.Unmarshal(data, &mod)
	if err != nil {
		return nil, errors.New(errors.ECUnmarshal, err.Error(), "utils.ToBson")
	}
	return mod, nil
}

// ToMap transform interface to a bson Model
func ToMap(v interface{}) (mod map[string]interface{}, err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.New(errors.ECMarshal, err.Error(), "utils.ToMap")
	}

	err = json.Unmarshal(data, &mod)
	if err != nil {
		return nil, errors.New(errors.ECUnmarshal, err.Error(), "utils.ToMap")
	}
	return mod, nil
}

// Copy copy value as json
func Copy(toValue interface{}, fromValue interface{}) error {
	data, err := json.Marshal(fromValue)
	if err != nil {
		return errors.New(errors.ECMarshal, err.Error(), "utils.Copy")
	}

	err = json.Unmarshal(data, toValue)
	if err != nil {
		return errors.New(errors.ECUnmarshal, err.Error(), "utils.Copy")
	}

	return nil
}

// RandomString create new random string have length n
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[seededRand.Intn(len(Charset))]
	}
	return string(b)
}

// ToJSONString parse object to string json
func ToJSONString(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}

	return string(b)
}
