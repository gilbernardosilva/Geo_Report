import { useState } from "react";
import bcrypt from "bcryptjs";
import { useNavigate } from 'react-router-dom';

function useLogin(setToken) {
    const navigate = useNavigate(); 
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [action, setAction] = useState("Sign Up");
    const [showWrongPasswordModal, setShowWrongPasswordModal] = useState(false);

    const [isValidEmail, setIsValidEmail] = useState(true);

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

    console.log(process.env.REACT_APP_API_URL);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const apiEndpoint = action === "Login" ? "login" : "register";
        const requestBody = action === "Login"
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
                setToken(data.Token);
                navigate('/dashboard'); 
            } else if (response.status === 400) {
               if (action === "Login") {
                   setShowWrongPasswordModal(true); 
               }
            } else {
                console.log(response);
            }
        } catch (error) {
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
      };

};

export default useLogin;