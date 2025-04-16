package controller

import (
	"dynamic/dao"
	"dynamic/helper"
	"dynamic/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.Redirect(http.StatusFound, "/dynamic/")
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "Error.html", gin.H{
		"Title":     "Page Not Found",
		"PageTitle": "Ops! Página não encontrada",
		"Message":   "A página que você está procurando pode ter sido removida ou está temporariamente indisponível.",
		"Error":     404,
	})
}

func GetDynamic(c *gin.Context) {
	search := c.Query("search")
	sortOrder := c.Query("sort")
	products, _ := dao.GetProducts()

	if search != "" {
		var root *model.Node
		for _, p := range products {
			root = helper.Insert(root, p)
		}
		var results []model.Product
		helper.Autocomplete(root, search, &results)
		products = results
	}

	if sortOrder == "asc" {
		products = helper.MergeSort(products, true)
	} else if sortOrder == "desc" {
		products = helper.MergeSort(products, false)
	}

	if c.Query("k-type") == "min" {
		products = helper.GetKMinProducts(products, 5)
	} else if c.Query("k-type") == "max" {
		products = helper.GetKMaxProducts(products, 5)
	}

	c.HTML(http.StatusOK, "Page.html", gin.H{
		"Products": products,
	})
}
