import os
import json
import uuid
import subprocess
import shutil
from flask import Flask, abort, request, jsonify, send_from_directory
from flask_cors import CORS


app = Flask(__name__)
CORS(app)

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
    for index, file in enumerate(files, start=1):
        if file.filename == '':
            return jsonify({"error": "One or more files have no selected filename"}), 400

        file_extension = os.path.splitext(file.filename)[1]
        new_filename = f"{index}_{uuid.uuid4()}{file_extension}"
        file_path = os.path.join(upload_folder, new_filename)
        file.save(file_path)
        file_paths.append(file_path)

    process = subprocess.Popen(['mokuro', upload_folder, '--force_cpu', '--disable_confirmation'])

    process.wait()

    return jsonify({"message": f"{len(file_paths)} files uploaded and created project '{title}'."}), 200


@app.route('/projects', methods=['GET'])
def get_projects():
    projects = get_projects_data()
    return jsonify({'projects': projects}), 200

@app.route('/projects/<string:project_name>', methods=['GET'])
def get_project_content(project_name):
    project_data = get_projects_data(project_name)
    if not project_data:
        return jsonify({'error': 'Project not found'}), 404

    return jsonify(project_data), 200

def get_projects_data(project_name=None):
    projects = []
    base_path = BASE_UPLOAD_FOLDER

    if project_name is None:
        for name in os.listdir(base_path):
            project_data = fetch_project_data(name)
            if project_data:
                projects.append(project_data)
        return projects

    project_data = fetch_project_data(project_name)
    return project_data if project_data else None


def fetch_project_data(project_name):
    project_path = os.path.join(BASE_UPLOAD_FOLDER, project_name)
    if not os.path.isdir(project_path):
        return None

    html_path = os.path.join(project_path, 'images.html')
    if not os.path.exists(html_path):
        return None

    images_folder = os.path.join(project_path, 'images')
    ocr_folder = os.path.join(project_path, '_ocr', 'images')

    images = [f for f in os.listdir(images_folder) if os.path.isfile(os.path.join(images_folder, f))]
    image_links = [f'/projects/{project_name}/images/{image}' for image in images]

    ocr_data = []
    for filename in os.listdir(ocr_folder):
        if filename.endswith('.json'):
            file_path = os.path.join(ocr_folder, filename)
            with open(file_path, 'r', encoding='utf-8') as json_file:
                json_content = json.load(json_file)
                ocr_data.append({
                    'data': json_content,
                    'name': filename
                })

    return {
        'name': project_name,
        'images': image_links,
        'ocrData': ocr_data
    }


@app.route('/projects/<string:project_name>/images/<path:filename>', methods=['GET'])
def serve_image(project_name, filename):
    images_folder = os.path.join(BASE_UPLOAD_FOLDER, project_name, 'images')

    if not os.path.exists(os.path.join(images_folder, filename)):
        abort(404)
    return send_from_directory(images_folder, filename)

@app.route('/projects/<string:project_name>', methods=['DELETE'])
def delete_project(project_name):
    project_folder = os.path.join(BASE_UPLOAD_FOLDER, project_name)

    if not os.path.exists(project_folder):
        return jsonify({"error": "Project not found"}), 404

    shutil.rmtree(project_folder)

    return jsonify({"message": f"Project '{project_name}' deleted successfully."}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
