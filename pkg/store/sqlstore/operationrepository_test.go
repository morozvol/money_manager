package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"testing"
	"time"
)

func TestOperationRepository_Create(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("operation", "account", "\"user\"")
	r := &UserRepository{
		store: store,
	}

	cr := &CategoryRepository{
		store: store,
	}

	ar := &AccountRepository{
		store: store,
	}

	u := model.GetUser()

	ac := model.GetAccount100(int(u.Id), model.Card, model.Currency{Id: 1})

	cat := model.GetCategory(model.Consumption, 0, 0, true, false)

	err := r.Create(u)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}

	err = cr.Create(cat)
	if err != nil {
		t.Fatal("категория не может быть создана")
	}

	err = ar.Create(ac)
	if err != nil {
		t.Fatal("счёт не может быть создан")
	}

	tests := []struct {
		name    string
		store   *Store
		args    []*model.Operation
		wantErr bool
	}{
		{
			name:    "test",
			store:   store,
			args:    []*model.Operation{{1, ac.Id, time.Now(), 50, *cat, ""}},
			wantErr: false,
		},
		{
			name:    "test",
			store:   store,
			args:    []*model.Operation{{1, ac.Id, time.Now(), 150, *cat, ""}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &OperationRepository{
				store: tt.store,
			}
			if err := r.Create(tt.args...); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperationRepository_Find(t *testing.T) {
	store, truncate := GetTestDBStore(t)
	defer truncate("operation", "account", "\"user\"")

	r := &UserRepository{
		store: store,
	}
	ar := &AccountRepository{
		store: store,
	}
	cr := &CategoryRepository{
		store: store,
	}
	or := &OperationRepository{
		store: store,
	}

	u := model.GetUser()
	err := r.Create(u)
	if err != nil {
		t.Fatal("пользователь не может быть создан")
	}

	cat := model.GetCategory(model.Consumption, 0, 0, true, false)
	err = cr.Create(cat)
	if err != nil {
		t.Fatal("Категория не может быть создана")
	}

	ac := model.GetAccount100(int(u.Id), model.Card, model.Currency{Id: 1})
	err = ar.Create(ac)
	if err != nil {
		t.Fatal("счёт не может быть создан")
	}

	o1 := model.GetOperation(*ac, 50, *cat)
	err = or.Create(o1)
	if err != nil {
		t.Fatal("операция не может быть создана")
	}

	tests := []struct {
		name    string
		store   *Store
		id      int
		wantErr bool
	}{
		{
			"test",
			store,
			int(o1.Id),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &OperationRepository{
				store: tt.store,
			}
			_, err := r.Find(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
