import React, { useState, useEffect } from 'react';
import { Container, Row, Col, Form, Button, Alert, Card, CardHeader, CardBody, CardFooter } from 'react-bootstrap';
import './index.css';
import CustomNavbar from './../Navbar';
import { useTranslation } from "react-i18next";
import { useAuth } from "../../hooks/AuthContext.jsx";
import { useAxiosWithToken } from "../../utils/api.js"
import CustomAuthorityNavbar from '../AuthorityNavbar/index.jsx';



function Profile() {
  const { userInfo, login } = useAuth();
  console.log(userInfo);
  const api = useAxiosWithToken();
  const [formData, setFormData] = useState({
    userid: userInfo?.id.toString() || "",
    username: userInfo?.username || "",
    email: userInfo?.email || "",
    firstName: userInfo?.firstName || "",
    lastName: userInfo?.lastName || "",
  });
  const [success, setSuccess] = useState(null);
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
      const response = await api.put(`/user/edit/${userInfo?.id}`, formData);
      if (response.status === 200) {
        setSuccess("Profile updated successfully!");
        login(response.data.token)
        setError(null);
      } else {
        throw new Error("Failed to update profile");
      }
    } catch (error) {
      setError(error.message);
    }
  };


  useEffect(() => {
    setFormData({
      userid: userInfo?.id.toString() || "",
      username: userInfo?.username || "",
      email: userInfo?.email || "",
      firstName: userInfo?.firstName || "",
      lastName: userInfo?.lastName || "",
    });
  }, [userInfo]);

  const { t } = useTranslation();
  return (
    <>
      {userInfo?.role === 0 ? (
        <CustomNavbar t={t} />
      ) : userInfo?.role === 1 ? (
        <CustomAuthorityNavbar t={t} />
      ) : null}
      <Container fluid className="dashboard-container">
        <Row className="justify-content-center">
          <Col md={6}>
            <Card>
              <CardHeader>
                <h1>{t("editProfile")}</h1>
                {error && <Alert variant="danger">{error}</Alert>}
                {success && (
                  <Alert variant="success" onClose={() => setSuccess(null)} dismissible>
                    {success}
                  </Alert>
                )}
              </CardHeader>

              <Form onSubmit={handleSubmit}>
                <CardBody>
                  <Form.Group controlId="userid" className="mb-3">
                    <Form.Control hidden
                      type="text"
                      value={formData.userid}
                      onChange={handleChange}
                    />
                  </Form.Group>
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


                </CardBody>
                <CardFooter>
                  <Button variant="primary" type="submit">
                    {t("saveChanges")}
                  </Button>
                </CardFooter>

              </Form>

            </Card>


          </Col>
        </Row>
      </Container>
    </>
  );
}

export default Profile;