# Advertisement-Go

**Advertisement-Go** is an open-source project built in Go, leveraging MongoDB to deliver a fast and responsive service for managing geolocation-based advertisements. This service allows clients to retrieve advertisements based on their geographical location by simply passing their coordinates. It also supports more refined queries that consider the distance from the user, enabling personalized and location-specific advertisement delivery.

## Features

- **Geolocation Queries**: Quickly retrieve advertisements relevant to your current location.
- **Distance Filtering**: Customize your search with distance parameters to find ads within a specific range.
- **High Performance**: Designed with performance in mind to handle requests swiftly.
- **Open Source**: Freedom to modify, distribute, and use the software according to your needs.

## Getting Started

### Prerequisites

- Go (Version 1.x or later)
- MongoDB (Version 4.x or later)
- Any REST client (for API testing)

### Installation

1. Clone the repository to your local machine:
   ```
   git clone https://github.com/yourgithubusername/advertisement-go.git
   ```
2. Navigate to the project directory:
   ```
   cd advertisement-go
   ```
3. Build the project (ensure Go is installed):
   ```
   go build
   ```
4. Run the server:
   ```
   ./advertisement-go
   ```

### Usage

**Retrieving Advertisements by Location**

Simply make a GET request to the server with your current longitude and latitude:

```
http://localhost:1323/<longitude>/<latitude>
```

Example:
```
http://localhost:1323/-46.6773326/-23.5800755
```

**Posting a Query with Distance Filtering**

To find advertisements within a specific distance from your location, send a POST request with the following JSON body:

```json
{
  "latitude": -23.5800755,
  "longitude": -46.6773326,
  "maxDistance": 10
}
```

This will return advertisements within 10 kilometers of the provided coordinates.

## Contributing

We welcome contributions from everyone. If you have an idea for improvement, please fork the repository, make your changes, and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is open-sourced under the [MIT License](LICENSE.md).

---

This template provides a basic structure for your README.md, including sections that explain what the project does, how to get started, how to use the service, and how to contribute. Make sure to replace placeholders (like the GitHub URL) with actual data from your project, and expand each section based on the specifics of your application and its deployment.