import { useState, useContext, useEffect } from "react";
import { CategoryContext } from '../../context';
import CategoryItem from "../../components/CategoryItem";
const EditCategory = () => {
    const [allCategories, setAllCategories] = useState(null);
    const [loading, setLoading] = useState(true);
    const categoryContext = useContext(CategoryContext);
    useEffect(() => {
        setAllCategories(categoryContext);
        setLoading(false);
    }, [categoryContext])

    if (loading) return <div>Loading...</div>
    return <div>
        {
            allCategories.map(item => <CategoryItem category={item.label} id={item.value} key={item.value} />)
        }
    </div>
}


export default EditCategory;