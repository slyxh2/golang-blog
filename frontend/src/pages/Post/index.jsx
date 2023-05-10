import { useEffect, useState } from 'react';
import { Pagination } from 'antd';
import { useParams } from 'react-router-dom';

import { getPosts } from '../../api';
import PostItem from '../../components/PostItem';
import { ITEM_PER_PAGE } from '../../const';

const Post = () => {
    let { categoryId } = useParams();
    const [page, setPage] = useState(1);
    const [total, setTotal] = useState(0);
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);
    const handlePageChange = (page) => {
        setPage(page);
    }
    useEffect(() => {
        getPosts(page, ITEM_PER_PAGE, categoryId).then((res) => {
            console.log(res);
            const { totalPage, posts } = res.data;
            setTotal(totalPage * ITEM_PER_PAGE);
            if (posts) {
                setPosts(posts);
            } else {
                setPosts([]);
            }
            setLoading(false);
        })
    }, [categoryId, page])

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
        <Pagination
            defaultCurrent={1}
            total={total}
            onChange={handlePageChange}
            style={{ transform: 'translateX(40%)', marginTop: '3em' }}
        />


    </>
}

export default Post;