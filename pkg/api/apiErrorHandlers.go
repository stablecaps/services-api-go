package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil/header"
)

////////////////////////////////////////////////////////
////////////////////////////////////////////////////////
type malformedRequest struct {
    status int
    msg    string
}

func (mr *malformedRequest) Error() string {
    return mr.msg
}

func errorCheckJSONBody(writer http.ResponseWriter, req *http.Request, dst interface{}) error {

    // check correct content-type header is present
    if req.Header.Get("Content-Type") != "" {
        value, _ := header.ParseValueAndParams(req.Header, "Content-Type")
        if value != "application/json" {
            msg := "Content-Type header is not application/json"
            return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
        }
    }


    // esure max 1 MB max payload in request body. A larget payload results in Decode()
    // returning a "http: request body too large" error.
    req.Body = http.MaxBytesReader(writer, req.Body, 1048576)

    // Setup the decoder and call the DisallowUnknownFields() method on it.
    // This will cause Decode() to return a "json: unknown field ..." error
    // if it encounters any extra unexpected fields in the JSON. Strictly
    // speaking, it returns an error for "keys which do not match any
    // non-ignored, exported fields in the destination".
    dec := json.NewDecoder(req.Body)
    dec.DisallowUnknownFields()

    err := dec.Decode(&dst)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError

        switch {
        // Catch any syntax errors in the JSON and send an error message
        // which interpolates the location of the problem to make it
        // easier for the client to fix.
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}

        // In some circumstances Decode() may also return an
        // io.ErrUnexpectedEOF error for syntax errors in the JSON.
        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprintf("Request body contains badly-formed JSON")
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}

        // Catch any type errors, like trying to assign a string in the
        // JSON request body to a int field in our Person struct. We can
        // interpolate the relevant field name and position into the error
        // message to make it easier for the client to fix.
        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}

        // Catch the error caused by extra unexpected fields in the request
        // body. We extract the field name from the error message and
        // interpolate it in our custom error message.
        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}

        // An io.EOF error is returned by Decode() if the request body is
        // empty.
        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
            return &malformedRequest{status: http.StatusBadRequest, msg: msg}

        // Catch the error caused by the request body being too large.
        case err.Error() == "http: request body too large":
            msg := "Request body must not be larger than 1MB"
            return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

        // Otherwise default to logging the error and sending a 500 Internal
        // Server Error response.
        default:
            log.Printf("Error default 500: %v", err.Error())
            msg := http.StatusText(http.StatusInternalServerError)
            return &malformedRequest{status: http.StatusInternalServerError, msg: msg}
        }
    }

    // Call decode again, using a pointer to an empty anonymous struct as
    // the destination. If the request body only contained a single JSON
    // object this will return an io.EOF error. So if we get anything else,
    // we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
        msg := "Request body must only contain a single JSON object"
        return &malformedRequest{status: http.StatusBadRequest, msg: msg}
    }

    return nil
}
