import { useState } from 'react';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { Modal, Input } from 'antd';
import './categoryItem.css';
import { deleteCategory, editCategoty } from '../../api';
import { useNavigate } from 'react-router-dom';
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
        <div className="category-item-container">
            <p>{category}</p>
            <div className="flex-left edit-icon" onClick={() => setEditOpen(true)}>
                <EditOutlined />
            </div>
            <div className='edit-icon delete-icon' onClick={() => setDeleteOpen(true)}>
                <DeleteOutlined />
            </div>
        </div>
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

export default CategoryItem;