package portobufmaper

import (
	"myapp/api/gen/store"
	"storeapp/internal/domain"
	"storeapp/internal/param"
)

func MapListResponseToProtobuf(g param.ListStoresByUserResponse) *store.ListStoresByUserResponse {
	r := &store.ListStoresByUserResponse{}
	for _, i := range g.Stores {
		r.Stores = append(r.Stores, &store.StoreInfo{
			Id:          uint64(i.ID),
			UserId:      uint64(i.UserID),
			Name:        i.Name,
			Description: i.Description,
			IsActive:    i.IsActive,
			Address: &store.Address{
				Street:      i.Address.Street,
				City:        i.Address.City,
				Province:    i.Address.Province,
				PostalCode:  i.Address.PostalCode,
				Description: i.Address.Description,
			},
			PhoneNumber: i.PhoneNumber,
			LogoUrl:     i.LogoURL,
		})

	}
	return r
}

func MapListStoreResponseFromProtobuf(g *store.ListStoresByUserResponse) param.ListStoresByUserResponse {
	r := param.ListStoresByUserResponse{}
	for _, i := range g.Stores {
		r.Stores = append(r.Stores, param.StoreInfo{
			ID:          uint(i.Id),
			UserID:      uint(i.UserId),
			Name:        i.Name,
			Description: i.Description,
			IsActive:    i.IsActive,
			Address: domain.Address{
				Street:      i.Address.Street,
				City:        i.Address.City,
				Province:    i.Address.Province,
				PostalCode:  i.Address.PostalCode,
				Description: i.Address.Description,
			},
			PhoneNumber: i.PhoneNumber,
			LogoURL:     i.LogoUrl,
		})

	}
	return r
}
