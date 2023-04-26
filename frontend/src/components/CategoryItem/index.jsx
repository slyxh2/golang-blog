import { useState } from 'react';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { Modal } from 'antd';
import './categoryItem.css';
import { deleteCategory } from '../../api';
import { useNavigate } from 'react-router-dom';
const CategoryItem = (props) => {
    const { category, id } = props;
    const [open, setOpen] = useState(false);
    const nagivate = useNavigate();
    const showModal = () => {
        setOpen(true);
    };
    const hideModal = () => {
        setOpen(false);
    };
    const handleDelete = (id) => {
        console.log(id);
        deleteCategory(id).then(() => {
            setOpen(false);
            nagivate(0);
        });
    };
    return <>
        <div className="category-item-container">
            <p>{category}</p>
            <div className="flex-left edit-icon">
                <EditOutlined />
            </div>
            <div className='edit-icon delete-icon' onClick={showModal}>
                <DeleteOutlined />
            </div>
        </div>
        <Modal
            title={`Are you sure to delete ${category}`}
            open={open}
            onOk={() => handleDelete(id)}
            onCancel={hideModal}
            okText="Confirm"
            cancelText="Cancel"
        >
            <p>The Delete Can Not be Recovered</p>
        </Modal>
    </>

}

export default CategoryItem;