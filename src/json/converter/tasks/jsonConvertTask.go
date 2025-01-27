package jsonConverter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	lists "github.com/RyuseiNomi/delay_reporter_lm/src/routeLists"
)

func ReadAllTrainRouteFromTsv() (lists.TrainRoutes, error) {

	// Read train route master tsv file
	tsv, err := os.Open("/tmp/tsv/train.tsv")
	if err != nil {
		return nil, fmt.Errorf("TSV Read Error:  %+v\n", err)
	}
	defer tsv.Close()

	render := csv.NewReader(tsv)
	render.Comma = '\t' // change delimiter

	// Append all of train route
	var trainRoutes lists.TrainRoutes
	for {
		record, err := render.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("TSV Render Error:  %+v\n", err)
		}
		trainRoute := lists.TrainRoute{
			Company:   record[0],
			RouteName: record[1],
			Region:    record[2],
			IsValid:   record[3],
		}
		trainRoutes = append(trainRoutes, trainRoute)
	}

	return trainRoutes, nil
}

func ConvertDelayLists(fetchedDelayList []lists.FetchedDelayList, trainRoutes lists.TrainRoutes) lists.ConvertedDelayLists {
	// determine delay status each train route
	var convertedDelayLists lists.ConvertedDelayLists
	for _, trainRoute := range trainRoutes {
		var status = "○"
		for _, delayRoute := range fetchedDelayList {
			if delayRoute.Name == trainRoute.RouteName {
				status = "△"
			}
		}
		convertedDelayList := lists.ConvertedDelayList{
			Name:    trainRoute.RouteName,
			Company: trainRoute.Company,
			Region:  trainRoute.Region,
			Status:  status,
			Source:  "鉄道com RSS",
		}
		convertedDelayLists = append(convertedDelayLists, convertedDelayList)
	}

	return convertedDelayLists
}
