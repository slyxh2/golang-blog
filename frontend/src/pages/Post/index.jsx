import { useEffect } from 'react';
import { useParams } from 'react-router-dom';

const Post = () => {
    let { categoryId } = useParams();

    useEffect(() => {
        console.log(categoryId);
    }, [categoryId])
    return <>
        POST
    </>
}

export default Post;