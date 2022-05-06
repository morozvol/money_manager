package system_category

import (
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
	"sync"
)

var lock = new(sync.RWMutex)

type systemCategory struct {
	IdComingTransfer      model.Category
	IdConsumptionTransfer model.Category
	IdUndefined           model.Category
	IdComing              model.Category
}

var singleInstance *systemCategory = nil

func GetCategory(store store.Store) *systemCategory {
	if singleInstance == nil {
		lock.Lock()
		sCategory, err := store.Category().GetSystem()
		if err != nil {
			return nil
		}
		singleInstance = &systemCategory{}
		for _, cat := range sCategory {
			switch cat.Name {
			case "Перевод из счёта":
				singleInstance.IdConsumptionTransfer = cat
			case "Перевод на счёт":
				singleInstance.IdComingTransfer = cat
			case "Не определено":
				singleInstance.IdUndefined = cat
			case "Поступление средств":
				singleInstance.IdComing = cat
			}
		}
		lock.Unlock()
	}
	return singleInstance
}
