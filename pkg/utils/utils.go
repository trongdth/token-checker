package utils

import (
	"encoding/json"

	errs "github.com/trongdth/token-checker/m/v2/app/errors"

	"go.mongodb.org/mongo-driver/bson"
)

// ToBson transform interface to a bson Model
func ToBson(v interface{}) (mod *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, errs.ErrorWithMessage(errs.ErrMarshal, err.Error())
	}

	err = bson.Unmarshal(data, &mod)
	if err != nil {
		return nil, errs.ErrorWithMessage(errs.ErrUnmarshal, err.Error())
	}
	return mod, nil
}

// Copy value as json
func Copy(toValue interface{}, fromValue interface{}) error {
	data, err := json.Marshal(fromValue)
	if err != nil {
		return errs.ErrorWithMessage(errs.ErrMarshal, err.Error())
	}

	err = json.Unmarshal(data, toValue)
	if err != nil {
		return errs.ErrorWithMessage(errs.ErrUnmarshal, err.Error())
	}

	return nil
}
