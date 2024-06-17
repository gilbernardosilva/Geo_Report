import React, { useState } from 'react';
import './index.css';
import logo from './../../img/logo/geo_report_no_text.png'
import { Container, Row, Col, Modal, Button } from 'react-bootstrap';
import home from './../../img/icons/home.svg'
import issues from './../../img/icons/issues.svg'
import logout from './../../img/icons/logout.svg'
import settings from './../../img/icons/settings.svg'
import { Link } from 'react-router-dom';

function Dashboard() {

    return (
        <Container fluid>
            <Row className="justify-content-center mt-4">
                <Col xs={12} md={3} className="d-flex justify-content-end" style={{width: 300}}>
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={logo} alt="Your Logo" className="img-fluid logo-icon" />
                    </Link>
                </Col>
                <Col xs={12} md={4} className="d-flex justify-content-center">
                </Col>
                <Col xs={12} md={3} className="d-flex justify-content-end">
                </Col>
            </Row>

            <Row className="justify-content-center">

                <Col xs={12} md={3} className="sidebar order-md-1 d-flex">
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={home} alt="" className="img-fluid logo-icon me-2" />
                        <h2>Home</h2>
                    </Link>
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={issues} alt="" className="img-fluid logo-icon me-2" />
                        <h2>My issues</h2>
                    </Link>
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={settings} alt="" className="img-fluid logo-icon me-2" />
                        <h2>Settings</h2>
                    </Link>
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={logout} alt="" className="img-fluid logo-icon me-2" />
                        <h2>Logout</h2>
                    </Link>
                </Col>
                

                <Col xs={12} md={4} className="map-area order-md-2"></Col>

                <Col xs={12} md={3} className="info-sidebar order-md-3">
                </Col>
            </Row>
        </Container>
    );
}

export default Dashboard;