import { useEffect, useState } from "react"

const useLogin = () => {
    const token = process.env.LOGIN_TOKEN;
    const [isLog, setIsLog] = useState(false);
    useEffect(() => {
        sessionStorage.getItem(token) ? setIsLog(true) : setIsLog(false);
    }, [])
    return isLog;
}

export default useLogin;