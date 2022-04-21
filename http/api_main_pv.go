package http

import (
	"github.com/ArtalkJS/ArtalkGo/lib"
	"github.com/ArtalkJS/ArtalkGo/model"
	"github.com/labstack/echo/v4"
)

type ParamsPV struct {
	PageKey   string `mapstructure:"page_key" param:"required"`
	PageTitle string `mapstructure:"page_title"`

	SiteName string `mapstructure:"site_name"`

	SiteID  uint
	SiteAll bool
}

func (a *action) PV(c echo.Context) error {
	var p ParamsPV
	if isOK, resp := ParamsDecode(c, ParamsPV{}, &p); !isOK {
		return resp
	}

	// find site
	if isOK, resp := CheckSite(c, &p.SiteName, &p.SiteID, &p.SiteAll); !isOK {
		return resp
	}

	// find page
	page := model.FindCreatePage(p.PageKey, p.PageTitle, p.SiteName)

	// ip := c.RealIP()
	// ua := c.Request().UserAgent()

	page.PV++
	lib.DB.Save(&page)

	return RespData(c, Map{
		"pv": page.PV,
	})
}
