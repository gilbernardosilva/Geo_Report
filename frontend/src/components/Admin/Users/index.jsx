import React, { useState, useEffect } from "react";
import { Container, Row, Col, Table, Button, Modal, Form } from "react-bootstrap";
import { useAuth } from "../../../hooks/AuthContext.jsx";
import { useAxiosWithToken } from "../../../utils/api.js";
import CustomAdminNavbar from "../AdminNavbar/index.jsx";
import { useTranslation } from "react-i18next";
import UserModal from "./UserModal/index.jsx";
function UserManagement() {

    const { userInfo } = useAuth();
    const api = useAxiosWithToken();
    const [users, setUsers] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [selectedUser, setSelectedUser] = useState(null);
    const [editMode, setEditMode] = useState(false);
    const { t } = useTranslation();
    // Fetch users on component mount
    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await api.get("admin/user");
                if (response.status === 200) {
                    console.log(response.data);
                    setUsers(response.data.users);
                } else {
                    // Handle errors
                }
            } catch (error) {
                // Handle errors
            }
        };
        fetchUsers();
    }, [userInfo]);

    const handleAddUser = () => {
        setSelectedUser(null);
        setEditMode(false);
        setShowModal(true);
    };

    const handleEditUser = (user) => {
        setSelectedUser(user);
        setEditMode(true);
        setShowModal(true);
    };

    const handleDeleteUser = async (userId) => {
        if (window.confirm("Are you sure you want to delete this user?")) {
            try {
                await api.delete(`/api/v1/admin/user/delete/${userId}`);
                setUsers(users.filter((user) => user.id !== userId));
            } catch (error) {
                // Handle error
            }
        }
    };
    
    return (
        <>
            <CustomAdminNavbar t={t}></CustomAdminNavbar>
            <Container>
                <h2>{t('userManagement')}</h2>

                <Table striped bordered hover>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>{t('username')}</th>
                            <th>Email</th>
                            <th>{t('role')}</th>
                            <th>{t('actions')}</th>
                        </tr>
                    </thead>
                    <tbody>
                        {users.map((user) => (
                            <tr key={user.id}>
                                <td>{user.id}</td>
                                <td>{user.username}</td>
                                <td>{user.email}</td>
                                <td>{user.role.name}</td>
                                <td>
                                    <Button variant="outline-primary" onClick={() => handleEditUser(user)}>{t('edit')}</Button>
                                    <Button variant="outline-danger" className="ms-2" onClick={() => handleDeleteUser(user.id)}>
                                        {t('delete')}
                                    </Button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </Table>
                <UserModal key={selectedUser ? selectedUser.id : "new"}  showModal={showModal} setShowModal={setShowModal} editMode={editMode} existingUserData={selectedUser} setSelectedUser={setEditMode}></UserModal>
                <Button variant="primary" onClick={handleAddUser}>{t('addUser')}</Button>

            </Container>
        </>
    );
}

export default UserManagement;