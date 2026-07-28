package visualising

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"pathfinder/data"
)

var TEMPLATE = template.Must(template.ParseFiles("visualising/index.html"))

func Viewbox(padding int) string {
	var minX, minY, width, height int

	minX = data.MIN_X_COORDINATE - padding
	minY = data.MIN_Y_COORDINATE - padding
	width = (data.MAX_X_COORDINATE - minX) + padding
	height = (data.MAX_Y_COORDINATE - minY) + padding

	if width < 60 {
		width = 60
	}
	if height < 60 {
		height = 60
	}

	return fmt.Sprintf("%d %d %d %d", minX, minY, width, height)
}

func Roothandler(start, end *data.Station, fastest_route *[]data.Station, turns [][]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		turnsJSON, _ := json.Marshal(turns)

		// Map station names to coordinates for JS train positioning
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
				ViewBox:           Viewbox(25), // Increased from 10 to 25 for spacing
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
