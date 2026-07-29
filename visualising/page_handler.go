package visualising

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"pathfinder/data"
)

var TEMPLATE = template.Must(template.ParseFiles("visualising/index.html"))

func Viewbox() string {
	minX := float64(data.MIN_X_COORDINATE)
	maxX := float64(data.MAX_X_COORDINATE)
	minY := float64(data.MIN_Y_COORDINATE)
	maxY := float64(data.MAX_Y_COORDINATE)

	rangeX := maxX - minX
	rangeY := maxY - minY

	if rangeX <= 0 {
		rangeX = 5
	}
	if rangeY <= 0 {
		rangeY = 5
	}

	// Calculate proportional padding (15% margin around the graph edges)
	padX := rangeX * 0.15
	if padX < 1.5 {
		padX = 1.5
	}
	padY := rangeY * 0.15
	if padY < 1.5 {
		padY = 1.5
	}

	vbMinX := minX - padX
	vbMinY := minY - padY
	vbWidth := rangeX + (2 * padX)
	vbHeight := rangeY + (2 * padY)

	return fmt.Sprintf("%.2f %.2f %.2f %.2f", vbMinX, vbMinY, vbWidth, vbHeight)
}

func Roothandler(start, end *data.Station, fastest_route *[]data.Station, turns [][]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		turnsJSON, _ := json.Marshal(turns)

		coordsMap := make(map[string][2]int)
		for name, st := range data.StationsMap {
			coordsMap[name] = st.Coordinates
		}
		stationsJSON, _ := json.Marshal(coordsMap)

		pageData := struct {
			data.VisualisingData
			TurnsJSON    template.JS
			StationsJSON template.JS
		}{
			VisualisingData: data.VisualisingData{
				Graph:             &data.StationsMap,
				ViewBox:           Viewbox(),
				Start:             start,
				End:               end,
				Fastest_route_ids: fastest_route,
				Turns:             turns,
			},
			TurnsJSON:    template.JS(turnsJSON),
			StationsJSON: template.JS(stationsJSON),
		}

		err := TEMPLATE.Execute(w, pageData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
