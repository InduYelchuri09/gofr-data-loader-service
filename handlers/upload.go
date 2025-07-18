package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/tkrajina/gofr"
	"your-module-name/utils"
)

func UploadHandler(ctx *gofr.Context) (interface{}, error) {
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	defer file.Close()

	fileExt := strings.ToLower(fileHeader.Filename[strings.LastIndex(fileHeader.Filename, ".")+1:])
	var preview interface{}

	switch fileExt {
	case "csv":
		preview, err = utils.ParseCSV(file)
	case "json":
		preview, err = utils.ParseJSON(file)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", fileExt)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}

	return gofr.Response{
		Data: map[string]interface{}{
			"preview": preview,
		},
	}, nil
}
