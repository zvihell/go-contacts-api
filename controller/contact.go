package controller

import (
	"go-contacts-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.QueryRow("INSERT INTO contacts (name,lastname,organization,dolzhnost,mobile,user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		contact.Name,
		contact.Lastname,
		contact.Organization,
		contact.Dolzhnost,
		contact.Mobile,
		contact.UserId,
	).Scan(&contact.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, contact)

}

func (h *Handler) GetContactbyUser(c *gin.Context) {
	id := c.Param("id")

	var contact []models.Contact

	rows, err := h.db.Query("SELECT * FROM contacts WHERE user_id = $1", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		var Id int
		var Name string
		var Lastname string
		var Organization string
		var Dolzhnost string
		var Mobile string
		var UserId int

		rows.Scan(&Id, &Name, &Lastname, &Organization, &Dolzhnost, &Mobile, &UserId)

		contact = append(contact, models.Contact{
			Id:           Id,
			Name:         Name,
			Lastname:     Lastname,
			Organization: Organization,
			Dolzhnost:    Dolzhnost,
			Mobile:       Mobile,
			UserId:       UserId,
		})
	}

	c.JSON(200, contact)

}

func (h *Handler) GetContacts(c *gin.Context) {
	var contact []models.Contact

	rows, err := h.db.Query("SELECT * FROM contacts")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		var Id int
		var Name string
		var Lastname string
		var Organization string
		var Dolzhnost string
		var Mobile string
		var UserId int

		rows.Scan(&Id, &Name, &Lastname, &Organization, &Dolzhnost, &Mobile, &UserId)

		contact = append(contact, models.Contact{
			Id:           Id,
			Name:         Name,
			Lastname:     Lastname,
			Organization: Organization,
			Dolzhnost:    Dolzhnost,
			Mobile:       Mobile,
			UserId:       UserId,
		})
	}

	c.JSON(200, contact)

}

func (h *Handler) UpdateContact(c *gin.Context) {
	contactID := c.Param("id")

	id, err := strconv.ParseInt(contactID, 10, 64)
	if err != nil {
		panic(err)
	}

	var contact models.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE contacts SET name=$1, lastname=$2, organization=$3, dolzhnost=$4, mobile=$5, user_id WHERE id=$7`

	_, err = h.db.Exec(query, contact.Name, contact.Lastname, contact.Organization, contact.Dolzhnost, contact.Mobile, contact.UserId, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "ok")

}
