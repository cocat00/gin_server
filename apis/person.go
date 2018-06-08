package apis

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	. "gin_server/models"
)

func IndexApi(c *gin.Context)  {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context)  {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	person := Person{FirstName:firstName, LastName:lastName}

	ra_rows, err := person.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("insert person Id {}", ra_rows)
	msg := fmt.Sprintf("insert successful %d", ra_rows)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context)  {
	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}

func GetPersonApi(c *gin.Context)  {
	cid := c.Param("id")

	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	p := Person{Id:id}

	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

func ModifyPersonApi(c *gin.Context)  {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	person := Person{Id: id}
	err = c.Bind(&person)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err := person.ModifyPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeletePersonApi(c *gin.Context)  {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}

	person := Person{Id: id}
	ra, err := person.DeletePerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Delete person %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
