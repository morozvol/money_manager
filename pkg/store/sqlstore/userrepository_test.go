package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("\"user\"")

	tests := []struct {
		name    string
		store   *Store
		u       *model.User
		wantErr bool
	}{
		{
			"create user",
			store,
			model.GetUser(),
			false,
		},
		{
			"create duplicate user",
			store,
			model.GetUser(),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				store: tt.store,
			}
			if err := r.Create(tt.u); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Find(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("\"user\"")

	u := model.GetUser()
	r := &UserRepository{
		store: store,
	}
	err := r.Create(u)
	if err != nil {
		t.Fatal("не удалось создать окружение")
	}

	tests := []struct {
		name    string
		store   *Store
		id      int
		want    *model.User
		wantErr bool
	}{
		{
			"find exist user",
			store,
			int(u.Id),
			u, false,
		},
		{
			"find non-existent user",
			store,
			int(u.Id + 1),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				store: tt.store,
			}
			got, err := r.Find(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}
