package pipedata

import (
	"encoding/json"
	"fmt"
	"service/files_server/demo/internal/dto"
)

type Codec interface {
	Json(b []byte) error
}

type jsonCodec struct{}

func (jsonCodec) Json(b []byte) error {

	var f interface{}
	json.Unmarshal(b, &f)

	myMap := f.(map[string]interface{})

	fmt.Println(myMap)
	return nil
}

var JSON Codec = jsonCodec{}

func ProcessData(m interface{}) {

	empData, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonStr := string(empData)
	b := []byte(jsonStr)
	JSON.Json(b)
}

func DataProcessStep1(sd <-chan []dto.OrderCreated, dash chan<- []dto.OrderCreated) error {
	var revenue float64 = 0.0
	for s := range sd {
		for n := range s {
			revenue += s[n].ProductsData.Price
			s[n].Total = revenue

			dash <- s
		}
	}
	close(dash)

	return nil
}

func DataProcessStep2(mark <-chan []dto.SalesMaketingData, res chan<- []dto.SalesMaketingData) error {

	var sl []string
	var g []string
	var a []string
	for s := range mark {
		for ds := range s {
			sl = append(sl, s[ds].SocialMedia)
			g = append(g, s[ds].Gender)
			a = append(a, s[ds].Age)
		}
	}
	var d []dto.SalesMaketingData

	slV := SliceCountValues(sl)
	gV := SliceCountValues(g)
	aV := SliceCountValues(a)

	for i := range d {
		d[i].TotalSocialMedia = slV
		d[i].TotalGender = gV
		d[i].TotalAge = aV
	}

	res <- d
	close(res)

	return nil
}

func DataProcessStep3(mark <-chan []dto.SalesMaketingData, res chan<- []dto.SalesMaketingData) error {

	var sl []string
	var g []string
	var a []string
	for s := range mark {
		for ds := range s {
			sl = append(sl, s[ds].SocialMedia)
			g = append(g, s[ds].Gender)
			a = append(a, s[ds].Age)
		}
	}
	var d dto.SalesMaketingData

	slV := SliceCountValues(sl)
	gV := SliceCountValues(g)
	aV := SliceCountValues(a)

	d.TotalSocialMedia = slV
	d.TotalGender = gV
	d.TotalAge = aV

	return nil
}

// m := map[string]interface{}{
// 	"key1": "aGuide", "lastname": "hub",
// }
