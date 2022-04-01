package cryptounifier

import (
	"github.com/cryptounifier/go-sdk/pkg/api/merchant"
	"github.com/cryptounifier/go-sdk/pkg/api/wallet"
	walletToken "github.com/cryptounifier/go-sdk/pkg/api/wallet-token"
)

func WalletAPI(walletKey string, secretKey, cryptoSymbol string) wallet.WalletAPIClient {
	return wallet.NewWalletApiClientFactory(walletKey, secretKey, cryptoSymbol)
}

func WalletTokenAPI(walletKey, secretKey, cryptoSymbol, tokenSymbol string) walletToken.WalletTokenAPIClient {
	return walletToken.NewWalletTokenApiClientFactory(walletKey, secretKey, cryptoSymbol, tokenSymbol)
}

func MerchantAPI(merchantKey, secretKey string) merchant.MerchantAPIClient {
	return merchant.NewMerchantApiClientFactory(merchantKey, secretKey)
}
