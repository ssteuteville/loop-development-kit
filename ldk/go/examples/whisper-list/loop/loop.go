package loop

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"

	ldk "github.com/open-olive/loop-development-kit/ldk/go/v2"
)

func Serve() error {
	l := ldk.NewLogger("example-whisper-list")
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

	sidekick ldk.Sidekick
	logger   *ldk.Logger
}

// NewLoop returns a pointer to a loop
func NewLoop(logger *ldk.Logger) (*Loop, error) {
	return &Loop{
		logger: logger,
	}, nil
}

// LoopStart is called by the host when the plugin is started to provide access to the host process
func (c *Loop) LoopStart(sidekick ldk.Sidekick) error {
	c.logger.Info("loopStart called")
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.sidekick = sidekick

	go func() {
		err := c.sidekick.Whisper().List(c.ctx, &whisper.WhisperContentList{
			Label:    "MCMG Location",
			Markdown: "",
			Elements: map[string]whisper.WhisperContentListElement{
				"topMessage": &whisper.WhisperContentListElementMessage{
					Style:  whisper.WhisperContentListElementStyleNone,
					Header: "Hello World, I am a subitle",
					Body:   "This is what body copy looks like. Just a bit, don’t overdo it!",
					Align:  whisper.WhisperContentListElementAlignLeft,
					Order:  0,
				},
				"successMessage": &whisper.WhisperContentListElementMessage{
					Align:  whisper.WhisperContentListElementAlignCenter,
					Header: "This is an alert message!",
					Body:   "It should be highlighted green.",
					Style:  whisper.WhisperContentListElementStyleSuccess,
					Order:  1,
				},
				"sectionDivider": &whisper.WhisperContentListElementDivider{
					Order: 2,
				},
				"sectionTitle": &whisper.WhisperContentListElementMessage{
					Style:  whisper.WhisperContentListElementStyleNone,
					Header: "Let’s set the table",
					Align:  whisper.WhisperContentListElementAlignCenter,
					Order:  3,
				},
				"name": &whisper.WhisperContentListElementPair{
					Label: "Name",
					Order: 4,
					Value: "David Simon MD",
				},
				"shoeSize": &whisper.WhisperContentListElementPair{
					Label: "Shoe Size",
					Order: 5,
					Value: "38",
				},
				"birthDate": &whisper.WhisperContentListElementPair{
					Style: whisper.WhisperContentListElementStyleWarning,
					Label: "Birth Date",
					Order: 6,
					Value: "Feb 30th, 1999",
				},
				"favoriteColor": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "FavoriteColor",
					Order: 7,
					Value: "Greige",
				},
				"favoriteAnimal": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "Dogs or Cats",
					Order: 8,
					Value: "Bats",
				},
				"streetName": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "Street Name",
					Order: 9,
					Value: "Main Street",
				},
				"zipCode": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "Zip Code",
					Order: 10,
					Value: "10000",
				},
				"city": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "City",
					Order: 11,
					Value: "Townsville",
				},
				"phone": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "Phone",
					Order: 12,
					Value: "123-456-7890",
				},
				"favoriteCondiment": &whisper.WhisperContentListElementPair{
					Style: whisper.WhisperContentListElementStyleWarning,
					Extra: true,
					Label: "Favorite Condiment",
					Order: 13,
					Value: "Pizza",
				},
				"100MeterDashTime": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "100m Dash Time",
					Order: 14,
					Value: "4 minutes",
				},
				"nickname": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "Nickname",
					Order: 15,
					Value: "Old Greg",
				},
				"notes": &whisper.WhisperContentListElementPair{
					Extra: true,
					Label: "Notes",
					Order: 16,
					Value: "Lorem ipsum sit amet dolor why does this always feel like a decree by the ancient Romans? It’s just filler text.",
				},
				"failureMessage": &whisper.WhisperContentListElementMessage{
					Align:  whisper.WhisperContentListElementAlignCenter,
					Body:   "It should be highlighted red.",
					Extra:  true,
					Header: "This is an alert message!",
					Order:  17,
					Style:  whisper.WhisperContentListElementStyleError,
				},
				"link": &whisper.WhisperContentListElementLink{
					Align: whisper.WhisperContentListElementAlignCenter,
					Extra: true,
					Href:  "https://isitchristmas.com/",
					Order: 18,
					Style: whisper.WhisperContentListElementStyleNone,
					Text:  "IsItChristmas.com",
				},
			},
		})
		if err != nil {
			c.logger.Error("failed to emit whisper", "error", err)
		}
	}()

	return nil
}

// LoopStop is called by the host when the plugin is stopped
func (c *Loop) LoopStop() error {
	c.logger.Info("loopStop called")
	c.cancel()

	return nil
}
