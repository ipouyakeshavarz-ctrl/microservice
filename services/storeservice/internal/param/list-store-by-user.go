package param

import "storeapp/internal/domain"

type ListStoresByUserRequest struct {
	UserID uint `json:"user_id"`
}
type ListStoresByUserResponse struct {
	Stores []StoreInfo `json:"stores"`
}

func StoreToInfo(s *domain.Store) StoreInfo {
	return StoreInfo{
		ID:          s.ID,
		UserID:      s.UserID,
		Name:        s.Name,
		Description: s.Description,
		LogoURL:     s.LogoURL,
		Address: domain.Address{
			Street:      s.Address.Street,
			City:        s.Address.City,
			Province:    s.Address.Province,
			PostalCode:  s.Address.PostalCode,
			Description: s.Address.Description,
		},
		PhoneNumber: s.PhoneNumber,
		IsActive:    s.IsActive,
	}
}

func InfoToStore(info *StoreInfo) domain.Store {
	return domain.Store{
		ID:          info.ID,
		UserID:      info.UserID,
		Name:        info.Name,
		Description: info.Description,
		LogoURL:     info.LogoURL,
		Address: domain.Address{
			Street:      info.Address.Street,
			City:        info.Address.City,
			Province:    info.Address.Province,
			PostalCode:  info.Address.PostalCode,
			Description: info.Address.Description,
		},
		PhoneNumber: info.PhoneNumber,
		IsActive:    info.IsActive,
	}
}
func StoresToInfos(stores []*domain.Store) []StoreInfo {
	result := make([]StoreInfo, 0, len(stores))
	for _, s := range stores {
		result = append(result, StoreToInfo(s))
	}
	return result
}
func InfosToStores(infos []StoreInfo) []domain.Store {
	result := make([]domain.Store, 0, len(infos))
	for i := range infos {
		store := InfoToStore(&infos[i])
		result = append(result, store)
	}
	return result
}
