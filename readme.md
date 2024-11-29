# Groopie Tracker

Groopie Tracker is a web application that provides information about music artists, their members, concert locations, and associated dates. It fetches data from external APIs, processes it, and displays it in a user-friendly interface.

---

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Integration](#api-integration)
  - [Testing API Endpoints with `curl`](#testing-api-endpoints-with-curl)
- [Project Structure](#project-structure)
- [Auditor Instructions](#auditor-instructions)
- [License](#license)

---

## Features
- Search for artists directly from the homepage.
- View artist details, including:
  - Band members
  - First album release date
  - Creation date
  - Concert locations and dates
- Fully responsive design for use on multiple devices.
- Data fetched dynamically from four APIs.

---

## Installation

### Prerequisites
1. [Go](https://go.dev/) installed on your system.
2. Git installed to clone the repository.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/groopie-tracker.git
   cd groopie-tracker
   ```
2. Run the application:
   ```bash
   go run .
   ```
3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

---

## Usage
- Search for an artist on the homepage.
- Click on an artist's card to view detailed information.
- Use the "Back to Home" button to return to the main page.

---

## API Integration

Groopie Tracker fetches data from the following external APIs:

1. **Artists API**: Provides details about artists, their members, and creation dates.
2. **Locations API**: Contains information about concert locations.
3. **Dates API**: Stores concert dates for artists.
4. **Relations API**: Links concert locations to corresponding dates.

---

### Testing API Endpoints with `curl`

To verify that the application correctly integrates with the APIs, use the following `curl` commands to fetch raw JSON data.

#### 1. Artists API
Fetch data about artists:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/artists" -H "Accept: application/json"
```

#### 2. Locations API
Fetch data about concert locations:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/locations" -H "Accept: application/json"
```

#### 3. Dates API
Fetch data about concert dates:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/dates" -H "Accept: application/json"
```

#### 4. Relations API
Fetch data about location-date relations:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/relation" -H "Accept: application/json"
```

#### Optional: Pretty-Print JSON Responses
To format the JSON response for easier readability, use `jq`:
```bash
curl -X GET "https://groupietrackers.herokuapp.com/api/artists" -H "Accept: application/json" | jq
```

---

## Project Structure
```
groopie-tracker/
├── main.go               # Main entry point
├── handlers/             # Route handlers
│   ├── home.go           # Homepage handler
│   ├── artist.go         # Artist detail handler
├── models/               # Data models
│   ├── artist.go         # Artist model
│   ├── location.go       # Location model
│   ├── date.go           # Date model
│   ├── relation.go       # Relations model
├── services/             # API service logic
│   ├── api.go            # Functions for fetching API data
├── templates/            # HTML templates
│   ├── layout.html       # Base layout
│   ├── home.html         # Homepage
│   ├── artist.html       # Artist details page
├── static/               # Static files
│   ├── css/              # CSS files
│       ├── styles.css    # Main stylesheet
```

---

## Auditor Instructions

To verify the application's functionality, follow these steps:

1. **Check API Integration**:
   Use the provided `curl` commands under [Testing API Endpoints with `curl`](#testing-api-endpoints-with-curl) to ensure that each API is accessible and returns the expected JSON data.

2. **Run the Application**:
   - Start the application using `go run .`.
   - Open `http://localhost:8080` in your browser.

3. **Verify API Data in the Application**:
   - Ensure that the homepage displays artist cards. This verifies that the **Artists API** is functioning.
   - Click on an artist card to view details. Check that:
     - Concert locations are displayed correctly (from the **Locations API**).
     - Concert dates are accurate (from the **Dates API**).
     - Location-date relations are correct (from the **Relations API**).

4. **Inspect Logs**:
   If any data is missing or incorrect, inspect the application logs for errors during API fetching. Logs will provide details about malformed responses or network issues.

5. **Test Responsiveness**:
   - Open the application in different screen sizes (e.g., desktop, tablet, mobile).
   - Ensure the layout adapts appropriately.
