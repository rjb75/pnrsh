package pnr

type PNR struct {
	Pnr     string
	Flights []Flight
}

type Flight struct {
	OriginAirportCode      string
	DestinationAirportCode string
	OperatingAirlineCode   string
	MarketingAirlineCode   string
	FlightNumber           string
	ScheduledDeparture     string
	ScheduledArrival       string
	Cabin                  string
	Status                 string
	Duration               int
	SabreCode              string
}

type GetPNRResponse struct {
	PseudoCityCode string `json:"aaaPseudoCityCode"`
	Eligibility    struct {
		Change struct {
			Eligible  bool   `json:"eligible"`
			ErrorCode string `json:"errorCode,omitempty"`
			RuleName  string `json:"ruleName,omitempty"`
		} `json:"CHG"`
		Cancel struct {
			Eligible  bool   `json:"eligible"`
			ErrorCode string `json:"errorCode,omitempty"`
			RuleName  string `json:"ruleName,omitempty"`
		} `json:"CXL"`
		Hold struct {
			Eligible  bool   `json:"eligible"`
			ErrorCode string `json:"errorCode,omitempty"`
			RuleName  string `json:"ruleName,omitempty"`
		} `json:"HOLD"`
		IChange struct {
			Eligible  bool   `json:"eligible"`
			ErrorCode string `json:"errorCode,omitempty"`
			RuleName  string `json:"ruleName,omitempty"`
		} `json:"ICHG"`
		ICancel struct {
			Eligible  bool   `json:"eligible"`
			ErrorCode string `json:"errorCode,omitempty"`
			RuleName  string `json:"ruleName,omitempty"`
		} `json:"ICXL"`
		Info struct {
			Eligible  bool   `json:"eligible"`
			ErrorCode string `json:"errorCode,omitempty"`
			RuleName  string `json:"ruleName,omitempty"`
		} `json:"Info"`
		Seat struct {
			Eligible  bool   `json:"eligible"`
			ErrorCode string `json:"errorCode,omitempty"`
			RuleName  string `json:"ruleName,omitempty"`
		} `json:"Seat"`
	} `json:"eligibility"`
	FullyUnconfirmed bool `json:"fullyUnconfirmed"`
	GroupBooking     bool `json:"groupBooking"`
	Guests           []struct {
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		NameId        string `json:"nameId"`
		PassengerType string `json:"passengerType"`
		Title         string `json:"title"`
		WestjetId     string `json:"westjetId,omitempty"`
	}
	OriginDestinations []struct {
		ArrivalDateTime string `json:"arrivalDateTime"`
		CarryOnBags     []struct {
			Allowed     bool   `json:"allowed"`
			GuestNameId string `json:"guestNameId"`
		} `json:"carryOnBags"`
		DepartureDateTime      string `json:"departureDateTime"`
		DestinationAirportCode string `json:"destinationAirportCode"`
		DurationMinutes        int    `json:"durationMinutes"`
		OriginAirportCode      string `json:"originAirportCode"`
		Segments               []struct {
			ArrivalDateTime        string `json:"arrivalDateTime"`
			DepartureDateTime      string `json:"departureDateTime"`
			DestinationAirportCode string `json:"destinationAirportCode"`
			DurationMinutes        int    `json:"durationMinutes"`
			FlightNumber           string `json:"flightNumber"`
			ItineraryAirlineType   string `json:"itineraryAirlineType"`
			Layover                bool   `json:"layover"`
			Legs                   []struct {
				ArrivalDateTime string `json:"arrivalDateTime"`
				Cabin           struct {
					Code      string `json:"code"`
					Lang      string `json:"lang"`
					Name      string `json:"name"`
					SabreCode string `json:"sabreCode"`
					ShortName string `json:"shortName"`
				} `json:"cabin"`
				Charter                bool   `json:"charter"`
				DepartureDateTime      string `json:"departureDateTime"`
				DestinationAirportCode string `json:"destinationAirportCode"`
				DurationMinutes        int    `json:"durationMinutes"`
				FlightNumber           string `json:"flightNumber"`
				// TODO: guests checked in array - guestsCheckedIn
				// TODO: irop history transaction - iropHistoryTransaction
				ItineraryAirlineType string `json:"itineraryAirlineType"`
				MarketingAirlineCode string `json:"marketingAirlineCode"`
				MarketingFareClass   string `json:"marketingFareClass"`
				OperatingAirlineCode string `json:"operatingAirlineCode"`
				OperatingAirlineName string `json:"operatingAirlineName"`
				OperatingFareClass   string `json:"operatingFareClass"`
				OriginAirportCode    string `json:"originAirportCode"`
				SabreAirSegmentId    string `json:"sabreAirSegmentId"`
				// TODO: scheduleChangeTransaction
				Seats []struct {
					GuestNameId string   `json:"guestNameId"`
					SeatFlags   []string `json:"seatFlags"`
					SeatNumber  string   `json:"seatNumber"`
				} `json:"seats"`
				// TODO: standby
				Status string `json:"status"`
				Type   string `json:"type"`
			} `json:"legs"`
			MarketingAirlineCode   string `json:"marketingAirlineCode"`
			MarketingFareClass     string `json:"marketingFareClass"`
			OperatingAirlineCode   string `json:"operatingAirlineCode"`
			OperatingAirlineName   string `json:"operatingAirlineName"`
			OperatingFareClass     string `json:"operatingFareClass"`
			OriginAirportCode      string `json:"originAirportCode"`
			StandbyStatusAvailable bool   `json:"standbyStatusAvailable"`
			Type                   string `json:"type"`
		} `json:"segments"`
		StandbyStatusAvailable bool `json:"standbyStatusAvailable"`
	} `json:"originDestinations"`
	Pnr string `json:"pnr"`
	// TODO: priorityCode
	ScheduleChangeType string `json:"scheduleChangeType"`
	// TODO: unconfirmedLegs
}
