package loop_test

import (
	"context"
	"errors"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	loop "github.com/open-olive/loop-development-kit/ldk/go/examples/whisper-list/loop"
	ldk "github.com/open-olive/loop-development-kit/ldk/go/v2"
	ldktest "github.com/open-olive/loop-development-kit/ldk/go/v2/ldk-test"
)

func TestLoop(t *testing.T) {
	type listRequest struct {
		ctx context.Context
		w   *whisper.WhisperContentList
	}
	type listResponse struct {
		err error
	}
	listRequestChan := make(chan listRequest)
	listResponseChan := make(chan listResponse)

	sidekick := &ldktest.Sidekick{
		WhisperService: &ldktest.WhisperService{
			Listf: func(ctx context.Context, w *whisper.WhisperContentList) error {
				listRequestChan <- listRequest{ctx, w}
				res := <-listResponseChan
				return res.err
			},
		},
	}

	l := ldk.NewLogger("example-whisper-list")
	c, err := loop.NewLoop(l)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.LoopStart(sidekick); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := c.LoopStop(); err != nil {
			t.Fatal(err)
		}
	}()

	// stage 1: get list whisper
	if err := (func() error {
		timeout := time.NewTimer(time.Second)
		defer timeout.Stop()
		select {
		case <-timeout.C:
			return errors.New("timeout")
		case req := <-listRequestChan:
			exp := &whisper.WhisperContentList{
				Label: "MCMG Location",
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
			}
			if got := req.w; !cmp.Equal(got, exp) {
				t.Errorf("unexpected whisper content:\n%s\n", cmp.Diff(got, exp))
			}
			return nil
		}
	}()); err != nil {
		t.Errorf("stage 1 failed\n%v", err)
		return
	}

	// stage 2: simulate closing list whisper
	if err := (func() error {
		timeout := time.NewTimer(time.Second)
		defer timeout.Stop()
		select {
		case <-timeout.C:
			return errors.New("timeout")
		case listResponseChan <- listResponse{
			err: nil,
		}:
			return nil
		}
	}()); err != nil {
		t.Errorf("stage 2 failed\n%v", err)
		return
	}
}
