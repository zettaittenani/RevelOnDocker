package controllers

import (
    "github.com/revel/revel"
    _ "myapp/app/controllers"
    "myapp/app/models"
    "gopkg.in/validator.v2"
)

type ApiV1Users struct {
    ApiV1Controller
}

func (c ApiV1Users) Index() revel.Result {
    users := []models.User{}

    if err := models.DB.Order("id desc").Find(&users).Error; err != nil {
        return c.HandleInternalServerError("Record Find Failure")
    }

    r := Response{users}
    return c.RenderJSON(r)
}

func (c ApiV1Users) Show(id int) revel.Result {
    user := &models.User{}

    if err := models.DB.First(&user, id).Error; err != nil {
        return c.HandleNotFoundError(err.Error())
    }

    r := Response{user}
    return c.RenderJSON(r)
}

func (c ApiV1Users) Create() revel.Result {
    user := &models.User{}

    if err := c.Params.BindJSON(user); err != nil {
        return c.HandleBadRequestError(err.Error())
    }

    if err := validator.Validate(user); err != nil {
        return c.HandleBadRequestError(err.Error())
    }

    if err := models.DB.Create(user).Error; err != nil {
        return c.HandleInternalServerError("Record Create Failure")
    }

    r := Response{user}
    return c.RenderJSON(r)
}

func (c ApiV1Users) Delete(id int) revel.Result {
    user := models.User{}

    if err := models.DB.First(&user, id).Error; err != nil {
        return c.HandleNotFoundError(err.Error())
    }

    if err := models.DB.Delete(&user).Error; err != nil {
	return c.HandleInternalServerError("Record Delete Failure")
    }

    r := Response{"success"}
    return c.RenderJSON(r)
}
