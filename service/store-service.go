package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/driver"
	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type StoreService interface {
	AddStore(entity.Store) error
	FindAllStores() []entity.Store
	FindActiveStores(govId int) []entity.Store
	FindNotActiveStores() []entity.Store
	FindStore(id entity.StoreInfoRequest) (entity.Store, error)
	EditStore(storeEditInfo entity.StoreEditRequest) error
	ActivateStore(storeEditInfo entity.StoreInfoRequest) error
	DeactivateStore(storeEditInfo entity.StoreInfoRequest) error
	DeleteStore(store entity.StoreDeleteRequest) error

	// AddMockStores(stores []entity.Store)
}

type storeService struct {
	// stores []entity.Store
	driver driver.StoreDriver
}

func NewStoreService(driver driver.StoreDriver) StoreService {
	return &storeService{
		driver: driver,
	}
}

func (service *storeService) AddStore(store entity.Store) error {
	storesList, err := service.driver.FindAllStores()
	if err != nil {
		return err
	}
	for i := 0; i < len(storesList); i++ {
		if storesList[i].Name == store.Name && storesList[i].AreaID == store.AreaID {
			return errors.New("store name with area_id already exists")
		}
	}
	err = service.driver.AddStore(store)
	if err != nil {
		return err
	}
	return nil
	// successStore := store
	// for i := 0; i < len(service.stores); i++ {
	// 	if service.stores[i].Name == store.Name {
	// 		return errors.New("store category already exisists")
	// 	}
	// }
	// if len(service.stores) > 0 {
	// 	successStore.ID = service.stores[len(service.stores)-1].ID + 1
	// } else {
	// 	successStore.ID = 1
	// }
	// if successStore.Discount < 0 || successStore.Discount > 100 {
	// 	return errors.New("store discount must be in 0-100 range")
	// }
	// service.stores = append(service.stores, successStore)
	// return nil
}

func (service *storeService) FindAllStores() []entity.Store {
	allStores, err := service.driver.FindAllStores()
	if err != nil {
		return make([]entity.Store, 0)
	}
	return allStores
	// return service.stores
}

func (service *storeService) FindActiveStores(areaId int) []entity.Store {
	activeStores, err := service.driver.FindActiveStores(areaId)
	if err != nil {
		return make([]entity.Store, 0)
	}
	return activeStores
	// var activeStores []entity.Store
	// for i := 0; i < len(service.stores); i++ {
	// 	if service.stores[i].AreaID == areaId {
	// 		if service.stores[i].Active {
	// 			activeStores = append(activeStores, service.stores[i])
	// 		}
	// 	}

	// }
	// return activeStores
}

func (service *storeService) FindNotActiveStores() []entity.Store {
	notActiveStores, err := service.driver.FindNotActiveStores()
	if err != nil {
		return make([]entity.Store, 0)
	}
	return notActiveStores
	// var notActiveStores []entity.Store
	// for i := 0; i < len(service.stores); i++ {
	// 	if !service.stores[i].Active {
	// 		notActiveStores = append(notActiveStores, service.stores[i])
	// 	}
	// }
	// return notActiveStores
}

func (service *storeService) FindStore(id entity.StoreInfoRequest) (entity.Store, error) {
	store, _ := service.driver.FindStore(id.ID)
	if store.Name == "" {
		return store, errors.New("the store couldn't be found")
	}
	return store, nil
	// stores := service.FindAllStores()
	// var store entity.Store
	// for i := 0; i < len(stores) && len(stores) != 0; i++ {
	// 	if id.ID != 0 {
	// 		if stores[i].ID == id.ID {
	// 			store = stores[i]
	// 			return store, nil
	// 		}
	// 	} else {
	// 		return store, errors.New("store id cannot be zero")
	// 	}
	// }
	// return store, errors.New("the store couldn't be found")
}

