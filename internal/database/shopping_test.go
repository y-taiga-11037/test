package database

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func NewTestModelDB(d Database) *ConnectDB {
	return &ConnectDB{d}
}

func TestConnectDB_GetShoppingLists(t *testing.T) {

	type fields struct {
		testModelDB Database
	}

	tests := []struct {
		name    string
		fields  fields
		want    []Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "nomal1",
			fields: fields{
				testModelDB: &ConnectDB{},
			},
			want: []Response{
				{
					ShoppingId:  1,
					ShoppingDay: "2021-09-01",
					Products: []Product{
						{
							ProductName: "たまねぎ",
							Quantity:    1,
							Price:       190,
						}, {
							ProductName: "にんじん",
							Quantity:    1,
							Price:       190,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestModelDB(tt.fields.testModelDB)
			got, err := c.GetShoppingLists()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDB.GetShoppingLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectDB.GetShoppingLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectDB_InsertShopping(t *testing.T) {
	type fields struct {
		testModelDB Database
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			// TODO: Add test cases.
			name: "nomal1",
			fields: fields{
				testModelDB: &ConnectDB{},
			},
			args: args{
				body: []byte(`{"shopping_day": "2020-09-01"}`),
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestModelDB(tt.fields.testModelDB)
			got, err := c.InsertShopping(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDB.InsertShopping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConnectDB.InsertShopping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectDB_GetInsertLists(t *testing.T) {

	type fields struct {
		testModelDB Database
	}
	type args struct {
		lastInsertID int64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "nomal1",
			fields: fields{
				testModelDB: &ConnectDB{},
			},
			args: args{
				lastInsertID: 1,
			},
			want: Response{
				ShoppingId:  1,
				ShoppingDay: "2021-09-01",
				Products: []Product{
					{
						ProductName: "たまねぎ",
						Quantity:    1,
						Price:       190,
					}, {
						ProductName: "にんじん",
						Quantity:    1,
						Price:       190,
					},
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestModelDB(tt.fields.testModelDB)
			got, err := c.GetInsertLists(tt.args.lastInsertID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDB.GetInsertLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectDB.GetInsertLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
