package jwt

import (
	"fmt"
	"gitlab.com/0xscore/degen-sdk/constants"
	pkgErrors "gitlab.com/0xscore/degen-sdk/errors"
)

func validateIssuer(issuer string) error {
	var acceptedIssuers = [...]string{
		// Official.
		"api.0xscore.pro",
		"api.0xscore.io",

		// Forks.
		"api.blastscore.io",

		// Stage.
		"stage.0xscore.io",
	}

	for _, acceptedIssuer := range acceptedIssuers {
		if acceptedIssuer == issuer {
			return nil
		}
	}

	return fmt.Errorf("wrong domain provided: %s", issuer)
}

func validateWalletAddr(walletAddress string) error {
	if len(walletAddress) != constants.EthAddrLength {
		return pkgErrors.WalletAddrLengthErr
	}
	return nil
}
