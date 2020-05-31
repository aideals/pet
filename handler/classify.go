package handler

import (
	"log"
	"Pet/models"
	"Pet/constants"
)

func Dog(c *gin.Context) {
	classify := c.Query("classify")

	if classify != "" {
		dogs,err := models.QueryDogByClassify(classify)
		if err != nil {
			log.Printf(constants.DB_ERROR)
		}

		c.JSON(http.StatusOK, gin.H{
			"dogs" : dogs,
		})
	}
}

func Cat(c *gin.Context) {
	classify := c.Query("classify")

	if classify != "" {
		cats,err := models.QueryCatByClassify(classify)
		if err != nil {
			log.Printf(constants.DB_ERROR)
		}

		c.JSON(http.StatusOK, gin.H{
			"cats" : cats,
		})
	}
}

func PetSupplies(c *gin.Context) {
	classify := c.Query("classify")

	if classify != "" {
		petSupplies,err := models.QueryPetSuppliesByClassify(classify)
		if err != nil {
			log.Printf(constants.DB_ERROR)
		}

		c.JSON(http.StatusOK, gin.H{
			"petSupplies" : petSupplies,
		})
	}
}