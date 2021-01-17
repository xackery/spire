package crudcontrollers

import (
	"eoc/database"
	"eoc/http/routes"
	"eoc/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterAuraController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterAuraController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterAuraController {
	return &CharacterAuraController {
		db:     db,
		logger: logger,
	}
}

func (e *CharacterAuraController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_aura/:character_aura", e.deleteCharacterAura, nil),
		routes.RegisterRoute(http.MethodGet, "character_aura/:character_aura", e.getCharacterAura, nil),
		routes.RegisterRoute(http.MethodGet, "character_auras", e.listCharacterAuras, nil),
		routes.RegisterRoute(http.MethodPatch, "character_aura/:character_aura", e.updateCharacterAura, nil),
		routes.RegisterRoute(http.MethodPut, "character_aura", e.createCharacterAura, nil),
	}
}

// listCharacterAuras godoc
// @Id listCharacterAuras
// @Summary Lists CharacterAuras
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterAura
// @Failure 500 {string} string "Bad query request"
// @Router /character_auras [get]
func (e *CharacterAuraController) listCharacterAuras(c echo.Context) error {
	var results []models.CharacterAura
	err := e.db.QueryContext(models.CharacterAura{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterAura godoc
// @Id getCharacterAura
// @Summary Gets CharacterAura
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterAura
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_aura/{id} [get]
func (e *CharacterAuraController) getCharacterAura(c echo.Context) error {
	characterAuraId, err := strconv.Atoi(c.Param("character_aura"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterAura
	err = e.db.QueryContext(models.CharacterAura{}, c).First(&result, characterAuraId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterAura godoc
// @Id updateCharacterAura
// @Summary Updates CharacterAura
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param id path int true "Id"
// @Param character_aura body models.CharacterAura true "CharacterAura"
// @Success 200 {array} models.CharacterAura
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_aura/{id} [patch]
func (e *CharacterAuraController) updateCharacterAura(c echo.Context) error {
	characterAura := new(models.CharacterAura)
	if err := c.Bind(characterAura); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.CharacterAura{}, c).Model(&models.CharacterAura{}).First(&models.CharacterAura{}, characterAura.ID).Error
	if err != nil || characterAura.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterAura{}, c).Model(&models.CharacterAura{}).Update(&characterAura).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterAura)
}

// createCharacterAura godoc
// @Id createCharacterAura
// @Summary Creates CharacterAura
// @Accept json
// @Produce json
// @Param character_aura body models.CharacterAura true "CharacterAura"
// @Tags CharacterAura
// @Success 200 {array} models.CharacterAura
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_aura [put]
func (e *CharacterAuraController) createCharacterAura(c echo.Context) error {
	characterAura := new(models.CharacterAura)
	if err := c.Bind(characterAura); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.CharacterAura{}, c).Model(&models.CharacterAura{}).Create(&characterAura).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterAura)
}

// deleteCharacterAura godoc
// @Id deleteCharacterAura
// @Summary Deletes CharacterAura
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_aura/{id} [delete]
func (e *CharacterAuraController) deleteCharacterAura(c echo.Context) error {
	characterAuraId, err := strconv.Atoi(c.Param("character_aura"))
	if err != nil {
		e.logger.Error(err)
	}

	characterAura := new(models.CharacterAura)
	err = e.db.Get(models.CharacterAura{}, c).Model(&models.CharacterAura{}).First(&characterAura, characterAuraId).Error
	if err != nil || characterAura.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterAura{}, c).Model(&models.CharacterAura{}).Delete(&characterAura).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
