import React, { useState } from 'react';
import { Modal, Button, Form } from 'react-bootstrap';
import { toast } from 'react-toastify'
import 'react-toastify/ReactToastify.css'

const ForgotPasswordModal = ({ show, handleClose }) => {
    const [email, setEmail] = useState('');
    const [message, setMessage] = useState('');

    const handleResetPassword = async () => {
        try {
            const response = await fetch(`${process.env.REACT_APP_API_URL}/user/forgotPassword`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({email}),
            });

            if (response.status === 200) {
                toast.success("E-mail successfully sent!")
                handleClose()
            } else {
                toast.error("Error sending e-mail.")
            }
        } catch (error) {
            setMessage(error.response.data.error);
        }
    };

    return (
        <Modal show={show} onHide={handleClose}>
            <Modal.Header closeButton>
                <Modal.Title>Forgot Password</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Group controlId="formEmail">
                        <Form.Label>Email address</Form.Label>
                        <Form.Control
                            type="email"
                            placeholder="Enter email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                        />
                    </Form.Group>
                </Form>
                {message && <p>{message}</p>}
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={handleClose}>
                    Close
                </Button>
                <Button variant="primary" onClick={handleResetPassword}>
                    Reset Password
                </Button>
            </Modal.Footer>
        </Modal>
    );
};

export default ForgotPasswordModal;
