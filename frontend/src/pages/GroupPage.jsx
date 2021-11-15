import Template from "../components/template";
import ModalHeader from "../components/modalHeader.jsx";
import LeftPanel from "../components/leftPanel";
import Button from "../components/button";
import { useState, React, useEffect } from "react";
import AddMemberWindow from "../components/addMemberWindow";
import "../css/personalPage.css";
import UploadNotesWindow from "../components/uploadNotesWindow";
import { useParams } from "react-router";
import SectionTitle from "../components/sectionTitle.jsx";

function GroupPage(props) {
    const [buttonAddMember, setButtonAddMember] = useState(false);
    const [buttonGroupProfile, setButtonGroupProfile] = useState(false);
    const [members, setMembers] = useState([]);
    const [groups, setGroups] = useState([]);
    const [notes, setNotes] = useState([]);
    const [buttonUpload, setButtonUpload] = useState(false);
    const [group, setGroup] = useState("");
    const [groupDescription, setGroupDescription] = useState("");
    const { groupId } = useParams();

    const getGroupInfo = async () => {
        let success = true;
        const groupInfoResponse = await fetch(
            "http://localhost:3000/v1/group/" + groupId,
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const groupInfoResponseJSON = await groupInfoResponse.json();
            if (groupInfoResponseJSON.group_name) {
                setGroup(groupInfoResponseJSON.group_name);
                setGroupDescription(groupInfoResponseJSON.description);
            } else {
                console.error("Could not get group information.");
            }
        }
    };

    const getGroupMembers = async () => {
        let success = true;
        const membersResponse = await fetch(
            "http://localhost:3000/v1/group/" + groupId + "/members",
            {
                method: "GET",
                headers: {
                    Authorization:
                        "bearer " + window.localStorage.getItem("authToken"),
                },
            }
        ).catch((err) => {
            console.error(err);
            success = false;
        });

        if (success) {
            const memberResponseJSON = await membersResponse.json();
            setMembers([]);
            if (memberResponseJSON.members) {
                for (const member of memberResponseJSON.members) {
                    setMembers((arr) => [...arr, member]);
                }
            }
        }
    };

    useEffect(() => {
        getGroupInfo();
        getGroupMembers();
    }, []);

    return (
        <>
            <Template
                showSearch={true}
                showProfile={true}
                body={
                    <div className="d-flex flex-column left-side">
                        <ModalHeader title={`Hi Group ${group}`} />
                        <div className="d-flex flex-column align-items-left">
                            <LeftPanel
                                body={
                                    <div className="flex-row">
                                        <ModalHeader title="Members" />
                                        <div className="d-flex row agenda">
                                            {members.map((member) => (
                                                <h3>{member.username}</h3>
                                            ))}
                                        </div>
                                        <div className="flex-row">
                                            <Button
                                                title="ADD MEMBER"
                                                type="primary"
                                                clicked={() =>
                                                    setButtonAddMember(true)
                                                }
                                            />
                                            <AddMemberWindow
                                                trigger={buttonAddMember}
                                                setTrigger={setButtonAddMember}
                                            />
                                        </div>
                                        <div className="flex-row">
                                            <Button
                                                title="EDIT PROFILE"
                                                type="secondary"
                                                clicked={() =>
                                                    setButtonGroupProfile(true)
                                                }
                                            />
                                        </div>
                                    </div>
                                }
                            />
                        </div>
                        <div className="d-flex flex-column right-side-top">
                            <SectionTitle title="Our Group" />
                            <div className="agenda big">{groupDescription}</div>
                        </div>
                        <div className="d-flex right-side-middle">
                            <SectionTitle title="Our Notes" />
                            <div className="d-flex flex-row right-side-down">
                                <Button
                                    title="UPLOAD"
                                    type="primary"
                                    clicked={() => setButtonUpload(true)}
                                />
                                <UploadNotesWindow
                                    trigger={buttonUpload}
                                    setTrigger={setButtonUpload}
                                />
                            </div>
                        </div>
                    </div>
                }
            />
        </>
    );
}
export default GroupPage;