import { useContext, useEffect, useRef, useState } from 'react';
import { DownOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom';

import { CategoryContext } from '../../context';
import SidebarItem from './SidebarItem';
import './sidebar.css'

const Sidebar = () => {
    const dropdownList = useRef();
    const dropdownIcon = useRef();
    const [category, setCategory] = useState(null);
    const categoryContext = useContext(CategoryContext);
    const [isLoading, setIsLoading] = useState(true);
    const nagivate = useNavigate();
    const handleCollapse = () => {
        let isCollapse = true;
        return () => {
            isCollapse = !isCollapse
            if (isCollapse) {
                // dropdownIcon.current
                dropdownIcon.current.classList.remove('collapse-icon');
                dropdownList.current.classList.remove('collapse-list');

            } else {
                dropdownIcon.current.classList.add('collapse-icon');
                dropdownList.current.classList.add('collapse-list');
            }
        }
    }
    const handleClickBarItem = (id) => {
        nagivate("/" + id);
    }
    useEffect(() => {
        setCategory(categoryContext);
        setIsLoading(false);
    }, [categoryContext])
    // useEffect(() => {
    //     getAllCategory().then(res => {
    //         setCategory(res.data.categories);
    //         setIsLoading(false);
    //     })
    // }, [])
    if (isLoading) return <div>LOADING...</div>

    return <div id="sidebar">
        <SidebarItem
            onClick={() => nagivate("/add-post")}
        >
            New Post
        </SidebarItem>
        <SidebarItem
            onClick={() => nagivate("/edit-category")}
        >Edit Category</SidebarItem>
        <SidebarItem
            onClick={() => nagivate("/")}
        >
            All Post
        </SidebarItem>

        <SidebarItem>
            <div className='dropdown-item' onClick={handleCollapse()}>
                <div>Category</div>
                <DownOutlined className='dropdown-icon' ref={dropdownIcon} />
            </div>
        </SidebarItem>
        <ul className='drop-list' ref={dropdownList}>
            {
                category.map(item => <li
                    key={item.value}
                    onClick={() => handleClickBarItem(item.value)}>
                    <SidebarItem>
                        {item.label}
                    </SidebarItem>
                </li>)
            }
        </ul>

    </div >
}


export default Sidebar;