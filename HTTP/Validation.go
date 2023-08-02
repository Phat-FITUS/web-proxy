package HTTP

func ValidateHeader(header string) error {
	var err error
	_, err = GetMethod(header)
	if (err != nil) {
		return err
	}

	_, err = GetURL(header)
	if (err != nil) {
		return err
	}

	return nil
}