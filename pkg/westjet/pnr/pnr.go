package pnr

func Retrieve(lastName, confirmationCode string) (PNR, error) {
	res, err := performRequest(lastName, confirmationCode)

	if err != nil {
		return PNR{}, err
	}

	return convertResponse(res), nil
}
