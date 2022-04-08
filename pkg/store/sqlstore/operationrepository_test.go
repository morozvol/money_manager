package sqlstore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"reflect"
	"testing"
)

func TestOperationRepository_Create(t *testing.T) {
	//store, truncate := GetTestStore(t)
	type fields struct {
		store *Store
	}
	type args struct {
		o []*model.Operation
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
			r := &OperationRepository{
				store: tt.fields.store,
			}
			if err := r.Create(tt.args.o...); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	//truncate("operation")
}

func TestOperationRepository_Find(t *testing.T) {
	type fields struct {
		store *Store
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Operation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &OperationRepository{
				store: tt.fields.store,
			}
			got, err := r.Find(tt.args.id)
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
