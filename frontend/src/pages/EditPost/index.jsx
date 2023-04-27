import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getPost } from "../../api";
import PostPlayground from "../../components/PostPlayground";

const EditPost = () => {
    const { postId } = useParams();
    const [loading, setLoading] = useState(true);
    const [postInf, setPostInf] = useState({});

    useEffect(() => {
        getPost(postId).then((res) => {
            console.log(res.data);
            const { header, category, content, id } = res.data.post;
            setPostInf({
                header,
                category,
                content,
                id
            })
            setLoading(false);
        })
    }, [postId])
    if (loading) return <div>Loading...</div>
    return <PostPlayground
        header={postInf.header}
        category={postInf.category}
        content={postInf.content}
        id={postInf.id}
    />
}

export default EditPost;