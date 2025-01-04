package repository

import (
	"context"

	"github.com/sena1267/cycle-note/domain/model"
)

type MedicineRepository interface {
	ListMedicines(ctx context.Context) (model.Medicines, error)
	GetMedicineByID(ctx context.Context, medicineID model.MedicineID) (model.Medicine, error)
	CreateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error)
	UpdateMedicine(ctx context.Context, medicine model.Medicine) (model.Medicine, error)
	DeleteMedicine(ctx context.Context, medicineID model.MedicineID) error
}
