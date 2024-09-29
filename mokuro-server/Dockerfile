FROM python:3.9-slim

RUN apt-get update && apt-get install -y \
    libgl1-mesa-glx \
    libglib2.0-0 \
    git \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

RUN pip install --upgrade pip setuptools wheel

RUN pip install flask
RUN pip install mokuro

WORKDIR /app

COPY app.py /app/

RUN mkdir -p /app/projects/uploads

EXPOSE 5001

CMD ["python", "app.py"]