func (service *storeService) EditStore(storeEditInfo entity.StoreEditRequest) error {
	_, err := service.driver.EditStore(storeEditInfo)
	if err != nil {
		return err
	}
	return nil
	// stores := service.FindAllStores()
	// for i := 0; i < len(stores) && len(stores) != 0; i++ {
	// 	if storeEditInfo.ID != 0 {
	// 		if stores[i].ID == storeEditInfo.ID {
	// 			// store = entity.Store{
	// 			// 	ID:     storeEditInfo.ID,
	// 			// 	Name:   storeEditInfo.Name,
	// 			// 	Active: storeEditInfo.Active,
	// 			// }
	// 			if storeEditInfo.Name != "" {
	// 				stores[i].Name = storeEditInfo.Name
	// 			}
	// 			if storeEditInfo.Discount > 0 && storeEditInfo.Discount < 100 {
	// 				stores[i].Discount = storeEditInfo.Discount
	// 			}
	// 			if storeEditInfo.Balance != 0 {
	// 				stores[i].Balance = storeEditInfo.Balance
	// 			}
	// 			if storeEditInfo.AreaID != 0 {
	// 				stores[i].AreaID = storeEditInfo.AreaID
	// 			}
	// 			if storeEditInfo.StoreCategoryId != 0 {
	// 				stores[i].StoreCategoryId = storeEditInfo.StoreCategoryId
	// 			}
	// 			if storeEditInfo.Image != "" {
	// 				stores[i].Image = storeEditInfo.Image
	// 			}
	// 			if storeEditInfo.DeliveryRent != 0 {
	// 				stores[i].DeliveryRent = storeEditInfo.DeliveryRent
	// 			}
	// 			stores[i].Active = storeEditInfo.Active
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store id cannot be zero")
	// 	}
	// }
	// return errors.New("the store couldn't be found")
}

func (service *storeService) ActivateStore(storeEditInfo entity.StoreInfoRequest) error {
	err := service.driver.ActivateStore(storeEditInfo)
	if err != nil {
		return err
	}
	return nil
	// stores := service.FindAllStores()
	// for i := 0; i < len(stores) && len(stores) != 0; i++ {
	// 	if storeEditInfo.ID != 0 {
	// 		if stores[i].ID == storeEditInfo.ID {
	// 			stores[i].Active = true
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store id cannot be zero")
	// 	}
	// }
	// return errors.New("the store couldn't be found")
}

func (service *storeService) DeactivateStore(storeEditInfo entity.StoreInfoRequest) error {
	err := service.driver.DeactivateStore(storeEditInfo)
	if err != nil {
		return err
	}
	return nil
	// stores := service.FindAllStores()
	// for i := 0; i < len(stores) && len(stores) != 0; i++ {
	// 	if storeEditInfo.ID != 0 {
	// 		if stores[i].ID == storeEditInfo.ID {
	// 			stores[i].Active = false
	// 			return nil
	// 		}
	// 	} else {
	// 		return errors.New("store id cannot be zero")
	// 	}
	// }
	// return errors.New("the store couldn't be found")
}

func (service *storeService) DeleteStore(storeId entity.StoreDeleteRequest) error {
	err := service.driver.DeleteStore(storeId.ID)
	if err != nil {
		return err
	}
	return nil
	// stores := service.FindAllStores()
	// var tempStore []entity.Store
	// if storeId.ID == 0 {
	// 	return errors.New("store id cannot be zero")
	// }
	// for i := 0; i < len(stores) && len(stores) != 0; i++ {
	// 	if stores[i].ID != storeId.ID {
	// 		tempStore = append(tempStore, stores[i])
	// 	}
	// }
	// if len(stores) != len(tempStore)+1 {
	// 	return errors.New("store could not be found")
	// }
	// service.stores = tempStore
	// return nil
}

// func (service *storeService) AddMockStores(stores []entity.Store) {
// 	service.stores = append(service.stores, stores...)
// }
