import os
import uuid
import subprocess
import shutil
from flask import Flask, request, jsonify

app = Flask(__name__)

BASE_UPLOAD_FOLDER = 'projects/uploads'

@app.route('/upload/<string:title>', methods=['POST'])
def upload_file(title):

    project_folder = os.path.join(BASE_UPLOAD_FOLDER, title)
    upload_folder = os.path.join(project_folder, "images")

    if os.path.exists(project_folder):
        shutil.rmtree(project_folder)

    os.makedirs(upload_folder, exist_ok=True)

    if 'file' not in request.files:
        return jsonify({"error": "No file part"}), 400

    files = request.files.getlist('file')
    if not files:
        return jsonify({"error": "No selected files"}), 400

    file_paths = []
    for file in files:
        if file.filename == '':
            return jsonify({"error": "One or more files have no selected filename"}), 400

        file_extension = os.path.splitext(file.filename)[1]
        new_filename = f"{uuid.uuid4()}{file_extension}"
        file_path = os.path.join(upload_folder, new_filename)
        file.save(file_path)
        file_paths.append(file_path)

    subprocess.Popen(['mokuro', upload_folder, '--force_cpu', '--disable_confirmation'])

    return jsonify({"message": f"{len(file_paths)} files uploaded and Mokuro started for '{title}'."}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
