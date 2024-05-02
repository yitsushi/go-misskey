package core

import (
	"reflect"
	"strings"
)

const numberOfPartsKeyValue = 2

type multipartField struct {
	Type      string
	Name      string
	Value     []byte
	Ref       string
	OmitEmpty bool
}

func parseMultipartFields(r BaseRequest) map[string]multipartField {
	fields := map[string]multipartField{}

	value := reflect.ValueOf(r)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for index := range value.NumField() {
		tag := value.Type().Field(index).Tag.Get("multipart")
		if tag == "" || tag == "-" {
			continue
		}

		field := parseTag(tag)
		if field.OmitEmpty && value.Field(index).IsZero() {
			continue
		}

		if field.Name == "" {
			field.Name = value.Type().Field(index).Name
		}

		field.Value = parseValueToBytes(value.Field(index))

		fields[field.Name] = field
	}

	return fields
}

func parseTag(tag string) multipartField {
	field := multipartField{}

	for _, part := range strings.Split(tag, ",") {
		if !strings.Contains(part, "=") {
			if part == "omitempty" {
				field.OmitEmpty = true
			} else {
				field.Name = part
			}

			continue
		}

		parts := strings.SplitN(part, "=", numberOfPartsKeyValue)
		key, value := parts[0], parts[1]

		switch key {
		case "type":
			field.Type = value
		case "ref":
			field.Ref = value
		}
	}

	return field
}

func parseValueToBytes(item reflect.Value) []byte {
	switch item.Type().Name() {
	case "string":
		return []byte(item.String())
	case "bool":
		if item.Bool() {
			return []byte("true")
		}

		return []byte("false")
	default:
		return item.Bytes()
	}
}
