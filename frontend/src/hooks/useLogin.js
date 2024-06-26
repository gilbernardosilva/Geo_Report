import { useState, useEffect } from "react";
import { useTranslation } from 'react-i18next';
import { useNavigate } from "react-router-dom";
import { useAuth } from '../hooks/AuthContext';
import { jwtDecode } from "jwt-decode";




function useLogin() {
    const { t } = useTranslation();

    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [action, setAction] = useState("Sign Up");
    const [showWrongPasswordModal, setShowWrongPasswordModal] = useState(false);

    const [isValidEmail, setIsValidEmail] = useState(true);

    const navigate = useNavigate();
    const { login, setLoggedIn, isLoggedIn, userInfo } = useAuth();


    const handleEmailChange = (e) => {
        const newEmail = e.target.value;
        setEmail(newEmail);

        const isValid = /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(newEmail);
        setIsValidEmail(isValid);
    };

    const [isValidPassword, setIsValidPassword] = useState(true);
    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;

    const handlePasswordChange = (e) => {
        const newPassword = e.target.value;
        setPassword(newPassword);

        const isValid = passwordRegex.test(newPassword);
        setIsValidPassword(isValid);
    };

    const [showPasswordModal, setShowPasswordModal] = useState(false);

    const handleShowPasswordModal = () => setShowPasswordModal(true);
    const handleClosePasswordModal = () => setShowPasswordModal(false);


    const handleSubmit = async (e) => {
        e.preventDefault();
        const apiEndpoint = action === t('login') ? "login" : "register";
        const requestBody = action === t('login')
            ? { username, password }
            : { email, password, username, firstName, lastName };

        try {
            const response = await fetch(`${process.env.REACT_APP_API_URL}/user/${apiEndpoint}`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(requestBody),
            });

            if (response.ok) {                
                const data = await response.json();
                login(data.token);
                const decodedToken = jwtDecode(data.token);
                if(decodedToken.role === 0)
                    navigate("/dashboard");
                else if(decodedToken.role === 1)
                    navigate("/authority");
                else if(decodedToken.role === 2)
                    navigate("/admin");
                else
                    navigate("/");

            } else {
                const errorData = await response.json();

                if (response.status === 400) {
                    switch (errorData.message) {
                        case "email already exists":
                            alert("Email already exists");
                            break;
                        case "username already exists":
                            alert("Username already exists");
                            break;
                        default:
                            if (action === t("login")) {
                                setShowWrongPasswordModal(true);
                            } else {
                                console.error("Registration error:", errorData);
                            }
                    }
                }
            }
        } catch (error) {
            console.log(error);
        }
    };


    return {
        email, setEmail,
        username, setUsername,
        password, setPassword,
        firstName, setFirstName,
        lastName, setLastName,
        action, setAction,
        showWrongPasswordModal, setShowWrongPasswordModal,
        isValidEmail, setIsValidEmail,
        isValidPassword, setIsValidPassword,
        showPasswordModal, setShowPasswordModal,
        handleShowPasswordModal,
        handleClosePasswordModal,
        handleEmailChange,
        handlePasswordChange,
        handleSubmit,
        isLoggedIn, setLoggedIn
    };

};

export default useLogin;