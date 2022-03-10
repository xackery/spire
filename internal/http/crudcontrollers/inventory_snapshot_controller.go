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

type InventorySnapshotController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewInventorySnapshotController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *InventorySnapshotController {
	return &InventorySnapshotController{
		db:	 db,
		logger: logger,
	}
}

func (e *InventorySnapshotController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "inventory_snapshot/:timeIndex", e.getInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodGet, "inventory_snapshots", e.listInventorySnapshots, nil),
		routes.RegisterRoute(http.MethodPut, "inventory_snapshot", e.createInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodDelete, "inventory_snapshot/:timeIndex", e.deleteInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodPatch, "inventory_snapshot/:timeIndex", e.updateInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodPost, "inventory_snapshots/bulk", e.getInventorySnapshotsBulk, nil),
	}
}

// listInventorySnapshots godoc
// @Id listInventorySnapshots
// @Summary Lists InventorySnapshots
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InventorySnapshot
// @Failure 500 {string} string "Bad query request"
// @Router /inventory_snapshots [get]
func (e *InventorySnapshotController) listInventorySnapshots(c echo.Context) error {
	var results []models.InventorySnapshot
	err := e.db.QueryContext(models.InventorySnapshot{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getInventorySnapshot godoc
// @Id getInventorySnapshot
// @Summary Gets InventorySnapshot
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InventorySnapshot
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /inventory_snapshot/{id} [get]
func (e *InventorySnapshotController) getInventorySnapshot(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	timeIndex, err := strconv.Atoi(c.Param("timeIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TimeIndex]"})
	}
	params = append(params, timeIndex)
	keys = append(keys, "time_index = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [slotid] position [3] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.InventorySnapshot
	query := e.db.QueryContext(models.InventorySnapshot{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.TimeIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateInventorySnapshot godoc
// @Id updateInventorySnapshot
// @Summary Updates InventorySnapshot
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param id path int true "Id"
// @Param inventory_snapshot body models.InventorySnapshot true "InventorySnapshot"
// @Success 200 {array} models.InventorySnapshot
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /inventory_snapshot/{id} [patch]
func (e *InventorySnapshotController) updateInventorySnapshot(c echo.Context) error {
	request := new(models.InventorySnapshot)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	timeIndex, err := strconv.Atoi(c.Param("timeIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TimeIndex]"})
	}
	params = append(params, timeIndex)
	keys = append(keys, "time_index = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [slotid] position [3] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.InventorySnapshot
	query := e.db.QueryContext(models.InventorySnapshot{}, c)
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

// createInventorySnapshot godoc
// @Id createInventorySnapshot
// @Summary Creates InventorySnapshot
// @Accept json
// @Produce json
// @Param inventory_snapshot body models.InventorySnapshot true "InventorySnapshot"
// @Tags InventorySnapshot
// @Success 200 {array} models.InventorySnapshot
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /inventory_snapshot [put]
func (e *InventorySnapshotController) createInventorySnapshot(c echo.Context) error {
	inventorySnapshot := new(models.InventorySnapshot)
	if err := c.Bind(inventorySnapshot); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.InventorySnapshot{}, c).Model(&models.InventorySnapshot{}).Create(&inventorySnapshot).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, inventorySnapshot)
}

// deleteInventorySnapshot godoc
// @Id deleteInventorySnapshot
// @Summary Deletes InventorySnapshot
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param id path int true "timeIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /inventory_snapshot/{id} [delete]
func (e *InventorySnapshotController) deleteInventorySnapshot(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	timeIndex, err := strconv.Atoi(c.Param("timeIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, timeIndex)
	keys = append(keys, "time_index = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [slotid] position [3] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.InventorySnapshot
	query := e.db.QueryContext(models.InventorySnapshot{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.InventorySnapshot{}, c).Model(&models.InventorySnapshot{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getInventorySnapshotsBulk godoc
// @Id getInventorySnapshotsBulk
// @Summary Gets InventorySnapshots in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags InventorySnapshot
// @Success 200 {array} models.InventorySnapshot
// @Failure 500 {string} string "Bad query request"
// @Router /inventory_snapshots/bulk [post]
func (e *InventorySnapshotController) getInventorySnapshotsBulk(c echo.Context) error {
	var results []models.InventorySnapshot

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

	err := e.db.QueryContext(models.InventorySnapshot{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}