package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/invest-market-backend/appl_row"
)

func (h *Handler) createCollection(c *gin.Context) {
	type Body struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.CreateCollection(appl_row.CollectionCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное создание коллекции",
	})
}

func (h *Handler) getAllCollections(c *gin.Context) {
	res, statusCode, err := h.services.GetAllCollections()
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if len(res) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение имеющихся коллекций",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение имеющихся коллекций",
			"result":  []int{},
		})
	}
}

func (h *Handler) createToken(c *gin.Context) {
	type Body struct {
		Name          string  `json:"name"`
		Price         float64 `json:"price"`
		Author        string  `json:"author"`
		Blockchain    string  `json:"blockchain"`
		UidCollection string  `json:"uid_collection"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.CreateToken(appl_row.TokenCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное создание токена",
	})
}

func (h *Handler) getAllTokensCollection(c *gin.Context) {
	type Body struct {
		UidCollection string `json:"uid_collection"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.GetAllTokensCollection(body.UidCollection)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if len(res) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение токенов коллекции",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение токенов коллекции",
			"result":  []int{},
		})
	}
}

func (h *Handler) getToken(c *gin.Context) {
	type Body struct {
		TokenUid string `json:"token_uid"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.GetToken(body.TokenUid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if len(res) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение токена",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение токена",
			"result":  []int{},
		})
	}
}
