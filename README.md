
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
