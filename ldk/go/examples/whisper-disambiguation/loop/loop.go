package loop

import (
	"context"
	"encoding/json"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/client"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/utils"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"
	"time"

	ldk "github.com/open-olive/loop-development-kit/ldk/go/v2"
)

type recall struct {
	Address1       string `json:"address_1"`
	Address2       string `json:"address_2"`
	Codes          string `json:"code_info"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Classification string `json:"classification"`
	Date           string `json:"recall_initiation_date"`
	Description    string `json:"product_description"`
	Distribution   string `json:"distribution_pattern"`
	Firm           string `json:"recalling_firm"`
	ID             string `json:"recall_number"`
	Quantity       string `json:"product_quantity"`
	Reason         string `json:"reason_for_recall"`
	RecallType     string `json:"voluntary_mandated"`
	State          string `json:"state"`
	Type           string `json:"product_type"`
	Zip            string `json:"postal_code"`
}

type apiResponse struct {
	Meta    interface{} `json:"meta"`
	Results []recall    `json:"results"`
}

// Serve allows Olive Helps to have access the loop
func Serve() error {
	l := utils.NewLogger("example-whisper-disambiguation")
	loop, err := NewLoop(l)
	if err != nil {
		return err
	}
	ldk.ServeLoopPlugin(l, loop)
	return nil
}

// Loop is a structure for generating SideKick whispers
type Loop struct {
	ctx    context.Context
	cancel context.CancelFunc

	sidekick client.Sidekick
	logger   *utils.Logger
}

// NewLoop returns a pointer to a loop
func NewLoop(logger *utils.Logger) (*Loop, error) {
	return &Loop{
		logger: logger,
	}, nil
}

// LoopStart is called by the host when the plugin is started to provide access to the host process
func (c *Loop) LoopStart(sidekick client.Sidekick) error {
	c.logger.Info("Starting example whisper disambiguation loop")
	c.ctx, c.cancel = context.WithCancel(context.Background())

	c.sidekick = sidekick

	now := time.Now()

	response, err := sidekick.Network().HTTPRequest(c.ctx, &service.HTTPRequest{
		URL:    "https://api.fda.gov/food/enforcement.json?search=report_date:[" + now.AddDate(0, -3, 0).Format("20060102") + "+TO+" + now.Format("20060102") + "]&limit=10",
		Method: "GET",
		Body:   nil,
	})
	if err != nil {
		c.logger.Error("received error from callback", err)
		return err
	}

	if response.ResponseCode == 200 {
		var data apiResponse

		if err := json.Unmarshal(response.Data, &data); err != nil {
			c.logger.Error("Error unmarshaling response payload", err)
			return err
		}

		elements := make(map[string]whisper.WhisperContentDisambiguationElement)

		for index := range data.Results {
			item := data.Results[index]
			elements[item.ID] = &whisper.WhisperContentDisambiguationElementOption{
				Label: "😝" + item.Firm + " (" + item.Date + ")",
				Order: uint32(index) + 1,
				OnChange: func(key string) {
					go func() {
						err := c.sidekick.Whisper().List(c.ctx, &whisper.WhisperContentList{
							Label: item.Firm + " Recall",
							Elements: map[string]whisper.WhisperContentListElement{
								"topMessage": &whisper.WhisperContentListElementMessage{
									Style: whisper.WhisperContentListElementStyleNone,
									Body:  item.Description,
									Align: whisper.WhisperContentListElementAlignLeft,
									Order: 0,
								},
								"sectionDivider": &whisper.WhisperContentListElementDivider{
									Order: 1,
								},
								"reason": &whisper.WhisperContentListElementPair{
									Label: "Reason",
									Order: 2,
									Value: item.Reason,
								},
								"distribution": &whisper.WhisperContentListElementPair{
									Label: "Distribution",
									Order: 3,
									Value: item.Distribution,
								},
								"quantity": &whisper.WhisperContentListElementPair{
									Label: "Quantity",
									Order: 4,
									Value: item.Quantity,
								},
								"codes": &whisper.WhisperContentListElementPair{
									Extra: true,
									Label: "Codes",
									Order: 5,
									Value: item.Codes,
								},
								"id": &whisper.WhisperContentListElementPair{
									Extra: true,
									Label: "Recall number",
									Order: 6,
									Value: item.ID,
								},
								"date": &whisper.WhisperContentListElementPair{
									Extra: true,
									Label: "Date initiated",
									Order: 7,
									Value: item.Date,
								},
								"recallType": &whisper.WhisperContentListElementPair{
									Extra: true,
									Label: "Recall type",
									Order: 8,
									Value: item.RecallType,
								},
								"type": &whisper.WhisperContentListElementPair{
									Extra: true,
									Label: "Product type",
									Order: 9,
									Value: item.Type,
								},
								"classification": &whisper.WhisperContentListElementPair{
									Extra: true,
									Label: "Classification",
									Order: 10,
									Value: item.Classification,
								},
								"address": &whisper.WhisperContentListElementPair{
									Extra: true,
									Label: "Company address",
									Order: 11,
									Value: item.Address1 + " " + item.Address2 + " " + item.City + ", " + item.State + " " + item.Zip + " " + item.Country,
								},
							},
						})
						if err != nil {
							c.logger.Error("failed to emit whisper", "error", err)
						}
					}()
				},
			}
		}

		elements["header1"] = &whisper.WhisperContentDisambiguationElementText{
			Body:  "# Text header\n## Text subheader",
			Order: 0,
		}
		elements["header2"] = &whisper.WhisperContentDisambiguationElementText{
			Body:  "# Text header",
			Order: 6,
		}

		go func() {
			_, err := c.sidekick.Whisper().Disambiguation(c.ctx, &whisper.WhisperContentDisambiguation{
				Label:    "Latest FDA Food Recalls",
				Markdown: "The Latest FDA Food Recalls",
				Elements: elements,
			})
			if err != nil {
				c.logger.Error("failed to emit whisper", "error", err)
			}
		}()

		return nil
	}

	c.logger.Error("received response code of", response.ResponseCode)
	return nil
}

// LoopStop is called by the host when the plugin is stopped
func (c *Loop) LoopStop() error {
	c.logger.Info("controller LoopStop called")
	c.cancel()

	return nil
}
