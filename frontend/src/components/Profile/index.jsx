import React, {  useState } from 'react';
import { Container, Row,Col,Form,Button } from 'react-bootstrap';
import './index.css';
import CustomNavbar from './../Navbar';
import { useTranslation } from "react-i18next";
import { useAuth } from "../../hooks/AuthContext.jsx";
import { useAxiosWithToken} from "../../utils/api.js"



function Profile() {
    const { userInfo } = useAuth();
    const api = useAxiosWithToken();
    const [formData, setFormData] = useState({
      firstName: userInfo?.firstName || "",
      lastName: userInfo?.lastName || "",
      email: userInfo?.email || "",
      username: userInfo?.username || "",
    });
  
    const [error, setError] = useState(null);

    const handleChange = (event) => {
      setFormData({
        ...formData,
        [event.target.id]: event.target.value,
      });
    };
  
    const handleSubmit = async (event) => {
      event.preventDefault();
      try {
        const response = await api.put("/api/v1/profile", formData);
        if (!response.ok) {
          throw new Error("Failed to update profile");
        } else {
          console.log("Profile updated successfully!");
        }
      } catch (error) {
        setError(error.message);
      }
    };
    const { t } = useTranslation();
    return (
        <>
            <CustomNavbar t={t}></CustomNavbar>
            <Container fluid className="dashboard-container">
            <Row className="justify-content-center">
        <Col md={6}>
          <h1>{t("editProfile")}</h1>
          {error && <div className="alert alert-danger">{error}</div>}
          <Form onSubmit={handleSubmit}>
            <Form.Group controlId="firstName" className="mb-3">
              <Form.Label>{t("firstName")}</Form.Label>
              <Form.Control
                type="text"
                value={formData.firstName}
                onChange={handleChange}
              />
            </Form.Group>

            <Form.Group controlId="lastName" className="mb-3">
              <Form.Label>{t("lastName")}</Form.Label>
              <Form.Control
                type="text"
                value={formData.lastName}
                onChange={handleChange}
              />
            </Form.Group>

            <Form.Group controlId="email" className="mb-3">
              <Form.Label>{t("email")}</Form.Label>
              <Form.Control
                type="email"
                value={formData.email}
                onChange={handleChange}
              />
            </Form.Group>

            <Form.Group controlId="username" className="mb-3">
              <Form.Label>{t("username")}</Form.Label>
              <Form.Control disabled
                type="text"
                value={formData.username}
                onChange={handleChange}
              />
            </Form.Group>


            <Button variant="primary" type="submit">
              {t("saveChanges")}
            </Button>
          </Form>
        </Col>
      </Row>
            </Container>
        </>
    );
}

export default Profile;