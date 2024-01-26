// pages.go

package page

import (
	// "log"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"whitetail/config"
	"whitetail/constants"
	"whitetail/logger"
	"whitetail/operation"
	"whitetail/utils"

	"github.com/gin-gonic/gin"
)

type GraphObj struct {
	Observer string `json:"observer"`
	Stream   string `json:"stream"`
	GraphDef string `json:"graph_def"`
}

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
	gs := d.Layout.Graphs
	sort.Slice(gs,
		func(i, j int) bool {
			return gs[i].XCoord < gs[j].YCoord
		})
	sort.Slice(gs,
		func(i, j int) bool {
			return gs[i].YCoord < gs[j].YCoord
		})
	// currentRow := 0
	// row := map[string][]interface{}{[]interface{}{}}
	def := []interface{}{}
	rows := []float64{}
	maxY := -1
	for _, g := range d.Layout.Graphs {
		if g.YCoord > maxY {
			maxY = g.YCoord
		}
		b, err := json.Marshal(g)
		if err != nil {
			utils.Error(err, c, http.StatusServiceUnavailable)
			return
		}
		var gg map[string]interface{}
		if err := json.Unmarshal(b, &gg); err != nil {
			utils.Error(err, c, http.StatusServiceUnavailable)
			return
		}
		gg["string_def"] = string(b)
		def = append(def, gg)

		// gg := GraphObj{
		// 	Observer: g.Observer,
		// 	Stream:   g.Stream,
		// 	GraphDef: string(b),
		// }
		// graphs = append(graphs, gg)
		// if g.YCoord == currentRow {
		// col := map[string]interface{}{
		// 	"observer": g.Observer,
		// 	"stream":   g.Stream,
		// 	"colspan":  g.ColSpan,
		// 	"rowspan":  g.RowSpan,
		// 	"width":    g.Width,
		// 	"height":   g.Height,
		// }
		// row[]
		// 	row["cols"] = append(row["cols"], g)
		// } else {
		// 	tableDef = append(tableDef, row)
		// 	row = nil
		// 	row = map[string][]interface{}{"cols": []interface{}{}}
		// }
	}
	// tableDef = append(tableDef, row)
	for i := 0; i < maxY+1; i++ {
		rows = append(rows, float64(i))
	}

	dashboards := []string{}
	for name := range operation.Operations.Dashboards {
		dashboards = append(dashboards, name)
	}

	logger.Debugf("", "Sending dashboard data: %v", gin.H{
		"title":      "Dashboard",
		"location":   fmt.Sprintf("Dashboard / %s", n),
		"dashboards": dashboards,
		"graphs":     def,
		"rows":       rows,
		"version":    constants.VERSION,
	})

	render(c, gin.H{
		"title":      "Dashboard",
		"location":   fmt.Sprintf("Dashboard / %s", n),
		"dashboards": dashboards,
		"graphs":     def,
		"rows":       rows,
		"version":    constants.VERSION,
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
