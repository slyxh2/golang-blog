import { Button, Checkbox, Form, Input } from 'antd';
import { useNavigate } from "react-router-dom";
import { handleLogin } from '../../api';
import './login.css'
const Login = () => {
    const navigate = useNavigate();
    const onFinish = (values) => {
        const { name, password } = values;
        handleLogin(name, password).then(res => {
            console.log(res);
            sessionStorage.setItem(process.env.AUTH_TOKEN, res.data.token);
            navigate(0);
        })
    };
    return <div id="login-form">
        <h1>Blog Management Platform</h1>
        <Form
            name="basic"
            labelCol={{
                span: 8,
            }}
            wrapperCol={{
                span: 16,
            }}
            style={{
                maxWidth: 600,
            }}
            initialValues={{
                remember: true,
            }}
            onFinish={onFinish}
            autoComplete="off"
        >
            <Form.Item
                label="Username"
                name="name"
                rules={[
                    {
                        required: true,
                        message: 'Please input your username!',
                    },
                ]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                label="Password"
                name="password"
                rules={[
                    {
                        required: true,
                        message: 'Please input your password!',
                    },
                ]}
            >
                <Input.Password />
            </Form.Item>

            <Form.Item
                wrapperCol={{
                    offset: 8,
                    span: 16,
                }}
            >
                <Button type="primary" htmlType="submit">
                    Login
                </Button>
            </Form.Item>
        </Form>
    </div>

}

export default Login;