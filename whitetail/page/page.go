// pages.go

package page

import (
	// "log"

	"encoding/json"
	"fmt"
	"net/http"
	"whitetail/config"
	"whitetail/constants"
	"whitetail/logger"
	"whitetail/operation"
	"whitetail/utils"

	"github.com/gin-gonic/gin"
)

func RedirectIndexPage(c *gin.Context) {
	c.Redirect(301, config.Config.BasePath+"/ui/home")
}

func ShowHomePage(c *gin.Context) {
	dashboards := []string{}
	for name := range operation.Operations.Dashboards {
		dashboards = append(dashboards, name)
	}

	render(c, gin.H{
		"title":      "Home",
		"location":   "Home",
		"dashboards": dashboards,
		"version":    constants.VERSION,
	},
		"index.html")
}

func ShowQueryPage(c *gin.Context) {
	dashboards := []string{}
	for name := range operation.Operations.Dashboards {
		dashboards = append(dashboards, name)
	}

	render(c, gin.H{
		"title":      "Logs",
		"location":   "Query",
		"dashboards": dashboards,
		"version":    constants.VERSION,
	},
		"query.html",
	)
}

func ShowDashboardPage(c *gin.Context) {
	n := c.Param("name")
	d := operation.Operations.Dashboards[n]
	// graphs := []GraphObj{}
	// tableDef := []interface{}{}
	// ps := d.Panels
	// sort.Slice(ps,
	// 	func(i, j int) bool {
	// 		return ps[i].XCoord < ps[j].YCoord
	// 	})
	// sort.Slice(ps,
	// 	func(i, j int) bool {
	// 		return ps[i].YCoord < ps[j].YCoord
	// 	})

	// panelDef := []interface{}{}
	// currentRow := 0
	// row := map[string][]interface{}{[]interface{}{}}

	scripts := []string{}

	panels := [][]interface{}{}
	maxY := -1
	maxX := -1
	for _, p := range d.Panels {
		logger.Tracef("", "%v", p)
		if p.YCoord+p.RowSpan-1 > maxY {
			maxY = p.YCoord + p.RowSpan - 1
		}
		if p.XCoord+p.ColSpan-1 > maxX {
			maxX = p.XCoord + p.ColSpan - 1
		}
		if p.Kind == "button" {
			script := fmt.Sprintf("<script>\n%s\n</script>", p.JS)
			scripts = append(scripts, script)
		}
		// b, err := json.Marshal(p)
		// if err != nil {
		// 	utils.Error(err, c, http.StatusServiceUnavailable)
		// 	return
		// }
		// var pp map[string]interface{}
		// if err := json.Unmarshal(b, &pp); err != nil {
		// 	utils.Error(err, c, http.StatusServiceUnavailable)
		// 	return
		// }
		// pp["string_def"] = string(b)
		// panelDef = append(panelDef, pp)
	}

	logger.Tracef("", "max x: %d, max y: %d", maxX, maxY)

	// panels := []interface{}{}

	// tableDef = append(tableDef, row)
	for y := 0; y < maxY+1; y++ {
		temp := []interface{}{}
		for x := 0; x < maxX+1; x++ {
			empty := map[string]string{"kind": "empty"}
			temp = append(temp, empty)
		}
		panels = append(panels, temp)
	}

	for _, p := range d.Panels {
		logger.Tracef("", "%v", p)

		for rr := 0; rr < p.RowSpan; rr++ {
			for cc := 0; cc < p.ColSpan; cc++ {
				row := p.YCoord + rr
				col := p.XCoord + cc
				if rr == 0 && cc == 0 {
					b, err := json.Marshal(p)
					if err != nil {
						utils.Error(err, c, http.StatusServiceUnavailable)
						return
					}
					var pp map[string]interface{}
					if err := json.Unmarshal(b, &pp); err != nil {
						utils.Error(err, c, http.StatusServiceUnavailable)
						return
					}
					pp["string_def"] = string(b)
					panels[row][col] = pp
					logger.Tracef("", "Adding %v to panels", p)
					continue
				}
				null := map[string]string{"kind": "null"}
				panels[row][col] = null
			}
		}
	}
	// 		found := false
	// 		for _, panel := range panelDef {
	// 			if int(panel.(map[string]interface{})["x_coord"].(float64)) == x {
	// 				panels = append(panels, panel)
	// 				found = true
	// 				break
	// 			}
	// 		}
	// 		if !found {
	// 			null := map[string]string{"kind": "null"}
	// 			panels = append(panels, null)
	// 		}
	// 	}
	// 	rows = append(rows, float64(y))
	// }

	dashboards := []string{}
	for name := range operation.Operations.Dashboards {
		dashboards = append(dashboards, name)
	}

	b, _ := json.MarshalIndent(panels, "", "    ")
	logger.Debugf("", "Sending panels: %s", string(b))

	logger.Debugf("", "Sending dashboard data: %v", gin.H{
		"title":      "Dashboard",
		"location":   fmt.Sprintf("Dashboard / %s", n),
		"dashboards": dashboards,
		"panels":     panels,
		"name":       n,
		"version":    constants.VERSION,
		"scripts":    scripts,
	})

	render(c, gin.H{
		"title":      "Dashboard",
		"location":   fmt.Sprintf("Dashboard / %s", n),
		"dashboards": dashboards,
		"panels":     panels,
		"name":       n,
		"version":    constants.VERSION,
		"scripts":    scripts,
	},
		"dashboard.html",
	)
}

// func ShowSettingsPage(c *gin.Context) {
// 	// Render the logs-selection.html page
// 	render(c, gin.H{
// 		"title":           "Settings",
// 		"location":        "Settings",
// 		"primary_color":   config.Config.Branding.PrimaryColor,
// 		"secondary_color": config.Config.Branding.SecondaryColor,
// 		"tertiary_color":  config.Config.Branding.TertiaryColor,
// 		"INFO_color":      config.Config.Branding.INFOColor,
// 		"WARN_color":      config.Config.Branding.WARNColor,
// 		"DEBUG_color":     config.Config.Branding.DEBUGColor,
// 		"TRACE_color":     config.Config.Branding.TRACEColor,
// 		"ERROR_color":     config.Config.Branding.ERRORColor},
// 		"page.settings.html")
// }

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
