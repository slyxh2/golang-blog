import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import './item.css';
const Item = (props) => {
    const { name, onEdit, onDelete } = props;
    return <div className="item-container">
        <p>{name}</p>
        <div className="flex-left edit-icon" onClick={onEdit}>
            <EditOutlined />
        </div>
        <div className='edit-icon delete-icon' onClick={onDelete}>
            <DeleteOutlined />
        </div>
    </div>
}


export default Item;