package v1

import (
	"mime/multipart"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func AddMaterial(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddMaterialForm
	)

	userName, exists := c.Get("currentUser")
	if !exists {
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_NOT_FOUND, nil)
		return
	}

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
		UserName: userName.(string),
	}

	if err := materialService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_MATERIAL_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary		Get all materials
// @Produce 	json
// @Success		200 {object} app.Response
// @Failure		500 {object} app.Response
// @Router		/api/v1/materials [get]
func GetMaterials(c *gin.Context) {
	appG := app.Gin{C: c}
	materials, err := material_service.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MATERIALS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"materials": materials,
	})
}

// @Summary Get a single material
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/materials/{id} [get]
func GetMaterial(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	materialService := material_service.Material{ID: id}
	exists, err := materialService.ExistsByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MATERIAL_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MATERIAL, nil)
		return
	}
}
