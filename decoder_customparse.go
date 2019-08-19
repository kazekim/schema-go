/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package schema

import (
	"fmt"
	"reflect"
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
	pm := *d.parserMap
	name := fmt.Sprint(v.Type())

	_, ok := pm[name]
	return ok
}