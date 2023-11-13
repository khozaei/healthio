package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/khozaei/healthio/ent"
	"github.com/khozaei/healthio/ent/reception"
	"github.com/khozaei/healthio/ent/visit"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateVisit(c *gin.Context) {
	v := ent.Visit{}
	rid := c.Param("rid")
	_rid, err := strconv.Atoi(rid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	newVisit := handlerConfig.Client.Visit.Create()
	body := make(map[string]any)

	err = c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	err = c.ShouldBindBodyWith(&v, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	for k := range body {
		switch k {
		case visit.FieldVisitPrice:
			newVisit.SetVisitPrice(v.VisitPrice)
		case visit.FieldPaymentType:
			newVisit.SetPaymentType(v.PaymentType)
		case visit.FieldVisitedAt:
			newVisit.SetVisitedAt(v.VisitedAt)
		case visit.FieldIsPaid:
			newVisit.SetIsPaid(v.IsPaid)
		}
	}
	newVisit.SetReceptionID(_rid)
	v2, err := newVisit.Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &v2)
}

func deleteVisit(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	_, err = handlerConfig.Client.Visit.UpdateOneID(_id).
		SetDeletedAt(time.Now()).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &gin.H{"deleted_id": id})
}

func updateVisit(c *gin.Context) {
	id := c.Param("id")
	v := ent.Visit{}
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	body := make(map[string]any)

	err = c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	err = c.ShouldBindBodyWith(&v, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	updVisit := handlerConfig.Client.Visit.UpdateOneID(_id)
	for k := range body {
		switch k {
		case visit.FieldVisitPrice:
			updVisit.SetVisitPrice(v.VisitPrice)
		case visit.FieldPaymentType:
			updVisit.SetPaymentType(v.PaymentType)
		case visit.FieldVisitedAt:
			updVisit.SetVisitedAt(v.VisitedAt)
		case visit.FieldIsPaid:
			updVisit.SetIsPaid(v.IsPaid)
		}
	}
	up, err := updVisit.SetUpdatedAt(time.Now()).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &up)
}

func getVisit(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	v, err := handlerConfig.Client.Visit.Query().
		Where(visit.ID(_id)).Only(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &v)
}

func getVisitByReception(c *gin.Context) {
	rid := c.Param("rid")
	_rid, err := strconv.Atoi(rid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	vs, err := handlerConfig.Client.Visit.Query().
		Where(visit.HasReceptionWith(reception.ID(_rid)), visit.DeletedAtIsNil()).
		All(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &vs)
}

func setVisitsRoutes(rg *gin.RouterGroup) {
	visits := rg.Group("/visits")

	visits.POST("/new/:rid", CreateVisit)
	visits.DELETE("/:id", deleteVisit)
	visits.POST("/update/:id", updateVisit)
	visits.GET("/:id", getVisit)
	visits.GET("/by/reception/:rid", getVisitByReception)
}
