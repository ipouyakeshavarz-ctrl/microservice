package mapper

import (
	"gatewayapp/internal/dto"

	storepb "myapp/api/gen/store"
)

func ToAddress(address *storepb.Address) dto.Address {
	if address == nil {
		return dto.Address{}
	}

	return dto.Address{
		Street:      address.Street,
		City:        address.City,
		Province:    address.Province,
		PostalCode:  address.PostalCode,
		Description: address.Description,
	}
}

func ToStoreInfo(store *storepb.StoreInfo) dto.StoreInfo {
	if store == nil {
		return dto.StoreInfo{}
	}

	return dto.StoreInfo{
		ID:          store.Id,
		UserID:      store.UserId,
		Name:        store.Name,
		Description: store.Description,
		LogoURL:     store.LogoUrl,
		Address:     ToAddress(store.Address),
		PhoneNumber: store.PhoneNumber,
		IsActive:    store.IsActive,
	}
}

func ToCreateStoreRequest(req *dto.CreateStoreRequest) *storepb.CreateStoreRequest {
	return &storepb.CreateStoreRequest{
		UserId:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
		PhoneNumber: req.PhoneNumber,
		LogoUrl:     req.LogoURL,
		Address: &storepb.Address{
			Street:      req.Address.Street,
			City:        req.Address.City,
			Province:    req.Address.Province,
			PostalCode:  req.Address.PostalCode,
			Description: req.Address.Description,
		},
	}
}

func ToCreateStoreResponse(res *storepb.CreateStoreResponse) *dto.CreateStoreResponse {
	return &dto.CreateStoreResponse{
		Store: ToStoreInfo(res.Store),
	}
}

func ToUpdateStoreRequest(req *dto.UpdateStoreRequest) *storepb.UpdateStoreRequest {
	return &storepb.UpdateStoreRequest{
		StoreId:     req.StoreID,
		UserId:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		LogoUrl:     req.LogoURL,
		PhoneNumber: req.PhoneNumber,
		Address: &storepb.Address{
			Street:      req.Address.Street,
			City:        req.Address.City,
			Province:    req.Address.Province,
			PostalCode:  req.Address.PostalCode,
			Description: req.Address.Description,
		},
		IsActive: req.IsActive,
	}
}

func ToUpdateStoreResponse(res *storepb.UpdateStoreResponse) *dto.UpdateStoreResponse {
	return &dto.UpdateStoreResponse{
		Store: ToStoreInfo(res.Store),
	}
}

func ToListStoresByUserRequest(req *dto.ListStoresByUserRequest) *storepb.ListStoresByUserRequest {
	return &storepb.ListStoresByUserRequest{
		UserId: req.UserID,
	}

}

func ToListStoresByUserResponse(res *storepb.ListStoresByUserResponse) *dto.ListStoresByUserResponse {
	stores := make([]dto.StoreInfo, 0, len(res.Stores))

	for _, store := range res.Stores {
		stores = append(stores, ToStoreInfo(store))
	}

	return &dto.ListStoresByUserResponse{
		Stores: stores,
	}
}

func ToDeleteStoreRequest(req *dto.DeleteStoreRequest) *storepb.DeleteStoreRequest {
	return &storepb.DeleteStoreRequest{
		UserId:  req.UserID,
		StoreId: req.StoreID,
	}
}

func ToDeleteStoreResponse(res *storepb.DeleteStoreResponse) *dto.DeleteStoreResponse {
	return &dto.DeleteStoreResponse{
		Success: res.Success,
	}
}
