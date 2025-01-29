package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
	"github.com/kyzrfranz/interventure-cli/internal/scraper"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Extractor struct {
	apiToken string
}

type Extract struct {
	Summary string
	Contact []ContactDetails
}

type ContactDetails struct {
	Type  string
	Name  string
	Value string
}

func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	return reflector.Reflect(v)
}

func NewExtractor(token string) *Extractor {
	return &Extractor{
		apiToken: token,
	}
}

var extractSchema = GenerateSchema[Extract]()

func (e Extractor) Extract(url string) (Extract, error) {

	prompt := "ich gebe dir den content einer webseite eines MdB. Du sollst einen kurzen zusammenfassungstext über die Person insbes. bezüglich arbeits und sozialpolitischen themen vornehmen, sowie folgende kontakldaten sofern vorhanden: email, anschrift bundestag, anschrift wahlkreis, telefonnummer, faxnummer, homepage"

	data, err := scraper.Scrape(url)

	if err != nil {
		return Extract{}, err
	}

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("extract"),
		Description: openai.F("Extract data from a website"),
		Schema:      openai.F(extractSchema),
		Strict:      openai.Bool(true),
	}

	client := openai.NewClient(
		option.WithAPIKey(e.apiToken), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			}),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("%s. Die website daten sind: %s", prompt, data)),
		}),
		Model: openai.F(openai.ChatModelGPT4o2024_08_06),
	})
	if err != nil {
		return Extract{}, err
	}

	xtr := Extract{}
	err = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &xtr)
	if err != nil {
		return Extract{}, err
	}

	return xtr, nil
}
