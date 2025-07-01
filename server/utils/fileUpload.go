// package utils

// import (
// 	"fmt"
// 	"os"
// 	"time"
// 	"path/filepath"

// 	"github.com/gin-gonic/gin"
// )

// func UploadFile(c *gin.Context) (string, error) {
// 	// Form field name should be 'file' in Postman
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return "", err
// 	}

// 	// Create folder if not exists
// 	err = os.MkdirAll("uploads", os.ModePerm)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Create simple & safe file name
// 	fileName := fmt.Sprint(time.Now().Unix()) + "_" + file.Filename
// 	filePath := filepath.Join("uploads", fileName)

// 	// Save file
// 	err = c.SaveUploadedFile(file, filePath)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Return file path or just filename
// 	return filePath, nil
// }

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c*gin.Context)(string, error){
	// create Formfile
	file,err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	// create uploads folder
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return "", err
	}


	// create filename nd filepath
	fileName := fmt.Sprint(time.Now().Unix())+ "_" + file.Filename
	filePath := filepath.Join("uploads", fileName)

	//save the changes
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		return "", err
	}

	//return file path nd nil
	return filePath, nil
}