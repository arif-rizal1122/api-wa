package controller

import (
	"api-wa/app/domain/types/request"
	"api-wa/app/helper"
	"api-wa/app/usecase"
	"os"
	"path/filepath"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	usecase usecase.StatusUsecase
}

func NewStatusController(Usecase usecase.StatusUsecase) *StatusController {
	return &StatusController{usecase: Usecase}
}



func (c *StatusController) CreateStatus(ctx *gin.Context) {
    var input request.RequestCreateStatus
    // Tangani upload file
    file, errFile := ctx.FormFile("picture")
    if errFile != nil {
        ctx.String(http.StatusBadRequest, "Bad Request: %v", errFile)
        return
    }
    saveDir := filepath.Join("public", "img")
    if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
        ctx.String(http.StatusInternalServerError, "Gagal membuat direktori: %v", err)
        return
    }
    savePath := filepath.Join(saveDir, file.Filename)
    if err := ctx.SaveUploadedFile(file, savePath); err != nil {
        ctx.String(http.StatusInternalServerError, "Gagal menyimpan file: %v", err)
        return
    }
    if err := ctx.ShouldBind(&input); err != nil {
        response := helper.NewErrorsResponse("BAD_REQUEST", http.StatusUnprocessableEntity, err.Error())
        ctx.JSON(http.StatusUnprocessableEntity, response)
        return
    }
    input.Picture = savePath
    userId, _ := helper.AuthUserID(ctx)
    data, err := c.usecase.CreateStatus(input, userId)
    if err != nil {
        response := helper.NewErrorsResponse("ERROR", http.StatusInternalServerError, "INVALID SERVER ERROR")
        ctx.JSON(http.StatusInternalServerError, response)
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": data.Message, "data": data.Data})
}




/*
*fix
*
 */

func (c *StatusController) FindById(ctx *gin.Context) {
	userId, errAuth := helper.AuthUserID(ctx)
	if errAuth != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusUnauthorized, "PARAMETER INVALID")
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusUnprocessableEntity, "PARAMETER INVALID")
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	status, err := c.usecase.FindById(id)
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusNotFound, "NOT FOUND")
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	userdID := status.StatusResponseFind.UserId
	if userdID != userId {
		errRes := helper.NewErrorsResponse("ERROR", http.StatusUnauthorized, "ACCESS DENIED")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	ctx.JSON(http.StatusOK, status)
}

/*
*fix
*
 */

func (c *StatusController) FindAll(ctx *gin.Context) {
	userId, _ := helper.AuthUserID(ctx)
	status, err := c.usecase.FindAll(userId)
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusInternalServerError, "INVALID SERVER ERROR")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	for _, item := range status.StatusResponseFinds {
		if item.UserId == userId {
			ctx.JSON(http.StatusOK, status)
			return
		}
	}
	errRes := helper.NewErrorsResponse("INTERNAL SERVER ERROR", http.StatusForbidden, "ACCESS DENIED")
	ctx.JSON(http.StatusForbidden, errRes)
}

/*
*fix
*
 */

func (c *StatusController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusUnprocessableEntity, "PARAMETER INVALID")
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	statusId, errFind := c.usecase.FindById(id)
	if errFind != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusNotFound, "NOT_FOUND")
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	idStr, errAuth := helper.AuthUserID(ctx)
	if errAuth != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusUnauthorized, "PARAMETER INVALID")
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	if statusId.StatusResponseFind.UserId != idStr {
		errRes := helper.NewErrorsResponse("ERROR", http.StatusUnauthorized, "ACCESS DENIED")
		ctx.JSON(http.StatusUnauthorized, errRes)
		return
	}
	errDelete := c.usecase.Delete(id)
	if errDelete != nil {
		response := helper.NewErrorsResponse("ERROR", http.StatusInternalServerError, "INVALID SERVER ERROR")
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "status deleted successfully"})
}

/*
*fix
*
 */

 func (c *StatusController) Update(ctx *gin.Context) {
    var input request.RequestUpdateStatus

    file, errFile := ctx.FormFile("picture")
    if errFile == nil {
        saveDir := filepath.Join("public", "img")
        if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
            ctx.String(http.StatusInternalServerError, "Gagal membuat direktori: %v", err)
            return
        }
        savePath := filepath.Join(saveDir, file.Filename)
        if err := ctx.SaveUploadedFile(file, savePath); err != nil {
            ctx.String(http.StatusInternalServerError, "Gagal menyimpan file: %v", err)
            return
        }
        input.Picture = savePath
    }
    if err := ctx.ShouldBind(&input); err != nil {
        response := helper.NewErrorsResponse("VALIDATION ERROR", http.StatusUnprocessableEntity, err.Error())
        ctx.JSON(http.StatusUnprocessableEntity, response)
        return
    }
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        response := helper.NewErrorsResponse("ERROR", http.StatusUnprocessableEntity, "PARAMETER INVALID")
        ctx.JSON(http.StatusUnprocessableEntity, response)
        return
    }
    userId, errFind := c.usecase.FindById(id)
    if errFind != nil {
        response := helper.NewErrorsResponse("ERROR", http.StatusNotFound, "NOT_FOUND")
        ctx.JSON(http.StatusNotFound, response)
        return
    }
    idStr, _ := helper.AuthUserID(ctx)
    if userId.StatusResponseFind.UserId != idStr {
        errRes := helper.NewErrorsResponse("BAD_REQUEST", http.StatusUnauthorized, "ACCESS DENIED")
        ctx.JSON(http.StatusUnauthorized, errRes)
        return
    }
    err = c.usecase.Update(id, input)
    if err != nil {
        response := helper.NewErrorsResponse("ERROR", http.StatusInternalServerError, "INVALID SERVER ERROR")
        ctx.JSON(http.StatusInternalServerError, response)
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Updated success success"})
}

