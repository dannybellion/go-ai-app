# Go AI Chat Application

A modern chat application built with Go and React that integrates with OpenAI's API to provide AI-powered conversations. The application includes a backend API server, persistent conversation storage, and both web and CLI interfaces.

## Features

- **AI Chat**: Communicate with OpenAI's GPT models through a clean interface
- **Conversation History**: Automatically saves all conversations to a local database
- **Multiple Interfaces**: 
  - Web interface built with React
  - CLI tool for terminal-based interactions
- **RESTful API**: Well-structured API endpoints for chat and history retrieval

## Tech Stack

### Backend
- **Go**: Server-side programming language
- **Gin**: Web framework for API routing and handling HTTP requests
- **SQLite**: Lightweight database for persistent storage
- **OpenAI API**: Integration with GPT models

### Frontend
- **React**: UI library for building the web interface
- **CSS3**: Modern styling with responsive design

## Project Structure

```
go-ai-app/
cmd/
    cli/        # Command-line interface tool
config/     # Application configuration
database/   # Database setup and connections
frontend/   # React web application
handlers/   # HTTP request handlers
models/     # Data structures and database operations
routes/     # API routing
utils/      # Utility functions and HTTP client
```

## Getting Started

### Prerequisites

- Go 1.17 or higher
- Node.js 14.x or higher
- npm
- OpenAI API key

### Environment Setup

Create a `.env` file in the root directory with your OpenAI API key:

```
OPENAI_API_KEY=your_openai_api_key_here
```

### Running the Backend

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/go-ai-app.git
   cd go-ai-app
   ```

2. Install Go dependencies:
   ```
   go mod download
   ```

3. Run the server:
   ```
   go run main.go
   ```

The API server will start on http://localhost:8080.

### Running the CLI Tool

1. In a separate terminal, run:
   ```
   go run cmd/cli/main.go
   ```

2. Use the interactive CLI to chat with the AI or view history.

### Running the Web Frontend

1. Navigate to the frontend directory:
   ```
   cd frontend
   ```

2. Install dependencies:
   ```
   npm install
   ```

3. Start the development server:
   ```
   npm start
   ```

The web application will open in your browser at http://localhost:3000.

## API Endpoints

- **POST /chat**
  - Sends a message to the AI and receives a response
  - Request body: `{ "prompt": "Your message here" }`
  - Response: `{ "response": "AI's response" }`

- **GET /history**
  - Retrieves conversation history
  - Response: Array of `{ "id": 1, "prompt": "User message", "response": "AI response", "timestamp": "2023-01-01T00:00:00Z" }`

## Development

### Building for Production

1. Build the backend:
   ```
   go build -o go-ai-app
   ```

2. Build the frontend:
   ```
   cd frontend
   npm run build
   ```