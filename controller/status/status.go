package controller

import (
	"fmt"
	"github.com/diegoasencio96/go-microservice/config"
	"github.com/diegoasencio96/go-microservice/model/status"
	"github.com/diegoasencio96/go-microservice/server/general"
	"github.com/labstack/echo/v4"
	"net/http"
	"os/exec"
	"time"
)


func StatusController(c echo.Context) error {
	ctx := general.GetContextFromEcho(c)
	general.ServerLog(ctx).Infof("%s", "controller: StatusController")
	out, _ := exec.Command(
		"sh",
		"-c",
		`{ git log --pretty=format:'%h - %ae - %ad' -n 1;
            echo '('$(git branch | grep \\* | cut -d ' ' -f2)')';}
            | tr '\n' ' '`,
	).Output()

	return c.JSON(http.StatusOK, &status.Status{
		Name:    config.Settings.ProjectName,
		Git:     fmt.Sprintf("%s", out),
		Version: config.Settings.ProjectVersion,
		Date:    time.Now(),
	})
}
