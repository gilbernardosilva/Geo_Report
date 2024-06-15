import { useState } from "react";
import user_icon from './../../img/icons/person.png'
import email_icon from './../../img/icons/email.png'
import password_icon from './../../img/icons/password.png'
import address_icon from './../../img/icons/address.svg'
import './index.css';
import logo from './../../img/logo/geo_report_white_text.svg'

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min";

import { Container, Row, Col, Form, Button } from 'react-bootstrap';



const Login = ({ setToken }) => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [ok, setOk] = useState("");

    const [action, setAction] = useState("Sign Up");

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch("http://localhost:3000/auth/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ email, password }),
            });
            if (response.ok) {
                const data = await response.json();
                sessionStorage.setItem("token", data.token);
                sessionStorage.setItem("email", data.email);
                sessionStorage.setItem("object", JSON.stringify(data));
                setOk(true);
                setToken(true);
                setError("");
            } else if (response.status === 401) {
                sessionStorage.removeItem("token");
                setError("Incorrect email or password.");
                setOk(false);
            } else {
                sessionStorage.removeItem("token");
                setError("An error occurred. Please try again.");
                setOk(false);
            }
        } catch (error) {
            sessionStorage.removeItem("token");

            setError("An error occurred. Please try again.");
            setOk(false);
        }
    };


    return (
        <Container fluid className="vh-1000 d-flex align-items-center bg-black">
            <Row className="vh-1000">
                <Col md={6} className="d-flex align-items-center justify-content-center">
                    <img src={logo} alt="Your Logo" className="img-fluid" style={{ maxWidth: '500px' }} /> {/* Add maxWidth style */}
                </Col>
                <Col md={6} className="d-flex align-items-center justify-content-center flex-column">
                    <div className="card login-signup-forgot" style={{display: 'flex', justifyContent: 'center', maxWidth: '650px', minWidth: '400px', 
                        marginTop: '5vh', marginBottom: '10vh', marginLeft: '5vh', marginRight: '5vh'}}>
                        <div className="header">
                            <div className="text">{action}</div>
                            <div className="underline"></div>
                        </div>
                        <form>
                            <div className="inputs">    
                            {action==="Login"?<div></div>:
                                <><div className="input">
                                    <img className="svg" src={address_icon} alt="Last Name Icon" />
                                    <input type="text" className="form-control" required />
                                    <label className="placeholderx">First Name</label>
                                </div>
                                <div className="input">
                                    <img className="svg" src={address_icon} alt="Last Name Icon" />
                                    <input type="text" className="form-control" required />
                                    <label className="placeholderx">Last Name</label>
                                </div>
                                <div className="input">
                                    <img src={email_icon} alt="E-mail Icon" />
                                    <input type="text" className="form-control" required />
                                    <label className="placeholderx">E-Mail</label>
                                </div></>}
                            <div className="input">
                                <img src={user_icon} alt="User Icon"/>
                                <input type="text" className="form-control" required />
                                <label className="placeholderx">Username</label>
                            </div>
                            <div className="input">
                                <img src={password_icon} alt="Password Icon" />
                                <input type="password" className="form-control" required />
                                <label className="placeholderx">Password</label>

                            </div>
                        </div>
                        </form>
                        {action === "Sign Up" ? <div></div> :
                        <div className="forgot-password"> Forgot Password? <span>Click Here!</span></div>}
                        <div className="submit-container">
                            <div className={action==="Login"?"submit gray":"submit"} onClick={action === "Sign Up" ? handleSubmit : ()=>{setAction("Sign Up")}}>Sign Up</div>
                            <div className={action === "Sign Up" ? "submit gray" : "submit"} onClick={action === "Login" ? handleSubmit : ()=>{setAction("Login")}} >Login</div>
                        </div>
                    </div>
                </Col>
            </Row>
        </Container>
    );

};

export default Login;