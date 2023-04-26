import './sidebarItem.css';
const SidebarItem = (props) => {
    const { children, onClick } = props;
    return <div id="sidebar-item" onClick={onClick}>
        {children}
    </div>
}

export default SidebarItem;