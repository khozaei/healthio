package routes

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/khozaei/healthio/ent"
	"github.com/khozaei/healthio/ent/patient"
)

func createPatient(c *gin.Context) {
	p := ent.Patient{}
	newPatient := handlerConfig.Client.Patient.Create()
	body := make(map[string]any)

	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	err = c.ShouldBindBodyWith(&p, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	for k := range body {
		switch k {
		case patient.FieldFirstName:
			newPatient.SetFirstName(p.FirstName)
		case patient.FieldLastName:
			newPatient.SetLastName(p.LastName)
		case patient.FieldFatherName:
			newPatient.SetFatherName(p.FatherName)
		case patient.FieldPhone:
			newPatient.SetPhone(p.Phone)
		case patient.FieldNationalCode:
			newPatient.SetNationalCode(p.NationalCode)
		case patient.FieldIdentityCode:
			newPatient.SetIdentityCode(p.IdentityCode)
		}
	}
	p2, err := newPatient.Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &p2)
}

func deletePatient(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	_, err = handlerConfig.Client.Patient.UpdateOneID(_id).
		SetDeletedAt(time.Now()).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &gin.H{"deleted_id": id})
}

func updatePatient(c *gin.Context) {
	id := c.Param("id")
	p := ent.Patient{}
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
	err = c.ShouldBindBodyWith(&p, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	updPatient := handlerConfig.Client.Patient.UpdateOneID(_id)
	for k := range body {
		switch k {
		case patient.FieldFirstName:
			updPatient.SetFirstName(p.FirstName)
		case patient.FieldLastName:
			updPatient.SetLastName(p.LastName)
		case patient.FieldFatherName:
			updPatient.SetFatherName(p.FatherName)
		case patient.FieldPhone:
			updPatient.SetPhone(p.Phone)
		case patient.FieldNationalCode:
			updPatient.SetNationalCode(p.NationalCode)
		case patient.FieldIdentityCode:
			updPatient.SetIdentityCode(p.IdentityCode)
		}
	}
	up, err := updPatient.SetUpdatedAt(time.Now()).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &up)
}

func getPatient(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	p, err := handlerConfig.Client.Patient.Query().
		Where(patient.ID(_id)).Only(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &p)
}

func searchPatient(c *gin.Context) {
	term := c.Param("term")
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
	ps, err := handlerConfig.Client.Patient.Query().
		Where(patient.Or(
			patient.FirstNameContains(term),
			patient.LastNameContains(term),
			patient.FatherNameContains(term),
			patient.PhoneContains(term),
			patient.NationalCodeContains(term),
			patient.IdentityCodeContains(term),
		), patient.DeletedAtIsNil()).
		Offset(_offset).Limit(_limit).
		All(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, &ps)
}

func setPatientRoutes(rg *gin.RouterGroup) {
	patients := rg.Group("/patients")

	patients.POST("/new", createPatient)
	patients.DELETE("/:id", deletePatient)
	patients.POST("/update/:id", updatePatient)
	patients.GET("/:id", getPatient)
	patients.GET("/query/:term/:offset/:limit", searchPatient)
}
