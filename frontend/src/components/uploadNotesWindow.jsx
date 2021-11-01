import Button from "./button";
import InputBox from "./inputBox";
import ModalHeader from "./modalHeader.jsx";
import { useHistory } from "react-router-dom";
import { useState } from "react";
import "../css/uploadNotesWindow.css";
import RadioButton from "./radioButton";
import ModalWindow from "./modalWindow";
import UploadDropzone from "./uploadDropzone";

function UploadNotesWindow({ trigger, setTrigger }) {
    const history = useHistory();
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [tag, setTag] = useState("");
    const [type, setType] = useState("");
    const [file, setFile] = useState(null);

    const setPublic = () => {
        setType("Public");
    };

    const setGroup = () => {
        setType("Group");
    };

    const setPrivate = () => {
        setType("Private");
    };

    const attemptUpload = async () => {
        let formData = new FormData();
        formData.append("noteFile", file);

        const pdfResponse = await fetch("http://localhost:3000/v1/notes/file", {
            method: "POST",
            headers: {
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: formData,
        });

        const pdfResponseJSON = await pdfResponse.json();
        if (!pdfResponseJSON.note_reference) {
            alert("Something went wrong with PDF upload!");
            return;
        }

        const note_reference = pdfResponseJSON.note_reference;

        const style = "dummy";
        const response = await fetch("http://localhost:3000/v1/notes/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization:
                    "bearer " + window.localStorage.getItem("authToken"),
            },
            body: JSON.stringify({
                title,
                description,
                tag,
                type,
                note_reference,
                style
            }),
        });

        const responseJSON = await response.json();
        if (responseJSON.note_id) {
            history.push("/note/" + responseJSON.note_id);
        } else {
            alert("Something went wrong with note upload!");
            setTrigger(false);
        }
    };

    return trigger ? (
        <div className="popup">
            <ModalWindow
                body={
                    <div className="d-flex flex-column align-items-center">
                        <ModalHeader title="Upload Notes" />
                        <div className="d-flex flex-column align-items-end">
                            <InputBox
                                label="Title"
                                placeholder="title"
                                onChange={setTitle}
                            />
                            <InputBox
                                textArea
                                label="Description"
                                placeholder="description"
                                onChange={setDescription}
                            />
                            <InputBox
                                label="Tag"
                                placeholder="tag"
                                onChange={setTag}
                            />
                        </div>
                        <br></br>
                        <label className="agenda">Sharing</label>
                        <div className="d-flex flex-row ">
                            <RadioButton
                                group="sharing"
                                label="Public"
                                onChange={setPublic}
                            />
                            <RadioButton
                                group="sharing"
                                label="Group"
                                onChange={setGroup}
                            />
                            <RadioButton
                                group="sharing"
                                label="Private"
                                onChange={setPrivate}
                            />
                        </div>
                        <br></br>
                        <label className="agenda">Attach File</label>
                        <UploadDropzone setFile={setFile} />
                        <div className="d-flex flex-row ">
                            <Button
                                title="CANCEL"
                                type="secondary"
                                clicked={() => setTrigger(false)}
                            />
                            <Button
                                title="UPLOAD"
                                type="primary"
                                clicked={attemptUpload}
                            />
                        </div>
                    </div>
                }
            />
        </div>
    ) : (
        ""
    );
}

export default UploadNotesWindow;
