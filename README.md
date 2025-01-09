# Readcommend

Readcommend is a book recommendation web app for the true book aficionados and disavowed human-size bookworms. It allows to search for book recommendations with best ratings, based on different search criteria.

# Instructions

The front-end single-page app has already been developed using Node/TypeScript/React (see `/app`) and a PostgreSQL database with sample data is also provided (see `/migrate.*`). Your mission - if you accept it - is to implement a back-end microservice that will fetch the data from the database and serve it to the front-end app.

- In the `service` directory, write a back-end microservice, using the back-end language of your choice (Golang, Java, Python, NodeJS, C# or the language that was aligned during your initial interview) that listens on `http://localhost:5001`.
- Write multiple REST endpoints, all under the `/api/v1` route, as specified in the `/open-api.yaml` swagger file.
- The most important endpoint, `/books`, must return book search results in order of descending ratings (from 5.0 to 1.0 stars) and filtered according to zero, one or multiple user selected criteria: author(s), genre(s), min/max pages, start/end publication date (the "era"). A maximum number of results can also be specified.
- It's OK to use libraries for http handling/routing and SQL (ie: query builders), but try to refrain from relying heavily on end-to-end frameworks (ie: Django) and ORMs that handle everything and leave little room to showcase your coding skills! ;)
- Write some documentation (ie: at the end of this file) to explain how to deploy and run your service. *Please assume that the reviewer should be able to run your solution without needing to install extra dependencies on his machine.*
- Keep your code simple, clean and well-organized.
- During development commit and push regularly to the CodeSubmit git repository and, when you are done, submit your assignment in CodeSubmit.
- Don't hesitate to come back to us with any questions along the way. We prefer that you ask questions, rather than assuming and misinterpreting requirements.
- You have no time limit to complete this exercise, but the more time you take, the higher our expectations in terms of quality and completeness.
- No changes in the front-end app are needed. The front-end app will call your back-end microservice to fetch the data.
- If you see improvement to API design, explain your recommendation in the readme.MD.  The spec should not be changed.
- You will be evaluated mainly based on how well you respect the above instructions. Please refer to the **Expectations** section below for more details on how you will be evaluated.

# Expectations

nesto recruits for all proficiency levels. Depending on your experience and the position you are applying for, what we expect to see in the test will differ. We do, however, understand that you may have a life (some people do). If you don't have the time to respect all the instructions, simply do your best and focus on what you deem most important.

**Junior Submission**

- Provides a solution that is bug free, fast and meets all the requirements listed above.
- Code provided shows a good proficiency in the chosen language.
- Shows good reflexes with regards to 3rd party dependencies. You should add dependencies that add value to your solution without solving everything for you.
- Basic testing of the core functionalities.

**Senior Submission**

We expect everything from a *Junior Submission* and that you treat this as production ready code without necessarily providing full test coverage
- Solution is architectured properly to allow for adequate test coverage and future extensions.
- Each layer of the provided solution must be protected against typical attack vectors. You should **NOT** rely on your API contract to protect you from malicious inputs.
- Showcases the type of tests needed to validate your architectural choices and to provide confidence that your solution satisfies the 'business' and performance requirements.

# Development environment

## Docker Desktop

Make sure you have the latest version of Docker Desktop installed, with sufficient memory allocated to it, otherwise you might run into errors such as:

```
app_1         | Killed
app_1         | npm ERR! code ELIFECYCLE
app_1         | npm ERR! errno 137.
```

If that happens, first try running the command again, but if it doesn't help, try increasing the amount of memory allocated to Docker in Preferences > Resources.

## Starting front-end app and database

In this repo's root dir, run this command to start the front-end app (on port 8080) and PostgreSQL database (on port 5432):

```bash
$ docker-compose up --build
```

(later you can press Ctrl+C to stop this docker composition when you no longer need it)

Wait for everything to build and start properly.

## Connecting to database

During development, you can connect to and experiment with the PostgreSQL database by running this command:

```bash
$ ./psql.sh
```

Then, on the psql CLI, test as follows:

```psql
readcommend=# \dt
```

If everything went well, you should get this result:

```psql
    List of relations
 Schema |  Name  | Type  |  Owner
--------+--------+-------+----------
 public | author | table | postgres
 public | book   | table | postgres
 public | era    | table | postgres
 public | genre  | table | postgres
 public | size   | table | postgres
(5 rows)
```

To exit the PostgreSQL session, type `\q` and press `ENTER`.

## Accessing front-end app

Point your browser to http://localhost:8080

Be patient, the first time it might take up to 1 or 2 minutes for parcel to build and serve the front-end app.

You should see the front-end app appear, with all components displaying error messages because the back-end service does not exist yet.

# Deploying and running back-end microservice
Here's a README.md template for setting up a Golang project:

## Prerequisites

Before you start, make sure you have the following installed:

1. **Docker Desktop** (latest version)
   - Allocate sufficient memory under Preferences > Resources.
2. **Git** (to clone the repository)
3. **Go (Golang)** (latest stable version)
4. **PostgreSQL** (or use a Docker container for the database)

---

## Setup Instructions

### 1. Clone the Repository

Run the following commands to clone the repository:

```
$ git clone <repository-url>
$ cd <repository-folder>/service

for example:  
$ git clone http://nesto-cccglh@git.codesubmit.io/nesto/back-end-code-challenge-tsvihr
$ cd back-end-code-challenge-tsvihr/service
```
### 2. Install Dependencies
Ensure you are in the project’s root directory. Run the following command to install all necessary dependencies:

```Copy code
$ go mod tidy
```
### 3. Set Up Environment Variables
Create a .env file in the root directory of the project with the following contents:

env
```
PORT=5001
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password123
DB_NAME=readcommend
```
Adjust these values based on your setup (especially database credentials).


### 4. Start PostgreSQL Database (Optional)
If you have Docker installed, you can use the following Docker command to spin up a PostgreSQL container. Alternatively, if you already have PostgreSQL set up, skip this step.

```bash
$ docker run --name readcommend-db -e POSTGRES_PASSWORD=password123 -e POSTGRES_DB=readcommend -p 5432:5432 -d postgres:latest
```
This will run a PostgreSQL instance on port 5432.

### 5. Start the Golang Application
In the root directory of the project, run the following command to start the backend service:

```bash
$ go run main.go
```
This will start the backend service, which will be available at http://localhost:5001.

### 6. Verify the Service
Open your browser or use tools like Postman or cURL to test the API:

Open http://localhost:5001/api/v1/books to check if the service is running correctly.
## Endpoints Overview
### Base URL
All endpoints are prefixed with /api/v1. Example: http://localhost:5001/api/v1/books

### Key Endpoints
### 1. GET /books**
Fetch books filtered by criteria and sorted by ratings.
- Query Parameters:
   - `authors` (optional, comma-separated list)
   - `genres` (optional, comma-separated list)
   - `min_pages` / `max_pages` (optional)
   - `start_date` / `end_date` (optional, YYYY-MM-DD format)
   - `max_results` (optional, integer)
- Example Request:

```http
GET /api/v1/books?authors=J.K.Rowling&genres=Fantasy&max_results=10
```
- Response:
```json
[
  {
    "id": 1,
    "title": "Harry Potter and the Philosopher's Stone",
    "author": "J.K. Rowling",
    "genre": "Fantasy",
    "rating": 4.9,
    "pages": 223,
    "publication_date": "1997-06-26"
  }
]
```
### 2. GET /authors
Retrieve available authors.

### 3. GET /genres
Retrieve available genres.

### 4. GET /sizes
Retrieve available book sizes (e.g., small, medium, large).

### 5. GET /eras
Retrieve available publication eras.

## Testing the Service
### Automated Tests
Run the following command to execute unit tests:

```bash
$ go test ./...
```
Manual Testing
You can manually test the API using tools like Postman or cURL.

Example with cURL:

```bash
$ curl "http://localhost:5001/api/v1/books?max_results=5"
```
## Notes and Recommendations
### Improvements to API Design
1. **Pagination: Add support for paginated responses to manage large datasets.
-Query Parameters: page and page_size.
2. **Error Handling: Implement consistent error structures for responses.

Example Error Response:

```json
{
  "error": "Invalid parameter",
  "details": "The start_date parameter is incorrectly formatted."
}
```
### Security Considerations
- Sanitized all user inputs to prevent SQL injection attacks.
- Use HTTPS in production to ensure secure API communication.
---
## Clean-Up
- After testing, you can stop all services by running:

```bash
$ docker-compose down
```
Or, if you are using PostgreSQL in a Docker container:

```bash
Copy code
$ docker stop readcommend-db
$ docker rm readcommend-db
```