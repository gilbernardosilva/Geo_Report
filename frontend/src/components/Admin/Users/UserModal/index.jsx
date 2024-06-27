import { Form, Button, Modal } from "react-bootstrap";
import { useState, useEffect } from "react";
import { useTranslation } from "react-i18next";
import generatePassword from "../../../../utils/passwordGenerator";

function UserModal({ showModal, setShowModal, editMode, existingUserData, setEditMode }) {
    const [formData, setFormData] = useState({ name: "", email: "", role: "" });

    const { t } = useTranslation();
    useEffect(() => {
        debugger;
        if (editMode && existingUserData) {
            setFormData(existingUserData);
        }
    }, [editMode, existingUserData]);

    const handleChange = (event) => {
        setFormData({
            ...formData,
            [event.target.name]: event.target.value,
        });
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        // Logic to send formData to your backend (add/update user)
        setShowModal(false); // Close the modal after submission
    };

    const handleClose = () => {
        setEditMode(false);
        setShowModal(false);
        setFormData({ name: "", email: "", role: "", password: "" }); // Reset form data
    };

    return (
        <Modal show={showModal} onHide={handleClose}>
            <Modal.Body>
                <Form onSubmit={handleSubmit}>
                    <Form.Group controlId="username">
                        <Form.Label>{t('username')}</Form.Label>
                        <Form.Control
                            type="text"
                            name="username"
                            placeholder={t('username')}
                            value={formData.username}
                            onChange={handleChange}
                        />
                    </Form.Group>
                    <Form.Group controlId="firstname">
                        <Form.Label>{t('firstName')}</Form.Label>
                        <Form.Control
                            type="text"
                            name="firstname"
                            placeholder={t('firstName')}
                            value={formData.first_name}
                            onChange={handleChange}
                        />
                    </Form.Group>

                    <Form.Group controlId="lastName">
                        <Form.Label>{t('lastName')}</Form.Label>
                        <Form.Control
                            type="text"
                            name="lastname"
                            placeholder={t('lastName')}
                            value={formData.last_name}
                            onChange={handleChange}
                        />
                    </Form.Group>


                    <Form.Group controlId="role">
                        <Form.Label>{t('role')}</Form.Label>

                        <Form.Select
                            value={formData.role.id}
                            onChange={(e) =>
                                setFormData({
                                    ...formData,
                                    role: parseInt(e.target.value, 10)
                                })
                            }>
                            <option value={0} selected={editMode && existingUserData.role.id === 0}>{t('user')}</option>
                            <option value={1} selected={editMode && existingUserData.role.id === 1}>{t('authority')}</option>
                            <option value={2} selected={editMode && existingUserData.role.id === 2}>{t('admin')}</option>
                        </Form.Select>
                    </Form.Group>
                    <Form.Group controlId="email">
                        <Form.Label>{t('email')}</Form.Label>
                        <Form.Control
                            type="email"
                            name="email"
                            placeholder={t('email')}
                            value={formData.email}
                            onChange={handleChange}
                        />
                    </Form.Group>
                    {!editMode && (
                        <Form.Group controlId="password">
                            <Form.Label>{t('password')}</Form.Label>
                            <Form.Control
                                disabled
                                type="text"
                                name="password"
                                placeholder={t('password')}
                                value={generatePassword()}
                                onChange={handleChange}
                            />
                        </Form.Group>
                    )
                    }


                    <Button variant="primary" type="submit">
                        {editMode ? t('saveChanges') : t('addUser')}
                    </Button>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

export default UserModal;
