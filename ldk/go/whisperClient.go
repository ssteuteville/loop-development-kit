package ldk

import (
	"context"
	"errors"
	"fmt"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/server"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"
	"io"

	"github.com/golang/protobuf/ptypes"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/proto"
)

// WhisperClient is used by the controller plugin to facilitate plugin initiated communication with the host
type WhisperClient struct {
	WhisperClient proto.WhisperClient
	Session       *server.Session
}

// Markdown is used by loops to create markdown whispers
// This method is blocking until the Whisper is closed, or the context provided is cancelled.
func (m *WhisperClient) Markdown(ctx context.Context, content *whisper.WhisperContentMarkdown) error {
	_, err := m.WhisperClient.WhisperMarkdown(ctx, &proto.WhisperMarkdownRequest{
		Meta: &proto.WhisperMeta{
			Label: content.Label,
		},
		Markdown: content.Markdown,
		Session:  m.Session.ToProto(),
	})
	return err
}

// Confirm is used by loops to create confirm whispers
// This method is blocking until the Whisper is closed, or the context provided is cancelled.
func (m *WhisperClient) Confirm(ctx context.Context, content *whisper.WhisperContentConfirm) (bool, error) {
	response, err := m.WhisperClient.WhisperConfirm(ctx, &proto.WhisperConfirmRequest{
		Meta: &proto.WhisperMeta{
			Label: content.Label,
		},
		Markdown:     content.Markdown,
		RejectLabel:  content.RejectLabel,
		ResolveLabel: content.ResolveLabel,
		Session:      m.Session.ToProto(),
	})
	if err != nil {
		return false, err
	}

	if response == nil {
		return false, errors.New("no response")
	}

	return response.Response, nil
}

// Disambiguation is used by loops to create disambiguation whispers
// This method is blocking until the Whisper is closed, or the context provided is cancelled.
func (m *WhisperClient) Disambiguation(ctx context.Context, content *whisper.WhisperContentDisambiguation) (bool, error) {
	req, err := content.ToProto()
	if err != nil {
		return false, fmt.Errorf("failed to encode content to proto: %w", err)
	}

	req.Session = m.Session.ToProto()

	client, err := m.WhisperClient.WhisperDisambiguation(ctx, req)
	if err != nil {
		return false, err
	}

	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return false, errors.New("unexpected end of stream")
		}
		if err != nil {
			return false, err
		}

		genericInput := content.Elements[resp.Key]
		if element := genericInput.(*whisper.WhisperContentDisambiguationElementOption); element.OnChange != nil {
			element.OnChange(resp.Key)
		}
	}

	// return false, nil
}

// Form is used by loops to create form whispers
// This method is blocking until the Whisper is closed, or the context provided is cancelled.
func (m *WhisperClient) Form(ctx context.Context, content *whisper.WhisperContentForm) (bool, map[string]whisper.WhisperContentFormOutput, error) {
	inputs := make(map[string]*proto.WhisperFormInput, len(content.Inputs))
	for key, input := range content.Inputs {
		protoInput, err := input.ToProto()
		if err != nil {
			return false, nil, err
		}

		inputs[key] = protoInput
	}

	client, err := m.WhisperClient.WhisperForm(ctx, &proto.WhisperFormRequest{
		Meta: &proto.WhisperMeta{
			Label: content.Label,
		},
		Markdown:    content.Markdown,
		SubmitLabel: content.SubmitLabel,
		CancelLabel: content.CancelLabel,
		Inputs:      inputs,
		Session:     m.Session.ToProto(),
	})
	if err != nil {
		return false, nil, err
	}

	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return false, nil, errors.New("unexpected end of stream")
		}
		if err != nil {
			return false, nil, err
		}

		switch respContainer := resp.WhisperFormResponseOneof.(type) {
		case *proto.WhisperFormStreamResponse_Result:
			outputs := make(map[string]whisper.WhisperContentFormOutput, len(respContainer.Result.Outputs))
			for key, protoOutput := range respContainer.Result.Outputs {
				output, err := whisper.WhisperContentFormOutputFromProto(protoOutput)
				if err != nil {
					return false, nil, err
				}
				outputs[key] = output
			}
			return respContainer.Result.Submitted, outputs, nil

		case *proto.WhisperFormStreamResponse_Update:
			genericInput := content.Inputs[respContainer.Update.Key]
			switch inputContainer := respContainer.Update.Output.OutputOneof.(type) {
			case *proto.WhisperFormOutput_Checkbox_:
				if input := genericInput.(*whisper.WhisperContentFormInputCheckbox); input.OnChange != nil {
					input.OnChange(inputContainer.Checkbox.Value)
				}
			case *proto.WhisperFormOutput_Email_:
				if input := genericInput.(*whisper.WhisperContentFormInputEmail); input.OnChange != nil {
					input.OnChange(inputContainer.Email.Value)
				}
			case *proto.WhisperFormOutput_Markdown_:
				if input := genericInput.(*whisper.WhisperContentFormInputMarkdown); input.OnChange != nil {
					input.OnChange(inputContainer.Markdown.Value)
				}
			case *proto.WhisperFormOutput_Number_:
				if input := genericInput.(*whisper.WhisperContentFormInputNumber); input.OnChange != nil {
					input.OnChange(inputContainer.Number.Value)
				}
			case *proto.WhisperFormOutput_Password_:
				if input := genericInput.(*whisper.WhisperContentFormInputPassword); input.OnChange != nil {
					input.OnChange(inputContainer.Password.Value)
				}
			case *proto.WhisperFormOutput_Radio_:
				if input := genericInput.(*whisper.WhisperContentFormInputRadio); input.OnChange != nil {
					input.OnChange(inputContainer.Radio.Value)
				}
			case *proto.WhisperFormOutput_Select_:
				if input := genericInput.(*whisper.WhisperContentFormInputSelect); input.OnChange != nil {
					input.OnChange(inputContainer.Select.Value)
				}
			case *proto.WhisperFormOutput_Tel_:
				if input := genericInput.(*whisper.WhisperContentFormInputTel); input.OnChange != nil {
					input.OnChange(inputContainer.Tel.Value)
				}
			case *proto.WhisperFormOutput_Text_:
				if input := genericInput.(*whisper.WhisperContentFormInputText); input.OnChange != nil {
					input.OnChange(inputContainer.Text.Value)
				}
			case *proto.WhisperFormOutput_Time_:
				if input := genericInput.(*whisper.WhisperContentFormInputTime); input.OnChange != nil {
					value, err := ptypes.Timestamp(inputContainer.Time.Value)
					if err != nil {
						return false, nil, err
					}

					input.OnChange(value)
				}
			default:
				return false, nil, errors.New("received unknown content type in response")
			}
		default:
			return false, nil, fmt.Errorf("content had unexpected type %T", respContainer)
		}
	}
}

// List is used by loops to create list whispers
// This method is blocking until the Whisper is closed, or the context provided is cancelled.
func (m *WhisperClient) List(ctx context.Context, content *whisper.WhisperContentList) error {
	req, err := content.ToProto()
	if err != nil {
		return fmt.Errorf("failed to encode content to proto: %w", err)
	}

	req.Session = m.Session.ToProto()

	_, err = m.WhisperClient.WhisperList(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
