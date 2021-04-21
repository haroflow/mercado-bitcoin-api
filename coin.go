package mercadobitcoin

// Coin is the ticker (identifier, name) of a coin
type Coin string

// Coins is a map of coin ticker (identifier, name) to a description
var Coins = map[Coin]string{
	"ACMFT":     "Fan Token ASR",
	"ACORDO01":  "None",
	"ASRFT":     "Fan Token ASR",
	"ATMFT":     "Fan Token ATM",
	"BCH":       "Bitcoin Cash",
	"BTC":       "Bitcoin",
	"CAIFT":     "Fan Token CAI",
	"CHZ":       "Chiliz",
	"ETH":       "Ethereum",
	"GALFT":     "Fan Token GAL",
	"IMOB01":    "None",
	"JUVFT":     "Fan Token JUV",
	"LINK":      "CHAINLINK",
	"LTC":       "Litecoin",
	"MBCONS01":  "Cota de Consórcio 01",
	"MBCONS02":  "Cota de Consórcio 02",
	"MBFP01":    "None",
	"MBFP02":    "None",
	"MBFP03":    "None",
	"MBPRK01":   "Precatório MB SP01",
	"MBPRK02":   "Precatório MB SP02",
	"MBPRK03":   "Precatório MB BR03",
	"MBPRK04":   "Precatório MB RJ04",
	"MBVASCO01": "MBVASCO01",
	"MCO2":      "MCO2",
	"OGFT":      "Fan Token ASR",
	"PAXG":      "PAX Gold",
	"PSGFT":     "Fan Token PSG",
	"USDC":      "USD Coin",
	"WBX":       "WiBX",
	"XRP":       "XRP",
}
