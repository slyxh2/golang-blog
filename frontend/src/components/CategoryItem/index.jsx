import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import './categoryItem.css';

const CategoryItem = (props) => {
    const { category, id } = props;
    return <div className="category-item-container">
        <p>{category}</p>
        <div className="flex-left edit-icon">
            <EditOutlined />
        </div>
        <div className='edit-icon delete-icon'>
            <DeleteOutlined />
        </div>

    </div>
}

export default CategoryItem;