# jp-reader

jp-reader is a monorepository that includes backends in Go and Flask for processing images using Mokuro. It also provides an interface for uploading manga, viewing processing statuses, and accessing a reader.

The project is designed to aid in learning the Japanese language, facilitate manga reading, and reduce the need to work with the terminal for image processing through Mokuro. Instead, users can interactively upload manga and have free access to reading on their local devices.

Recommend - using Yomichan/Yomitan for more effective reading.

## Key Features

1. **Manga Uploading**: The application allows users to upload manga images, which are processed using the Mokuro library for efficient text recognition.
2. **Image Processing**: Mokuro is used for text recognition and image conversion to facilitate further content handling.
3. **Message Queues**: RabbitMQ provides asynchronous data processing, organizing task queues for managing file uploads and processing.
4. **Processing Status Tracking**: Users can view the status of their uploaded manga, providing insights into the processing progress and enabling a seamless reading experience.

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
4. **Processing Status Tracking**: The system updates database records to show the current processing status of each uploaded manga.
5. **User Interface**: The React-based frontend offers an intuitive interface for users to upload manga images easily, view processing statuses of uploaded files, access a reader for seamless manga viewing, Interact with features that enhance the reading experience.

## Installation and Setup

### For local use

If you want to use the jp-reader service locally without diving into development, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
    ```

2. **Run the entire application with Docker**:

    In the root directory, run:
    ```bash
    docker-compose up
    ```
    This will start all necessary services, including the Mokuro server, Reader API, PostgreSQL, and RabbitMQ.

3. **Run reader app**

    Navigate to the reader-app folder
    ```bash
    cd ./reader-app
    ```

    Install dependencies
    
     ```bash
        pnpm install
     ```
    Build

    ```bash
        pnpm build
    ``` 
    For Access the Reader App
    ```bash
        pnpm preview
    ``` 
    Navigate to http://localhost:4173/ in your web browser to use the application.
    
    From there, you can upload images, view the processing results and use reader.  


### For Developers: Running and Modifying Each Service
For detailed instructions on running and modifying each service, please refer to the respective README files in each service directory:

Mokuro Server: [mokuro-server/README.md](https://github.com/sieugene/jp-reader/blob/master/mokuro-server/Readme.md)

Reader API: [reader-api/README.md](https://github.com/sieugene/jp-reader/blob/master/reader-api/readme.md)

Reader App: [reader-app/README.md](https://github.com/sieugene/jp-reader/blob/master/reader-app/README.md)

Note: The project is still under development. Please report any issues you encounter.

### Contributing
If you would like to contribute to the project, please follow these steps:

Fork the repository.

Create a new branch for your feature or bugfix.

Make your changes and commit them.

Push your branch and open a pull request.