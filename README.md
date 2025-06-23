# Model API Backend

A Go backend API built with Gin framework and GORM for managing states, models, and FAQs.

## 🚀 Deployment to Vercel

### Prerequisites
- Node.js and npm installed
- Vercel CLI installed: `npm install -g vercel`
- GitHub account for repository hosting

### Steps to Deploy

1. **Install Vercel CLI**
   ```bash
   npm install -g vercel
   ```

2. **Initialize Git and push to GitHub**
   ```bash
   git init
   git add .
   git commit -m "Initial commit: Go API for Vercel"
   # Create a new repository on GitHub and push
   git remote add origin https://github.com/yourusername/your-repo-name.git
   git push -u origin main
   ```

3. **Deploy to Vercel**
   ```bash
   vercel
   ```
   
   Follow the prompts:
   - Link to existing project? → No
   - What's your project name? → model-api (or your preferred name)
   - Which directory? → . (press Enter)
   - Want to overwrite settings? → No

4. **Set Environment Variables**
   
   Go to Vercel Dashboard > Your Project > Settings > Environment Variables
   
   Add these variables:
   ```
   DB_HOST=your-mysql-host
   DB_USER=your-mysql-user
   DB_PASSWORD=your-mysql-password
   DB_NAME=your-mysql-database
   DB_PORT=3306
   ```

5. **Redeploy with Environment Variables**
   ```bash
   vercel --prod
   ```

### API Endpoints

After deployment, your API will be available at:
`https://your-project-name.vercel.app/api/`

#### Available Endpoints:
- `GET /api/states` - Get all states
- `POST /api/states` - Create a new state
- `GET /api/states/:id` - Get state by ID
- `PUT /api/states/:id` - Update state
- `DELETE /api/states/:id` - Delete state
- `GET /api/states/:id/models` - Get models by state ID

- `GET /api/models` - Get all models
- `POST /api/models` - Create a new model
- `GET /api/models/:id` - Get model by ID
- `PUT /api/models/:id` - Update model
- `DELETE /api/models/:id` - Delete model

- `GET /api/faq` - Get FAQ
- `POST /api/faq` - Create/Update FAQ

- `GET /api/global-phone` - Get global phone number
- `POST /api/global-phone` - Create/Update global phone number

## 🏗️ Project Structure

```
├── api/
│   └── main.go          # Vercel serverless function
├── config/
│   └── database.go      # Database configuration
├── controller/
│   ├── model_controller.go
│   ├── state_controller.go
│   ├── faq_controller.go
│   └── global_phone_controller.go
├── models/
│   └── models.go        # Data models
├── repository/
│   ├── model_repository.go
│   ├── state_repository.go
│   ├── faq_repository.go
│   └── global_phone_repository.go
├── routes/
│   └── routes.go        # API routes
├── service/
│   ├── model_service.go
│   ├── state_service.go
│   ├── faq_service.go
│   └── global_phone_service.go
├── main.go              # Local development server
├── go.mod
├── go.sum
├── vercel.json          # Vercel configuration
└── README.md
```

## 🔧 Local Development

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📝 Notes

- The API uses MySQL database
- CORS is configured for multiple domains
- Auto-migration is enabled for database schema
- File uploads are handled as base64 strings
- Global phone number maintains only one record in the database

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Make (optional, for using Makefile commands)

## Setup

1. Clone the repository
2. Create a PostgreSQL database named `model_db`
3. Copy `.env.example` to `.env` and update the database credentials if needed
4. Install dependencies:
   ```bash
   go mod download
   ```
5. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

### States

- `POST /states` - Create a new state
- `GET /states` - Get all states
- `GET /states/:id` - Get a specific state
- `GET /states/:id/models` - Get all models for a specific state

### Models

- `POST /models` - Create a new model
- `GET /models` - Get all models
- `GET /models/:id` - Get a specific model

## Project Structure

```
.
├── config/         # Configuration files
├── controller/     # HTTP handlers
├── models/         # Database models
├── repository/     # Database operations
├── routes/         # Route definitions
├── service/        # Business logic
├── .env           # Environment variables
├── go.mod         # Go module file
├── main.go        # Application entry point
└── README.md      # This file
```

## Database Schema

### States Table
- id (primary key)
- name
- created_at
- updated_at
- deleted_at

### Models Table
- id (primary key)
- state_id (foreign key)
- phone_number
- description
- created_at
- updated_at
- deleted_at 