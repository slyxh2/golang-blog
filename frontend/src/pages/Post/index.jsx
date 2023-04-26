import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';


import { getPosts } from '../../api';
import PostItem from '../../components/PostItem';
import { ITEM_PER_PAGE } from '../../const';
const Post = () => {
    let { categoryId } = useParams();
    const [page, setPage] = useState(1);
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);
    useEffect(() => {
        getPosts(page, ITEM_PER_PAGE, categoryId).then((res) => {
            console.log(res);

            if (res.data.posts) {
                setPosts(res.data.posts);
            } else {
                setPosts([]);
            }
            setLoading(false);
        })
    }, [categoryId])

    if (loading) return <div>Loading...</div>
    return <>
        {
            posts.length ?
                posts.map(post => <PostItem
                    header={post.header}
                    key={post.id}
                    id={post.id}
                />)
                :
                <div>No Post</div>
        }



    </>
}

export default Post;