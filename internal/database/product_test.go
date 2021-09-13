package database

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectDB_InsertProduct(t *testing.T) {

	type fields struct {
		testModelDB Database
	}
	type args struct {
		body         []byte
		lastInsertID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			// TODO: Add test cases.
			name: "nomal1",
			fields: fields{
				testModelDB: &ConnectDB{},
			},
			args: args{
				body:         []byte(`{"products":[{"product_name": "たまねぎ", "quantity": 1, "price": 190}]}`),
				lastInsertID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestModelDB(tt.fields.testModelDB)
			if err := c.InsertProduct(tt.args.body, tt.args.lastInsertID); (err != nil) != tt.wantErr {
				t.Errorf("ConnectDB.InsertProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMain(m *testing.M) {
	DB = Connect()
	code := m.Run()
	os.Exit(code)
}
