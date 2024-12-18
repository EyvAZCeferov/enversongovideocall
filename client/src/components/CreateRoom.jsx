import React from "react";

const CreateRoom = (props) => {
    const create = async (e) => {
        e.preventDefault();

        const resp = await fetch("http://31.171.73.195:3290/create");
        const { room_id } = await resp.json();

		props.history.push(`/room/${room_id}`)
    };

    return (
        <div>
            <button onClick={create}>Create Room</button>
        </div>
    );
};

export default CreateRoom;
