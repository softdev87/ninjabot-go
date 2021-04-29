package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/markcheno/go-talib"
	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/pkg/exchange"
	"github.com/rodrigo-brito/ninjabot/pkg/model"
)

type Example struct{}

func (e Example) Init(settings model.Settings) {}

func (e Example) Timeframe() string {
	return "1m"
}

func (e Example) WarmupPeriod() int {
	return 14
}

func (e Example) Indicators(dataframe *model.Dataframe) {
	dataframe.Metadata["rsi"] = talib.Rsi(dataframe.Close, 14)
	dataframe.Metadata["ema"] = talib.Ema(dataframe.Close, 9)
}

func (e Example) OnCandle(dataframe *model.Dataframe, broker exchange.Broker) {
	fmt.Println("New Candle = ", dataframe.Pair, dataframe.LastUpdate, model.Last(dataframe.Close, 0))

	if model.Last(dataframe.Metadata["rsi"], 0) < 30 {
		broker.OrderMarket(exchange.BuyOrder, dataframe.Pair, 1)
	}

	if model.Last(dataframe.Metadata["rsi"], 0) > 70 {
		broker.OrderMarket(exchange.SellOrder, dataframe.Pair, 1)
	}
}

func main() {
	var (
		apiKey    = os.Getenv("API_KEY")
		secretKey = os.Getenv("API_SECRET")
		ctx       = context.Background()
	)

	settings := model.Settings{
		Pairs: []string{
			"BTCUSDT",
			"ETHUSDT",
		},
	}

	binance := exchange.NewBinance(apiKey, secretKey)
	strategy := Example{}

	bot, err := ninjabot.NewBot(settings, binance, strategy)
	if err != nil {
		log.Fatalln(err)
	}

	err = bot.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
