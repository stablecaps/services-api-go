// package tests

// import (
// 	"encoding/json"
// 	"net/http"
// 	"os"
// 	"testing"

// 	"net/http"
// 	"testing"

// 	_ "github.com/stablecaps/services-api-go/pkg/api"
// 	"github.com/stablecaps/services-api-go/pkg/models"
// )

// func TestAPIServer_handleGetServiceById(t *testing.T) {
// 	type fields struct {
// 		listenAddr string
// 		db         models.Dbase
// 	}
// 	type args struct {
// 		writer http.ResponseWriter
// 		req    *http.Request
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			server := &APIServer{
// 				listenAddr: tt.fields.listenAddr,
// 				db:         tt.fields.db,
// 			}
// 			if err := server.handleGetServiceById(tt.args.writer, tt.args.req); (err != nil) != tt.wantErr {
// 				t.Errorf("APIServer.handleGetServiceById() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }


// func TestHealthCheck(t *testing.T) {

// 	response, err := http.Get("http://localhost:8969/health")
// 	if err != nil {
// 		t.Errorf("expected no errors, but got %v", err)
// 	}

// 	if response.StatusCode != http.StatusOK {
// 		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
// 	}

// 	responseBody := ""
// 	json.NewDecoder(response.Body).Decode(&responseBody)
// 	response.Body.Close()

// 	if responseBody != "service is up and running" {
// 		t.Errorf(`expected message to be "service is up and running", but got %v`, responseBody)
// 	}

// 	os.Interrupt.Signal()
// }
