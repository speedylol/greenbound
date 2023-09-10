# greenbound
 
# greenbound
 
## Inspiration

**GreenBound** is a web application/Framework dedicated to uniting communities impacted by environmental racism and advocating for environmental justice. Our mission is to empower individuals, amplify voices, and foster collaboration to combat the unequal distribution of environmental burdens
## What it does

Using either the search bar or dropdown menu, user can select any of the 100 counties in North Carolina, GreenBound's page will update showing environmental data for that county.

## How we built it
GreenBound is built with **Golang (Go)** , **HTMX**, and the **EJSCREEN API** [link](https://www.epa.gov/ejscreen/what-ejscreen)

Go as the Backend:

Go serves as the backend of the GreenBound application, handling data processing, API requests, and rendering HTML templates. Here's how it works:

    HTTP Routing: Go's "net/http" package is used to define routes and handlers for incoming HTTP requests. Each route corresponds to a specific URL endpoint, such as "/search."

    API Integration: Go makes HTTP requests to external APIs, such as the EJSCREEN API, to fetch environmental data. The data is retrieved in JSON format.

    Data Processing: Go processes the JSON data received from the external APIs. It can parse the JSON into Go structs, making it easier to work with the data.

    Template Rendering: Go uses an HTML template engine (e.g., the standard library's "html/template" package) to render dynamic HTML pages. These templates can incorporate data fetched from external APIs and data stored in the backend.



## Challenges we ran into

While building the Go script to make HTTP requests to EJScreen, I ran into troubles parsing the JSON response. I am more familiar with this process in Python or Javascript.


## Accomplishments that we're proud of

I got the JSON Streams to work!
```go
type Main struct {
	Data Data `json:"data"`
}

type Data struct {
	Demographics Demographics       `json:"demographics"`
	Main         map[string]*string `json:"main"`
	Extras       map[string]*string `json:"extras"`
}
```

GreenBound can now retrieve data from the API and use it to dynamically update the webpage

## What I learned

I learned a lot about the HTMX framework and gained experience using Golang for web development!

## What's next for greenbound

GreenBound will need polishing for the UI ()

It is intended to be used not just to search









