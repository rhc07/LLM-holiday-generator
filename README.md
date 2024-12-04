# LLM-holiday-generator

This is a small project to demonstrate how to build an API in Go (Golang) that leverages a Large Language Model (LLM) to generate vacation ideas based on user preferences. The API uses the following parameters to personalise the suggestions:

- `favourite_season` (e.g., summer, winter)
- `hobbies` (e.g., hiking, photography, skiing, surfing)
- `budget` (e.g., 1000) - this unit is in USD 💰

## Key Features

- **User-friendly API** built using the [GinGonic](https://github.com/gin-gonic/gin) framework.
- **LLM Integration** powered by [LangchainGo](https://github.com/langchain-ai/langchaingo).
- **Unique Request Tracking** using Google's [UUID package](https://pkg.go.dev/github.com/google/uuid).

## Project Structure

```
go-llm/
├── main.go            # Entry point of the application
├── routes/            # API route handlers
│   ├── vacation.go    # Endpoint logic for vacation generation
│   ├── types.go       # Request and response types for the vacation endpoints
├── chains/          # Service layer for LLM and business logic
│   ├── generator.go   # LLM request and response handling
│   ├── types.go       # Vaction type - this is an acting database
├── go.mod             # Dependency management file
├── go.sum             # Checksums for dependencies
└── README.md          # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.18 or later
- Access to an LLM service compatible with LangchainGo
- OpenAPI account and API Key

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/llm-vacation-api.git
   cd llm-vacation-api
   ```

2. Install dependencies:

   ```
   go mod tidy && go mod vendor
   ```

3. Set up environment variables for the OPENAI API:

   Run this command in the terminal when you have your OPENAI key:

   - `export OPENAI_API_KEY=$OPENAI_API_KEY`: Your API key for the OPENAI .

     **or**

   Create a `.env` file and add this variable:

   - `OPENAI_API_KEY=your-api-key`

### Running the API

1. Start the server:

   ```
   go run main.go
   ```

2. The API will run on `http://localhost:8080` by default.

### API Endpoint

- **POST `/vacations`**
  - **Description**: Generate a vacation idea based on user preferences.
  - **Request Body** (JSON):
    ```json
    {
      "favourite_season": "winter",
      "hobbies": ["skiing", "sledging"],
      "budget": "1000"
    }
    ```
  - **Response** (JSON):
    ```json
    {
      "id": "f5aa253f-35e0-4110-9927-136c0464fe5e",
      "completed": false,
      "idea": ""
    }
    ```
- **GET `/vacations/:id`**
  - **Description**: Fetch vacation idea from mock database.
  - **Response** (JSON):
    ```json
    {
      "id": "ef5cbd33-97f0-4d50-9741-c8142b224870",
      "completed": true,
      "idea": "Destination: Whistler, British Columbia, Canada Duration: 7 days Day 1: Arrival in Whistler - Check into a cozy ski-in/ski-out lodge in Whistler Village - Explore the village and enjoy a delicious dinner at a local restaurant Day 2-3: Skiing in Whistler Blackcomb..."
    }
    ```

## Technologies Used

- [Go (Golang)](https://golang.org/): Programming language.
- [GinGonic](https://github.com/gin-gonic/gin): Web framework for API development.
- [LangchainGo](https://github.com/langchain-ai/langchaingo): Go library for interacting with LLMs.
- [Google UUID](https://pkg.go.dev/github.com/google/uuid): Generate unique identifiers for tracking requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- The creators of GinGonic, LangchainGo, and UUID for their excellent tools.
- Inspiration from the Go development community.
