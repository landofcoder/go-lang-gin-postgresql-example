package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/landofcoder/go-lang-gin-postgresql-example/models"
)

// GET /contacts
// Get all contacts
func FindContacts(c *gin.Context) {
	var contacts []models.Contact
	models.DB.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

// GET /contacts/:id
// Find a contact
func FindContact(c *gin.Context) {  // Get model if exist
	var contact models.Contact
  
	if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": book})
  }

func CreateContact(c *gin.Context) {
	// Validate input
	var input CreateContactInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create contact
	contact := models.Contact{
		FirstName: input.FirstName, 
		LastName: input.LastName, 
		PhoneNumber: input.PhoneNumber, 
		Email: input.Email,
		Company: input.Company,
		JobTitle: input.JobTitle,
		Address: input.Address, 
		City: input.City, 
		State: input.State, 
		ZipCode: input.ZipCode, 
		Country: input.Country, 
		Tags: input.Tags
	}
	models.DB.Create(&contact)

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// PATCH /contacts/:id
// Update a contact
func UpdateContact(c *gin.Context) {
	// Get model if exist
	var contact models.Contact
	if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input UpdateContactInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	models.DB.Model(&contact).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// DELETE /contacts/:id
// Delete a contact
func DeleteContact(c *gin.Context) {
	// Get model if exist
	var contact models.Contact
	if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&contact)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
}