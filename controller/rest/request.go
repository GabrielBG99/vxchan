package rest

import (
	"fmt"
	"net/http"
	"reflect"
)

type logResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// NewLogResponseWriter returns a response writer that saves data for future logging
func NewLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK}
}

func (w *logResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// DecodeMultipartForm decodes the request body and stores it in the value pointed to by z.
func DecodeMultipartForm(r *http.Request, z interface{}) error {
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		return err
	}

	fields := make(map[string]string)
	elm := reflect.ValueOf(z).Elem()
	for i := 0; i < elm.NumField(); i++ {
		field := elm.Type().Field(i)

		name := field.Name
		tag := field.Tag.Get("multipart")
		if tag == "" {
			tag = name
		}

		fields[tag] = name
	}

	for k := range r.MultipartForm.Value {
		fieldName, ok := fields[k]
		if !ok {
			continue
		}

		v := r.FormValue(k)

		field := elm.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			expectedType := field.Type()

			formValue := reflect.ValueOf(v)
			if !formValue.CanConvert(expectedType) {
				return fmt.Errorf("The field \"%s\" is not of type \"%s\"", fieldName, expectedType.String())
			}

			field.Set(formValue.Convert(expectedType))
		}
	}

	for k := range r.MultipartForm.File {
		fieldName, ok := fields[k]
		if !ok {
			continue
		}

		f, _, err := r.FormFile(k)
		if err != nil {
			return err
		}

		field := elm.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(f))
		}
	}

	return nil
}
