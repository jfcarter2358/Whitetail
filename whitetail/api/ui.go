package api

import (
	"log"
	"net/http"
	"path/filepath"
	"whitetail/config"

	"github.com/gin-gonic/gin"

	// "whitetail/ast"

	"io"
	"os"
)

func UpdateLogo(c *gin.Context) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	if extension != ".png" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File is of wrong format",
		})
		return
	}

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "static/img/logo.png"); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	config.UpdateBranding()

	c.Redirect(302, config.Config.BasePath+"/ui/settings")
}

func UpdateIcon(c *gin.Context) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	if extension != ".png" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File is of wrong format",
		})
		return
	}

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "static/img/favicon.png"); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	config.UpdateBranding()

	c.Redirect(302, config.Config.BasePath+"/ui/settings")
}

func UpdateColors(c *gin.Context) {
	var input config.BrandingConfigObject
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Unable to bind JSON")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.Config.Branding.PrimaryColor = input.PrimaryColor
	config.Config.Branding.SecondaryColor = input.SecondaryColor
	config.Config.Branding.TertiaryColor = input.TertiaryColor
	config.Config.Branding.INFOColor = input.INFOColor
	config.Config.Branding.WARNColor = input.WARNColor
	config.Config.Branding.DEBUGColor = input.DEBUGColor
	config.Config.Branding.TRACEColor = input.TRACEColor
	config.Config.Branding.ERRORColor = input.ERRORColor
	config.UpdateBranding()
	c.JSON(http.StatusOK, gin.H{})
}

func DefaultColors(c *gin.Context) {
	config.Config.Branding.PrimaryColor = config.Defaults.PrimaryColor
	config.Config.Branding.SecondaryColor = config.Defaults.SecondaryColor
	config.Config.Branding.TertiaryColor = config.Defaults.TertiaryColor
	config.Config.Branding.INFOColor = config.Defaults.INFOColor
	config.Config.Branding.WARNColor = config.Defaults.WARNColor
	config.Config.Branding.DEBUGColor = config.Defaults.DEBUGColor
	config.Config.Branding.TRACEColor = config.Defaults.TRACEColor
	config.Config.Branding.ERRORColor = config.Defaults.ERRORColor

	config.UpdateBranding()
	c.JSON(http.StatusOK, gin.H{})
}

func DefaultLogo(c *gin.Context) {
	log.Println("Copying default logo file")
	source, err := os.Open("static/img/logo.default.png")
	if err != nil {
		panic(err)
	}
	defer source.Close()

	destination, err := os.Create("static/img/logo.png")
	if err != nil {
		panic(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	c.JSON(http.StatusOK, gin.H{})
}

func DefaultIcon(c *gin.Context) {
	log.Println("Copying default icon file")
	source, err := os.Open("static/img/favicon.default.png")
	if err != nil {
		panic(err)
	}
	defer source.Close()

	destination, err := os.Create("static/img/favicon.png")
	if err != nil {
		panic(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	c.JSON(http.StatusOK, gin.H{})
}
