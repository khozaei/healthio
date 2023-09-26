package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/khozaei/healthio/ent"
	"github.com/khozaei/healthio/ent/reception"
)

func createReception(c *gin.Context) {
	r := ent.Reception{}
	newReception := handlerConfig.Client.Reception.Create()
	body := make(map[string]any)

	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	err = c.ShouldBindBodyWith(&r, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	for k := range body {
		switch k {
		case reception.FieldReceptionFor:
			newReception.SetReceptionFor(r.ReceptionFor)
		case reception.FieldInsuranceCode:
			newReception.SetInsuranceCode(r.InsuranceCode)
		case reception.FieldVisitDuration:
			newReception.SetVisitDuration(r.VisitDuration)
		case reception.FieldDescription:
			newReception.SetDescription(r.Description)
		}
	}
	r2, err := newReception.Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &r2)
}

func deleteReception(c *gin.Context) {
}

func updateReception(c *gin.Context) {
}

func getReception(c *gin.Context) {
}

func searchReception(c *gin.Context) {
}

func setReceptionRoutes(rg *gin.RouterGroup) {
	rs := rg.Group("/receptions")

	rs.POST("/new", createReception)
	rs.DELETE("/:id", deleteReception)
	rs.POST("/update/:id", updateReception)
	rs.GET("/:id", getReception)
	rs.GET("/query/:term/:offset/:limit", searchReception)
}
