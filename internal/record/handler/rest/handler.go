package rest

import (
	"math"
	"strconv"
	"time"

	"github.com/KimNattanan/exprec-backend/internal/record/dto"
	"github.com/KimNattanan/exprec-backend/internal/record/usecase"
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpRecordHandler struct {
	recordUseCase usecase.RecordUseCase
}

func NewHttpRecordHandler(useCase usecase.RecordUseCase) *HttpRecordHandler {
	return &HttpRecordHandler{recordUseCase: useCase}
}

func (h *HttpRecordHandler) Save(c *fiber.Ctx) error {
	req := new(dto.RecordSaveRequest)
	if err := c.BodyParser(req); err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	record, err := dto.FromRecordSaveRequest(req)
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	userID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	record.UserID = userID
	if err := h.recordUseCase.Save(record); err != nil {
		return responses.Error(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToRecordResponse(record))
}

func (h *HttpRecordHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	if err := h.recordUseCase.Delete(id); err != nil {
		return responses.Error(c, err)
	}
	return responses.Message(c, fiber.StatusOK, "record deleted")
}

func (h *HttpRecordHandler) FindByUserID(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	var (
		pageStr  = c.Query("page")
		limitStr = c.Query("limit")
		page     = 1
		limit    = 5
	)
	if x, err := strconv.Atoi(pageStr); err == nil {
		page = x
	}
	if x, err := strconv.Atoi(limitStr); err == nil {
		limit = x
	}
	offset := (page - 1) * limit

	records, totalRecords, err := h.recordUseCase.FindByUserID(userID, offset, limit)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(fiber.Map{
		"data": dto.ToRecordResponseList(records),
		"pagination": fiber.Map{
			"totalRecords": totalRecords,
			"totalPages":   int(math.Ceil(float64(totalRecords) / float64(limit))),
			"currentPage":  page,
			"limit":        limit,
		},
	})
}

func (h *HttpRecordHandler) GetUserDashboardData(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	var (
		timeStartStr = c.Query("timeStart")
		timeEndStr   = c.Query("timeEnd")
		timeStart    time.Time
		timeEnd      = time.Now()
	)
	if x, err := time.Parse(time.RFC3339, timeStartStr); err == nil {
		timeStart = x
	}
	if x, err := time.Parse(time.RFC3339, timeEndStr); err == nil {
		timeEnd = x
	}

	dashboardData, err := h.recordUseCase.GetDashboardDataByUserID(userID, timeStart, timeEnd)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dashboardData)
}
