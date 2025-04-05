package v1

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vgbhj/SKAT/pkg/app"
	"github.com/vgbhj/SKAT/pkg/e"
	"github.com/vgbhj/SKAT/service/material_service"
)

type AddMaterialForm struct {
	Title    string                `form:"title" valid:"Required;MaxSize(100)"`
	Desc     string                `form:"desc" valid:"Required;MaxSize(255)"`
	fileData *multipart.FileHeader `form:"file" binding:"required"`
	// user_id integer [ref: > users.id]
	// faculty_id integer [ref: > faculty.id]
	// subject_id integer [ref: > subject.id]
	// year_id integer [ref: > year.id]
	// university_id integer [ref: > university.id]
}

// @Summary Add material
// @Produce  json
// @Accept multipart/form-data
// @Param title formData string true "Title"
// @Param desc formData string true "Desc"
// @Param file formData file true "Material File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/materials [post]
func AddMaterial(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddMaterialForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	var err error
	form.fileData, err = c.FormFile("file")
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_INVALID_FILE, nil)
		return
	}

	materialService := material_service.Material{
		Title:    form.Title,
		Desc:     form.Desc,
		FileData: form.fileData,
	}

	if err := materialService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_MATERIAL_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
