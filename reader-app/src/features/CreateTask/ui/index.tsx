import { CONFIG } from "@/shared/config";
import classNames from "classnames";
import { ChangeEvent, FormEvent, useRef, useState } from "react";

type ErrorState = {
  title: boolean;
  files: boolean;
};

export const CreateTask = () => {
  const modal = useRef<HTMLDialogElement>(null);
  const [files, setFiles] = useState<FileList | null>(null);
  const [title, setTitle] = useState<string>("");
  const [errors, setErrors] = useState<ErrorState>({
    title: false,
    files: false,
  });

  const open = () => {
    modal.current?.showModal();
  };

  const clearErrors = () => {
    setErrors({ files: false, title: false });
  };
  const validate = () => {
    const errorsState: ErrorState = {
      files: !files?.length,
      title: title.length < 3,
    };
    setErrors(errorsState);
    return errorsState.files || errorsState.title;
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const hasErrors = validate();
    if (hasErrors) {
      return;
    }

    const formData = new FormData();
    formData.append("title", title);

    Array.from(files!).forEach((file) => {
      formData.append("file", file);
    });

    await fetch(`${CONFIG.READER_API}/upload`, {
      method: "POST",
      body: formData,
    });

    setFiles(null);
    setTitle("");
  };

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (event.target.files) {
      setFiles(event.target.files);
    }
  };
  const handleTitleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  return (
    <div>
      <button className="btn btn-primary" onClick={open}>
        Create new
      </button>

      <dialog id="my_modal_1" className="modal" ref={modal}>
        <div className="modal-box">
          <h3 className="font-bold text-lg">Create new task</h3>
          <p className="py-4">
            Lorem ipsum dolor sit amet consectetur adipisicing elit.
            Consectetur, nihil. Voluptates placeat sed nihil ab eaque laborum
            distinctio sit quis.
          </p>
          <form
            onSubmit={handleSubmit}
            onChange={() => {
              clearErrors();
            }}
          >
            <div className="flex gap-5 flex-col">
              <label className="form-control w-full max-w-xs">
                <div className="label">
                  <span className="label-text">Name</span>
                </div>
                <input
                  type="text"
                  className={classNames(
                    "input input-bordered w-full max-w-xs",
                    {
                      "input-error": errors.title,
                    },
                  )}
                  placeholder="My project"
                  value={title}
                  onChange={handleTitleChange}
                  required
                />
              </label>

              <input
                type="file"
                className={classNames("file-input w-full max-w-xs", {
                  "file-input-error": errors.files,
                })}
                multiple
                accept="image/*"
                onChange={handleFileChange}
                required
              />
              <button className="btn btn-neutral" type="submit">
                Upload
              </button>
            </div>
          </form>

          <div className="modal-action">
            <form method="dialog">
              <button className="btn">Close</button>
            </form>
          </div>
        </div>
      </dialog>
    </div>
  );
};
