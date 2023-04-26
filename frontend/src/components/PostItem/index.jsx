import { Modal } from 'antd';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

import Item from "../Item";
import { deletePost } from '../../api';
const PostItem = (props) => {
    const { header, id } = props;
    const nagivate = useNavigate();
    const [deleteOpen, setDeleteOpen] = useState(false);
    const handleDelete = (id) => {
        deletePost(id).then(() => {
            setDeleteOpen(false);
            nagivate(0);
        })
    }
    return <>
        <Item
            name={header}
            onDelete={() => setDeleteOpen(true)}
        />
        <Modal
            title={`Are you sure to delete ${header}`}
            open={deleteOpen}
            onOk={() => handleDelete(id)}
            onCancel={() => setDeleteOpen(false)}
            okText="Confirm"
            cancelText="Cancel"
        >
            <p>The Delete Can Not be Recovered</p>
        </Modal>
    </>
}

export default PostItem;