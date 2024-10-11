import os
import re
import uuid
import subprocess
import shutil
from flask import Flask, abort, request, jsonify, send_from_directory

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

    process = subprocess.Popen(['mokuro', upload_folder, '--force_cpu', '--disable_confirmation'])

    process.wait()

    return jsonify({"message": f"{len(file_paths)} files uploaded and created project '{title}'."}), 200


@app.route('/projects', methods=['GET'])
def get_projects():
    projects = []

    for name in os.listdir(BASE_UPLOAD_FOLDER):
        project_path = os.path.join(BASE_UPLOAD_FOLDER, name)
        if os.path.isdir(project_path):

            html_path = os.path.join(project_path, 'images.html')
            if os.path.exists(html_path):

                projects.append({
                    'name': name,
                    'link': f'/projects/{name}'
                })

    return jsonify({'projects': projects})


def fix_image_paths(html_content, project_name):
    pattern = r'background-image:url\(&quot;images/([^)]+)&quot;\)'
    fixed_html = re.sub(pattern, f'background-image:url(&quot;/projects/{project_name}/images/\\1&quot;)', html_content)
    return fixed_html

# Serve html
@app.route('/projects/<string:project_name>', methods=['GET'])
def serve_html(project_name):
    project_folder = os.path.join(BASE_UPLOAD_FOLDER, project_name)
    html_path = os.path.join(project_folder, 'images.html')

    if not os.path.exists(html_path):
        abort(404)

    with open(html_path, 'r', encoding='utf-8') as f:
        html_content = f.read()

    fixed_html_content = fix_image_paths(html_content, project_name)

    return fixed_html_content

# Serve html
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
