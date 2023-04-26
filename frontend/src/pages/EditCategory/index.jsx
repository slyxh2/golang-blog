import { useState } from "react";
import CategoryItem from "../../components/CategoryItem";
const EditCategory = () => {
    const [loading, setLoading] = useState(true);


    if (loading) return <div>Loading...</div>
    return <div>
        <CategoryItem category={"typescript"} />
        <CategoryItem category={"typescript"} />
    </div>
}


export default EditCategory;