package routes

import (
	"github.com/khozaei/healthio/ent/patient"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/khozaei/healthio/ent"
	"github.com/khozaei/healthio/ent/reception"
)

func createReception(c *gin.Context) {
	r := ent.Reception{}
	patientID := c.Param("pid")
	_patientID, err := strconv.Atoi(patientID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	newReception := handlerConfig.Client.Reception.Create()
	body := make(map[string]any)
	err = c.ShouldBindBodyWith(&body, binding.JSON)
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

	r2, err := newReception.SetPatientID(_patientID).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &r2)
}

func deleteReception(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	_, err = handlerConfig.Client.Reception.UpdateOneID(_id).
		SetDeletedAt(time.Now()).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &gin.H{"deleted_id": id})
}

func updateReception(c *gin.Context) {
	id := c.Param("id")
	r := ent.Reception{}
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
	err = c.ShouldBindBodyWith(&r, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	updReception := handlerConfig.Client.Reception.UpdateOneID(_id)
	for k := range body {
		switch k {
		case reception.FieldReceptionFor:
			updReception.SetReceptionFor(r.ReceptionFor)
		case reception.FieldInsuranceCode:
			updReception.SetInsuranceCode(r.InsuranceCode)
		case reception.FieldVisitDuration:
			updReception.SetVisitDuration(r.VisitDuration)
		case reception.FieldDescription:
			updReception.SetDescription(r.Description)
		}
	}
	up, err := updReception.SetUpdatedAt(time.Now()).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &up)
}

func getReception(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	r, err := handlerConfig.Client.Reception.Query().
		Where(reception.ID(_id)).Only(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &r)
}

func getReceptionByPatient(c *gin.Context) {
	pid := c.Param("pid")
	_pid, err := strconv.Atoi(pid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	rs, err := handlerConfig.Client.Reception.Query().
		Where(reception.HasPatientWith(patient.ID(_pid)), reception.DeletedAtIsNil()).
		All(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &rs)
}

func searchReception(c *gin.Context) {
	from := c.Param("from")
	to := c.Param("to")
	offset := c.Param("offset")
	limit := c.Param("limit")
	_offset, err := strconv.Atoi(offset)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	_limit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	_from, err := time.Parse("2006-01-02", from)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	_to, err := time.Parse("2006-01-02", to)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	rs, err := handlerConfig.Client.Reception.Query().
		Where(reception.And(
			reception.ReceptionForGTE(_from),
			reception.ReceptionForLTE(_to),
		), reception.DeletedAtIsNil()).
		Offset(_offset).Limit(_limit).
		All(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &rs)
}

func setReceptionRoutes(rg *gin.RouterGroup) {
	rs := rg.Group("/receptions")

	rs.POST("/new/:pid", createReception)
	rs.DELETE("/:id", deleteReception)
	rs.POST("/update/:id", updateReception)
	rs.GET("/:id", getReception)
	rs.GET("/by/patient/:pid", getReceptionByPatient)
	rs.GET("/query/:from/:to/:offset/:limit", searchReception)
}
