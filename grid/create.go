// Essentially stations are the nodes and connections are pointers to stations.
// Here grid means a linked list structure.

// [nodes out]

package grid

import (
	"errors"
)

func InitGrid(input_path string) error {
	path, err := GetAbsPath(input_path)

	if err != nil {
		return err
	}

	strings, err := ReadFileToStringRows(path)

	if err != nil {
		return err
	}

	stations, connections, err := ExtractStationsConnections(TrimLines(strings))

	if err != nil {
		return err
	}

	for _, station_data := range stations {

		valid, err := BuildStation(station_data)

		if err != nil && ALLOW_CORRUPT_DATA {
			continue

		} else if err != nil {
			return errors.Join(err, errors.New("Invalid data for station: "+station_data))
		}

		_, station_exists := StationsMap[valid.Name]

		if station_exists && !ALLOW_CORRUPT_DATA {
			return errors.New("Duplicate station name: " + valid.Name)
		}

		StationsMap[valid.Name] = &valid

	}

	for _, connection := range connections {

		err := CreateConnection(connection)

		if err != nil && ALLOW_CORRUPT_DATA {
			continue

		} else if err != nil {
			return err
		}
	}

	return nil
}
