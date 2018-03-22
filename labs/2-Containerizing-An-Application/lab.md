# Containerizing an Application

This lab will begin our DevOps journey by running a simple Golang application locally using Docker. Containers are at the core of how Kubernetes, and many other DevOps implementations, operate. This lab will set the stage for later, more advanced, labs.

## About the Application
The source code for the application located in the `src/link-unshorten` directory is a simple API that we will eventually Dockerize and deploy using Kubernetes. The API takes a shortened link as a parameter in the URL and returns the destination URL as JSON. The application is built using a lightweight HTTP web framework called [Gin-Gonic](https://github.com/gin-gonic/gin).

### Task 1: Browse the Application
Open up the files in `src/link-unshorten` in your favorite IDE or text editor and familiarize yourself with the application.

### Task 2 (optional - do not spend too much time on this): Run the Golang Application
If you have Golang and [Dep](https://github.com/golang/dep) (Golang dependency management tool) installed locally, go ahead and run the following commands to fire up the application:

1. Copy the contents of the `src/link-unshorten` directory to your GOPATH. The GOPATH is often located at `~/go/src` on Mac OS. 

2. Run `dep ensure` to ensure the `vendor` directory is in the correct state for your configuration.

3. In the `src/link-unshorten` directory of the directory copied to your GOPATH, run `go run *.go` to start the application.

4. Visit `http://localhost:8080/api/check?url=bit.ly/test`

### Task 3: Build the Docker Container Locally
In the `src/link-unshorten` directory run the following command (substituting <yourname> with your own identifier) to build the image on you laptop:
```
docker build -t <yourname>/link-unshorten:0.1 .
```

Inspect your local Docker images:
```
docker images
```

### Task 4: Run the Docker Image
To run the application using Docker use the following command:
```
docker run -d -p 8080:8080 <yourname>/link-unshorten:0.1
```

Make sure the image is running without any errors:
```
docker ps
```
Visit `http://localhost:8080/api/check?url=bit.ly/test`

Note: If you run into port conflict errors, make sure to kill the Golang application if you ran the application manually in Task 2 or any other applications that may be running on port 8080.

### Bonus: Make the port in main.go an environment variable instead of a hardcoded value

Hint 1: Check out the [OS](https://golang.org/pkg/os) package for Golang

Hint 2: Use the `-e` flag to pass an environment variable into your `docker run` command

Hint 3: Yes, the answer is commented in the source code

Hint 4: You will need to run `docker stop` on the first running container before running another one with the same port
 
### Bonus+ : Scan the Dockerfile for vulnerabilities using Clair 
Documentation can be found in the Clair [Github](https://github.com/arminc/clair-scanner) repository:


Hint 1: After Clair is installed, the command will looking something like this:
```
./clair-scanner --ip 192.168.67.98 link-unshorten:0.1
```
Hint 2: You will need to use the IP address of en0 (192.168.67.98 in the above command) to run Clair

### Bonus++ Create a whitelist for a few CVEs and print the output to .json
Hint: Enforce the whitelist located in the `clair-whitelist` directory (only use this in production if you want your AppSec team to cry).

### Discussion Question: How would you plug this into your current CI/CD pipeline? 