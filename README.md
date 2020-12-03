# Background

Our production environments utilise AWS cloud platform with the infrastructure being created and controlled by Hashicorp's Terraform. Our deployment process takes the app code repository, generates a series of docker images which are then placed into AWS Elastic Cluster Service.

# Task

This repository contains a Go application (with `./app` folder) that acts as an API; it is configured to respond on port `8080` and has two endpoints (`/` and `/status`) that return `json` content.

The `/status` endpoint returns different content and `http` response code depending on the value of an environment variable called `APP_STATUS`. When `APP_STATUS=OK` the `/status` endpoint will return a `200` code and message, otherwise it will return a `500` code.

We would like you to create a fork of this repository which expands this code to provide a way to create a working docker image from this application with both endpoints returning successful response codes.

Please include an updated README with steps to produce and run the docker image.

**Do not worry if you are not able to fully complete the task, we are more interested in your approach then a perfect solution.**


## Outline

To help guide you in this task, we've outlined some typical steps (in no particular order) that would be needed.

- Creating a Dockerfile
- Building the Go app
- Mapping ports
- Setting up Docker on your local machine
- Passing environment variables


## My Approach to the tech test

Both Golang and Docker are technologies that I am aware of and completed training on, albeit my usage of them have been limited.

So I started by identifying steps to take to complete the task.

First of all, do I have all the applications required? 

    1 - Golang
    
    2 - Docker
    
Both of these were already installed so we are good here. 

Secondly, at a high-level outline steps to deliver the requirements 

    1 - Fork the application and create a local copy
    
    2 - Locate Docker images from the official Docker website
    
        a - I am setting this up on a Windows laptop but will create a linux based container so a linux distro
        
        b - The Golang image compatible for linux distro- https://hub.docker.com/_/golang
        
    3 - I familiarised myself with the application file and run it to observe the behaviour.
    
    4 - This was followed by jotting down 2 ways I thought I could take to complete the task.

My initial thoughts were to set the environment variable in the Docker file and then overwrite when running the container.

Then, I thought about potentially setting the environment variable dynamically through an API call and that would involve changing the script, apologies for that, and adding another endpoint to cater for this. And this is the route I went down to, purely to enable the toggling of the response code and message through the API.

You may have also noticed that I have changed couple of bits in the script. I altered the variable name from status_from_env to statusFromEnv as my editor was complaining about the underscores. I love my underscores in variable names in Python but after searching online for variable naming conventions it turns out it is frowned upon by Gophers... over dramatic :-)

I have also removed couple of else blocks and indented out the statements contained within. My thought process for that was the presence of the 'return' statement in the if block kind of took care of the flow. I appreciate it is down to preference totally.

Once I made the changes, I compiled the go script and run it to test the output.

Once I was happy with the results I created the Docker file.

# Build Docker image
To build the docker file I used the below on a Cmder shell, but command prompt would work just as well.

    docker build -t go_on_alpine .

# Check Docker image
This will pull all required components and create your docker image. You can check it is there using

    docker image ls

# Run Docker image
Running the Docker image will create your container. I run it in detached mode and set to port to 8080.

    docker run -d -p 8080:8080 go_on_alpine

# Check Docker container is running
    docker container ls

Now you can perform the following tests to see behaviour.

    Navigating to http://localhost:8080/ produced
        {"code":200,"message":"Home"}

    Navigating to http://localhost:8080/status/ produced
        {"code":500,"message":""}

    Navigating to http://localhost:8080/setStatus/?key=OK produced
        {"code":200,"message":"APP_STATUS SET"}

    Re-Navigating to http://localhost:8080/status/ produced
        {"code":200,"message":"OK"}

# Stop Docker container
Once happy you can stop the container by using the first 4 digits of the container id displayed as part of the 'ls' command.

    docker container stop <container_id>

# Remove Docker container
    docker rm <container_id>

# Alternatively pull the image from Docker Hub
    docker pull gogi1973/go_on_alpine:latest

# And run container
    docker run -d -p 8080:8080 gogi1973/go_on_alpine
