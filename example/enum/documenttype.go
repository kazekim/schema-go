package enum

import (
	"reflect"
)

type DocumentType string

const (
	DOCUMENTTYPE_NONE      DocumentType = "none"
	DOCUMENTTYPE_CUSTOM                 = "custom"
	DOCUMENTTYPE_TAX                    = "tax"
	DOCUMENTTYPE_COMPANYID              = "companyID"
)


func DocumentTypeFromKey(key string) DocumentType {

	var documentType DocumentType
	rv := reflect.Indirect(reflect.ValueOf(documentType))
	rv.SetString(key)

	return rv.Interface().(DocumentType)
}

func ParseDocumentType(d DocumentType, key string) DocumentType {
	return d.ParseKey(&key)
}

func ParseDocumentTypeReflectValue(value string, rv reflect.Value, typeName string) error {

	rv.SetString(value)

	return nil
}

func (d DocumentType) String() string {
	return string(d)
}

func (d DocumentType) SetIndex(i int64) {

	v := reflect.ValueOf(d)
	v.SetInt(i)
}

func (d DocumentType) ParseKey(key *string) DocumentType {

	documentType := DOCUMENTTYPE_NONE

	if key != nil {
		documentType = DocumentTypeFromKey(*key)
	}
	return documentType
}
