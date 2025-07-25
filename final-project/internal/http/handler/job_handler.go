package handler

import (
	"net/http"

	"github.com/kriti-mauludin/final-project/internal/http/middleware"
	"github.com/kriti-mauludin/final-project/internal/parser"
	"github.com/kriti-mauludin/final-project/internal/presenter/json"
	job_usecase "github.com/kriti-mauludin/final-project/internal/usecase/job"
	"github.com/kriti-mauludin/final-project/internal/usecase/job/entity"

	fiber "github.com/gofiber/fiber/v2"
)

type JobHandler struct {
	parser         parser.Parser
	presenter      json.JsonPresenter
	jobCrudUsecase job_usecase.ICrudJobUsecase
}

func NewJobHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	jobCrudUsecase job_usecase.ICrudJobUsecase,
) *JobHandler {
	return &JobHandler{parser, presenter, jobCrudUsecase}
}

func (w *JobHandler) Register(app fiber.Router) {
	app.Get("/jobs/:id", middleware.VerifyJWTToken, w.GetByID)
	app.Post("/jobs", middleware.VerifyJWTToken, w.Create)
	app.Put("/jobs/:id", middleware.VerifyJWTToken, w.Update)
	app.Delete("/jobs/:id", middleware.VerifyJWTToken, w.Delete)
}

// @Summary         Get Job by ID
// @Description     Get a Job by its ID
// @Tags            Job
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Job"
// @Success			201 {object} entity.GeneralResponse{data=entity.JobResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/jobs/{id} [get]
func (w *JobHandler) GetByID(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.jobCrudUsecase.GetByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Retrieve Job by User ID
// @Description     Retrieve a list of Job belonging to a user by their User ID
// @Tags            Job
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Success			201 {object} entity.GeneralResponse{data=[]entity.JobResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/jobs [get]
func (w *JobHandler) GetByUserID(c *fiber.Ctx) error {
	userID, err := w.parser.ParserUserID(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.jobCrudUsecase.GetByUserID(c.Context(), userID)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary			Create a new Job
// @Description		Create a new Job
// @Tags			Job
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param			req body entity.JobReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse{data=entity.JobReq} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/jobs [post]
func (w *JobHandler) Create(c *fiber.Ctx) error {
	var req entity.JobReq

	err := w.parser.ParserBodyRequestWithUserID(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.jobCrudUsecase.Create(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Update an existing Job by ID
// @Description     Update an existing Job
// @Tags            Job
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Job"
// @Param			req body entity.JobReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/jobs [put]
func (w *JobHandler) Update(c *fiber.Ctx) error {
	var req entity.JobReq
	err := w.parser.ParserBodyWithIntIDPathParamsAndUserID(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.jobCrudUsecase.UpdateByID(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

// @Summary         Delete Job by ID
// @Description     Delete an existing Job by its ID
// @Tags            Job
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Job"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/jobs/{id} [delete]
func (w *JobHandler) Delete(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.jobCrudUsecase.DeleteByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
