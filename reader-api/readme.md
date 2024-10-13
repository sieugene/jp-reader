# jp-reader

jp-reader is a monorepository that includes a Go and Flask backend for image processing using Mokuro. The project is structured to enable efficient interaction between microservices and optimal handling of uploaded data.

## Key Features

1. **Image Uploading**: The application allows users to upload images, which are then processed using Mokuro.
2. **Image Processing**: Mokuro is used for text recognition and image conversion to facilitate further content handling.
3. **Message Queues**: RabbitMQ provides asynchronous data processing, organizing task queues for managing file uploads and processing.
4. **Processing Status Tracking**: The system tracks the status of uploaded images, allowing users to see the processing progress.

## Tech Stack

- **Go**: The primary backend language, handling routing, database interactions, message queue management, and file processing.
- **Flask (Python)**: Used for integrating with Mokuro and executing image processing tasks.
- **Mokuro**: A tool for text recognition on images, especially for working with vertical text.
- **RabbitMQ**: Manages tasks and message queues.
- **PostgreSQL**: The database for storing information about uploads and processing statuses.

## Architecture

1. **Monorepository**: Includes directories for each microservice and shared code.
2. **Microservices**: The Go backend handles file uploading and interaction with RabbitMQ, while the Flask subsystem performs processing with Mokuro.
3. **Asynchronous Processing**: RabbitMQ distributes tasks between services.
4. **Processing Status Tracking**: The system updates database records to show the current processing status of each file.

## Installation and Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/your-repo.git
   ```

2. **For the Mokuro server**:
   - Navigate to the `./mokuro-server` folder.
   - Install the requirements:
     ```bash
     pip install -r requirements.txt
     ```
   - Run the server:
     ```bash
     python app.py
     ```
   - Alternatively, you can start it using Docker:
     ```bash
     docker-compose up
     ```

3. **For the Reader API**:
   - Navigate to the `./reader-api` folder.
   - Start the Reader API using Docker:
     ```bash
     docker-compose up
     ```
   - Or run it directly:
     ```bash
     go run main.go
     ```

4. **Upload an image**:
   - Send a request to `http://localhost:3000/v1/upload` with form-data containing `[file, title]`.

5. **View current processed projects**:
   - Open [http://127.0.0.1:5001/projects](http://127.0.0.1:5001/projects) to see the current processed projects.
   - Open [http://127.0.0.1:5001/projects/[name]](http://127.0.0.1:5001/projects/[name]) to view the static results from Mokuro.

**Note**: The project is still under development.