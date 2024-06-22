import React, {  } from 'react';
import { Container } from 'react-bootstrap';
import './index.css';
import CustomNavbar from './../Navbar';
import { useTranslation } from "react-i18next";



function Profile() {

    const { t } = useTranslation();
    return (
        <>
            <CustomNavbar t={t}></CustomNavbar>
            <Container fluid className="dashboard-container">

            </Container>
        </>
    );
}

export default Profile;