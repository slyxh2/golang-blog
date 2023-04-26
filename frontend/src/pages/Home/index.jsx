import './home.css'
import Sidebar from '../../components/Sidebar';
import { Outlet } from 'react-router-dom';
const Home = () => {
    return <div id="home-container">
        <aside id="side">
            <Sidebar />
        </aside>
        <main id="main">
            <Outlet />
        </main>
    </div>
};
export default Home;