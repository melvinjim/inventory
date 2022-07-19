package actions

import (
	"fmt"
	"inventory/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func ListInformation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	lines := []models.Lines{}
	line := &models.Lines{}

	FilterInformation := c.Param("filter_information")

	q := tx.Q()

	if FilterInformation != "" {
		n := fmt.Sprintf("%%%v%%", FilterInformation)
		q = tx.Where("carrier like ? OR phone_line like ? OR end_contract_date like ?", n, n, n)
	}

	err := q.All(&lines)
	if err != nil {
		return err
	}

	c.Set("count", len(lines))
	c.Set("lines", lines)
	c.Set("line", line)

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

func Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	lines := models.Lines{}
	linesID := c.Param("user_id")

	err := tx.Find(&lines, linesID)
	if err != nil {
		return err
	}

	c.Set("lines", lines)

	return c.Render(http.StatusOK, r.HTML("list_lines/edit.plush.html"))
}

func EditLine(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	lines := models.Lines{}
	linesID := c.Param("user_id")

	err := tx.Find(&lines, linesID)
	if err != nil {
		return err
	}

	if err := c.Bind(&lines); err != nil {
		return err
	}

	err = tx.Update(&lines)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func DeleteLine(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	line := models.Lines{}
	lineID := c.Param("user_id")

	err := tx.Find(&line, lineID)
	if err != nil {
		return err
	}

	err = tx.Destroy(&line)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
