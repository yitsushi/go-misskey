package core

import (
	"reflect"
	"strings"
)

type multipartField struct {
	Type  string
	Name  string
	Value []byte
	Ref   string
}

func parseMultipartFields(r BaseRequest) map[string]multipartField {
	fields := map[string]multipartField{}

	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("multipart")
		if tag == "" || tag == "-" {
			continue
		}

		field := parseTag(tag)
		if field.Name == "" {
			field.Name = v.Type().Field(i).Name
		}

		field.Value = parseValueToBytes(v.Field(i))

		fields[field.Name] = field
	}

	return fields
}

func parseTag(tag string) multipartField {
	field := multipartField{}

	for _, part := range strings.Split(tag, ",") {
		if !strings.Contains(part, "=") {
			field.Name = part

			continue
		}

		parts := strings.SplitN(part, "=", 2)
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
