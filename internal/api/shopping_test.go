package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	db "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/database"
	_ "gh.iiji.jp/y-taiga/mdtd_bootcamp/internal/logger"
	"github.com/gorilla/mux"
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

func (t *testModelStub) InsertShopping(body []byte) (int64, error) {
	return 1, nil
}

func (t *testModelStub) InsertProduct(body []byte, lastInsertID int64) error {
	return nil
}

func (t *testModelStub) GetInsertLists(lastInsertID int64) (db.Response, error) {

	var responseSlice = db.Response{ShoppingId: 1, ShoppingDay: "2021-08-12", Products: []db.Product{
		{
			ProductName: "にんじん",
			Quantity:    1,
			Price:       190,
		},
	}}

	return responseSlice, nil

}

func TestGetShoppingListsHandler(t *testing.T) {

	router := mux.NewRouter()

	type fields struct {
		testModel Shopping
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
			h := NewShoppingHandler(tt.fields.testModel)
			r := httptest.NewRequest(tt.args.httpMethod, tt.args.path, nil)
			router.HandleFunc("/api/shopping", h.GetShoppingListsHandler).Methods("GET")
			router.ServeHTTP(tt.args.w, r)
			if tt.wantCode != tt.args.w.Code {
				t.Errorf("GetCategory() code = %v, wantCode %v", tt.wantCode, tt.args.w.Code)
			}
		})
	}
}

func TestPostShoppingListsHandler(t *testing.T) {
	router := mux.NewRouter()

	type fields struct {
		testModel Shopping
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
				httpMethod: http.MethodPost,
				path:       "/api/shopping",
			},
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewShoppingHandler(tt.fields.testModel)
			r := httptest.NewRequest(tt.args.httpMethod, tt.args.path, nil)
			r.Header.Set("Content-Type", "application/json")
			router.HandleFunc("/api/shopping", h.PostShoppingListsHandler).Methods("POST")
			router.ServeHTTP(tt.args.w, r)
			if tt.wantCode != tt.args.w.Code {
				t.Errorf("GetCategory() code = %v, wantCode %v", tt.wantCode, tt.args.w.Code)
			}
		})
	}
}
