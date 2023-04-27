import { useEffect, useState } from "react"
import env from "react-dotenv";

const useLogin = () => {
    const token = env.AUTH_TOKEN;
    const [isLog, setIsLog] = useState(false);
    useEffect(() => {
        if (sessionStorage.getItem(token)) {
            setIsLog(true)
        }
    }, [])
    return isLog;
}

export default useLogin;