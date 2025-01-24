
# Introduction
This project implements a RESTful API for asynchronous image processing. The API allows clients to submit jobs consisting of multiple store visits. Each visit includes a store ID and a list of image URLs. The API processes these jobs by downloading the images, calculating their perimeters, and returning the results.

# 1. Description
- Asynchronous processing: jobs are proecessed in background.
- Json Store Data validation: Data is validated through the json file.
- Error handling: handles errors like - invalid requests, invalid store ids, image download failure and responds.
- Status Tracking - Ask teh api to return the status of the job and it will return completed/ ongoing response.
- Concurrent processing - Using goroutines the images are processed concurrently.
- Dockerized - Can be easily containerized using Docker.

# 2. Assumptions:
- Assuming "Each" image takes around 0.1-0.4 sec not number of visits, eg Let delay be 5,changing the no of images be 3 ie 15 sec , unline no of visits 2 ->10 sec.
- Using of goroutines and mutexes for concurrent image processing, assuming it as a small load.
- Flag a whole transaction as rejection if either of the payload info is wrong. eg visit[0] has all the right information and visit[1] does not have all the right information both the visits are considered as error.
- Taking AreaCode as a string/Text.

# 3. Installation (setup)
###  Without Docker

1. **Clone the repository:**

```bash
git clone https://github.com/adityarevankarp/KiranaClub-Assignment.git
cd KiranaClub-Assignment

#Get Dependencies
go get github.com/gorilla/mux

#Setup Go module
go mod init server

#build and run
go build
go run main.go

```

###  With Docker
```bash
docker build -t server .


docker run -p 8080:8080 server
```


# 4. Testing

## 1. Submitting a job
### API Endpoint (POST) - http://localhost:8080/api/submit/
```bash
{
    "count": 2,
    "visits": [
        {
            "store_id": "RP00001",
            "image_url": [
                "https://www.gstatic.com/webp/gallery/2.jpg",
                "https://www.gstatic.com/webp/gallery/3.jpg"
            ],
            "visit_time": "2024-07-27T10:00:00Z"
        },
        {
            "store_id": "RP0002",
            "image_url": [
                "https://www.gstatic.com/webp/gallery/3.jpg"
            ],
            "visit_time": "2024-07-27T11:00:00Z"
        }
    ]
}
```
### ExpectedOUtput if everything is right (201 Created).
```bash
{"job_id":3} 
```
### ExptectedOutput if count!=Visits
```bash
{"error": "Count mismatch"}
```
### ExptedOutput if Missing Fields or wrong store_id
```
{"error": "Invalid request payload"}
```
## 2. Get Job Info
### API ENDPOINT (GET) - http://localhost:8080/api/status?jobid=1
#### ExpectedOutput if ImageUrls are correct, StoreId is correct.
##### Responds with the job_id and the perimeter of the corresponding image.
```bash
{"job_id":1,"results":[{"store_id":"RP04941","image_url":"https://www.gstatic.com/webp/gallery/2.jpg","perimeter":1908},{"store_id":"RP04944","image_url":"https://www.gstatic.com/webp/gallery/3.jpg","perimeter":4000},{"store_id":"RP04941","image_url":"https://www.gstatic.com/webp/gallery/3.jpg","perimeter":4000}],"status":"completed"}
```
#### ExptectedOutput if the job is under processing
```bash
{
    "status": "ongoing",
    "job_id": 1
}
```
#### ExpectedOUtput if the store id is wrong..
```bash
{"error":["Invalid store_id: RP049asdf41"],"job_id":2,"status":"failed"}
```

#### ExpectedOUtput if the job id is not found
```bash
{}
```

# 5. Work Environment

- Text Editor/IDE: VS Code
- Libraries: github.com/gorilla/mux
- Operating System: Windows 11
- Go Version: go1.23.5
- Other Tools: Postman (for API testing).
- Docker: Docker Desktop 4.28.0 

# 6. Future Improvements
- Implement caching (e.g., Redis) for frequently accessed data like job statuses.
- Replace in-memory storage with a database like PostgreSQL for scalability.
- Implement a message queue (e.g., RabbitMQ, Kafka) to handle job submission asynchronously for high throughput.
- Write comprehensive unit tests for every function (e.g., edge cases for job creation and status retrieval).

