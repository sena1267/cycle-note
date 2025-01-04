package repository

import (
	"context"
	"fmt"

	"github.com/sena1267/cycle-note/domain/model"
	"github.com/sena1267/cycle-note/domain/repository"
	"github.com/sena1267/cycle-note/infrastructure"
	"github.com/sena1267/cycle-note/util/appctx"
)

type medicineRepository struct {
	dbClient *infrastructure.DBClient
}

func NewMedicineRepository(dbClient *infrastructure.DBClient) repository.MedicineRepository {
	return &medicineRepository{dbClient: dbClient}
}

func (r medicineRepository) ListMedicines(ctx context.Context) (model.Medicines, error) {
	medicines := model.Medicines{}

	if err := r.dbClient.DB().NewSelect().Model(&medicines).Where("user_id = ?", appctx.User(ctx).ID).Scan(ctx); err != nil {
		return model.Medicines{}, fmt.Errorf("filed to find medicines. %w", err)
	}
	//query := r.dbClient.DB().NewInsert().AppendQuery()
	//query := infrastructure.ApplyFixedWhere(ctx, r.dbClient.DB().NewSelect().Model(&medicines))
	//if err := query.Scan(ctx); err != nil {
	//	return model.Medicines{}, nil
	//}
	return medicines, nil
}

func (r medicineRepository) GetMedicineByID(ctx context.Context, medicineID model.MedicineID) (model.Medicine, error) {
	medicine := model.Medicine{}
	if err := r.dbClient.DB().NewSelect().Model(&medicine).Where("id = ?", medicineID).Scan(ctx); err != nil {
		return model.Medicine{}, fmt.Errorf("failed to find medicine by id. %w", err)
	}
	return medicine, nil
}

func (r medicineRepository) CreateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error) {
	_, err := r.dbClient.DB().NewInsert().Model(&medicine).Exec(ctx)
	if err != nil {
		return model.Medicine{}, fmt.Errorf("failed to create medicine. %w", err)
	}
	return model.Medicine{}, nil
}

func (r medicineRepository) UpdateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error) {
	_, err := r.dbClient.DB().NewUpdate().Model(&medicine).WherePK().OmitZero().Exec(ctx)
	if err != nil {
		return model.Medicine{}, fmt.Errorf("failed to update medicine. %w", err)
	}
	return model.Medicine{}, nil
}

func (r medicineRepository) DeleteMedicine(ctx context.Context, medicineID model.MedicineID) error {
	_, err := r.dbClient.DB().NewDelete().Model(&model.Medicine{}).Where("id = ?", medicineID).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete medicine. %w", err)
	}
	return nil
}
