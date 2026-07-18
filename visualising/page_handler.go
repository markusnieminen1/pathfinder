package visualising

import (
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
	width = (data.MAX_X_COORDINATE - (data.MIN_X_COORDINATE - padding)) + 2*padding
	height = (data.MAX_Y_COORDINATE - (data.MIN_Y_COORDINATE - padding)) + 2*padding

	return fmt.Sprintf("%d %d %d %d", minX, minY, width, height)
}

func Roothandler(start, end *data.Station) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := TEMPLATE.Execute(w, data.VisualisingData{Graph: &data.StationsMap, ViewBox: Viewbox(10), Start: start, End: end})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
