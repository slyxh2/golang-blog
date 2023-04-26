import './home.css'
import Sidebar from '../../components/Sidebar';
import { Outlet } from 'react-router-dom';
import { useState, useMemo } from 'react';
import { getAllCategory } from '../../api';
import { CategoryContext } from '../../context'

const Home = () => {
    const [category, setCategory] = useState([]);
    useMemo(async () => {
        let res = await getAllCategory();
        let all = res.data.categories.map(category => {
            let obj = {};
            obj.value = category.id;
            obj.label = category.name;
            return obj;
        })
        setCategory(all);
    }, []);
    return <CategoryContext.Provider value={category}>
        <div id="background-container"></div>
        <div id="home-container">
            <aside id="side">
                <Sidebar />
            </aside>
            <main id="main">
                <Outlet />
            </main>
        </div>
    </CategoryContext.Provider>



};
export default Home;