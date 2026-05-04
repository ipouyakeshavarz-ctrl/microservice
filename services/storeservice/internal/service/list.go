package storeservice

import (
	"context"
	"myapp/pkg/richerror"
	"storeapp/internal/domain"
	"storeapp/internal/param"
)

func (s *Service) ListStoresByUser(ctx context.Context,
	req param.ListStoresByUserRequest) (param.ListStoresByUserResponse, error) {
	const op = "StoreService.ListStoresByUser"

	storeIDs, err := s.repo.ListStoreIDsByUser(ctx, req.UserID)
	if err != nil {
		return param.ListStoresByUserResponse{},
			richerror.New(op).WithErr(err)
	}

	if len(storeIDs) == 0 {
		return param.ListStoresByUserResponse{Stores: []param.StoreInfo{}}, nil
	}

	cachedMap, err := s.storeCache.GetManyByIDs(ctx, storeIDs)
	if err != nil {
		return param.ListStoresByUserResponse{},
			richerror.New(op).WithErr(err)
	}

	var missedIDs []uint
	for _, id := range storeIDs {
		if _, ok := cachedMap[id]; !ok {
			missedIDs = append(missedIDs, id)
		}
	}

	var freshStores []*domain.Store
	if len(missedIDs) > 0 {
		freshStores, err = s.repo.GetStoresByIDs(ctx, missedIDs)
		if err != nil {
			return param.ListStoresByUserResponse{},
				richerror.New(op).WithErr(err)
		}

		for _, st := range freshStores {
			_ = s.storeCache.SetByID(ctx, st.ID, st)
		}
	}
	out := make([]*domain.Store, 0, len(storeIDs))

	for _, id := range storeIDs {
		if st, ok := cachedMap[id]; ok {
			out = append(out, st)
		} else {
			for _, st := range freshStores {
				if st.ID == id {
					out = append(out, st)
				}
			}
		}
	}

	return param.ListStoresByUserResponse{Stores: param.StoresToInfos(out)}, nil
}
