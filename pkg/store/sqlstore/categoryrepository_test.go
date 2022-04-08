package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestCategoryRepository_Create(t *testing.T) {
	type fields struct {
		store *Store
	}
	type args struct {
		c *model.Category
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepository{
				store: tt.fields.store,
			}
			if err := r.Create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCategoryRepository_Get(t *testing.T) {
	type fields struct {
		store *Store
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Category
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepository{
				store: tt.fields.store,
			}
			got, err := r.Get(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCategoryRepository_GetSystem(t *testing.T) {
	type fields struct {
		store *Store
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Category
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepository{
				store: tt.fields.store,
			}
			got, err := r.GetSystem()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSystem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSystem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_category_toModel(t *testing.T) {
	type fields struct {
		Id       int64
		Name     string
		Type     model.OperationPaymentType
		IdOwner  sql.NullInt64
		IdParent sql.NullInt64
		IsEnd    bool
		IsSystem bool
	}
	tests := []struct {
		name   string
		fields fields
		want   model.Category
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := category{
				Id:       tt.fields.Id,
				Name:     tt.fields.Name,
				Type:     tt.fields.Type,
				IdOwner:  tt.fields.IdOwner,
				IdParent: tt.fields.IdParent,
				IsEnd:    tt.fields.IsEnd,
				IsSystem: tt.fields.IsSystem,
			}
			if got := c.toModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
