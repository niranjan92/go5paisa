package go5paisa

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Position represents a position object returned in the JSON response
type Position struct {
	BodQty        float32 `json:"BodQty"`
	BookedPL      float32 `json:"BookedPL"`
	BuyAvgRate    float32 `json:"BuyAvgRate"`
	BuyValue      float32 `json:"BuyValue"`
	Exch          string  `json:"Exch"`
	ExchType      float32 `json:"ExchType"`
	LTP           float32 `json:"LTP"`
	MTOM          float32 `json:"MTOM"`
	Multiplier    float32 `json:"Multiplier"`
	NetQty        float32 `json:"NetQty"`
	OrderFor      float32 `json:"OrderFor"`
	PreviousClose float32 `json:"PreviousClose"`
	ScripCode     float32 `json:"ScripCode"`
	ScripName     float32 `json:"ScripName"`
	SellAvgRate   float32 `json:"SellAvgRate"`
	SellQty       float32 `json:"SellQty"`
	SellValue     float32 `json:"SellValue"`
}

type positionResponseData struct {
	Head interface{} `json:"head"`
	Body Positions   `json:"body"`
}

// Positions contains Position objects
type Positions struct {
	PositionDetail []Position `json:"NetPositionDetail"`
}

func parseResponseBody(resBody []byte, obj Positions) {
	var body positionResponseData
	body.Body = obj
	if err := json.Unmarshal(resBody, &body); err != nil {
		log.Fatal("Error parsing JSON response:", err)
	}
}

// GetPositions fetches margins of the user
func (c *Client) GetPositions() (Positions, error) {
	var positions Positions
	c.appConfig.head.RequestCode = positionsRequestCode
	payloadBody := genericPayloadBody{
		ClientCode: c.clientCode,
	}
	payload := genericPayload{
		Head: c.appConfig.head,
		Body: payloadBody,
	}
	jsonValue, _ := json.Marshal(payload)
	res, err := c.connection.Post(baseURL+positionsRoute, contentType, bytes.NewBuffer(jsonValue))
	if err != nil {
		return positions, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return positions, err
	}
	parseResBody(resBody, &positions)
	return positions, nil
}
