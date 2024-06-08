package handlers

import (
	"net/http"

	"github.com/akposiyefa/go-gin-migration/core/models"
	"github.com/akposiyefa/go-gin-migration/internal"
	"github.com/akposiyefa/go-gin-migration/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func CreateRecords(c *gin.Context) {

	file, err := excelize.OpenFile("input.xlsx")

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Failed to open Excel file", false)
		return
	}

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry unable to read records", false)
		return
	}
	for _, row := range rows {
		record := models.Record{
			Genotype: row[0],
			Age:      row[1],
			Address:  row[2],
		}
		result := internal.DB.Create(&record)
		if result.Error != nil {
			utils.WriteError(c, http.StatusBadRequest, "Sorry unable to create records", false)
			return
		}
	}
	utils.WriteSuccess(c, http.StatusCreated, "Record created successfully", map[string]string{}, true)

}

func GetRecords(c *gin.Context) {

}

func GetSingleRecord(c *gin.Context) {

}

func UpdateRecord(c *gin.Context) {

}

func DeleteRecord(c *gin.Context) {

}
