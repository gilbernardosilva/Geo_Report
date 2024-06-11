import { useState } from "react";
import './index.css';

import { Container, Row, Col, Form, Button } from 'react-bootstrap';



const Login = ({ setToken }) => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [ok, setOk] = useState("");

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
        <Container fluid className="vh-100 d-flex align-items-center bg-black">
            <Row className="w-100">
                <Col md={6} className="d-flex align-items-center justify-content-center">
                    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/57/X_logo_2023_%28white%29.png/800px-X_logo_2023_%28white%29.png" alt="Your Logo" className="img-fluid" style={{ maxWidth: '400px' }} /> {/* Add maxWidth style */}
                </Col>
                <Col
                    md={6}
                    className="d-flex align-items-left justify-content-left flex-column">
                    <div>
                        <h2 className="fw-bold text-white" style={{ fontSize: '64px' }}>Helping others</h2>
                    </div>
                    <div>
                        <p className="fw-bold text-white mb-4" style={{ fontSize: '31px', marginTop: '30px' }}>Sign up today</p>
                    </div>
                    <Form onSubmit={handleSubmit}>
                        {/* ... your form fields (email, password, etc.) */}
                    </Form>
                    <Button variant="link" className="text-black">
                        Already got an account?
                    </Button>
                </Col>
            </Row>
        </Container>
    );

};

export default Login;