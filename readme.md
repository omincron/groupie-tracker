# Groopie Tracker

A web application that provides comprehensive information about music artists, including band members, concert locations, and performance dates. The application integrates with multiple APIs to deliver real-time data through a responsive user interface.

## Features

- **Artist Search**: Search functionality directly from the homepage
- **Detailed Artist Profiles**:
  - Band member information
  - First album release date
  - Band creation date
  - Concert locations and dates
- **Responsive Design**: Optimized for all device sizes
- **Real-time Data**: Integration with four distinct APIs

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (latest stable version)
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/groopie-tracker.git
cd groopie-tracker
```

2. Run the application:
```bash
go run .
```

3. Access the application:
```
http://localhost:8080
```

## Application Structure

```
groopie-tracker/
├── handlers/           # Route handlers
│   ├── artist.go      # Artist detail handler
│   ├── helpers.go     # Helper functions
│   ├── home.go        # Homepage handler
│   └── search.go      # Search functionality
├── models/            # Data models
│   ├── artist.go      # Artist model
│   ├── artist_full.go # Extended artist model
│   ├── date.go        # Date model
│   ├── location.go    # Location model
│   └── relation.go    # Relations model
├── services/          # API integration
│   └── api.go         # API service logic
├── static/            # Static assets
│   └── css/
│       └── styles.css # Main stylesheet
├── templates/         # HTML templates
│   ├── artist.html    # Artist detail page
│   ├── error.html     # Error page
│   ├── home.html      # Homepage
│   └── layout.html    # Base layout
├── main.go            # Application entry point
└── go.mod            # Go module file
```

## API Integration

The application integrates with four main APIs:

### 1. Artists API
Provides core artist information:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/artists" -H "Accept: application/json"
```

### 2. Locations API
Retrieves concert location data:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/locations" -H "Accept: application/json"
```

### 3. Dates API
Fetches concert dates:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/dates" -H "Accept: application/json"
```

### 4. Relations API
Manages location-date relationships:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/relation" -H "Accept: application/json"
```

**Tip**: For formatted JSON output, append `| jq` to any curl command:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/artists" -H "Accept: application/json" | jq
```

## Testing & Verification

1. **API Integration Test**:
   - Execute the curl commands above to verify API accessibility
   - Check response formats and data completeness

2. **Application Testing**:
   - Launch the application (`go run .`)
   - Navigate to `http://localhost:8080`
   - Verify homepage artist cards load correctly
   - Test artist detail pages for complete information
   - Confirm search functionality

3. **Responsive Design Test**:
   - Test on multiple screen sizes:
     - Desktop (1920×1080 and above)
     - Tablet (768×1024)
     - Mobile (375×667)
   - Verify layout adaptability
   - Check navigation usability

4. **Error Handling**:
   - Monitor application logs for API integration issues
   - Verify error page displays appropriately
   - Test search with invalid inputs

## Authors

Pkalliag
Cemvalot
Gpatoula
