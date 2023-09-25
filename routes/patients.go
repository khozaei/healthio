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
	}
	err = c.ShouldBindBodyWith(&p, binding.JSON)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
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
	}

	c.JSON(http.StatusOK, &p2)
}

func deletePatient(c *gin.Context) {
	id := c.Param("id")
	log.Printf("id= %v", id)
	// if err != nil {
	//	log.Println(err)
	//		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
	//	}
	_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
	}
	p, err := handlerConfig.Client.Patient.
		Query().Where(patient.ID(_id)).Only(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
	}
	_, err = p.Update().SetDeletedAt(time.Now()).Save(handlerConfig.Context)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &gin.H{"error": err})
	}
	c.JSON(http.StatusOK, &gin.H{"deleted_id": id})
}

func setPatientRoutes(rg *gin.RouterGroup) {
	patients := rg.Group("/patients")

	patients.POST("/new", createPatient)
	patients.DELETE("/:id", deletePatient)
}
