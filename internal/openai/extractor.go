package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
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

type ContactType string

const EmailType ContactType = "email"
const AddressType ContactType = "address"
const PhoneType ContactType = "phone"
const FaxType ContactType = "fax"

type ContactDetails struct {
	Type  ContactType
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

func (e Extractor) Extract(url string) (Extract, error) {

	prompt := "ich gebe dir eine bio zu einem mdb, inklusive liste über tätigkeit in arbeitskreisen des bundestags, sowie bezüge.\n\nverfasse basierend darauf einen kurzen, prägnanten persönlichen text der den kandiadten auf seine standpunkte zum thema arbeits & sozialploitik anspricht, mehr für selbständige in DE zu unternehmen:"

	client := openai.NewClient(
		option.WithAPIKey(e.apiToken), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)
	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("%s. Die personendaten sind: %s", prompt, "")),
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
