package jwt

import (
	"fmt"
)

func validateIssuer(issuer string) error {
	var acceptedIssuers = [...]string{
		// Official.
		"marketplace",
	}

	for _, acceptedIssuer := range acceptedIssuers {
		if acceptedIssuer == issuer {
			return nil
		}
	}

	return fmt.Errorf("wrong domain provided: %s", issuer)
}
