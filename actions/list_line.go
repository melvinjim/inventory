package actions

import (
	"inventory/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func ListInformation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	lines := []models.Lines{}
	err := tx.All(&lines)
	if err != nil {
		return err
	}

	c.Set("lines", lines)

	return c.Render(http.StatusOK, r.HTML("list_lines/index.plush.html"))
}

func InfoLine(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	lines := models.Lines{}
	lineID := c.Param("user_id")
	err := tx.Find(&lines, lineID)
	if err != nil {
		return err
	}

	c.Set("lines", lines)

	return c.Render(http.StatusOK, r.HTML("list_lines/info_line.plush.html"))
}
