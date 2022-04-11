package system_category

import (
	"github.com/morozvol/money_manager/pkg/store"
	"sync"
)

var lock = new(sync.RWMutex)

type systemCategory struct {
	IdComingTransfer      int
	IdConsumptionTransfer int
}

var singleInstance *systemCategory

func GetCategory(store store.Store) *systemCategory {

	sCategory, err := store.Category().GetSystem()
	if err != nil {
		return nil
	}
	singleInstance = &systemCategory{}
	for _, cat := range sCategory {
		switch cat.Name {
		case "Перевод из счёта":
			singleInstance.IdConsumptionTransfer = cat.Id
		case "Перевод на счёт":
			singleInstance.IdComingTransfer = cat.Id
		}
	}
	return singleInstance
}
