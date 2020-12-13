package querybuilder

import (
	"database/sql"
	"errors"
	"reflect"
)

// ScanStruct converts sql row.Scan into passed struct
func ScanStruct(str interface{}, r *sql.Rows) error {

	// Get the underlying element from passed interface
	v := reflect.ValueOf(str).Elem()

	// Check if passed interface is pointer to struct
	if v.Kind() != reflect.Struct {
		return errors.New("Passed interface is not a pointer to struct")
	}

	// Make slice to hold fields of structs
	fields := make([]interface{}, v.NumField())
	for kk := 0; kk < v.NumField(); kk++ {
		fields[kk] = v.Field(kk).Addr().Interface()
	}

	// Execute rows.Scan
	err := r.Scan(fields...)
	if err != nil {
		return err
	}
	return nil
}
