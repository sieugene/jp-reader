import os
import uuid
import subprocess
from flask import Flask, request, jsonify

app = Flask(__name__)

UPLOAD_FOLDER = 'projects/uploads'
os.makedirs(UPLOAD_FOLDER, exist_ok=True)

@app.route('/upload', methods=['POST'])
def upload_file():
    if 'file' not in request.files:
        return jsonify({"error": "No file part"}), 400

    file = request.files['file']
    if file.filename == '':
        return jsonify({"error": "No selected file"}), 400

    file_extension = os.path.splitext(file.filename)[1]
    new_filename = f"{uuid.uuid4()}{file_extension}"
    file_path = os.path.join(UPLOAD_FOLDER, new_filename)
    file.save(file_path)

    manga_volume_path = os.path.dirname(file_path)
    subprocess.Popen(['mokuro', manga_volume_path, '--force_cpu' ,'--disable_confirmation'])

    return jsonify({"message": "File uploaded and Mokuro started."}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
