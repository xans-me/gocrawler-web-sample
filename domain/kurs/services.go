package kurs

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"gocrawler-web-sample/infrastructure/environment"
	"gocrawler-web-sample/shared/times"
	"strconv"
	"strings"
)

type Service struct {
	appEnvironment environment.AppEnvironment
	repo           IKursRepository
}

func (s *Service) IndexingKurs() (data []DataKurs, err error) {
	data, err = scrappingKurs()
	if err != nil {
		return
	}

	go s.BulkInsertDataKurs(data)

	return
}

func (s *Service) BulkInsertDataKurs(data []DataKurs) {
	for _, kurs := range data {
		err := s.InsertDataKurs(kurs)
		if err != nil {
			log.Error("error when insert scrapped kurs ==> ", err.Error())
		}
	}
}

func (s *Service) InsertDataKurs(kurs DataKurs) (err error) {
	countBN, countTT, countERates, err := s.repo.IsExistKurs(kurs)
	if err != nil {
		return err
	}

	if countBN < 1 {
		err = s.repo.InsertBankNotes(kurs.Symbol, kurs.BankNote, kurs.Date)
		if err != nil {
			return
		}
	}

	if countTT < 1 {
		err = s.repo.InsertTTCounter(kurs.Symbol, kurs.TTCounter, kurs.Date)
		if err != nil {
			return
		}
	}

	if countERates < 1 {
		err = s.repo.InsertERates(kurs.Symbol, kurs.ERate, kurs.Date)
		if err != nil {
			return
		}
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
	if strings.Contains(value, ",") {
		replaceComma := strings.Replace(value, ",", ".", 1)
		removePointFirst := strings.Replace(replaceComma, ".", "", 1)
		value = removePointFirst
	}

	return strconv.ParseFloat(value, 64)
}

// NewService function to init service instance
func NewService(appEnvironment environment.AppEnvironment, repo IKursRepository) *Service {
	return &Service{appEnvironment: appEnvironment, repo: repo}
}
