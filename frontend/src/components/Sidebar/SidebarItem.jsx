import './sidebarItem.css';
const SidebarItem = (props) => {
    const { children } = props;
    return <div id="sidebar-item">
        {children}
    </div>
}

export default SidebarItem;