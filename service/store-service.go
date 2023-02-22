package service

import (
	"errors"

	"github.com/NomanSalhab/go_gin_my_first_project/entity"
)

type StoreService interface {
	AddStore(entity.Store) error
	FindAllStores() []entity.Store
	FindActiveStores() []entity.Store
	FindNotActiveStores() []entity.Store
	FindStore(id entity.StoreInfoRequest) (entity.Store, error)
	EditStore(storeEditInfo entity.StoreEditRequest) error
	DeleteStore(store entity.StoreDeleteRequest) error
}

type storeService struct {
	stores []entity.Store
}

func NewStoreService() StoreService {
	return &storeService{}
}

func (service *storeService) AddStore(store entity.Store) error {
	successStore := store
	for i := 0; i < len(service.stores); i++ {
		if service.stores[i].Name == store.Name {
			return errors.New("store category already exisists")
		}
	}
	if len(service.stores) > 0 {
		successStore.ID = service.stores[len(service.stores)-1].ID + 1
	} else {
		successStore.ID = 1
	}

	if !store.Active {
		store.Active = false
	}
	/*var s *storeCategoryService
	storeCategories := s.FindAllStoreCategories()
	if storeCategories != nil {
		for i := 0; i < len(storeCategories); i++ {
			if storeCategories[i].ID == successStore.StoreCategoryId {*/
	service.stores = append(service.stores, successStore)
	return nil
	/*}
		}
	}
	return errors.New("sent store category id does not exist")*/
}

func (service *storeService) FindAllStores() []entity.Store {
	return service.stores
}

func (service *storeService) FindActiveStores() []entity.Store {
	var activeStores []entity.Store
	for i := 0; i < len(service.stores); i++ {
		if service.stores[i].Active {
			activeStores = append(activeStores, service.stores[i])
		}
	}
	return activeStores
}

func (service *storeService) FindNotActiveStores() []entity.Store {
	var notActiveStores []entity.Store
	for i := 0; i < len(service.stores); i++ {
		if !service.stores[i].Active {
			notActiveStores = append(notActiveStores, service.stores[i])
		}
	}
	return notActiveStores
}

func (service *storeService) FindStore(id entity.StoreInfoRequest) (entity.Store, error) {
	stores := service.FindAllStores()
	var store entity.Store
	for i := 0; i < len(stores) && len(stores) != 0; i++ {
		if id.ID != 0 {
			if stores[i].ID == id.ID {
				store = stores[i]
			}
		} else {
			return store, errors.New("store id cannot be zero")
		}
	}
	if id.ID == 0 {
		return store, errors.New("store id cannot be zero")
	}
	if store.StoreCategoryId == 0 {
		return store, errors.New("the store category couldn't be found")
	}
	if store.Name == "" {
		return store, errors.New("the store couldn't be found")
	}
	return store, nil
}

func (service *storeService) EditStore(storeEditInfo entity.StoreEditRequest) error {
	stores := service.FindAllStores()
	var store entity.Store
	for i := 0; i < len(stores) && len(stores) != 0; i++ {
		if storeEditInfo.ID != 0 {
			if stores[i].ID == storeEditInfo.ID {
				store = entity.Store{
					ID:     storeEditInfo.ID,
					Name:   storeEditInfo.Name,
					Active: storeEditInfo.Active,
				}
				if storeEditInfo.Name != "" {
					stores[i].Name = storeEditInfo.Name
				}
				if storeEditInfo.Active {
					stores[i].Active = storeEditInfo.Active
				}
				if storeEditInfo.Balance != 0 {
					stores[i].Balance = storeEditInfo.Balance
				}
				if storeEditInfo.StoreCategoryId != 0 {
					stores[i].StoreCategoryId = storeEditInfo.StoreCategoryId
				}
				if storeEditInfo.Image != "" {
					stores[i].Image = storeEditInfo.Image
				}
			}
		} else {
			return errors.New("store id cannot be zero")
		}
	}
	if storeEditInfo.ID == 0 {
		return errors.New("store id cannot be zero")
	}
	if store.ID == 0 {
		return errors.New("the store couldn't be found")
	}
	return nil
}

func (service *storeService) DeleteStore(storeId entity.StoreDeleteRequest) error {
	stores := service.FindAllStores()
	var tempStore []entity.Store
	if storeId.ID == 0 {
		return errors.New("store id cannot be zero")
	}
	for i := 0; i < len(stores) && len(stores) != 0; i++ {
		if storeId.ID != 0 {
			if stores[i].ID != storeId.ID {
				tempStore = append(tempStore, stores[i])
			}
		}
	}
	if len(stores) != len(tempStore)+1 {
		return errors.New("store could not be found")
	}
	service.stores = tempStore
	return nil
}
