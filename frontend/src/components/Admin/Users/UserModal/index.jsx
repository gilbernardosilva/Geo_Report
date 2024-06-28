import { Form, Button, Modal } from "react-bootstrap";
import { useState, useEffect } from "react";
import { useTranslation } from "react-i18next";
import generatePassword from "../../../../utils/passwordGenerator";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useAxiosWithToken } from "../../../../utils/api.js";

function UserModal({ showModal, setShowModal, editMode, existingUserData, setEditMode }) {
    const [formData, setFormData] = useState({
        username: "",
        first_name: "",
        last_name: "",
        role: 0,
        email: "",
        password: "",
    });
    console.log(existingUserData);
    const api = useAxiosWithToken();
    const { t } = useTranslation();
    useEffect(() => {
        debugger;
        if (editMode && existingUserData) {
            setFormData(existingUserData);
        }else {
            setFormData((prevFormData) => ({
                ...prevFormData,
                password: generatePassword(),
            }));
        }
    }, [editMode, existingUserData]);

    const handleChange = (event) => {
        setFormData({
            ...formData,
            [event.target.name]: event.target.value,
        });
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        try {
            if (editMode) {
                const response = await api.put(`user/edit/${existingUserData.id}`, formData);
                if (response.status === 200) {
                    toast.success("User updated successfully");
                } else {
                    toast.error("Error updating user:", response.data);
                }
            } else {
                const response = await api.post("/user/register", formData);
                if (response.status === 200) {
                    window.location.reload(true);
                    toast.success("User added successfully");
                } else {
                    toast.error("Error adding user:", response.data);
                }
            }

            setShowModal(false);
        } catch (error) {
            toast.error("An error occurred during submission:", error);
        }
    };

    const handleClose = () => {
        setEditMode(false);
        setShowModal(false);
        setFormData({ name: "", email: "", role: "", password: "" });
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
                            value={formData.firstname}
                            onChange={handleChange}
                        />
                    </Form.Group>

                    <Form.Group controlId="lastName">
                        <Form.Label>{t('lastName')}</Form.Label>
                        <Form.Control
                            type="text"
                            name="lastname"
                            placeholder={t('lastName')}
                            value={formData.lastname}
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
