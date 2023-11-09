// package models

// import (
// 	"reflect"
// 	"testing"
// )

// func TestNewService(t *testing.T) {
// 	type args struct {
// 		ServiceName        string
// 		ServiceDescription string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *Service
// 	}{
// 		{
// 			name: "should not work",
// 			args: args{
// 				ServiceName: "ded",
// 				ServiceDescription: "crwedfc",
// 			},
// 			// a problem due to random elements in new service
// 			want: NewService{
// 				ServiceName: "won't work",
// 				ServiceDescription: " atest that won't work",
// 				ServiceVersions: "v1,v2,v3",
// 				CreatedAt: "2023-11-07 21:38:52.360551",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewService(tt.args.ServiceName, tt.args.ServiceDescription); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewService() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
