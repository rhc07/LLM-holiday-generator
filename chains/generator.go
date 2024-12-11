package chains

import (
	"context"
	"errors"
	"log"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

var Vacations []*Vacation

func GetVacationFromDb(id uuid.UUID) (Vacation, error) {
	// Use the slices package to find the index of the object with
	// matching ID in the database. If it does not exist, this will return
	// -1
	idx := slices.IndexFunc(Vacations, func(v *Vacation) bool { return v.ID == id })

	// If the ID didn't exist, return an error and let the caller
	// handle it
	if idx < 0 {
		return Vacation{}, errors.New("ID Not Found")
	}

	return *Vacations[idx], nil
}

func GenerateVacationIdeaChange(id uuid.UUID, budget int, weather string, hobbies []string, travellingMonth string, flyingFrom string, flyingTime int) {
	log.Printf("Generating new vacation idea with ID: %s", id)

	// Soft create the mock database instance
	vacation := &Vacation{ID: id, Idea: "", Completed: false}
	Vacations = append(Vacations, vacation)

	// Connect to the openai LLM
	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Create a system prompt with the season, hobbies, and budget parameters
	// Helps tell the LLM how to act / respond to queries
	systemPromptMessageString := "You are an AI travel agent that will help me create a vacation idea.\n" +
		"My favorite weather is {{.weather}}.\n" +
		"My hobbies include {{.hobbies}}.\n" +
		"I plan to travel during {{.travellingMonth}}.\n" +
		"I plan to fly from {{.flyingFrom}}.\n" +
		"I do not wish to fly for more than {{.flyingTime}} hours.\n" +
		"My budget is {{.budget}} GBP.\n"

	systemPromptMessage := prompts.NewSystemMessagePromptTemplate(systemPromptMessageString, []string{"weather", "hobbies", "budget", "travellingMonth", "flyingFrom", "flyingTime"})

	// Create a human prompt with the request that a human would have
	humanPromptMessageString := "Write up a travel itinerary"
	humanPromptMessage := prompts.NewHumanMessagePromptTemplate(humanPromptMessageString, []string{})

	// Create a chat prompt consisting of the system messages and human messages
	// At this point, we will also inject the values into the prompts
	// and turn them into message content objects which we can feed through
	// to our LLM
	chatPrompt := prompts.NewChatPromptTemplate([]prompts.MessageFormatter{systemPromptMessage, humanPromptMessage})

	vals := map[string]any{
		"weather":         weather,
		"budget":          budget,
		"hobbies":         strings.Join(hobbies, ","),
		"travellingMonth": travellingMonth,
		"flyingFrom":      flyingFrom,
		"flyingTime":      flyingTime,
	}

	msgs, err := chatPrompt.FormatMessages(vals)

	if err != nil {
		log.Printf("Error formatting messages; %v", err)
		return
	}

	content := []llms.MessageContent{
		llms.TextParts(msgs[0].GetType(), msgs[0].GetContent()),
		llms.TextParts(msgs[1].GetType(), msgs[1].GetContent()),
	}

	// Invoke the LLM with the messages
	completion, err := llm.GenerateContent(ctx, content)

	if err != nil {
		log.Printf("Error invoking llm: %v", err)
		return
	}
	vacation.Idea = completion.Choices[0].Content
	vacation.Completed = true

	log.Printf("Generation for %s is done!", vacation.ID)
}
