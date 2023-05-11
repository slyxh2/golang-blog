import { useEffect, useState } from "react"


const useLogin = () => {
    const token = 'Authorization';
    const [isLog, setIsLog] = useState(false);
    useEffect(() => {
        if (sessionStorage.getItem(token)) {
            setIsLog(true)
        }
    }, [])
    return isLog;
}

export default useLogin;