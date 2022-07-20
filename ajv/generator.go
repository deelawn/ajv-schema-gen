package ajv

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"reflect"
)

func Generate(r io.Reader) (schema Schema, err error) {

	rawJSON, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	var mappedValues interface{}
	{
		mappedValues = map[string]interface{}{}
		err = json.Unmarshal(rawJSON, &mappedValues)
	}

	if err != nil {
		mappedValues = []interface{}{}
		if err2 := json.Unmarshal(rawJSON, &mappedValues); err2 != nil {
			err = errors.New("input must be a json object or array")
			return
		}
	}

	return resolveSchema(mappedValues)
}

func resolveSchema(value interface{}) (schema Schema, err error) {

	switch value.(type) {

	case string:
		schema.Typ = "string"

	case float64:
		schema.Typ = "number"

	case bool:
		schema.Typ = "boolean"

	default:
		vtyp, vval := reflect.TypeOf(value), reflect.ValueOf(value)
		if vtyp == nil {
			schema = Schema{
				Typ:      "null",
				Nullable: true,
			}
			break
		}

		switch vtyp.Kind() {

		case reflect.Map:
			schema.Typ = "object"
			if len(vval.MapKeys()) == 0 {
				break
			}

			schema.Properties = make(map[string]Schema)

			iter := vval.MapRange()
			for iter.Next() {
				key := iter.Key()
				value := iter.Value()

				valueSchema, err := resolveSchema(value.Interface())
				if err != nil {
					return schema, err
				}

				schema.Properties[key.String()] = valueSchema
				schema.Required = append(schema.Required, key.String())
			}
		case reflect.Array, reflect.Slice:
			schema.Typ = "array"
			if vval.Len() == 0 {
				schema.Nullable = true
				break
			}

			// We only care about the schema of the first item in the array for now.
			itemSchema, err := resolveSchema(vval.Index(0).Interface())
			if err != nil {
				return schema, err
			}

			schema.Items = &itemSchema
		default:
			err = errors.New("unknown type" + vtyp.String())
		}
	}

	return
}
