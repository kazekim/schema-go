/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package schema

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func (d *Decoder) CustomParser(parserMap CustomParserMap) {
	d.parserMap = &parserMap
}

func (d *Decoder) parseCustomParser(value string, v reflect.Value) error {

	typeName := fmt.Sprint(v.Type())
	if d.parserMap == nil {
		return nil
	}

	m := *d.parserMap
	f := m[typeName]

	if f == nil {
		err := fmt.Errorf("unsupport field type %s", typeName)
		return err
	}

	err := f(value, v, typeName)
	return err
}

func (d *Decoder) hasCustomParser(v reflect.Value) bool {

	if d.parserMap != nil {
		pm := *d.parserMap
		name := fmt.Sprint(v.Type())

		_, ok := pm[name]
		return ok
	}

	return false
}

func (d *Decoder) findMatchRecursiveStructType(t reflect.Type, path string) (reflect.Type, *string, error) {
	struc := d.cache.get(t)

	if struc == nil {
		// unexpect, cache.get never return nil
		return nil, nil, errors.New("cache fail")
	}

	for _, f := range struc.fields {
		if f.typ.Kind() == reflect.Struct {
			if f.isSubStructParse {
				for i := 0; i < f.typ.NumField() ; i++ {
					sf := f.typ.Field(i)
					if strings.EqualFold(snakeCase(sf.Name), path) {
						return f.typ, &f.name, nil
					}
				}
			}
		}
	}

	return nil, nil, errors.New("no recursive struct field found")
}