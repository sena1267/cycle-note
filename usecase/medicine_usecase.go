package usecase

import (
	"context"
	"fmt"

	"github.com/sena1267/cycle-note/domain/model"
	"github.com/sena1267/cycle-note/domain/repository"
	"github.com/sena1267/cycle-note/pkg/util/xidgen"
	"github.com/sena1267/cycle-note/util/appctx"
)

type MedicineUsecase interface {
	ListMedicine(ctx context.Context) (model.Medicines, error)
	CreateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error)
	UpdateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error)
	DeleteMedicine(ctx context.Context, medicineID model.MedicineID) error
}

type medicineUsecase struct {
	medicineRepo repository.MedicineRepository
}

func NewMedicineUsecase(medicineRepo repository.MedicineRepository) MedicineUsecase {
	return medicineUsecase{
		medicineRepo: medicineRepo,
	}
}

func (u medicineUsecase) ListMedicine(ctx context.Context) (model.Medicines, error) {
	medicines, err := u.medicineRepo.ListMedicines(ctx)
	if err != nil {
		return model.Medicines{}, fmt.Errorf("failed to list medicines. %w", err)
	}
	return medicines, nil
}

func (u medicineUsecase) CreateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error) {
	// TODO: util的な処理をまとめたCoreを作る
	medicine.ID = model.MedicineID(xidgen.XIDGenerator{}.Generate())
	medicine.UserID = appctx.User(ctx).ID
	_, err := u.medicineRepo.CreateMedicine(ctx, medicine)
	if err != nil {
		return model.Medicine{}, err
	}
	return medicine, nil
}

func (u medicineUsecase) UpdateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error) {
	_, err := u.medicineRepo.UpdateMedicine(ctx, medicine)
	if err != nil {
		return model.Medicine{}, err
	}
	return medicine, nil
}

func (u medicineUsecase) DeleteMedicine(ctx context.Context, medicineID model.MedicineID) error {
	if err := u.medicineRepo.DeleteMedicine(ctx, medicineID); err != nil {
		return err
	}
	return nil
}
