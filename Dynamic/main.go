package main

import (
	"dynamic/controller"
	"dynamic/helper"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

/*
	|--------------------------------------------------|
	| Dynamic WebServer                                |
	| -------------------------------------------------|
	| Version: 1.0.0                                   |
	| Authors:                                         |
	| Diogo Julio       - RM 553837                    |
	| Jonata Rafael     - RM 552939                    |
	| Matheus Zottis    - RM 94119                     |
	| Victor Didoff     - RM 552965                    |
	| Vinicius Silva    - RM 553240                    |
	|--------------------------------------------------|
*/

func main() {
	/* Import Config Init */
	var configFile string
	flag.StringVar(&configFile, "config", "./config.json", "Path to the configuration file")
	flag.Parse()
	helper.DefineConfigPath(configFile)
	/* Import Config End */

	/* Config WebServer */
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	/* Import WebServer Template Files and Static Files Init */
	files, err := loadTemplates(helper.GetConfig().Template)
	if err != nil {
		fmt.Println(err)
	}
	r.LoadHTMLFiles(files...)
	r.Static("/static", helper.GetConfig().Static)
	/* Import WebServer Template Files and Static Files End */

	/* Import WebServer Middleware Init */
	dynamic := r.Group("/dynamic/")
	dynamic.Use(basic())
	/* Import WebServer Middleware End */

	/* Config WebServer Routes Init */
	r.NoRoute(controller.NotFound)
	r.GET("/", controller.Index)
	dynamic.GET("/", controller.GetDynamic)
	/* Config WebServer Routes End */

	/* Run WebServer */
	err = r.Run(fmt.Sprintf(":%d", helper.GetConfig().Ambient.Port))
	if err != nil {
		fmt.Println(err)
	}
}

// basic is a middleware that does nothing.
func basic() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// loadTemplates loads all template files from the specified root directory.
func loadTemplates(root string) (files []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			if path != root {
				loadTemplates(path)
			}
		} else {
			files = append(files, path)
		}
		return err
	})
	return files, err
}
