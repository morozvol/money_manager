package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestCategoryRepository_Create(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("category")
	tests := []struct {
		name    string
		store   *Store
		c       *model.Category
		wantErr bool
	}{
		{
			"valid",
			store,
			model.GetCategory(model.Consumption, 0, 0, true, false),
			false,
		},
		{
			"invalid user non exist",
			store,
			model.GetCategory(model.Consumption, 10, 0, true, false),
			true,
		},
		{
			"invalid parent category non exist",
			store,
			model.GetCategory(model.Consumption, 0, 15, true, false),
			true,
		},
		{
			"invalid user non exist",
			store,
			model.GetCategory(model.Consumption, 0, 15, true, false),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepository{
				store: tt.store,
			}
			if err := r.Create(tt.c); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCategoryRepository_Get(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("category", "\"user\"")

	u1 := model.GetUser()
	u2 := model.GetUser()
	u2.Id = 5
	cat1 := model.GetCategory(model.Consumption, 0, 0, true, false)
	cat2 := model.GetCategory(model.Consumption, int(u2.Id), 0, true, false)

	r := &UserRepository{
		store: store,
	}
	err := r.Create(u1)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}
	err = r.Create(u2)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}

	c := &CategoryRepository{
		store: store,
	}
	err = c.Create(cat1)
	if err != nil {
		t.Fatal("категория не может быть создана")
	}
	err = c.Create(cat2)
	if err != nil {
		t.Fatal("категория не может быть создана")
	}

	tests := []struct {
		name    string
		store   *Store
		userId  int
		want    []model.Category
		wantErr bool
	}{
		{
			"get categories for user without individual categories",
			store,
			int(u1.Id),
			[]model.Category{*cat1},
			false,
		},
		{
			"get categories for user with individual categories",
			store,
			int(u2.Id),
			[]model.Category{*cat1, *cat2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepository{
				store: tt.store,
			}
			got, err := r.Get(tt.userId)
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
	store, truncate := GetTestDBStore(t)
	defer truncate()
	tests := []struct {
		name    string
		store   *Store
		wantErr bool
	}{
		{
			"test get system",
			store,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepository{
				store: tt.store,
			}
			_, err := r.GetSystem()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSystem() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		{
			"test 1",
			fields{
				Id:       1,
				Name:     "test",
				Type:     1,
				IsEnd:    true,
				IsSystem: false,
			},
			model.Category{Id: 1, Name: "test", Type: 1, IsEnd: true},
		},
		{
			"id parent not null",
			fields{
				Id:       1,
				Name:     "test",
				Type:     1,
				IdParent: sql.NullInt64{Int64: 12, Valid: true},
				IsEnd:    true,
				IsSystem: false,
			},
			model.Category{Id: 1, Name: "test", Type: 1, IdParent: 12, IsEnd: true},
		},
		{
			"id parent not null",
			fields{
				Id:       1,
				Name:     "test",
				Type:     1,
				IdParent: sql.NullInt64{Int64: 12, Valid: true},
				IdOwner:  sql.NullInt64{Int64: 2, Valid: true},
				IsEnd:    true,
				IsSystem: false,
			},
			model.Category{Id: 1, Name: "test", Type: 1, IdOwner: 2, IdParent: 12, IsEnd: true},
		},
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
