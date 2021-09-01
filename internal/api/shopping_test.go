package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"
	"github.com/gorilla/mux"

	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
)

type testModelStub struct{}

func (t *testModelStub) GetShoppingLists() ([]db.Response, error) {

	//responseSlice := make([]Response, 0)

	var responseSlice = []db.Response{
		{
			ShoppingId:  1,
			ShoppingDay: "2021-08-12",
			Products: []db.Product{
				{
					ProductName: "にんじん",
					Quantity:    1,
					Price:       190,
				},
			},
		},
	}

	return responseSlice, nil
}

func TestGetShoppingListsHandler(t *testing.T) {

	router := mux.NewRouter()

	type fields struct {
		testModel TestModel
	}

	type args struct {
		w          *httptest.ResponseRecorder
		httpMethod string
		path       string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				testModel: &testModelStub{},
			},
			args: args{
				w:          httptest.NewRecorder(),
				httpMethod: http.MethodGet,
				path:       "/api/shopping",
			},
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewGetHandler(tt.fields.testModel)
			r := httptest.NewRequest(tt.args.httpMethod, tt.args.path, nil)
			router.HandleFunc("/api/shopping", h.GetShoppingListsHandler).Methods("GET")
			router.ServeHTTP(tt.args.w, r)
			if tt.wantCode != tt.args.w.Code {
				t.Errorf("GetCategory() code = %v, wantCode %v", tt.wantCode, tt.args.w.Code)
			}
		})
	}
}
