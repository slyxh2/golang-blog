import React from 'react';
import { useState } from 'react';

import { Modal, Input } from 'antd';
import { useNavigate } from 'react-router-dom';

import { deleteCategory, editCategoty } from '../../api';
import Item from '../Item';

const CategoryItem = (props) => {
    const { category, id } = props;
    const [deleteOpen, setDeleteOpen] = useState(false);
    const [editOpen, setEditOpen] = useState(false);
    const [name, setName] = useState(category);
    const nagivate = useNavigate();
    const handleDelete = (id) => {
        deleteCategory(id).then(() => {
            setDeleteOpen(false);
            nagivate(0);
        });
    };
    const handleEdit = (id) => {
        editCategoty(id, name).then(() => {
            setEditOpen(false);
            nagivate(0);
        })
    }
    const handleCancelEdit = () => {
        setName(category);
        setEditOpen(false)
    }
    return <>
        <Item
            name={category}
            onEdit={() => setEditOpen(true)}
            onDelete={() => setDeleteOpen(true)}
        />
        <Modal
            title={`Are you sure to delete ${category}`}
            open={deleteOpen}
            onOk={() => handleDelete(id)}
            onCancel={() => setDeleteOpen(false)}
            okText="Confirm"
            cancelText="Cancel"
        >
            <p>The Delete Can Not be Recovered</p>
        </Modal>
        <Modal
            title={`Edit ${category}`}
            open={editOpen}
            onOk={() => handleEdit(id)}
            onCancel={handleCancelEdit}
            okText="Edit"
            cancelText="Cancel"
        >
            <Input
                placeholder="Enter Category Name"
                defaultValue={category}
                onChange={(e) => setName(e.target.value)}
            />
        </Modal>
    </>

}

export default React.memo(CategoryItem);