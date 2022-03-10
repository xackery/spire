package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterInstanceSafereturnController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterInstanceSafereturnController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterInstanceSafereturnController {
	return &CharacterInstanceSafereturnController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterInstanceSafereturnController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_instance_safereturn/:id", e.getCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodGet, "character_instance_safereturns", e.listCharacterInstanceSafereturns, nil),
		routes.RegisterRoute(http.MethodPut, "character_instance_safereturn", e.createCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodDelete, "character_instance_safereturn/:id", e.deleteCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodPatch, "character_instance_safereturn/:id", e.updateCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodPost, "character_instance_safereturns/bulk", e.getCharacterInstanceSafereturnsBulk, nil),
	}
}

// listCharacterInstanceSafereturns godoc
// @Id listCharacterInstanceSafereturns
// @Summary Lists CharacterInstanceSafereturns
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 500 {string} string "Bad query request"
// @Router /character_instance_safereturns [get]
func (e *CharacterInstanceSafereturnController) listCharacterInstanceSafereturns(c echo.Context) error {
	var results []models.CharacterInstanceSafereturn
	err := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterInstanceSafereturn godoc
// @Id getCharacterInstanceSafereturn
// @Summary Gets CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_instance_safereturn/{id} [get]
func (e *CharacterInstanceSafereturnController) getCharacterInstanceSafereturn(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.CharacterInstanceSafereturn
	query := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterInstanceSafereturn godoc
// @Id updateCharacterInstanceSafereturn
// @Summary Updates CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param id path int true "Id"
// @Param character_instance_safereturn body models.CharacterInstanceSafereturn true "CharacterInstanceSafereturn"
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_instance_safereturn/{id} [patch]
func (e *CharacterInstanceSafereturnController) updateCharacterInstanceSafereturn(c echo.Context) error {
	request := new(models.CharacterInstanceSafereturn)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.CharacterInstanceSafereturn
	query := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterInstanceSafereturn godoc
// @Id createCharacterInstanceSafereturn
// @Summary Creates CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Param character_instance_safereturn body models.CharacterInstanceSafereturn true "CharacterInstanceSafereturn"
// @Tags CharacterInstanceSafereturn
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_instance_safereturn [put]
func (e *CharacterInstanceSafereturnController) createCharacterInstanceSafereturn(c echo.Context) error {
	characterInstanceSafereturn := new(models.CharacterInstanceSafereturn)
	if err := c.Bind(characterInstanceSafereturn); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterInstanceSafereturn{}, c).Model(&models.CharacterInstanceSafereturn{}).Create(&characterInstanceSafereturn).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterInstanceSafereturn)
}

// deleteCharacterInstanceSafereturn godoc
// @Id deleteCharacterInstanceSafereturn
// @Summary Deletes CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_instance_safereturn/{id} [delete]
func (e *CharacterInstanceSafereturnController) deleteCharacterInstanceSafereturn(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.CharacterInstanceSafereturn
	query := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.CharacterInstanceSafereturn{}, c).Model(&models.CharacterInstanceSafereturn{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterInstanceSafereturnsBulk godoc
// @Id getCharacterInstanceSafereturnsBulk
// @Summary Gets CharacterInstanceSafereturns in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterInstanceSafereturn
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 500 {string} string "Bad query request"
// @Router /character_instance_safereturns/bulk [post]
func (e *CharacterInstanceSafereturnController) getCharacterInstanceSafereturnsBulk(c echo.Context) error {
	var results []models.CharacterInstanceSafereturn

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}