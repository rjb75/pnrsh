package pnr

func convertPnrResponse(response GetPNRResponse, pnr *PNR) {
	pnr.Pnr = response.Pnr
}

func convertFlights(response GetPNRResponse, pnr *PNR) {
	for _, od := range response.OriginDestinations {
		for _, segment := range od.Segments {
			pnr.Flights = append(pnr.Flights, Flight{
				OriginAirportCode:      od.OriginAirportCode,
				DestinationAirportCode: od.DestinationAirportCode,
				OperatingAirlineCode:   segment.OperatingAirlineCode,
				MarketingAirlineCode:   segment.MarketingAirlineCode,
				FlightNumber:           segment.FlightNumber,
				ScheduledDeparture:     segment.DepartureDateTime,
				ScheduledArrival:       segment.ArrivalDateTime,
				Cabin:                  segment.Legs[0].Cabin.Name,
				Status:                 segment.Legs[0].Status,
				Duration:               segment.DurationMinutes,
				SabreCode:              segment.Legs[0].Cabin.SabreCode,
			})
		}
	}
}
