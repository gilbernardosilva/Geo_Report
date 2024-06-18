import React, { } from 'react';
import './index.css';
import logo from './../../img/logo/geo_report_no_text.png'
import { Container, Row, Col } from 'react-bootstrap';
import home from './../../img/icons/home.svg'
import issues from './../../img/icons/issues.svg'
import logout from './../../img/icons/logout.svg'
import profile from './../../img/icons/profile.svg'
import settings from './../../img/icons/settings.svg'
import { Link } from 'react-router-dom';
import { MapContainer, TileLayer, useMap, Marker, Popup } from 'react-leaflet'
import "leaflet/dist/leaflet.css";
import useLogout from './../../hooks/useLogout.js'
import { useTranslation } from 'react-i18next';

function Dashboard({ setToken }) {

    const { handleLogout } = useLogout(setToken);
    const { t } = useTranslation();

    return (
        <Container fluid className="dashboard-container">
            <Row className="justify-content-center mt-4 dashboard-row">
                <Col xs={12} md={3} className="d-flex justify-content-end" style={{ width: 300 }}>
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={logo} alt="Your Logo" className="img-fluid logo-icon" />
                    </Link>
                </Col>
                <Col xs={12} md={4} className="d-flex justify-content-center">
                    <h2>REPORT AN ISSUE!</h2>
                </Col>
                <Col xs={12} md={3} className="d-flex justify-content-end">
                </Col>
            </Row>

            <Row className="justify-content-center dashboard-row flex-grow-1" >

                <Col xs={12} md={3} className="sidebar order-md-1 d-flex">
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={home} alt="" className="img-fluid logo-icon me-2" />
                        <h2>{t('home')}</h2>
                    </Link>
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={issues} alt="" className="img-fluid logo-icon me-2" />
                        <h2>{t('myIssues')}</h2>
                    </Link>

                    <div className="mt-auto"></div>
                    <Link to="/profile" className="sidebar-header text-end btn">
                        <img src={profile} alt="" className="img-fluid logo-icon me-2" />
                        <h2>{t('profile')}</h2>
                    </Link>
                    <Link to="/dashboard" className="sidebar-header text-end btn">
                        <img src={settings} alt="" className="img-fluid logo-icon me-2" />
                        <h2>{t('settings')}</h2>
                    </Link>
                    <Link to="/" className="sidebar-header text-end btn">
                        <img src={logout} alt="" className="img-fluid logo-icon me-2" />
                        <h2 onClick={handleLogout}>{t('logout')}</h2>
                    </Link>
                </Col>


                <Col xs={12} md={4} className="map-area order-md-2 d-flex">
                        <MapContainer
                            center={[51.505, -0.09]}
                            zoom={13}
                            scrollWheelZoom={false}
                            style={{ height: "100%", width: "100%" }}>
                            <TileLayer
                                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                            />
                            <Marker position={[51.505, -0.09]}>
                                <Popup>
                                    A pretty CSS3 popup. <br /> Easily customizable.
                                </Popup>
                            </Marker>
                        </MapContainer>
                </Col>

                <Col xs={12} md={3} className="info-sidebar order-md-3">
                </Col>
            </Row>
        </Container>
    );
}

export default Dashboard;