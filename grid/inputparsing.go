package grid

import (
	"errors"
	"pathfinder/data"
	"slices"
	"strconv"
	"strings"
)

var RunningID int = 1

func isValidStationName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for _, ch := range name {
		if !((ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_') {
			return false
		}
	}
	return true
}

// Function to extract valid characters if the row is a comment
func TrimLines(s_slice []string) []string {

	s_out := []string{}
	for _, line := range s_slice {
		line_with_text := ""

		for i := 0; i < len(line); i++ {
			if line[i] < 33 {
				continue
			}
			// Skip all after # as it's a comment
			if line[i] == '#' {
				break
			}

			line_with_text += string(line[i])
		}

		if line_with_text != "" {
			s_out = append(s_out, line_with_text)
		}
	}

	return s_out
}

// Simple logic for extracting the stations and connections from trimmed list.
func ExtractStationsConnections(s_slice []string) (stations, connections []string, err error) {
	hasStationsHeader := false
	hasConnectionsHeader := false

	for _, line := range s_slice {
		if strings.EqualFold(line, "stations:") {
			hasStationsHeader = true
			continue
		}
		if strings.EqualFold(line, "connections:") {
			hasConnectionsHeader = true
			continue
		}

		if strings.Count(line, ",") == 2 {
			stations = append(stations, line)
			continue
		}
		if strings.Count(line, "-") == 1 {
			connections = append(connections, line)
			continue
		}
	}

	if !hasStationsHeader {
		return nil, nil, errors.New("Missing 'stations:' section header")
	}

	if !hasConnectionsHeader {
		return nil, nil, errors.New("Missing 'connections:' section header")
	}

	if len(stations) < 2 || len(connections) < 2 {
		return nil, nil, errors.New("Not enough stations (min 2) or connections (min 2)")
	}

	if len(stations) > 10000  {
		return nil, nil, errors.New("Too many stations (max 10000)")
	}
	return
}

func GetStationItems(s string) (string, [2]int, error) {

	splitted := strings.Split(s, ",")

	if len(splitted) != 3 {
		return "", [2]int{}, errors.New("Too many commas for a station. " + s)
	}

	name := splitted[0]

	if name == "" {
		return "", [2]int{}, errors.New("Station name cannot be empty. (index 0)" + s)
	}

	if !isValidStationName(name) {
		return "", [2]int{}, errors.New("Invalid station name format: " + name)
	}

	coord1, err := strconv.Atoi(splitted[1])
	if err != nil || coord1 < 0 {
		return "", [2]int{}, errors.Join(err, errors.New("Cannot parse or invalid X coordinate: "+s))
	}

	if data.MAX_X_COORDINATE < coord1 {
		data.MAX_X_COORDINATE = coord1
	} else if data.MIN_X_COORDINATE > coord1 {
		data.MIN_X_COORDINATE = coord1
	}

	coord2, err := strconv.Atoi(splitted[2])
	if err != nil || coord2 < 0 {
		return "", [2]int{}, errors.Join(err, errors.New("Cannot parse or invalid Y coordinate: "+s))
	}

	if data.MAX_Y_COORDINATE < coord2 {
		data.MAX_Y_COORDINATE = coord2
	} else if data.MIN_Y_COORDINATE > coord2 {
		data.MIN_Y_COORDINATE = coord2
	}

	return name, [2]int{coord1, coord2}, nil

}

func BuildStation(s string) (data.Station, error) {

	name, coords, err := GetStationItems(s)

	if err != nil {
		return data.Station{}, err
	}

	_, locationExists := data.CoordsMap[coords]

	if locationExists {
		return data.Station{}, errors.New("Duplicate station by coordinates. " + s)
	}

	station := data.Station{Coordinates: coords, Name: name, ID: RunningID}
	RunningID += 1

	data.CoordsMap[coords] = &station

	return station, nil
}

func CreateConnection(s string) error {

	stations := strings.Split(s, "-")

	if len(stations) != 2 || stations[0] == "" || stations[1] == "" {
		return errors.New("Connection format invalid. Should have station1-station2. " + s)
	}

	station1, found := data.StationsMap[stations[0]]

	if !isValidStationName(stations[0]) {
		return errors.New("Invalid station name format in connections: " + stations[0])
	}

	if !found {
		return errors.New("Station does not exist in the grid: " + stations[0])
	}

	station2, found := data.StationsMap[stations[1]]

	if !isValidStationName(stations[1]) {
		return errors.New("Invalid station name format in connections: " + stations[1])
	}

	if !found {
		return errors.New("Station does not exist in the grid: " + stations[1])
	}

	// Check if the connection exists
	if slices.Contains(station1.Connections, station2) || slices.Contains(station2.Connections, station1) {
		return errors.New("Route already declared: " + s)
	}

	station1.Connections = append(station1.Connections, station2)
	station2.Connections = append(station2.Connections, station1)

	return nil
}
