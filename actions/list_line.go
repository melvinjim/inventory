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

func AddLine(c buffalo.Context) error {
	line := &models.Lines{}
	c.Set("line", line)

	return c.Render(http.StatusOK, r.HTML("list_lines/add_line.plush.html"))
}

func CreateLine(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	line := models.Lines{}
	if err := c.Bind(&line); err != nil {
		return err
	}

	verrs, err := line.Validate(tx)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("line", line)
		c.Set("errors-phone", verrs.Get("phone_line"))
		c.Set("errors-carrier", verrs.Get("carrier"))
		c.Set("errors-end_contract_date", verrs.Get("end_contract_date"))
		return c.Render(http.StatusOK, r.HTML("list_lines/add_line.plush.html"))
	}

	errCreate := tx.Create(&line)
	if errCreate != nil {
		return errCreate
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
