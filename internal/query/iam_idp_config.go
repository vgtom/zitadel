package query

import (
	"context"
	"github.com/caos/zitadel/internal/domain"
)

func (r *Queries) DefaultIDPConfigByID(ctx context.Context, idpConfigID string) (*domain.IDPConfigView, error) {
	idpConfig := NewIAMIDPConfigReadModel(r.iamID, idpConfigID)
	err := r.eventstore.FilterToQueryReducer(ctx, idpConfig)
	if err != nil {
		return nil, err
	}

	return readModelToIDPConfigView(idpConfig), nil
}
