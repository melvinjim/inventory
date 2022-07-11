package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func ListInformation(c buffalo.Context) error {

	return c.Render(http.StatusOK, r.HTML("associated/index.plush.html"))
}
