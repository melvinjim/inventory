package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func ListAssociated(c buffalo.Context) error {

	return c.Render(http.StatusOK, r.HTML("associated/index.plush.html"))
}
