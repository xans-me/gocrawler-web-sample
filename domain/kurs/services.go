package kurs

import (
	"context"
	"github.com/gocolly/colly"
	"gocrawler-web-sample/infrastructure/environment"
	"gocrawler-web-sample/shared/times"
	"strconv"
	"strings"
)

type Service struct {
	appEnvironment   environment.AppEnvironment
	publicRepository IKursRepository
}

func (s *Service) IndexingKurs(_ context.Context) (data []DataKurs, err error) {
	data, err = scrappingKurs()
	if err != nil {
		return
	}

	return
}

func scrappingKurs() (ResultIndexing []DataKurs, err error) {
	c := colly.NewCollector()
	c.OnHTML("tbody", func(tab *colly.HTMLElement) {
		tab.ForEach("tr", func(indexTr int, tr *colly.HTMLElement) {

			eRateBuy, err := convertStringKursToFloat64(tr.ChildText("p[rate-type=\"ERate-buy\"]"))
			if err != nil {
				return
			}
			eRateSell, err := convertStringKursToFloat64(tr.ChildText("p[rate-type=\"ERate-sell\"]"))
			if err != nil {
				return
			}
			ttBuy, err := convertStringKursToFloat64(tr.ChildText("p[rate-type=\"TT-buy\"]"))
			if err != nil {
				return
			}
			ttSell, err := convertStringKursToFloat64(tr.ChildText("p[rate-type=\"TT-sell\"]"))
			if err != nil {
				return
			}
			bnBuy, err := convertStringKursToFloat64(tr.ChildText("p[rate-type=\"BN-buy\"]"))
			if err != nil {
				return
			}
			bnSell, err := convertStringKursToFloat64(tr.ChildText("p[rate-type=\"BN-sell\"]"))
			if err != nil {
				return
			}

			dataKurs := DataKurs{
				Symbol: tr.Attr("code"),
				ERate: ERate{
					ERateBuy:  eRateBuy,
					ERateSell: eRateSell,
				},
				TTCounter: TTCounter{
					TTBuy:  ttBuy,
					TTSell: ttSell,
				},
				BankNote: BankNote{
					BNBuy:  bnBuy,
					BNSell: bnSell,
				},
				Date: times.Now(times.TimeGmt, times.DateFormat),
			}
			ResultIndexing = append(ResultIndexing, dataKurs)
		})
	})

	err = c.Visit("https://www.bca.co.id/en/informasi/kurs")
	if err != nil {
		return nil, err
	}

	return
}

func convertStringKursToFloat64(value string) (result float64, err error) {
	valueString := strings.Replace(strings.Replace(value, ",", ".", 1), ".", "", 1)
	return strconv.ParseFloat(valueString, 64)
}

// NewService function to init service instance
func NewService(appEnvironment environment.AppEnvironment, publicRepository IKursRepository) *Service {
	return &Service{appEnvironment: appEnvironment, publicRepository: publicRepository}
}
