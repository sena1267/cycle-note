package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/sena1267/cycle-note/domain/model"
	medicinev1 "github.com/sena1267/cycle-note/gen/protobuf/medicine/v1"
	"github.com/sena1267/cycle-note/gen/protobuf/medicine/v1/medicinev1connect"
	"github.com/sena1267/cycle-note/usecase"
)

type medicineHandler struct {
	medicineUsecase usecase.MedicineUsecase
}

func NewMedicineHandler(medicineUsecase usecase.MedicineUsecase) medicinev1connect.MedicineServiceHandler {
	return &medicineHandler{medicineUsecase: medicineUsecase}
}

func (h *medicineHandler) GetMedicine(ctx context.Context, req *connect.Request[medicinev1.GetMedicineRequest]) (*connect.Response[medicinev1.GetMedicineResponse], error) {
	return nil, nil
}

func (h *medicineHandler) ListMedicines(ctx context.Context, req *connect.Request[medicinev1.ListMedicinesRequest]) (*connect.Response[medicinev1.ListMedicinesResponse], error) {
	medicines, err := h.medicineUsecase.ListMedicine(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: 別パッケージに分ける
	responseMedicines := make([]*medicinev1.Medicine, 0, len(medicines))
	for _, medicine := range medicines {
		responseMedicines = append(responseMedicines, &medicinev1.Medicine{
			Id:   string(medicine.ID),
			Name: medicine.Name,
			Note: medicine.Note,
		})
	}

	return connect.NewResponse(&medicinev1.ListMedicinesResponse{Medicines: responseMedicines}), nil
}

func (h *medicineHandler) CreateMedicine(ctx context.Context, req *connect.Request[medicinev1.CreateMedicineRequest]) (*connect.Response[medicinev1.CreateMedicineResponse], error) {
	medicine, err := h.medicineUsecase.CreateMedicine(ctx, model.Medicine{
		UserID: model.UserID(1),
		Name:   req.Msg.Name,
		Note:   req.Msg.Note,
	})
	if err != nil {
		return nil, err
	}

	responseMedicine := &medicinev1.Medicine{
		Id:   string(medicine.ID),
		Name: medicine.Name,
		Note: medicine.Note,
	}
	return connect.NewResponse(&medicinev1.CreateMedicineResponse{Medicine: responseMedicine}), nil
}

func (h *medicineHandler) UpdateMedicine(ctx context.Context, req *connect.Request[medicinev1.UpdateMedicineRequest]) (*connect.Response[medicinev1.UpdateMedicineResponse], error) {
	requestMedicine := req.Msg.Medicine
	_, err := h.medicineUsecase.UpdateMedicine(ctx, model.Medicine{
		ID:   model.MedicineID(requestMedicine.Id),
		Name: requestMedicine.Name,
		Note: requestMedicine.Note,
	})
	if err != nil {
		return nil, err
	}
	responseMedicine := &medicinev1.Medicine{
		Id:   string(requestMedicine.Id),
		Name: requestMedicine.Name,
		Note: requestMedicine.Note,
	}
	return connect.NewResponse(&medicinev1.UpdateMedicineResponse{Medicine: responseMedicine}), nil
}

func (h *medicineHandler) DeleteMedicine(ctx context.Context, req *connect.Request[medicinev1.DeleteMedicineRequest]) (*connect.Response[medicinev1.DeleteMedicineResponse], error) {
	if err := h.medicineUsecase.DeleteMedicine(ctx, model.MedicineID(req.Msg.Id)); err != nil {
		return nil, err
	}
	return connect.NewResponse(&medicinev1.DeleteMedicineResponse{}), nil
}
