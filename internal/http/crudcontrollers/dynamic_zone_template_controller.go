package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type DynamicZoneTemplateController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewDynamicZoneTemplateController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *DynamicZoneTemplateController {
	return &DynamicZoneTemplateController{
		db:	 db,
		logger: logger,
	}
}

func (e *DynamicZoneTemplateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "dynamic_zone_template/:id", e.getDynamicZoneTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "dynamic_zone_templates", e.listDynamicZoneTemplates, nil),
		routes.RegisterRoute(http.MethodPut, "dynamic_zone_template", e.createDynamicZoneTemplate, nil),
		routes.RegisterRoute(http.MethodDelete, "dynamic_zone_template/:id", e.deleteDynamicZoneTemplate, nil),
		routes.RegisterRoute(http.MethodPatch, "dynamic_zone_template/:id", e.updateDynamicZoneTemplate, nil),
		routes.RegisterRoute(http.MethodPost, "dynamic_zone_templates/bulk", e.getDynamicZoneTemplatesBulk, nil),
	}
}

// listDynamicZoneTemplates godoc
// @Id listDynamicZoneTemplates
// @Summary Lists DynamicZoneTemplates
// @Accept json
// @Produce json
// @Tags DynamicZoneTemplate
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZoneTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_templates [get]
func (e *DynamicZoneTemplateController) listDynamicZoneTemplates(c echo.Context) error {
	var results []models.DynamicZoneTemplate
	err := e.db.QueryContext(models.DynamicZoneTemplate{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDynamicZoneTemplate godoc
// @Id getDynamicZoneTemplate
// @Summary Gets DynamicZoneTemplate
// @Accept json
// @Produce json
// @Tags DynamicZoneTemplate
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZoneTemplate
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_template/{id} [get]
func (e *DynamicZoneTemplateController) getDynamicZoneTemplate(c echo.Context) error {
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
	var result models.DynamicZoneTemplate
	query := e.db.QueryContext(models.DynamicZoneTemplate{}, c)
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

// updateDynamicZoneTemplate godoc
// @Id updateDynamicZoneTemplate
// @Summary Updates DynamicZoneTemplate
// @Accept json
// @Produce json
// @Tags DynamicZoneTemplate
// @Param id path int true "Id"
// @Param dynamic_zone_template body models.DynamicZoneTemplate true "DynamicZoneTemplate"
// @Success 200 {array} models.DynamicZoneTemplate
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /dynamic_zone_template/{id} [patch]
func (e *DynamicZoneTemplateController) updateDynamicZoneTemplate(c echo.Context) error {
	request := new(models.DynamicZoneTemplate)
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
	var result models.DynamicZoneTemplate
	query := e.db.QueryContext(models.DynamicZoneTemplate{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.DynamicZoneTemplate{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createDynamicZoneTemplate godoc
// @Id createDynamicZoneTemplate
// @Summary Creates DynamicZoneTemplate
// @Accept json
// @Produce json
// @Param dynamic_zone_template body models.DynamicZoneTemplate true "DynamicZoneTemplate"
// @Tags DynamicZoneTemplate
// @Success 200 {array} models.DynamicZoneTemplate
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /dynamic_zone_template [put]
func (e *DynamicZoneTemplateController) createDynamicZoneTemplate(c echo.Context) error {
	dynamicZoneTemplate := new(models.DynamicZoneTemplate)
	if err := c.Bind(dynamicZoneTemplate); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.DynamicZoneTemplate{}, c).Model(&models.DynamicZoneTemplate{}).Create(&dynamicZoneTemplate).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, dynamicZoneTemplate)
}

// deleteDynamicZoneTemplate godoc
// @Id deleteDynamicZoneTemplate
// @Summary Deletes DynamicZoneTemplate
// @Accept json
// @Produce json
// @Tags DynamicZoneTemplate
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /dynamic_zone_template/{id} [delete]
func (e *DynamicZoneTemplateController) deleteDynamicZoneTemplate(c echo.Context) error {
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
	var result models.DynamicZoneTemplate
	query := e.db.QueryContext(models.DynamicZoneTemplate{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = query.Limit(10000).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDynamicZoneTemplatesBulk godoc
// @Id getDynamicZoneTemplatesBulk
// @Summary Gets DynamicZoneTemplates in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DynamicZoneTemplate
// @Success 200 {array} models.DynamicZoneTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_templates/bulk [post]
func (e *DynamicZoneTemplateController) getDynamicZoneTemplatesBulk(c echo.Context) error {
	var results []models.DynamicZoneTemplate

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

	err := e.db.QueryContext(models.DynamicZoneTemplate{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}