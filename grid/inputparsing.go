package grid

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

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

	for _, line := range s_slice {
		if strings.Count(line, ",") == 2 {
			stations = append(stations, line)
			continue
		}
		if strings.Count(line, "-") == 1 {
			connections = append(connections, line)
			continue
		}
	}

	if len(stations) < 2 || len(connections) < 2 {
		return nil, nil, errors.New("Not enough stations (min 2) or connections (min 2)")
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

	coord1, err := strconv.Atoi(splitted[1])

	if err != nil {
		return "", [2]int{}, errors.Join(err, errors.New("Cannot parse X coordinate (index 1)"+s))
	}

	coord2, err := strconv.Atoi(splitted[2])

	if err != nil {
		return "", [2]int{}, errors.Join(err, errors.New("Cannot parse X coordinate (index 2)"+s))
	}

	return name, [2]int{coord1, coord2}, nil

}

func BuildStation(s string) (Station, error) {

	name, coords, err := GetStationItems(s)

	if err != nil {
		return Station{}, err
	}

	_, locationExists := CoordsMap[coords]

	if locationExists {
		return Station{}, errors.New("Duplicate station by coordinates. " + s)
	}

	station := Station{Coordinates: coords, Name: name}

	CoordsMap[coords] = &station

	return station, nil
}

func CreateConnection(s string) error {

	stations := strings.Split(s, "-")

	if len(stations) != 2 || stations[0] == "" || stations[1] == "" {
		return errors.New("Connection format invalid. Should have station1-station2. " + s)
	}

	station1, found := StationsMap[stations[0]]

	if !found {
		return errors.New("Station does not exist in the grid: " + stations[0])
	}

	station2, found := StationsMap[stations[1]]

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
