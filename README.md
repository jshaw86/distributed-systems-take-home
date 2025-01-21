# Project Overview

## Purpose
This project aims to evaluate your ability to comprehend and make critical decisions within the domain of distributed systems and to get you thinking about
considerations around trust and safety as it pertains to rate limiting and anomaly detection. 

## Installing Docker
To get started, install Docker by following the appropriate guide for your operating system:

- [Docker Desktop for Mac](https://docs.docker.com/desktop/setup/install/mac-install/)
- [Docker Desktop for Windows](https://docs.docker.com/desktop/setup/install/windows-install/) 

Alternatively, on macOS, you can install Docker using Homebrew:

```bash
brew install docker docker-compose
```

Ensure that the necessary environment variables are correctly set in your terminal (`PATH`) and your docker daemon is running through homebrew.

## Quick Start
Run the provided Go HTTP client alongside the echo server. Ensure Docker and Docker Compose are installed and running.

```bash
docker-compose up --build
```

### Sample Output

```bash
docker-compose up
[+] Running 2/2
 ✔ Container distributed-systems-take-home-echo-server-1  Created                                                                                                                                                0.0s
 ✔ Container distributed-systems-take-home-client-1       Created                                                                                                                                                0.0s
Attaching to client-1, echo-server-1
client-1       | 2025/01/21 03:59:00 Servers:  [http://echo-server:80]
echo-server-1  | Listening on port 80.
echo-server-1  | {"name":"echo-server","hostname":"2b48b8df0746","pid":1,"level":30,"host":{"hostname":"echo-server","ip":"::ffff:172.20.0.3","ips":[]},"http":{"method":"POST","baseUrl":"","originalUrl":"/send","protocol":"http"},"request":{"params":{},"query":{},"cookies":{},"body":{},"headers":{"host":"echo-server:80","user-agent":"Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X)","content-length":"39","x-ip-address":"134.45.149.242","x-session-id":"XKwCRPXbjS","accept-encoding":"gzip"}},"environment":{"PATH":"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin","HOSTNAME":"2b48b8df0746","NODE_VERSION":"20.11.0","YARN_VERSION":"1.22.19","HOME":"/root"},"msg":"Tue, 21 Jan 2025 03:59:00 GMT | [POST] - http://echo-server:80/send","time":"2025-01-21T03:59:00.327Z","v":0}
```

> **Note:** It is normal for some initial client HTTP requests to fail due to an imperfect readiness detection mechanism of the echo server.

## Client Overview
The provided `docker-compose.yml` file sets up a custom Go HTTP client that generates random HTTP requests to a backend server. An echo server is included to facilitate inspection of the requests and headers.

### Sample Request
The key fields to observe in the request headers are `x-session-id` and `x-ip-address`:

```json
{
    "name": "echo-server",
    "request": {
        "params": {},
        "query": {},
        "cookies": {},
        "body": {},
        "headers": {
            "host": "echo-server:80",
            "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
            "content-length": "0",
            "x-ip-address": "220.127.113.2",
            "x-session-id": "WPQAgorGfK",
            "accept-encoding": "gzip"
        }
    }
}
```

## Request Headers Explained
- **`x-session-id`**: Represents a user ID or unique session identifier. There is a probability that a new session/user id is generated on each request up to a max number of users.
- **`x-ip-address`**: Simulates an `X-Forwarded-For` header, representing the client's IP address.
- **`user-agent`**: Mimics standard user-agent behavior; the client randomly selects a user agent.
- **Sticky Behavior**: The `user-agent` and `x-ip-address` remain constant for a session ID, except in cases of anomalies.

## Anomalies
During each `POST /send` request, there's a probability that the same `x-session-id` will be assigned a different `x-ip-address` and `user-agent`. The likelihood varies by session:

- Some session IDs may rarely, if ever, change their IP address or user-agent.
- Others may experience changes frequently.
- It is up to you to determine which requests should be allowed or blocked.

## Project Requirements
1. **Rate Limiting Implementation**  
   Develop a backend service to enforce the following rate limiting policies:
   - Maximum of **30 requests per minute** globally.
   - Maximum of **5 requests per minute** per session ID.
   The backend should return a `429 Too Many Requests` status code when these limits are exceeded. The client will backoff and retry if it receives a `429` response.

> **Note:** The client has a hard max of 20 total user ids, consider this when architecting the state for your solution. 

2. **Scalability Consideration**  
   Deploy a second backend instance and ensure the rate-limiting policies still hold. Update the `SERVERS` environment variable in `docker-compose.yaml` with multiple backend instances:

   ```yaml
   environment:
     - SERVERS=http://backend1:8080,http://backend2:8080
   ```

   You may use a load balancer, but the frontend was designed to accept multiple backends to simplify the infrastructure setup.

3. **Anomaly Detection and Mitigation**  
   Identify and block suspicious request patterns based on changes in `x-ip-address` and `user-agent` for the same session ID.

4. **Technology Choice**  
   Use any programming language, runtime, or framework as long as it functions within the Docker Compose environment. Running `docker-compose up --build` should work seamlessly.

5. **External Dependencies**  
   Any additional services or datastores must operate within an ~8GB memory footprint and integrate with the `docker-compose` setup.

6. **Immutable Client Code**  
   The provided client code should be considered immutable. If you discover a bug, feel free to fix it or let us know and we'll work with you to fix it. Don't add/update the request with additional headers or data.

7. **Documentation**  
   Extend the bottom of this `README.md` file to include an explanation of your solution, rationale for architectural choices, and anomaly detection/enforcement strategies.


# Suggestions
1. **Get something working first** 
    Start with the distributed rate limiting problem on a single instance and build out from there.
2. **Don't build for massive scale**  
    For this exercise we don't expect your solution to support millions of requests and users. In your README.md, explain how you would scale your solution 
    if you had to support a much larger user base and request throughput. We like to see pragmatic solutions that can be iterated on.
3. **Code Style, Testing and your time**
   Code style, effective use of comments and testing are important but we understand your time is valuable and limited. Don't get bogged down if those considerations are adding too much time to the exercise.
   We're more interested in a working solution and how you approach the problem. This project is designed to be completed in 2-4 hours depending on your experience level. If it's taking you longer than that
   consider submitting your solution without the anomaly detection and mitigation portion and document in your README.md how you would have approached that section.

