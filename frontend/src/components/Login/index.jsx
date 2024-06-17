import { useState } from "react";
import user_icon from './../../img/icons/person.png'
import email_icon from './../../img/icons/email.png'
import password_icon from './../../img/icons/password.png'
import address_icon from './../../img/icons/address.svg'
import information_icon from './../../img/icons/information.svg'
import './index.css';
import logo from './../../img/logo/geo_report_white_text.svg'

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min";

import { Container, Row, Col, Modal, Button } from 'react-bootstrap';



const Login = ({ setToken }) => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [action, setAction] = useState("Sign Up");

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
                // ... (Handle successful login/signup: store token, redirect, etc.)
            } else if (response.status === 401) {
                // ... (Handle authentication error)
            } else {
                console.log(response);
                // ... (Handle other errors))
            }
        } catch (error) {
            // ... (Handle network or other errors)
        }
    };


    return (
        <Container fluid className="bg-black">
            <Row className="vh-100">
                <Col sm={12} md={6} className="d-flex align-items-center justify-content-center">
                    <img src={logo} alt="Your Logo" className="img-fluid" />
                </Col>
                <Col sm={12} md={6} className="d-flex align-items-center justify-content-center flex-column">
                    <div className="card login-signup-forgot" style={{
                        display: 'flex', justifyContent: 'center', maxWidth: '650px', minWidth: '400px',
                    }}>
                        <div className="header">
                            <div className="text">{action}</div>
                            <div className="underline"></div>
                        </div>
                        <form>
                            <div className="inputs">
                                {action === "Login" ? <div></div> :
                                    <><div className="input">
                                        <img className="svg" src={address_icon} alt="Last Name Icon" />
                                        <input type="text" className="form-control" value={firstName} onChange={(e) => setFirstName(e.target.value)} required />
                                        <label className="placeholderx">First Name</label>
                                    </div>
                                        <div className="input">
                                            <img className="svg" src={address_icon} alt="Last Name Icon" />
                                            <input type="text" className="form-control" value={lastName} onChange={(e) => setLastName(e.target.value)} required />
                                            <label className="placeholderx">Last Name</label>
                                        </div>
                                        <div className="input">
                                            <img src={email_icon} alt="E-mail Icon" />
                                            <input
                                                type="email"
                                                className={`form-control ${!isValidEmail ? "invalid-input" : ""}`}
                                                value={email}
                                                onChange={handleEmailChange}
                                                required
                                            />
                                            <label className="placeholderx">E-Mail</label>
                                        </div></>}
                                <div className="input">
                                    <img src={user_icon} alt="User Icon" />
                                    <input type="text" className="form-control" value={username} onChange={(e) => setUsername(e.target.value)} required />
                                    <label className="placeholderx">Username</label>
                                </div>

                                <div className="input">
                                    {!isValidPassword && action === "Sign Up" && (
                                        <img src={information_icon} alt="Information Icon" onClick={handleShowPasswordModal} style={{ cursor: 'pointer' }} />
                                    )}
                                    <img src={password_icon} alt="Password Icon" />
                                    <input
                                        type="password"
                                        className={`form-control ${!isValidPassword && action === "Sign Up" ? "invalid-input" : ""}`}
                                        value={password}
                                        onChange={handlePasswordChange}
                                        required
                                    />
                                    <label className="placeholderx">Password</label>
                                </div>
                            </div>
                        </form>
                        {action === "Sign Up" ? <div></div> :
                            <div className="forgot-password"> Forgot Password? <span>Click Here!</span></div>}
                        <div className="submit-container">
                            <div className={action === "Login" ? "submit gray" : "submit"} onClick={action === "Sign Up" ? handleSubmit : () => { setAction("Sign Up") }}>Sign Up</div>
                            <div className={action === "Sign Up" ? "submit gray" : "submit"} onClick={action === "Login" ? handleSubmit : () => { setAction("Login") }} >Login</div>
                        </div>
                    </div>
                </Col>
            </Row>
            <Modal show={showPasswordModal} onHide={handleClosePasswordModal}>
                <Modal.Header closeButton>
                    <Modal.Title>Password Requirements</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <p>Password must be at least 8 characters long and contain:</p>
                    <ul>
                        <li>At least one lowercase letter</li>
                        <li>At least one uppercase letter</li>
                        <li>At least one number</li>
                        <li>At least one special character (@$!%*?&)</li>
                    </ul>
                </Modal.Body>
            </Modal>
        </Container>
    );

};

export default Login;