import React, { useEffect, useRef } from 'react';
import { useState, useContext } from 'react';
import { Select, Input, Button } from 'antd';
import MdEditor from 'react-markdown-editor-lite';

import { UploadOutlined, CloudUploadOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

import Markdown from '../../components/Markdown';
import { debounce } from '../../utils';
import { uploadPost, editPost } from '../../api';
import { CategoryContext } from '../../context';

import './postPlayground.css';
import 'react-markdown-editor-lite/lib/index.css';

const PostPlayground = (props) => {
    const { header: defaultHeader, category: defaultCategory, content: defaultContent, id: defaultId } = props;
    // const textRef = useRef();
    const uploadRef = useRef();
    const editor = useRef();

    const nagivate = useNavigate();
    const categoryContext = useContext(CategoryContext);
    const [input, setInput] = useState(defaultContent || ""); // markdown input
    const [allCategories, setAllCategories] = useState([]);
    const [selectedCategory, setSelectedCategory] = useState(defaultCategory || "");
    const [postHeader, setPostHeader] = useState(defaultHeader || "");
    const [isContentEdit, setIsContentEdit] = useState(false);
    const handleInput = debounce(({ text }) => {
        // console.log(text);
        setInput(text);
        if (isContentEdit === false) setIsContentEdit(true);
    }, 500);
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
            editor.current.setText(fileContent);
        };
        reader.readAsText(file);
    };
    const clickUpload = () => {
        uploadRef.current.click();
    };
    const handlePost = () => {
        const file = new Blob([input], { type: "text/markdown" });
        uploadPost(file, postHeader, selectedCategory).then(data => {
            nagivate('/');
        }).catch(err => console.log(err));
    }
    const handleEdit = () => {
        const params = {};
        params.id = defaultId;
        if (defaultHeader !== postHeader) params.header = postHeader;
        if (defaultCategory !== selectedCategory) params.category = selectedCategory;
        if (isContentEdit) {
            const file = new Blob([input], { type: "text/markdown" });
            params.file = file;
        }
        editPost(params).then(data => {
            nagivate('/');
        }).catch(err => console.log(err));
    }
    useEffect(() => {
        setAllCategories(categoryContext);
    }, [categoryContext])
    useEffect(() => {
        if (defaultContent) editor.current.setText(defaultContent);
    }, [defaultContent])
    return <>
        <div id="tools-container">
            <Select
                placeholder="Category"
                style={{
                    width: 120,
                }}
                onChange={handleCategotySelect}
                options={allCategories}
                defaultValue={defaultCategory}
            />
            <Input
                placeholder="Input Post Header"
                onChange={handelPostHeader}
                style={{
                    width: 240,
                }}
                defaultValue={defaultHeader || ""}
            />

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
            <Button icon={<CloudUploadOutlined />} onClick={defaultHeader ? handleEdit : handlePost}>
                {defaultHeader ? "EDIT" : "POST"}
            </Button>
        </div>

        {/* <textarea id="text-area" onInput={(e) => handleInput(e)} ref={textRef}></textarea>
            <div id="markdown-container">
                <Markdown post={input} />
            </div> */}
        <MdEditor
            style={{ height: '700px' }}
            renderHTML={text => <Markdown post={text} />}
            onChange={handleInput}
            ref={editor}
        />

    </>

}
export default React.memo(PostPlayground);