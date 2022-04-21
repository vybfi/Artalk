package http

import (
	"github.com/labstack/echo/v4"
)

func (a *action) AdminSettingGet(c echo.Context) error {
	if isOK, resp := AdminOnly(c); !isOK {
		return resp
	}

	return RespSuccess(c)
}

type ParamsAdminSettingSave struct {
	ID   uint   `mapstructure:"id" param:"required"`
	Name string `mapstructure:"name"`
	Url  string `mapstructure:"url"`
}

func (a *action) AdminSettingSave(c echo.Context) error {
	if isOK, resp := AdminOnly(c); !isOK {
		return resp
	}

	var p ParamsAdminSettingSave
	if isOK, resp := ParamsDecode(c, ParamsAdminSettingSave{}, &p); !isOK {
		return resp
	}

	return RespSuccess(c)
}
