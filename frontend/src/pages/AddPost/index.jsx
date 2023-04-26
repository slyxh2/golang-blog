import React, { useRef } from 'react';
import { useMemo, useState } from 'react';
import { Select, Input, Button, Alert } from 'antd';
import { UploadOutlined, CloudUploadOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

import Markdown from '../../components/Markdown';
import { debounce } from '../../utils';
import { getAllCategory, uploadPost } from '../../api';
import './addPost.css';
const AddPost = () => {
    const textRef = useRef();
    const uploadRef = useRef();
    const nagivate = useNavigate();
    const [input, setInput] = useState(""); // markdown input
    const [allCategories, setAllCategories] = useState([]);
    const [selectedCategory, setSelectedCategory] = useState("");
    const [postHeader, setPostHeader] = useState("");
    const handleInput = debounce((e) => {
        setInput(e.target.value);
    }, 500);
    useMemo(async () => {
        let res = await getAllCategory();
        let all = res.data.categories.map(category => {
            let obj = {};
            obj.value = category.id;
            obj.label = category.name;
            return obj;
        })
        setAllCategories(all);
    }, []);
    const handleCategotySelect = (val) => {
        setSelectedCategory(val);
    }
    const handelPostHeader = debounce((e) => {
        setPostHeader(e.target.value);
    }, 1000);
    const handleFileSelect = (event) => {
        const file = event.target.files[0];
        const reader = new FileReader();
        reader.onload = () => {
            const fileContent = reader.result;
            console.log(fileContent);
            textRef.current.value = fileContent;
            setInput(fileContent);
        };
        reader.readAsText(file);
    };
    const clickUpload = () => {
        uploadRef.current.click();
    }
    const handlePost = () => {
        const file = new Blob([input], { type: "text/markdown" });

        uploadPost(file, postHeader, selectedCategory).then(data => {
            nagivate('/');
        }).catch(err => console.log(err));
    }
    return <>
        <div id="tools-container">
            <Select
                placeholder="Category"
                style={{
                    width: 120,
                }}
                onChange={handleCategotySelect}
                options={allCategories}
            />
            <Input placeholder="Input Post Header" onChange={handelPostHeader} style={{
                width: 240,
            }} />

            <div id="markdown-uploader">
                <Button icon={<UploadOutlined />} onClick={clickUpload}>Click to Upload</Button>
                <input
                    id="file-input"
                    type="file"
                    accept=".md"
                    onChange={handleFileSelect}
                    ref={uploadRef}
                    hidden
                />

            </div>
            <Button icon={<CloudUploadOutlined />} onClick={handlePost}>POST</Button>
        </div>
        <div id="add-post-container">
            <textarea id="text-area" onInput={(e) => handleInput(e)} ref={textRef}></textarea>
            <div id="markdown-container">
                <Markdown post={input} />
            </div>
        </div>
    </>

}
export default React.memo(AddPost);