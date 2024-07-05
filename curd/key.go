package curd

import (
	"github.com/god-jason/bucket/config"
	"github.com/god-jason/bucket/web"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"reflect"
)

func GenerateXID[T any]() func(data *T) error {
	return func(data *T) error {
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("Id")
		//使用UUId作为Id
		//field.IsZero() 如果为空串时，生成UUID
		if field.Len() == 0 {
			key := xid.New().String()
			field.SetString(key)
		}
		return nil
	}
}

func GenerateKSUID[T any]() func(data *T) error {
	return func(data *T) error {
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("Id")
		//使用UUId作为Id
		//field.IsZero() 如果为空串时，生成UUID
		if field.Len() == 0 {
			key := ksuid.New().String()
			field.SetString(key)
		}
		return nil
	}
}

func GenerateID[T any]() func(data *T) error {
	return func(data *T) error {
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("Id")
		//使用UUId作为Id
		//field.IsZero() 如果为空串时，生成UUID
		if field.Len() == 0 {
			var key string
			switch config.GetString(web.MODULE, "id") {
			case "ksuid":
				key = ksuid.New().String()
			case "xid":
				key = xid.New().String()
			default:
				key = xid.New().String()
			}
			field.SetString(key)
		}
		return nil
	}
}
