import { useState, useContext, useEffect } from "react";
import { Button, Modal, Input } from 'antd';
import { PlusOutlined } from '@ant-design/icons';

import { CategoryContext } from '../../context';
import CategoryItem from "../../components/CategoryItem";
import { createCategory } from "../../api";
import { useNavigate } from "react-router-dom";

const EditCategory = () => {
    const nagivate = useNavigate();
    const [allCategories, setAllCategories] = useState(null);
    const [loading, setLoading] = useState(true);
    const [addOpen, setAddOpen] = useState(false);
    const [name, setName] = useState('');
    const categoryContext = useContext(CategoryContext);
    const handleCancelAdd = () => {
        setAddOpen(false);
        setName('');
    }
    const handleAdd = () => {
        createCategory(name).then(() => {
            setAddOpen(false);
            nagivate(0);
        })
    }
    useEffect(() => {
        setAllCategories(categoryContext);
        setLoading(false);
    }, [categoryContext])

    if (loading) return <div>Loading...</div>
    return <>
        <div>
            <Button
                icon={<PlusOutlined />}
                style={{ 'marginBottom': '1em' }}
                onClick={() => setAddOpen(true)}
            >
                Add New Category
            </Button>

            {
                allCategories.map(item => <CategoryItem category={item.label} id={item.value} key={item.value} />)
            }
        </div>
        <Modal
            title="Add New Category"
            open={addOpen}
            onOk={handleAdd}
            onCancel={handleCancelAdd}
            okText="Add"
            cancelText="Cancel"
        >
            <Input
                placeholder="Enter Category Name"
                onChange={(e) => setName(e.target.value)}
            />
        </Modal>
    </>
}


export default EditCategory;