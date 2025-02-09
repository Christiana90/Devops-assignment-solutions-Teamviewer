Sure! Here's the updated version of the `README.md` file, which now includes the Helm chart instructions:

---

# DevOps Interview Assignment

## Overview

This repository contains the solution for the DevOps Interview Assignment for TeamViewer GmbH. The task is to deploy a Go application that calculates the mass of a sphere or cube based on the dimensions passed in through HTTP requests. The application is containerized using Docker, and Kubernetes resources are deployed using Helm.

### The Project Includes:
1. **Go Application**: The core Go application that calculates the mass of a sphere and cube based on the dimension provided.
2. **Dockerfile**: The Dockerfile to build and run the Go application inside a container.
3. **Helm Chart**: A Helm chart to deploy the containerized Go application on a Kubernetes cluster. The port is configurable via the Helm chart's `values.yaml` file.

---

## Getting Started

### 1. Prerequisites
- Install **Go** version 1.23 or higher
- Install **Docker** for containerization
- Install **Helm** for Kubernetes deployments
- A **Kubernetes** cluster (Minikube or a cloud-based cluster like GKE, AKS, or EKS)

---

### 2. Running the Go Application Locally

#### Step 1: Clone this repository

```bash
git clone https://github.com/your-repository.git
cd your-repository
```

#### Step 2: Run the Go application

Run the Go application using the following command:

```bash
go run main.go 8083
```

This will start the server on `http://localhost:8083`. You can access the endpoints:

- `/aluminium/sphere?dimension=<radius>`: Calculates the mass of an aluminum sphere based on the radius.
- `/iron/cube?dimension=<side>`: Calculates the mass of an iron cube based on the side length.

http://localhost:8083/aluminium/sphere?dimension=3

---

### 3. Building and Running the Docker Container

#### Step 1: Build the Docker image

Build the Docker image using the following command:

```bash
docker build -t my-go-app .
```

#### Step 2: Run the Docker container

Run the container with the desired port (default is 8081):

```bash
docker run -p 8083:8083 my-go-app
```

You can change the port by setting the `PORT` environment variable:

```bash
docker run -p 8080:8080 -e PORT=8080 my-go-app
```
This will start the container server on `http://localhost:8083`. You can access the endpoints:

- `/aluminium/sphere?dimension=<radius>`: Calculates the mass of an aluminum sphere based on the radius.
- `/iron/cube?dimension=<side>`: Calculates the mass of an iron cube based on the side length.

http://localhost:8083/aluminium/sphere?dimension=3
http://localhost:8083/iron/cube?dimension=3

---

### 4. Using Helm to Deploy the Application on Kubernetes

#### Step 1: Install Helm (if not already installed)

Follow the installation instructions from the [Helm website](https://helm.sh/docs/intro/install/).

#### Step 2: Create a Helm Chart

The Helm chart is located in the `my-go-app-chart/` directory. It contains the resources required to deploy the Go application on Kubernetes.

#### Step 3: Update the Helm chart values

Open the `values.yaml` file and modify the `port` value to the desired port if necessary.

#### Step 4: Install the Helm chart

You can install the Helm chart on your Kubernetes cluster using the following command:

```bash
helm install my-go-app ./my-go-app-chart --set port=8081
```

This will deploy the Go application on your Kubernetes cluster and expose it on the configured port.

#### Step 5: Access the application

Once the application is running on the Kubernetes cluster, access it via the service URL provided by your Kubernetes provider or using port-forwarding:

```bash
kubectl port-forward svc/my-go-app-service 8081:8081
```

You can now access the application at `http://localhost:8081`.

---

## Folder Structure

```
my-go-app/
├── main.go
├── Dockerfile
├── README.md
└── chart/
    ├── Chart.yaml
    ├── values.yaml
    └── templates/
        ├── deployment.yaml
        └── service.yaml
```

---

## Files Overview

- **`main.go`**: The Go application that calculates the mass of a sphere or cube based on dimensions passed via HTTP queries.
- **`Dockerfile`**: The Dockerfile used to containerize the Go application.
- **`chart/Chart.yaml`**: The metadata for the Helm chart.
- **`chart/values.yaml`**: The configurable values, including the port where the application will run.
- **`chart/templates/deployment.yaml`**: Kubernetes deployment configuration.
- **`chart/templates/service.yaml`**: Kubernetes service configuration for exposing the application.

---

## Conclusion

This solution provides a simple yet complete setup for deploying a Go application with Docker and Kubernetes using Helm. The port for the Go application is made configurable via the Helm chart's `values.yaml` file, making it easy to deploy and configure the app in different environments.

---

