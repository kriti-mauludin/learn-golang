package handler

import (
	"net/http"

	"github.com/kriti-mauludin/final-project/internal/http/middleware"
	"github.com/kriti-mauludin/final-project/internal/parser"
	"github.com/kriti-mauludin/final-project/internal/presenter/json"
	role_usecase "github.com/kriti-mauludin/final-project/internal/usecase/role"
	"github.com/kriti-mauludin/final-project/internal/usecase/role/entity"

	fiber "github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	parser          parser.Parser
	presenter       json.JsonPresenter
	roleCrudUsecase role_usecase.ICrudRoleUsecase
}

func NewRoleHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	roleCrudUsecase role_usecase.ICrudRoleUsecase,
) *RoleHandler {
	return &RoleHandler{parser, presenter, roleCrudUsecase}
}

func (w *RoleHandler) Register(app fiber.Router) {
	app.Get("/roles/:id", middleware.VerifyJWTToken, w.GetByID)
	app.Post("/roles", middleware.VerifyJWTToken, w.Create)
	app.Put("/roles/:id", middleware.VerifyJWTToken, w.Update)
	app.Delete("/roles/:id", middleware.VerifyJWTToken, w.Delete)
}

// @Summary         Get Todo List by ID
// @Description     Get a Todo List by its ID
// @Tags            Todo List
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List"
// @Success			201 {object} entity.GeneralResponse{data=entity.RoleResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/roles/{id} [get]
func (w *RoleHandler) GetByID(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.roleCrudUsecase.GetByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Retrieve Todo Lists by User ID
// @Description     Retrieve a list of Todo Lists belonging to a user by their User ID
// @Tags            Todo List
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Success			201 {object} entity.GeneralResponse{data=[]entity.RoleResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/roles [get]
func (w *RoleHandler) GetByUserID(c *fiber.Ctx) error {
	userID, err := w.parser.ParserUserID(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.roleCrudUsecase.GetByUserID(c.Context(), userID)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary			Create a new Todo List
// @Description		Create a new Todo List
// @Tags			Todo List
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param			req body entity.RoleReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse{data=entity.RoleReq} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/roles [post]
func (w *RoleHandler) Create(c *fiber.Ctx) error {
	var req entity.RoleReq

	err := w.parser.ParserBodyRequestWithUserID(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.roleCrudUsecase.Create(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Update an existing Todo List by ID
// @Description     Update an existing Todo List
// @Tags            Todo List
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the role"
// @Param			req body entity.RoleReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/roles [put]
func (w *RoleHandler) Update(c *fiber.Ctx) error {
	var req entity.RoleReq
	err := w.parser.ParserBodyWithIntIDPathParamsAndUserID(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.roleCrudUsecase.UpdateByID(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

// @Summary         Delete Todo List by ID
// @Description     Delete an existing Todo List by its ID
// @Tags			Todo List
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param           id path int true "ID of the role"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/roles/{id} [delete]
func (w *RoleHandler) Delete(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.roleCrudUsecase.DeleteByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
